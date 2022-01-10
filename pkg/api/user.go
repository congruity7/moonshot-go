package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/congruity7/moonshot-go/pkg/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (c *Context) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	logrus.Info("in get user by id")
	var user models.User
	id := ps.ByName("user_id")

	idValue, _ := strconv.ParseInt(id, 10, 64)
	c.ds.Db.Table("user").First(&user, "id = ?", idValue)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(user)
	w.Write(bytes)
}

func (c *Context) GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info("in get user by id")
	var users []models.User

	c.ds.Db.Table("user").Find(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(users)
	w.Write(bytes)
}

func (c *Context) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (c *Context) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (c *Context) DeleteUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
