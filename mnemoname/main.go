package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/matoous/mnemoname"
)

func main() {
	var n int
	rootCmd := &cobra.Command{
		Use:   "mnemoname",
		Short: "Generate mnemonic name(s)",
		Long: `Generate easy to memorable names for whatever you need.

  Based on: https://web.archive.org/web/20090918202746/http://tothink.com/mnemonic/wordlist.html`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if n == 1 {
				cmd.Println(mnemoname.New())
				return nil
			}
			names, err := mnemoname.NewN(n)
			if err != nil {
				return err
			}
			for i := range names {
				cmd.Println(names[i])
			}
			return nil
		},
	}
	rootCmd.LocalFlags().IntVarP(&n, "num", "n", 1, "numer of names to generate")

	err := rootCmd.Execute()
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "%v", err)
	}
}
