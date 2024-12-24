# rocket-hexagonal

A brief description of the project, what it does, and its purpose.

---

## Table of Contents

1. [Getting Started](#getting-started)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Configuration](#configuration)
6. [Testing](#testing)
7. [Contributing](#contributing)
8. [License](#license)

---

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

---

## Prerequisites

- Go (minimum version 1.22) installed on your system.
- [Optional] Docker for containerized deployment.

---

## Installation

### Clone the Repository

```bash
git clone github.com/muhfaris/rocket-hexagonal
cd repository
```

### Initialize the Project

Run the following command to prepare the project:

```bash
make init
```

This command will:

- Run `go mod tidy` to clean up dependencies.
- Run `go mod vendor` to vendor dependencies.

### Quick Start

Run the following command to start the services, gomod application and run the application:

```bash
make start
```
---

## Usage

### Running the Application

```bash
go run main.go rest
```

or

```bash
make run
```

### Build the Application

```bash
make build
```

---

## Configuration

Configuration is handled through `config/config.yaml`:

```env
app:
  name: rocket-hexagonal
  port: 7000
  debug:
    config: false
  fiber:
    enable_print_routes: true
    enable_splitting_on_parsers: true
  cache:
    redis:
      addr: localhost:6379
      username:
      password:
      db: 0
  datastore:
    psql:
      host: localhost
      port: 5432
      username: myapp
      password: myapp
      db: myapp
```

---

### Shortcuts Commands

The project comes with a set of shortcuts commands to make development easier:

- `make start` - Run the services, gomod application and run the application.
- `make services` - Start necessary services and remove old volume.
- `make run` - Run the application.
- `make build` - Build the application for deployment and store it in the `$(GOPATH)/bin` directory.
- `make init` - Run `go mod tidy` and `go mod vendor` commands.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).
