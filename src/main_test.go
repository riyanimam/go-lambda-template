package main

import (
	"context"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		request     Request
		wantMessage string
		wantStatus  int
		wantErr     bool
	}{
		{
			name:        "successful invocation with default name",
			ctx:         context.Background(),
			request:     Request{},
			wantMessage: "Hello, world!",
			wantStatus:  200,
			wantErr:     false,
		},
		{
			name:        "successful invocation with custom name",
			ctx:         context.Background(),
			request:     Request{Name: "Alice"},
			wantMessage: "Hello, Alice!",
			wantStatus:  200,
			wantErr:     false,
		},
		{
			name:        "invocation with custom context",
			ctx:         context.WithValue(context.Background(), "key", "value"),
			request:     Request{Name: "Bob"},
			wantMessage: "Hello, Bob!",
			wantStatus:  200,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler(tt.ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Message != tt.wantMessage {
				t.Errorf("handler() Message = %v, want %v", got.Message, tt.wantMessage)
			}
			if got.Status != tt.wantStatus {
				t.Errorf("handler() Status = %v, want %v", got.Status, tt.wantStatus)
			}
		})
	}
}

func TestHandlerContext(t *testing.T) {
	ctx := context.Background()
	request := Request{Name: "Test"}

	// Test that handler doesn't panic with background context
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("handler() panicked: %v", r)
		}
	}()

	response, err := handler(ctx, request)
	if err != nil {
		t.Errorf("handler() returned unexpected error: %v", err)
	}
	if response.Status != 200 {
		t.Errorf("handler() returned status %d, want 200", response.Status)
	}
}

func BenchmarkHandler(b *testing.B) {
	ctx := context.Background()
	request := Request{Name: "Benchmark"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = handler(ctx, request)
	}
}

func TestResponse(t *testing.T) {
	tests := []struct {
		name     string
		response Response
	}{
		{
			name: "valid response structure",
			response: Response{
				Message: "Test message",
				Status:  200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.response.Message == "" {
				t.Error("Response Message should not be empty")
			}
			if tt.response.Status != 200 {
				t.Errorf("Response Status = %v, want 200", tt.response.Status)
			}
		})
	}
}
