package flagenv

import "strings"

func (fe *FlagEnv) Shorthand() (shorthand string) {
	switch fe.shorthand {
	case "-":
		shorthand = strings.ToLower(fe.FlagName()[0:1])
	case "-^":
		shorthand = strings.ToUpper(fe.FlagName()[0:1])
	default:
		shorthand = fe.shorthand
	}
	return shorthand
}

func (fe *FlagEnv) SetShorthand(shorthand string) *FlagEnv {
	fe.shorthand = shorthand
	return fe
}

func WithShorthand(shorthand string) Modifier {
	return func(fe *FlagEnv) {
		fe.SetShorthand(shorthand)
	}
}
