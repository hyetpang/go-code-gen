package strategy

import "go-code-gen/config"

type Strategy interface {
	Gen(*config.Config)
}
