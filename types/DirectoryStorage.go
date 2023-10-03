package types

import (
	"io/fs"
)

type DirectoryStorage struct {
	Path            string
	InnerComponents []fs.DirEntry
	PreviousDir     *DirectoryStorage
}

func NewDirectoryStorage(path string, entries []fs.DirEntry, previousDir *DirectoryStorage) *DirectoryStorage {
	return &DirectoryStorage{
		Path:            path,
		InnerComponents: entries,
		PreviousDir:     previousDir,
	}
}
