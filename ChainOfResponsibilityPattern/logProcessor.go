package main

type LogType int

const (
	INFO LogType = iota + 1
	DEBUG
	ERROR
)

func (o LogType) String() string {
	return [...]string{"INFO", "DEBUG", "ERROR"}[o-1]
}

type ILogProcessor interface {
	log(logtype LogType, message string)
}

type LogProcessor struct {
	nextLogProcessor ILogProcessor
}

func NewLogProcessor(nextLogProcessor ILogProcessor) *LogProcessor {
	return &LogProcessor{nextLogProcessor: nextLogProcessor}
}

func (p *LogProcessor) log(logtype LogType, message string) {
	if p.nextLogProcessor != nil {
		p.nextLogProcessor.log(logtype, message)
	}
}
