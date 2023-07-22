package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/msales/kage"
	"github.com/msales/kage/kafka"
	"github.com/msales/kage/server"
	"github.com/msales/kage/testutil"
	"github.com/msales/kage/testutil/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBrokersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/brokers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	monitor := new(mocks.MockMonitor)
	monitor.On("Brokers").Return([]kafka.Broker{{ID: 0, Connected: false}})

	app := &kage.Application{Monitor: monitor}

	srv := server.New(app)
	srv.ServeHTTP(rr, req)

	want := "[{\"id\":0,\"connected\":false}]"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, want, rr.Body.String())
}

func TestBrokersHealthHandler(t *testing