package library

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
	"taxi-platcont-go/src/controller"

	"github.com/google/uuid"
)

func GetSession_key(key string) string {
	data, err := controller.SessionMgr.GetSessionVal(controller.SessionID, key)
	if !err {
		fmt.Println("Error session:", key, "=>", err)
		return ""
	}
	// data := dataSession.(map[string]interface{})
	if data == nil {
		return ""
	}
	l := ""
	if fmt.Sprintf("%T", data) == "int64" {
		n := data.(int64)
		l = fmt.Sprintf("%v", n)
	} else {
		l = data.(string)
	}
	return l
}

func GetSession_key_interface(key string) interface{} {
	data, err := controller.SessionMgr.GetSessionVal(controller.SessionID, key)
	if !err {
		fmt.Println("Error session:", key, "=>", err)
		return nil
	}
	if data == nil {
		return nil
	}

	return data
}

func CheckSession(user string) bool {
	_, err := controller.SessionMgr.GetSessionVal(controller.SessionID, "ot")
	if !err {
		controller.SessionMgr.EndSessionBy(controller.SessionID)
		return false
	}
	us, err := controller.SessionMgr.GetSessionVal(controller.SessionID, "us")
	if !err {
		controller.SessionMgr.EndSessionBy(controller.SessionID)
		return false
	}

	if us != user {
		controller.SessionMgr.EndSessionBy(controller.SessionID)
		return false
	}
	return true
}

func GetDestructing(parameter interface{}) []interface{} {

	s := reflect.ValueOf(parameter).Elem()

	numCols := s.NumField()
	fmt.Print(numCols)
	columns := make([]interface{}, numCols)

	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	return columns
}

// devuelve el indice del array donde se encuentra tipo del valor buscado
func ConcurrentArraySearch(array []reflect.Type, value reflect.Type) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

//recibe un valor interface que no se reconoce su tipo y devuelve un string
func InterfaceToString(params ...interface{}) string {
	typeValue := reflect.TypeOf(params[0]).String()
	value := params[0]
	valueReturn := ""
	if strings.Contains(typeValue, "string") {
		toSql := false
		if len(params) == 2 && reflect.TypeOf(params[1]).Kind() == reflect.Bool {
			toSql = params[1].(bool)
		}

		if toSql {
			valueReturn = fmt.Sprintf("'%s'", value)
		} else {
			valueReturn = fmt.Sprintf("%s", value)
		}
	} else if strings.Contains(typeValue, "int") {
		valueReturn = fmt.Sprintf("%d", value)
	} else if strings.Contains(typeValue, "float") {
		valueReturn = fmt.Sprintf("%f", value)
	} else if strings.Contains(typeValue, "bool") {
		valueReturn = fmt.Sprintf("%t", value)
	}
	return valueReturn
}

// convierte un byte a float64
func BytesToFloat64(bytes []byte) float64 {

	text := bytes // A decimal value represented as Latin-1 text

	f, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		fmt.Print("Error Conv:", err)
	}

	return f
}

// genera nombre de usuario
func UserGenerator(data map[string]string, conteo int) (user string, status bool, error string) {
	status = true
	if _, ok := data["first_name"]; !ok {
		return "", false, "Nombre  is required"
	}
	if _, ok := data["first_last_name"]; !ok {
		return "", false, "Nombre  is required"
	}
	if _, ok := data["second_last_name"]; !ok {
		return "", false, "Nombre  is required"
	}
	fullName := strings.Trim(data["first_name"], " ")
	arrayNames := strings.Split(fullName, " ")
	firstName := strings.ToUpper(strings.Trim(arrayNames[0], " "))
	secondName := ""
	if len(arrayNames) > 1 {
		secondName = strings.ToUpper(strings.Trim(arrayNames[1], " "))
	}
	firstLastName := strings.ToUpper(strings.Trim(data["first_last_name"], " "))
	secondLastName := strings.ToUpper(strings.Trim(data["second_last_name"], " "))

	switch conteo {
	case 0:
		user = firstName[0:1]
		user += firstLastName
	case 1:
		if len(secondName) > 0 {
			user = secondName[0:1]
			user += firstLastName
		} else {
			status = false
		}
	case 2:
		user = firstName[0:1]
		user += firstLastName
		user += secondLastName[0:1]
	case 3:
		if len(secondName) > 0 {
			user = secondName[0:1]
			user += firstLastName
			user += secondLastName[0:1]
		} else {
			status = false
		}
	case 4:
		user = firstName[0:1]
		user += secondLastName
	case 5:
		if len(secondName) > 0 {
			user = secondName[0:1]
			user += secondLastName
		} else {
			status = false
		}
	case 6:
		user = firstName[0:1]
		user += firstLastName[0:1]
		user += secondLastName
	case 7:
		if len(secondName) > 0 {
			user = secondName[0:1]
			user += firstLastName[0:1]
			user += secondLastName
		} else {
			status = false
		}
	default:
		user = firstName[0:1]
		user += firstLastName
		user += strconv.Itoa(conteo - 7)
	}

	return user, status, error
}

func IndexOf_String(arreglo []string, search string) int {
	for indice, valor := range arreglo {
		if valor == search {
			return indice
		}
	}
	// -1 porque no existe
	return -1
}

func SaveFile(file multipart.File, handle *multipart.FileHeader) (string, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	uuid := uuid.New().String()
	filename := "tmp/" + uuid + handle.Filename

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func IndexOf_String_Map(arreglo []map[string]interface{}, key, search string) int {
	for indice, valor := range arreglo {
		if valor[key] == search {
			return indice
		}
	}
	// -1 porque no existe
	return -1
}
