package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kurt/internal/version"
)

var cmdNamespaces = &cobra.Command{
	Use:     "namespaces",
	Short:   "Only print namespace-wide restart counts",
	Long:    "Only print namespace-wide restart counts",
	Aliases: []string{"ns"},
	Run: func(cmd *cobra.Command, args []string) {
		printNS = true
		printAll = false
		clientset := auth()
		collect(clientset, inamespace, ilabels)
	},
}

var cmdNodes = &cobra.Command{
	Use:     "nodes",
	Short:   "Only print node restart counts",
	Long:    "Only print node restart counts",
	Aliases: []string{"no", "node"},
	Run: func(cmd *cobra.Command, args []string) {
		printNode = true
		printAll = false
		clientset := auth()
		collect(clientset, inamespace, ilabels)
	},
}

var cmdPods = &cobra.Command{
	Use:     "pods",
	Short:   "Only print pod restart counts",
	Long:    "Only print pod restart counts",
	Aliases: []string{"po"},
	Run: func(cmd *cobra.Command, args []string) {
		printPods = true
		printAll = false
		clientset := auth()
		collect(clientset, inamespace, ilabels)
	},
}

var cmdLabels = &cobra.Command{
	Use:   "labels",
	Short: "Only print restart counts grouped by labels",
	Long:  "Only print restart counts grouped by labels",
	Run: func(cmd *cobra.Command, args []string) {
		printLabel = true
		printAll = false
		clientset := auth()
		collect(clientset, inamespace, ilabels)
	},
}

var cmdAll = &cobra.Command{
	Use:   "all",
	Short: "Print all groupings collected by kurt!",
	Long:  "Print all groupings collected by kurt!",
	Run: func(cmd *cobra.Command, args []string) {
		printAll = true
		clientset := auth()
		collect(clientset, inamespace, ilabels)
	},
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the current version and exit",
	Long:  `Print the current version and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kurt:  %s\n", version.Version())
	},
}

func init() {
	rootCmd.AddCommand(cmdNamespaces)
	rootCmd.AddCommand(cmdNodes)
	rootCmd.AddCommand(cmdPods)
	rootCmd.AddCommand(cmdLabels)
	rootCmd.AddCommand(cmdAll)
	rootCmd.AddCommand(cmdVersion)
}
