package logrustools

import (
	"context"
	"net/http"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const ContextKeyRequestID = "requestID"

type recorder struct {
	http.ResponseWriter

	// Response data.
	StatusCode  int
	CharsWriten int
}

func (r *recorder) Write(b []byte) (int, error) {

	if r.StatusCode == 0 {
		r.StatusCode = http.StatusOK
	}

	chars, err := r.ResponseWriter.Write(b)
	r.CharsWriten += chars

	return chars, err
}

func (r *recorder) WriteHeader(s int) {
	r.ResponseWriter.WriteHeader(s)
	r.StatusCode = s
}

func (r *recorder) Header() http.Header {
	return r.ResponseWriter.Header()
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Setup the uuid.
		id := uuid.New()

		// Setup context.
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextKeyRequestID, id.String())
		r = r.WithContext(ctx)

		// Setup logger.
		logger := logrus.New()
		logger.SetFormatter(&nested.Formatter{
			HideKeys: false,
			FieldsOrder: []string{
				"id", "addr", "status", "method", "url",
				"time", "chars", "referer", "user-agent",
			},
		})

		// Start timer.
		start := time.Now()

		// Serve next.
		rec := &recorder{w, 0, 0}
		next.ServeHTTP(rec, r)

		// Log.
		fields := logrus.Fields{
			"id":         id,
			"addr":       r.RemoteAddr,
			"status":     rec.StatusCode,
			"method":     r.Method,
			"url":        r.URL,
			"time":       time.Since(start),
			"chars":      rec.CharsWriten,
			"referer":    r.Referer(),
			"user-agent": r.UserAgent(),
		}

		logger.WithFields(fields).Infof("Handled request!")
	}
}
