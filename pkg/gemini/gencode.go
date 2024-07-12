package gemini

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/vertexai/genai"
)

// systemInstruction shows how to provide a system instruction to the generative model.
func GenCode(w io.Writer, instruction, prompt, projectID, location, modelName string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return "", fmt.Errorf("unable to create client: %w", err)
	}
	defer client.Close()

	// The System Instruction is set at model creation
	model := client.GenerativeModel(modelName)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(instruction)},
	}

	// config := genai.GenerationConfig{}
	// config.ResponseMIMEType = "application/json"
	// model.GenerationConfig = config
	res, err := model.GenerateContent(ctx, genai.Text(prompt))
	if res == nil {
		return "", fmt.Errorf("unable to generate contents: %w", err)
	}
	if err != nil {
		return "", fmt.Errorf("unable to generate contents: %w", err)
	}
	if len(res.Candidates) == 0 ||
		res.Candidates[0].Content == nil ||
		len(res.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("empty response from model")
	}

	s := fmt.Sprintf("%s", res.Candidates[0].Content.Parts[0])

	return s, nil
}
