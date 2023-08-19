# spinners

A small sandbox for spinning spinners in different directions.

## Setup

A local version of [`briandowns/spinner`][spinner] can be used for easy
modification. Prepare this with the following command:

```sh
git clone https://github.com/briandowns/spinner ../../briandowns/spinner
```

To run with the latest version, comment this line in `go.mod`:

```go
replace github.com/briandowns/spinner v1.23.0 => ../../briandowns/spinner
```

## Running

Build the branch and test any recent changes with the following commands:

```sh
$ make build
$ ./wait
```

### Flags

For a fancier experience, flags can be used to try different cases. There's
only one now:

```sh
# Capture termination signals while the spinner is active
$ ./wait --signal
```

Without a flag, the spinner will automatically spin after a moment.

### Interrupts

More information on testing interrupts and signal behaviors is shared in
[`SIGNALS.md`][signals].

<!-- a collection of links -->
[spinner]: https://github.com/briandowns/spinner
[signals]: ./SIGNALS.md
