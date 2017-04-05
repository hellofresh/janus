package middleware_test

import (
	"testing"

	"net/http"

	"github.com/hellofresh/janus/pkg/middleware"
	"github.com/hellofresh/janus/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulLog(t *testing.T) {
	mw := middleware.NewLogger()
	w, err := test.Record(
		"GET",
		"/",
		map[string]string{
			"Content-Type": "application/json",
		},
		recovery(mw.Handler(http.HandlerFunc(test.Ping))),
	)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}