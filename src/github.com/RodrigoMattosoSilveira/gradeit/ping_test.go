package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/RodrigoMattosoSilveira/gradeit/configs"
	"github.com/RodrigoMattosoSilveira/gradeit/routes"
)

func TestPingRoute(t *testing.T) {
	router := configs.GetRouter()
	routes.RoutesPing(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
