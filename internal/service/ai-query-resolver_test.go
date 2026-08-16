package service

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateSqlQueryResolver_Success(t *testing.T) {
	mockResponse := `{
		"id": "test_id",
		"status": "completed",
		"steps": [
			{"type": "thought"},
			{
				"type": "model_output",
				"content": [
					{"text": "SELECT * FROM public.telecall_number WHERE business_id = '123' LIMIT 10;", "type": "text"}
				]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.Header.Get("x-goog-api-key") != "test-key" {
			t.Errorf("expected test-key, got %s", r.Header.Get("x-goog-api-key"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected application/json, got %s", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	svc := &QueryResolverService{
		apiKey:     "test-key",
		apiURL:     server.URL,
		model:      "gemini-3.5-flash",
		httpClient: server.Client(),
	}

	query, err := svc.CreateSqlQueryResolver("generate query for test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "SELECT * FROM public.telecall_number WHERE business_id = '123' LIMIT 10;"
	if query != expected {
		t.Errorf("expected %q, got %q", expected, query)
	}
}

func TestCreateSqlQueryResolver_EmptyPrompt(t *testing.T) {
	svc := NewQueryResolverService("test-key")
	_, err := svc.CreateSqlQueryResolver("   ")
	if err == nil {
		t.Fatal("expected error for empty prompt, got nil")
	}
}

func TestCreateSqlQueryResolver_MissingAPIKey(t *testing.T) {
	os.Unsetenv("GEMINI_API_KEY")
	svc := &QueryResolverService{}
	_, err := svc.CreateSqlQueryResolver("generate query")
	if err == nil {
		t.Fatal("expected error when API key is missing, got nil")
	}
}
