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
		config.WithModeName("UltraboxNotice"),                                                // 模型名字
		config.WithMethodName("Validate"),                                                    // 要生成的方法
		config.WithDocDesc("活动详情"),                                                           // 文档描述
		config.WithDocUrl("/ultrabox/api/v1/activity"),                                       // url
		config.WithDocUrlMethod("GET"),                                                       // 请求method
		config.WithDocTag("游戏盒子客户端-活动"),                                                      // 文档分类tags
		config.WithLogicPath("/projects/ultrasdk/ultrasdk.hub.go/cmd/ultrabox_client/logic"), // 仓库中的logic目录，
		// config.WithLogicPath("/projects/ultrasdk/ultrasdk.hub.go/logic"), // 仓库中的logic目录，
		// config.WithRspParamType(),                                                            // services层逻辑处理返回的响应，如果不指定返回对象，指定返回数组
		config.WithDependencyName("github.com/HyetPang/go-frame"),  // 依赖库
		config.WithRepoName("ultrasdk.hub.go/cmd/ultrabox_client"), // 包含logic目录的仓库目录
	))
}
