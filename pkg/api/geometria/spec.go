package geometria


type sinRequest struct {
	DegreeSinAangle float64`json:"sin_a_angle"`
	DegreeSinBangle float64`json:"sin_b_angle"`
	DegreeSinCangle float64`json:"sin_c_angle"`
	Avalor float64`json:"A_valor"`
	Bvalor float64`json:"B_valor"`
	Cvalor float64`json:"C_valor"`
}

type cosRequest struct {
	DegreeCosAangle float64`json:"cos_a_angle"`
	DegreeCosBangle float64`json:"cos_b_angle"`
	DegreeCosCangle float64`json:"cos_c_angle"`
	Avalor float64`json:"A_valor"`
	Bvalor float64`json:"B_valor"`
	Cvalor float64`json:"C_valor"`
}

type tanRequest struct {
	DegreeTanAangle float64`json:"tan_a_angle"`
	DegreeTanBangle float64`json:"tan_b_angle"`
	DegreeTanCangle float64`json:"tan_c_angle"`
	Avalor float64`json:"A_valor"`
	Bvalor float64`json:"B_valor"`
	Cvalor float64`json:"C_valor"`
	
}