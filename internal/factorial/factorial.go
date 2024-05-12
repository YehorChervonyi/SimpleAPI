package factorial

type IFactorialRepository interface {
	Calculate(n int) int
}

type FactorialRepository struct{}

func NewFactorialRepository() *FactorialRepository {
	return &FactorialRepository{}
}
func (f *FactorialRepository) Calculate(n int) int {
	if n == 0 {
		return 1
	}
	return n * f.Calculate(n-1)
}
