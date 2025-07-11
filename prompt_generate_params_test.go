//go:build !wasi && !wasm
// +build !wasi,!wasm

package dllama

import (
	"encoding/json"
	"reflect"
	"testing"
	"unsafe"

	"github.com/taubyte/dllama-go/symbols"
)

// ptrToString converts a memory region starting at ptr of length len bytes into a Go string.
func ptrToString(ptr *byte, length uint32) string {
	bytes := make([]byte, length)
	for i := 0; i < int(length); i++ {
		bytes[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i)))
	}
	return string(bytes)
}

func TestGenerate_WithParams(t *testing.T) {
	origNewPrompt := symbols.NewPrompt
	defer func() { symbols.NewPrompt = origNewPrompt }()

	var (
		gotModelName string
		gotPrompt    string
		gotParams    map[string]any
	)

	symbols.NewPrompt = func(
		modelNamePtr *byte, modelNameLen uint32,
		promptPtr *byte, promptLen uint32,
		paramsJsonPtr *byte, paramsJsonLen uint32,
		errBufPtr *byte, errLenPtr *uint32,
	) int64 {
		gotModelName = ptrToString(modelNamePtr, modelNameLen)
		gotPrompt = ptrToString(promptPtr, promptLen)
		jsonStr := ptrToString(paramsJsonPtr, paramsJsonLen)
		if err := json.Unmarshal([]byte(jsonStr), &gotParams); err != nil {
			t.Fatalf("failed to unmarshal params JSON: %v", err)
		}
		return 99 // arbitrary positive id
	}

	const (
		modelName = "gpt-neo"
		promptStr = "Hello"
	)

	_, err := Generate(
		modelName,
		promptStr,
		WithMaxTokens(100),
		WithTemperature(0.8),
		WithIgnoreEos(true),
	)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if gotModelName != modelName {
		t.Fatalf("modelName mismatch: expected %q, got %q", modelName, gotModelName)
	}
	if gotPrompt != promptStr {
		t.Fatalf("prompt mismatch: expected %q, got %q", promptStr, gotPrompt)
	}

	expectedParams := map[string]any{
		"maxTokens":   float64(100), // numbers decode as float64
		"temperature": 0.8,
		"ignoreEos":   true,
	}

	for k, v := range expectedParams {
		if !reflect.DeepEqual(gotParams[k], v) {
			t.Fatalf("param %q mismatch: expected %v, got %v", k, v, gotParams[k])
		}
	}
}
