// Package app serves as a dependency manager
package app

import "wt/internal/git"

type Dependencies struct {
	Git git.Git
}
