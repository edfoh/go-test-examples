# Testable Examples

Godoc examples are snippets of Go code that can be executed in your package documentation. These are also executed as normal tests, so it ensures that your documentation is always kept up to date with your API!

The way to do that is to add an [example_test.go](./example_test.go). We have a func called `SayHi` in [example5.go](./example5.go), and in [example_test.go](./example_test.go), the test func needs to start with `Example`, so in this case `func ExampleSayHi`. This is recognized as a test and will run as part of your test suite.

