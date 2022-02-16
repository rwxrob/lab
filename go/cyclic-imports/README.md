# Breaking Cyclical Imports Problem with Interfaces

> "Pass interfaces, return structs (usually)"

1. Create a subpackage (because subpackages are awesome)
1. Create a package or command in parent directory
1. Create a dependency from parent to subpackage
1. Create a dependency from subpackage to parent
1. Create an interface in both parent and subpackage
1. Allow subpackage interface to be subset of parent interface
