package publicapi

import (
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// PublicApp is the application for the public-api.
type PublicApp interface {
	Run() error
}

type publicApp struct {
	Args           []string
	Stdout, Stderr io.Writer
	ExitFunc       func(int)
	StopChan       chan os.Signal
	HTTPPort       string
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello world!"))
}

// NewApp returns a new instance of PublicApp with the given globals passed in.
func NewApp(args []string, stdout, stderr io.Writer, exit func(int), httpPort string) PublicApp {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	return &publicApp{
		Args:     args,
		Stdout:   stdout,
		Stderr:   stderr,
		ExitFunc: exit,
		StopChan: stop,
		HTTPPort: httpPort,
	}
}

func (pa *publicApp) Run() error {
	go func() {
		<-pa.StopChan
		pa.ExitFunc(1)
	}()
	return http.ListenAndServe(pa.HTTPPort, http.HandlerFunc(HomeHandler))
}
