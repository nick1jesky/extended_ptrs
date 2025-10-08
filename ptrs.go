package  extendedptrs

func Ptr[T any](t T) *T {
	return &t
}

// safety for pointer dereference
func Val[T any](t *T) T {
	if t == nil {
		return Zero[T]()
	}
	return *t
}

func Zero[T any]() T {
	var zero T
	return zero
}

func IsZero[T comparable](t T) bool {
	return t == Zero[T]()
}

func ValOr[T any](t *T, defaultValue T) T {
	if t == nil {
		return defaultValue
	}
	return *t
}

func ValOrFunc[T any](t *T, defaultFunc func() T) T {
	if t == nil {
		return defaultFunc()
	}
	return *t
}
