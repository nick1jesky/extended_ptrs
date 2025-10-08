package extendedptrs

// create a deep copy
func DeepCopy[T any](t *T) *T {
	if t == nil {
		return nil
	}

	var copyT T
	switch v := any(*t).(type) {
	case interface{ DeepCopy() T }:
		copyT = v.DeepCopy()
	case interface{ Copy() T }:
		copyT = v.Copy()
	default:
		copyT = *t
	}

	return &copyT
}

// safely applies function to pointer
func Transform[T any](t *T, fn func(T) T) *T {
	if t == nil {
		return nil
	}
	result := fn(*t)
	return &result
}

// check equality of two pointers
func Equal[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// returns first non-nil pointer
func Coalesce[T any](pointers ...*T) *T {
	for _, p := range pointers {
		if p != nil {
			return p
		}
	}
	return nil
}

// applies function and returning new pointer
func Map[T, U any](t *T, fn func(T) U) *U {
	if t == nil {
		return nil
	}
	result := fn(*t)
	return &result
}

// returns pointer if a value satisfies the condition
func Filter[T any](t *T, predicate func(T) bool) *T {
	if t != nil && predicate(*t) {
		return t
	}
	return nil
}
