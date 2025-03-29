package controller

import (
	"net/http"
	"time"

	"github.com/loveuer/uzone/pkg/uapi"
)

func kvCreate(c uapi.Context) error {
	type Req struct {
		Key string `json:"key"`
		Val string `json:"val"`
	}

	var (
		err error
		req = new(Req)
	)

	if err = c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if req.Key == "" || req.Val == "" {
		return c.Status(http.StatusBadRequest).SendString("invalid kv")
	}

	if err = c.UseCache().SetEx(c.Context(), req.Key, req.Val, 6*time.Hour); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	c.UseLogger().Info("saved kv = %s, %s", req.Key, req.Val)

	return c.JSON(req)
}

func kvGet(c uapi.Context) error {
	type Req struct {
		Key string `json:"key" query:"key"`
		Val string `json:"val"`
	}

	var (
		err error
		req = new(Req)
	)

	if err = c.QueryParser(req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// key := c.Query("key")

	if req.Key == "" {
		return c.Status(http.StatusBadRequest).SendString("invalid key")
	}

	var val string
	if err = c.UseCache().GetScan(c.Context(), req.Key).Scan(&val); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	req.Val = val

	return c.JSON(req)
}
