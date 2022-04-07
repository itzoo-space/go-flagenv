package flagenv

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Modifier func(*FlagEnv)
type ValueSetter func(*FlagEnv)

type FlagEnv struct {
	flagSet *pflag.FlagSet

	envName     string
	normalizers []func(string) string
	flagName    string
	shorthand   string

	usage string
}

func (fe *FlagEnv) FlagSet() *pflag.FlagSet {
	return fe.flagSet
}

func NewFlagEnv(flagSet *pflag.FlagSet) *FlagEnv {
	return &FlagEnv{flagSet: flagSet}
}

func New(
	flagSet *pflag.FlagSet,
	value ValueSetter,
	opts ...Modifier,
) {
	fe := NewFlagEnv(flagSet)

	for _, modifier := range opts {
		modifier(fe)
	}

	if fe.FlagSet().Lookup(fe.FlagName()) != nil {
		cobra.CheckErr(errors.New("flag already set"))
	}

	value(fe)

	if fe.EnvName() != "" {
		cobra.CheckErr(viper.BindPFlag(
			fe.EnvName(),
			fe.FlagSet().Lookup(fe.FlagName()),
		))
	}
}
