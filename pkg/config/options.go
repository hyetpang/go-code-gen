package config

import (
	"net/http"
	"path/filepath"
	"strings"

	"go-code-gen/pkg/common"

	"golang.org/x/exp/slices"
)

type Option func(*Config)

func WithMethodName(methodName string) Option {
	return func(c *Config) {
		if !common.MustAlpha(methodName) {
			panic("methodName参数必须是字母!")
		}
		c.MethodName = common.UpperFirst(common.CamelString(methodName))
	}
}

func WithModeName(modelName string) Option {
	return func(c *Config) {
		if !common.MustAlpha(modelName) {
			panic("modelName参数必须是字母!")
		}
		c.FilePrefix = common.SnakeString(common.CamelString(modelName))
		c.ModelName = common.LowerFirst(modelName)
		c.DModelName = modelName
	}
}

func WithServicePath(servicePath string) Option {
	return func(c *Config) {
		c.ServicesPath = servicePath
	}
}

func WithHandlerPath(handlerPath string) Option {
	return func(c *Config) {
		c.HandlersPath = handlerPath
	}
}

func WithMsgPath(msgPath string) Option {
	return func(c *Config) {
		c.MsgPath = msgPath
	}
}

func WithRepoName(repoName string) Option {
	return func(c *Config) {
		c.RepoName = repoName
	}
}

func WithDependencyName(dependencyName string) Option {
	return func(c *Config) {
		c.DependencyName = dependencyName
	}
}

var methods []string = []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace}

func WithDocUrlMethod(docUrlMethod string) Option {
	return func(c *Config) {
		method := common.Upper(docUrlMethod)
		if !slices.Contains(methods, method) {
			panic("无效的http method:" + docUrlMethod + "，http method 只能是[" + strings.Join(methods, ",") + "]中的一个")
		}
		c.DocUrlMethod = method
	}
}

func WithDocUrl(DocUrl string) Option {
	return func(c *Config) {
		c.DocUrl = DocUrl
	}
}

func WithDocTag(docTag string) Option {
	return func(c *Config) {
		c.DocTag = docTag
	}
}

func WithDocDesc(docDesc string) Option {
	return func(c *Config) {
		c.DocDesc = docDesc
		if len(c.DocSummary) < 1 {
			c.DocSummary = docDesc
		}
	}
}

func WithDocSummary(docSummary string) Option {
	return func(c *Config) {
		c.DocSummary = docSummary
	}
}

func WithProjectRoot(projectRootDir string) Option {
	return func(c *Config) {
		c.ProjectRootDir = projectRootDir
		if len(c.HandlersPath) < 1 {
			c.HandlersPath = filepath.Join(c.ProjectRootDir, "logic", "handlers")
		}
		if len(c.ServicesPath) < 1 {
			c.ServicesPath = filepath.Join(c.ProjectRootDir, "logic", "services")
		}
		if len(c.MsgPath) < 1 {
			c.MsgPath = filepath.Join(c.ProjectRootDir, "pkg", "msg")
		}
	}
}

func WithRspParamType() Option {
	return func(c *Config) {
		c.RspParamType = RspParamTypeObject
	}
}

// 是否请求参数需要生成ip字段
func WithAddIpToReqParam() Option {
	return func(c *Config) {
		c.IsAddIp = true
	}
}
