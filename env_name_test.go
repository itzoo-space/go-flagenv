package flagenv_test

import (
	"github.com/mazen160/go-random"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlagEnv_EnvName(t *testing.T) {
	t.Parallel()

	tCMD := getCMD(t)
	tFlagEnv := getFlagEnv(t, tCMD.Flags())

	rStr, err := random.StringRange(1, 1000)
	require.NoError(t, err)

	tFlagEnv.SetEnvName(rStr)
	require.Equal(t, rStr, tFlagEnv.EnvName())
}
