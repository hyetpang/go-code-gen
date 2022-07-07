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
		config.WithModeName("DataSync"),                                  // 模型名字
		config.WithMethodName("Sync"),                                    // 要生成的方法
		config.WithDocDesc("数据同步,服务器接口，不对前端开放"),                          // 文档描述
		config.WithDocUrl("/ultra/api/data/sync"),                        // url
		config.WithDocUrlMethod("POST"),                                  // 请求method
		config.WithDocTag("数据同步"),                                        // 文档分类tags
		config.WithLogicPath("/projects/ultrasdk/ultrasdk.hub.go/logic"), // 仓库中的logic目录，
		config.WithRepoName("ultrasdk.hub.go"),                           // 包含logic目录的仓库目录
		config.WithDependencyName("github.com/HyetPang/go-frame"),        // 依赖库
		config.WithAddIpToReqParam(),                                     // 给请求参数增加ip字段
		config.WithAddUserIdToReqParam(),                                 // 给请求参数增加userId字段
		config.WithAddCompanyIdToReqParam(),                              // 给请求参数增加companyId字段
	))
}
