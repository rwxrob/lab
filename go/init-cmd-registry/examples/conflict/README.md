# Example of a Naming Conflict

* [cmds](cmds) - the base library that contains all the commands
* [cmd-hi](hi) - simple `hi` command to say "Hi!"
* [cmd-hello](hello) - simple `hello` command to say "Hello!"
* [cmd-hiagain](hiagain) - another conflicting `hi` command
* [greet](greet) - monolith to contain all greeting commands

In this scenario the `Add()` function neither overrides the previous
addition nor panics because of the conflict. Instead, it adds an
underscore (`_`) to the end of the conflicting command name and provides
a `Rename()` function that can be called from `init()` of the highest
level `greet/main.go` (the same in which the entry point `Run()` is
called). Since the addition of underscore is predictable based on the
set order of `init()` invocations this fix is consistent.

This solution allows any combination of any commands to immediately be
used despite the flat nature of the internal commands registry. It also
avoids unnecessary complications from adding fully qualified command
names (ex `local/cmd-hello` instead of `hello`).

