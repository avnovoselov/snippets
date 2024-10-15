package test

func KeepOLD[T any](old *T, temporary *T) (restore func()) {
	keep := new(T)
	*keep, *old = *old, *temporary

	return func() {
		*old = *keep
	}
}
