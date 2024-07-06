package server

import (
	"net/http"
	"testing"
)

func TestServerStart(t *testing.T) {
	go func() {
		err := StartServer()
		if err != nil {
			t.Errorf("Error starting server: %v", err)
		}
	}()

	get, getErr := http.Get("http://localhost:8080/")
	if getErr != nil {
		t.Errorf("Error starting server: %v", getErr)
	}
	if get.StatusCode != http.StatusNotFound {
		t.Errorf("Server returned wrong status code: got %v want %v", get.StatusCode, http.StatusNotFound)
	}
}
