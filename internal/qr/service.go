package qr

import "github.com/skip2/go-qrcode"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Generate(url string, size int) ([]byte, error) {

	return qrcode.Encode(url, qrcode.Medium, size)

}