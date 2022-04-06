package flagenv

import (
	"github.com/spf13/pflag"
)

type Modifier func(*FlagEnv)

type FlagEnv struct {
	flags *pflag.FlagSet

	flagIsSet bool

	normalizers []func(string) string
	flagName    string
	shorthand   string
	envName     string
	usage       string
}

func New(flagSet *pflag.FlagSet, opts ...Modifier) *FlagEnv {
	fe := &FlagEnv{
		flags: flagSet,
	}

	for _, modifier := range opts {
		modifier(fe)
	}

	return fe
}
