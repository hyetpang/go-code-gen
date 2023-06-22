package strategy

import (
	"bufio"
	"bytes"
	"go-code-gen/pkg/common"
	"go-code-gen/pkg/platform/newline"
	"io"
	"log"
	"os"
	"strings"
)

// 为生成的service以及handler结构体放到fx中注册
func structRegister(isHandler bool, providerFile, structName string) error {
	file, err := os.Open(providerFile)
	if err != nil {
		log.Fatalf("打开文件(%s)出错:%s", providerFile, err.Error())
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)
	var newFile bytes.Buffer
	var isFxOption bool
	var isStruct bool
	suffix := "Service"
	if isHandler {
		suffix = "Handler"
	}
	for {
		var isWrite = true
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("读取文件(%s)出错:%s", providerFile, err.Error())
		}
		lineString := string(line)
		if !isFxOption && strings.Contains(lineString, "[]fx.Option{") {
			isFxOption = true
		}
		if isFxOption && strings.TrimSpace(lineString) == "}" {
			// fileReader
			_, err = newFile.WriteString("		fx.Provide(new" + structName + suffix + ")," + string(newline.NewLine) + "	}" + string(newline.NewLine))
			if err != nil {
				log.Fatal("文件内容写入出错:", err.Error())
			}
			isWrite = false
			isFxOption = false
		}
		if isHandler && strings.Contains(lineString, "struct") {
			isStruct = true
		}
		if isHandler && isStruct && strings.TrimSpace(lineString) == "}" {
			_, err = newFile.WriteString("	" + structName + " *" + common.LowerFirst(structName) + suffix + string(newline.NewLine) + "}" + string(newline.NewLine))
			if err != nil {
				log.Fatal("文件内容写入出错:", err.Error())
			}
			isWrite = false
			isStruct = false
		}
		if isWrite {
			_, err = newFile.WriteString(lineString + string(newline.NewLine))
			if err != nil {
				log.Fatal("文件内容写入出错:", err.Error())
			}
		}
	}
	newProviderFile, err := os.Create(providerFile)
	if err != nil {
		log.Fatalf("创建文件(%s)出错:%s", providerFile, err.Error())
	}
	_, err = newProviderFile.WriteString(newFile.String())
	if err != nil {
		log.Fatalf("写入文件(%s)内容出错:%s", providerFile, err.Error())
	}
	return nil
}
