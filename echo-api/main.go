package main

import (
	"echo-api/controller"
	"echo-api/db"
	"echo-api/repository"
	"echo-api/router"
	"echo-api/usecase"
	"echo-api/validator"
)

func main() {
	db := db.NewDB()

	// user
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	// task
	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
