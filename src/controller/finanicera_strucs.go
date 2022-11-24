package controller

type Cronograma struct {
	F_venc string  `json:"f_venc"` // fecha de vencimiento
	N_cuot int     `json:"n_cuot"` // numero de cuota
	S_scap float64 `json:"s_scap"` // saldo capital
	S_capi float64 `json:"s_capi"` // capital de cuota
	S_inte float64 `json:"s_inte"` // interes de cuota
	S_cuot float64 `json:"s_cuot"` // cuota a pagar
}
