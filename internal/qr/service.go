package qr

import (
	"strconv"
	"sync"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

type Service struct {
	cache map[string][]byte
	store map[string]string
	mu    sync.RWMutex
}

func NewService() *Service {
	return &Service{
		cache: make(map[string][]byte),
		store: make(map[string]string),
	}
}

func (s *Service) Generate(url string, size int, format string) ([]byte, error) {

	key := url + "|" + strconv.Itoa(size) + "|" + format

	s.mu.RLock()
	data, ok := s.cache[key]
	s.mu.RUnlock()

	if ok {
		return data, nil
	}

	var err error

	if format == "png" {
		data, err = qrcode.Encode(url, qrcode.Medium, size)
	} else {
		data, err = qrcode.Encode(url, qrcode.Medium, size)
	}

	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.cache[key] = data
	s.mu.Unlock()

	return data, nil
}

func (s *Service) Create(url string) string {

	id := uuid.New().String()

	s.mu.Lock()
	s.store[id] = url
	s.mu.Unlock()

	return id
}

func (s *Service) Get(id string) (string, bool) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.store[id]

	return url, ok
}