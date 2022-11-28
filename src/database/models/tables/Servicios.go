package tables

import (
	"taxi-platcont-go/src/database/models"
	"taxi-platcont-go/src/libraries/date"

	"github.com/google/uuid"
)

func Servicios_GetSchema() ([]models.Base, string) {
	var servicios []models.Base
	tableName := "taxi_" + "servicios"
	id_serv := uuid.New().String()
	servicios = append(servicios, models.Base{ //id_serv
		Name:        "id_serv",
		Description: "id_serv",
		Required:    true,
		Important:   true,
		Default:     id_serv,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	servicios = append(servicios, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	servicios = append(servicios, models.Base{ //n_year
		Name:        "n_year",
		Description: "n_year",
		Update:      true,
		Required:    true,
		Default:     date.GetYear(),
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	servicios = append(servicios, models.Base{ //k_stad
		Name:        "k_stad",
		Description: "k_stad",
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 5,
		},
	})
	servicios = append(servicios, models.Base{ //f_digi
		Name:        "f_digi",
		Description: "f_digi",
		Default:     date.GetDateLocationString(),
	})
	servicios = append(servicios, models.Base{ //n_month
		Name:        "n_month",
		Description: "n_month",
		Required:    true,
		Default:     date.GetMonth(),
		Type:        "uint64",
		Uint: models.Uints{
			Max: 12,
		},
	})
	servicios = append(servicios, models.Base{ //f_fact
		Name:        "f_fact",
		Description: "f_fact",
		Required:    true,
		Update:      true,
		Default:     date.GetFechaLocationString(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	servicios = append(servicios, models.Base{ //n_docu
		Name:        "n_docu",
		Description: "n_docu",
		Update:      true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Number(),
			Min:  8,
			Max:  11,
		},
	})
	servicios = append(servicios, models.Base{ //c_plac
		Name:        "c_plac",
		Description: "c_plac",
		Update:      true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       7,
			Max:       7,
			UpperCase: true,
		},
	})
	servicios = append(servicios, models.Base{ //n_flot
		Name:        "n_flot",
		Description: "n_flot",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 1000,
		},
	})
	return servicios, tableName
}
