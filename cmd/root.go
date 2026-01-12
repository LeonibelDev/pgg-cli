/*
Copyright © 2026 LeonibelDev <leonibel.ramirez@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/leonibeldev/pgg-cli/internal/crypt"
	"github.com/leonibeldev/pgg-cli/internal/flags"
	"github.com/leonibeldev/pgg-cli/internal/utils"
	"github.com/spf13/cobra"
)

var length int
var types []string
var save string
var copy bool
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pgg-cli",
	Short: "Generate secure password",
	Long:  `Generate password based on crypt/rand, you can export all password in a file or just one password`,
	Run: func(cmd *cobra.Command, args []string) {

		/*
			Generate password with (-l, -t)
			GenPassword need two params length and types on a plain string
		*/
		types, err := flags.Types(types)
		if err != nil {
			fmt.Println(err)
		}

		password, _ := crypt.GenPassword(length, types.String())

		if verbose && !copy {
			fmt.Printf("Password: %v\n", password)
		} else if !verbose && !copy {
			fmt.Print(password)
		}

		// Save Passsword
		utils.Save(save, password)

		// Copy to clipboard
		utils.Copy(copy, password)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// flags
	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of password")
	rootCmd.Flags().StringSliceVarP(&types, "types", "t", []string{"numbers", "upper", "lower", "special"}, "Type of password, valid options: lower, upper, numbers, special")
	rootCmd.Flags().StringVarP(&save, "save", "s", "", "Save password in your db")
	rootCmd.Flags().BoolVarP(&copy, "copy", "c", false, "Copy password to clipboard")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

}
