package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

/*
Initializes the logger and makes it available globally
via

	zip.ReplaceGlobals()

method. The logger can be accessed
using

	zap.L()

method.
*/
func Init() {
	logger, err := zap.NewDevelopment()

	if err != nil {
		fmt.Println("could not instantiate logger, terminating the program")
		os.Exit(1)
	}

	defer logger.Sync()

	zap.ReplaceGlobals(logger)
}
