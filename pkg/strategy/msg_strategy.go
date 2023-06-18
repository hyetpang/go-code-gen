package strategy

import (
	"errors"
	"os"
	"path/filepath"

	"go-code-gen/pkg/config"
)

type msgStrategy struct{}

func (msgStrategy *msgStrategy) Gen(c *config.Config) {
	msgFile := filepath.Join(c.MsgPath, c.FilePrefix+"_msg.go")
	_, err := os.Stat(msgFile)
	if errors.Is(err, os.ErrNotExist) {
		// 不存在
		fileCreate(msgFile, "msg_file.tmpl", c)
		return
	}
	fileAppend(msgFile, "msg_func.tmpl", c)
}
