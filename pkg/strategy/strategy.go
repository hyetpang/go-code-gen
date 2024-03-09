package strategy

import "github.com/hyetpang/go-code-gen/pkg/config"

type Strategy interface {
	Gen(*config.Config)
}
