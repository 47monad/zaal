# Zaal

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/47monad/zaal.svg)](https://pkg.go.dev/github.com/47monad/zaal)
![Status: Under Development](https://img.shields.io/badge/Status-Under%20Development-orange)

> Named after Zāl (زال), the legendary Persian hero and father of Rostam in Ferdowsi's Shahnameh.

## Overview

Zaal is a robust configuration management package for Go applications that combines the power of [CUE](https://cuelang.org/) with environment variable handling. It allows you to:

1. Load structured configuration from CUE files
2. Override configuration values with environment variables
3. Access your configuration in a type-safe manner

## ⚠️ Development Status

**This package is in heavy development and is not ready for production use yet.**

Features and APIs may change significantly before the first stable release.

## Installation

```bash
go get github.com/47monad/zaal
```

## Basic Usage

```go
package main

import (
 "fmt"
 "log"

 "github.com/your-repo/zaal"
)

func main() {
 // Load configuration from CUE file
 var cfg zaal.Config
 
 // Option 1: Load config file directly
 if err := zaal.Build("config.cue", ".env"); err != nil {
  log.Fatalf("Failed to load config: %v", err)
 }
 
 // Use your configuration
 fmt.Printf("Application: %s v%s\n", cfg.Name, cfg.Version)
 fmt.Printf("Environment: %s\n", cfg.Env)
 
 if cfg.Mongodb != nil {
  fmt.Printf("MongoDB URI: %s\n", cfg.Mongodb.URI)
 }
}
```

## Configuration Structure

Zaal supports a flexible configuration structure with nested objects, optional components, and automatic environment variable binding:

```go
type Config struct {
 Name       string            `json:"name"`
 Title      string            `json:"title"`
 Version    string            `json:"version"`
 Env        string            `json:"env" env:"env"`
 Mode       string            `json:"mode" env:"mode"`
 Host       string            `json:"host" env:"host"`
 Logging    LoggingConfig     `json:"logging"`
 Mongodb    *MongodbConfig    `json:"mongodb,omitempty"`
 RabbiMQ    *RabbitMQConfig   `json:"rabbitmq,omitempty"`
 Prometheus *PrometheusConfig `json:"prometheus,omitempty"`
 GRPC       *GRPCConfig       `json:"grpc,omitempty"`
 HTTP       *HTTPConfig       `json:"http,omitempty"`
}
```

## Environment Variable Binding

Zaal automatically binds environment variables to your configuration based on the `env` struct tags:

```go
// Config struct field with env tag
Env string `json:"env" env:"env"`

// Environment variable: ENV="production"
// Result: cfg.Env == "production"
```

For nested fields, environment variables are bound in a flattened structure:

```go
// Config struct
type Config struct {
 Mongodb *MongodbConfig `json:"mongodb,omitempty"`
}

// MongodbConfig struct
type MongodbConfig struct {
 URI string `json:"uri" env:"mongodb_uri"`
}

// Environment variable: MONGODB_URI="mongodb://localhost:27017"
// Result: cfg.Mongodb.URI == "mongodb://localhost:27017"
```

## Features

- **Type Safety**: Strongly typed configuration with Go structs
- **Validation**: CUE schema validation ensures your configuration is correct
- **Environment Overlay**: Override configuration with environment variables
- **Optional Components**: Use pointer fields for optional configuration sections
- **Flexible Loading**: Load from files, environment, or both

## Roadmap

- [ ] Comprehensive documentation
- [ ] Schema validation using CUE
- [ ] Configuration merging from multiple sources
- [ ] Command-line argument support
- [ ] Secrets management integration
- [ ] Hot reloading support

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
