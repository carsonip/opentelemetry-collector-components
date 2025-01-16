module github.com/elastic/opentelemetry-collector-components/loadgen

go 1.22.7

require (
	github.com/elastic/opentelemetry-collector-components/processor/ratelimitprocessor v0.0.0-00010101000000-000000000000
	github.com/elastic/opentelemetry-collector-components/receiver/loadgenreceiver v0.0.0-00010101000000-000000000000
	github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor v0.117.0
	go.opentelemetry.io/collector/component v0.117.0
	go.opentelemetry.io/collector/confmap v1.23.0
	go.opentelemetry.io/collector/confmap/provider/envprovider v1.23.0
	go.opentelemetry.io/collector/confmap/provider/fileprovider v1.23.0
	go.opentelemetry.io/collector/confmap/provider/httpprovider v1.23.0
	go.opentelemetry.io/collector/confmap/provider/httpsprovider v1.23.0
	go.opentelemetry.io/collector/confmap/provider/yamlprovider v1.23.0
	go.opentelemetry.io/collector/connector v0.117.0
	go.opentelemetry.io/collector/exporter v0.117.0
	go.opentelemetry.io/collector/exporter/debugexporter v0.117.0
	go.opentelemetry.io/collector/exporter/nopexporter v0.117.0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.117.0
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.117.0
	go.opentelemetry.io/collector/extension v0.117.0
	go.opentelemetry.io/collector/otelcol v0.117.0
	go.opentelemetry.io/collector/processor v0.117.0
	go.opentelemetry.io/collector/receiver v0.117.0
	go.opentelemetry.io/collector/receiver/nopreceiver v0.117.0
)

require (
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/antchfx/xmlquery v1.4.3 // indirect
	github.com/antchfx/xpath v1.3.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/elastic/go-grok v0.3.1 // indirect
	github.com/elastic/lunes v0.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.23.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v0.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.117.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter v0.117.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil v0.117.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl v0.117.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.117.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.61.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/shirou/gopsutil/v4 v4.24.12 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/ua-parser/uap-go v0.0.0-20240611065828-3a4781585db6 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/collector v0.117.0 // indirect
	go.opentelemetry.io/collector/client v1.23.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.117.0 // indirect
	go.opentelemetry.io/collector/component/componenttest v0.117.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.117.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.23.0 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.117.0 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.117.0 // indirect
	go.opentelemetry.io/collector/config/confignet v1.23.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.23.0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.23.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.117.0 // indirect
	go.opentelemetry.io/collector/config/configtls v1.23.0 // indirect
	go.opentelemetry.io/collector/connector/connectortest v0.117.0 // indirect
	go.opentelemetry.io/collector/connector/xconnector v0.117.0 // indirect
	go.opentelemetry.io/collector/consumer v1.23.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.117.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror/xconsumererror v0.117.0 // indirect
	go.opentelemetry.io/collector/consumer/consumertest v0.117.0 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.117.0 // indirect
	go.opentelemetry.io/collector/exporter/exporterhelper/xexporterhelper v0.117.0 // indirect
	go.opentelemetry.io/collector/exporter/exportertest v0.117.0 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.117.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.117.0 // indirect
	go.opentelemetry.io/collector/extension/extensioncapabilities v0.117.0 // indirect
	go.opentelemetry.io/collector/extension/extensiontest v0.117.0 // indirect
	go.opentelemetry.io/collector/extension/xextension v0.117.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.23.0 // indirect
	go.opentelemetry.io/collector/internal/fanoutconsumer v0.117.0 // indirect
	go.opentelemetry.io/collector/pdata v1.23.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.117.0 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.117.0 // indirect
	go.opentelemetry.io/collector/pipeline v0.117.0 // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.117.0 // indirect
	go.opentelemetry.io/collector/processor/processortest v0.117.0 // indirect
	go.opentelemetry.io/collector/processor/xprocessor v0.117.0 // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.117.0 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.117.0 // indirect
	go.opentelemetry.io/collector/semconv v0.117.0 // indirect
	go.opentelemetry.io/collector/service v0.117.0 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.6.0 // indirect
	go.opentelemetry.io/contrib/config v0.10.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.57.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.31.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.54.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.7.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.31.0 // indirect
	go.opentelemetry.io/otel/log v0.8.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.7.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.9.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241206012308-a4fef0638583 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241206012308-a4fef0638583 // indirect
	google.golang.org/grpc v1.69.2 // indirect
	google.golang.org/protobuf v1.36.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/elastic/opentelemetry-collector-components/processor/ratelimitprocessor => ./../../../processor/ratelimitprocessor
	github.com/elastic/opentelemetry-collector-components/receiver/loadgenreceiver => ./../../../receiver/loadgenreceiver
)
