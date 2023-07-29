package cmd

import (
	"github.com/spf13/cobra"
	"npmw/watcher"
)

var script string
var root string

func init() {
	rootCmd.AddCommand(watch)
	watch.PersistentFlags().StringVarP(&root, "root", "r", ".", "root")
	watch.PersistentFlags().StringVarP(&script, "script", "s", "build", "script")
}

var watch = &cobra.Command{
	Use:   "watch",
	Short: "Start watching.",
	Long:  "Start watching a directory for any file changes.",
	Run: func(cmd *cobra.Command, args []string) {
		watcher.Start(script, root)
	},
}
