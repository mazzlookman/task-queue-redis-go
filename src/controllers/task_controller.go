package controllers

import (
	"redis-task-queue/src/services"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	TaskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

func (contr *TaskController) AddTask(c *fiber.Ctx) error {
	queueName := "task_queue"

	taskId, err := contr.TaskService.CreateTask(queueName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"errors": err.Error()},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"task_id": taskId,
	})
}

func (contr *TaskController) GetTaskStatus(c *fiber.Ctx) error {
	taskId := c.Params("id")

	status, err := contr.TaskService.GetTaskStatus(taskId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{"errors": "Task not found"},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": status,
	})
}
