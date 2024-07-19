package application

func NewLinuxFileSystem() FileSystem {
	return NewDirectory("/")
}
