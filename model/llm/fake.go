package llm

import (
	"context"

	"github.com/hupe1980/golc/schema"
)

// Compile time check to ensure Fake satisfies the LLM interface.
var _ schema.LLM = (*Fake)(nil)

type Fake struct {
	schema.Tokenizer
	response string
}

func NewFake(response string) *Fake {
	return &Fake{
		response: response,
	}
}

func (l *Fake) Generate(ctx context.Context, prompts []string, optFns ...func(o *schema.GenerateOptions)) (*schema.ModelResult, error) {
	return &schema.ModelResult{
		Generations: [][]schema.Generation{{schema.Generation{Text: l.response}}},
		LLMOutput:   map[string]any{},
	}, nil
}

func (l *Fake) Type() string {
	return "llm.Fake"
}

func (l *Fake) Verbose() bool {
	return false
}

func (l *Fake) Callbacks() []schema.Callback {
	return []schema.Callback{}
}

func (l *Fake) InvocationParams() map[string]any {
	return nil
}
