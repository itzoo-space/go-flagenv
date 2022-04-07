package flagenv_test

import (
	"github.com/mazen160/go-random"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func getNormalizers(t *testing.T, number int) (normalizers []func(string) string) {
	for i := 0; i < number; i++ {
		rL, err := random.String(1)
		require.NoError(t, err)

		rC, err := random.IntRange(0, 1)
		require.NoError(t, err)

		rL = strings.ToLower(rL)

		if rC == 1 {
			rL = strings.ToUpper(rL)
		}

		normalizers = append(normalizers, func(s string) string {
			return strings.ReplaceAll(s, rL, "")
		})
	}
	return
}

func TestFlagEnv_Normalizers(t *testing.T) {
	t.Parallel()

	tCMD := getCMD(t)
	tFlagEnv := getFlagEnv(t, tCMD.Flags())

	tNormalizers1 := getNormalizers(t, 1)
	tFlagEnv.SetNormalizers(tNormalizers1)
	require.Equal(t, tNormalizers1, tFlagEnv.Normalizers())

	rInt, err := random.IntRange(1, 16)
	require.NoError(t, err)

	tNormalizersN := getNormalizers(t, rInt)
	tFlagEnv.SetNormalizers(tNormalizersN)
	require.Equal(t, tNormalizersN, tFlagEnv.Normalizers())
}
