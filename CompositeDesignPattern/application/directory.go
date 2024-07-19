package application

import "errors"

type Directory struct {
	FileSystem
	name      string
	filenames []string
	files     map[string]FileSystem
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name:  name,
		files: make(map[string]FileSystem),
	}
}

func (d *Directory) Mkdir(directory string) error {
	if _, ok := d.files[directory]; ok {
		return errors.New("directory already exists")
	}
	d.files[directory] = NewDirectory(directory)
	d.filenames = append(d.filenames, directory)
	return nil
}

func (d *Directory) Ls() ([]string, error) {
	return d.filenames, nil
}

func (d *Directory) Cd(directory string) (FileSystem, error) {
	if _, ok := d.files[directory]; !ok {
		return nil, errors.New("Directory does not exist")
	}

	switch d.files[directory].(type) {
	case *Directory:
		return d.files[directory], nil
	default:
		return nil, errors.New("given directory name does not exist")
	}
}

func (d *Directory) Touch(filename string) error {
	if _, ok := d.files[filename]; ok {
		return errors.New("file already exists")
	}
	d.files[filename] = NewFile(filename)
	d.filenames = append(d.filenames, filename)
	return nil
}
