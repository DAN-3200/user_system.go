// recebe a requisição já tratada
package usecase

import (
	"app/internal/dto"
	"app/internal/model"
	"app/pkg/security"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (it *LayerUseCase) CreateUser(info dto.UserReq) error {
	hash, err := security.HashPassword(info.Password)
	if err != nil {
		return fmt.Errorf("Error Bycript HashPassword")
	}

	newUser := model.User{
		ID:              uuid.New().String(),
		Name:            info.Name,
		Email:           info.Email,
		PasswordHash:    hash,
		IsEmailVerified: false,
		IsActive:        true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Role:            info.Role,
	}

	err = it.Repo.CreateUserSQL(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerUseCase) GetUser(infoID string) (dto.UserRes, error) {
	result, err := it.Repo.GetUserSQL(infoID)
	if err != nil {
		return dto.UserRes{}, err
	}

	return result, nil
}

func (it *LayerUseCase) GetUserList() ([]dto.UserRes, error) {
	result, err := it.Repo.GetUserListSQL()
	if err != nil {
		return []dto.UserRes{}, err
	}

	return result, nil
}

func (it *LayerUseCase) EditUser(id string, info dto.EditUserReq) error {
	if info.Password != nil {
		hash, err := security.HashPassword(*info.Password)
		if err != nil {
			return fmt.Errorf("Error Bycript HashPassword")
		}
		info.Password = &hash
	}

	err := it.Repo.EditUserSQL(id, info)

	if err != nil {
		return err
	}
	return nil
}

func (it *LayerUseCase) DeleteUser(infoID string) error {
	var err = it.Repo.DeleteUserSQL(infoID)
	if err != nil {
		return err
	}
	return nil
}
