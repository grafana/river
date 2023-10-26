# Changelog

This document contains a historical list of changes between releases. Only
changes that impact end-user behavior are listed; changes to documentation or
internal API changes are not present.

Main (unreleased)
-----------------

v0.3.0 (2023-10-26)
-------------------

### Enhancements

- Allow the `[]` operator to return `null` when accessing non-existant keys in
  objects, rather than returning an error. For example, `{}["foo"]` now returns
  `null`. (@rfratto)

### Bugfixes

- Fix a bug where indexing an object with a non-string key would generate a
  type error informing users to supply a `number` instead of the actually
  expected `string` type. (@rfratto)

v0.2.0 (2023-10-20)
-------------------

### Features

- Add support for raw strings in river using backticks (@erikbaranowski)

- Add functions for validating and sanitizing River identifiers to the scanner package (@erikbaranowski)

v0.1.1 (2023-08-25)
-------------------

### Other changes

- Fix typos and expand README for documentation.

- `token/builder`: Update River encode handling of optional fields to compare values using
  DeepEqual even if they don't implement Equal when deciding if the field
  should be included. (@erikbaranowski)

v0.1.0 (2023-08-25)
-------------------

> First release!

### Features

- Publish a `riverfmt` binary for formatting River files.

- Publish River as a library:

  - `github.com/grafana/river/ast` contains the AST representation of River with some utilities.
  - `github.com/grafana/river/diag` contains types for River diagnostics (errors and warnings).
  - `github.com/grafana/river/encoding/riverjson` contains utilities to print River bodies as JSON.
  - `github.com/grafana/river/parser` contains utilities to parse River files.
  - `github.com/grafana/river/printer` contains utilities to format River files.
  - `github.com/grafana/river/rivertypes` contains useful capsule values.
  - `github.com/grafana/river/scanner` contains utilities to scan River files.
  - `github.com/grafana/river/token` contains token definitions for River.
  - `github.com/grafana/river/token/builder` contains utilities to build River files from Go code.
  - `github.com/grafana/river/vm` evalutes River blocks and expressions

  The top-level `github.com/grafana/river` module contains a high-level API for
  unmarshaling River files to Go types and marshaling Go types to River files.
