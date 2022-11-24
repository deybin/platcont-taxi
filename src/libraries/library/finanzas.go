package library

import (
	"math"
	"strings"
	"taxi-platcont-go/src/controller"
	"taxi-platcont-go/src/libraries/date"
)

func TIR(monto float64, numero_cuotas int, periodo int, tasa float64) float64 {
	//flujo_caja:-> las cuotas que se pagan en el periodo
	var m float64
	periodo_float := float64(periodo)
	numero_cuotas_float := float64(numero_cuotas)
	tasa = tasa / 100
	if periodo == 7 {
		m = (7.5 * numero_cuotas_float) / 30
	} else {
		m = (periodo_float * numero_cuotas_float) / 30
	}
	flujo_caja := (((monto * tasa) * m) + monto) / numero_cuotas_float
	TED := math.Pow(1+tasa, float64(1)/30) - 1
	TEP := math.Pow(1+TED, periodo_float) - 1
	vpn := float64(7)
	tr := float64(0)
	for vpn >= 0 {
		val := float64(0)
		for i := 1; i <= numero_cuotas; i++ {
			val += (flujo_caja / math.Pow(1+TEP, float64(i)))
		}

		vpn = -monto + val
		if vpn-1 >= 5 {
			TEP += .00001
		}
		tr = TEP
		TEP += .0000001
	}
	nt := math.Pow(1+tr, 1/periodo_float) - 1
	TIR := math.Pow(1+nt, float64(30)) - 1
	return TIR
}

// Retorna la cuota según el periodo que se deba pagar
func GenCuotasCredits(methodType int, monto float64, TEMR float64, periodo int, numero_cuotas int) float64 {
	//methodType -> es método utilizado para generar créditos esta la tabla fina_sisprop
	//monto -> monto del crédito
	//TEMR -> tasa de interés mensual obtenida del TIR
	//periodo -> periodo de pago
	//numero_cuotas -> numero de cuotas

	var cuota float64
	periodo_float := float64(periodo)
	numero_cuotas_float := float64(numero_cuotas)
	TED := (math.Pow(1+TEMR, float64(1)/30) - 1)
	TEP := math.Pow(1+TED, periodo_float) - 1

	if methodType == 0 {
		cuota = (monto * (TEP * math.Pow(1+TEP, numero_cuotas_float))) / (math.Pow(1+TEP, numero_cuotas_float) - 1)
	} else if methodType == 1 {
		if periodo == 1 {
			mes := (numero_cuotas_float * periodo_float) / 30
			cuota = (monto + ((monto * TEMR) * mes)) / numero_cuotas_float
		} else if periodo == 7 {
			mes := ((numero_cuotas_float * periodo_float) + (numero_cuotas_float / 2)) / 30
			cuota = (monto + ((monto * TEMR) * mes)) / numero_cuotas_float
		} else if periodo == 14 {
			mes := ((numero_cuotas_float * periodo_float) + (numero_cuotas_float / 1)) / 30
			cuota = (monto + ((monto * TEMR) * mes)) / numero_cuotas_float
		} else if periodo == 15 {
			mes := (numero_cuotas_float * periodo_float) / 30
			cuota = (monto + ((monto * TEMR) * mes)) / numero_cuotas_float
		} else if periodo == 30 {
			cuota = (monto + ((monto * TEMR) * numero_cuotas_float)) / numero_cuotas_float
		}
	}
	return cuota
}

//Retorna el Interés Diferido por El crédito según los Dias de gracia
func GetInterestDiferido(monto float64, TEMR float64, days int) float64 {
	/* INFORMACIÓN DE LA FUNCIÓN
	Formulas Utilizadas
		Cálculo de interés para periodos de gracia
		=========FORMULA-PERIODO DE GRACIA===============
			Interes_diferido=M*[(1+TED)^t-1]
			Donde:
				Interes_diferido : Interés para días de gracia
				M : Monto del préstamo
				t : Número de días de gracia
				TED : Tasa de interés Efectiva Diaria
		==================================================
	*/
	var interest float64
	TED := (math.Pow(1+TEMR, float64(1)/30) - 1)
	interest = monto * (math.Pow(1+TED, float64(days)-1))
	return interest
}

