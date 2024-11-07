# Fetcher image URI from API third-party

## Requirement

## Architeture project folder

The project folder structure for a Go-based tool to fetch image URLs from a third-party API can be organized as follows:

- `cmd/`: Contains the main applications for the project.
- `pkg/`: Contains library code that's ok to use by external applications.
- `internal/`: Contains private application and library code.
- `api/`: Contains code related to API interactions.
- `config/`: Contains configuration files and related code.
- `scripts/`: Contains scripts for various tasks like building, testing, etc.
- `docs/`: Contains documentation files.
- `test/`: Contains test-related files and test data.
