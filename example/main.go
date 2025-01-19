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
		note: if you need more granular Logger initialization use the NewConfigurable function
	*/

	logger.Warn("this is a warn log\n")
	// this will print "[WARN] this is a warn log" without changing the default level

	logger.Log("this is an info log\n")
	// this will print "[INFO] this an info log" as by default the logger has the LogInfoLevel set

	logger.LogWithLevel(plogs.LogDebugLevel, "this message will NOT print\n")
	logger.SetLoggerLevel(plogs.LogDebugLevel)
	logger.LogWithLevel(plogs.LogDebugLevel, "this message will print\n")

	/*
		final output:
		[WARN] this is a warn log
		[INFO] this is an info log
		[DEBUG] this message will print
	*/
}
