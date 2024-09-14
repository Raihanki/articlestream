package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Raihanki/articlestream/internal/entities"
	"github.com/Raihanki/articlestream/internal/helpers"
	"github.com/Raihanki/articlestream/internal/repositories"
)

type UserHandler struct {
	UserRepository repositories.UserRepository
}

func (u *UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	request := entities.CreateUser{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helpers.JsonResponse(w, 400, nil, "invalid request body")
		return
	}

	userId, err := u.UserRepository.CreateUser(request)
	if err != nil {
		log.Printf("error creating user ERR::%v", err)
		helpers.JsonResponse(w, 500, nil, "")
		return
	}

	user, err := u.UserRepository.GetUserById(userId)
	if err != nil {
		log.Printf("error get user by id:%d ERR::%v", userId, err)
		helpers.JsonResponse(w, 500, nil, "")
		return
	}

	helpers.JsonResponse(w, 201, user, http.StatusText(http.StatusCreated))
}

func (u *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("userId")
	userId, err := strconv.Atoi(param)
	if err != nil {
		helpers.JsonResponse(w, 404, nil, http.StatusText(http.StatusNotFound))
		return
	}

	user, err := u.UserRepository.GetUserById(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, nil, "user not found")
			return
		}
		log.Printf("error get user by id:%d ERR::%v", userId, err)
		helpers.JsonResponse(w, 500, nil, "")
		return
	}

	helpers.JsonResponse(w, 200, user, http.StatusText(http.StatusOK))
}
