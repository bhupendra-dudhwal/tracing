# Distributed Tracing with OpenTelemetry and Jaeger in Go

This repo shows how to trace requests across three microservices using OpenTelemetry and visualize them with Jaeger.

## 🧱 Architecture
- Service A → Service B → Service C
- Each service propagates context using OpenTelemetry
- Traces are exported to Jaeger for visualization

## 🚀 Getting Started
### 1. Run everything
```bash
docker-compose up --build
```

### 2. Trigger a trace
```bash
curl http://localhost:8001
```

### 3. Visualize in Jaeger
Visit: [http://localhost:16686](http://localhost:16686)

## 📦 Tech Stack
- Go 1.20+
- OpenTelemetry SDK
- Jaeger exporter
- Docker Compose

## 🔮 Future Improvements
- Add metrics
- Include structured logging
- Add OpenTelemetry Collector for scalable pipelines