package extension

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/robertkrimen/otto"
	"github.com/spf13/cobra"
)

var scriptsDir = defaultScriptsDir()

func defaultScriptsDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Sprint("faild to get user config dir: ", err.Error()))
	}

	return filepath.Join(dir, "cuc", "scripts")
}

type Extension struct {
	VM  *otto.Otto
	Cmd *cobra.Command
}

func NewExtension(cmd *cobra.Command) Extension {
	return Extension{
		VM:  otto.New(),
		Cmd: cmd,
	}
}

func (ext Extension) Setup() error {

	throw := func(str string) {
		panic(ext.VM.MakeCustomError("CucExtensionError", str))
	}

	ext.VM.Set("cuc", func(call otto.FunctionCall) otto.Value {
		args := make([]string, len(call.ArgumentList))
		for i, v := range call.ArgumentList {
			args[i] = v.String()
		}
		ext.Cmd.SetArgs(args)

		outBuf := new(bytes.Buffer)
		errBuf := new(bytes.Buffer)
		ext.setOutput(outBuf, errBuf)

		if err := ext.Cmd.Execute(); err != nil {
			throw(err.Error() + "\n" + errBuf.String())
		}

		str := outBuf.String()
		if json.Valid([]byte(str)) {
			obj, err := ext.VM.Object(str)
			if err != nil {
				throw(err.Error())
			}

			return obj.Value()
		}

		val, err := ext.VM.ToValue(str)
		if err != nil {
			throw(err.Error())
		}

		return val
	})

	return nil
}

func (ext Extension) setOutput(out, errOut io.Writer) {
	ext.Cmd.SetOut(out)
	ext.Cmd.SetErr(errOut)

	for _, cmd := range ext.Cmd.Commands() {
		ext := Extension{Cmd: cmd}
		ext.setOutput(out, errOut)
	}
}

func (ext Extension) Run(name string, args []string) error {
	f, err := lookup(name)
	if err != nil {
		return err
	}

	if _, err := ext.VM.Run(f); err != nil {
		return err
	}

	return nil
}

func lookup(name string) (*os.File, error) {
	f, err := os.Open(filepath.Join(scriptsDir, name+".js"))
	if err != nil {
		return nil, err
	}

	return f, nil
}

type ScriptInfo struct {
	Name      string
	Path      string
	UpdatedAt time.Time
	Size      int64
}

func List() ([]ScriptInfo, error) {
	d, err := os.Open(scriptsDir)
	if err != nil {
		return nil, err
	}

	entries, err := d.ReadDir(0)
	if err != nil {
		return nil, err
	}

	res := make([]ScriptInfo, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileExt := filepath.Ext(entry.Name())
		if fileExt != ".js" {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		res = append(res, ScriptInfo{
			Name:      strings.TrimSuffix(entry.Name(), fileExt),
			Path:      filepath.Join(scriptsDir, entry.Name()),
			UpdatedAt: info.ModTime(),
			Size:      info.Size(),
		})
	}

	return res, nil
}
