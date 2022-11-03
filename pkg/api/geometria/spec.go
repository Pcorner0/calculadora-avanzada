package geometria


type sinRequest struct {
	DegreeSinAangle float64`json:"sin_a_angle"`
	DegreeSinBangle float64`json:"sin_b_angle"`
	Avalor float64`json:"A_valor"`
	Bvalor float64`json:"B_valor"`
}

type cosRequest struct {
	DegreeCosBangle float64`json:"cos_b_angle"`
	Avalor float64`json:"A_valor"`
	Bvalor float64`json:"B_valor"`
	Cvalor float64`json:"C_valor"`
}