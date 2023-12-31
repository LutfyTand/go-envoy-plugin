package plugin

import (
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

func ConfigFactory(c interface{}) api.StreamFilterFactory {
	conf, ok := c.(*Config)
	if !ok {
		panic("unexpected config type")
	}
	return func(callbacks api.FilterCallbackHandler) api.StreamFilter {
		return &Filter{
			callback: callbacks,
			config:   conf,
		}
	}
}
