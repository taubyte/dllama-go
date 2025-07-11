package dllama

import (
	"fmt"
	"time"

	"github.com/taubyte/dllama-go/symbols"
)

type prompt struct {
	id uint64
}

type PromptStats struct {
	ReceivedAt               time.Time
	SubmittedAt              time.Time
	DoneAt                   time.Time
	TokensOutCount           uint64
	TokensProcessingDuration time.Duration
}

type Prompt interface {
	Next() (string, error)
	Stats() (*PromptStats, error)
}

func Generate(modelName string, promptStr string, paramsJson string) (Prompt, error) {
	modelNameBytes := []byte(modelName)
	promptBytes := []byte(promptStr)
	paramsJsonBytes := []byte(paramsJson)

	// error buffer for NewPrompt
	errBuf := make([]byte, 1024)
	var errLen uint32

	promptId := symbols.NewPrompt(
		&modelNameBytes[0], uint32(len(modelNameBytes)),
		&promptBytes[0], uint32(len(promptBytes)),
		&paramsJsonBytes[0], uint32(len(paramsJsonBytes)),
		&errBuf[0], &errLen,
	)

	if promptId < 0 {
		return nil, fmt.Errorf("%s", string(errBuf[:errLen]))
	}

	return &prompt{id: uint64(promptId)}, nil
}

func (p *prompt) Next() (string, error) {
	tokenBuf := make([]byte, 1024)
	errBuf := make([]byte, 1024)
	var errLen uint32

	n := symbols.NextToken(p.id, &tokenBuf[0], &errBuf[0], &errLen)
	if n < 0 {
		return "", fmt.Errorf("%s", string(errBuf[:errLen]))
	}

	if n == 0 {
		return "", nil
	}

	return string(tokenBuf[:n]), nil

}

func (p *prompt) Stats() (*PromptStats, error) {
	var receivedAt, submittedAt, doneAt uint64
	var tokensOutCount, tokensProcessingDuration uint64

	ret := symbols.GetPromptStats(
		p.id,
		&receivedAt,
		&submittedAt,
		&doneAt,
		&tokensOutCount,
		&tokensProcessingDuration,
	)

	if ret != 0 {
		return nil, fmt.Errorf("GetPromptStats returned %d", ret)
	}

	return &PromptStats{
		ReceivedAt:               time.Unix(0, int64(receivedAt)),
		SubmittedAt:              time.Unix(0, int64(submittedAt)),
		DoneAt:                   time.Unix(0, int64(doneAt)),
		TokensOutCount:           tokensOutCount,
		TokensProcessingDuration: time.Duration(tokensProcessingDuration),
	}, nil
}
