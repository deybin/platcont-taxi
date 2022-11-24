package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"taxi-platcont-go/src/controller"
	"taxi-platcont-go/src/database/models/tables"
	"taxi-platcont-go/src/database/orm"
	"taxi-platcont-go/src/middleware"

	"github.com/gorilla/mux"
)

func RutasCliente(r *mux.Router) {

	s := r.PathPrefix("/clientes").Subrouter()
	s.Handle("/list", middleware.AuthLogin(http.HandlerFunc(allCliente))).Methods("GET")
	s.Handle("/info/{n_docu}", middleware.AuthLogin(http.HandlerFunc(oneCLiente))).Methods("GET")
	s.Handle("/update/{n_docu}", middleware.AuthLogin(http.HandlerFunc(updateCliente))).Methods("PUT")
	s.Handle("/create", middleware.AuthLogin(http.HandlerFunc(insertCliente))).Methods("POST")
}

func allCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_cliente := orm.NewQuerys("requ_clientes").Select("c_docu,n_docu,l_clie").OrderBy("n_docu").Exec().All()

	if len(data_cliente) <= 0 {
		controller.ErrorsSuccess(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["clientes"] = data_cliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Clientes_GetSchema()
	_Clientes := orm.SqlExec{}
	err = _Clientes.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Clientes.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Clientes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	n_docu := params["n_docu"]
	if n_docu == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"n_docu": n_docu}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Clientes_GetSchema()
	_Clientes := orm.SqlExec{}
	err = _Clientes.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Clientes.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Clientes.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func oneCLiente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	n_docu := params["n_docu"]
	if n_docu == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	//get allData from database
	data_cliente := orm.NewQuerys("requ_clientes").Select("l_emai,c_docu,n_docu,l_clie,k_gene,f_naci,l_dire,l_refe,c_ubig,n_tele,n_celu,l_obse").Where("n_docu", "=", n_docu).Exec().One()

	if len(data_cliente) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data = data_cliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
