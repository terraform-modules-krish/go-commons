***WARNING: THIS REPO IS AN AUTO-GENERATED COPY.*** *This repo has been copied from [Gruntwork’s](https://gruntwork.io/) GitHub repositories so that you can consume it from your company’s own internal Git repositories. This copy is automatically created and updated by the `repo-copier` CLI tool. If you need to make changes to this repo, you should make the changes in a separate fork, and NOT make changes directly in this repo, as otherwise, the `repo-copier` will overwrite your changes! Please see the `repo-copier` [documentation](https://github.com/terraform-modules-krish/repo-copier) for more information on how the code is copied, how cross-references are updated, how the changelog is handled, etc.*

***

_You may find it valuable to view the following resources in the original repo. If these links give you a 404, visit https://app.gruntwork.io to gain access or email support@gruntwork.io if you need assistance._

[Home Page](https://github.com/gruntwork-io/go-commons/) |
[Pull Requests](https://github.com/gruntwork-io/go-commons/pulls) |
[Issues](https://github.com/gruntwork-io/go-commons/issues) |
[Releases and Assets](https://github.com/gruntwork-io/go-commons/releases)

_Alternatively, you can view a copied version of the resources listed above._

[Pull Requests](https://github.com/terraform-modules-krish/go-commons/blob/master/.github/PULL_REQUESTS.md) |
[Issues](https://github.com/terraform-modules-krish/go-commons/blob/master/.github/ISSUES.md) |
[ChangeLog](https://github.com/terraform-modules-krish/go-commons/blob/master/.github/CHANGELOG.md)

***

[![Maintained by Gruntwork.io](https://img.shields.io/badge/maintained%20by-gruntwork.io-%235849a6.svg)](https://gruntwork.io/?ref=repo_go-commons)
[![Go Report Card](https://goreportcard.com/badge/github.com/gruntwork-io/go-commons)](https://goreportcard.com/report/github.com/gruntwork-io/go-commons)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/gruntwork-io/go-commons?tab=overview)
![go.mod version](https://img.shields.io/github/go-mod/go-version/gruntwork-io/go-commons)

# Gruntwork Go Commons

This repo contains common libraries and helpers used by Gruntwork for building CLI tools in Go.

**NOTE: This repo was renamed from gruntwork-cli starting v0.8.0. Update your gruntwork-cli references to go-commons.**

## Packages

This repo contains the following packages:

* collections
* entrypoint
* errors
* files
* logging
* shell
* ssh
* retry
* awscommons

Each of these packages is described below.

### collections

This package contains useful helper methods for working with collections such as lists and maps, as Go has very few of
these built-in.

### entrypoint

Most CLI apps built by Gruntwork should use this package to run their app, as it takes
care of common tasks such as setting the proper exit code, rendering stack
traces, handling panics, and rendering help text in a standard format. Note
that this package assumes you are using
[urfave/cli](https://github.com/urfave/cli), which is currently our library of
choice for CLI apps.

Here is the typical usage pattern in `main.go`:

```go
package main

import (
        "github.com/urfave/cli"
        "github.com/terraform-modules-krish/go-commons/entrypoint"
        "github.com/terraform-modules-krish/go-commons/version"
)

func main() {
      // Create a new CLI app. This will return a urfave/cli App with some
      // common initialization.
      app := entrypoint.NewApp()

      app.Name = "my-app"
      app.Author = "Gruntwork <www.gruntwork.io>"

      // Set the version number from your app from the Version variable that is passed in at build time in `version` package
      // for more understanding see github.com/gruntwork-io/go-commons/version
      app.Version = version.GetVersion()

      app.Action = func(cliContext *cli.Context) error {
        // ( fill in your app details)
        return nil
      }

      // Run your app using the entrypoint package, which will take care of exit codes, stack traces, and panics
      entrypoint.RunApp(app)
}
```

### errors

In our CLI apps, we should try to ensure that:

1. Every error has a stacktrace. This makes debugging easier.
1. Every error generated by our own code (as opposed to errors from Go built-in functions or errors from 3rd party
   libraries) has a custom type. This makes error handling more precise, as we can decide to handle different types of
   errors differently.

To accomplish these two goals, we have created an `errors` package that has several helper methods, such as
`errors.WithStackTrace(err error)`, which wraps the given `error` in an Error object that contains a stacktrace. Under
the hood, the `errors` package is using the [go-errors](https://github.com/go-errors/errors) library, but this may
change in the future, so the rest of the code should not depend on `go-errors` directly.

Here is how the `errors` package should be used:

1. Any time you want to create your own error, create a custom type for it, and when instantiating that type, wrap it
   with a call to `errors.WithStackTrace`. That way, any time you call a method defined in our own code, you know the
   error it returns already has a stacktrace and you don't have to wrap it yourself.
1. Any time you get back an error object from a function built into Go or a 3rd party library, immediately wrap it with
   `errors.WithStackTrace`. This gives us a stacktrace as close to the source as possible.
1. If you need to get back the underlying error, you can use the `errors.IsError` and `errors.Unwrap` functions.
1. If you want to return an error that forces a specific exit code, wrap it with `errors.ErrorWithExitCode`.

Note that `entrypoint.RunApp` takes care of showing stack traces and handling exit codes.

### files

This package has a number of helpers for working with files and file paths, including one-liners for checking if a
given path is a file or a directory, reading a file as a string, and building relative and canonical file paths.

### logging

This package contains utilities for logging from our CLI apps. Instead of using Go's built-in logging library, we are
using [logrus](github.com/sirupsen/logrus), as it supports log levels (INFO, WARN, DEBUG, etc), structured logging
(making key=value pairs easier to parse), log formatting (including text and JSON), hooks to connect logging to a
variety of external systems (e.g. syslog, airbrake, papertrail), and even hooks for automated testing.

To get a Logger, call the `logging.GetLogger` method:

```go
logger := logging.GetLogger("my-app")
logger.Info("Something happened!")
```

To change the logging level globally, call the `SetGlobalLogLevel` function:

```go
logging.SetGlobalLogLevel(logrus.DebugLevel)
```

Note that this ONLY affects loggers created using the `GetLogger` function **AFTER** you call `SetGlobalLogLevel`, so
you need to call this as early in the life of your CLI app as possible!

### shell

This package contains two types of helpers:

* `cmd.go`: This file contains helpers for running shell commands.
* `prompt.go`: This file contains helpers for prompting the user for input (e.g. yes/no).

### ssh

This package contains helper methods for initiating SSH connections and running commands over the connection.

### retry

This package contains helper methods for retrying an action up to a limit.

### awscommons

This package contains routines for interacting with AWS. Meant to provide high level interfaces used throughout various Gruntwork CLIs.

Note that the routines in this package are adapted for `aws-sdk-go-v2`, not v1 (`aws-sdk-go`).


## Running tests

```
go test -v ./...
```

## License

This code is released under the MIT License. See [LICENSE.txt](LICENSE.txt).
