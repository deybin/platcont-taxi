package tables

import (
	"taxi-platcont-go/src/database/models"
	"taxi-platcont-go/src/libraries/date"
)

func Serviciosdetalle_GetSchema(id_serv string) ([]models.Base, string) {
	var serviciosdetalle []models.Base
	tableName := "taxi_" + "serviciosdetalle"
	serviciosdetalle = append(serviciosdetalle, models.Base{ //id_serv
		Name:        "id_serv",
		Description: "id_serv",
		Required:    true,
		Important:   true,
		Type:        "string",
		Default:     id_serv,
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	serviciosdetalle = append(serviciosdetalle, models.Base{ //n_year
		Name:        "n_year",
		Description: "n_year",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	serviciosdetalle = append(serviciosdetalle, models.Base{ //n_month
		Name:        "n_month",
		Description: "n_month",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 12,
		},
	})
	serviciosdetalle = append(serviciosdetalle, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	serviciosdetalle = append(serviciosdetalle, models.Base{ //k_stad
		Name:        "k_stad",
		Description: "k_stad",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 5,
		},
	})
	serviciosdetalle = append(serviciosdetalle, models.Base{ //f_pago
		Name:        "f_pago",
		Description: "f_pago",
		Required:    true,
		Default:     date.GetFechaLocationString(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	return serviciosdetalle, tableName
}
