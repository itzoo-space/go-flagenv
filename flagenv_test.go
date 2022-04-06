package flagenv_test

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"testing"
)

func getCMD(t *testing.T) (tCMD *cobra.Command) {
	tCMD = &cobra.Command{}
	require.NotNil(t, tCMD)
	require.False(t, tCMD.Flags().HasFlags())
	return
}

func TestNew(t *testing.T) {
	t.Parallel()

	var (
		flagParams [3]string
		tFlag      *pflag.Flag
		tCMD       *cobra.Command
		fe         *flagenv.FlagEnv
	)

	tCMD = getCMD(t)
	flagParams = getFlagNameParams(t)

	fe = nil
	fe = flagenv.New(
		tCMD.Flags(),
		flagenv.WithFlagName(flagParams[0]),
		flagenv.WithStringValue(),
	)
	require.NotNil(t, fe)
	require.True(t, tCMD.Flags().HasFlags())

	tFlag = tCMD.Flags().Lookup(flagParams[1])
	require.NotNil(t, tFlag)

	require.Equal(t, tCMD.Flags(), fe.Flags())
	require.Equal(t, flagParams[1], tFlag.Name)
	require.Equal(t, flagParams[1], fe.FlagName())

	require.Empty(t, fe.Usage())
	require.Empty(t, tFlag.Usage)
	require.Empty(t, fe.EnvName())
	require.Empty(t, fe.Shorthand())
	require.Empty(t, tFlag.Shorthand)
}
