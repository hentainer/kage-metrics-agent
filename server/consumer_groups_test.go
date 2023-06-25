package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/msales/kage"
	"github.com/msales/kage/server"
	"github.com/msales/kage/store"
	"github.com/msales/kage/testutil/mocks"
	"github.com/stretchr/testify/assert"
)

func TestConsumerGroupsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/consumers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	co := store.ConsumerOffsets{
		"test": map[string][]*store.ConsumerOffset{
			"test": {{Offset: 0, Lag: 100, Timestamp: 0}},
		},
	}

	store := new(mocks.MockStore)
	store.On("ConsumerOffsets").Retu