package main

import (
	"echo-api/controller"
	"echo-api/db"
	"echo-api/repository"
	"echo-api/router"
	"echo-api/usecase"
)

func main() {
	db := db.NewDB()

	// user
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	// task
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
