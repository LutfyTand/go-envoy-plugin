package main

import (
	"github.com/LutfyTand/go-envoy-plugin/plugin"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
)

func init() {
	http.RegisterHttpFilterConfigFactoryAndParser(plugin.Name, plugin.ConfigFactory, &plugin.ConfigParser{})
}

func main() {}
