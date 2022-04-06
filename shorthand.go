package flagenv

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

func (fe *FlagEnv) Shorthand() string {
	return fe.shorthand
}

func (fe *FlagEnv) SetShorthand(shorthand string) *FlagEnv {
	switch shorthand {
	case "-":
		fe.shorthand = fe.FlagName()[0:1]
	case "-^":
		fe.shorthand = strings.ToUpper(fe.FlagName()[0:1])
	default:
		fe.shorthand = shorthand
	}
	return fe
}

func WithShorthand(shorthand string) Modifier {
	return func(fe *FlagEnv) {
		if fe.FlagName() == "" {
			cobra.CheckErr(errors.New("flag name isn't set"))
		}
		fe.SetShorthand(shorthand)
	}
}
