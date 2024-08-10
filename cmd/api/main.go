package main

import (
	"fmt"
	"gin-be/internal/database"
	_ "gin-be/internal/ent/runtime" // important to make hooks working
	"gin-be/internal/server"
	"gin-be/internal/tool"
	"log"
)

func main() {
	fmt.Println("Starting api usaha")
	envData := tool.NewEnv(nil)
	db := database.New()
	db.Migrate()
	log.Default().Println("DB's Health")
	log.Default().Println(db.Health())
	srv := server.NewServer()
	err := srv.Run(":" + envData.PORT)
	if err != nil {
		log.Fatalf("failed starting server on port %s: %v", envData.PORT, err)
	}
}
