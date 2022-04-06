package flagenv

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (fe *FlagEnv) FlagIsSet() bool {
	return fe.flagIsSet
}

func WithValue(modifier Modifier) Modifier {
	return func(fe *FlagEnv) {
		if fe.FlagName() == "" {
			if fe.EnvName() == "" {
				cobra.CheckErr(errors.New("flag name and env name aren't set"))
			}
			fe.SetFlagName(fe.EnvName())
		}

		if fe.FlagIsSet() {
			cobra.CheckErr(errors.New("flag already set"))
		}

		modifier(fe)
		fe.flagIsSet = true

		if fe.EnvName() != "" {
			cobra.CheckErr(viper.BindPFlag(
				fe.EnvName(),
				fe.Flags().Lookup(fe.FlagName()),
			))
		}
	}
}

func getValue[T any](valuePool []T) (value T, err error) {
	switch len(valuePool) {
	case 0:
		return
	case 1:
		value = valuePool[0]
		return
	default:
		err = errors.New("wrong number of parameters for value")
		return
	}
}

func WithBoolValue(optValue ...bool) Modifier {
	value, err := getValue(optValue)
	cobra.CheckErr(err)

	return WithValue(
		func(fe *FlagEnv) {
			fe.Flags().BoolP(
				fe.FlagName(),
				fe.Shorthand(),
				value,
				fe.Usage(),
			)
		},
	)
}

func WithStringValue(optValue ...string) Modifier {
	value, err := getValue(optValue)
	cobra.CheckErr(err)

	return WithValue(
		func(fe *FlagEnv) {
			fe.Flags().StringP(
				fe.FlagName(),
				fe.Shorthand(),
				value,
				fe.Usage(),
			)
		},
	)
}

func WithIntValue(optValue ...int) Modifier {
	value, err := getValue(optValue)
	cobra.CheckErr(err)

	return WithValue(
		func(fe *FlagEnv) {
			fe.Flags().IntP(
				fe.FlagName(),
				fe.Shorthand(),
				value,
				fe.Usage(),
			)
		},
	)
}
