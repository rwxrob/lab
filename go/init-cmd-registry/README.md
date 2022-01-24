# Learn to Use Go `init()` for Command Registry

This lab presents ways that `init()` can be leveraged in its
"side effects" form (unnamed with `_`) to create a registry of commands
available to the caller from `main()`.

## Summary

This lab helps you explore different ways to leverage `init()` to create
importable modularity for a `greet` command that supports different
languages:

```
greet hi
greet hello
greet fr hi
greet fr hello
greet sp hi
greet sp hello
greet ru hi
greet ru hello
```

French and Spanish should be loaded as side-effects modules that
register their commands by `cmds.Add(...)` themselves into a
centralizing package `cmds`.

```go
import (
  ...
  _ "somewhere.io/user/greet-french
  _ "somewhere.io/user/greet-spanish
)
```

Then add a few *local* languages into their own `*.go` files with their
own `func init() {}` functions that are identical to lose imported but
just don't need to be imported.

## Caveats

* **Don't `go mod init` Top Level.** If you do a `go mod init` at the
  top level (the same as this `README.md`) you will overtake the other
  `go.mod` files within this repo and get strange results with your
  imports. Instead, create a `greet` directory.

* **Use `replace` in `go.mod`.** If you are going to use a mock host
  provider like `somewhere.io/user` then you will need to add `replace`
  directives pointing to the relative paths for these in their own
  directories. Otherwise, you *could* use private repos on GitHub or
  somewhere, but that is a bad idea and you should learn `replace`
  anyway.

## Requirements

* No command module should ever override the commands of another.
* Provide a `Rename()` function to fix conflicts
