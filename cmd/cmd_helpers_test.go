package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func AssertCommandProperties(t *testing.T, cmd *cobra.Command, use string, aliases []string, short string, long string) {
	if cmd.Use != use {
		t.Errorf("Expected use to be '%s', got '%s'", use, cmd.Use)
	}

	if len(cmd.Aliases) != len(aliases) {
		t.Errorf("Expected %d aliases, got %d", len(aliases), len(cmd.Aliases))
	} else {
		for i, alias := range aliases {
			if cmd.Aliases[i] != alias {
				t.Errorf("Expected alias %d to be '%s', got '%s'", i, alias, cmd.Aliases[i])
			}
		}
	}

	if cmd.Short != short {
		t.Errorf("Expected short description to be '%s', got '%s'", short, cmd.Short)
	}

	if cmd.Long != long {
		t.Errorf("Expected long description to be '%s', got '%s'", long, cmd.Long)
	}
}

func AssertFlagProperties(t *testing.T, cmd *cobra.Command, flagName string, shorthand string, defValue string, usage string) {
	flag := cmd.Flag(flagName)
	if flag == nil {
		t.Errorf("Expected flag '%s' to exist", flagName)
	} else {
		if flag.Name != flagName {
			t.Errorf("Expected flag name to be '%s', got '%s'", flagName, flag.Name)
		}

		if flag.Shorthand != shorthand {
			t.Errorf("Expected flag shorthand to be '%s', got '%s'", shorthand, flag.Shorthand)
		}

		if flag.DefValue != defValue {
			t.Errorf("Expected flag default value to be '%s', got '%s'", defValue, flag.DefValue)
		}

		if flag.Usage != usage {
			t.Errorf("Expected flag usage to be '%s', got '%s'", usage, flag.Usage)
		}
	}
}