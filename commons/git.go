package commons

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GitCloneRepository(url, directory string) error {
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return err
}

func GitIsValidGitRepository(directory string) bool {
	_, err := git.PlainOpen(directory)
	return err == nil
}

func GitPullRepository(directory string) error {
	// Open an existing repository
	r, err := git.PlainOpen(directory)
	if err != nil {
		return err
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return err
	}

	return nil
}

// func GitCheckoutBranch to checkout branch form name
func GitCheckoutBranch(directory, name string) error {
	// Open an existing repository
	r, err := git.PlainOpen(directory)
	if err != nil {
		return err
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Get the branch reference
	branch := plumbing.NewRemoteReferenceName("origin", name)

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Checkout(&git.CheckoutOptions{Branch: branch})
	if err != nil {
		return err
	}

	return nil
}

func GitResetUncommittedChanges(directory string) error {
	r, err := git.PlainOpen(directory)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = w.Reset(&git.ResetOptions{Mode: git.HardReset})
	if err != nil {
		return err
	}

	return nil
}
