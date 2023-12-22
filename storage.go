package acc

type DataStorage[T interface{}] interface {
	Save(data []T) error
	Load() ([]T, error)
}

type InMemoryStorage[T any] struct {
	data []Data[T]
}

func NewInMemoryStorage[T any]() *InMemoryStorage[T] {
	return &InMemoryStorage[T]{}
}

func (s *InMemoryStorage[T]) Save(data []Data[T]) error {
	for _, item := range data {
		s.data = append(s.data, item)
	}
	return nil
}

func (s *InMemoryStorage[T]) Load() ([]Data[T], error) {
	return s.data, nil
}
