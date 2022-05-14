package config

import (
	"net/http"
	"text/template"

	"go-code-gen/common"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	MethodName     string `validate:"required"` // 方法名
	ModelName      string `validate:"required"` // 模型名字
	FilePrefix     string `validate:"required"`
	LogicPath      string
	ServicesPath   string `validate:"required"` // 生成的services路径
	HandlersPath   string `validate:"required"` // 生成的handlers路径
	MsgPath        string `validate:"required"` // 生成的msg路径
	DependencyName string `validate:"required"`
	RepoName       string `validate:"required"`
	DocUrlMethod   string `validate:"required"` //
	DocUrl         string `validate:"required"` //
	DocTag         string `validate:"required"`
	DocDesc        string `validate:"required"`
	RspType        string

	DModelName   string // 首字母小写的模型名字
	ReqParamType string
	RspParamType string
	DocSummary   string
	Temps        *template.Template
}

func New(options ...Option) *Config {
	c := new(Config)
	c.ReqParamType = "body"
	c.RspParamType = "object"
	for _, o := range options {
		o(c)
	}
	v := validator.New()
	err := v.Struct(c)
	if err != nil {
		panic(err)
	}
	if common.UpperFirst(c.DocUrlMethod) == http.MethodGet {
		c.ReqParamType = "query"
	}
	if c.RspParamType == "array" {
		c.RspType = "[]"
	}
	c.Temps = template.Must(template.ParseGlob("./templates/*.tmpl"))
	return c
}
