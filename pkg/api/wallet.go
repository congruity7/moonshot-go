package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/congruity7/moonshot-go/pkg/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (c *Context) GetWalletByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	logrus.Info("in get wallet by id")
	var wallet models.Wallet
	id := ps.ByName("wallet_id")

	idValue, _ := strconv.ParseInt(id, 10, 64)
	c.ds.Db.Table("wallet").Preload("user").First(&wallet, "id = ?", idValue)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(wallet)
	w.Write(bytes)
}

func (c *Context) GetWallets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	queryValues := r.URL.Query()

	if uid := queryValues.Get("user_id"); uid != "" {
		var user models.User
		var wallet models.Wallet

		c.ds.Db.Table("user").First(&wallet, "id = ?", uid)
		c.ds.Db.Table("wallet").First(&wallet, "id", user.ID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		bytes, _ := json.Marshal(wallet)
		w.Write(bytes)
		return
	}

	var wallets []models.Wallet

	c.ds.Db.Table("wallet").Preload("user").Find(&wallets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(wallets)
	w.Write(bytes)
}

func (c *Context) CreateWallet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var wallet models.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)

	if err != nil {
		logrus.Error("creating wallet", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := c.ds.Db.Table("wallet").Create(&wallet)

	if result.Error != nil {
		logrus.Error("creating wallet", result.Error)
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(wallet)
	w.Write(bytes)
}

func (c *Context) UpdateWallet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (c *Context) DeleteWalletByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
