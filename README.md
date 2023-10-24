# Transact Ease API Documentation

## Introduction

An intuitive transaction management system, designed for seamless card operations and account handling.

## Production URL

The live application is deployed on [Heroku](https://www.heroku.com/) and can be accessed at [https://transact-ease-605ea8e7ab48.herokuapp.com/](https://transact-ease-605ea8e7ab48.herokuapp.com/).

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Project Structure](#project-structure)
3. [Postman Documentation](#postman-documentation)
4. [Postman Collection](#postman-collection)
5. [Installation](#installation)
6. [Database Structure](#database-structure)
7. [Usage](#usage)
8. [Endpoints](#endpoints)
9. [Testing](#testing)
10. [Release and Deployment](#release-and-deployment)

## Prerequisites

To run local, you will need to create an `.env` file on `internal/config`. Take the example and populate it.

### Docker

Ensure you have Docker and Docker Compose installed. If not, follow the installation guides:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Project Structure

Below is the directory structure of the Transact Ease API:

```
├── cmd
│   ├── api
│   │   ├── main.go
│   │   └── server
│   │       └── server.go
│   ├── helpers
│   │   └── common_helper.go
│   ├── migration
│   │   ├── main.go
│   │   └── migrations
│   │       ├── 1_create_accounts_table.up.sql
│   │       ├── 1_drop_accounts_table.down.sql
│   │       ├── 2_create_operations_types_table.up.sql
│   │       ├── 2_drop_operations_types_table.down.sql
│   │       ├── 3_create_transactions_table.up.sql
│   │       ├── 3_drop_transactions_table.down.sql
│   │       ├── 4_create_index_to_transaction_account_query.up.sql
│   │       └── 4_drop_index_to_transaction_account_query.down.sql
│   └── seed
│       ├── main.go
│       └── seeds
│           ├── 1_seed_accounts.sql
│           ├── 2_seed_operations_types.sql
│           └── 3_seed_transactions.sql
├── go.mod
├── go.sum
├── infra
│   └── local
│       ├── docker-compose.yaml
│       ├── init.sql
│       └── postgres
│           └── data
│               ├── PG_VERSION
│               ├── base
│               ├── global
│               ├── pg_commit_ts
│               ├── pg_dynshmem
│               ├── pg_hba.conf
│               ├── pg_ident.conf
│               ├── pg_logical
│               ├── pg_multixact
│               ├── pg_notify
│               ├── pg_replslot
│               ├── pg_serial
│               ├── pg_snapshots
│               ├── pg_stat
│               ├── pg_stat_tmp
│               ├── pg_subtrans
│               ├── pg_tblspc
│               ├── pg_twophase
│               ├── pg_wal
│               ├── pg_xact
│               ├── postgresql.auto.conf
│               ├── postgresql.conf
│               ├── postmaster.opts
│               └── postmaster.pid
├── internal
│   ├── config
│   │   ├── config.go
│   │   ├── setup_postgre_config.go
│   │   └── sqlx_driver.go
│   ├── constants
│   │   ├── env_constants.go
│   │   ├── error_constants.go
│   │   ├── logger_constants.go
│   │   ├── success_constants.go
│   │   └── time_constants.go
│   ├── controllers
│   │   ├── accounts
│   │   │   ├── account_controller.go
│   │   │   └── account_controller_test.go
│   │   ├── base_response.go
│   │   ├── base_response_test.go
│   │   └── transactions
│   │       ├── transaction_controller_test.go
│   │       └── transction_controller.go
│   ├── domains
│   │   ├── accounts
│   │   │   ├── account.go
│   │   │   ├── account_interface.go
│   │   │   └── account_mapper.go
│   │   ├── database_interface.go
│   │   └── transactions
│   │       ├── transaction.go
│   │       ├── transaction_interface.go
│   │       └── transaction_mapper.go
│   ├── dto
│   │   ├── requests
│   │   │   ├── account_request.go
│   │   │   ├── account_request_test.go
│   │   │   ├── transaction_request.go
│   │   │   └── transaction_request_test.go
│   │   └── responses
│   │       ├── account_response.go
│   │       ├── account_response_test.go
│   │       ├── transaction_response.go
│   │       └── transaction_response_test.go
│   ├── integrations
│   ├── mocks
│   │   ├── accounts
│   │   │   ├── account_repository.go
│   │   │   └── account_usecase.go
│   │   ├── mock_database.go
│   │   └── transactions
│   │       └── transaction_usecase.go
│   ├── repositories
│   │   ├── account_repository.go
│   │   ├── account_repository_test.go
│   │   └── transaction_repository.go
│   ├── routes
│   │   ├── accounts
│   │   │   └── routes.go
│   │   ├── router.go
│   │   └── transactions
│   │       └── routes.go
│   └── usecases
│       ├── account_usecase.go
│       ├── account_usecase_test.go
│       └── transaction_usecase.go
├── makefile
├── pkg
│   ├── helpers
│   │   ├── utils_helper.go
│   │   └── utils_helper_test.go
│   ├── logger
│   │   ├── log_logger.go
│   │   └── logrus_logger.go
│   └── validators
│       ├── payload_validator.go
│       └── payload_validator_test.go
└── scripts
```

This structure provides an organized view of the main components of the application.

## Postman Documentation

To access the documentation click [here](https://documenter.getpostman.com/view/10569183/2s9YRDyqKZ)

## Postman Collection

To test the API endpoints, you can use the provided Postman collection. [Download the collection here]() and import it into your Postman application.

To the **local** environment use [localhost:8080](http://localhost:3000) as the `url` variable value.

To the **production** environment use [https://transact-ease-605ea8e7ab48.herokuapp.com/](https://transact-ease-605ea8e7ab48.herokuapp.com/) as the `url` variable value.

## Installation

1. **Clone the Repository:**
   ```bash
   git clone git@github.com:OvictorVieira/transact.ease.git
   cd transact.ease
   ```

2. **Use `go mod download` to download the dependencies:**
   ```bash
   go mod download
   ```

3. **Build and Start the Docker Containers:**
   ```bash
   make docker-up
   ```

## Database Structure

### Accounts Table

| Column Name     | Data Type | Description                       |
|-----------------|-----------|-----------------------------------|
| id              | Integer   | Primary Key                       |
| document_number | String    | Document number                   |
| created_at      | Tiemstamp | Date that the account was created |
| updated_at      | Tiemstamp | Date that the account was updated |

### Operation Types Table

| Column Name     | Data Type | Description                                                            |
|-----------------|-----------|------------------------------------------------------------------------|
| id              | Integer   | Primary Key                                                            |
| description     | String    | Operation description (e.g., CASH PURCHASE, INSTALLMENT PURCHASE, etc) |
| password_digest | String    | Encrypted password                                                     |

### Transactions Table

| Column Name       | Data Type | Description                           |
|-------------------|-----------|---------------------------------------|
| id                | Integer   | Primary Key                           |
| account_id        | Integer   | Id of an account                      |
| operation_type_id | Integer   | Id of a operation type                |
| amount            | Float     | Value of a transaction                |
| event_date        | Tiemstamp | Date that the transaction occurred    |
| created_at        | Tiemstamp | Date that the transaction was created |
| updated_at        | Tiemstamp | Date that the transaction was updated |

## Usage

- **Configurations:** Ensure all configurations are set correctly in the `config` directory.
- **Starting the Server:** Use `run-server` to start the server.
- **Using the API:** Use tools like `curl` or Postman to interact with the API.

## Endpoints

- **Accounts:**
    - `GET /api/accounts/:accountId`: Fetches a specific account.
    - `POST /api/accounts`: Creates a new account.

- **Transactions:**
    - `POST /api/transactions`: Creates a new transaction.

## Testing

- **Running Tests:** Use `make run-tests` to run the tests.

## Test Coverage

At Transact Ease API, we prioritize the quality and reliability of our codebase. To ensure that our application is thoroughly tested, we utilize the `coverage` gem to monitor our test coverage.

### Checking Test Coverage

- **Running Tests:** Use `make run-tests-with-cover` to run the tests with coverage.

## Release and Deployment

### Publishing via Tags

At Transact Ease API, we believe in structured and organized releases. To ensure consistency and traceability, we use a tag-based release system. This means that every significant change or release of the application is marked with a specific tag, following the [Semantic Versioning](https://semver.org/) convention.

#### How it Works:

1. **Development & Testing:** Once a feature is developed and thoroughly tested, it's merged into the main branch.
2. **Tagging:** Before a release, a new tag is created. This tag represents the version of the application. For example, `v1.0.0`.
3. **Deployment:** Our CI/CD pipeline is triggered by the creation of tags. Once a new tag is pushed, the pipeline automates the deployment process, ensuring that the tagged version of the application is deployed.
4. **Accessing Releases:** You can view all the tags and releases in the "Releases" section of our GitHub repository. Each release will have detailed notes and changes that were introduced in that version.

#### Benefits:

- **Traceability:** Each release is traceable, and you can easily revert to a previous version if needed.
- **Stability:** By releasing through tags, we ensure that only stable and tested features make it to production.
- **Clarity:** Each tag comes with release notes, providing a clear understanding of the changes.
