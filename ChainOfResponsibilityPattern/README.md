
Implementing Chain of Responsibility Pattern in Golang for a Log Processor

Design patterns are essential tools for software developers, enabling them to solve common problems in a structured and efficient manner. One such pattern is the Chain of Responsibility pattern, which is particularly useful for handling requests in a decoupled and flexible way. In this blog post, we'll explore the Chain of Responsibility pattern through a practical example of a log processor, implemented in Golang.
What is the Chain of Responsibility Pattern?

The Chain of Responsibility pattern allows an object to send a command without knowing which object will handle it. This pattern creates a chain of receiver objects, where each object in the chain holds a reference to the next object. When a request is sent to the chain, each object can either handle the request or pass it to the next object in the chain.
Use Case: Log Processor

In logging, we often need to process logs at different levels (e.g., INFO, DEBUG, ERROR). Using the Chain of Responsibility pattern, we can create a series of log processors where each processor handles a specific log level. If a processor cannot handle the log, it passes the log to the next processor in the chain.
Implementation in Golang

Let's walk through the implementation of a log processor using the Chain of Responsibility pattern in Golang.
Step 1: Define the Log Levels

First, we'll define the log levels as constants:

```
package main

const (
    INFO = iota
    DEBUG
    ERROR
)
```

Step 2: Create the LogProcessor Interface

Next, we'll define the LogProcessor interface. This interface will have a single method LogMessage, which will be implemented by all concrete log processors.

```
package main

type LogProcessor interface {
    LogMessage(level int, message string)
}
```

Step 3: Implement the Concrete Log Processors

Now, let's implement concrete log processors for each log level: InfoLogProcessor, DebugLogProcessor, and ErrorLogProcessor. Each processor will have a reference to the next processor in the chain.

```
package main

import "fmt"

type BaseLogProcessor struct {
    next LogProcessor
}

func (b *BaseLogProcessor) SetNext(next LogProcessor) {
    b.next = next
}

type InfoLogProcessor struct {
    BaseLogProcessor
}

func (i *InfoLogProcessor) LogMessage(level int, message string) {
    if level == INFO {
        fmt.Println("INFO:", message)
    } else if i.next != nil {
        i.next.LogMessage(level, message)
    }
}

type DebugLogProcessor struct {
    BaseLogProcessor
}

func (d *DebugLogProcessor) LogMessage(level int, message string) {
    if level == DEBUG {
        fmt.Println("DEBUG:", message)
    } else if d.next != nil {
        d.next.LogMessage(level, message)
    }
}

type ErrorLogProcessor struct {
    BaseLogProcessor
}

func (e *ErrorLogProcessor) LogMessage(level int, message string) {
    if level == ERROR {
        fmt.Println("ERROR:", message)
    } else if e.next != nil {
        e.next.LogMessage(level, message)
    }
}
```

Step 4: Set Up the Chain

We'll now set up the chain of log processors. The main function will create instances of the log processors and link them together.

```
package main

func main() {
    errorLogProcessor := &ErrorLogProcessor{}
    debugLogProcessor := &DebugLogProcessor{}
    debugLogProcessor.SetNext(errorLogProcessor)
    infoLogProcessor := &InfoLogProcessor{}
    infoLogProcessor.SetNext(debugLogProcessor)

    infoLogProcessor.LogMessage(INFO, "This is an info message.")
    infoLogProcessor.LogMessage(DEBUG, "This is a debug message.")
    infoLogProcessor.LogMessage(ERROR, "This is an error message.")
}
```

Explanation

1. BaseLogProcessor: This struct provides a base implementation for setting the next processor in the chain.
2. InfoLogProcessor, DebugLogProcessor, ErrorLogProcessor: These structs implement the LogProcessor interface. Each processor checks if it can handle the log level. If it can, it processes the log. Otherwise, it passes the log to the next processor in the chain.
3. main: In the main function, we create instances of the log processors and set up the chain by linking them together. Then, we test the chain by logging messages at different levels.

Conclusion

The Chain of Responsibility pattern is a powerful tool for creating flexible and decoupled systems. In this example, we've demonstrated how to use this pattern to build a log processor in Golang. By following this approach, you can easily extend the system by adding new log processors without modifying the existing code.



