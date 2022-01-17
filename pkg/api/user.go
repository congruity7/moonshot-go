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

	queryValues := r.URL.Query()

	if wa := queryValues.Get("wallet_address"); wa != "" {
		var user models.User
		var wallet models.Wallet

		c.ds.Db.Table("wallet").First(&wallet, "wallet_address = ?", wa)

		logrus.Info("wallet", wallet)

		c.ds.Db.Table("user").Preload("wallet").First(&user, "id", wallet.UserID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		bytes, _ := json.Marshal(user)
		w.Write(bytes)
		return
	}

	var users []models.User

	c.ds.Db.Table("user").Find(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(users)
	w.Write(bytes)
}

func (c *Context) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		logrus.Error("creating user", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := c.ds.Db.Table("user").Create(&user)

	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		logrus.Error("creating user", result.Error)
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(user)
	w.Write(bytes)
}

func (c *Context) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		logrus.Error("updating user", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := c.ds.Db.Table("user").Save(user)

	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		logrus.Error("updating user", result.Error)
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(user)
	w.Write(bytes)
}

func (c *Context) DeleteUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
