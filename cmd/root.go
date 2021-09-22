package cmd

import (
	"github.com/spf13/cobra"
)

var inamespace []string
var ilabels []string

var rootCmd = &cobra.Command{
	Use:   "kurt",
	Short: "KUbernetes Restart Tracker",
	Long: `kurt: KUbernetes Restart Tracker

A restart tracker that gives context to what is restarting in your cluster
`,
}

func init() {

	rootCmd.PersistentFlags().StringSliceVarP(&inamespace, "namespace", "n", []string{""}, "Specify namespace for kurt to collect restart metrics.\nLeave blank to collect in all namespaces.")
	rootCmd.PersistentFlags().StringSliceVarP(&ilabels, "label", "l", []string{""}, "Specify multiple times for the label keys you want to see.\nFor example: -l app")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
