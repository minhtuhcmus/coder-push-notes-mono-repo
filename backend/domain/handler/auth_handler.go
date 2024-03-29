package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/controllers"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/domain/models"
	"github.com/minhtuhcmus/coder-push-notes-mono-repo/backend/utils"
	"net/http"
	"time"
)

func NewAuthHandler(authController controllers.AuthController) func(r chi.Router) {

	ah := &authHandler{authController: authController}

	return func(r chi.Router) {
		r.Post("/signin", ah.SignIn)
		r.Post("/signup", ah.SignUp)
		r.Post("/verify-token", ah.VerifyToken)
	}
}

type authHandler struct {
	authController controllers.AuthController
}

type AuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	VerifyToken(w http.ResponseWriter, r *http.Request)
}

func (ah *authHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	t := time.Now().UnixNano()

	var creds models.AuthRequest
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	accessToken, err := ah.authController.SignIn(r.Context(), &creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Domain: "http://localhost:8080",
		Name:   "access_token",
		Value:  accessToken,
		Secure: true,
		MaxAge: 300,
	})

	fmt.Println("SignIn Handler ", time.Now().UnixNano()-t)

}

func (ah *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {
}

func (ah *authHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("access_token")
	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}
	if authCookie.Value == "" {
		json.NewEncoder(w).Encode(false)
		return
	}

	authClaims, err := utils.JwtValidate(authCookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	if authClaims == nil || authClaims.Permissions == nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	json.NewEncoder(w).Encode(true)
}
