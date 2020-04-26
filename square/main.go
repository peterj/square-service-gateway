package main
import (
	"context"
	"fmt"
	"flag"
	"net/http"
	"time"
	"os"
	"os/signal"
	"syscall"

	"square/pkg/server"

	"github.com/sirupsen/logrus"
)
const (
	defaultPort = "8080"
	shutdownTimeout = 5 * time.Second
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	flag.Parse()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s := server.New(context.Background())
    hs := &http.Server{
        Handler:      s,
        Addr:         fmt.Sprintf(":%s", port),
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
	}

	go func() {
		logrus.Printf("Running on %s", defaultPort)
		if err := hs.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Fatalf("failed to start the server %+v", err)
		}
	}()

	shutdown(hs, shutdownTimeout)
}

// shutdown gracefully shuts down the HTTP server
func shutdown(h *http.Server, timeout time.Duration) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logrus.Printf("shutting down with timeout %s", timeout)
	if err := h.Shutdown(ctx); err != nil {
		logrus.Fatalf("shutdown failed: %v", err)
	} else {
		logrus.Printf("shutdown completed")
	}
}