# Exercism's Go Analyzer

This is Exercism's automated analyzer for the Go track.

## Executing the Analyzer

The analyser takes two parameters:
- the exercise slug, e.g. `two_fer`
- the path to the solution to analize

Example to execute with binary:

```bash
analyze two_fer ~/solution-238382y7sds7fsadfasj23j/
```

From source with Go installed:

```bash
go run ./main.go two_fer ~/solution-238382y7sds7fsadfasj23j/
```

## Build Executable

This will create an executable called `analyze`.

```bash
go generate .
go build -o analyze .
```

`go generate` is called before the build to incorporate all necessary files within the binary.
