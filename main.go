package main

import (
	"fmt"

	"github.com/SigNoz/signoz-collector-config-parser/service"
	"github.com/SigNoz/signoz-otel-collector/components"
	yamlParser "github.com/knadh/koanf/parsers/yaml"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"gopkg.in/yaml.v2"
)

func IsValid(content []byte) (*service.Config, error) {
	var rawConf map[string]interface{}
	if err := yaml.Unmarshal(content, &rawConf); err != nil {
		return nil, err
	}

	conf := confmap.NewFromStringMap(rawConf)
	factories, err := components.Components()
	if err != nil {
		return nil, err
	}

	cfg, err := service.Unmarshal(conf, factories)
	if err != nil {
		return nil, err
	}

	return &service.Config{
		Receivers:  cfg.Receivers.GetReceivers(),
		Processors: cfg.Processors.GetProcessors(),
		Exporters:  cfg.Exporters.GetExporters(),
		Extensions: cfg.Extensions.GetExtensions(),
		Service:    cfg.Service,
	}, nil
}

func main() {
	content := []byte(`
receivers:
  otlp:
    protocols:
      grpc:
      http:
processors:
  batch:
exporters:
  clickhousetraces:
extensions:
  health_check:
service:
  extensions: [health_check]
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [clickhousetraces]
`)
	cfgMain, err := IsValid(content)
	if err != nil {
		panic(err)
	}
	for k, v := range cfgMain.Receivers {
		switch k.Type() {
		case "otlp":
			v := v.(*otlpreceiver.Config)
			v.HTTP = &confighttp.HTTPServerSettings{}
		default:
		}

	}

	coreCfg := CoreComponents()

	getByType := func(cfg map[string]component.Config, t string) component.Config {
		for k, v := range cfg {
			if k == t {
				return v
			}
		}
		return nil
	}

	otlp := getByType(coreCfg.Receivers, "otlp").(*otlpreceiver.Config)
	otlp.HTTP = &confighttp.HTTPServerSettings{
		Endpoint: ":1234",
	}
	err = otlp.Validate()
	yamlBytes, err := yaml.Marshal(otlp)

	c, err := yamlParser.Parser().Unmarshal(content)
	conf := confmap.NewFromStringMap(c)

	c2, err := yamlParser.Parser().Unmarshal(yamlBytes)
	conf2 := confmap.NewFromStringMap(c2)

	conf2.Merge(conf)

	factories, err := components.Components()
	cfg, err := service.Unmarshal(conf2, factories)
	newCfg := &service.Config{
		Receivers:  cfg.Receivers.GetReceivers(),
		Processors: cfg.Processors.GetProcessors(),
		Exporters:  cfg.Exporters.GetExporters(),
		Extensions: cfg.Extensions.GetExtensions(),
		Service:    cfg.Service,
	}
	configBytes, err := yaml.Marshal(newCfg)
	println(string(configBytes))

	// Or
	fmt.Println("koanf")
	b, err := yamlParser.Parser().Marshal(conf2.ToStringMap())
	fmt.Println(string(b))
}
