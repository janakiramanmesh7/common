package logging

import (
	"fmt"
	"os"
	"github.com/sirupsen/logrus"
)

// Setup configures a global logrus logger to output to stderr.
// It populates the standard logrus logger as well as the global logging instance.
func Setup(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("error parsing log level: %v", err)
	}
	if err != nil {
		return err
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	SetGlobal(Logrus(logrus.StandardLogger()))
	return nil
}
