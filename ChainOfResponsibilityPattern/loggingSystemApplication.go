package main

func main() {
	logger := NewLogProcessor(NewInfoLogProcessor(NewDebugLogProcessor(NewErrorLogProcessor(nil))))

	logger.log(DEBUG, "Debugging request")
	logger.log(ERROR, "Failed to serve request")
	logger.log(INFO, "Request successfully served")
}
