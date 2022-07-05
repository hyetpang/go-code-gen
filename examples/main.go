/*
 * @Date: 2022-05-16 13:58:11
 * @LastEditTime: 2022-05-16 18:57:55
 * @FilePath: /go-code-gen/examples/main.go
 * @Author: guangming.zhang hyetpang@yeah.net
 * @LastEditors: guangming.zhang hyetpang@yeah.net
 * @Description:
 *
 * Copyright (c) 2022 by hero, All Rights Reserved.
 */
package main

import (
	"go-code-gen/config"
	"go-code-gen/strategy"
)

func main() {
	strategy.Run(config.New(
		config.WithModeName("GameAgentChannel"),                          // 模型名字
		config.WithMethodName("Subordinate"),                             // 要生成的方法
		config.WithDocDesc("获取我的下级用户id"),                                 // 文档描述
		config.WithDocUrl("/ultra//api/services/game/agent/subordinate"), // url
		config.WithDocUrlMethod("GET"),                                   // 请求method
		config.WithDocTag("经销商"),                                         // 文档分类tags
		config.WithLogicPath("/projects/ultrasdk/ultrasdk.hub.go/logic"), // 仓库中的logic目录，
		config.WithDependencyName("github.com/HyetPang/go-frame"),        // 依赖库
		config.WithRepoName("ultrasdk.hub.go"),                           // 包含logic目录的仓库目录
	))
}
