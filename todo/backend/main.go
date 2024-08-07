package main

import (
		"github.com/Matama091/go_exercise/todo/backend/config"
		"github.com/Matama091/go_exercise/todo/backend/routers"
)

func main() {
    config.InitDB()
    defer config.DB.Close()

    r := routers.SetupRouter()
    r.Run(":8080")
}
