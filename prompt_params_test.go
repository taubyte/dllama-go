//go:build !wasi && !wasm
// +build !wasi,!wasm

package dllama

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/taubyte/dllama-go/symbols"
)

func TestPromptParams_JSON_Roundtrip(t *testing.T) {
	// Build a PromptParams using every available helper so that each helper
	// function is executed at least once.
	params := []PromptParam{
		WithSeed(123),
		WithTemperature(0.5),
		WithDynatempRange(0.2),
		WithDynatempExponent(0.9),
		WithTopK(40),
		WithTopP(0.95),
		WithMinP(0.1),
		WithXtcProbability(0.25),
		WithXtcThreshold(0.3),
		WithMirostat(2),
		WithMirostatTau(5.0),
		WithMirostatEta(0.04),
		WithRepeatPenalty(1.1),
		WithPresencePenalty(0.15),
		WithFrequencyPenalty(0.02),
		WithDryMultiplier(1.2),
		WithDryBase(0.3),
		WithDryAllowedLength(100),
		WithDryPenaltyLastN(64),
		WithMaxTokens(128),
		WithMinKeep(12),
		WithIgnoreEos(true),
		WithNoPerf(true),
		WithTimingPerToken(false),
	}

	pp := &PromptParams{}
	for _, p := range params {
		if err := p(pp); err != nil {
			t.Fatalf("unexpected error applying PromptParam: %v", err)
		}
	}

	jsonBytes, err := pp.Json()
	if err != nil {
		t.Fatalf("PromptParams.Json returned error: %v", err)
	}

	// Ensure the JSON generated is valid.
	if !json.Valid(jsonBytes) {
		t.Fatalf("generated JSON is not valid: %s", string(jsonBytes))
	}

	// Round-trip via ParsePromptParams.
	parsed, err := ParsePromptParams(jsonBytes)
	if err != nil {
		t.Fatalf("ParsePromptParams returned error: %v", err)
	}

	// Simple reflect.DeepEqual comparison works because the fields inside the
	// structs are pointers; the pointed-to values are compared, not the pointer
	// addresses.
	if !reflect.DeepEqual(pp, parsed) {
		t.Fatalf("round-tripped PromptParams mismatch.\noriginal: %+v\nparsed:   %+v", pp, parsed)
	}
}

func TestParsePromptParams_InvalidJSON(t *testing.T) {
	if _, err := ParsePromptParams([]byte("{")); err == nil {
		t.Fatalf("expected error for invalid JSON, got nil")
	}
}

func TestPrompt_Next_ErrorPath(t *testing.T) {
	origNextToken := symbols.NextToken
	defer func() { symbols.NextToken = origNextToken }()

	const errMsg = "next token error"
	symbols.NextToken = func(
		promptId uint64,
		tokenBufPtr *byte,
		errBufPtr *byte,
		errLenPtr *uint32,
	) int64 {
		copyStringToPtr(errBufPtr, errMsg)
		*errLenPtr = uint32(len(errMsg))
		return -1
	}

	p := &prompt{id: 1}
	if _, err := p.Next(); err == nil {
		t.Fatalf("expected error, got nil")
	}
}
