package routes

import (
	"net/http"

	productViewHandler "github.com/Lechufane/go-next-crud/pkg/useCases/handlers/consumableApiHandler"
	responseHelper "github.com/Lechufane/go-next-crud/pkg/useCases/helpers/response"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type PlayerRouter struct {
}


// Get single product... Gets a single product
// @Description Get single product
// @Summary Endpoint to get a single product
// @Tags product
// @Param playerId path string true "playerId"
// @Success 201 {object} response.Response
// @Failure 400,401,409 {object} response.Response
// @Router /player/{playerId} [get]
func (pr *PlayerRouter) GetPlayerById(w http.ResponseWriter, r *http.Request) {
	playerId := chi.URLParam(r, "playerId")
	customers, status := productViewHandler.GetPlayerById(playerId)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), customers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	w.WriteHeader(status.StatusCode())
	responseHelper.ResponseStatusChecker(w, resp)
}

func (pr *PlayerRouter) Routes() http.Handler {
	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"https://*", "http://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowOriginFunc:    func(r *http.Request, origin string) bool { return true },
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              true,
		MaxAge:             300,
	}))

	r.Get("/{productId}", pr.GetPlayerById)

	return r
}
