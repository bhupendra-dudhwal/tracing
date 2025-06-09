module servicee

go 1.23.9

require github.com/bhupendra-dudhwal/tracing v0.0.0-20250609181431-538d7a15fa66

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0 // indirect
	go.opentelemetry.io/otel/metric v1.22.0 // indirect
	go.opentelemetry.io/otel/sdk v1.22.0 // indirect
	go.opentelemetry.io/otel/trace v1.22.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
)

replace github.com/bhupendra-dudhwal/tracing => ../../tracing
