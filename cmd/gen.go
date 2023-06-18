/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		confFile, err := cmd.Flags().GetString(genFileConfFlag)
		if err != nil {
			fmt.Println("获取指定的配置文件路径出错:", err.Error())
			return
		}
		if len(confFile) <= 0 {
			confFile = "./file_gen.conf"
		}
		configData := conf.ParseConfig(confFile)
		configs := config.NewFromMethods(configData.Methods)
		strategy.Runs(configs)
	},
}

const genFileConfFlag = "conf"

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringP(genFileConfFlag, "c", "", "指定配置文件路径")
}
