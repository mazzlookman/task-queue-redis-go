package workers

import (
	"log"
	"redis-task-queue/src/repositories"
	"time"
)

func StartWorker(taskRepo *repositories.TaskRepository, queueName string) {
	for {
		taskId, err := taskRepo.PopTask(queueName)
		if err != nil {
			log.Println("Error popping task: ", err)
			continue
		}

		statusKey := "task:" + taskId + ":status"
		taskRepo.UpdateTaskStatus(statusKey, "in-progress")

		// simulate task processing
		time.Sleep(5 * time.Second)

		taskRepo.UpdateTaskStatus(statusKey, "done")
		log.Println("Task", taskId, "completed")
	}
}
