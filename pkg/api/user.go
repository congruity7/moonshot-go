package api

import (
	"context"
	"net/http"
	"time"

	"github.com/congruity7/moonshot-go/pkg/models"
	"github.com/julienschmidt/httprouter"
)

func (c *Context) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	var user models.User
	var id int

	const selectSQL = `SELECT * FROM users WHERE id = $1;`

	err := c.ds.Db.QueryRowContext(ctx, selectSQL, id).Scan(&user)

	if err != nil {
		c._log.Error("fetching user", "id", id, "err", err)
	}

	c._log.Info("fetched user details", "user", user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *Context) GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (c *Context) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (c *Context) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (c *Context) DeleteUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
