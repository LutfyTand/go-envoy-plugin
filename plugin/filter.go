package plugin

import (
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

const Name = "http_plugin"

type Filter struct {
	callback api.FilterCallbackHandler
	config   *Config
}

func (f *Filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) (status api.StatusType) {
	// header.Set("x-city", f.config.AddRequestHeaders["x-city"])
	// header.Set("x-country-code", f.config.AddRequestHeaders["x-country-code"])
	for k, v := range f.config.AddRequestHeaders {
		header.Set(k, v)
	}
	return api.Continue
}

func (f *Filter) EncodeHeaders(header api.ResponseHeaderMap, endSteam bool) (status api.StatusType) {
	for k, v := range f.config.AddResponseHeaders {
		header.Set(k, v)
	}
	// header.Set("via", f.config.AddResponseHeaders["via"])
	return api.Continue
}

// DecodeData is not implemented
func (f *Filter) DecodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	// NOT IMPLEMENTED, quick no-op
	return api.Continue
}

// EncodeData is not implemented
func (f *Filter) EncodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	// NOT IMPLEMENTED, quick no-op
	return api.Continue
}

// DecodeTrailers is not implemented
func (f *Filter) DecodeTrailers(trailers api.RequestTrailerMap) api.StatusType {
	// NOT IMPLEMENTED, quick no-op
	return api.Continue
}

// EncodeTrailers is not implemented
func (f *Filter) EncodeTrailers(trailers api.ResponseTrailerMap) api.StatusType {
	// NOT IMPLEMENTED, quick no-op
	return api.Continue
}

// OnDestroy is not implemented
func (f *Filter) OnDestroy(reason api.DestroyReason) {
}

func (f *Filter) OnLog() {
	api.LogError("call log in OnLog")
}

func (f *Filter) OnLogDownstreamPeriodic() {

}

func (f *Filter) OnLogDownstreamStart() {

}
