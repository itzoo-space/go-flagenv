package flagenv_test

import (
	"github.com/mazen160/go-random"
	"github.com/segmentio/go-camelcase"
	"github.com/stretchr/testify/require"
	"testing"
)

func getCamelCased(toCamelCase string) (camelCased string) {
	camelCased = camelcase.Camelcase(toCamelCase)

	for toCamelCase != camelCased {
		toCamelCase = camelCased
		camelCased = camelcase.Camelcase(toCamelCase)
	}
	return
}

func TestFlagEnv_FlagName(t *testing.T) {
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
	tFlagEnv.SetFlagName("")
	// ==================

	// ==================
	rStr, err = random.StringRange(1, 1000)
	require.NoError(t, err)

	tFlagEnv.SetEnvName(rStr)
	require.Equal(t, rStr, tFlagEnv.EnvName())

	tFlagName = getCamelCased(rStr)

	require.Equal(t, tFlagName, tFlagEnv.FlagName())
	tFlagEnv.SetFlagName("")
	// ==================

	// ==================
	rStr, err = random.StringRange(1, 1000)
	require.NoError(t, err)

	rInt, err := random.IntRange(1, 16)
	require.NoError(t, err)

	tNormalizersN := getNormalizers(t, rInt)
	tFlagEnv.SetNormalizers(tNormalizersN)
	tFlagName = rStr

	for _, n := range tNormalizersN {
		tFlagName = n(tFlagName)
	}

	tFlagName = getCamelCased(tFlagName)
	tFlagEnv.SetFlagName(rStr)
	require.Equal(t, tFlagName, tFlagEnv.FlagName())
	// ==================
}
