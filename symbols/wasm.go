//go:build wasi || wasm
// +build wasi wasm

package symbols

//go:wasm-module dllama
//export list_models
func ListModels() uint64

//go:wasm-module dllama
//export list_models_next
func ListModelsNext(
	listId uint64,
	namePtr *byte,
) int32

//go:wasm-module dllama
//export fetch_hf_model
func FetchHFModel(
	namePtr *byte, nameLen uint32,
	repoPtr *byte, repoLen uint32,
	filePtr *byte, fileLen uint32,
	errorPtr *byte, errorLenPtr *uint32,
) int32

//go:wasm-module dllama
//export new_prompt
func NewPrompt(
	modelNamePtr *byte, modelNameLen uint32,
	promptPtr *byte, promptLen uint32,
	paramsJsonPtr *byte, paramsJsonLen uint32,
	errBufPtr *byte, errBufLen *uint32,
) int64

//go:wasm-module dllama
//export next_token
func NextToken(
	promptId uint64,
	tokenBufferPtr,
	errBufPtr *byte,
	errBufLen *uint32,
) int64

//go:wasm-module dllama
//export get_prompt_stats
func GetPromptStats(
	promptId uint64,
	receivedAtPtr,
	submittedAtPtr,
	doneAtPtr,
	tokensOutCountPtr,
	tokensProcessingDurationPtr *uint64,
) int32
