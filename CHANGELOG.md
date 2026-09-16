# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Changed

- `PostgresConfig.Port` changed from `string` to `int` and is now validated
  against the `0-65535` range (fixes #22, #13).
- MongoDB env variable `MONGODB_DBNAME` renamed to `MONGODB_DB_NAME`.
  The old name is still accepted with a deprecation warning (fixes #19).

### Added

- `PostgresConfig.Mode`: `"pool"` (default) or `"single"`, to declare whether
  the application uses a connection pool or a single connection.
- `PostgresConfig.Pool` with pool tuning knobs: `maxConns`, `minConns`,
  `maxConnLifetime`, `maxConnIdleTime`, `healthCheckInterval` (in seconds).
- `PostgresConfig.SSLMode`, `PostgresConfig.AppName`, `PostgresConfig.ConnTimeout`.
- `PostgresConfig.DSN()` helper: returns the URI when set, otherwise composes
  a connection string from the decomposed fields.
- Optional config sections absent from the CUE file can now be activated via
  environment variables (scoped to `postgres` for now, fixes #14).
- Environment variable overrides are re-validated after loading, so values
  such as `POSTGRES_MODE=garbage` fail the build (scoped to postgres, fixes #12).
