package strategy

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"go-code-gen/pkg/config"
)

type serviceStrategy struct{}

func (serviceStrategy *serviceStrategy) Gen(c *config.Config) {
	serviceFile := filepath.Join(c.ServicesPath, c.FilePrefix+"_service.go")
	_, err := os.Stat(serviceFile)
	if errors.Is(err, os.ErrNotExist) {
		// 不存在
		fileCreate(serviceFile, "service_file.tmpl", c)
		// 需要解析provider.go文件，吧当前handler加入provider
		err = structRegister(true, filepath.Join(c.HandlersPath, "provider.go"), c.DModelName)
		if err != nil {
			log.Fatal("生成结构体注册代码出错:", err.Error())
		}
		return
	}
	fileAppend(serviceFile, "service_func.tmpl", c)
}
