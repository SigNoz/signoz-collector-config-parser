package main

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/fluentbitextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarder"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/extension/ballastextension"
	"go.opentelemetry.io/collector/extension/zpagesextension"
	"go.opentelemetry.io/collector/processor/batchprocessor"
	"go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
)

// func Components() (component.Factories, error) {
// 	var errs []error
// 	factories, err := CoreComponents()
// 	if err != nil {
// 		return component.Factories{}, err
// 	}

// 	var extensions []component.ExtensionFactory
// 	for _, ext := range factories.Extensions {
// 		extensions = append(extensions, ext)
// 	}
// 	factories.Extensions, err = component.MakeExtensionFactoryMap(extensions...)
// 	if err != nil {
// 		errs = append(errs, err)
// 	}

// 	receivers := []component.ReceiverFactory{
// 		apachereceiver.NewFactory(),
// 		carbonreceiver.NewFactory(),
// 		collectdreceiver.NewFactory(),
// 		couchdbreceiver.NewFactory(),
// 		dockerstatsreceiver.NewFactory(),
// 		dotnetdiagnosticsreceiver.NewFactory(),
// 		elasticsearchreceiver.NewFactory(),
// 		filelogreceiver.NewFactory(),
// 		fluentforwardreceiver.NewFactory(),
// 		hostmetricsreceiver.NewFactory(),
// 		httpcheckreceiver.NewFactory(),
// 		jaegerreceiver.NewFactory(),
// 		jmxreceiver.NewFactory(),
// 		journaldreceiver.NewFactory(),
// 		k8sclusterreceiver.NewFactory(),
// 		k8seventsreceiver.NewFactory(),
// 		k8sobjectsreceiver.NewFactory(),
// 		kafkametricsreceiver.NewFactory(),
// 		kafkareceiver.NewFactory(),
// 		kubeletstatsreceiver.NewFactory(),
// 		memcachedreceiver.NewFactory(),
// 		mongodbatlasreceiver.NewFactory(),
// 		mongodbreceiver.NewFactory(),
// 		mysqlreceiver.NewFactory(),
// 		nginxreceiver.NewFactory(),
// 		opencensusreceiver.NewFactory(),
// 		oracledbreceiver.NewFactory(),
// 		otlpjsonfilereceiver.NewFactory(),
// 		podmanreceiver.NewFactory(),
// 		postgresqlreceiver.NewFactory(),
// 		prometheusexecreceiver.NewFactory(),
// 		prometheusreceiver.NewFactory(),
// 		pulsarreceiver.NewFactory(),
// 		rabbitmqreceiver.NewFactory(),
// 		receivercreator.NewFactory(),
// 		redisreceiver.NewFactory(),
// 		sqlqueryreceiver.NewFactory(),
// 		sqlserverreceiver.NewFactory(),
// 		simpleprometheusreceiver.NewFactory(),
// 		statsdreceiver.NewFactory(),
// 		syslogreceiver.NewFactory(),
// 		tcplogreceiver.NewFactory(),
// 		udplogreceiver.NewFactory(),
// 		zipkinreceiver.NewFactory(),
// 		zookeeperreceiver.NewFactory(),
// 	}
// 	for _, rcv := range factories.Receivers {
// 		receivers = append(receivers, rcv)
// 	}
// 	factories.Receivers, err = component.MakeReceiverFactoryMap(receivers...)
// 	if err != nil {
// 		errs = append(errs, err)
// 	}

// 	exporters := []component.ExporterFactory{
// 		carbonexporter.NewFactory(),
// 		clickhousemetricsexporter.NewFactory(),
// 		clickhousetracesexporter.NewFactory(),
// 		clickhouselogsexporter.NewFactory(),
// 		fileexporter.NewFactory(),
// 		jaegerexporter.NewFactory(),
// 		jaegerthrifthttpexporter.NewFactory(),
// 		kafkaexporter.NewFactory(),
// 		loadbalancingexporter.NewFactory(),
// 		opencensusexporter.NewFactory(),
// 		parquetexporter.NewFactory(),
// 		prometheusexporter.NewFactory(),
// 		prometheusremotewriteexporter.NewFactory(),
// 		pulsarexporter.NewFactory(),
// 		zipkinexporter.NewFactory(),
// 	}
// 	for _, exp := range factories.Exporters {
// 		exporters = append(exporters, exp)
// 	}
// 	factories.Exporters, err = component.MakeExporterFactoryMap(exporters...)
// 	if err != nil {
// 		errs = append(errs, err)
// 	}

