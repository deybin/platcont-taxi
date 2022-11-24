package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"taxi-platcont-go/src/auth"
	"taxi-platcont-go/src/controller"
	"taxi-platcont-go/src/libraries/cryptoAes"
	"taxi-platcont-go/src/libraries/date"
	"taxi-platcont-go/src/libraries/library"

	"github.com/gorilla/mux"
)

func AuthLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := controller.NewResponseManager()
		token := r.Header.Get("access-token")
		if token == "" {
			response.Msg = "Intento de Violación de Seguridad, usuario no autorizado"
			response.Status = "error"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		data, err := auth.ValidateToken(token)
		if err != nil {
			response.Msg = "token invalido o expirado"
			response.Status = "Error Grave"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		array_modules := strings.Split(data.Modules, ",")
		if library.IndexOf_String(array_modules, "12") == -1 {
			response.Msg = "No cuenta con los permisos necesarios para acceder a este modulo"
			response.Status = "Error Grave"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		if library.CheckSession(data.Id_user) != true {
			controller.SessionID = controller.SessionMgr.StartSession(w, r)
			database, _ := cryptoAes.AesDecrypt_PHP([]byte(data.Id_other), auth.GetKey_PrivateCrypto())
			controller.SessionMgr.SetSessionVal(controller.SessionID, "en", data.Id_ente)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "pr", data.Id_prod)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "og", data.Id_orga)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "us", data.Id_user)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "ot", string(database))
			controller.SessionMgr.SetSessionVal(controller.SessionID, "mod", data.Modules)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "month", data.Month)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "year", data.Year)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "date", data.Date)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "cargo", data.Cargo)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "branch", data.Branch)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "store_house", data.StoreHouse)
		}
		next.ServeHTTP(w, r)

	})
}

func AuthLogin_Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := controller.NewResponseManager()
		token := r.Header.Get("access-token")
		if token == "" {
			response.Msg = "Intento de Violación de Seguridad, usuario no autorizado"
			response.Status = "error"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}
		data, err := auth.ValidateToken(token)
		if err != nil {
			response.Msg = "token invalido o expirado"
			response.Status = "Error Grave"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		if data.Cargo > 2 {
			response.Msg = "Usuario no autorizado"
			response.Status = "Error de autorización"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		array_modules := strings.Split(data.Modules, ",")
		if library.IndexOf_String(array_modules, "12") == -1 {
			response.Msg = "No cuenta con los permisos necesarios para acceder a este modulo"
			response.Status = "Error Grave"
			response.StatusCode = 401
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		if library.CheckSession(data.Id_user) != true {
			controller.SessionID = controller.SessionMgr.StartSession(w, r)
			database, _ := cryptoAes.AesDecrypt_PHP([]byte(data.Id_other), auth.GetKey_PrivateCrypto())
			controller.SessionMgr.SetSessionVal(controller.SessionID, "en", data.Id_ente)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "pr", data.Id_prod)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "og", data.Id_orga)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "us", data.Id_user)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "ot", string(database))
			controller.SessionMgr.SetSessionVal(controller.SessionID, "mod", data.Modules)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "month", data.Month)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "year", data.Year)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "date", data.Date)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "cargo", data.Cargo)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "branch", data.Branch)
			controller.SessionMgr.SetSessionVal(controller.SessionID, "store_house", data.StoreHouse)
		}
		next.ServeHTTP(w, r)

	})
}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Auth-Date, Auth-Periodo, Access-Token")
			next.ServeHTTP(w, req)
		})
}

func AuthenticationDate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := controller.NewResponseManager()
		Auth_Date := r.Header.Get("Auth-Date")
		if Auth_Date == "" {
			response.Msg = "Intento de Violación de Seguridad, no se encontró la cabecera Auth-KeyDate-Date"
			response.Status = "error"
			response.StatusCode = 409
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(response)
			return
		}
		date_server := date.GetFechaLocationString()

		if Auth_Date != date_server {
			response.Msg = "Para realizar cualquier acción la fecha de autenticación debe ser igual a la fecha del servidor"
			response.Status = "error"
			response.StatusCode = 409
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(response)
			return
		} else {
			Auth_Periodo := r.Header.Get("Auth-Periodo")

			if Auth_Periodo == "" {
				response.Msg = "Intento de Violación de Seguridad, no se encontró la cabecera Auth-KeyPeriodo-Periodo"
				response.Status = "error"
				response.StatusCode = 409
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(response)
				return
			}

			split_periodo := strings.Split(Auth_Periodo, "-")
			var year_client string
			var month_client string
			if len(split_periodo) == 2 {
				year_client = split_periodo[0]
				month_client = split_periodo[1]
			}
			year_server := date.GetYearString()
			month_server := date.GetMonthString()
			if year_server != year_client && month_server != month_client {
				response.Msg = "Para realizar cualquier acción el periodo de autenticación debe ser igual al periodo del servidor"
				response.Status = "error"
				response.StatusCode = 409
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(response)
				return
			} else {
				next.ServeHTTP(w, r)
			}
		}

	})
}

func EnableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(MiddlewareCors)
}
