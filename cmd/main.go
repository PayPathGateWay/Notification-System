// cmd/main.go
package main

import (
	"Notification-System/internal/presentation"
)

func main() {
	print("Starting the notification system...\n")
	presentation.StartThegRPCServer()
}
