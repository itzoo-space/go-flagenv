package flagenv_test

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/mazen160/go-random"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"testing"
)

//import (
//	"github.com/itzoo-space/go-flagenv"
//	"github.com/spf13/cobra"
//	"github.com/spf13/pflag"
//	"github.com/stretchr/testify/require"
//	"testing"
//)

func getCMD(t *testing.T) (tCMD *cobra.Command) {
	tCMD = &cobra.Command{}
	require.NotNil(t, tCMD)
	return
}

func getFlagEnv(t *testing.T, flagSet *pflag.FlagSet) (fe *flagenv.FlagEnv) {
	fe = flagenv.NewFlagEnv(flagSet)
	require.NotNil(t, fe)

	require.Empty(t, fe.Usage())
	require.Empty(t, fe.EnvName())
	require.Empty(t, fe.Shorthand())
	require.Empty(t, fe.Normalizers())

	require.Equal(t, flagSet, fe.FlagSet())
	return
}

func getNew(
	t *testing.T,
	tCMD *cobra.Command,
	value flagenv.ValueSetter,
	opts ...flagenv.Modifier,
) {

	flagenv.New(tCMD.Flags(), value, opts...)
	require.True(t, tCMD.Flags().HasFlags())
}

func TestNew(t *testing.T) {
	t.Parallel()

	type NewTC struct {
		name        string
		valueSetter flagenv.ValueSetter
		opts        []flagenv.Modifier
		checker     func(*cobra.Command)
	}

	tShorthand, err := random.String(1)
	require.NoError(t, err)

	tFlagName, err := random.StringRange(1, 100)
	require.NoError(t, err)

	tEnvName, err := random.StringRange(1, 1000)
	require.NoError(t, err)

	tUsage, err := random.StringRange(1, 10000)
	require.NoError(t, err)

	rInt, err := random.IntRange(1, 5)
	require.NoError(t, err)

	tNormalizers := getNormalizers(t, rInt)

	tCCNormalizedFlagName := tFlagName
	for _, normalizer := range tNormalizers {
		tCCNormalizedFlagName = normalizer(tCCNormalizedFlagName)
	}

	tCCNormalizedFlagName = getCamelCased(tCCNormalizedFlagName)
	tCCFlagName := getCamelCased(tFlagName)
	tCCEnvName := getCamelCased(tEnvName)

	tc := []NewTC{
		{
			"WithFlagName",
			flagenv.String(),
			[]flagenv.Modifier{
				flagenv.WithFlagName(tFlagName),
			},
			func(cmd *cobra.Command) {
				tFlag := cmd.Flags().Lookup(tCCFlagName)
				require.NotNil(t, tFlag)
			},
		},
		{
			"WithEnvName",
			flagenv.String(),
			[]flagenv.Modifier{
				flagenv.WithEnvName(tEnvName),
			},
			func(cmd *cobra.Command) {
				tFlag := cmd.Flags().Lookup(tCCEnvName)
				require.NotNil(t, tFlag)
			},
		},
		{
			"WithShorthand",
			flagenv.String(),
			[]flagenv.Modifier{
				flagenv.WithShorthand(tShorthand),
				flagenv.WithFlagName(tFlagName),
			},
			func(cmd *cobra.Command) {
				tFlag := cmd.Flags().Lookup(tCCFlagName)
				require.NotNil(t, tFlag)

				require.Equal(t, tShorthand, tFlag.Shorthand)
			},
		},
		{
			"WithUsage",
			flagenv.String(),
			[]flagenv.Modifier{
				flagenv.WithUsage(tUsage),
				flagenv.WithFlagName(tFlagName),
			},
			func(cmd *cobra.Command) {
				tFlag := cmd.Flags().Lookup(tCCFlagName)
				require.NotNil(t, tFlag)

				require.Equal(t, tUsage, tFlag.Usage)
			},
		},
		{
			"WithNormalizers",
			flagenv.String(),
			[]flagenv.Modifier{
				flagenv.WithFlagName(tFlagName),
				flagenv.WithNormalizers(tNormalizers...),
			},
			func(cmd *cobra.Command) {
				tFlag := cmd.Flags().Lookup(tCCNormalizedFlagName)
				require.NotNil(t, tFlag)
			},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			tCMD := getCMD(t)

			getNew(
				t,
				tCMD,
				c.valueSetter,
				c.opts...,
			)

			c.checker(tCMD)
		})
	}

}

//func TestNew(t *testing.T) {
//	t.Parallel()
//
//	var (
//		flagParams [3]string
//		tFlag      *pflag.Flag
//		tCMD       *cobra.Command
//		fe         *flagenv.FlagEnv
//	)
//
//	tCMD = getCMD(t)
//	flagParams = getFlagNameParams(t)
//
//	fe = nil
//	fe = flagenv.New(
//		tCMD.Flags(),
//		flagenv.WithFlagName(flagParams[0]),
//		flagenv.WithStringValue(),
//	)
//	require.NotNil(t, fe)
//	require.True(t, tCMD.Flags().HasFlags())
//
//	tFlag = tCMD.Flags().Lookup(flagParams[1])
//	require.NotNil(t, tFlag)
//
//	require.Equal(t, tCMD.Flags(), fe.Flags())
//	require.Equal(t, flagParams[1], tFlag.Name)
//	require.Equal(t, flagParams[1], fe.FlagName())
//
//	require.Empty(t, fe.Usage())
//	require.Empty(t, tFlag.Usage)
//	require.Empty(t, fe.EnvName())
//	require.Empty(t, fe.Shorthand())
//	require.Empty(t, tFlag.Shorthand)
//}
