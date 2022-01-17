package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

func (c *Context) GetKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	key := ps.ByName("key")

	conn, err := c.rs.Pool.GetContext(ctx)
	defer conn.Close()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("getting connection", err)
		return
	}

	i, err := redis.String(conn.Do("GET", key))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("getting key", key, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(i)
	w.Write(bytes)
}

func (c *Context) CreateKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	key := ps.ByName("key")
	value := ps.ByName("value")
	conn, err := c.rs.Pool.GetContext(ctx)
	defer conn.Close()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("getting connection", err)
		return
	}

	i, err := redis.String(conn.Do("SET", key, value))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("creating key", key, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(i)
	w.Write(bytes)
}

func (c *Context) DeleteKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	key := ps.ByName("key")
	conn, err := c.rs.Pool.GetContext(ctx)
	defer conn.Close()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("getting connection", err)
		return
	}

	i, err := redis.String(conn.Do("DEL", key))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("delete key", key, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(i)
	w.Write(bytes)
}
