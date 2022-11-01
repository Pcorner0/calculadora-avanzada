package aritmetica

type Service interface {
	SumNumbers(req numbersRequest) (customNumber, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) SumNumbers(req numbersRequest) (customNumber, error) {
	var total customNumber

	for _, number := range req.ListNumbers {
		total = total + number
	}
	return total, nil
}
