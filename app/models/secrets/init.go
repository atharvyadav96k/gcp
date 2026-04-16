package secrets

import (
	"sync"
)

var (
	instance *Env
	once     sync.Once
)

func NewSecrets() *Env {
	once.Do(func() {
		instance = &Env{
			secrets: make(map[string]string),
		}
		instance.LoadSecrets()
	})
	return instance
}
