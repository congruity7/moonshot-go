package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/congruity7/moonshot-go/pkg/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (c *Context) GetRoundByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var round models.Round
	id := ps.ByName("id")

	idValue, _ := strconv.ParseInt(id, 10, 64)
	c.ds.Db.Table("round").First(&round, "id = ?", idValue)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(round)
	w.Write(bytes)
}

func (c *Context) GetRounds(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var rounds []models.Round

	c.ds.Db.Table("round").Find(&rounds).Limit(50).Order("created_at desc")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(rounds)
	w.Write(bytes)
}

func (c *Context) CreateRound(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var round models.Round
	err := json.NewDecoder(r.Body).Decode(&round)

	if err != nil {
		logrus.Error("creating round", err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := c.ds.Db.Table("round").Create(&round)

	if result.Error != nil {
		logrus.Error("creating round", result.Error)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(round)
	w.Write(bytes)
}

func (c *Context) UpdateRound(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var round models.Round
	err := json.NewDecoder(r.Body).Decode(&round)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		logrus.Error("updating round", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := c.ds.Db.Table("round").Save(round)

	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		logrus.Error("updating round", result.Error)
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(round)
	w.Write(bytes)

}

func (c *Context) DeleteRoundByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
