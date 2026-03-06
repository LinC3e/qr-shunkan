package qr

import "github.com/google/uuid"

type QRItem struct {
	ID  string
	URL string
}

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Create(url string) string {

	id := uuid.New().String()

	s.data[id] = url

	return id
}

func (s *Store) Get(id string) (string, bool) {
	u, ok := s.data[id]
	return u, ok
}