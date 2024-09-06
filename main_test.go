package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Reu")
	want := "Hello, Reu"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDownloadFile(t *testing.T) {
	// create mock server and file
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test File"))
	}))
	defer mockServer.Close()

	filepath := "testfile.txt"
	url := mockServer.URL

	// test DownloadFile function
	err := DownloadFile(filepath, url)
	if err != nil {
		t.Fatalf("Expected to download a file but got %v instead", err)
	}

	// check if file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		t.Fatalf("File does not exist")
	}

	// check for file content
	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Expected no error reading file, but got %v", err)
	}

	expectedContent := "Test File"
	if string(content) != expectedContent {
		t.Fatalf("Expected file content to be %q, but got %q instead", expectedContent, string(content))
	}
}
