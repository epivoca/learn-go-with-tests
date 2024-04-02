# Some stuff to learn

#### The problem about making our code passing new tests () is we can either

- Break the existing API by changing the argument to Sum to be a slice rather than an array. When we do this, we will potentially ruin someone's day because our other test will no longer compile!

- Create a new function

In my case, no one else is using my function, so rather than having two functions to maintain, let's have just one.

#### It is important to question the value of your tests.

It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. **Every test has a cost**.

#### Excursion into history

According to the author of the course:

You could write a function to iterate over each got and want slice and check their values but for convenience sake, we can use `reflect.DeepEqual` which is useful for seeing if any two variables are the same.

It's important to note that `reflect.DeepEqual` is not "type safe" - the code will compile even if you did something a bit silly.
