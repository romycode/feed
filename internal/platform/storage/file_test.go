package storage

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/romycode/go-feeder/internal"
)

func TestFileSKUStore_Save(t *testing.T) {
	filename := fmt.Sprintf("sku-%s.log", time.Now().Format("2006-02-01_15:03:04"))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("could not open output file: %s\n", err.Error())
	}
	defer func() {
		_ = file.Close()
		_ = os.RemoveAll(filename)
	}()

	got, _ := NewFileSKUStore(file)

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		t.Errorf("NewFileSKUStore(): not work correctly")
	}

	sku := internal.NewSKU()
	_ = got.Save([]internal.SKU{sku})

	rawFile, _ := os.ReadFile(filename)

	if string(rawFile) != sku.Value() {
		t.Errorf("NewFileSKUStore(): want %s, got %s", sku.Value(), string(rawFile))
	}
}
