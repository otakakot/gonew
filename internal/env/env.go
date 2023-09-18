package env

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type Env string

const (
	prd   Env = "prd"
	stg   Env = "stg"
	dev   Env = "dev"
	local Env = "local"
	empty Env = ""
)

var (
	env  Env        //nolint:gochecknoglobals
	lock sync.Mutex //nolint:gochecknoglobals
)

func Init() {
	if env != empty {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	e := Env(os.Getenv("ENV"))

	slog.Info(fmt.Sprintf("environment is %s", e))

	env = e
}

func Get() Env {
	return env
}

func (e Env) String() string {
	return string(e)
}

func (e Env) IsProduction() bool {
	return e == prd
}

func (e Env) IsStaging() bool {
	return e == stg
}

func (e Env) IsDevelopment() bool {
	return e == dev
}

func (e Env) IsLocal() bool {
	return e == local
}
