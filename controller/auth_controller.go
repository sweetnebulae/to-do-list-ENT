package controller

import (
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"todo-list/data/request"
	"todo-list/helper"
	"todo-list/service"
	"todo-list/utils"
)

type UserController struct {
	UserService  service.UserService
	CacheService utils.CacheService
}

func NewUserController(userService service.UserService, cacheService utils.CacheService) *UserController {
	return &UserController{
		UserService:  userService,
		CacheService: cacheService,
	}
}

func (s *UserController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var req request.RegisterUser
	if err := helper.DecodeJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := s.UserService.Register(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.RespondJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})

}

func (s *UserController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var req request.LoginUser
	if err := helper.DecodeJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := s.UserService.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	helper.RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := r.Context().Value("id").(uuid.UUID)

	if err := c.UserService.Logout(r.Context(), id.String()); err != nil {
		http.Error(w, `{"error":"failed to logout"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message":"logout successful"}`))
	if err != nil {
		return
	}
}
