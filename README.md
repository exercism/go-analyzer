# Exercism's Go Analyzer

This is Exercism's automated analyzer for the Go track.

## Executing the Analyzer

The analyser takes two parameters:
- the exercise `slug`, e.g. `two-fer`
- the `path` to the solution to analize

Example to execute with binary:

```bash
analyze two-fer ~/solution-238382y7sds7fsadfasj23j/
```

From source with Go installed:

```bash
go run ./main.go two-fer ~/solution-238382y7sds7fsadfasj23j/
```

## Build Executable

This will create an executable called `analyze`.

```bash
go generate .
go build -tags build -o analyze .
```

`go generate` is called before the build to incorporate all necessary files within the binary.

## Docker

To `build` execute the following from the repositories `root` directory:

```bash
docker build -t exercism/go-analyzer .
```

To `run` from docker pass in the solutions path as a volume and execute with the necessary parameters:

```bash
docker run -v $(PATH_TO_SOLUTION):/solution exercism/go-analyzer ${SLUG} /solution
```

Example:

```bash
docker run -v ~/solution-238382y7sds7fsadfasj23j:/solution exercism/go-analyzer two-fer /solution
```


## Stats

### Twofer

Out of 500 real world solutions we get:

```
approve_as_optimal      10
approve_with_comment    24
disapprove_with_comment 463
refer_to_mentor         3
ejected (failed)        0
```
