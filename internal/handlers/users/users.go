package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/kevannnn/dineder-backend/internal/api"
	users "github.com/kevannnn/dineder-backend/internal/dataaccess"
	"github.com/kevannnn/dineder-backend/internal/database"
	"github.com/pkg/errors"
)

const (
	ListUsers = "users.HandleList"

	SuccessfulListUsersMessage = "Successfully listed users"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveUsers           = "Failed to retrieve users in %s"
	ErrEncodeView              = "Failed to retrieve users in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListUsers))
	}

	users, err := users.List(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}

	data, err := json.Marshal(users)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListUsers))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}

func PostUser(w http.ResponseWriter, req *http.Request) {

	var newUser models.User
	json.NewDecoder(req.Body).Decode(newUser)

	//err := dataaccess.CreateUser(database.DB, newUser)
	//response, _ := utils.HandlerFormatter(err, newUser, "PostUser", constants.SuccessfulPostMessage)
	response := newUser.id
	json.NewEncoder(w).Encode(response)
}