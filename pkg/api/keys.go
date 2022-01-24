package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (c *Context) PingStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	// conn, err := c.rs.Pool.GetContext(ctx)
	// defer conn.Close()

	//i, err := redis.String(conn.Do("PING"))
	cmd := c.rs.Client.Ping()
	err := cmd.Err() //Set(ctx, "key", "value", 0).Err()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("pinging memory store", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(cmd)
	w.Write(bytes)
}

func (c *Context) GetKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	key := ps.ByName("key")

	//conn, err := c.rs.Pool.GetContext(ctx)
	//defer conn.Close()

	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	c._log.Error("getting connection", err)
	// 	return
	// }

	//i, err := redis.String(conn.Do("GET", key))
	result, err := c.rs.Client.Get(key).Result()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("getting key", key, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}

func (c *Context) CreateKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	key := ps.ByName("key")
	value := ps.ByName("value")
	// conn, err := c.rs.Pool.GetContext(ctx)
	// defer conn.Close()

	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	c._log.Error("getting connection", err)
	// 	return
	// }

	// i, err := redis.String(conn.Do("SET", key, value))
	result, err := c.rs.Client.Set(key, value, 36000*time.Hour).Result()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("creating key", key, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}

func (c *Context) DeleteKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	key := ps.ByName("key")
	// conn, err := c.rs.Pool.GetContext(ctx)
	// defer conn.Close()

	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	c._log.Error("getting connection", err)
	// 	return
	// }

	//i, err := redis.String(conn.Do("DEL", key))
	result, err := c.rs.Client.Del().Result()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		c._log.Error("delete key", key, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(result)
	w.Write(bytes)
}
