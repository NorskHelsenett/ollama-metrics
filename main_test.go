package main

import (
	"encoding/json"
	"testing"
)

func TestParseOpenAIChatCompletionsUsageNonStreaming(t *testing.T) {
	model, prompt, completion := parseOpenAIChatCompletionsUsage([]byte(`{
  "model": "gemma4:e4b",
  "usage": {"prompt_tokens": 42, "completion_tokens": 7}
}`))
	if model != "gemma4:e4b" || prompt != 42 || completion != 7 {
		t.Fatalf("got model=%q prompt=%d completion=%d", model, prompt, completion)
	}
}

func TestParseOpenAIChatCompletionsUsageStreaming(t *testing.T) {
	model, prompt, completion := parseOpenAIChatCompletionsUsage([]byte("data: {\"model\":\"gemma4:e4b\",\"choices\":[{\"delta\":{\"content\":\"Hi\"}}]}\n\n" +
		"data: {\"model\":\"gemma4:e4b\",\"choices\":[],\"usage\":{\"prompt_tokens\":42,\"completion_tokens\":7}}\n\n" +
		"data: [DONE]\n"))
	if model != "gemma4:e4b" || prompt != 42 || completion != 7 {
		t.Fatalf("got model=%q prompt=%d completion=%d", model, prompt, completion)
	}
}

func TestIncludeOpenAIStreamingUsage(t *testing.T) {
	got := includeOpenAIStreamingUsage([]byte(`{"model":"gemma4:e4b","stream":true,"messages":[]}`))
	var request map[string]interface{}
	if err := json.Unmarshal(got, &request); err != nil {
		t.Fatal(err)
	}
	options, ok := request["stream_options"].(map[string]interface{})
	if !ok || options["include_usage"] != true {
		t.Fatalf("stream_options = %#v, want include_usage=true", request["stream_options"])
	}
}

func TestIncludeOpenAIStreamingUsageLeavesNonStreamingRequestAlone(t *testing.T) {
	input := []byte(`{"model":"gemma4:e4b","stream":false,"messages":[]}`)
	if got := string(includeOpenAIStreamingUsage(input)); got != string(input) {
		t.Fatalf("got %s, want request unchanged", got)
	}
}
