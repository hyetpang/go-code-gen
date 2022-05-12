package main

import (
	"go-code-gen/config"
	"go-code-gen/strategy"
)

func main() {
	strategy.Run(config.New(
		config.WithModeName("User"),
		config.WithMethodName("Login"),
		config.WithLogicPath("F:/projects/demo/logic"),
		config.WithDocUrl("user/login"),
		config.WithDocUrlMethod("post"),
		config.WithDocTag("用户"),
		config.WithDocDesc("用户登录"),
		config.WithDependencyName("github.com/HyetPang/go-frame"),
		config.WithRepoName("demo"),
	))
}
