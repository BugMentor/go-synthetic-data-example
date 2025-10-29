# Go Synthetic Data Tool for Cloud SDET

This repository demonstrates a modular, scalable Go application for generating high-volume, quality-assured synthetic data, suitable for testing enterprise-scale, multi-tenant cloud applications (like OpenGov).

The architecture emphasizes:

1. **Modularity:** Logic is separated into `generator`, `models`, and `validator` packages.

2. **Concurrency:** Uses Go's goroutines to achieve high generation throughput.

3. **Validation:** Includes a dedicated `validator` package to enforce business and structural rules on the generated data.

## 📝 Table of Contents

- [Go Synthetic Data Tool for Cloud SDET](#go-synthetic-data-tool-for-cloud-sdet)
  - [📝 Table of Contents](#-table-of-contents)
  - [🚀 Structure](#-structure)
  - [✨ Features](#-features)
  - [🛠️ How to Run](#️-how-to-run)

## 🚀 Structure

```
/go-synthetic-data-tool
├── cmd
│   └── main.go       <-- Executes the generation and validation pipeline
├── internal
│   ├── generator     <-- Core logic for creating data and managing concurrency
│   ├── models        <-- Go Struct definitions (Schemas)
│   └── validator     <-- Logic to enforce business rules and data constraints
└── README.md
```

## ✨ Features

* **High Performance:** Concurrent data generation using a worker pool.

* **Realistic Data:** Simulates statistical distributions (e.g., Normal Distribution) and enforces categorical constraints (`Department`, `TenantID`).

* **Quality Gate:** The built-in `validator` ensures generated data meets all compliance and structural rules (e.g., multitenancy checks, minimum amount constraints) before being used in automated tests.

## 🛠️ How to Run

1. Make sure you have Go installed.

2. Initialize the Go module (if running locally):

   ```
   go mod init go-synthetic-data-tool 
   
   ```

3. Run the main application from the root directory:

   ```
   go run cmd/main.go
   ```