package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"todo-list/data/request"
	"todo-list/ent"
	"todo-list/ent/task"
	"todo-list/ent/user"
)

type TaskService struct {
	Client *ent.Client
}

func NewTaskService(client *ent.Client) *TaskService {
	return &TaskService{
		Client: client,
	}
}

func (t *TaskService) Create(ctx context.Context, userID uuid.UUID, request request.CreateTask) (*ent.Task, error) {
	return t.Client.Task.
		Create().
		SetTitle(request.Title).
		SetNote(request.Note).
		SetDueDate(request.Due).
		SetOwnerID(userID).
		Save(ctx)
}

func (t *TaskService) GetAllTask(ctx context.Context, userID uuid.UUID) ([]*ent.Task, error) {
	return t.Client.Task.
		Query().
		Where(task.HasOwnerWith(user.ID(userID))).
		All(ctx)
}

func (t *TaskService) GetTaskByID(ctx context.Context, taskID uuid.UUID, userID uuid.UUID) (*ent.Task, error) {
	return t.Client.Task.
		Query().
		Where(
			task.ID(taskID),
			task.HasOwnerWith(user.ID(userID)),
		).
		Only(ctx)
}

func (t *TaskService) Update(ctx context.Context, taskID uuid.UUID, userID uuid.UUID, request request.UpdateTask) (*ent.Task, error) {
	updateQuery := t.Client.Task.
		Update().
		Where(
			task.ID(taskID),
			task.HasOwnerWith(user.ID(userID)),
		)

	// Conditional updates
	if request.Title != nil {
		updateQuery.SetTitle(*request.Title)
	}
	if request.Note != nil {
		updateQuery.SetNote(*request.Note)
	}
	if request.Complete != nil {
		updateQuery.SetComplete(*request.Complete)
	}
	if request.Due != nil {
		updateQuery.SetDueDate(*request.Due)
	}

	// Execute dan return
	if _, err := updateQuery.Save(ctx); err != nil {
		return nil, err
	}

	return t.GetTaskByID(ctx, taskID, userID)
}

func (t *TaskService) Delete(ctx context.Context, taskID uuid.UUID, userID uuid.UUID) error {
	exists, err := t.Client.Task.
		Query().
		Where(task.ID(
			taskID),
			task.HasOwnerWith(user.ID(userID))).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("task not found or access denied")
	}
	_, err = t.Client.Task.
		Delete().
		Where(
			task.ID(taskID),
			task.HasOwnerWith(user.ID(userID))).
		Exec(ctx)
	return err
}
