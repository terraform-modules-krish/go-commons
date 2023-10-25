package git

import (
	"github.com/terraform-modules-krish/go-commons/files"
	"github.com/terraform-modules-krish/go-commons/shell"
)

// Clone runs git clone to clone the specified repository into the given target directory.
func Clone(repo string, targetDir string) error {
	if !files.IsDir(targetDir) {
		return TargetDirectoryNotExistsErr{dirPath: targetDir}
	}

	opts := shell.NewShellOptions()
	return shell.RunShellCommand(opts, "git", "clone", repo, targetDir)
}

// Checkout checks out the given ref for the repo cloned in the target directory.
func Checkout(ref string, targetDir string) error {
	if !files.IsDir(targetDir) {
		return TargetDirectoryNotExistsErr{dirPath: targetDir}
	}

	opts := shell.NewShellOptions()
	opts.WorkingDir = targetDir
	return shell.RunShellCommand(opts, "git", "checkout", ref)
}
