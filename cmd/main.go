package main

import (
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(restApi)

	err := rootCmd.Execute()
	collection.PanicIfErr(err)
}
