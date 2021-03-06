package flagenv_test

import (
	"github.com/mazen160/go-random"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlagEnv_Usage(t *testing.T) {
	t.Parallel()

	tCMD := getCMD(t)
	tFlagEnv := getFlagEnv(t, tCMD.Flags())

	rStr, err := random.StringRange(1, 1000000)
	require.NoError(t, err)

	tFlagEnv.SetUsage(rStr)
	require.Equal(t, rStr, tFlagEnv.Usage())
}
