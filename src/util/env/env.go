package env

import (
	"os"
)

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

func Get() string {
	env := os.Getenv("LCENV")
	if env == "" {
		env = EnvDevelopment
	}

	if env != EnvDevelopment && env != EnvProduction && env != EnvStaging {
		env = EnvDevelopment
	}
	return env
}

func IsProduction() bool {
	if Get() == EnvDevelopment {
		return true
	}
	return false
}

func IsDevelopent() bool {
	if Get() == EnvDevelopment {
		return true
	}
	return false
}
func IsSaging() bool {
	if Get() == EnvStaging {
		return true
	}
	return false
}
