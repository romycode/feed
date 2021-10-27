package storage

import (
	"fmt"
	"os"

	"github.com/romycode/go-feeder/internal"
)

type FileSKUStore struct {
	file *os.File
}

func NewFileSKUStore(f *os.File) (internal.SKUStore, error) {
	return &FileSKUStore{file: f}, nil
}

func (f FileSKUStore) Save(skus []internal.SKU) error {
	for _, sku := range skus {
		_, err := f.file.WriteString(sku.Value())
		if err != nil {
			return fmt.Errorf("error saving into file: %s, error: %w", f.file.Name(), err)
		}
	}

	return nil
}
