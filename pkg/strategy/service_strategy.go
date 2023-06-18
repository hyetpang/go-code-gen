package strategy

import (
	"errors"
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
		return
	}
	fileAppend(serviceFile, "service_func.tmpl", c)
}
