package main

import (
	"log"
	"reflect"

	"go-code-gen/common"
	"go-code-gen/conf"
	"go-code-gen/config"
	"go-code-gen/strategy"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {
	configData := parseConfig()
	configs := config.NewFromMethods(configData.Methods)
	strategy.Runs(configs)
}

func parseConfig() *conf.Config {
	configData := new(conf.Config)
	viper.SetConfigFile("./conf.toml")
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
	var False bool
	if configData.Global.AddCompanyIdToReqParam == nil {
		configData.Global.AddCompanyIdToReqParam = &False
	}
	if configData.Global.AddIpToReqParam == nil {
		configData.Global.AddIpToReqParam = &False
	}
	if configData.Global.AddUserIdToReqParam == nil {
		configData.Global.AddUserIdToReqParam = &False
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
