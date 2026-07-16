package cache

import "sync/atomic"

type Config struct {
	HTTPAdrr string
	Timeout  int
	Version  string
}

type ConfigStore struct {
	pointer atomic.Pointer[Config]
}

func NewConfigStore(c *Config) *ConfigStore {
	s := &ConfigStore{}
	s.pointer.Store(c)

	return s
}

func (s *ConfigStore) Get() Config {
	c := s.pointer.Load()

	return *c
}

func (s *ConfigStore) Update(c *Config) {
	s.pointer.Store(c)
}
