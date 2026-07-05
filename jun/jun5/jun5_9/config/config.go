package config

import (
	"fmt"
	"sync"
)

type Config struct{}

type ConfigLoader struct {
	once sync.Once
	cfg  *Config
}

func (l *ConfigLoader) Get() *Config {
	l.once.Do(func() {
		fmt.Println("Инициализация Config")
		l.cfg = &Config{}
	})

	return l.cfg
}
