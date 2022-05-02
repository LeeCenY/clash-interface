package clash

import (
	"encoding/json"

	"github.com/Dreamacro/clash/tunnel"
)

func HealthCheck() {
	providers := tunnel.Providers()
	for _, provider := range providers {
		provider.HealthCheck()
	}
}

func Providers() []byte {
	providers := tunnel.Providers()
	data, _ := json.Marshal(providers)
	return data
}

func Provider(name string) []byte {
	providers := tunnel.Providers()
	provider, exist := providers[name]
	if !exist {
		return nil
	}
	data, _ := json.Marshal(provider)
	return data
}
