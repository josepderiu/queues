# Go Queues Project

This project implements a queue data structure in Go.

## Project Structure

The project has the following file structure:

- `go.mod`: This file defines the module and dependencies of the project.
- `queues.go`: This file contains the implementation of the queue data structure.
- `queues_test.go`: This file contains the unit tests for the queue data structure.

## Usage

To use the queue data structure, you first need to import it in your Go file:

`import "github.com/josepderiu/queues"`

Then, you can create a new queue and use its `Enqueue`, `Dequeue`, and `PrintElements` methods:

`q := &queues.Queue{}`
`q.Enqueue(1)`
`q.Enqueue(2)`
`q.PrintElements()`
`q.Dequeue()`

## Tests

To run the unit tests, use the following command in your terminal:

`go test [-v] [--cover]`


## Versioning

This project uses [SemVer](https://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/josepderiu/queues/tags).