# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2025-09-28
### Fixed
- Fixed CSV writer delimiter issue where uninitialized delimiter was set to null character, causing empty output
- Enhanced CSV writer to properly flatten nested JSON objects for CSV conversion
- Added support for flattening nested structures like {"address": {"city": "NYC"}} to address_city=NYC in CSV output

## [0.1.0] - 2025-09-28
### Added
- Initial release with support for JSON, CSV, YAML, XML, and TOML conversion
- Automatic format detection
- Batch processing support
- Command-line interface with various options (--to, --pretty, --batch, --validate)
- Schema inference for optimal output structure