package aritmetica

import (
	"errors"
	"math"
)

type Service interface {
	SumSubtNumbers(req numbersRequest) (float64, error)
	DivideTwoNumbers(req divideRequest) (float64, error)
	MultiplyNumbers(req multiplyRequest) (float64, error)
	RootNumbers(req RootRequest) (float64, error)
	PowerNumbers(req PowerRequest) (float64, error)
	SinNumbers(req SinRequest) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) SumSubtNumbers(req numbersRequest) (float64, error) {
	var total float64

	for _, number := range req.ListNumbers {
		total = total + number
	}
	return total, nil
}

func (s service) DivideTwoNumbers(req divideRequest) (float64, error) {

	if req.Denominator == 0 {
		return 0, errors.New("indefinido")
	}

	total := req.Numerator / req.Denominator

	return total, nil
}

func (s service) MultiplyNumbers(req multiplyRequest) (float64, error) {
	var total float64 = 1

	for _, number := range req.ListNumbers {
		total = total * number
	}
	return total, nil
}

func (s service) RootNumbers(req RootRequest) (float64, error) {
	var total float64
	var RootDecimal float64

	if req.RootNumber%2 == 0 && req.NumberToRoot < 0 {
		return 0, errors.New("raíz negativa")
	}
	RootDecimal = 1 / float64(req.RootNumber)
	total = math.Pow(req.NumberToRoot, RootDecimal)

	return total, nil
}

func (s service) PowerNumbers(req PowerRequest) (float64, error) {
	var total float64 = 1

	total = math.Pow(req.NumberToPower, req.PowerNumber)

	return total, nil
}

func (s service) SinNumbers(req SinRequest) (float64, error) {
	if req.Avalor > 0 && req.Bvalor > 0 && req.DegreeSinAangle > 0 {
		var DegreeToRadianAangle float64
		//math.Pi/180 convierte grados a radianes
		DegreeToRadianAangle = req.DegreeSinAangle * (math.Pi / 180)

		RadianAngle := math.Sin(DegreeToRadianAangle)

		//180/math.Pi convierte radianes a grados
		total := (180 / math.Pi) * (math.Asin((req.Bvalor * RadianAngle) / req.Avalor))

		return total, nil
	} else if req.Avalor > 0 && req.DegreeSinAangle > 0 && req.DegreeSinBangle > 0 {
		var DegreeToRadianAangle float64
		var DegreeToRadianBangle float64
		DegreeToRadianAangle = req.DegreeSinAangle * (math.Pi / 180)
		DegreeToRadianBangle = req.DegreeSinBangle * (math.Pi / 180)
		total := ((req.Avalor*math.Sin(DegreeToRadianBangle)))/math.Sin(DegreeToRadianAangle)

		return total,nil
	}
	return 0, errors.New("faltan valores")

}
