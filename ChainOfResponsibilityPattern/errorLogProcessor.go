package main

import "fmt"

type ErrorLogProcessor struct {
	nextLogProcessor ILogProcessor
}

func NewErrorLogProcessor(nextLogProcessor ILogProcessor) *ErrorLogProcessor {
	return &ErrorLogProcessor{nextLogProcessor: nextLogProcessor}
}

func (p *ErrorLogProcessor) log(logType LogType, message string) {
	if logType == ERROR {
		fmt.Println("ERROR: ", message)
	} else if p.nextLogProcessor != nil {
		p.nextLogProcessor.log(logType, message)
	}
}
