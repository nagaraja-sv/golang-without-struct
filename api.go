package forcego

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router := httprouter.New()

	router.POST("/api/organization", GetAccountHandler)
	//router.POST("/api/uploadimage")

	http.Handle("/", router)
	http.HandleFunc("/api/uploadimage", handler)
}
