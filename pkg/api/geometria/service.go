package geometria

import (
	"errors"
	"math"
)

type Service interface {
	CosNumbers(req cosRequest) (float64, error)
	SinNumbers(req sinRequest) (float64, error)
	TanNumbers(req tanRequest) ([]float64, error)
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
	switch {
	case req.Avalor < 0 || req.Bvalor < 0 || req.Cvalor < 0 || req.DegreeSinAangle < 0 || req.DegreeSinBangle < 0 || req.DegreeSinCangle < 0:
		return 0, errors.New("los valores no deben ser negativos")
	case req.Avalor > 0 && req.Bvalor > 0 && req.DegreeSinBangle > 0:

		//math.Pi/180 convierte grados a radianes
		DegreeToRadianBangle := req.DegreeSinBangle * (math.Pi / 180)

		RadianAngle := math.Sin(DegreeToRadianBangle)

		//180/math.Pi convierte radianes a grados
		Aangle := (180 / math.Pi) * (math.Asin((req.Avalor * RadianAngle) / req.Bvalor))

		return Aangle, nil
	case req.Bvalor > 0 && req.DegreeSinAangle > 0 && req.DegreeSinBangle > 0:

		DegreeToRadianAangle := req.DegreeSinAangle * (math.Pi / 180)
		DegreeToRadianBangle := req.DegreeSinBangle * (math.Pi / 180)
		Avalor := (req.Bvalor * math.Sin(DegreeToRadianAangle)) / math.Sin(DegreeToRadianBangle)

		return Avalor, nil
	case req.Avalor > 0 && req.Cvalor > 0 && req.DegreeSinCangle > 0:

		//math.Pi/180 convierte grados a radianes
		DegreeToRadianCangle := req.DegreeSinCangle * (math.Pi / 180)

		RadianAngle := math.Sin(DegreeToRadianCangle)

		//180/math.Pi convierte radianes a grados
		Aangle := (180 / math.Pi) * (math.Asin((req.Avalor * RadianAngle) / req.Cvalor))

		return Aangle, nil
	case req.Bvalor > 0 && req.DegreeSinAangle > 0 && req.DegreeSinCangle > 0:

		DegreeToRadianAangle := req.DegreeSinAangle * (math.Pi / 180)
		DegreeToRadianCangle := req.DegreeSinCangle * (math.Pi / 180)
		Avalor := (req.Cvalor * math.Sin(DegreeToRadianAangle)) / math.Sin(DegreeToRadianCangle)

		return Avalor, nil
	case req.Avalor > 0 && req.Bvalor > 0 && req.DegreeSinAangle > 0:

		//math.Pi/180 convierte grados a radianes
		DegreeToRadianAangle := req.DegreeSinAangle * (math.Pi / 180)

		RadianAngle := math.Sin(DegreeToRadianAangle)

		//180/math.Pi convierte radianes a grados
		Bangle := (180 / math.Pi) * (math.Asin((req.Bvalor * RadianAngle) / req.Avalor))

		return Bangle, nil
	case req.Avalor > 0 && req.DegreeSinAangle > 0 && req.DegreeSinBangle > 0:

		DegreeToRadianAangle := req.DegreeSinAangle * (math.Pi / 180)
		DegreeToRadianBangle := req.DegreeSinBangle * (math.Pi / 180)
		Bvalor := (req.Avalor * math.Sin(DegreeToRadianBangle)) / math.Sin(DegreeToRadianAangle)

		return Bvalor, nil
	case req.Cvalor > 0 && req.Bvalor > 0 && req.DegreeSinCangle > 0:

		//math.Pi/180 convierte grados a radianes
		DegreeToRadianCangle := req.DegreeSinCangle * (math.Pi / 180)

		RadianAngle := math.Sin(DegreeToRadianCangle)

		//180/math.Pi convierte radianes a grados
		Bangle := (180 / math.Pi) * (math.Asin((req.Bvalor * RadianAngle) / req.Cvalor))

		return Bangle, nil
	case req.Cvalor > 0 && req.DegreeSinCangle > 0 && req.DegreeSinBangle > 0:

		DegreeToRadianCangle := req.DegreeSinCangle * (math.Pi / 180)
		DegreeToRadianBangle := req.DegreeSinBangle * (math.Pi / 180)
		Bvalor := (req.Cvalor * math.Sin(DegreeToRadianBangle)) / math.Sin(DegreeToRadianCangle)

		return Bvalor, nil
	case req.Cvalor > 0 && req.Avalor > 0 && req.DegreeSinAangle > 0:

		//math.Pi/180 convierte grados a radianes
		DegreeToRadianAangle := req.DegreeSinAangle * (math.Pi / 180)

		RadianAngle := math.Sin(DegreeToRadianAangle)

		//180/math.Pi convierte radianes a grados
		Cangle := (180 / math.Pi) * (math.Asin((req.Cvalor * RadianAngle) / req.Avalor))

		return Cangle, nil
	case req.Avalor > 0 && req.DegreeSinCangle > 0 && req.DegreeSinAangle > 0:

		DegreeToRadianCangle := req.DegreeSinCangle * (math.Pi / 180)
		DegreeToRadianAangle := req.DegreeSinAangle * (math.Pi / 180)
		Cvalor := (req.Avalor * math.Sin(DegreeToRadianCangle)) / math.Sin(DegreeToRadianAangle)

		return Cvalor, nil
	case req.Cvalor > 0 && req.Bvalor > 0 && req.DegreeSinBangle > 0:

		//math.Pi/180 convierte grados a radianes
		DegreeToRadianBangle := req.DegreeSinBangle * (math.Pi / 180)

		RadianAngle := math.Sin(DegreeToRadianBangle)

		//180/math.Pi convierte radianes a grados
		Cangle := (180 / math.Pi) * (math.Asin((req.Cvalor * RadianAngle) / req.Bvalor))

		return Cangle, nil
	case req.Bvalor > 0 && req.DegreeSinCangle > 0 && req.DegreeSinBangle > 0:

		DegreeToRadianCangle := req.DegreeSinCangle * (math.Pi / 180)
		DegreeToRadianBangle := req.DegreeSinBangle * (math.Pi / 180)
		Cvalor := (req.Bvalor * math.Sin(DegreeToRadianCangle)) / math.Sin(DegreeToRadianBangle)

		return Cvalor, nil
	}
	return 0, errors.New("faltan valores")

}

