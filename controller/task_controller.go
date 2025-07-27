package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todo-list/data/request"
	"todo-list/data/response"
	"todo-list/ent"
	"todo-list/helper"
	"todo-list/service"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *TaskController {
	return &TaskController{
		TaskService: taskService,
	}
}

func (tc *TaskController) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID := r.Context().Value("id").(uuid.UUID)

	if userID == uuid.Nil {
		helper.RespondJSON(w, http.StatusUnauthorized, "Unauthorized")
	}

	var req request.CreateTask
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	_, err := tc.TaskService.Create(r.Context(), userID, req)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondJSON(w, http.StatusOK, "task created")
}

func (tc *TaskController) GetAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID := r.Context().Value("id").(uuid.UUID)
	if userID == uuid.Nil {
		helper.RespondJSON(w, http.StatusUnauthorized, response.ErrorResponse{Code: 401, Message: "Unauthorized"})
		return
	}

	tasks, err := tc.TaskService.GetAllTask(r.Context(), userID)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}

	var taskResponses []response.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, convertToTaskResponse(task))
	}

	helper.RespondJSON(w, http.StatusOK, response.TaskListResponse{
		Success: true,
		Data:    taskResponses,
	})
}

func (tc *TaskController) GetByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID := r.Context().Value("id").(uuid.UUID)
	if userID == uuid.Nil {
		helper.RespondJSON(w, http.StatusUnauthorized, response.ErrorResponse{Code: 401, Message: "Unauthorized"})
		return
	}

	paramID := param.ByName("task_id")
	taskID, err := uuid.Parse(paramID)
	if err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, response.ErrorResponse{Code: 400, Message: "Invalid task ID"})
		return
	}

	task, err := tc.TaskService.GetTaskByID(r.Context(), taskID, userID)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}

	helper.RespondJSON(w, http.StatusOK, response.TaskDetailResponse{
		Success: true,
		Data:    convertToTaskResponse(task),
	})
}

func (tc *TaskController) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID := r.Context().Value("id").(uuid.UUID)
	paramID := param.ByName("task_id")
	taskID, err := uuid.Parse(paramID)
	if err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, "invalid task id")
	}
	if userID == uuid.Nil {
		helper.RespondJSON(w, http.StatusUnauthorized, "Unauthorized")
	}
	var req request.UpdateTask
	_, err = tc.TaskService.Update(r.Context(), taskID, userID, req)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondJSON(w, http.StatusOK, response.TaskUpdateResponse{
		Success: true,
		Message: "task updated",
	})
}

func (tc *TaskController) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID := r.Context().Value("id").(uuid.UUID)
	if userID == uuid.Nil {
		helper.RespondJSON(w, http.StatusUnauthorized, response.ErrorResponse{Code: 401, Message: "Unauthorized"})
		return
	}

	paramID := param.ByName("task_id")
	taskID, err := uuid.Parse(paramID)
	if err != nil {
		helper.RespondJSON(w, http.StatusBadRequest, response.ErrorResponse{Code: 400, Message: "Invalid task ID"})
		return
	}

	err = tc.TaskService.Delete(r.Context(), taskID, userID)
	if err != nil {
		helper.RespondJSON(w, http.StatusInternalServerError, response.ErrorResponse{Code: 500, Message: err.Error()})
		return
	}

	helper.RespondJSON(w, http.StatusOK, response.TaskDeleteResponse{
		Success: true,
		Message: "Task deleted successfully",
	})
}

func convertToTaskResponse(task *ent.Task) response.TaskResponse {
	return response.TaskResponse{
		ID:         task.ID,
		Title:      task.Title,
		Note:       task.Note,
		IsComplete: task.Complete,
		DueDate:    &task.DueDate,
	}
}
