package config

import (
	"embed"
	"net/http"
	"text/template"

	"github.com/hyetpang/go-code-gen/pkg/common"
	"github.com/hyetpang/go-code-gen/pkg/conf"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	MethodName     string `validate:"required"` // 方法名
	ModelName      string `validate:"required"` // 模型名字
	FilePrefix     string `validate:"required"`
	ProjectRootDir string
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

	DModelName     string // 首字母小写的模型名字
	ReqParamType   string
	RspParamType   string
	IsAddIp        bool
	IsAddUserId    bool
	IsAddCompanyId bool
	DocSummary     string
	Temps          *template.Template
}

//go:embed templates
var templateFiles embed.FS

func New(options ...Option) *Config {
	c := new(Config)
	c.ReqParamType = RspParamTypeBody
	c.RspParamType = RspParamTypeObject
	for _, o := range options {
		o(c)
	}
	v := validator.New()
	err := v.Struct(c)
	if err != nil {
		panic(err)
	}
	if common.UpperFirst(c.DocUrlMethod) == http.MethodGet {
		c.ReqParamType = RspParamTypeQuery
	}
	if c.RspParamType == RspParamTypeArray {
		c.RspType = "[]"
	}
	c.Temps = template.Must(template.ParseFS(templateFiles, "templates/*.tmpl"))
	return c
}

// 开始生成配置
func NewFromMethods(methods []*conf.Method) []*Config {
	configs := make([]*Config, len(methods))
	for methodIndex, method := range methods {
		options := make([]Option, 0, 10)
		if *method.AddIpToReqParam {
			options = append(options, WithAddIpToReqParam())
		}
		options = append(options,
			WithModeName(method.ModelName),            // 模型
			WithMethodName(method.MethodName),         // 要生成的方法
			WithDocDesc(method.DocDesc),               // 文档描述
			WithDocUrl(method.DocUrl),                 // url
			WithDocUrlMethod(method.DocUrlMethod),     // 请求method
			WithDocTag(method.DocTag),                 // 文档分类tags
			WithProjectRoot(method.ProjectRootDir),    // 仓库中的logic目录，
			WithRepoName(method.RepoName),             // 包含logic目录的仓库目录
			WithDependencyName(method.DependencyName), // 依赖库)
		)
		configs[methodIndex] = New(options...)
	}
	return configs
}
