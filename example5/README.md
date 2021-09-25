# Other tips and tricks

## Run tests separately

### Using build tags

You can use build tags to indicate which files will be included in a package during the build process. This can also be used for separating your tests. The examples below demonstrate the rules behind build tags.

The following is an `OR` operation. So this file will only be included if either `tag1` or `tag2` are specified.
```
// +build tag1 tag2
```

The following is an `AND` operation. So this file will only be included if `tag1` and `tag2` are specified.
```
// +build tag1, tag2
```

The following is a `negation` operation. So this file will only be included if `tag1` is not specified.
```
// +build !tag1
```

Consider [example5.1_test.go](./example5.1_test.go), [example5.2_test.go](./example5.2_test.go) and [example5.3_test.go](./example5.3_test.go) with build tags added. Try the following and see what happens.

```
# directory example5

go test -v ./... -tags=feat1
go test -v ./... -tags=feat2
go test -v ./... -tags=feat3
go test -v ./... -tags=feat4
```

### Using -run option

Another way to separate tests is to name your tests in a specific manner and use the [test flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags) `-run` which uses a regex pattern. Try these and observe the output

```
# directory example5

go test -v ./... -tags=all -run=One
go test -v ./... -tags=all -run=Two
go test -v ./... -tags=all -run=Three
```

## testing with time

At times, we might need to deal with test that using `time.Now`. A simple way to test this is to declare a `var` to return a func for `time.Now`. We can override that in our test. See [example5.go](./exampl5.go) and [example5.4_test.go](./example5.4_test.go).


## 3. test helper functions

* pass in `*testing.T` to perform assertions
* don't return `error` as your test code now has to do the error handling
* use [t.Helper](https://pkg.go.dev/testing#T.Helper) for more concise failure reporting
* return cleanup code in a func()

See `testSetEnvVar` in [example5.5_test.go](./example5.5_test.go)
