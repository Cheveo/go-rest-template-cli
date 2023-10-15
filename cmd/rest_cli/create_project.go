package rest_cli

import (
	"fmt"

	"github.com/Cheveo/go-rest-cli/pkg/rest_cli"
	"github.com/spf13/cobra"
)

var domain string

var createProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Creates a whole REST Project from scratch",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			fmt.Println("The name of the domain is required.")
		}
		if modName == "" {
			fmt.Println("The name of the module is required.")
		}
		
		err := rest_cli.CreateProjectFromScratch(domain, modName, directory)
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&domain, "domain", "d", "", "The name of the domain")
	createProjectCmd.Flags().StringVarP(&modName, "goModName", "m", "", "The Go mod name")
	createProjectCmd.Flags().StringVarP(&directory, "directory", "p", "", "The Directory to create the project")

	rootCmd.AddCommand(createProjectCmd)
}
