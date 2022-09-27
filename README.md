# levelog
Log levels for golang's built in [log](https://pkg.go.dev/log) library 

## Usage
```go
import (
    "github.com/matansh/levelog"
    "github.com/matansh/levelog/loglevel"
)

func main() {
    ctx := context.Background()

    logger := levelog.NewLogger(
        loglevel.Info,
        log.New(os.Stdout, "", log.LstdFlags),
        log.New(os.Stderr, "", log.LstdFlags),
    )
    // from this line onward panics will be logged as errors
    defer logger.PanicRecovery()
    // embedding the logger instance into the applications context
    ctx = levellogger.ContextWithLogger(ctx, logger)

    logger.Info("tada")
}
```
### Usage in your applications tests
```go
import "github.com/matansh/levelog"

func Test(t *testing.T) {
    logger := levelog.NewTestLogger(t)
    ctx := levelog.ContextWithLogger(context.Background(), logger)

    logger.Debug("I am a log statement that will show up as part of the tests output")
    // log statements made as part of the tested code will also show up as part of the tests output
    TestedCode(ctx)
}
```
### Testing against logs
Outputted log statements can be asserted in tests by passing a bytes buffer as the loggers output
```go
import "github.com/matansh/levelog"

func Test(t *testing.T) {
    stdout := bytes.Buffer{}
    stderr := bytes.Buffer{}

    logger := levelog.NewLogger(
        loglevel.Debug,
        log.New(&stdout, "", 0),
        log.New(&stderr, "", 0),
    )
    ctx := levelog.ContextWithLogger(context.Background(), logger)

    TestedCode(ctx)
    // assertions regarding what is or is not in the buffers
```
As an example, see the [tests implemented for the lib itself](logger_test.go)

## Background
While there are plenty of logging libraries in the golang sphere I failed to find one that simply wrapped the languages pre existing logging capabilities in a level based interface.

This library does exactly that, it is a nicer packaging for something we already know and love.

### footnote
This library is intentionally dependency-less in order to minimize the dependency trees of its importers, you are welcome ;)