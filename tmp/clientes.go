package tables
import "taxi-platcont-go/src/models" 
func clientes_GetSchema() ([]models.Base, string) {
	var clientes []models.Base
	tableName := "requ_" + "clientes"
	clientes = append(clientes, models.Base{//k_gene
		Name:"k_gene",
		Description:"k_gene",
		Required: true,
		Update: true,
		Type:"uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	clientes = append(clientes, models.Base{//n_docu
		Name:"n_docu",
		Description:"n_docu",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:1.100000,
			Max:11,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//l_clie
		Name:"l_clie",
		Description:"l_clie",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:10.000000,
			Max:100,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//f_naci
		Name:"f_naci",
		Description:"f_naci",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Date:true,
		},
	})
	clientes = append(clientes, models.Base{//l_dire
		Name:"l_dire",
		Description:"l_dire",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:40.000000,
			Max:400,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//l_dire2
		Name:"l_dire2",
		Description:"l_dire2",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:40.000000,
			Max:400,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//l_refe
		Name:"l_refe",
		Description:"l_refe",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:40.000000,
			Max:400,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//l_refe2
		Name:"l_refe2",
		Description:"l_refe2",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:40.000000,
			Max:400,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//c_ubig
		Name:"c_ubig",
		Description:"c_ubig",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:0.600000,
			Max:6,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//c_ubig2
		Name:"c_ubig2",
		Description:"c_ubig2",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:0.600000,
			Max:6,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//l_srcs
		Name:"l_srcs",
		Description:"l_srcs",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:30.000000,
			Max:300,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//n_tele
		Name:"n_tele",
		Description:"n_tele",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:3.000000,
			Max:30,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//n_celu
		Name:"n_celu",
		Description:"n_celu",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:3.000000,
			Max:30,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//l_emai
		Name:"l_emai",
		Description:"l_emai",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:10.000000,
			Max:100,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//l_obse
		Name:"l_obse",
		Description:"l_obse",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:10.000000,
			Max:100,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//c_docu
		Name:"c_docu",
		Description:"c_docu",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:0.200000,
			Max:2,
			UpperCase:true,
		},
	})
	clientes = append(clientes, models.Base{//id_regi
		Name:"id_regi",
		Description:"id_regi",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
		},
	})
	return clientes, tableName
}
