package flagenv

import (
	"github.com/segmentio/go-camelcase"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

type Flag struct {
	flagSet     *pflag.FlagSet
	normalizers []func(string) string
	envName     string
	name        string
	shorthand   string
	value       string
	usage       string
}

func (f *Flag) FlagSet() *pflag.FlagSet {
	return f.flagSet
}

func (f *Flag) EnvName() string {
	return f.envName
}

func (f *Flag) Name() string {
	return f.name
}

func (f *Flag) Shorthand() string {
	return f.shorthand
}

func (f *Flag) Value() string {
	return f.value
}

func (f *Flag) Usage() string {
	return f.usage
}

func (f *Flag) AddNormalizer(normalizers ...func(string) string) *Flag {
	f.normalizers = append(f.normalizers, normalizers...)
	return f
}

func (f *Flag) SetEnvName(envName string) *Flag {
	f.envName = envName
	return f
}

func (f *Flag) SetName(name string) *Flag {
	f.name = name

	for _, normalizer := range f.normalizers {
		f.name = normalizer(f.Name())
	}

	f.name = camelcase.Camelcase(f.Name())
	return f
}

func (f *Flag) SetShorthand(shorthand string) *Flag {
	switch shorthand {
	case "-":
		f.shorthand = f.Name()[0:1]
	case "-^":
		f.shorthand = strings.ToUpper(f.Name()[0:1])
	default:
		f.shorthand = shorthand
	}
	return f
}

func (f *Flag) SetValue(value string) *Flag {
	f.value = value
	return f
}

func (f *Flag) SetUsage(usage string) *Flag {
	f.usage = usage
	return f
}

func (f *Flag) SetCLIFlag() *Flag {
	f.flagSet.StringP(
		f.Name(),
		f.Shorthand(),
		f.Value(),
		f.Usage(),
	)
	return f
}

func (f *Flag) SetEnv() *Flag {
	cobra.CheckErr(viper.BindPFlag(
		f.EnvName(),
		f.flagSet.Lookup(f.Name()),
	))
	return f
}

func NewFlag(flagSet *pflag.FlagSet) *Flag {
	return &Flag{
		flagSet: flagSet,
	}
}

// New creates flag, add it to the flagSet and bind environment variable.
func New(
	flagSet *pflag.FlagSet,
	envName, flagShort, value, usage string,
	flagNormalizers ...func(string) string,
) {
	f := NewFlag(flagSet)
	f.AddNormalizer(flagNormalizers...)
	f.SetEnvName(envName)
	f.SetName(envName)
	f.SetShorthand(flagShort)
	f.SetValue(value)
	f.SetUsage(usage)
	f.SetCLIFlag()
	f.SetEnv()
}
