package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/romycode/go-feeder/internal/platform/server"
)

type App struct {
	ttl          time.Duration     // ttl time to life up for the app
	svr          *server.TCPServer // svr tcp listener for connections
	sigint       chan os.Signal    // sigint chan for os interrupt event
	msgs         chan string       // msgs chan for messages to process
	deduplicator *Deduplicator     // deduplicator service to process received messages
	skuStorage   SKUStore          // skuStorage service to store skus
}

func NewApp(addr string, maxConn int, ttl time.Duration, deduplicator *Deduplicator, store SKUStore) *App {
	app := &App{ttl: ttl, sigint: make(chan os.Signal, 1), msgs: make(chan string, 10), deduplicator: deduplicator, skuStorage: store}
	app.svr = server.NewTCPServer(addr, maxConn, app.msgs)

	return app
}

// Start wake-up app and start TCP listener
func (a *App) Start() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.ttl)
	listenOsInterruptSignal(a.sigint, cancel)

	go a.handleOutput(cancel)

	if err := a.svr.Start(ctx); err != nil {
		return fmt.Errorf("error starting tcp server: %w", err)
	}

	a.Close()

	err := a.skuStorage.Save(a.deduplicator.values)
	if err != nil {
		return fmt.Errorf("error saving skus: %w", err)
	}

	report := a.deduplicator.Report()
	fmt.Printf(
		"Received %d unique product skus, %d duplicates, %d discard values\n",
		report.ValidCount,
		report.DuplicatedCount,
		report.InvalidCount,
	)

	return nil
}

func (a App) handleOutput(cancel context.CancelFunc) {
	for msg := range a.msgs {
		if "terminate\n" == msg {
			log.Println("client sent 'terminate' sequence gracefully shutdown started...")
			cancel()
			break
		}

		_ = a.deduplicator.Process(msg)
	}
}

func (a App) Close() {
	close(a.sigint)
	close(a.msgs)
}

// listenOsInterruptSignal will cancel the context for CTRL+C
func listenOsInterruptSignal(c chan os.Signal, cancel func()) {
	signal.Notify(c, os.Interrupt)
	go func() {
		_, ok := <-c
		if ok {
			log.Println("SIGINT: stopping server gracefully...")
		}
		cancel()
	}()
}
