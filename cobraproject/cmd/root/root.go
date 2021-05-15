package root

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

)

func NewRootCommand() cobra.Command {
	rootCMd := cobra.Command{
		Use:   "app [sub]",
		Short: "My app command",
		// PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 	fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
		// },
		// PreRun: func(cmd *cobra.Command, args []string) {
		// 	fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
		// },
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Printf("Inside rootCmd Run with args: %v\n", args)
		// },
		// PostRun: func(cmd *cobra.Command, args []string) {
		// 	fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
		// },
		// PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 	fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
		// },
	}
	var cmdHTTPServer = &cobra.Command{
			Run:
		}
	// rootCMd.SetArgs
	rootCMd.AddCommand()
	return rootCMd
}

