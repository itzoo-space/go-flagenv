package flagenv_test

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyz"
	letterBytesU = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbolBytes  = ".,;:|?$%@][{}#&/()*"
	numberBytes  = "0123456789"

	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	src    = rand.NewSource(time.Now().UnixNano())
	flagTC = map[string][]string{
		"flag": {
			"flag",
			"Flag",
			"_flag",
			".flag",
			"-flag",
			"FLAG",
		},
		"flagEnv": {
			"flagEnv",
			"FlagEnv",
			"flag_env",
			"flag.env",
			"flag-env",
			"FLAG_ENV",
		},
		"flagEnvTC": {
			"flagEnvT C",
			"FlagEnvT C",
			"flag_env_t_c",
			"flag.env.t.c",
			"flag-env-t-c",
			"FLAG_ENV_T_C",
		},
	}
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randString(n int, pools ...string) string {
	pool := letterBytes + letterBytesU + symbolBytes + numberBytes
	if len(pools) > 0 {
		pool = strings.Join(pools, "")
	}

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(pool) {
			b[i] = pool[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func TestNewFlag(t *testing.T) {
	t.Parallel()

	testCmd := &cobra.Command{}
	tf := testCmd.Flags()
	f := flagenv.NewFlag(testCmd.Flags())
	require.Equal(t, tf, f.FlagSet())
}

func TestFlag_Usage(t *testing.T) {
	t.Parallel()

	f := flagenv.NewFlag((&cobra.Command{}).Flags())
	require.Empty(t, f.Usage())

	var words []string

	for w := 0; w < randInt(0, 1000); w++ {
		wl := randInt(1, 32)
		words = append(words, randString(wl))
	}

	usage := strings.Join(words, " ")

	f.SetUsage(usage)
	require.Equal(t, usage, f.Usage())
}

func TestFlag_Value(t *testing.T) {
	t.Parallel()

	f := flagenv.NewFlag((&cobra.Command{}).Flags())
	require.Empty(t, f.Value())

	value := randString(randInt(0, 1000))

	f.SetValue(value)
	require.Equal(t, value, f.Value())
}

func TestFlag_EnvName(t *testing.T) {
	t.Parallel()

	f := flagenv.NewFlag((&cobra.Command{}).Flags())
	require.Empty(t, f.EnvName())

	var words []string

	for w := 0; w < randInt(0, 100); w++ {
		wl := randInt(1, 50)
		words = append(words, randString(wl))
	}

	usage := strings.Join(words, "_")

	f.SetEnvName(usage)
	require.Equal(t, usage, f.EnvName())
}

func TestFlag_Name(t *testing.T) {
	t.Parallel()

	f := flagenv.NewFlag((&cobra.Command{}).Flags())
	require.Empty(t, f.Name())

	for outFlag, inFlags := range flagTC {
		for _, inFlag := range inFlags {
			f.SetName(inFlag)
			require.Equal(t, outFlag, f.Name())
		}
	}

	f.SetName("")
	require.Empty(t, f.Name())

	p := randString(randInt(0, 500))
	f.AddNormalizer(func(n string) string {
		return strings.Replace(n, p, "", 1)
	})

	for outFlag, inFlags := range flagTC {
		for _, inFlag := range inFlags {
			f.SetName(p + inFlag)
			require.Equal(t, outFlag, f.Name())
		}
	}

}

func TestFlag_Shorthand(t *testing.T) {
	t.Parallel()

	f := flagenv.NewFlag((&cobra.Command{}).Flags())
	require.Empty(t, f.Shorthand())

	for outFlag, inFlags := range flagTC {
		for _, inFlag := range inFlags {
			f.SetName(inFlag)
			require.Equal(t, outFlag, f.Name())

			s := randString(1, letterBytes+letterBytesU)
			f.SetShorthand(s)
			require.Equal(t, s, f.Shorthand())

			f.SetShorthand("")
			require.Empty(t, f.Shorthand())

			f.SetShorthand("-")
			require.Equal(t, f.Name()[0:1], f.Shorthand())

			f.SetShorthand("-^")
			require.Equal(t, strings.ToUpper(f.Name()[0:1]), f.Shorthand())
		}
	}
}

func TestFlag_SetCLIFlag(t *testing.T) {
	t.Parallel()

	testCmd := &cobra.Command{}
	tf := testCmd.Flags()
	f := flagenv.NewFlag(testCmd.Flags())
	require.Equal(t, tf, f.FlagSet())

	f.SetName(randString(randInt(0, 32), letterBytes))
	f.SetValue(randString(randInt(0, 500)))
	f.SetUsage(randString(randInt(0, 1000)))
	f.SetShorthand(randString(1, letterBytes+letterBytesU))

	cmdF := testCmd.Flags().Lookup(f.Name())
	require.Nil(t, cmdF)

	f.SetCLIFlag()

	cmdF = testCmd.Flags().Lookup(f.Name())
	require.Equal(t, cmdF.Name, f.Name())
	require.Equal(t, cmdF.Value.String(), f.Value())
	require.Equal(t, cmdF.Usage, f.Usage())
	require.Equal(t, cmdF.Shorthand, f.Shorthand())
}

func TestFlag_SetEnv(t *testing.T) {
	t.Parallel()

	testCmd := &cobra.Command{}
	tf := testCmd.Flags()
	f := flagenv.NewFlag(testCmd.Flags())
	require.Equal(t, tf, f.FlagSet())

	f.SetName(randString(randInt(0, 32), letterBytes))
	f.SetValue(randString(randInt(0, 500)))
	f.SetEnvName(randString(randInt(0, 500), letterBytes+letterBytesU+numberBytes))

	envV := viper.GetString(f.EnvName())
	require.Empty(t, envV)

	f.SetCLIFlag()
	f.SetEnv()

	envV = viper.GetString(f.EnvName())
	require.Equal(t, envV, f.Value())
}

func TestNew(t *testing.T) {
	t.Parallel()

	testCmd := &cobra.Command{}

	f := testCmd.Flags()
	v := randString(randInt(0, 500))
	u := randString(randInt(0, 1000))
	s := randString(1, letterBytes+letterBytesU)
	e := randString(randInt(0, 500), letterBytes+letterBytesU+numberBytes)

	n := flagenv.NewFlag(f).SetName(e).Name()

	cmdF := f.Lookup(n)
	require.Nil(t, cmdF)

	envV := viper.GetString(e)
	require.Empty(t, envV)

	flagenv.New(f, e, s, v, u)

	cmdF = f.Lookup(n)
	require.Equal(t, cmdF.Name, n)
	require.Equal(t, cmdF.Usage, u)
	require.Equal(t, cmdF.Shorthand, s)
	require.Equal(t, cmdF.Value.String(), v)

	envV = viper.GetString(e)
	require.Equal(t, envV, v)
}
