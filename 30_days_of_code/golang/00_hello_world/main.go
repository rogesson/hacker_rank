package main

import "fmt"
import "bufio"
import "os"

func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    messages := []string{"Hello, roger.", text}

    PrintMessages(messages)
}

    func PrintMessages(messages []string) bool {
    for message := range messages {
        fmt.Println(message)
    }

    return true
}
