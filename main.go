package main

import "todo-list/config"

func main() {
	client := config.ConnectDB()
	defer config.DisconnectDB(client)\

	config.StartServer(routes)
	select {}
}
