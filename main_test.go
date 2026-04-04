package main

import (
	"bytes"
	"testing"
)

func TestFixDoneReason(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "numeric done_reason is converted to string",
			input:    []byte(`{"done_reason":0}`),
			expected: []byte(`{"done_reason":"0"}`),
		},
		{
			name:     "multi-digit numeric done_reason",
			input:    []byte(`{"done_reason":123}`),
			expected: []byte(`{"done_reason":"123"}`),
		},
		{
			name:     "string done_reason is unchanged",
			input:    []byte(`{"done_reason":"stop"}`),
			expected: []byte(`{"done_reason":"stop"}`),
		},
		{
			name:     "no done_reason field",
			input:    []byte(`{"model":"llama3","done":true}`),
			expected: []byte(`{"model":"llama3","done":true}`),
		},
		{
			name:     "empty input",
			input:    []byte(``),
			expected: []byte(``),
		},
		{
			name:     "multiple fields with numeric done_reason",
			input:    []byte(`{"model":"llama3","done":true,"done_reason":1,"eval_count":42}`),
			expected: []byte(`{"model":"llama3","done":true,"done_reason":"1","eval_count":42}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fixDoneReason(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("fixDoneReason(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestEnsureModelTag(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string returns empty",
			input:    "",
			expected: "",
		},
		{
			name:     "model without tag gets :latest",
			input:    "llama3",
			expected: "llama3:latest",
		},
		{
			name:     "model with tag is unchanged",
			input:    "llama3:8b",
			expected: "llama3:8b",
		},
		{
			name:     "model with :latest is unchanged",
			input:    "llama3:latest",
			expected: "llama3:latest",
		},
		{
			name:     "namespaced model without tag gets :latest",
			input:    "library/llama3",
			expected: "library/llama3:latest",
		},
		{
			name:     "namespaced model with tag is unchanged",
			input:    "library/llama3:8b",
			expected: "library/llama3:8b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ensureModelTag(tt.input)
			if result != tt.expected {
				t.Errorf("ensureModelTag(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
