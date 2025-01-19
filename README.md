# prettylogs

A simple logging library in Go. Pretty logs currently is supporting the following features, with some of them planned for the future:

- [x] Logging different log levels (e.g. Debug, Warning, Error, etc.)
- [x] Support for any io.Writer
- [x] Ability to disable prefixes (e.g. `[DEBUG]`, `[INFO]`, etc.)
- [x] Add timestamps option
- [ ] Add flag to enable colorful output to os.Stdout
- [ ] Add structured logging support with JSON

# ⚙️ Installation

```
go get github.com/tatucosmin/prettylogs
```

# 🔨 Example

```go
package main

import (
	plogs "github.com/tatucosmin/prettylogs"
)

func main() {
	logger := plogs.New()
	/*
		New will intialize with this data as default
		- os.Stdout as the io.Writer
		- LogInfoLevel which will only print logs with level >= LogInfoLevel
		- prefixes (e.g. [DEBUG], [INFO], etc) enabled
		- timestamps enabled
		note: if you need more granular Logger initialization use the NewConfigurable function
	*/

	logger.Warn("this is a warn log\n")
	// prints "[WARN] this is a warn log" without changing the default level

	logger.Log("this is an info log\n")
	// prints "[INFO] this an info log" as by default the logger has the LogInfoLevel set

	logger.LogWithLevel(plogs.LogDebugLevel, "this message will NOT print\n")
	logger.SetLoggerLevel(plogs.LogDebugLevel)
	logger.LogWithLevel(plogs.LogDebugLevel, "this message will print\n")

	/*
		final output:
		2025-01-19 15:22:31 [WARN] this is a warn log
		2025-01-19 15:22:31 [INFO] this is an info log
		2025-01-19 15:22:31 [DEBUG] this message will print
	*/
}
```

# 🙋‍♂️ Why?

I wanted to practice my Go-lang skills and was finding myself in the situation where this was all around my code which I found confusing:

```go
fmt.Printf("debug value: %v\n")
```

To simplify this, I created a library that provides a cleaner and more structured way to debug values with minimal boilerplate.

# 👏 Contributing

I'd love some help on this project. If you'd like to contribute, please fork this project and open pull requests.

Please make sure all existing tests pass before submitting your pull request and your own tests if that's the case.
