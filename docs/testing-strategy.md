# Testing Strategy

Manny uses a series of test tables to validate internal functions. The test tables and functions can be divided into the
following categories:

- Examples
- Negative Tests
- Unit tests

### Examples

The purpose of testing against the `examples` directory is two fold:
- We establish a contract for use of Manny so that we don't unintentionally break core functionality
- We educate users on how to properly use Manny without detailed instruction pages

### Negative Tests

The `tests` directory consists of our negative tests and showcases how _not_ to use Manny and establishes a contract of
what will definitively be an end-user error.

### Unit Tests

Although all of the above tests are ran in unit style testing, they actually accomplish more E2E level testing but with
better introspection/coverage support.

These tests are specifically more unit test oriented and make singular method calls and evaluate output.

### Conventions

- All tests contain a case named "Valid" that demonstrates how to use the function or method.
- All tests use a field named "Identifier" to differentiate one test from the other.
- All tests are runnable by `go test` and contribute to coverage
- Core functionality should include a test in `examples`