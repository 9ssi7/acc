# Accumulator: A Go-based Data Accumulation Mechanism

## Overview

Accumulator is a lightweight Go library designed to aggregate, manage, and process data items. Whether you're accumulating data from various sources, processing them at specific intervals, or simply managing a stream of data, this library provides the tools you need.

## Features

- **Flexible Storage**: Choose from built-in storage options like in-memory storage or integrate with your custom storage mechanism.
- **Customizable Processing**: Define how data items are processed using your custom function.
- **Configurable Interval**: Set the interval at which accumulated data is processed.
- **Scheduled Processing**: Optionally start processing data at a specific time.

## Installation

To install the package, simply run:

```bash
go get github.com/9ssi7/acc
```


## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/9ssi7/acc.svg)](https://pkg.go.dev/github.com/9ssi7/acc)

## Usage

### Basic Setup

Here's a basic example demonstrating how to set up an accumulator:

```go
package main

import (
	"fmt"
	"time"

	"github.com/9ssi7/acc"
)

func main() {
	config := acc.Config[int]{
		Processor: processFunction,
		// ... other configurations
	}

	accumulator := acc.New(config)
	// Start using the accumulator
}

func processFunction(data []int) {
	// Your logic to process data
	fmt.Println("Processing data:", data)
}
```

### Advanced Configurations

For more advanced configurations like choosing a custom storage or setting a specific start time, refer to the documentation or the examples provided.

## Contributing

We welcome contributions! Please see our [Contribution Guidelines](CONTRIBUTING.md) for details.

## License

This project is licensed under the Apache License. See [LICENSE](LICENSE) for more details.