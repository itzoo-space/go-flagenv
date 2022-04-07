package flagenv_test

import (
	"github.com/mazen160/go-random"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestFlagEnv_Shorthand(t *testing.T) {
	t.Parallel()

	var tFlagName string

	tCMD := getCMD(t)
	tFlagEnv := getFlagEnv(t, tCMD.Flags())

	// ==================
	rStr, err := random.StringRange(1, 1000)
	require.NoError(t, err)

	tFlagName = getCamelCased(rStr)

	tFlagEnv.SetFlagName(rStr)
	require.Equal(t, tFlagName, tFlagEnv.FlagName())

	tFlagEnv.SetShorthand("")
	require.Empty(t, tFlagEnv.Shorthand())

	tFlagEnv.SetShorthand("-")
	require.Equal(t, strings.ToLower(tFlagName[0:1]), tFlagEnv.Shorthand())

	tFlagEnv.SetShorthand("-^")
	require.Equal(t, strings.ToUpper(tFlagName[0:1]), tFlagEnv.Shorthand())

	rL, err := random.String(1)
	require.NoError(t, err)

	tFlagEnv.SetShorthand(rL)
	require.Equal(t, rL, tFlagEnv.Shorthand())
	// ==================
}
