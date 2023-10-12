package mapfunc

// DeepCopy создает настоящую копию мап
func DeepCopy[K comparable, V any](src map[K]V) map[K]V {
	result := make(map[K]V, len(src))
	for k, v := range src {
		result[k] = v
	}

	return result
}

func AppendMap[T comparable, K any](src, app map[T]K) {
	for t, k := range app {
		src[t] = k
	}
}

func SliceToMap[T comparable](src ...T) map[T]struct{} {
	res := make(map[T]struct{}, len(src))

	for _, item := range src {
		res[item] = struct{}{}
	}

	return res
}
