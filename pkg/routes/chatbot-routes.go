package routes

import (
	"github.com/MirInjamamul/go-chatbot/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterChatUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/chat_user/", controllers.CreateChatUser).Methods("POST")
	router.HandleFunc("/chat_user/", controllers.GetChatUser).Methods("GET")
	router.HandleFunc("/chat_user/{userId}", controllers.GetChatUserById).Methods("GET")
	router.HandleFunc("/chat_user/{userId}", controllers.UpdateChatUser).Methods("PUT")
	router.HandleFunc("/chat_user/{userId}", controllers.DeleteChatUser).Methods("DELETE")

}
