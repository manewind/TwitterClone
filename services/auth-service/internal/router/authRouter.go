package router

import (
	"auth-services/internal/controller"
	"auth-services/internal/service"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func AuthRouter(pool *pgxpool.Pool, logger *zap.Logger) *mux.Router {
	r := mux.NewRouter()
	userService := service.NewUserService(pool, logger)
	userController := controller.NewUserController(userService)
	r.HandleFunc("/create", userController.MainHandler)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "aboba")
	})
	return r
}
