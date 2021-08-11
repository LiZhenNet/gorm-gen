package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gorm-gen",
	Short: "gorm-gen is a generator for gorm v2",
	Long:  `gorm-gen is a generator for gorm v2`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
