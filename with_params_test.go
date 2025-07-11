//go:build !wasi && !wasm
// +build !wasi,!wasm

package dllama

import "testing"

func TestWithParams_Merge(t *testing.T) {
	// Existing params with Seed preset to 1
	seed1 := 1
	base := &PromptParams{Seed: &seed1}

	// Params that should override Seed and add Temperature
	seed2 := 2
	temp := float32(0.7)
	overrides := &PromptParams{Seed: &seed2, Temperature: &temp}

	// Apply WithParams
	if err := WithParams(overrides)(base); err != nil {
		t.Fatalf("WithParams returned error: %v", err)
	}

	// Ensure Seed updated
	if base.Seed == nil || *base.Seed != seed2 {
		t.Fatalf("Seed not overridden correctly, got %+v", base.Seed)
	}

	// Ensure new field set
	if base.Temperature == nil || *base.Temperature != temp {
		t.Fatalf("Temperature not set correctly, got %+v", base.Temperature)
	}

	// Ensure an untouched field remains nil
	if base.MaxTokens != nil {
		t.Fatalf("unexpected overwrite of unrelated field MaxTokens: %v", *base.MaxTokens)
	}
}
