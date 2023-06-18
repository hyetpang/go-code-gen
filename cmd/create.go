/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"go-code-gen/pkg/common"
	"go-code-gen/pkg/consts"
	"go-code-gen/pkg/platform/newline"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "创建Go项目目录",
	Long: `创建Go项目目录,创建好的项目目录如下:


	`,
	Run: func(cmd *cobra.Command, args []string) {
		createDirString, _ := cmd.Flags().GetString(createDirFlagName)
		if len(createDirString) <= 0 {
			execFile, err := os.Getwd()
			if err != nil {
				fmt.Println("获取程序可执行路径出错:", err.Error())
				return
			}
			createDirString = execFile
		}
		createRepoNameString, _ := cmd.Flags().GetString(createRepoNameFlagName)
		if len(createRepoNameString) <= 0 {
			fmt.Println("请指定生成目录的仓库名字")
			return
		}
		// cmd/gorm 目录
		common.ExitIfErr(createDir(filepath.Join(createDirString, "cmd/gorm")))
		// cmd/gorm/main.go 文件
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "cmd/gorm/main.go"), consts.GormFile, createRepoNameString))
		// config 目录
		common.ExitIfErr(createDir(filepath.Join(createDirString, "config")))
		// config/app.toml 文件
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "config/app.toml"), consts.AppToml, nil))
		// logic/router
		common.ExitIfErr(createDir(filepath.Join(createDirString, "logic/routers")))
		// logic/router/router.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "logic/routers/routers.go"), consts.RouterFile, createRepoNameString))
		// logic/service
		common.ExitIfErr(createDir(filepath.Join(createDirString, "logic/services")))
		// logic/services/provider.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "logic/services/provider.go"), consts.ServicesProvider, nil))
		// logic/handlers
		common.ExitIfErr(createDir(filepath.Join(createDirString, "logic/handlers")))
		// logic/handlers/provider.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "logic/handlers/provider.go"), consts.HandlersProvider, nil))
		// pkg/common
		common.ExitIfErr(createDir(filepath.Join(createDirString, "pkg/common")))
		// pkg/common/codes.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "pkg/common/codes.go"), consts.CodesFile, nil))
		// pkg/config
		common.ExitIfErr(createDir(filepath.Join(createDirString, "pkg/config")))
		// pkg/config/provider.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "pkg/config/provider.go"), consts.ConfigProvider, nil))
		// pkg/config/config.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "pkg/config/config.go"), consts.ConfigFile, nil))
		// pkg/consts
		common.ExitIfErr(createDir(filepath.Join(createDirString, "pkg/consts")))
		// pkg/consts/consts.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "pkg/consts/consts.go"), consts.ConstsFile, nil))
		// pkg/orm/models_register
		common.ExitIfErr(createDir(filepath.Join(createDirString, "pkg/orm/models_register")))
		// pkg/orm/models_register/inject.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "pkg/orm/models_register/inject.go"), consts.ModelsRegisterInject, createRepoNameString))
		// pkg/orm/models
		common.ExitIfErr(createDir(filepath.Join(createDirString, "pkg/orm/models")))
		// pkg/orm/models/user.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "pkg/orm/models/user.go"), consts.ModelsUserFile, nil))
		// ./gitignore
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, ".gitignore"), consts.GitIgnoreFile, nil))
		// ./api.rest
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "api.rest"), consts.ApiRestFile, nil))
		// ./main.go
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "main.go"), consts.MainFile, createRepoNameString))
		// Makefile
		common.ExitIfErr(createAndWriteFile(filepath.Join(createDirString, "Makefile"), consts.MakeFile, nil))
		// 执行初始化命令
		common.ExitIfErr(execCommand(createDirString, "go", []string{"mod", "init", createRepoNameString}...))
		// 执行tidy命令
		common.ExitIfErr(execCommand(createDirString, "go", []string{"mod", "tidy"}...))
		// 执行生成models命令
		common.ExitIfErr(execCommand(filepath.Join(createDirString,"cmd/gorm"),"go","run","main.go"))
	},
}

const (
	createDirFlagName      = "dir"
	createRepoNameFlagName = "repo_name"
)

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().StringP(createDirFlagName, "d", "", "指定要生成的目录")
	createCmd.Flags().StringP("repo_name", "r", "", "当前仓库的名字")
}

func createDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Println("创建"+dirName+"目录出错:", err.Error())
		return err
	}
	return nil
}

func createAndWriteFile(fileName, content string, data any) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建"+fileName+"文件出错:", err.Error())
		return err
	}
	contentTemp, err := template.New(fileName).Parse(content)
	if err != nil {
		fmt.Println("解析模板内容:"+string(newline.NewLine)+content+string(newline.NewLine)+"出错:", err.Error())
		return err
	}
	if err = contentTemp.Execute(file, data); err != nil {
		fmt.Println("执行模板"+fileName+"出错", err.Error())
		return err
	}
	if err = file.Close(); err != nil {
		fmt.Println("关闭文件"+fileName+"出错", err.Error())
		return err
	}
	return nil
}

func execCommand(wd, command string, args ...string) error {
	initCmd := exec.Command(command, args...)
	initCmd.Dir = wd
	stdout, err := initCmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取命令输出管道出错:", err)
		return err
	}
	if err := initCmd.Start(); err != nil {
		fmt.Println("执行命令出错:", err.Error())
		return err
	}
	initCmd.Stderr = os.Stderr
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString(newline.NewLine)
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	err = initCmd.Wait()
	if err != nil {
		fmt.Println("等待命令执行完成出错:", err.Error())
		return err
	}
	return nil
}
