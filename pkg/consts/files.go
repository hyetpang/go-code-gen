package consts

import "go-code-gen/pkg/platform/newline"

const (
	GormFile = `package main

import (
	"os"

	"{{.}}/pkg/orm/models"

	"gorm.io/gen"
)

func main() {
	outPath := "../../pkg/orm/q"
	if len(os.Args) >= 2 {
		outPath = os.Args[1]
	}
	cfg := gen.Config{
		OutPath: outPath,
		Mode:    gen.WithDefaultQuery | gen.WithoutContext,
	}
	g := gen.NewGenerator(cfg)
	g.ApplyBasic(new(models.User))
	g.Execute()
}`

	AppToml = `[app]
run_mode = "dev" # 按具体情况配置: prod,dev,test

[http] 											# http 配置
addr = ":8001"                                  # http服务监听地址
is_doc = true                                   # 是否开启swagger文档
doc_path = "/game-platform/api/v1/swagger/*any" # 文档路由
is_pprof = true                                 # 是否开启pprof
pprof_prefix = ""                               # pprof 路径前缀
is_metrics = true                               # 是否使用指标
metrics_path = "/game-platform/api/v1/metrics"  # 指标导出url
is_prod = false                                 # 是否在线上环境

[mysql]																						   # mysql 配置
connect_string = "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" # mysql连接字符串
max_idle_time = 30                                                                             # 单位：分钟
max_life_time = 60                                                                             # 单位：分钟
max_idle_conns = 10                                                                            # 最大的空闲连接
max_open_conns = 100                                                                           # 最大打开的连接数
table_prefix = ""                                                                            # 表前缀
name = "default"

[redis] 				# redis 配置
addr = "127.0.0.1:6379" # redis连接地址
pwd = ""                # redis密码
db = 3                  # 使用的数据库

[zap_log] # 日志配置
level = -1 # 可能的值:-1=>debug,0=>info,1=>warn,2=>error
`
	RouterFile = `package routers

import (
	"{{.}}/logic/handlers"

	"github.com/gin-gonic/gin"
)

func Inject(g *gin.IRouter, handler handlers.Handlers) {

}`
	ServicesProvider = `package services

import (
	"go.uber.org/fx"
)

func Provider() []fx.Option {
	return []fx.Option{}
}

type Handlers struct {
	fx.In
}`
	HandlersProvider = `package handlers

import (
	"go.uber.org/fx"
)

func Provider() []fx.Option {
	return []fx.Option{}
}

type Handlers struct {
	fx.In
}`
	CodesFile = `package common

import "github.com/hyetpang/go-frame/pkgs/base"

var ErrPlaceHolder = base.NewCodeErr(100, "错误占位符")
`
	ConfigProvider = `package config

import (
	"log"

	"github.com/hyetpang/go-frame/pkgs/common"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func newConfig() *Config {
	conf := new(Config)
	err := viper.UnmarshalKey("app", conf)
	if err != nil {
		log.Fatalf("序列化app配置数据到对象出错:%s", err.Error())
	}
	common.MustValidate(conf)
	return conf
}

func Provider() []fx.Option {
	return []fx.Option{fx.Provide(newConfig)}
}`
	ConfigFile = `package config

type Config struct {
	RunMode     string         ` + "`mapstructure:\"run_mode\" validate:\"oneof= dev test prod\"` // 运行模式，dev,test,prod" + string(newline.NewLine) + "}"
	ConstsFile           = `package consts`
	ModelsRegisterInject = `package model_register

import (
	"{{.}}/pkg/orm/models"
	"{{.}}/pkg/orm/q"

	"github.com/hyetpang/go-frame/pkgs/common"
	"gorm.io/gorm"
)

func Inject(db *gorm.DB) {
	common.Panic(db.AutoMigrate(
		&models.User{},
	))
	q.SetDefault(db)
}`
	GitIgnoreFile = `logs
.vscode`
	ApiRestFile   = `@host=http://app-api.maxity.io
@token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzYWx0IjoiXnptI0k5T05id1dpVEpsNnI4aGVyM0BKemhXaFpzWGxsRE1Dc016JTMxJENQUHhJIiwiYXVkIjoibWF4LWFuZHJvaWQiLCJleHAiOjE2ODYzMDk5ODEsInN1YiI6IjMyMCJ9.8LeTEPbdX9G8mAw64nnM6QRbc5GS0RhABIB5rhrW7k9RnYBA_LVD_nBkNuZP8fF-roNhmbvo2vSv5GKSLZFt7g`
	MainFile = `package main

import (
	"{{.}}/logic/handlers"
	"{{.}}/logic/routers"
	"{{.}}/logic/services"
	"{{.}}/pkg/config"
	"{{.}}/pkg/orm/models_register"

	"github.com/hyetpang/go-frame/pkgs/app"
	"github.com/hyetpang/go-frame/pkgs/options"
)

func main() {
	app.Run(
		options.WithTasks(),
		options.WithHttp(nil),
		options.WithMysql(),
		options.WithRedis(),
		options.WithFxOptions(handlers.Provider(), services.Provider(),config.Provider()),
		options.WithInvokes(model_register.Inject, routers.Inject),
	)
}`
	MakeFile = `gorm:
	go run cmd/gorm/main.go
doc:
	swag fmt --dir=./
	swag init --parseDependency --dir=./,./logic/handlers --quiet=true --ot go
fieldAlign:
	betteralign -apply ./pkg/config ./pkg/orm/models ./pkg/msg # go install github.com/dkorunic/betteralign/cmd/betteralign@latest
build: fieldAlign doc
	env GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags="sonic avx" -a -ldflags '-s -w -extldflags "-static"' -o main main.go
	upx -9 -o main-upx main
	mv main-upx main
scp: build
	mv main maxit
	scp maxit akswork001@akswork.plant.test:/home/jorjan/maxit_server
	rm maxit
deploy: build
	scp main akswork001@akswork.plant.test:/home/docker/maxoauth
	ssh akswork.plant.test "sudo docker restart maxoauth"
	rm main`
	ModelsUserFile = `package models

import "github.com/hyetpang/go-frame/pkgs/base"

type User struct {
	base.Model
}`
)
