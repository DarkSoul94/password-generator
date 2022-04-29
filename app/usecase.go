package app

type Usecase interface {
	GeneratePassword(length, digitsCount int, withUpper, allowRepeat bool) (string, error)
}
