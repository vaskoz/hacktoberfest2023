package publicapi

import (
	"io"
	"net/http"
	"os"
	"testing"
)

var fakeExit = func(int) {}

func BenchmarkEndpoint(b *testing.B) {
	app := NewApp(nil, os.Stdout, os.Stderr, fakeExit, ":8080")
	go func() {
		app.Run()
	}()

	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			b.Fail()
		}
		bytez, err := io.ReadAll(resp.Body)
		if err != nil {
			b.Fail()
		}
		msg := string(bytez)
		if msg != "hello world!" {
			b.Fail()
		}
	}

	app.(*publicApp).StopChan <- os.Interrupt
}
