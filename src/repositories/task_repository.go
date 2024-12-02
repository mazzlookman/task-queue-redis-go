package repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type TaskRepository struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

// constructor
func NewTaskRepository(client *redis.Client, ctx context.Context) *TaskRepository {
	return &TaskRepository{
		RedisClient: client,
		Ctx:         ctx,
	}
}

func (r *TaskRepository) AddTask(queueName, taskId, statusKey string) error {
	if err := r.RedisClient.LPush(r.Ctx, queueName, taskId).Err(); err != nil {
		return err
	}

	if err := r.RedisClient.Set(r.Ctx, statusKey, "pending", 0).Err(); err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) GetTaskStatus(statusKey string) (string, error) {
	return r.RedisClient.Get(r.Ctx, statusKey).Result()
}

func (r *TaskRepository) UpdateTaskStatus(statusKey, status string) error {
	return r.RedisClient.Set(r.Ctx, statusKey, status, 0).Err()
}

func (r *TaskRepository) PopTask(queueName string) (string, error) {
	task, err := r.RedisClient.BRPop(r.Ctx, 0, queueName).Result()
	if err != nil {
		return "", err
	}

	return task[1], nil // BRPop returns a slice [queueName, value]
}
