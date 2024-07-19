package application

import "errors"

type File struct {
	FileSystem
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) Ls() ([]string, error) {
	return []string{f.name}, nil
}

func (f *File) Mkdir(directory string) error {
	return errors.New("mkdir cannot be used with file")
}

func (f *File) Touch(path string) error {
	return nil
}

func (f *File) Cd(directory string) (FileSystem, error) {
	return nil, errors.New("cd cannot be used with file")
}
