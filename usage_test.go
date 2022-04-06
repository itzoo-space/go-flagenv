package flagenv_test

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/mazen160/go-random"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlagEnv_Usage(t *testing.T) {
	t.Parallel()

	var (
		err        error
		rStr       string
		flagParams [3]string
		tFlag      *pflag.Flag
		tCMD       *cobra.Command
		fe         *flagenv.FlagEnv
	)

	tCMD = getCMD(t)
	flagParams = getFlagNameParams(t)

	rStr, err = random.StringRange(1, 10000)
	require.NoError(t, err)

	fe = nil
	fe = flagenv.New(
		tCMD.Flags(),
		flagenv.WithFlagName(flagParams[0]),
		flagenv.WithUsage(rStr),
		flagenv.WithStringValue(),
	)
	require.NotNil(t, fe)
	require.True(t, tCMD.Flags().HasFlags())

	tFlag = tCMD.Flags().Lookup(flagParams[1])
	require.NotNil(t, tFlag)

	require.Equal(t, tCMD.Flags(), fe.Flags())
	require.Equal(t, flagParams[1], tFlag.Name)
	require.Equal(t, flagParams[1], fe.FlagName())

	require.Empty(t, fe.EnvName())
	require.Empty(t, fe.Shorthand())
	require.Empty(t, tFlag.Shorthand)

	require.Equal(t, rStr, fe.Usage())
	require.Equal(t, rStr, tFlag.Usage)
}
