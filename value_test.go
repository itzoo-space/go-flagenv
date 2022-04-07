package flagenv_test

import (
	"github.com/itzoo-space/go-flagenv"
	"github.com/mazen160/go-random"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
)

func TestBool(t *testing.T) {
	t.Parallel()

	type BoolTC struct {
		name         string
		tValueSetter flagenv.ValueSetter
		tValue       bool
	}

	tc := []BoolTC{
		{
			"Default",
			flagenv.Bool(),
			false,
		},
		{
			"False",
			flagenv.Bool(false),
			false,
		},
		{
			"True",
			flagenv.Bool(true),
			true,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			tCMD := getCMD(t)
			name := strings.ToLower(c.name)

			getNew(
				t,
				tCMD,
				c.tValueSetter,
				flagenv.WithFlagName(name),
			)

			tFlag := tCMD.Flags().Lookup(name)
			require.NotNil(t, tFlag)

			tValue, err := strconv.ParseBool(tFlag.DefValue)
			require.NoError(t, err)

			require.Equal(t, c.tValue, tValue)
		})
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	type BoolTC struct {
		name         string
		tValueSetter flagenv.ValueSetter
		tValue       int
	}

	rInt, err := random.IntRange(1, 1000000000000)
	require.NoError(t, err)

	tc := []BoolTC{
		{
			"Default",
			flagenv.Int(),
			0,
		},
		{
			"0",
			flagenv.Int(0),
			0,
		},
		{
			"rInt",
			flagenv.Int(rInt),
			rInt,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			tCMD := getCMD(t)
			name := strings.ToLower(c.name)

			getNew(
				t,
				tCMD,
				c.tValueSetter,
				flagenv.WithFlagName(name),
			)

			tFlag := tCMD.Flags().Lookup(name)
			require.NotNil(t, tFlag)

			tValue, err := strconv.ParseInt(tFlag.DefValue, 10, 64)
			require.NoError(t, err)

			require.Equal(t, c.tValue, int(tValue))
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	type BoolTC struct {
		name         string
		tValueSetter flagenv.ValueSetter
		tValue       string
	}

	rStr, err := random.StringRange(1, 100000)
	require.NoError(t, err)

	tc := []BoolTC{
		{
			"Default",
			flagenv.String(),
			"",
		},
		{
			"Empty",
			flagenv.String(""),
			"",
		},
		{
			"rStr",
			flagenv.String(rStr),
			rStr,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			tCMD := getCMD(t)
			name := strings.ToLower(c.name)

			getNew(
				t,
				tCMD,
				c.tValueSetter,
				flagenv.WithFlagName(name),
			)

			tFlag := tCMD.Flags().Lookup(name)
			require.NotNil(t, tFlag)

			require.Equal(t, c.tValue, tFlag.DefValue)
		})
	}
}
