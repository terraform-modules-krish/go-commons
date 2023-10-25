//go:build gittest

package gittest

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/terraform-modules-krish/go-commons/files"
	"github.com/terraform-modules-krish/go-commons/git"
	"github.com/terraform-modules-krish/go-commons/logging"
	"github.com/terraform-modules-krish/terratest/modules/environment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	gitPATEnvName = "GITHUB_OAUTH_TOKEN"
)

var (
	logger = logging.GetLogger("testlogger")
)

// NOTE: All these tests should be run in the provided docker environment to avoid polluting the local git configuration
// settings. The tests will assert that it is running in the docker environment, and will fail if it is not.
// All these tests are also run in serial to avoid race conditions on the git config file.

func TestHTTPSAuth(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)
	require.Equal(t, "/workspace/go-commons/git/test", currentDir)

	environment.RequireEnvVar(t, gitPATEnvName)
	gitPAT := os.Getenv(gitPATEnvName)
	require.NoError(t, git.ConfigureHTTPSAuth(logger, "git", gitPAT, "github.com"))

	tmpDir, err := ioutil.TempDir("", "git-test")
	require.NoError(t, err)
	require.NoError(t, git.Clone(logger, "https://github.com/gruntwork-io/terraform-aws-lambda.git", tmpDir))
	assert.True(t, files.IsDir(filepath.Join(tmpDir, "modules/lambda")))
}

func TestForceHTTPS(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)
	require.Equal(t, "/workspace/go-commons/git/test", currentDir)

	environment.RequireEnvVar(t, gitPATEnvName)
	gitPAT := os.Getenv(gitPATEnvName)
	require.NoError(t, git.ConfigureHTTPSAuth(logger, "git", gitPAT, "github.com"))
	require.NoError(t, git.ConfigureForceHTTPS(logger))

	tmpDir, err := ioutil.TempDir("", "git-test")
	require.NoError(t, err)
	require.NoError(t, git.Clone(logger, "git@github.com:gruntwork-io/terraform-aws-lambda.git", tmpDir))
	assert.True(t, files.IsDir(filepath.Join(tmpDir, "modules/lambda")))
}