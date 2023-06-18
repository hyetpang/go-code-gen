package strategy

import "go-code-gen/pkg/config"

type Strategy interface {
	Gen(*config.Config)
}
