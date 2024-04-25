# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Updated golang.org/x/sys so the project can be built on mac and with go 1.18. See: https://github.com/golang/go/issues/49219
- Upgrade dependencies
- Upgrade to Go 1.22

### Fixed

- Ensure separator obey the --no-color flag

## [0.1.2] - 2022-02-21

### Changed

- No changes since [0.1.1]

## [0.1.1] - 2020-07-06

### Added

- Format new microerror@0.2.x JSON formatted error.
- Print operatorkit resource name in bold.

### Fixed

- Print default color instead of white.

## [0.1.0] - 2020-07-01

### Changed

- First release.

[Unreleased]: https://github.com/giantswarm/luigi/compare/v0.1.2...HEAD
[0.1.2]: https://github.com/giantswarm/luigi/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/giantswarm/luigi/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/luigi/releases/tag/v0.1.0
