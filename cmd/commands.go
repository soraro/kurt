package cmd

import (
	"github.com/spf13/cobra"
)

var cmdNamespaces = &cobra.Command{
	Use:   "namespaces",
	Short: "Only print namespace-wide restart counts",
	Long:  "Only print namespace-wide restart counts",
	Run: func(cmd *cobra.Command, args []string) {
		printNS = true
		printAll = false
		clientset := auth()
		collect(clientset, inamespace, ilabels)
	},
}

var cmdPods = &cobra.Command{
	Use:   "pods",
	Short: "Only print pod restart counts",
	Long:  "Only print pod restart counts",
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

func init() {
	rootCmd.AddCommand(cmdNamespaces)
	rootCmd.AddCommand(cmdPods)
	rootCmd.AddCommand(cmdLabels)
	rootCmd.AddCommand(cmdAll)
}
