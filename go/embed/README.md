This is an experiment to see how much RAM is consumed by a file that is
embedded with `//go:embed 1gbfile`.

To create a file of a specific size use `fallocate` (just don't commit
it to git).

```
fallocate -l $((1*1024*1024*1024)) 1gbfile
```

Note that `go:embed` is limited to files that are less than
approximately 2GB:

```
$ go build -o foo
# github.com/rwxrob/lab/go/embed
./main.go:11:5: embed 2gbfile: file too large (2147483648 bytes > 2000000000 bytes)
```

Now see `main.go` for how to add it in and build it with `go build -o
foo`. Note that it will take quite a bit longer to build/compile when
embedding a large file (or collection of files).

Note that the size is slightly more than the embedded file.

```
-rwxrwxr-x 1 rwxrob rwxrob 1.1G Aug  1 15:40 foo
```

Now run it and watch RAM utilization.

```
watch free -h
```

Note that the memory is consumed while building, but while running if
only the `fs.FileInfo` is queried the memory for the embedded file
is not consumed.


Also check out.

```
readelf -Wl foo
size foo
objdump -s -j .rodata foo
strace
https://www.youtube.com/watch?v=zDKM7uXsKrQ
```

## Conclusion

The Go `embed` package adds the files and directories to the `.rodata`
section of the ELF binary where there are managed just like any other
constant by the kernel. Most importantly, they *do not consume RAM until
they are explicitly read* (just as with a regular file system). Queries
of the `fs.FileInfo` data do *not* load any portion of the actual file
into memory and are safe to use with `Stat` (and family) without
triggering buffered file loads, which must be explicit. However, such
behavior cannot be guaranteed consistently for every operating system
since this is highly dependent on the architecture of each OS kernel.
