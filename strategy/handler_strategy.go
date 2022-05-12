package strategy

import (
	"errors"
	"os"
	"path/filepath"

	"go-code-gen/config"
)

type handlerStrategy struct{}

func (handlerStrategy *handlerStrategy) Gen(c *config.Config) {
	handlerFile := filepath.Join(c.HandlersPath, c.FilePrefix+"_handler.go")
	_, err := os.Stat(handlerFile)
	if errors.Is(err, os.ErrNotExist) {
		// 不存在
		fileCreate(handlerFile, "handler_file.tmpl", c)
		return
	}
	fileAppend(handlerFile, "handler_func.tmpl", c)
}
