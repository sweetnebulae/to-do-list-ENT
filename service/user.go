package service

import (
	"context"
	"errors"
	"time"
	"todo-list/data/request"
	"todo-list/ent"
	"todo-list/ent/user"
	"todo-list/utils"
)

type UserService struct {
	Client       *ent.Client
	CacheService *utils.CacheService
	secretKey    string
}

func NewUserService(client *ent.Client, secretKey string, cacheService *utils.CacheService) *UserService {
	return &UserService{
		client,
		cacheService,
		secretKey,
	}
}
func (s *UserService) Register(ctx context.Context, req request.RegisterUser) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	_, err = s.Client.User.
		Create().
		SetUsername(req.Username).
		SetName(req.Name).
		SetPassword(hashedPassword).
		Save(ctx)
	return err
}

func (s *UserService) Login(ctx context.Context, req request.LoginUser) (string, error) {
	u, err := s.Client.User.
		Query().
		Where(user.UsernameEQ(req.Username)).
		Only(ctx)

	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if err := utils.CheckPasswordHash(req.Password, u.Password); err != nil {
		return "", errors.New("invalid password")
	}
	token, err := utils.GenerateTokens(u.ID, u.Username, s.secretKey)
	if err != nil {
		return "", err
	}

	if err := s.CacheService.Set(u.ID.String(), token, 24*time.Hour); err != nil {
		return "", err
	}

	return token, nil
}

//func (s *UserService) UpdateUser(ctx context.Context, req request.UpdateUser) error {
//	_, err := s.Client.User.
//		Update().
//		SetName(req.Name).
//		SetProfilePict(req.ProfilePict).
//		Save(ctx)
//	if err != nil {
//		fmt.Println("failed to update", err)
//	}
//	return nil
//}

func (s *UserService) Logout(ctx context.Context, userID string) error {
	return s.CacheService.Delete(userID)
}
