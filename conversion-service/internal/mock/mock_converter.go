package mock

import (
	"ConversionService/internal/domain/model"
)

type MockConverter struct {
	pdf      []byte
	filename string
	error    error
}

func NewMockConverter(pdf []byte, filename string, error error) *MockConverter {
	return &MockConverter{
		pdf:      pdf,
		filename: filename,
		error:    error,
	}
}

func (receiver MockConverter) Convert(data model.Data) ([]byte, string, error) {
	return receiver.pdf, receiver.filename, receiver.error
}
