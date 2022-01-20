package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/congruity7/moonshot-go/pkg/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (c *Context) GetConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var config models.Config

	c.ds.Db.Table("config").First(&config)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(config)
	w.Write(bytes)
}

func (c *Context) CreateConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var config models.Config

	err := json.NewDecoder(r.Body).Decode(&config)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logrus.Error("creating config", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := c.ds.Db.Table("config").First(&config)

	if result.Error != gorm.ErrRecordNotFound {
		logrus.Error("creating config", errors.New("record already exists"))
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	c.ds.Db.Table("config").Create(&config)
	bytes, _ := json.Marshal(config)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}

func (c *Context) UpdateConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var requestConfig models.Config
	var config models.Config

	err := json.NewDecoder(r.Body).Decode(&requestConfig)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logrus.Error("updating config", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := c.ds.Db.Table("config").First(&config)
	if result.Error == gorm.ErrRecordNotFound {
		logrus.Error("updating config", result.Error)
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	c.ds.Db.Table("config").Save(&requestConfig)
	bytes, _ := json.Marshal(requestConfig)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return

}
