package gemini

import (
	"io"
	"os"
	"testing"
)

func TestSystemInstruction(t *testing.T) {
	w := io.Writer(os.Stdout)
	projectID := "go-algo-429119"
	location := "us-central1"
	modelName := "gemini-1.5-flash-001"
	instruction := `
		You are an experienced software developer.
		You receive the name of an algorithm and you respond with an implementation
		of the algorithm using the C programming language. 
		`
	prompt := `bubble sort`
	code, err := GenCode(w, instruction, prompt, projectID, location, modelName)
	if err != nil {
		t.Errorf("SystemInstruction() error = %v", err)
	}
	if code == "" {
		t.Errorf("SystemInstruction() code is empty")
	}
	t.Logf("code: %s", code)
}
