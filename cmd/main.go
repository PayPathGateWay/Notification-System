// cmd/main.go
package main

import (
	"Notification-System/internal/gRPC/handlers"
)

func main() {
	print("Starting the notification system...\n")
	handlers.StartThegRPCServer()
	//if len(os.Args) < 2 {
	//	fmt.Println("Expected 'api' or 'scheduler' subcommands")
	//	os.Exit(1)
	//}
	//
	//cfg := config.MustLoadConfig()
	//
	//switch os.Args[1] {
	//case "api":
	//	api := internal.NewAPI(cfg)
	//	api.StartAPI()
	//case "scheduler":
	//	scheduler := internal.NewScheduler(cfg)
	//	scheduler.StartScheduler()
	//default:
	//	fmt.Println("Unknown command. Expected 'api' or 'scheduler'")
	//	os.Exit(1)
	//}
}
