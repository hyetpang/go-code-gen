package main

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Content struct {
	MethodName     string // 方法名
	ModelName      string // 模型名字
	DModelName     string // 首字母小写的模型名字
	ServicesPath   string // 生成的services路径
	HandlersPath   string // 生成的handlers路径
	MsgPath        string // 生成的msg路径
	DependencyName string
	ParamType      string
	RepoName       string
	DocUrlMethod   string //
	DocUrl         string //
	DocTag         string
	DocDesc        string
}

const (
	repoName       = "github.com/HyetPang/go-code-gen"
	servicesPath   = "../logic/services"
	handlersPath   = "../logic/handlers"
	msgPath        = "../logic/msg"
	dependencyName = "github.com/HyetPang/go-frame"
	templateDir    = "./templates"
)

func newContent(methodName, modelName, urlMethod, url, tag, desc string) *Content {
	paramType := "body"
	if urlMethod == http.MethodGet {
		paramType = "query"
	}
	return &Content{
		MethodName:     cases.Title(language.English).String(methodName),
		ModelName:      cases.Title(language.English).String(modelName),
		DModelName:     modelName,
		ServicesPath:   servicesPath,
		HandlersPath:   handlersPath,
		MsgPath:        msgPath,
		RepoName:       repoName,
		DependencyName: dependencyName,
		DocUrlMethod:   urlMethod,
		ParamType:      paramType,
		DocUrl:         url,
		DocTag:         tag,
		DocDesc:        desc,
	}
}

var temps *template.Template

func main() {
	temps = template.Must(template.ParseGlob("./templates/*.tmpl"))
	content := newContent("login", "user", "POST", "/ultra/backend/api/v1/user/login", "用户", "用户登录")
	handlerGen(content)
	serviceGen(content)
	msgGen(content)
}

func msgGen(c *Content) {
	msgFile := filepath.Join(c.MsgPath, c.DModelName+"_msg.go")
	_, err := os.Stat(msgFile)
	if errors.Is(err, os.ErrNotExist) {
		// 不存在
		msgFileCreate(msgFile, c)
		return
	}
	msgFileAppend(msgFile, c)
}

func msgFileCreate(msgFile string, c *Content) {
	file, err := os.Create(msgFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = temps.ExecuteTemplate(file, "msg_file.tmpl", c)
	if err != nil {
		panic(err)
	}
}

func msgFileAppend(msgFile string, content *Content) {
	file, err := os.OpenFile(msgFile, os.O_APPEND, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = temps.ExecuteTemplate(file, "msg_func.tmpl", content)
	if err != nil {
		panic(err)
	}
}

func serviceGen(c *Content) {
	serviceFile := filepath.Join(c.ServicesPath, c.DModelName+"_service.go")
	_, err := os.Stat(serviceFile)
	if errors.Is(err, os.ErrNotExist) {
		// 不存在
		serviceFileCreate(serviceFile, c)
		return
	}
	serviceFileAppend(serviceFile, c)
}

func serviceFileCreate(serviceFile string, c *Content) {
	file, err := os.Create(serviceFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = temps.ExecuteTemplate(file, "service_file.tmpl", c)
	if err != nil {
		panic(err)
	}
}

func serviceFileAppend(serviceFile string, content *Content) {
	file, err := os.OpenFile(serviceFile, os.O_APPEND, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = temps.ExecuteTemplate(file, "service_func.tmpl", content)
	if err != nil {
		panic(err)
	}
}

// 代码生成
func handlerGen(c *Content) {
	handlerFile := filepath.Join(c.HandlersPath, c.DModelName+"_handler.go")
	_, err := os.Stat(handlerFile)
	if errors.Is(err, os.ErrNotExist) {
		// 不存在
		handlerFileCreate(handlerFile, c)
		return
	}
	handlerFileAppend(handlerFile, c)
}

// handler文件生成
func handlerFileCreate(handlerFile string, content *Content) {
	file, err := os.Create(handlerFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = temps.ExecuteTemplate(file, "handler_file.tmpl", content)
	if err != nil {
		panic(err)
	}
}

func handlerFileAppend(handlerFile string, content *Content) {
	file, err := os.OpenFile(handlerFile, os.O_APPEND, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = temps.ExecuteTemplate(file, "handler_func.tmpl", content)
	if err != nil {
		panic(err)
	}
}
