# Container Doctor

Container Doctor is a unified diagnostic engine and framework designed to analyze and troubleshoot container environments, orchestrators, and projects with CNCF-quality standards.

## Architecture

At its core, Container Doctor is an extensible framework:
- **Engine**: Core execution logic that loads providers, runs checks, and builds structured reports.
- **Providers**: Plugins that implement specific domain checks (e.g., Docker, Kubernetes, Helm).
- **Checks**: Individual diagnostics yielding a `Result` (Pass, Warn, Fail).
- **Renderers**: Output formatters like Terminal, JSON, Markdown, and HTML.

## Getting Started

### Prerequisites
- Go 1.22+

### Installation

```bash
go install github.com/raphgm/container-doctor@latest
```

### Usage

```bash
# Run all diagnostics
container-doctor check

# Inspect a specific project
container-doctor inspect .

# Output as JSON
container-doctor check --output json
```
