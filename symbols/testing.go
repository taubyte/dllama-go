//go:build !wasi && !wasm
// +build !wasi,!wasm

package symbols

var ListModels = func() uint64 {
	return 0
}

var ListModelsNext = func(
	listId uint64,
	namePtr *byte,
) int32 {
	return 0
}

var FetchHFModel = func(
	namePtr *byte, nameLen uint32,
	repoPtr *byte, repoLen uint32,
	filePtr *byte, fileLen uint32,
	errorPtr *byte, errorLenPtr *uint32,
) int32 {
	return 0
}

var NewPrompt = func(
	modelNamePtr *byte, modelNameLen uint32,
	promptPtr *byte, promptLen uint32,
	paramsJsonPtr *byte, paramsJsonLen uint32,
	errBufPtr *byte, errBufLen *uint32,
) int64 {
	return 0
}

var NextToken = func(
	promptId uint64,
	tokenBufferPtr,
	errBufPtr *byte,
	errBufLen *uint32,
) int64 {
	return 0
}

var GetPromptStats = func(
	promptId uint64,
	receivedAtPtr,
	submittedAtPtr,
	doneAtPtr,
	tokensOutCountPtr,
	tokensProcessingDurationPtr *uint64,
) int32 {
	return 0
}
