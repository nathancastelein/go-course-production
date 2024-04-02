# Exercise 4

## Install benchstat

From your home directory, let's install [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) to compare results:

```
$ go install golang.org/x/perf/cmd/benchstat@latest
```

## Run benchmarks

First of all, let's try to run a first benchmark:

```
go test -bench .
```

Read the result, then try to understand it.

### Current implementation

Now, run the benchmark 10 times and output the result into a file, named `old.txt`.

```
go test -count 10 -benchmem -bench .
```

See `go help testflag` to understand the flags.

### New implementation

Your lead tech comes with a great idea: let's use another library to generate request id if not found in the request.

They propose to use this one: [https://github.com/nu7hatch/gouuid](https://github.com/nu7hatch/gouuid)

Rewrite the logger to generate an uuid using this package.

Once the change is done, run the benchmark 10 times and output the result into a file, named `new.txt`.

## Compare benchmarks

Once you have two files, `old.txt` and `new.txt`, compare them with `benchstat`:

```
benchstat old.txt new.txt
```

Which implementation is the most performant one?