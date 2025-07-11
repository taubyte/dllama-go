//go:build !wasi && !wasm
// +build !wasi,!wasm

package dllama

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/taubyte/dllama-go/symbols"
)

// copyStringToPtr writes the contents of s into the memory starting at ptr.
// It is the caller's responsibility to ensure that the memory region is large
// enough to hold the contents of s.
func copyStringToPtr(ptr *byte, s string) {
	for i := 0; i < len(s); i++ {
		*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i))) = s[i]
	}
}

func TestModels(t *testing.T) {
	t.Run("ListModels", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			origListModels := symbols.ListModels
			origListModelsNext := symbols.ListModelsNext
			defer func() {
				symbols.ListModels = origListModels
				symbols.ListModelsNext = origListModelsNext
			}()

			expected := []string{"alpaca", "llama", "dolly"}

			symbols.ListModels = func() uint64 { return 1 }
			var idx int
			symbols.ListModelsNext = func(listId uint64, namePtr *byte) int32 {
				if idx >= len(expected) {
					return 0
				}
				s := expected[idx]
				idx++
				copyStringToPtr(namePtr, s)
				return int32(len(s))
			}

			got, err := ListModels()
			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("expected %v, got %v", expected, got)
			}
		})

		t.Run("ErrorOnListId", func(t *testing.T) {
			origListModels := symbols.ListModels
			defer func() { symbols.ListModels = origListModels }()
			symbols.ListModels = func() uint64 { return ^uint64(0) }
			if _, err := ListModels(); err == nil {
				t.Fatalf("expected error, got nil")
			}
		})

		t.Run("ErrorOnNext", func(t *testing.T) {
			origListModels := symbols.ListModels
			origListModelsNext := symbols.ListModelsNext
			defer func() {
				symbols.ListModels = origListModels
				symbols.ListModelsNext = origListModelsNext
			}()
			symbols.ListModels = func() uint64 { return 1 }
			symbols.ListModelsNext = func(listId uint64, namePtr *byte) int32 { return -5 }
			if _, err := ListModels(); err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	})

	t.Run("FetchHuggingFaceModel", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			orig := symbols.FetchHFModel
			defer func() { symbols.FetchHFModel = orig }()
			symbols.FetchHFModel = func(
				namePtr *byte, nameLen uint32,
				repoPtr *byte, repoLen uint32,
				filePtr *byte, fileLen uint32,
				errBufPtr *byte, errLenPtr *uint32,
			) int32 {
				return 0
			}
			if err := FetchHuggingFaceModel("model", "repo", "file"); err != nil {
				t.Fatalf("expected success, got error %v", err)
			}
		})

		t.Run("Error", func(t *testing.T) {
			orig := symbols.FetchHFModel
			defer func() { symbols.FetchHFModel = orig }()
			const errMsg = "network error"
			symbols.FetchHFModel = func(
				namePtr *byte, nameLen uint32,
				repoPtr *byte, repoLen uint32,
				filePtr *byte, fileLen uint32,
				errBufPtr *byte, errLenPtr *uint32,
			) int32 {
				copyStringToPtr(errBufPtr, errMsg)
				*errLenPtr = uint32(len(errMsg))
				return -1
			}
			if err := FetchHuggingFaceModel("model", "repo", "file"); err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	})
}
