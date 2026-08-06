package core

import (
	"github.com/elzorrorebelde/remedy/internal/pkg/fs"
	"github.com/elzorrorebelde/remedy/internal/pkg/helper"
	"github.com/elzorrorebelde/remedy/internal/pkg/vcs"
)

func DefaultEnvironment() *Environment {
	return &Environment{
		VersioningClient: &vcs.Client{
			Vcs: &vcs.Git{},
		},
		FileSystem:     fs.DefaultFileSystem(),
		Clock:          helper.SystemClock{},
		SchemaLocation: "https://raw.githubusercontent.com/elzorrorebelde/remedy/master/docs/schema.json",
	}
}

type Environment struct {
	VersioningClient vcs.VersioningClient
	FileSystem       *fs.FileSystem
	Clock            helper.Clock
	SchemaLocation   string
}
