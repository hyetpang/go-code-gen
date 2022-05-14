package main

import (
	"go-code-gen/config"
	"go-code-gen/strategy"
)

func main() {
	strategy.Run(config.New(
		config.WithModeName("UltraboxGame"),                                                  // 模型名字
		config.WithMethodName("SearchRecommend"),                                             // 要生成的方法
		config.WithLogicPath("/projects/ultrasdk/ultrasdk.hub.go/cmd/ultrabox_client/logic"), // 仓库中的logic目录，
		config.WithDocUrl("/ultrabox/api/v1/game/search/recommend"),                          // url
		config.WithDocUrlMethod("GET"),                                                       // 请求method
		config.WithRspParamType(),                                                            // services层逻辑处理返回的响应，如果不指定返回对象，指定返回数组
		config.WithDocTag("游戏盒子客户端-游戏"),                                                      // 文档分类tags
		config.WithDocDesc("游戏搜索推荐"),                                                         // 文档描述
		config.WithDependencyName("github.com/HyetPang/go-frame"),                            // 依赖库
		config.WithRepoName("ultrasdk.hub.go/cmd/ultrabox_client"),                           // 包含logic目录的仓库目录
	))
}
