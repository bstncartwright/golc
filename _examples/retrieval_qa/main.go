package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hupe1980/golc/chain"
	"github.com/hupe1980/golc/llm"
	"github.com/hupe1980/golc/schema"
)

type mockRetriever struct{}

func (r *mockRetriever) GetRelevantDocuments(ctx context.Context, query string) ([]schema.Document, error) {
	return []schema.Document{
		{PageContent: "Why don't scientists trust atoms? Because they make up everything!"},
		{PageContent: "Why did the bicycle fall over? Because it was two-tired!"},
	}, nil
}

func main() {
	openai, err := llm.NewOpenAI(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	retrievalQAChain, err := chain.NewRetrievalQAFromLLM(openai, &mockRetriever{})
	if err != nil {
		log.Fatal(err)
	}

	result, err := retrievalQAChain.Run(context.Background(), "Why don't scientists trust atoms?")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
