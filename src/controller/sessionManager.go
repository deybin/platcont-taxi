package controller

import "taxi-platcont-go/src/helper"

var SessionMgr *helper.SessionMgr = nil
var SessionID string

/*
	SesssionMgr={
			"en":"", id_ente => id de empresa registrado en la base de datos platcont tabla clientes
			"og":"", id_orga =>id de organizacion registrado en la base en la tabla requ_organization
			"us":"", id_user => id del usurario que se conecto a la api
			"ot":"", database_name => nombre de la base de datos del cliente
			"month":"", month => mes actual en formato MM(01,02,03,04,05,06,07,08,09,10,11,12)
			"year":"", year => año actual en formato YYYY(2020,2021,2022,2023,2024,2025,2026,2027,2028,2029,2030)
			"date":"", date=> fecha actual en formato DD/MM/YYYY(02/032019)
			"cargo":"", cargo => cargo del usuario que se conecto a la api
			"branch":"", c_bran => código de la sucursal del usuario que se conecto a la api
			"store_house":"" c_stor=>código del almacén del usuario que se conecto a la api
	}
*/
