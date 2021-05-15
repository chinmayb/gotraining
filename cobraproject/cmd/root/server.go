package root

import "github.com/spf13/cobra"

func NewServer() {

}


func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
}
