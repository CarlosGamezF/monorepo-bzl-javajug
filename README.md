# Monorepo with Bazel

This monorepo demonstrates how to set up and manage multiple microservices using [Bazel](https://bazel.build/) as the build tool. It includes services written in Spring Boot (Java), Go, Python, and a React frontend, each with their own tests. The setup also supports different build flavors, such as `dev` and `prod`. Bazel was chosen for its ability to handle large, multi-language projects efficiently, providing fast builds, reproducible results, and scalability.

## Prerequisites

- [Bazel](https://bazel.build/) (it is recommended to use [Aspect Cli](https://docs.aspect.build/cli/) for version management)

## Setup

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd monorepo-bzl-javajug
   ```

2. Ensure Bazel is installed. If not, you can install [Installation](https://docs.aspect.build/cli/#installation):

## Project Structure

The monorepo is structured as follows:

| Directory          | Description                                                                 |
|--------------------|-----------------------------------------------------------------------------|
| `java_service/`    | Spring Boot microservice with source code, configuration files, and tests   |
| `go_service/`      | Go microservice with source code and tests                                  |
| `python_service/`  | Python microservice with source code, dependencies, and tests               |
| `react_frontend/`  | React frontend with source code, npm dependencies, and tests                |
| `WORKSPACE`        | Bazel workspace configuration                                               |
| `MODULE.bazel`      | Externals deps                                                             |
| `.bazelrc`         | Bazel configuration file                                                    |

This structure allows for modular development and easy management of multiple services within a single repository.

## Building and Running

To build all services:
```bash
bazel build //...
```

To run a specific service:
- **Java (Spring Boot)**: `bazel run //java_service:java_service` (listens on port 8080 for dev, 80 for prod)
- **Go**: `bazel run //go_service:go_service` (listens on port 8081)
- **Python**: `bazel run //python_service:python_service` (listens on port 8082)
- **React**: Build with `bazel build //react_frontend:react_frontend`, then serve with `python -m http.server -d bazel-bin/react_frontend` and access at http://localhost:8000

### Using Build Flavors

The project supports different build flavors, such as `dev` and `prod`. For the Java service, this affects the configuration files used (e.g., `application-dev.properties` for dev, `application-prod.properties` for prod). To build or run with a specific flavor, use the `--config` flag:
```bash
bazel build //java_service:java_service --config=prod
bazel run //java_service:java_service --config=prod
```

For other services, the flavor might not have an effect unless specifically configured.

## Testing

To run tests for a specific service:
- **Java**: `bazel test //java_service:java_service_test`
- **Go**: `bazel test //go_service:go_service_test`
- **Python**: `bazel test //python_service:python_service_test`
- **React**: `bazel test //react_frontend:react_frontend_test`

Tests are integrated into the Bazel build system, allowing for easy execution with `bazel test`.

##  Bazel by JetBrains (EAP) for IntelliJ IDEA
This plugin is available as EAP (Early Access Program). This means that it may already be useful, but many corner cases might not be supported yet.

* [IntelliJ IDEA](https://plugins.jetbrains.com/plugin/22977-bazel-eap-)

## Contributing

Contributions are welcome! Please follow the standard pull request workflow.

## License

[Insert license information if applicable]

## Best Practices

- This monorepo follows Bazel's best practices for structuring projects, including modular organization and clear dependency management.
- For more details on Bazel best practices, refer to the [official documentation](https://docs.bazel.build/versions/master/best-practices.html).
- Dependencies for each language are managed through the `WORKSPACE` file, ensuring reproducibility and consistency across environments.

Repository used to https://www.meetup.com/madridjug/events/307175844/?eventOrigin=group_upcoming_events
