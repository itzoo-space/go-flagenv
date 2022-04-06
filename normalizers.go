package flagenv

func (fe *FlagEnv) Normalizers() []func(string) string {
	return fe.normalizers
}

func (fe *FlagEnv) SetNormalizers(normalizers []func(string) string) *FlagEnv {
	fe.normalizers = normalizers
	return fe
}

func WithNormalizers(normalizers ...func(string) string) Modifier {
	return func(fe *FlagEnv) {
		fe.SetNormalizers(append(fe.normalizers, normalizers...))
	}
}
