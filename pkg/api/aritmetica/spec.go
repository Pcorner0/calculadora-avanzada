package aritmetica


type numbersRequest struct {
	ListNumbers []float64 `json:"list_numbers"`
}
type divideRequest struct {
	Numerator float64 `json:"numerator"`
	Denominator float64 `json:"denominator"`
}
type multiplyRequest struct {
	ListNumbers []float64 `json:"list_numbers"`
}
type RootRequest struct {
	RootNumber int64`json:"root_number"`
	NumberToRoot float64 `json:"number_to_root"`
}
type PowerRequest struct {
	PowerNumber float64`json:"power_number"`
	NumberToPower float64 `json:"number_to_power"`
}
type SinRequest struct {
	DegreeSinAangle float64`json:"sin_a_angle"`
	DegreeSinBangle float64`json:"sin_b_angle"`
	Avalor float64`json:"A_valor"`
	Bvalor float64`json:"B_valor"`
}

