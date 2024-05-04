# HA-TDD-Chess

- **Author:** Jim Törnvall
- **Year:** 2024
- **School:** Åland University of Applied Sciences
- **Course:** Test Driven Development
- **Status:** [![Go Build and Test](https://github.com/JimTornvall/HA-TDD-Chess/actions/workflows/go.yml/badge.svg)](https://github.com/JimTornvall/HA-TDD-Chess/actions/workflows/go.yml)

## Info

Chess constructor and testing for school

## Packages
- Testify
    - https://github.com/stretchr/testify
    - Assert: Used for assertions in the tests
    - Suite: Used for creating test suites, used to mimic @BeforeEach and @AfterEach in JUnit
    - Mock: Isnt used as i like Mockio better
- Mockio
    - https://github.com/ovechkin-dm/mockio
    - Mockito for golang, used for mocking interfaces in the tests

## Setup

```shell
go mod tidy
```

## Run

```shell
go run .
```

## Test

```shell
go test -v ./...
```

## TODO

- board move needs to use piece move, to fully get all functionality
  - verify that this is the case
- Future:
  - refactor into more packages
  - rewrite to support
    - charm/bubbletea - TUI interface
    - charm/wish - SSH interface
- switch all movement and movement checks to use error instead of bool and error
- **Board is the wrong way** switch places of BLACK and WHITE, in init and piecePawn
- Extend asserts with more information, so you see what went wrong
- Make sure all pieces test illegal moves
- Check when a player moves if opponent king becomes in check