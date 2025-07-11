package dllama

import (
	"encoding/json"
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

type PromptParams struct {
	Seed             *int     `json:"seed,omitempty"`
	Temperature      *float32 `json:"temperature,omitempty"`
	DynatempRange    *float32 `json:"dynatempRange,omitempty"`
	DynatempExponent *float32 `json:"dynatempExponent,omitempty"`
	TopK             *int     `json:"topK,omitempty"`
	TopP             *float32 `json:"topP,omitempty"`
	MinP             *float32 `json:"minP,omitempty"`
	XtcProbability   *float32 `json:"xtcProbability,omitempty"`
	XtcThreshold     *float32 `json:"xtcThreshold,omitempty"`
	Mirostat         *int     `json:"mirostat,omitempty"`
	MirostatTau      *float32 `json:"mirostatTau,omitempty"`
	MirostatEta      *float32 `json:"mirostatEta,omitempty"`
	RepeatPenalty    *float32 `json:"repeatPenalty,omitempty"`
	PresencePenalty  *float32 `json:"presencePenalty,omitempty"`
	FrequencyPenalty *float32 `json:"frequencyPenalty,omitempty"`
	DryMultiplier    *float32 `json:"dryMultiplier,omitempty"`
	DryBase          *float32 `json:"dryBase,omitempty"`
	DryAllowedLength *int     `json:"dryAllowedLength,omitempty"`
	DryPenaltyLastN  *int     `json:"dryPenaltyLastN,omitempty"`
	MaxTokens        *int     `json:"maxTokens,omitempty"`
	MinKeep          *int     `json:"minKeep,omitempty"`
	IgnoreEos        *bool    `json:"ignoreEos,omitempty"`
	NoPerf           *bool    `json:"noPerf,omitempty"`
	TimingPerToken   *bool    `json:"timingPerToken,omitempty"`
}

type PromptParam func(*PromptParams) error

func (p *PromptParams) Json() ([]byte, error) {
	return json.Marshal(p)
}

func ParsePromptParams(jsonBytes []byte) (*PromptParams, error) {
	p := &PromptParams{}
	if err := json.Unmarshal(jsonBytes, p); err != nil {
		return nil, err
	}
	return p, nil
}

func WithSeed(seed int) PromptParam {
	return func(p *PromptParams) error {
		p.Seed = &seed
		return nil
	}
}

func WithTemperature(temperature float32) PromptParam {
	return func(p *PromptParams) error {
		p.Temperature = &temperature
		return nil
	}
}

func WithDynatempRange(dynatempRange float32) PromptParam {
	return func(p *PromptParams) error {
		p.DynatempRange = &dynatempRange
		return nil
	}
}

func WithDynatempExponent(dynatempExponent float32) PromptParam {
	return func(p *PromptParams) error {
		p.DynatempExponent = &dynatempExponent
		return nil
	}
}

func WithTopK(topK int) PromptParam {
	return func(p *PromptParams) error {
		p.TopK = &topK
		return nil
	}
}

func WithTopP(topP float32) PromptParam {
	return func(p *PromptParams) error {
		p.TopP = &topP
		return nil
	}
}

func WithMinP(minP float32) PromptParam {
	return func(p *PromptParams) error {
		p.MinP = &minP
		return nil
	}
}

func WithXtcProbability(xtcProbability float32) PromptParam {
	return func(p *PromptParams) error {
		p.XtcProbability = &xtcProbability
		return nil
	}
}

func WithXtcThreshold(xtcThreshold float32) PromptParam {
	return func(p *PromptParams) error {
		p.XtcThreshold = &xtcThreshold
		return nil
	}
}

func WithMirostat(mirostat int) PromptParam {
	return func(p *PromptParams) error {
		p.Mirostat = &mirostat
		return nil
	}
}

func WithMirostatTau(mirostatTau float32) PromptParam {
	return func(p *PromptParams) error {
		p.MirostatTau = &mirostatTau
		return nil
	}
}

func WithMirostatEta(mirostatEta float32) PromptParam {
	return func(p *PromptParams) error {
		p.MirostatEta = &mirostatEta
		return nil
	}
}

func WithRepeatPenalty(repeatPenalty float32) PromptParam {
	return func(p *PromptParams) error {
		p.RepeatPenalty = &repeatPenalty
		return nil
	}
}

func WithPresencePenalty(presencePenalty float32) PromptParam {
	return func(p *PromptParams) error {
		p.PresencePenalty = &presencePenalty
		return nil
	}
}

func WithFrequencyPenalty(frequencyPenalty float32) PromptParam {
	return func(p *PromptParams) error {
		p.FrequencyPenalty = &frequencyPenalty
		return nil
	}
}

func WithDryMultiplier(dryMultiplier float32) PromptParam {
	return func(p *PromptParams) error {
		p.DryMultiplier = &dryMultiplier
		return nil
	}
}

func WithDryBase(dryBase float32) PromptParam {
	return func(p *PromptParams) error {
		p.DryBase = &dryBase
		return nil
	}
}

func WithDryAllowedLength(dryAllowedLength int) PromptParam {
	return func(p *PromptParams) error {
		p.DryAllowedLength = &dryAllowedLength
		return nil
	}
}

func WithDryPenaltyLastN(dryPenaltyLastN int) PromptParam {
	return func(p *PromptParams) error {
		p.DryPenaltyLastN = &dryPenaltyLastN
		return nil
	}
}

func WithMaxTokens(maxTokens int) PromptParam {
	return func(p *PromptParams) error {
		p.MaxTokens = &maxTokens
		return nil
	}
}

func WithMinKeep(minKeep int) PromptParam {
	return func(p *PromptParams) error {
		p.MinKeep = &minKeep
		return nil
	}
}

func WithIgnoreEos(ignoreEos bool) PromptParam {
	return func(p *PromptParams) error {
		p.IgnoreEos = &ignoreEos
		return nil
	}
}

func WithNoPerf(noPerf bool) PromptParam {
	return func(p *PromptParams) error {
		p.NoPerf = &noPerf
		return nil
	}
}

func WithTimingPerToken(timingPerToken bool) PromptParam {
	return func(p *PromptParams) error {
		p.TimingPerToken = &timingPerToken
		return nil
	}
}

func WithParams(params *PromptParams) PromptParam {
	return func(p *PromptParams) error {
		*p = *params
		return nil
	}
}

func Generate(modelName string, promptStr string, params ...PromptParam) (Prompt, error) {
	modelNameBytes := []byte(modelName)
	promptBytes := []byte(promptStr)

	pparams := &PromptParams{}
	for _, param := range params {
		if err := param(pparams); err != nil {
			return nil, err
		}
	}

	paramsJsonBytes, err := pparams.Json()
	if err != nil {
		return nil, err
	}

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
