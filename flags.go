package flagenv

import "github.com/spf13/pflag"

func (fe *FlagEnv) Flags() *pflag.FlagSet {
	return fe.flags
}
