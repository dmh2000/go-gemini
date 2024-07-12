package main

import (
	"log"
	"os"

	"sqirvy.xyz/go-gemini/cmd"
)

func main() {
	log.SetOutput(os.Stderr)
	cmd.Execute()
}
