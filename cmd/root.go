package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var inamespace []string
var ilabels []string
var ishowContainers bool
var limitFlag int
var output string

var rootCmd = &cobra.Command{
	Use:   "kurt",
	Short: "KUbernetes Restart Tracker",
	Long: `kurt: KUbernetes Restart Tracker

A restart tracker that gives context to what is restarting in your cluster
`,
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&inamespace, "namespace", "n", []string{""}, "Specify namespace for kurt to collect restart metrics.\nLeave blank to collect in all namespaces.")
	rootCmd.PersistentFlags().StringSliceVarP(&ilabels, "label", "l", []string{""}, "Specify multiple times for the label keys you want to see.\nFor example: \"kurt all -l app\"")
	rootCmd.PersistentFlags().IntVarP(&limitFlag, "limit", "c", 5, "Limit the number of resources you want to see. Set limit to 0 for no limits. Must be positive.\nFor example: \"kurt all -c=10\"")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "standard", "Specify output type. Options are: json, yaml, standard\nFor example: \"kurt all -o json\"")

	// command specific flags
	cmdPods.PersistentFlags().BoolVarP(&ishowContainers, "show-containers", "", false, "Show specific container restart counts for pods\nFor example: \"kurt pods --show-containers\"")
	cmdAll.PersistentFlags().BoolVarP(&ishowContainers, "show-containers", "", false, "Show specific container restart counts for pods\nFor example: \"kurt pods --show-containers\"")

	if strings.HasPrefix(filepath.Base(os.Args[0]), "kubectl-") {
		rootCmd.SetUsageTemplate(strings.NewReplacer(
			"{{.UseLine}}", "kubectl {{.UseLine}}",
			"{{.CommandPath}}", "kubectl {{.CommandPath}}").Replace(rootCmd.UsageTemplate()))
	}

}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
