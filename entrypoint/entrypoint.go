package entrypoint

import (
	"os"
	"github.com/urfave/cli"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/errors"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/logging"
)

const defaultSuccessExitCode = 0
const defaultErrorExitCode = 1

// Run the given app, handling errors, panics, and stack traces where possible
func RunApp(app *cli.App) {
	defer errors.Recover(checkForErrorsAndExit)
	err := app.Run(os.Args)
	checkForErrorsAndExit(err)
}

// If there is an error, display it in the console and exit with a non-zero exit code. Otherwise, exit 0.
func checkForErrorsAndExit(err error) {
	exitCode := defaultSuccessExitCode

	if err != nil {
		logging.GetLogger("").WithError(err).Error()

		errorWithExitCode, isErrorWithExitCode := err.(errors.ErrorWithExitCode)
		if isErrorWithExitCode {
			exitCode = errorWithExitCode.ExitCode
		} else {
			exitCode = defaultErrorExitCode
		}
	}

	os.Exit(exitCode)
}
