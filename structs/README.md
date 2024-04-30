## Interfaces

---

This is quite different to interfaces in most other programming languages. Normally you have to write code to say My type Foo implements interface Bar.

But in our case

    Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface

    Circle has a method called Area that returns a float64 so it satisfies the Shape interface

    string does not have such a method, so it doesn't satisfy the interface

    etc.

In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

## Table driven tests

Useful when you want to build a list of test cases that can be tested in the same manner.

> The test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations

> Make sure your test output is helpful
