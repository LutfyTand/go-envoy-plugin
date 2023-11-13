package plugin

import (
	"encoding/json"

	xds "github.com/cncf/xds/go/xds/type/v3"
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"google.golang.org/protobuf/types/known/anypb"
)

type Config struct {
	AddRequestHeaders  map[string]string `json:"add_request_headers"`
	AddResponseHeaders map[string]string `json:"add_response_headers"`
}

type ConfigParser struct{}

func (p *ConfigParser) Parse(any *anypb.Any, callbacks api.ConfigCallbackHandler) (interface{}, error) {
	configStruct := &xds.TypedStruct{}
	if err := any.UnmarshalTo(configStruct); err != nil {
		return nil, err
	}

	v := configStruct.GetValue()
	jsonByte, err := json.Marshal(v.AsMap())
	if err != nil {
		return nil, err
	}

	var c Config
	if err := json.Unmarshal(jsonByte, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (p *ConfigParser) Merge(parent interface{}, child interface{}) interface{} {
	parentConfig := parent.(*Config)
	childConfig := child.(*Config)

	margedConfig := *parentConfig
	margedConfig.AddRequestHeaders = childConfig.AddRequestHeaders
	margedConfig.AddResponseHeaders = childConfig.AddResponseHeaders

	return &margedConfig
}
