package flagenv

import (
	"errors"
	"github.com/segmentio/go-camelcase"
	"github.com/spf13/cobra"
)

func (fe *FlagEnv) FlagName() string {
	flagName := fe.flagName

	if flagName == "" {
		if fe.EnvName() == "" {
			cobra.CheckErr(errors.New("flag name and env name aren't set"))
		}
		flagName = fe.EnvName()
	}

	for _, normalizer := range fe.Normalizers() {
		flagName = normalizer(flagName)
	}

	camelCased := camelcase.Camelcase(flagName)
	for flagName != camelCased {
		flagName = camelCased
		camelCased = camelcase.Camelcase(flagName)
	}

	return flagName
}

func (fe *FlagEnv) SetFlagName(flagName string) *FlagEnv {
	fe.flagName = flagName

	return fe
}

func WithFlagName(flagName string) Modifier {
	return func(fe *FlagEnv) {
		fe.SetFlagName(flagName)
	}
}
