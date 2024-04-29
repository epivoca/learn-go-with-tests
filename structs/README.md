## Interfaces

This is quite different to interfaces in most other programming languages. Normally you have to write code to say My type Foo implements interface Bar.

But in our case

    Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface

    Circle has a method called Area that returns a float64 so it satisfies the Shape interface

    string does not have such a method, so it doesn't satisfy the interface

    etc.

In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
