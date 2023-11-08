package routes

import (
	"backend/api/handlers"
	"net/http"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	// allowedOrigins := []string{"http://facturacion.lumonidy.studio", "http://localhost:3000"}

	// c := middleware.CorsMiddleware(allowedOrigins)
	// r.Use(c)

	r.Handle("/api/facturacion", http.HandlerFunc(handlers.HomeHandler))

	r.Handle("/api/facturacion/user_paquetes", http.HandlerFunc(handlers.ObtenerPaquetesByUser))
	// Agrega más configuraciones de rutas aquí si es necesario
	//rutas google login
	r.Handle("/login-google", http.HandlerFunc(handlers.LoginGoogle))
	r.Handle("/callback-google", http.HandlerFunc(handlers.CallbackGoogle))
	//rutas facebook 

	r.Handle("/login-facebook" http.HandlerFunc(handlers.LoginFacebook))
	r.Handle("/callback-facebook", http.HandlerFunc(handlers.CallbackFacebook))

	r.Handle("/user/register", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterUser(w, r, app)
	})).Methods("POST")

	r.Handle("/user/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginUser(w, r, app)
	})).Methods("POST")

	r.Handle("/user/reset-password", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.ResetPassword(w, r)
	})).Methods("POST")

}

