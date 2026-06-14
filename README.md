# ServStore

A cloud-native, distributed-ready, S3-compatible object storage engine. ServStore serves as an open-source alternative to MinIO, designed for strong consistency, high scalability, and high performance.

Currently, this repository contains the **MVP release**, featuring a robust single-node storage engine with native S3 compatibility, Object Versioning, AWS Signature V4 verification, S3 Multipart Upload support, a built-in Glassmorphic Admin Console, and lightweight OpenTelemetry tracing.

---

## Key Features
* **S3-Compatible REST API**: Native support for creating/deleting buckets, uploading/retrieving objects, HEAD requests, delete markers, and listing bucket contents.
* **S3 Multipart Uploads**: Supports standard S3 multipart operations (`InitiateMultipartUpload`, `UploadPart`, `CompleteMultipartUpload`, and `AbortMultipartUpload`) for uploading large files.
* **Object Versioning**: Supports versioning states (Enabled, Suspended, Disabled) matching AWS S3 versioning specs.
* **Authentication**: Decodes and verifies AWS Signature V4 (header-based and query-based signature verification).
* **OpenTelemetry Tracing**: A custom, lightweight, zero-dependency tracing client (inspired by the `Serv-lang` project) to export trace spans of HTTP routes and storage I/O operations to any OTel collector.
* **Console Dashboard**: A premium, responsive Web UI with dark mode, drag-and-drop uploads, bucket management, and version history viewer.
* **Single-Binary Deployment**: Frontend assets are embedded directly into the Go compiled binary for simple, zero-dependency distribution.

---

## Directory Structure
```text
ServStore/
├── cmd/
│   └── servstore/
│       └── main.go             # Application entry point & CLI flag configuration
├── pkg/
│   ├── auth/
│   │   └── auth.go             # AWS Signature V4 authentication handler
│   ├── otel/
│   │   ├── otel.go             # Lightweight OpenTelemetry tracing client
│   │   └── otel_test.go        # Unit tests for OTel tracing
│   ├── s3/
│   │   ├── api.go              # S3 API Router, Gateway handlers & HTTP tracing
│   │   └── xml.go              # S3-compliant XML request/response models
│   ├── storage/
│   │   ├── store.go            # Storage engine interface definition
│   │   ├── local_store.go      # Local versioned storage & multipart staging
│   │   └── local_store_test.go # Storage engine test suite (including multipart tests)
│   └── web/
│       ├── server.go           # Web Console static asset and API router wrapper
│       └── assets/             # Web Console files (index.html, style.css, app.js)
├── roadmap.md                  # Phase-wise roadmap requirements
└── README.md                   # Product documentation
```

---

## Getting Started

### Prerequisites
* Go 1.22 or higher

### 1. Run Tests
Validate the versioned local storage engine and tracing modules by running the test suite:
```bash
go test -v ./...
```

### 2. Build the Server
Compile the single-binary executable:
```bash
go build -o servstore ./cmd/servstore
```

### 3. Run the Server
Launch the storage engine (by default it listens on port `8080` and stores data inside `./data` with authentication disabled for local console convenience):
```bash
./servstore
```

To run with AWS Signature V4 verification enabled:
```bash
./servstore -auth -access-key "yourAccessKey" -secret-key "yourSecretKey"
```

To enable OpenTelemetry tracing (e.g. exporting to a local OTel collector or Jaeger):
```bash
# Set OTel endpoint environment variables before running the binary
$env:OTEL_ENDPOINT="http://localhost:4318"
$env:OTEL_SERVICE_NAME="servstore"
./servstore
```

### 4. Open the Web Console
Navigate to [http://localhost:8080](http://localhost:8080) in your web browser. From here you can:
* Create and delete buckets.
* Toggle versioning (Enabled / Suspended).
* Drag and drop files to upload them.
* Inspect object metadata, download past versions, or permanently delete them.

---

## Roadmap

To see the development roadmap and requirements for building ServStore into a multi-node, Raft-replicated distributed system, see [roadmap.md](roadmap.md).