func (s service) CosNumbers(req cosRequest) (float64, error) {
	switch {
	case req.Avalor < 0 || req.Bvalor < 0 || req.Cvalor < 0 || req.DegreeCosAangle < 0 || req.DegreeCosBangle < 0 || req.DegreeCosCangle < 0:
		return 0, errors.New("los valores no deben ser negativos")
	case req.Avalor > 0 && req.Bvalor > 0 && req.Cvalor > 0:

		//180/math.Pi convierte radianes a grados
		Bangle := (180 / math.Pi) * (math.Acos((math.Pow(req.Avalor, 2) + math.Pow(req.Cvalor, 2) - math.Pow(req.Bvalor, 2)) / (2 * req.Avalor * req.Cvalor)))

		return Bangle, nil
	case req.Avalor > 0 && req.Cvalor > 0 && req.DegreeCosBangle > 0:
		DegreeToRadianBangle := req.DegreeCosBangle * (math.Pi / 180)
		Bvalor := math.Sqrt(math.Pow(req.Avalor, 2) + math.Pow(req.Cvalor, 2) - 2*req.Avalor*req.Cvalor*math.Cos(DegreeToRadianBangle))

		return Bvalor, nil
	case req.Bvalor > 0 && req.Cvalor > 0 && req.DegreeCosAangle > 0:
		DegreeToRadianAangle := req.DegreeCosAangle * (math.Pi / 180)
		Avalor := math.Sqrt(math.Pow(req.Bvalor, 2) + math.Pow(req.Cvalor, 2) - 2*req.Bvalor*req.Cvalor*math.Cos(DegreeToRadianAangle))

		return Avalor, nil
	case req.Avalor > 0 && req.Bvalor > 0 && req.DegreeCosCangle > 0:
		DegreeToRadianCangle := req.DegreeCosCangle * (math.Pi / 180)
		Cvalor := math.Sqrt(math.Pow(req.Avalor, 2) + math.Pow(req.Bvalor, 2) - 2*req.Avalor*req.Bvalor*math.Cos(DegreeToRadianCangle))

		return Cvalor, nil
	}
	return 0, errors.New("faltan valores")

}

