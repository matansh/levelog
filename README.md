# levelog
Log levels for golang's built in [log](https://pkg.go.dev/log) library 

## Usage
```go
import "github.com/matansh/levelog/log"

func main() {
    ctx := context.Background()

    logger := levellogger.NewLogger(
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

## Background
While there are plenty of logging libraries in the golang sphere I failed to find one that simply wrapped the languages pre existing logging capabilities in a level based interface.

This library does exactly that, it is a nicer packaging for something we already know and love.

### footnote
This library is intentionally dependency-less in order to minimize the dependency trees of its importers, you are welcome ;)