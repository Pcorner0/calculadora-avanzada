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
type rootRequest struct {
	RootNumber int64`json:"root_number"`
	NumberToRoot float64 `json:"number_to_root"`
}
type powerRequest struct {
	PowerNumber float64`json:"power_number"`
	NumberToPower float64 `json:"number_to_power"`
}


