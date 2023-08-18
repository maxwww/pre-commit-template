package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestPingHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/ping", pingHandler)

	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	responseData, _ := io.ReadAll(resp.Body)

	if string(responseData) != "pong" {
		t.Errorf("unexpected response. expexted: pong, got: %s", responseData)
		return
	}
}
