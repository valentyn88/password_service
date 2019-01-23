package secure_password

import (
	"math/rand"

	"github.com/valentyn88/password_service"
)

const (
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specCharacters = "!\"#$%&'()*+,-./:;<=>?@"
	numbers        = "0123456789"
	vowels         = "aeioyAEIOY"
)

// PasswordService implementation of Creator.
type PasswordService struct {
}

// Create create password.
func (ps PasswordService) Create(lettersLen, specChLen, numbersLen int) password_service.Password {
	var res []rune

	if lettersLen > 0 {
		res = append(res, ps.randRunes([]rune(letters), lettersLen)...)
	}

	if specChLen > 0 {
		res = append(res, ps.randRunes([]rune(specCharacters), specChLen)...)
	}

	if numbersLen > 0 {
		res = append(res, ps.randRunes([]rune(numbers), numbersLen)...)
	}

	res = ps.shuffle(res)

	ps.vowelsToNumbers(res, []rune(numbers))

	return password_service.Password(string(res))
}

// randRunes pick random runes.
func (ps PasswordService) randRunes(rr []rune, n int) []rune {
	var res []rune

	for i := 0; i < n; i++ {
		res = append(res, rr[rand.Intn(len(rr))])
	}

	return res
}

// shuffle mix runes.
func (ps PasswordService) shuffle(rr []rune) []rune {
	var (
		n    = len(rr)
		res  = make([]rune, n)
		perm = rand.Perm(n)
	)

	for i, randIndex := range perm {
		res[i] = rr[randIndex]
	}

	return res
}

// vowelsToNumbers change vowels to random numbers.
func (ps PasswordService) vowelsToNumbers(rr, numbers []rune) {
	for i, r := range rr {
		for _, v := range vowels {
			if r == v {
				n := ps.randRunes(numbers, 1)
				rr[i] = n[0]
			}
		}
	}
}
