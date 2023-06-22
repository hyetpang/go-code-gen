package strategy

import (
	_ "embed"
	"errors"
	"log"
	"os"
	"path/filepath"

	"go-code-gen/pkg/config"
)

type handlerStrategy struct{}

func (handlerStrategy *handlerStrategy) Gen(c *config.Config) {
	handlerFile := filepath.Join(c.HandlersPath, c.FilePrefix+"_handler.go")
	_, err := os.Stat(handlerFile)
	if errors.Is(err, os.ErrNotExist) {
		// 不存在
		fileCreate(handlerFile, "handler_file.tmpl", c)
		// 需要解析provider.go文件，吧当前handler加入provider
		err = structRegister(true, filepath.Join(c.HandlersPath, "provider.go"), c.DModelName)
		if err != nil {
			log.Fatal("生成结构体注册代码出错:", err.Error())
		}
		return
	}
	fileAppend(handlerFile, "handler_func.tmpl", c)
}