// 	processors := []component.ProcessorFactory{
// 		attributesprocessor.NewFactory(),
// 		cumulativetodeltaprocessor.NewFactory(),
// 		deltatorateprocessor.NewFactory(),
// 		filterprocessor.NewFactory(),
// 		groupbyattrsprocessor.NewFactory(),
// 		groupbytraceprocessor.NewFactory(),
// 		k8sattributesprocessor.NewFactory(),
// 		metricsgenerationprocessor.NewFactory(),
// 		metricstransformprocessor.NewFactory(),
// 		probabilisticsamplerprocessor.NewFactory(),
// 		redactionprocessor.NewFactory(),
// 		resourcedetectionprocessor.NewFactory(),
// 		resourceprocessor.NewFactory(),
// 		routingprocessor.NewFactory(),
// 		schemaprocessor.NewFactory(),
// 		servicegraphprocessor.NewFactory(),
// 		signozspanmetricsprocessor.NewFactory(),
// 		spanmetricsprocessor.NewFactory(),
// 		spanprocessor.NewFactory(),
// 		tailsamplingprocessor.NewFactory(),
// 		transformprocessor.NewFactory(),
// 		logstransformprocessor.NewFactory(),
// 	}
// 	for _, pr := range factories.Processors {
// 		processors = append(processors, pr)
// 	}
// 	factories.Processors, err = component.MakeProcessorFactoryMap(processors...)
// 	if err != nil {
// 		errs = append(errs, err)
// 	}

// 	return factories, multierr.Combine(errs...)
// }

type CoreComponentsConfig struct {
	Extensions map[string]component.Config
	Receivers  map[string]component.Config
	Processors map[string]component.Config
	Exporters  map[string]component.Config
}

func CoreComponents() CoreComponentsConfig {

	ext := make(map[string]component.Config)
	ext[string(ballastextension.NewFactory().Type())] = &ballastextension.Config{}
	ext[string(basicauthextension.NewFactory().Type())] = &basicauthextension.Config{}
	ext[string(bearertokenauthextension.NewFactory().Type())] = &bearertokenauthextension.Config{}
	ext[string(dockerobserver.NewFactory().Type())] = &dockerobserver.Config{}
	ext[string(ecsobserver.NewFactory().Type())] = &ecsobserver.Config{}
	ext[string(ecstaskobserver.NewFactory().Type())] = &ecstaskobserver.Config{}
	ext[string(filestorage.NewFactory().Type())] = &filestorage.Config{}
	ext[string(fluentbitextension.NewFactory().Type())] = &fluentbitextension.Config{}
	ext[string(hostobserver.NewFactory().Type())] = &hostobserver.Config{}
	ext[string(httpforwarder.NewFactory().Type())] = &httpforwarder.Config{}
	ext[string(jaegerremotesampling.NewFactory().Type())] = &jaegerremotesampling.Config{}
	ext[string(k8sobserver.NewFactory().Type())] = &k8sobserver.Config{}
	ext[string(oauth2clientauthextension.NewFactory().Type())] = &oauth2clientauthextension.Config{}
	ext[string(oidcauthextension.NewFactory().Type())] = &oidcauthextension.Config{}
	ext[string(healthcheckextension.NewFactory().Type())] = &healthcheckextension.Config{}
	ext[string(pprofextension.NewFactory().Type())] = &pprofextension.Config{}
	ext[string(zpagesextension.NewFactory().Type())] = &zpagesextension.Config{}

	rec := make(map[string]component.Config)
	rec[string(otlpreceiver.NewFactory().Type())] = &otlpreceiver.Config{}

	exp := make(map[string]component.Config)
	exp[string(loggingexporter.NewFactory().Type())] = &loggingexporter.Config{}
	exp[string(otlpexporter.NewFactory().Type())] = &otlpexporter.Config{}
	exp[string(otlphttpexporter.NewFactory().Type())] = &otlphttpexporter.Config{}

	prc := make(map[string]component.Config)
	prc[string(batchprocessor.NewFactory().Type())] = &batchprocessor.Config{}
	prc[string(memorylimiterprocessor.NewFactory().Type())] = &memorylimiterprocessor.Config{}

	return CoreComponentsConfig{
		Extensions: ext,
		Receivers:  rec,
		Processors: prc,
		Exporters:  exp,
	}
}
