package service

import "github.com/YehorChervonyi/SimpleAPI/internal/factorial"

type IFactorialService interface {
	CalculateFacrotial(n int) int
}

type FactorialService struct {
	FactorialRepository *factorial.FactorialRepository
}

func NewFactorialService(repository *factorial.FactorialRepository) *FactorialService {
	return &FactorialService{FactorialRepository: repository}
}

func (s FactorialService) CalculateFactorial(n int) int {
	return s.FactorialRepository.Calculate(n)
}
