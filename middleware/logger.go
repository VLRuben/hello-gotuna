package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/alcalbg/gotdd/i18n"
	"github.com/alcalbg/gotdd/templating"
	"github.com/gorilla/mux"
)

func Logger(logger *log.Logger, locale i18n.Locale, staticPrefix string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			defer func() {
				if err := recover(); err != nil {
					stacktrace := string(debug.Stack())

					logger.Printf("PANIC RECOVERED: %v", err)
					logger.Println(stacktrace)

					//fmt.Println(err, stacktrace)

					w.WriteHeader(http.StatusInternalServerError)
					templating.GetEngine(locale, nil, staticPrefix). // TODO lang per user
												Set("error", err).
												Set("stacktrace", string(debug.Stack())).
												Render(w, r, "app.html", "error.html")
				}
			}()

			next.ServeHTTP(w, r)

			logger.Printf("%s %s %s %s", start.Format(time.RFC3339), r.Method, r.URL.Path, time.Since(start))
		})
	}
}
