package bcrypt

import (
	"github.com/fwojciec/gqlmeetup"
	"golang.org/x/crypto/bcrypt"
)

var _ gqlmeetup.PasswordService = (*PasswordService)(nil)

// PasswordService implements gqlmeetup.PasswordService interface.
type PasswordService struct{}

// Check compares a password with the stored hash.
func (s *PasswordService) Check(pwdHash, pwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwd)); err != nil {
		return gqlmeetup.ErrPwdCheck
	}
	return nil
}

// Hash creates a hash of a password.
func (s *PasswordService) Hash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
