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
	if err !