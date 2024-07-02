package main

import "fmt"

type DebugLogProcessor struct {
	nextLogProcessor ILogProcessor
}

func NewDebugLogProcessor(nextLogProcessor ILogProcessor) *DebugLogProcessor {
	return &DebugLogProcessor{nextLogProcessor: nextLogProcessor}
}

func (p *DebugLogProcessor) log(logType LogType, message string) {
	if logType == DEBUG {
		fmt.Println("DEBUG: ", message)
	} else if p.nextLogProcessor != nil {
		p.nextLogProcessor.log(logType, message)
	}
}
