package geometria

import (
	"errors"
	"math"
)

type Service interface {
	CosNumbers(req cosRequest) (float64, error)
	SinNumbers(req sinRequest) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) SinNumbers(req sinRequest) (float64, error) {
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

func (s service) CosNumbers(req cosRequest) (float64, error) {
	if req.Avalor > 0 && req.Bvalor > 0 && req.Cvalor > 0 {

		//180/math.Pi convierte radianes a grados
		total := (180 / math.Pi) * (math.Acos((math.Pow(req.Avalor,2) + math.Pow(req.Cvalor,2) - math.Pow(req.Bvalor,2))/(2 * req.Avalor * req.Cvalor)))

		return total, nil
	} else if req.Avalor > 0 && req.Cvalor > 0 && req.DegreeCosBangle > 0 {
		var DegreeToRadianBangle float64
		DegreeToRadianBangle = req.DegreeCosBangle * (math.Pi / 180)
		total := math.Sqrt(math.Pow(req.Avalor,2) + math.Pow(req.Cvalor,2) - 2 * req.Avalor * req.Cvalor * math.Cos(DegreeToRadianBangle))

		return total, nil
	}
	return 0, errors.New("faltan valores")

}