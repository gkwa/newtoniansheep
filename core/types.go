package core

type Processor interface {
	Process(input []string) ([]string, error)
}

type FileHandler interface {
	Read(path string) ([]string, error)
	Write(path string, content []string) error
}
