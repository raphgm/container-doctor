# Container Doctor 🩺

`container-doctor` is a fast, extensible CLI tool designed to diagnose your system's readiness for running Docker and container workloads. It performs health checks across your system, Docker daemon, and Compose installation, alerting you to critical issues like low memory, insufficient disk space, or missing dependencies.

## Features

- **System Diagnostics**: Checks your operating system, CPU architecture, memory (RAM), and disk space to ensure you meet minimum requirements.
- **Docker Health**: Validates that Docker is installed, the daemon is running, the socket is accessible, and client/server versions are compatible.
- **Docker Compose & Buildx**: Ensures crucial container plugins are installed and ready.
- **Static Inspection**: The `inspect` command statically analyzes your `Dockerfile` and `docker-compose.yml` for common anti-patterns (like using the `:latest` tag in production).
- **Multiple Output Formats**: View reports natively in your terminal, or export them to JSON, Markdown, or HTML.

---

## Installation

### Using Go
If you have Go installed, you can install the CLI directly:
```bash
go install github.com/raphgm/container-doctor@latest
```

### From Releases
Pre-compiled binaries for **macOS**, **Linux**, and **Windows** are automatically generated for every release. 
Head over to the [Releases page](https://github.com/raphgm/container-doctor/releases) to download the binary for your operating system.

---

## Usage

### 1. Run Health Checks
Run the diagnostic engine to check your environment:
```bash
container-doctor check
```

You can export the results to different formats using the `--format` flag:
```bash
# JSON Output
container-doctor check --format json > report.json

# Markdown Output
container-doctor check --format markdown > report.md

# HTML Output (Standalone Webpage)
container-doctor check --format html > report.html
```

### 2. Inspect a Project
Statically analyze the configuration files in your current directory (or any specified path):
```bash
container-doctor inspect .
```

---

## Architecture & Extensibility

`container-doctor` is built with a highly decoupled, plugin-based architecture. 

- **Engine**: The core `internal/engine` runs the execution loop and aggregates results.
- **Providers**: Checks are grouped by Providers (e.g., `system`, `docker`, `compose`).
- **Renderers**: Custom output rendering is handled via the `Renderer` interface in `internal/renderer`.

Adding a new check is as simple as creating a struct that implements the `check.Check` interface and registering it with a Provider!

## License

This project is licensed under the MIT License.
