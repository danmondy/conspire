package routes
import (
	"github.com/gorilla/mux"
	"net/http"
)

//////////api router

func RegisterApiRoutes(r *mux.Router){
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/{object}", apiListHandler)
	//api.HandleFunc("/{object}/{id}", apiObjectHandler())
}


func apiListHandler(w http.ResponseWriter, r *http.Request){

}

