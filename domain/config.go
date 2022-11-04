package domain

import "github.com/recative/recative-backend/definition"

type Config struct {
	CrossMicroServiceConfig definition.CrossServiceConfig `mapstructure:"cross_micro_service"`
}
