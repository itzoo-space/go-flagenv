package flagenv

func (fe *FlagEnv) EnvName() string {
	return fe.envName
}

func (fe *FlagEnv) SetEnvName(envName string) *FlagEnv {
	fe.envName = envName
	return fe
}

func WithEnvName(envName string) Modifier {
	return func(fe *FlagEnv) {
		fe.SetEnvName(envName)
	}
}
