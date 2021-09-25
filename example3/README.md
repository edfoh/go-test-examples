# Test Fixtures and Golden Files

## Test fixtures

It is sometimes necessary to store test files. Taken from [Go Test Packages documentation](https://pkg.go.dev/cmd/go#hdr-Test_packages)

>The go tool will ignore a directory named "testdata", making it available to hold ancillary data needed by the tests.

See [example3.1_test.go](./example3.1_test.go) to see how this works.

## Golden files

> A "golden file" is the expected output of some test (usually automated), stored as a separate file rather than as a string literal inside the test code. So when the test is executed, it will read in the file and compare it to the output produced by the system under test.

These usually are outputs that are cumbersome to include inside your test code. The idea would be to allow your code to generate the output, and you would manually check the output for correctness, instead of attempting to craft the actual output yourself manually. See [example3.2_test.go](./example3.2_test.go) for an example of how this work.

Notice there's a `update flag` set in the code. This allows you to run the test in `write mode` where you can save the generated output into the specified golden files, stored in the `testdata` directory. By default, that flag is set to false. To save the generated golden files, you would run this

```
go test example3.2_test.go -update
```
