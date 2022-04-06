package flagenv

func (fe *FlagEnv) Usage() string {
	return fe.usage
}

func (fe *FlagEnv) SetUsage(usage string) *FlagEnv {
	fe.usage = usage
	return fe
}

func WithUsage(Usage string) Modifier {
	return func(fe *FlagEnv) {
		fe.SetUsage(Usage)
	}
}
