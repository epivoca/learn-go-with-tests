# A lesson to learn

It is good to separate your "domain" code from the outside world (side-effects).
Talking about "hello world", the `fmt.Println` is a side-effect, and the string we send in is our domain.

The idea is to separate these concerns so it's easier to test

### Random thoughts about `go.mod`

```bash
require example.com/othermodule v1.2.2
replace example.com/othermodule => example.com/othermodule v1.2.3
```

Use the replace directive to temporarily substitute a module path value with another value when you want Go to use the other path to find the module’s source. This has the effect of redirecting Go’s search for the module to the replacement’s location. You needn’t change package import paths to use the replacement path.

Use the exclude and replace directives to control build-time dependency resolution when building the current module. These directives are ignored in modules that depend on the current module.

The replace directive can be useful in situations such as the following:
- You’re developing a new module whose code is not yet in the repository. You want to test with clients using a local version.
- You’ve identified an issue with a dependency, have cloned the dependency’s repository, and you’re testing a fix with the local repository.

Note that a replace directive alone does not add a module to the module graph.
A require directive that refers to a replaced module version is also needed, either in the main module’s go.mod
file or a dependency’s go.mod file.
If you don’t have a specific version to replace, you can use a fake version, as in the example below.
Note that this will break modules that depend on your module, since replace directives are only applied in the main module.
