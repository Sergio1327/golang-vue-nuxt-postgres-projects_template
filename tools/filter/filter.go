package filter

func sliceToMap[T comparable](src ...T) map[T]struct{} {
	res := make(map[T]struct{}, len(src))

	for _, item := range src {
		res[item] = struct{}{}
	}

	return res
}

type Filter[T comparable, K comparable] struct {
	db map[T]map[K]struct{}
}

func NewFilter[T comparable, K comparable](filter map[T][]K) Filter[T, K] {
	l := Filter[T, K]{
		make(map[T]map[K]struct{}),
	}

	for k, v := range filter {
		l.db[k] = sliceToMap[K](v...)
	}

	return l
}

func (l Filter[T, K]) In(filter T, value K) bool {
	filterMap, exists := l.db[filter]
	if !exists {
		return false
	}

	_, exists = filterMap[value]

	return exists
}
