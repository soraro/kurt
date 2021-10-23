package cmd

import (
	"github.com/spf13/cobra"
)

var inamespace []string
var ilabels []string
var limitFlag int

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

}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
