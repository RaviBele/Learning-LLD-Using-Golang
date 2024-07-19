package application

type FileSystem interface {
	Mkdir(directory string) error
	Ls() ([]string, error)
	Touch(path string) error
	Cd(directory string) (FileSystem, error)
}
