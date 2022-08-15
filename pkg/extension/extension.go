package extension

import (
	"bytes"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/robertkrimen/otto"
	"github.com/spf13/cobra"
)

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
	ext.VM.Set("cuc", func(call otto.FunctionCall) otto.Value {
		args := make([]string, len(call.ArgumentList))
		for i, v := range call.ArgumentList {
			args[i] = v.String()
		}
		ext.Cmd.SetArgs(args)

		outBuf := new(bytes.Buffer)
		errBuf := new(bytes.Buffer)

		if err := ext.Cmd.Execute(); err != nil {
			fmt.Println("error", err.Error())
			return otto.Value{}
		}

		fmt.Println("0............")
		fmt.Println(outBuf.String())
		fmt.Println("1............")
		fmt.Println(errBuf.String())
		fmt.Println("2............")

		return otto.Value{}
	})

	return nil
}

func (ext Extension) Run(name string, args []string) error {
	_, err := ext.VM.Run(heredoc.Doc(`
	console.log('============================');
	cuc();
	console.log('============================');
	cuc('team');
	console.log('============================');
`))
	if err != nil {
		return err
	}

	return nil
}
