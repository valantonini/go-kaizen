package sync_once

import (
	"fmt"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// As opposed to init functions that are called at application start, this code executes when Do is called Intended for
// initialisation, if once.Do(f) is called multiple times, only the first call will invoke f, even if f has a different
// value in each invocation https://pkg.go.dev/sync#Once.
func Test_SyncOnce(t *testing.T) {
	Is := is.New(t)

	t.Run("it should run once regardless of parameters", func(t *testing.T) {
		var once sync.Once
		count := 0
		incrementCount := func() {
			count++
			if count > 1 {
				Is.Fail()
			}
		}

		for i := 0; i < 10; i++ {
			once.Do(incrementCount)
		}

		Is.Equal(count, 1)
	})

	t.Run("can be used with http.handlerFunc to create handlers with closures that don't slow down application startup", func(t *testing.T) {
		var fileContents string

		// https://www.youtube.com/watch?v=5DVV36uqQ4E (21mins)
		handleRequest := func(filename string) http.HandlerFunc {
			var init sync.Once
			return func(writer http.ResponseWriter, request *http.Request) {
				init.Do(func() {
					fileContents = fmt.Sprintf("file contents of %v", filename)
				})
				writer.Write([]byte(fileContents))
			}
		}

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/file", nil)

		handler := handleRequest("foo.txt")

		// handler has not been invoked yet so Do has not been run yet
		Is.Equal(fileContents, "")

		// first request will execute Do
		handler.ServeHTTP(rr, req)

		Is.Equal(rr.Body.String(), "file contents of foo.txt")
	})
}
