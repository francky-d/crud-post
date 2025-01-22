package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("it return all posts", func(t *testing.T) {
		// Arrange
		want, err := json.Marshal(getPosts())
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest("GET", "/posts", nil)
		resp := httptest.NewRecorder()

		// Act
		Router().ServeHTTP(resp, req)

		// Assert
		if resp.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, resp.Result().StatusCode)
		}

		body := resp.Result().Body
		defer body.Close()
		got, err := io.ReadAll(body)

		if err != nil {
			t.Fatal(err)
		}

		if string(want) != string(got) {
			t.Errorf("expected %s but got %s", want, got)
		}

	})

	t.Run("it return post by id", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/posts/2", nil)
		resp := httptest.NewRecorder()

		want, err := json.Marshal(getPosts()[1])

		if err != nil {
			t.Fatal(err)
		}
		// Act
		Router().ServeHTTP(resp, req)

		if resp.Result().StatusCode != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, resp.Result().StatusCode)
		}

		// Assert
		body := resp.Result().Body

		got, err := io.ReadAll(body)

		if err != nil {
			t.Fatal(err)
		}

		if string(want) != string(got) {
			t.Errorf("expected %s but got %s", want, got)
		}
	})
}
