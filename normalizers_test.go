package flagenv_test

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/mazen160/go-random"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"testing"
)

func getNormalizers(src string) (normalizers [3]func(s string) string) {
	normalizers[0] = func(s string) string {
		return s + src
	}
	normalizers[1] = func(s string) string {
		return src + s
	}
	normalizers[2] = func(s string) string {
		return src + s + src
	}
	return
}

func getNormalizersList(t *testing.T, listLen int, src string) (normalizersList []func(s string) string) {
	if listLen == 0 {
		return
	}

	var (
		rInt int
		err  error
	)

	normalizers := getNormalizers(src)

	for i := 0; i < listLen; i++ {
		rInt, err = random.IntRange(0, len(normalizers))
		require.NoError(t, err)

		normalizersList = append(normalizersList, normalizers[rInt])
	}

	return
}

func TestFlagEnv_Normalizers(t *testing.T) {
	t.Parallel()

	var (
		param,
		tFlagName string
		rInt            int
		err             error
		flagParams      [3]string
		tFlag           *pflag.Flag
		tCMD            *cobra.Command
		fe              *flagenv.FlagEnv
		normalizer      func(s string) string
		normalizersList []func(s string) string
		normalizers     [3]func(s string) string
	)

	for i := 0; i < 10; i++ {
		flagParams = getFlagNameParams(t)

		for _, param = range flagParams {
			normalizers = getNormalizers(param)

			for _, normalizer = range normalizers {
				tFlagName = getCamelCased(normalizer(param))
				tCMD = getCMD(t)

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

			rInt, err = random.IntRange(0, 10)
			require.NoError(t, err)

			normalizersList = getNormalizersList(t, rInt, param)

			tFlagName = param
			for _, normalizer = range normalizersList {
				tFlagName = normalizer(tFlagName)
			}
			tFlagName = getCamelCased(tFlagName)

			tCMD = getCMD(t)

			fe = nil
			fe = flagenv.New(
				tCMD.Flags(),
				flagenv.WithNormalizers(normalizersList...),
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
}
