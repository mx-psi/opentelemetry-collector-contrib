module github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy

go 1.22.0

require (
	github.com/aws/aws-sdk-go v1.55.5
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.108.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/config/confignet v0.108.2-0.20240829190554-7da6b618a7ee
	go.opentelemetry.io/collector/config/configtls v1.14.2-0.20240904075637-48b11ba1c5f8
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.14.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.opentelemetry.io/collector/featuregate v1.14.2-0.20240904075637-48b11ba1c5f8 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../../internal/common

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
