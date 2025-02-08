# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.1.0] - 2023-11-19
### Added
- Base version of the logging library

## [v1.0.0] - 2024-03-21
### Changed
- Changed the name of EncoderConfig TimeKey to ts_date
- Changed method name `Zap` to `UnZap`
- Method `Zap` now sets the logger
### Added
- Option to add prefix to log keys `WithKeyPrefix`
- Option to add skip caller level `WithCallerSkip`

## [v1.1.0] - 2024-05-29
### Removed
- Removed field `KeyPrefix` from config
### Changed
- Config field `callerSkip` of set public
- Update Goland version to 1.22.3
- Changed [.golangci.yml](.golangci.yml)

## [v1.2.0] - 2024-08-18
### Added
- `CasinoID` field to the config
- Taskfile.yml
### Fixed
- Missing `contextEncoder` field on copy logger

## [v1.2.1] - 2024-09-11
### Added
- `NewDebugConfig` function to create a new logger config for debugging or testing purposes