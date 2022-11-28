package routes

import (
	"errors"
	"taxi-platcont-go/src/controller"
	"taxi-platcont-go/src/database/models/tables"
	"taxi-platcont-go/src/database/orm"
	"taxi-platcont-go/src/middleware"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RutasClientesCars(r *mux.Router) {
	s := r.PathPrefix("/clientes-cars").Subrouter()
	s.Handle("/info/{c_plac}", middleware.AuthLogin(http.HandlerFunc(listOneClienteCar))).Methods("GET")
	s.Handle("/list-documento/{n_docu}", middleware.AuthLogin(http.HandlerFunc(listDocumento))).Methods("GET")
	s.Handle("/list", middleware.AuthLogin(http.HandlerFunc(listarAllClientesCars))).Methods("GET")
	s.Handle("/create/{n_docu}", middleware.AuthLogin(http.HandlerFunc(insertClienteCar))).Methods("POST")
	s.Handle("/update/{c_plac}", middleware.AuthLogin(http.HandlerFunc(updateClienteCar))).Methods("PUT")
	s.Handle("/asignar/{c_plac}/{n_docu}", middleware.AuthLogin(http.HandlerFunc(asignar))).Methods("PUT")
}

func listarAllClientesCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	data_clientes_cars := orm.NewQuerys("taxi_clientescars").Select("c_plac,l_marc,l_mode,l_color,n_docu").OrderBy("c_plac").Exec().All()

	if len(data_clientes_cars) <= 0 {
		controller.ErrorsSuccess(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["cars"] = data_clientes_cars
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func listOneClienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	c_plac := params["c_plac"]
	if c_plac == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	data_clientes_cars := orm.NewQuerys("taxi_clientescars").Select().Where("c_plac", "=", c_plac).Exec().One()

	if len(data_clientes_cars) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data = data_clientes_cars
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func listDocumento(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	n_docu := params["n_docu"]
	if n_docu == "" {
		controller.ErrorsWaning(w, errors.New("Cliente no cuenta con vehiculo asignado"))
		return
	}

	data_clientes_cars := orm.NewQuerys("taxi_clientescars").Select("c_plac,l_mode,l_marc").Where("n_docu", "=", n_docu).Exec().All()

	if len(data_clientes_cars) <= 0 {
		controller.ErrorsWaning(w, errors.New("Cliente no cuenta con vehiculo asignado"))
		return
	}
	response.Data["cars"] = data_clientes_cars
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertClienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	n_docu := params["n_docu"]
	if n_docu == "" {
		controller.ErrorsSuccess(w, errors.New("no se encontraron resultados para la consulta"))
	}

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Clientescars_GetSchema(n_docu)
	_ClientesCars := orm.SqlExec{}
	err = _ClientesCars.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _ClientesCars.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _ClientesCars.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateClienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	c_plac := params["c_plac"]
	if c_plac == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"c_plac": c_plac}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Clientescars_GetSchema("")
	_ClientesCars := orm.SqlExec{}
	err = _ClientesCars.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _ClientesCars.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _ClientesCars.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func asignar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	c_plac := params["c_plac"]
	n_docu := params["n_docu"]
	if c_plac == "" && n_docu == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	data_request := map[string]interface{}{"n_docu": n_docu}
	data_request["where"] = map[string]interface{}{"c_plac": c_plac}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Clientescars_GetSchema("")
	_ClientesCars := orm.SqlExec{}
	err := _ClientesCars.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _ClientesCars.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _ClientesCars.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
