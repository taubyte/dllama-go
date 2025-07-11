//go:build !wasi && !wasm
// +build !wasi,!wasm

package dllama

import (
	"testing"
	"time"
	"unsafe"

	"github.com/taubyte/dllama-go/symbols"
)

func TestPrompt(t *testing.T) {
	t.Run("GenerateSuccessAndNext", func(t *testing.T) {
		origNewPrompt := symbols.NewPrompt
		origNextToken := symbols.NextToken
		defer func() {
			symbols.NewPrompt = origNewPrompt
			symbols.NextToken = origNextToken
		}()

		symbols.NewPrompt = func(
			modelNamePtr *byte, modelNameLen uint32,
			promptPtr *byte, promptLen uint32,
			paramsJsonPtr *byte, paramsJsonLen uint32,
			errBufPtr *byte, errLenPtr *uint32,
		) int64 {
			return 42
		}

		var call int
		const tokenStr = "Hello"
		symbols.NextToken = func(
			promptId uint64,
			tokenBufPtr *byte,
			errBufPtr *byte,
			errLenPtr *uint32,
		) int64 {
			if call == 0 {
				for i := 0; i < len(tokenStr); i++ {
					*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(tokenBufPtr)) + uintptr(i))) = tokenStr[i]
				}
				call++
				return int64(len(tokenStr))
			}
			return 0
		}

		p, err := Generate("model", "prompt", "{}")
		if err != nil {
			t.Fatalf("Generate failed: %v", err)
		}

		tok, err := p.Next()
		if err != nil {
			t.Fatalf("Next failed: %v", err)
		}
		if tok != tokenStr {
			t.Fatalf("expected %q, got %q", tokenStr, tok)
		}

		tok, err = p.Next()
		if err != nil {
			t.Fatalf("Next second call error: %v", err)
		}
		if tok != "" {
			t.Fatalf("expected empty token, got %q", tok)
		}
	})

	t.Run("GenerateError", func(t *testing.T) {
		origNewPrompt := symbols.NewPrompt
		defer func() { symbols.NewPrompt = origNewPrompt }()

		const errMsg = "bad model"
		symbols.NewPrompt = func(
			modelNamePtr *byte, modelNameLen uint32,
			promptPtr *byte, promptLen uint32,
			paramsJsonPtr *byte, paramsJsonLen uint32,
			errBufPtr *byte, errLenPtr *uint32,
		) int64 {
			for i := 0; i < len(errMsg); i++ {
				*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(errBufPtr)) + uintptr(i))) = errMsg[i]
			}
			*errLenPtr = uint32(len(errMsg))
			return -1
		}

		if _, err := Generate("model", "prompt", "{}"); err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("StatsSuccess", func(t *testing.T) {
		origGetStats := symbols.GetPromptStats
		defer func() { symbols.GetPromptStats = origGetStats }()

		symbols.GetPromptStats = func(
			promptId uint64,
			receivedAtPtr,
			submittedAtPtr,
			doneAtPtr,
			tokensOutCountPtr,
			tokensProcessingDurationPtr *uint64,
		) int32 {
			*receivedAtPtr = uint64(1e9)  // 1s
			*submittedAtPtr = uint64(2e9) // 2s
			*doneAtPtr = uint64(3e9)      // 3s
			*tokensOutCountPtr = 5
			*tokensProcessingDurationPtr = uint64(4e9) // 4s
			return 0
		}

		p := &prompt{id: 55}
		stats, err := p.Stats()
		if err != nil {
			t.Fatalf("Stats failed: %v", err)
		}

		if stats.ReceivedAt != time.Unix(0, 1e9) || stats.SubmittedAt != time.Unix(0, 2e9) || stats.DoneAt != time.Unix(0, 3e9) {
			t.Fatalf("unexpected time conversions: %+v", stats)
		}
		if stats.TokensOutCount != 5 {
			t.Fatalf("expected TokensOutCount 5, got %d", stats.TokensOutCount)
		}
		if stats.TokensProcessingDuration != 4*time.Second {
			t.Fatalf("expected duration 4s, got %v", stats.TokensProcessingDuration)
		}
	})

	t.Run("StatsError", func(t *testing.T) {
		origGetStats := symbols.GetPromptStats
		defer func() { symbols.GetPromptStats = origGetStats }()

		symbols.GetPromptStats = func(
			promptId uint64,
			receivedAtPtr,
			submittedAtPtr,
			doneAtPtr,
			tokensOutCountPtr,
			tokensProcessingDurationPtr *uint64,
		) int32 {
			return -1
		}

		p := &prompt{id: 1}
		if _, err := p.Stats(); err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
