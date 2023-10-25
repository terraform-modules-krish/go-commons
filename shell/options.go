package shell

import (
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/logging"
	"github.com/sirupsen/logrus"
)

type ShellOptions struct {
	NonInteractive bool
	Logger         *logrus.Logger
	WorkingDir     string
	SensitiveArgs  bool // If true, will not log the arguments to the command
}

func NewShellOptions() *ShellOptions {
	return &ShellOptions{
		NonInteractive: false,
		Logger:         logging.GetLogger(""),
		WorkingDir:     ".",
		SensitiveArgs:  false,
	}
}
