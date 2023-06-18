/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-code-gen/pkg/conf"
	"go-code-gen/pkg/config"
	"go-code-gen/pkg/strategy"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成模板代码",
	Long:  `生成模板代码`,
	Run: func(cmd *cobra.Command, args []string) {
		configData := conf.ParseConfig()
		configs := config.NewFromMethods(configData.Methods)
		strategy.Runs(configs)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
