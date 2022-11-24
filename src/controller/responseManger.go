package controller

type ResponseManager struct {
	Msg        string                 "json:\"msg\""
	StatusCode int                    "json:\"statusCode\""
	Status     string                 "json:\"status\""
	Data       map[string]interface{} "json:\"data\""
}

func NewResponseManager() *ResponseManager {
	return &ResponseManager{
		Msg:        "successfully",
		StatusCode: 200,
		Status:     "success",
		Data:       make(map[string]interface{}),
	}
}

/*
	StatusCode:
		100: success info
		200: Success
		201: Success-Created
		300: Error-Multiple
		400: Error-BadRequest
		401: Unauthorized->No autorizado->perdida de session o no tiene permisos
		404: NotFound -> No Encontrado->usuario no encontrado
		409: Conflict -> Error de conflicto->fecha del cliente no es igual a la del servidor
		410: Gone -> Error de borrado->periodo que tiene el cliente es diferente seleccionado en el servido

		500: Error

*/
