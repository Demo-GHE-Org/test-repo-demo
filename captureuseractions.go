package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// Open the text file for writing
	file, err := os.OpenFile("user_actions.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffered writer for better performance
	writer := bufio.NewWriter(file)

	// Start capturing user actions
	fmt.Println("Capturing user actions... Press Ctrl+C to stop.")
	fmt.Println("User actions will be logged to user_actions.txt")

	// Capture keyboard events
	go func() {
		for {
			if robotgo.AddEvent("k") {
				key := robotgo.LastKey()
				action := "Pressed"
				if robotgo.KeyTap(key) {
					action = "Released"
				}
				logAction(writer, "Keyboard", action, key)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Capture mouse events
	go func() {
		for {
			if robotgo.AddEvent("m") {
				evt := robotgo.GetMouseMsg()
				action := ""
				switch evt.Kind {
				case robotgo.MouseMove:
					action = "Moved"
				case robotgo.MouseWheel:
					action = "Scrolled"
				default:
					action = "Clicked"
				}
				logAction(writer, "Mouse", action, evt.X, evt.Y)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Keep the program running until interrupted
	select {}
}

// Log the user action to the file
func logAction(writer *bufio.Writer, device, action string, args ...interface{}) {
	logTime := time.Now().Format("2006-01-02 15:04:05")
	logString := fmt.Sprintf("[%s] %s %s: %v\n", logTime, device, action, args)
	writer.WriteString(logString)
	writer.Flush()
}

