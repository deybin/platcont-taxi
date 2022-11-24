package tables

import (
	"taxi-platcont-go/src/database/models"
	"taxi-platcont-go/src/libraries/library"
)

func Clientes_GetSchema() ([]models.Base, string) {
	var clientes []models.Base
	tableName := "requ_" + "clientes"
	clientes = append(clientes, models.Base{ //c_docu
		Name:        "c_docu",
		Description: "c_docu",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  2,
			Max:  2,
		},
	})
	clientes = append(clientes, models.Base{ //n_docu
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Number(),
			Min:  1,
			Max:  11,
		},
	})
	clientes = append(clientes, models.Base{ //l_clie
		Name:        "l_clie",
		Description: "l_clie",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3,
			Max:       100,
			LowerCase: true,
		},
	})
	clientes = append(clientes, models.Base{ //k_gene
		Name:        "k_gene",
		Description: "k_gene",
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	clientes = append(clientes, models.Base{ //f_naci
		Name:        "f_naci",
		Description: "f_naci",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	clientes = append(clientes, models.Base{ //l_dire
		Name:        "l_dire",
		Description: "l_dire",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       20,
			Max:       400,
			LowerCase: true,
		},
	})
	clientes = append(clientes, models.Base{ //l_refe
		Name:        "l_refe",
		Description: "l_refe",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       400,
			LowerCase: true,
		},
	})
	clientes = append(clientes, models.Base{ //c_ubig
		Name:        "c_ubig",
		Description: "c_ubig",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  6,
			Max:  6,
		},
	})
	clientes = append(clientes, models.Base{ //n_tele
		Name:        "n_tele",
		Description: "n_tele",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  3,
			Max:  30,
		},
	})
	clientes = append(clientes, models.Base{ //n_celu
		Name:        "n_celu",
		Description: "n_celu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3,
			Max:       30,
			LowerCase: true,
		},
	})
	clientes = append(clientes, models.Base{ //l_obse
		Name:        "l_obse",
		Description: "l_obse",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       100,
			LowerCase: true,
		},
	})
	clientes = append(clientes, models.Base{ //id_regi
		Name:        "id_regi",
		Description: "id_regi",
		Required:    true,
		Default:     library.GetSession_key("us"),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	clientes = append(clientes, models.Base{ //l_emai
		Name:        "l_emai",
		Description: "l_emai",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       9,
			Max:       100,
			LowerCase: true,
		},
	})
	return clientes, tableName

}
