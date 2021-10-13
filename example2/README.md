# Test Setup and Teardown

It is somewhat common (especially coming from different languages) to have test frameworks that support test case setups and teardowns. 
Some common examples that would be useful to have setup and teardown

* loading a file and cleaning that up afterwards
* database setup and teardown

I will share two different ways to achieve these using some examples.

## Creating setup and teardown functions

One approach is to create your own setup and teardown functions that you can add to each test. A convenient way is to use `defer` to invoke your teardown functions. [example2.1_test.go](./example2.1_test.go) illustrates this.

## Using TestMain in Go 1.4

In Go 1.4, [TestMain](https://pkg.go.dev/testing#hdr-Main) was introduced. As mentioned in the [release notes](https://golang.org/doc/go1.4)

> The testing package has a new facility to provide more control over running a set of tests. If the test code contains a function
>```
>func TestMain(m *testing.M)
>```
> that function will be called instead of running the tests directly. The M struct contains methods to access and run the tests.

[example2.2_test.go](./example2.2_test.go) illustrates this.

---

Besides test setup and teardown, this can also be a useful way to test subprocesses. 

Another use-case is to ensure code runs in the main goroutine that is locked to the main OS thread, where the tests can be run in another goroutine. For example, graphics related code / libraries are very particular about this. Consider

```
func init() {
    // LockOSThread wires the calling goroutine to its current operating system thread
    runtime.LockOSThread() 
}

func TestMain(m *testing.Main) {
    go func() {
        os.Exit(m.Run())
    }()
    runCodeInMainThread()
}
```

The example above ensures `runCodeInMainThread()` runs in the main thread, and the package tests will assumingly be testing APIs from that function.

