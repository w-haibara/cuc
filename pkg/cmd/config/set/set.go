package set

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/w-haibara/cuc/pkg/config"
	"github.com/w-haibara/cuc/pkg/util"
)

type SetOptions struct {
	Key   string
	Value string
}

func NewCmdConfigSet(opts SetOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Update configuration with a value for the given key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Key = args[0]
			opts.Value = args[1]

			return setRun(opts, cmd.OutOrStdout(), cmd.OutOrStderr(), util.JsonFlag(cmd))
		},
	}

	return cmd
}

func setRun(opts SetOptions, out, errOut io.Writer, jsonFlag bool) error {
	if !jsonFlag {
		if err := ValidateKey(opts.Key); err != nil {
			warningIcon := "⚠️" //iostreams.IO.ColorScheme().WarningIcon()
			fmt.Fprintf(errOut, "%s warning: '%s' is not a known configuration key\n", warningIcon, opts.Key)
		}
	}

	if err := ValidateValue(opts.Key, opts.Value); err != nil {
		var invalidValue InvalidValueError
		if !errors.As(err, &invalidValue) {
			return fmt.Errorf("unknown error: %s", err.Error())
		}

		var values []string
		for _, v := range invalidValue.ValidValues {
			values = append(values, fmt.Sprintf("'%s'", v))
		}

		return fmt.Errorf("failed to set %q to %q: valid values are %v", opts.Key, opts.Value, strings.Join(values, ", "))
	}

	config.Set(opts.Key, opts.Value)
	config.Write()

	return nil
}

func ValidateKey(key string) error {
	for _, cfg := range config.Configs() {
		if key == cfg.Key {
			return nil
		}
	}

	return fmt.Errorf("invalid key")
}

type InvalidValueError struct {
	ValidValues []string
}

func (e InvalidValueError) Error() string {
	return "invalid value"
}

func ValidateValue(key, value string) error {
	var validValues []string

	for _, cfg := range config.Configs() {
		if cfg.Key == key {
			validValues = cfg.AllowedValues
			break
		}
	}

	if validValues == nil {
		return nil
	}

	for _, v := range validValues {
		if v == value {
			return nil
		}
	}

	return InvalidValueError{ValidValues: validValues}
}
