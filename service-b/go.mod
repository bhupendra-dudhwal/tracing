module github.com/bhupendra-dudhwal/distributed-tracing-sample/service-b

go 1.23.9

// replace github.com/bhupendra-dudhwal/tracing => ../tracing

// github.com/bhupendra-dudhwal/tracing v0.0.0-00010101000000-000000000000
require go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.61.0

require github.com/bhupendra-dudhwal/tracing v0.0.0-20250602211137-d76baa704091

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0 // indirect
	go.opentelemetry.io/otel/metric v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.36.0 // indirect
	go.opentelemetry.io/otel/trace v1.36.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)
