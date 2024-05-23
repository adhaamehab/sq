# SQ

A strong consistency FIFO queue with at-least-once delivery semantics.


## Installation

Download the binary from the [releases page](https://github.com/adhaamehab/sq/releases) and add it to your PATH.

## Usage

### Starting the server on port 6565

```bash
sq-server --port 6565
```

### CLI

```bash
sq-cli

> PING # PONG
> ENQUEUE foo
> ENQUEUE bar
> DEQUEUE # foo
> DEQUEUE # bar
> DEQUEUE # ""
> ENQUEUE 1
> ENQUEUE 2.1
> DEQUEUE # 1
> DEQUEUE # 2.1
> DEQUEUE # ""
```


## Features

### V1 (MVP)
- [ ] A standalone server.
- [ ] A basic CLI client.
- [ ] Basic queue operations: `ENQUEUE`, `DEQUEUE`.
- [ ] Support JSON payloads.
- [ ] Examples:
  - [ ] A Go pub-sub application
  - [ ] A Python pub-sub application
- [ ] Proper logging.
- [ ] Presistent storage.
- [ ] Tests
  - [ ] Unit tests
  - [ ] E2E tests
- [ ] Benchmarking

### V2
Coming soon...