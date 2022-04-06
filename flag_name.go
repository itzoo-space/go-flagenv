package flagenv

import "github.com/segmentio/go-camelcase"

func (fe *FlagEnv) FlagName() string {
	return fe.flagName
}

func (fe *FlagEnv) SetFlagName(flagName string) *FlagEnv {
	fe.flagName = flagName

	for _, normalizer := range fe.Normalizers() {
		fe.flagName = normalizer(fe.FlagName())
	}

	fe.flagName = camelcase.Camelcase(fe.FlagName())
	reCamelCased := camelcase.Camelcase(fe.FlagName())

	for fe.FlagName() != reCamelCased {
		fe.flagName = reCamelCased
		reCamelCased = camelcase.Camelcase(fe.FlagName())
	}

	return fe
}

func WithFlagName(flagName string) Modifier {
	return func(fe *FlagEnv) {
		fe.SetFlagName(flagName)
	}
}
