The `github.com/matansh/levelog/log` package is intended to provide pre-existing projects an easy migration path away from the built in `log` package.

Simply replace any `import "log"` statements with `import "github.com/matansh/levelog/log"` for a painless and refactoring-less transition.

On a personal note: \
I do not encourage the use of globals, I would highly recommend instantiating and reusing your own logger instance.
