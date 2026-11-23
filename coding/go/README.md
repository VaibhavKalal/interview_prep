# Go Coding Practice

One package per problem: `coding/go/<pattern>/<problem>/`, with a solution file and a table-driven test.

## Workflow

1. Read the problem in the comment at the top of the `.go` file.
2. Say your approach and its time/space complexity out loud (as you would in an interview).
3. Implement it. Replace the `panic("TODO")` line.
4. Run the tests:

```bash
go test ./coding/go/arrays/twosum/
```

5. Add a test case of your own (edge cases: empty input, duplicates, negatives).
6. Update the progress table in the root README.

Run everything:

```bash
go test ./...
```

Check for races (useful once you reach concurrency problems):

```bash
go test -race ./...
```

## Adding a new problem

Copy an existing folder, rename the package, write the problem statement as a comment, and add test cases first.
