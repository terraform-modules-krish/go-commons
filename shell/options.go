package shell

import (
	"github.com/Sirupsen/logrus"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/logging"
)

type ShellOptions struct {
	NonInteractive bool
	Logger         *logrus.Entry
	WorkingDir     string
}

func NewShellOptions() *ShellOptions {
	return &ShellOptions{
		NonInteractive: false,
		Logger: logging.GetLogger(""),
		WorkingDir: ".",
	}
}
