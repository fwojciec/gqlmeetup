package scs

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/fwojciec/gqlmeetup"
)

// SessionService is an implementation of the gqlgen.SessionService interface.
// This implementation requires that the graphql handler is wrapped in the scs
// LoadAndSave middleware.
type SessionService struct {
	sm *scs.SessionManager
}

var _ gqlmeetup.SessionService = (*SessionService)(nil)

// New returns a new instance of SessionService.
func New() *SessionService {
	sm := scs.New()
	return &SessionService{sm: sm}
}

// Login saves the ID of the user in the session.
func (s *SessionService) Login(ctx context.Context, user *gqlmeetup.User) error {
	if err := s.sm.RenewToken(ctx); err != nil {
		return err
	}
	b, err := s.marshalUser(user)
	if err != nil {
		return err
	}
	s.sm.Put(ctx, "user", b)
	return nil
}

// Logout logs the user out.
func (s *SessionService) Logout(ctx context.Context) error {
	return s.sm.Destroy(ctx)
}

// GetUser gets the user email from the current session. The zero value for
// a string ("") is returned if the email does not exist in the session (i.e.
// the user is not logged in).
func (s *SessionService) GetUser(ctx context.Context) *gqlmeetup.User {
	b := s.sm.GetBytes(ctx, "user")
	if b == nil {
		return nil
	}
	user, err := s.unmarshalUser(b)
	if err != nil {
		return nil
	}
	return user
}

// Middleware returns a middleware that enables session functionality for the
// wrapped handlers.
func (s *SessionService) Middleware(next http.Handler) http.Handler {
	return s.sm.LoadAndSave(next)
}

func (s *SessionService) marshalUser(user *gqlmeetup.User) ([]byte, error) {
	// no need to store the password hash in the session
	user.Password = ""
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(user); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *SessionService) unmarshalUser(b []byte) (*gqlmeetup.User, error) {
	user := &gqlmeetup.User{}
	if err := json.NewDecoder(bytes.NewBuffer(b)).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
