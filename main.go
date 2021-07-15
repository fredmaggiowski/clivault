package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fredmaggiowski/clivault/cmd/save"
	"github.com/spf13/cobra"
)

func cmdEntrypoint() {
	rootCmd := &cobra.Command{
		Use: "cliv",
	}
	rootCmd.AddCommand(save.NewSaveCmd())

	ctx := context.TODO()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	cmdEntrypoint()
}
