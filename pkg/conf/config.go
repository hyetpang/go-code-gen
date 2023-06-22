package conf

import (
	"go-code-gen/pkg/common"
	"log"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Global  Global    `mapstructure:"global"`
	Methods []*Method `mapstructure:"methods" validate:"min=1"`
}

// 全局配置
type Global struct {
	AddIpToReqParam *bool  `mapstructure:"add_ip_to_req_param"`                  // 	请求参数增加Ip
	ModelName       string `mapstructure:"model_name"`                           // 方法名字
	DocTag          string `mapstructure:"doc_tag"`                              // 文档显示的标签
	ProjectRootDir  string `mapstructure:"project_root_dir" validate:"required"` // logic 路径
	RepoName        string `mapstructure:"repo_name"`                            // 仓库名字
	DependencyName  string `mapstructure:"dependency_name"`                      // 依赖名字
}

// 自定义配置
type Method struct {
	ModelName       string `mapstructure:"model_name" validate:"required"`
	DocTag          string `mapstructure:"doc_tag" validate:"required"`
	ProjectRootDir  string `mapstructure:"project_root_dir" validate:"required"`
	RepoName        string `mapstructure:"repo_name" validate:"required"`
	DependencyName  string `mapstructure:"dependency_name" validate:"required"`
	AddIpToReqParam *bool  `mapstructure:"add_ip_to_req_param" validate:"required"`
	MethodName      string `mapstructure:"method_name" validate:"required"`
	DocDesc         string `mapstructure:"doc_desc" validate:"required"`
	DocUrl          string `mapstructure:"doc_url" validate:"required"`
	DocUrlMethod    string `mapstructure:"doc_url_method" validate:"required"`
}

func ParseConfig(confFile string) *Config {
	configData := new(Config)
	viper.SetConfigFile(confFile)
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件出错:%s", err.Error())
	}
	err = viper.Unmarshal(configData)
	if err != nil {
		log.Fatalf("序列化配置文件出错:%s", err.Error())
	}
	validator := validator.New()
	err = validator.Struct(configData)
	if err != nil {
		log.Fatalf("配置验证出错:%s", err.Error())
	}
	globalValue := reflect.ValueOf(configData.Global)
	globalType := reflect.TypeOf(configData.Global)
	for _, method := range configData.Methods {
		for i := 0; i < globalValue.NumField(); i++ {
			globalFieldValue := globalValue.Field(i)
			globalFieldType := globalType.Field(i)
			methodFieldValue := reflect.ValueOf(method).Elem().FieldByName(globalFieldType.Name)
			switch globalFieldType.Type.Kind() {
			case reflect.Pointer:
				// 布尔
				underlyingKind := globalFieldType.Type.Elem().Kind()
				if underlyingKind == reflect.Bool {
					if methodFieldValue.IsNil() {
						methodFieldValue.Set(globalFieldValue)
					}
				} else {
					log.Fatalf("未知的类型:%s,%s", underlyingKind.String(), globalFieldType.Name)
				}
			case reflect.String:
				// 字符串
				if len(globalFieldValue.String()) < 1 && len(methodFieldValue.String()) < 1 {
					log.Fatalf("配置参数[%s]必须在global或者method段配置", common.SnakeString(globalFieldType.Name))
				} else if len(globalFieldValue.String()) > 0 && len(methodFieldValue.String()) < 1 {
					methodFieldValue.SetString(globalFieldValue.String())
				}
			default:
				log.Fatalf("未知类型:%s", globalFieldValue.Type().Kind().String())
			}
		}
		err = validator.Struct(method)
		if err != nil {
			log.Fatalf("method配置出错:%s", err.Error())
		}
	}
	return configData
}
