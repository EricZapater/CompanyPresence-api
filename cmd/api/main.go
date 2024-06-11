package main

import (
	"companypresence-api/internal/handlers"
	"companypresence-api/internal/repositories"
	"companypresence-api/internal/server"
	timetrackingservice "companypresence-api/internal/services/timetracking"
	userservice "companypresence-api/internal/services/users"
	workscheduleservice "companypresence-api/internal/services/workschedule"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//App para registrar usuarios
//Crear clockins y clockouts
//Cálcular clockins y clockouts
func main() {
	fmt.Println("Running APP")
	timeTrackingRepository := repositories.NewTimeTrackingRepository()
	userRepository := repositories.NewUserRepository()
	userService := userservice.NewUserService(*userRepository)
	workScheduleRepository := repositories.NewWorkScheduleRepository()
	workscheduleservice := workscheduleservice.NewWorkScheduleService(*workScheduleRepository)

	timeTrackingService := timetrackingservice.NewTimeTrackingService(*timeTrackingRepository, *userService, *workscheduleservice)
	userHandler := handlers.NewUserHandler(userService)

	app := server.Setup(userHandler)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go timeTrackingService.Scheduler()

	go func(){
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-quit
	fmt.Println("Shutting down app")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("App shutdown failed %v", err)
	}
	fmt.Println("App stopped")
}