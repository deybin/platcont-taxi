package tables

import "taxi-platcont-go/src/database/models"

func Clientescars_GetSchema(n_docu string) ([]models.Base, string) {
	var clientescars []models.Base
	tableName := "taxi_" + "clientescars"
	clientescars = append(clientescars, models.Base{ //c_plac
		Name:        "c_plac",
		Description: "c_plac",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       7,
			Max:       7,
			UpperCase: true,
		},
	})
	clientescars = append(clientescars, models.Base{ //n_year
		Name:        "n_year",
		Description: "n_year",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	clientescars = append(clientescars, models.Base{ //l_marc
		Name:        "l_marc",
		Description: "l_marc",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			LowerCase: true,
		},
	})
	clientescars = append(clientescars, models.Base{ //l_mode
		Name:        "l_mode",
		Description: "l_mode",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			LowerCase: true,
		},
	})
	clientescars = append(clientescars, models.Base{ //l_color
		Name:        "l_color",
		Description: "l_color",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  4,
			Max:  70,
		},
	})
	clientescars = append(clientescars, models.Base{ //n_mode
		Name:        "n_mode",
		Description: "n_mode",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 3000,
		},
	})
	clientescars = append(clientescars, models.Base{ //n_seri
		Name:        "n_seri",
		Description: "n_seri",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       17,
			Max:       17,
			UpperCase: true,
		},
	})
	clientescars = append(clientescars, models.Base{ //n_pasa
		Name:        "n_pasa",
		Description: "n_pasa",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	clientescars = append(clientescars, models.Base{ //l_obse
		Name:        "l_obse",
		Description: "l_obse",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       100,
			LowerCase: true,
		},
	})
	clientescars = append(clientescars, models.Base{ //n_docu
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Update:      true,
		Default:     n_docu,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Number(),
			Min:  8,
			Max:  11,
		},
	})
	clientescars = append(clientescars, models.Base{ //k_stad
		Name:        "k_stad",
		Description: "k_stad",
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 2,
		},
	})
	return clientescars, tableName
}
