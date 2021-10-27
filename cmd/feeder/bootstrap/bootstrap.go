package bootstrap

import (
	"os"
	"time"

	"github.com/romycode/go-feeder/internal"
	"github.com/romycode/go-feeder/internal/platform/storage"
)

func BuildApp(addr string, maxConn int, ttl time.Duration, logFile *os.File) (*internal.App, error) {
	deduplicator := internal.NewDeduplicator()
	skuStore, err := storage.NewFileSKUStore(logFile)
	if err != nil {
		return &internal.App{}, err
	}

	return internal.NewApp(addr, maxConn, ttl, deduplicator, skuStore), nil
}
