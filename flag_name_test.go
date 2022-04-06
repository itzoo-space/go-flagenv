package flagenv_test

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/mazen160/go-random"
	"github.com/segmentio/go-camelcase"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

func getFlagNameParams(t *testing.T) (params [3]string) {
	pL, err := random.Random(1, random.ASCIILettersLowercase, true)
	require.NoError(t, err)

	pU, err := random.Random(1, random.ASCIILettersUppercase, true)
	require.NoError(t, err)

	pD, err := random.Random(1, random.Digits, true)
	require.NoError(t, err)

	rInt, err := random.IntRange(1, 1000)
	require.NoError(t, err)

	pA, err := random.Random(rInt, random.ASCIICharacters, true)
	require.NoError(t, err)

	params[0] = pD + pL + pU + pA
	params[1] = getCamelCased(params[0])
	params[2] = camelcase.Camelcase(params[0])
	require.NotEqual(t, params[0], params[1])
	require.NotEqual(t, params[1], params[2])
	return
}

func TestFlagEnv_FlagName(t *testing.T) {
	t.Parallel()

	var (
		rStr,
		value,
		tFlagName,
		param string
		rInt        int
		err         error
		flagParams  [3]string
		tFlag       *pflag.Flag
		tCMD        *cobra.Command
		fe          *flagenv.FlagEnv
		normalizer  func(string) string
		normalizers [3]func(s string) string
	)

	flagParams = getFlagNameParams(t)

	for _, param = range flagParams {
		tCMD = getCMD(t)

		value, err = random.StringRange(0, 1000)
		require.NoError(t, err)

		fe = nil
		fe = flagenv.New(
			tCMD.Flags(),
			flagenv.WithFlagName(param),
			flagenv.WithStringValue(value),
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

		require.Equal(t, value, tFlag.DefValue)

		rInt, err = random.IntRange(1, len(param)-1)
		require.NoError(t, err)

		fe.SetFlagName(param[:rInt])
		require.Equal(t, flagParams[1][:rInt], fe.FlagName())
		require.NotEqual(t, tFlag.Name, fe.FlagName())

		tCMD = getCMD(t)
		rStr, err = random.StringRange(1, len(param)-1)
		require.NoError(t, err)

		normalizers = getNormalizers(rStr)
		rInt, err = random.IntRange(0, len(normalizers))
		require.NoError(t, err)

		normalizer = normalizers[rInt]
		tFlagName = getCamelCased(normalizer(param))
		fe = nil
		fe = flagenv.New(
			tCMD.Flags(),
			flagenv.WithNormalizers(normalizer),
			flagenv.WithFlagName(param),
			flagenv.WithStringValue(),
		)
		require.NotNil(t, fe)
		require.True(t, tCMD.Flags().HasFlags())

		tFlag = tCMD.Flags().Lookup(tFlagName)
		require.NotNil(t, tFlag)

		require.Equal(t, tFlagName, tFlag.Name)
		require.Equal(t, tFlagName, fe.FlagName())
	}
}