func (s service) TanNumbers(req tanRequest) ([]float64, error) {
	switch {
	case req.Avalor < 0 || req.Bvalor < 0 || req.Cvalor < 0 || req.DegreeTanAangle < 0 || req.DegreeTanBangle < 0 || req.DegreeTanCangle < 0:
		return []float64{0}, errors.New("los valores no deben ser negativos")
	case req.Bvalor > 0 && req.DegreeTanAangle > 0 && req.DegreeTanBangle > 0:
		//(math.Pi / 180) convierte radianes a grados
		DegreeToRadianAangle := req.DegreeTanAangle * (math.Pi / 180)
		DegreeToRadianBangle := req.DegreeTanBangle * (math.Pi / 180)
		AplusB := DegreeToRadianAangle + DegreeToRadianBangle
		AminusB := DegreeToRadianAangle - DegreeToRadianBangle
		Avalor := (req.Bvalor * ((math.Tan(AminusB / 2)) + (math.Tan(AplusB / 2)))) / ((math.Tan(AplusB / 2)) - (math.Tan(AminusB / 2)))

		return []float64{Avalor}, nil

	case req.Avalor > 0 && req.DegreeTanAangle > 0 && req.DegreeTanBangle > 0:
		// (math.Pi / 180) convierte radianes a grados
		DegreeToRadianAangle := req.DegreeTanAangle * (math.Pi / 180)
		DegreeToRadianBangle := req.DegreeTanBangle * (math.Pi / 180)
		AplusB := DegreeToRadianAangle + DegreeToRadianBangle
		AminusB := DegreeToRadianAangle - DegreeToRadianBangle
		Bvalor := (req.Avalor * ((math.Tan(AplusB / 2)) - (math.Tan(AminusB / 2)))) / ((math.Tan(AminusB / 2)) + (math.Tan(AplusB / 2)))

		return []float64{Bvalor}, nil

	case req.Cvalor > 0 && req.DegreeTanBangle > 0 && req.DegreeTanCangle > 0:
		//(math.Pi / 180) convierte radianes a grados
		DegreeToRadianBangle := req.DegreeTanBangle * (math.Pi / 180)
		DegreeToRadianCangle := req.DegreeTanCangle * (math.Pi / 180)
		BplusC := DegreeToRadianBangle + DegreeToRadianCangle
		BminusC := DegreeToRadianBangle - DegreeToRadianCangle
		Bvalor := (req.Cvalor * ((math.Tan(BminusC / 2)) + (math.Tan(BplusC / 2)))) / ((math.Tan(BplusC / 2)) - (math.Tan(BminusC / 2)))

		return []float64{Bvalor}, nil

	case req.Avalor > 0 && req.DegreeTanAangle > 0 && req.DegreeTanCangle > 0:
		//(math.Pi / 180) convierte radianes a grados
		DegreeToRadianCangle := req.DegreeTanCangle * (math.Pi / 180)
		DegreeToRadianAangle := req.DegreeTanAangle * (math.Pi / 180)
		CplusA := DegreeToRadianCangle + DegreeToRadianAangle
		CminusA := DegreeToRadianCangle - DegreeToRadianAangle
		Cvalor := (req.Avalor * ((math.Tan(CminusA / 2)) + (math.Tan(CplusA / 2)))) / ((math.Tan(CplusA / 2)) - (math.Tan(CminusA / 2)))

		return []float64{Cvalor}, nil

	case req.Bvalor > 0 && req.DegreeTanBangle > 0 && req.DegreeTanCangle > 0:
		// (math.Pi / 180) convierte radianes a grados
		DegreeToRadianBangle := req.DegreeTanBangle * (math.Pi / 180)
		DegreeToRadianCangle := req.DegreeTanCangle * (math.Pi / 180)
		BplusC := DegreeToRadianBangle + DegreeToRadianCangle
		BminusC := DegreeToRadianBangle - DegreeToRadianCangle
		Cvalor := (req.Bvalor * ((math.Tan(BplusC / 2)) - (math.Tan(BminusC / 2)))) / ((math.Tan(BminusC / 2)) + (math.Tan(BplusC / 2)))

		return []float64{Cvalor}, nil
	}

	switch {
	case req.DegreeTanAangle < 0 || req.DegreeTanBangle < 0 || req.DegreeTanCangle < 0:
		return []float64{0}, errors.New("los valores no deben ser negativos")
	case req.DegreeTanAangle > 0 && req.DegreeTanBangle > 0:
		Cangle := 180 - req.DegreeTanAangle - req.DegreeTanBangle
		return []float64{Cangle}, nil

	case req.DegreeTanBangle > 0 && req.DegreeTanCangle > 0:
		Aangle := 180 - req.DegreeTanBangle - req.DegreeTanCangle
		return []float64{Aangle}, nil

	case req.DegreeTanAangle > 0 && req.DegreeTanCangle > 0:
		Bangle := 180 - req.DegreeTanAangle - req.DegreeTanCangle
		return []float64{Bangle}, nil
	}

	switch {
	case req.Avalor < 0 || req.Bvalor < 0 || req.Cvalor < 0 || req.DegreeTanAangle < 0 || req.DegreeTanBangle < 0 || req.DegreeTanCangle < 0:
		return []float64{0}, errors.New("los valores no deben ser negativos")
	case req.DegreeTanCangle > 0 && req.Avalor > 0 && req.Bvalor > 0:
		AplusB := 180 - req.DegreeTanCangle
		AplusBtoRadian := AplusB * (math.Pi / 180)

		Bangle := ((((AplusBtoRadian / 2) * 2) - (2 * math.Atan(((req.Avalor-req.Bvalor)*math.Tan(AplusBtoRadian/2))/(req.Avalor+req.Bvalor)))) / 2) * (180 / math.Pi)

		Aangle := 180 - Bangle - req.DegreeTanCangle

		return []float64{Aangle, Bangle}, nil

	case req.DegreeTanBangle > 0 && req.Avalor > 0 && req.Cvalor > 0:
		AplusC := 180 - req.DegreeTanBangle
		AplusCtoRadian := AplusC * (math.Pi / 180)

		Cangle := ((((AplusCtoRadian / 2) * 2) - (2 * math.Atan(((req.Avalor-req.Cvalor)*math.Tan(AplusCtoRadian/2))/(req.Avalor+req.Cvalor)))) / 2) * (180 / math.Pi)

		Aangle := 180 - Cangle - req.DegreeTanBangle

		return []float64{Aangle, Cangle}, nil

	case req.DegreeTanAangle > 0 && req.Bvalor > 0 && req.Cvalor > 0:
		BplusC := 180 - req.DegreeTanAangle
		BplusCtoRadian := BplusC * (math.Pi / 180)

		Cangle := ((((BplusCtoRadian / 2) * 2) - (2 * math.Atan(((req.Bvalor-req.Cvalor)*math.Tan(BplusCtoRadian/2))/(req.Bvalor+req.Cvalor)))) / 2) * (180 / math.Pi)

		Bangle := 180 - Cangle - req.DegreeTanAangle

		return []float64{Bangle, Cangle}, nil
	}
	return []float64{0}, errors.New("faltan valores")

}
