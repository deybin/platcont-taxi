package routes

import (
	"errors"
	"fmt"
	"io/ioutil"
	"taxi-platcont-go/src/controller"
	"taxi-platcont-go/src/database/models/tables"
	"taxi-platcont-go/src/database/orm"
	"taxi-platcont-go/src/libraries/date"
	"taxi-platcont-go/src/libraries/library"
	"taxi-platcont-go/src/middleware"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RutasServicios(r *mux.Router) {
	s := r.PathPrefix("/servicios").Subrouter()
	s.Handle("/list", middleware.AuthLogin(http.HandlerFunc(listarAllservicios))).Methods("GET")
	s.Handle("/info/{id_serv}", middleware.AuthLogin(http.HandlerFunc(listOneServicio))).Methods("GET")
	s.Handle("/info-cliente/{n_docu}", middleware.AuthLogin(http.HandlerFunc(servicioCliente))).Methods("GET")
	s.Handle("/list-factura/{id_serv}", middleware.AuthLogin(http.HandlerFunc(detalleFactura))).Methods("GET")
	s.Handle("/create", middleware.AuthLogin(http.HandlerFunc(insertServicios))).Methods("POST")
	s.Handle("/create-factura/{id_serv}", middleware.AuthLogin(http.HandlerFunc(regDetalleFactura))).Methods("POST")
	s.Handle("/update/{id_serv}", middleware.AuthLogin(http.HandlerFunc(updateServicio))).Methods("PUT")
	s.Handle("/update-baja/{id_serv}", middleware.AuthLogin(http.HandlerFunc(darBaja))).Methods("PUT")
	s.Handle("/update-alta/{id_serv}", middleware.AuthLogin(http.HandlerFunc(darAlta))).Methods("PUT")

}

func listarAllservicios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_servicio := orm.NewQuerys("taxi_servicios").Select("c_plac,f_fact,id_serv,k_stad,n_flot,n_month,n_year").OrderBy("n_flot").Exec().All()

	if len(data_servicio) <= 0 {
		controller.ErrorsSuccess(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data["servicios"] = data_servicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func listOneServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_serv := params["id_serv"]
	if id_serv == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	data_servicio := orm.NewQuerys("taxi_servicios as a").Select("a.*,b.l_clie").InnerJoin("requ_clientes as b", "a.n_docu = b.n_docu").Where("id_serv", "=", id_serv).Exec().One()

	if len(data_servicio) <= 0 {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	response.Data = data_servicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertServicios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Servicios_GetSchema()
	_Servicios := orm.SqlExec{}
	err = _Servicios.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Servicios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Servicios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func servicioCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	// id_serv := params["id_serv"]
	n_docu := params["n_docu"]
	if n_docu == "" {
		controller.ErrorsSuccess(w, errors.New("cliente no cuenta con ningun servicio"))
	}

	data_servicio := orm.NewQuerys("taxi_servicios").Select("n_year,n_month,f_fact,s_impo,c_plac,k_stad,f_digi,id_serv,n_flot").Where("n_docu", "=", n_docu).Exec().All()
	if len(data_servicio) <= 0 {
		controller.ErrorsInfo(w, errors.New("cliente no cuenta con ningun servicio"))
		return
	}
	response.Data["servicio_cliente"] = data_servicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func detalleFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_serv := params["id_serv"]

	if id_serv == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	fact := orm.NewQuerys("taxi_serviciosdetalle").Select("n_year || '-' || n_month as periodo").Where("id_serv", "=", id_serv).And("k_stad", "=", 0).Exec().All()

	var newFact []string
	for _, v := range fact {
		newFact = append(newFact, v["periodo"].(string))
	}
	// fmt.Println(newFact)

	data_servicio := orm.NewQuerys("taxi_servicios").Select("n_year", "n_month", "f_fact", "s_impo", "k_stad").Where("id_serv", "=", id_serv).Exec().One()
	date_fact := date.GetDate(data_servicio["f_fact"].(string))
	date_now := date.GetDateLocation()

	month_init := int64(date_fact.Month())
	year_init := date_fact.Year()
	month_now := int64(date_now.Month())
	year_now := date_now.Year()

	var data_facturaciones []map[string]interface{}
	var month = int64(12)

	impo := data_servicio["s_impo"].(float64)

	for i := year_init; i <= year_now; i++ {
		if i == year_now {
			month = month_now
		}
		for e := month_init; e <= month; e++ {
			// fmt.Println(i, e)
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%d", e)
			if library.IndexOf_String(newFact, year+"-"+month) == -1 {
				data_facturaciones = append(data_facturaciones, map[string]interface{}{
					"n_year":  year,
					"n_month": month,
					"s_impo":  impo,
				})
			}
		}

		month_init = 1
	}

	response.Data["detalle_fact"] = data_facturaciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func regDetalleFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id_serv := params["id_serv"]
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}

	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_request)

	schema, table := tables.Serviciosdetalle_GetSchema(id_serv)
	_Servicios := orm.SqlExec{}
	err = _Servicios.New(data_insert, table).Insert(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Servicios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Servicios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id_serv := params["id_serv"]
	if id_serv == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}
	data_request, err := controller.CheckBody(w, r)
	if err != nil {
		return
	}
	data_request["where"] = map[string]interface{}{"id_serv": id_serv}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_request)

	schema, table := tables.Servicios_GetSchema()
	_Servicios := orm.SqlExec{}
	err = _Servicios.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Servicios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Servicios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func darBaja(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)

	id_serv := params["id_serv"]
	if id_serv == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para dar de baja el servicio"))
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		controller.ErrorServer(w, err)
		return
	}

	data_request := make(map[string]interface{})
	json.Unmarshal(reqBody, &data_request)

	cargo := library.GetSession_key_interface("cargo").(int64)

	if cargo <= 2 {

		controller.ErrorsWaning(w, errors.New("No se pueden hacer cambios"))
		return
	}

	data_pagos := orm.NewQuerys("taxi_serviciosdetalle").Select().Where("id_serv", "=", id_serv).And("k_stad", "=", 1).Exec().All()

	if len(data_pagos) <= 0 {
		controller.ErrorsInfo(w, errors.New("No se encontraron resultados para la consulta"))
		return
	}

	data_body := make(map[string]interface{})
	data_body["k_stad"] = int64(1)
	// data_body["s_impo"] = int64(0)

	data_body["where"] = map[string]interface{}{"id_serv": id_serv}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Servicios_GetSchema()
	_Servicios := orm.SqlExec{}
	err = _Servicios.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Servicios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Servicios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func darAlta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id_serv := params["id_serv"]
	if id_serv == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para dar de alta el servicio"))
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		controller.ErrorServer(w, err)
		return
	}

	data_request := make(map[string]interface{})
	json.Unmarshal(reqBody, &data_request)

	data_body := make(map[string]interface{})
	if data_request["s_impo"] != nil {
		data_body["s_impo"] = data_request["s_impo"]
	}

	data_body["k_stad"] = int64(0)
	// retorna fecha de formato string dd/mm/yyyy (America Bogota)
	data_body["f_fact"] = date.GetFechaLocationString()

	data_body["where"] = map[string]interface{}{"id_serv": id_serv}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Servicios_GetSchema()
	_Servicios := orm.SqlExec{}
	err = _Servicios.New(data_update, table).Update(schema)
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	err = _Servicios.Exec()
	if err != nil {
		controller.ErrorsWaning(w, err)
		return
	}

	response.Data = _Servicios.Data[0]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func pagarTributo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_serv := params["id_serv"]

	if id_serv == "" {
		controller.ErrorsWaning(w, errors.New("no se encontraron resultados para la consulta"))
		return
	}

	fact := orm.NewQuerys("taxi_serviciosdetalle").Select("n_year || '-' || n_month as periodo").Where("id_serv", "=", id_serv).And("k_stad", "=", 0).Exec().All()

	var newFact []string
	for _, v := range fact {
		newFact = append(newFact, v["periodo"].(string))
	}
	// fmt.Println(newFact)

	data_servicio := orm.NewQuerys("taxi_servicios").Select("n_year", "n_month", "f_fact", "s_impo", "k_stad").Where("id_serv", "=", id_serv).Exec().One()
	date_fact := date.GetDate(data_servicio["f_fact"].(string))
	date_now := date.GetDateLocation()

	month_init := int64(date_fact.Month())
	year_init := date_fact.Year()
	month_now := int64(date_now.Month())
	year_now := date_now.Year()

	var data_facturaciones []map[string]interface{}
	var month = int64(12)

	impo := data_servicio["s_impo"].(float64)
	
	if pagoTributo == true {
		return

	} else {
		return false

	}

	for i := year_init; i <= year_now; i++ {
		if i == year_now {
			month = month_now
		}
		for e := month_init; e <= month; e++ {
			// fmt.Println(i, e)
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%d", e)
			if library.IndexOf_String(newFact, year+"-"+month) == -1 {
				data_facturaciones = append(data_facturaciones, map[string]interface{}{
					"n_year":  year,
					"n_month": month,
					"s_impo":  impo,
				})
			}
		}

		month_init = 1
	}

	response.Data["detalle_fact"] = data_facturaciones
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
