package cmd

import (
	"github.com/spf13/cobra"
)

//var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cache-projects",
	Short: "Cache Insights Projects",
	Long:  "Caches insights projects to s3 ",
	Run: func(cmd *cobra.Command, args []string) {
		Handler()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
