package jwt

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fwojciec/gqlmeetup"
)

const (
	defaultATDuration = 15 * time.Minute
	defaultRTDuration = 14 * 24 * time.Hour
	key               = ctxKey("jwt")
)

type ctxKey string

type accessTokenClaims struct {
	UserEmail string `json:"userEmail"`
	IsAdmin   bool   `json:"admin"`
	jwt.StandardClaims
}

type refreshTokenClaims struct {
	UserEmail string `json:"userEmail"`
	jwt.StandardClaims
}

var _ gqlmeetup.TokenService = (*TokenService)(nil)

// TokenService implements gqlmeetup.TokenService interface.
type TokenService struct {
	Secret               []byte
	AccessTokenDuration  time.Duration    // optional
	RefreshTokenDuration time.Duration    // optional
	Now                  func() time.Time // optional
}

// Issue issues a pair of tokens.
func (t *TokenService) Issue(userEmail string, isAdmin bool, pwdHash string) (*gqlmeetup.Tokens, error) {
	at, expAt, err := t.issueAccessToken(userEmail, isAdmin)
	if err != nil {
		return nil, err
	}
	rt, err := t.issueRefreshToken(userEmail, []byte(pwdHash))
	if err != nil {
		return nil, err
	}
	return &gqlmeetup.Tokens{
		Access:    at,
		Refresh:   rt,
		ExpiresAt: expAt,
	}, nil
}

func (t *TokenService) issueAccessToken(userEmail string, isAdmin bool) (string, int, error) {
	expAt := t.now().Add(t.atDuration()).Unix()
	atClaims := &accessTokenClaims{
		UserEmail: userEmail,
		IsAdmin:   isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(t.Secret)
	if err != nil {
		return "", 0, err
	}
	return token, int(expAt), nil
}

func (t *TokenService) issueRefreshToken(userEmail string, pwdHash []byte) (string, error) {
	rtClaims := &refreshTokenClaims{
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.now().Add(t.rtDuration()).Unix(),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	// we append the hash of the password to the refresh token secret so
	// that we are able to invalidate the token by changing the password
	token, err := rt.SignedString(append(t.Secret, pwdHash...))
	if err != nil {
		return "", err
	}
	return token, nil
}

// DecodeRefreshToken retrieves the user ID encoded in the Refresh Token without
// checking the validity of the token.
func (t *TokenService) DecodeRefreshToken(token string) (string, error) {
	claims := &refreshTokenClaims{}
	p := jwt.Parser{}
	_, _, err := p.ParseUnverified(token, claims)
	if err != nil || claims.UserEmail == "" {
		return "", gqlmeetup.ErrUnauthorized
	}
	return claims.UserEmail, nil
}

// CheckRefreshToken validates the refresh token and returns its payload.
func (t *TokenService) CheckRefreshToken(token, pwdHash string) (*gqlmeetup.RefreshTokenPayload, error) {
	claims := &refreshTokenClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, gqlmeetup.ErrUnauthorized
		}
		return append(t.Secret, []byte(pwdHash)...), nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, gqlmeetup.ErrUnauthorized
	}
	return &gqlmeetup.RefreshTokenPayload{UserEmail: claims.UserEmail}, nil
}

// CheckAccessToken validates the access token and returns its payload.
func (t *TokenService) CheckAccessToken(token string) (*gqlmeetup.AccessTokenPayload, error) {
	claims := &accessTokenClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, gqlmeetup.ErrUnauthorized
		}
		return t.Secret, nil
	})
	if err != nil || !parsedToken.Valid {
		return nil, gqlmeetup.ErrUnauthorized
	}
	return &gqlmeetup.AccessTokenPayload{UserEmail: claims.UserEmail, IsAdmin: claims.IsAdmin}, nil
}

// Retrieve retrieves the access token payload from the request context.
func (t *TokenService) Retrieve(ctx context.Context) (*gqlmeetup.AccessTokenPayload, error) {
	ap, ok := ctx.Value(key).(*gqlmeetup.AccessTokenPayload)
	if !ok {
		return nil, gqlmeetup.ErrUnauthorized
	}
	return ap, nil
}

// Store stores the access token payload in the request context.
func (t *TokenService) Store(ctx context.Context, ap *gqlmeetup.AccessTokenPayload) context.Context {
	return context.WithValue(ctx, key, ap)
}

func (t *TokenService) atDuration() time.Duration {
	if t.AccessTokenDuration == 0 {
		return defaultATDuration
	}
	return t.AccessTokenDuration
}

func (t *TokenService) rtDuration() time.Duration {
	if t.RefreshTokenDuration == 0 {
		return defaultRTDuration
	}
	return t.RefreshTokenDuration
}

func (t *TokenService) now() time.Time {
	if t.Now == nil {
		return time.Now()
	}
	return t.Now()
}
