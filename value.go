package flagenv

import (
	"errors"
	"github.com/spf13/cobra"
)

func getValue[T any](valuePool []T) (value T) {
	switch len(valuePool) {
	case 0:
		return
	case 1:
		value = valuePool[0]
		return
	default:
		cobra.CheckErr(errors.New("wrong number of parameters for value"))
		return
	}
}

func Bool(optValue ...bool) ValueSetter {
	return func(fe *FlagEnv) {
		fe.FlagSet().BoolP(
			fe.FlagName(),
			fe.Shorthand(),
			getValue(optValue),
			fe.Usage(),
		)
	}

}

func String(optValue ...string) ValueSetter {
	return func(fe *FlagEnv) {
		fe.FlagSet().StringP(
			fe.FlagName(),
			fe.Shorthand(),
			getValue(optValue),
			fe.Usage(),
		)
	}
}

func Int(optValue ...int) ValueSetter {
	return func(fe *FlagEnv) {
		fe.FlagSet().IntP(
			fe.FlagName(),
			fe.Shorthand(),
			getValue(optValue),
			fe.Usage(),
		)
	}
}
