# Blackbox vs Whitebox testing

Prefer blackbox testing (`_test package`) over whitebox testing (`same package`). Why?

* remove cyclic dependencies in test
* testing your exported functions and using them like a public API. Helps determine if easy to use.
* reduce brittleness if you overly rely on testing unexported functions.
* tests are less coupled, makes refactoring easier.


## Cyclic Dependency

An example of cycle dependency in whitebox testing:

```
/--A
  /--a.go
  /--a_test.go
/--B
  /--b.go
```

where `package B` relies on `package A`. `a_test.go` also uses the same package `A` and in the test, it relies on a function in `Package B`. This results in a cyclic dependency. The simple way around this is to make `a_test.go` under `package A_test`. See [A/a_test.go](./A/a_test.go) to see how this works.

## How to test unexported (private) functions then?

The golang standard library uses an interesting way to make export private functions for testing, as seen in this [export_test.go in fmt package](https://github.com/golang/go/blob/master/src/fmt/export_test.go). The same example is show here using [export_test.go](./export_test.go) and [example4_test.go](./example4_test.go)

Another method is to mark a test file as doing whitebox testing, such as [example4_internal_test.go](./example4_internal_test.go), using `_internal_test.go` to make it obvious.