// retorna cronograma de pago de un crédito
func GetCronograma(_type int, monto float64, fecha string, TEM float64, periodo int, numero_cuotas int, arg_days_gracia ...int) []controller.Cronograma {

	/* INFORMACIÓN DE LA FUNCIÓN
	Variables Solicitadas:
		-types->método que se utilizara para la amortización
			0- método Frances
			1- Peru interés y capital fijo
			2- método alemán/italiano

		-monto->monto del crédito
		-fecha-> la fecha que se desembolsara desde ahi se generara comprobante
		-TEM->TEM tasa real porque el sistema almacena dos tasas TEM -> Tasa Efectiva Mensual TEMR-> Tasa Efectiva Mensual Real
		-periodo-> numero de periodos o también llamado frecuencia de pago
		-numero_cuotas-> numero de cuotas ya se mensual semanal o quincenal hast mensual
		-days_gracia-> Dias de gracias otorgados (Opcional) no se exige esta variable

	Formulas Utilizadas
		Cálculo de las cuotas a pagar por periodo
		=========FORMULA-PERIODO DE GRACIA================
			cuota= [M*(TEP*(1+TEP)^n)]/[(1+TEP)^n)-1]
			Donde:
				cuota : cuota periódicas
				M : Monto del préstamo
				n : Número cuotas
				TEP : Tasa de interés Efectiva Periodo
		==================================================
	*/

	var resultado []controller.Cronograma

	var cuota float64
	periodo_float := float64(periodo)
	numero_cuotas_float := float64(numero_cuotas)
	days_gracias := 0
	year_days_feriados := ""
	var lista_days_feriados []string

	TED := (math.Pow(1+TEM, float64(1)/30) - 1)
	TEP := math.Pow(1+TED, periodo_float) - 1
	if _type == 0 {
		cuota = (monto * (TEP * math.Pow(1+TEP, numero_cuotas_float))) / (math.Pow(1+TEP, numero_cuotas_float) - 1)
		// cuota_redondeada := math.Ceil(cuota*10) / 10
		// diferencia_cuotas := cuota_redondeada - cuota
		// fmt.Println("Diferencia de cuotas: ", diferencia_cuotas)
		var interest_diferido float64
		if len(arg_days_gracia) == 1 {
			if arg_days_gracia[0] != 0 {
				days_gracias = arg_days_gracia[0]
				interest_diferido = GetInterestDiferido(monto, TED, days_gracias)
			}
		}

		saldo_capital := monto
		date_interaction := fecha
		date_base := date_interaction
		diferencia_date_acumulado := 0
		y := 0

		for i := 0; i <= numero_cuotas; i++ {
			interest := saldo_capital * TEP
			if i == 1 && days_gracias != 0 {
				diferencia_date_acumulado += days_gracias
				date_interaction = date.SumarDate(date_interaction, diferencia_date_acumulado)
			}
			capital := cuota - interest

			if i > 0 {
				diferencia_date_acumulado += periodo
				date_interaction = date.SumarDate(date_base, diferencia_date_acumulado)
				year_get := strings.Split(date_interaction, "/")[2]
				if year_get != year_days_feriados {
					year_days_feriados = year_get
					lista_days_feriados = GetDiasFeriados(year_days_feriados)
				}

				if date.IsItHoliday(date_interaction, lista_days_feriados) {
					date_temp := ""
					for k := 1; k <= 7; k++ {
						date_temp = date.SumarDate(date_interaction, k)
						year_get := strings.Split(date_temp, "/")[2]
						if year_get != year_days_feriados {
							year_days_feriados = year_get
							lista_days_feriados = GetDiasFeriados(year_days_feriados)
						}
						if date.IsItHoliday(date_temp, lista_days_feriados) == false {
							if periodo == 1 {
								diferencia_date_acumulado += k
							}
							date_interaction = date_temp
							break
						}
					}
				}
			}

			if i == 0 {
				resultado = append(resultado, controller.Cronograma{
					F_venc: date_interaction,
					N_cuot: i,
					S_scap: saldo_capital,
					S_capi: 0,
					S_inte: 0,
					S_cuot: 0,
				})
				capital = 0
			} else if i == 1 && days_gracias != 0 {
				resultado = append(resultado, controller.Cronograma{
					F_venc: date_interaction,
					N_cuot: i,
					S_scap: saldo_capital,
					S_capi: capital,
					S_inte: interest + interest_diferido,
					S_cuot: cuota + interest_diferido,
				})
			} else {
				resultado = append(resultado, controller.Cronograma{
					F_venc: date_interaction,
					N_cuot: i,
					S_scap: saldo_capital,
					S_capi: capital,
					S_inte: interest,
					S_cuot: cuota,
				})

			}

			saldo_capital -= capital
			y++
		}
	}
	return resultado
}

func GetDiasFeriados(year_filter string) []string {
	// fmt.Print(year_filter)
	var days_feriados []string
	days_feriados = append(days_feriados, "01-01")
	days_feriados = append(days_feriados, "05-01")
	days_feriados = append(days_feriados, "06-29")
	days_feriados = append(days_feriados, "07-28")
	days_feriados = append(days_feriados, "07-29")
	days_feriados = append(days_feriados, "08-30")
	days_feriados = append(days_feriados, "10-08")
	days_feriados = append(days_feriados, "11-01")
	days_feriados = append(days_feriados, "12-08")
	days_feriados = append(days_feriados, "12-25")

	return days_feriados
}

/**
 * genera el interés moratoria de una cuota según  la diferencia de Dias entre la fecha de vencimiento y la fecha actual
 * @param cuota {float64}: Monto de cuota a pagar.
 * @param importe_mora {float64}: Importe o Tasa de interés moratoria de periodo mensual.
 * @param diferencia_days {int64}: diferencia de días entre la fecha de hoy y la fecha de vencimiento de la cuota.
 * @param typ_mora{int64}: Tipo de mora: 0 = po porcentaje TMEM, 1 = por importe.
 * @return {float64}: Importe de la mora.
 */
func GenteMoraByCuotaCredit(cuota float64, importe_mora float64, diferencia_days float64, typ_mora int64) (mora float64) {

	if typ_mora == 0 {
		TMED := (math.Pow(1+importe_mora, float64(1)/30) - 1)
		mora = cuota * TMED
		mora = mora * diferencia_days

	} else {
		mora := diferencia_days * importe_mora
		mora = mora * diferencia_days
	}
	mora = math.Ceil(mora*10) / 10
	return
}
