package main

import (
	"redis-task-queue/pkg/redis"
	"redis-task-queue/src/controllers"
	"redis-task-queue/src/repositories"
	"redis-task-queue/src/services"
	"redis-task-queue/src/workers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initialize redis
	redisClient := redis.NewRedisClient("localhost:6379", "", 0)

	// initialize repository, service, and controller
	taskRepo := repositories.NewTaskRepository(redisClient, redis.Ctx)
	taskService := services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)

	// start worker
	go workers.StartWorker(taskRepo, "task_queue")

	// setup fiber app
	app := fiber.New()

	// routes
	app.Post("/task", taskController.AddTask)
	app.Get("/task/:id/status", taskController.GetTaskStatus)

	// start server
	app.Listen(":3000")
}
