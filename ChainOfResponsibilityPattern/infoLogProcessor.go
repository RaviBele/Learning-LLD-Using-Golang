package main

import "fmt"

type InfoLogProcessor struct {
	nextLogProcessor ILogProcessor
}

func NewInfoLogProcessor(nextLogProcessor ILogProcessor) *InfoLogProcessor {
	return &InfoLogProcessor{nextLogProcessor: nextLogProcessor}
}

func (p *InfoLogProcessor) log(logType LogType, message string) {
	if logType == INFO {
		fmt.Println("INFO: ", message)
	} else if p.nextLogProcessor != nil {
		p.nextLogProcessor.log(logType, message)
	}
}
