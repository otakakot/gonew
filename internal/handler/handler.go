package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/otakakot/gonew/internal/cache"
	"github.com/otakakot/gonew/internal/database"
	"github.com/otakakot/gonew/internal/env"
)

type Handler struct {
	db *database.Database
	cc *cache.Cache
}

func New(
	db *database.Database,
	cc *cache.Cache,
) *Handler {
	return &Handler{
		db: db,
		cc: cc,
	}
}

func (hdl *Handler) Database(w http.ResponseWriter, r *http.Request) {
	hdl.log(r)

	_, _ = w.Write([]byte("Hello Database"))
}

func (hdl *Handler) Cache(w http.ResponseWriter, r *http.Request) {
	hdl.log(r)

	_, _ = w.Write([]byte("Hello Cache"))
}

func (hdl *Handler) Env(w http.ResponseWriter, r *http.Request) {
	hdl.log(r)

	_, _ = w.Write([]byte(fmt.Sprintf("Hello %s", env.Get())))
}

func (hdl *Handler) log(r *http.Request) {
	slog.Info(fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.UserAgent()))
}
