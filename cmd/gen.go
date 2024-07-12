/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sqirvy.xyz/go-gemini/gemini"
)

var systemTemplate string = `
You are an experienced software developer.
You will receive a description of an algorithm or computer program.
you will respond with an implementation of the algorithm or program using the %v programming language.
if the specified programming language is not supported, you will report an error.
if no algorithm or program is known, you will report an error.
add comments explaining the algorithm and the implementation.'
`

var prompt string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate code",
	Long:  `Use the Google Vertex AI API to generate code.`,
	Run: func(cmd *cobra.Command, args []string) {
		// all config parameters and flags are in viper

		system := fmt.Sprintf(systemTemplate, viper.GetString("language"))
		log.Println(viper.GetString("project"))
		log.Println(viper.GetString("location"))
		log.Println(viper.GetString("model"))
		log.Println(viper.GetString("language"))
		log.Println(system)
		log.Println(prompt)

		code, err := gemini.GenCode(os.Stdout,
			system, prompt,
			viper.GetString("project"),
			viper.GetString("location"),
			viper.GetString("model"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(code)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringP("id", "i", "", "The Google Cloud project ID")
	genCmd.Flags().StringP("location", "l", "", "The Google Cloud server location")
	genCmd.Flags().StringP("model", "m", "", "The Google AI model name")
	genCmd.Flags().StringP("language", "g", "", "The programming language for the output of the model")

	genCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "The prompt for the generative model (required)")
	genCmd.MarkFlagRequired("prompt")

	// override config file with flags.
	viper.BindPFlag("project", genCmd.Flags().Lookup("id"))
	viper.BindPFlag("location", genCmd.Flags().Lookup("location"))
	viper.BindPFlag("model", genCmd.Flags().Lookup("model"))
	viper.BindPFlag("language", genCmd.Flags().Lookup("language"))
}
