package jwt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Claims contains the claims of a jwt.
type Claims struct {
	claimsMap map[string]string
}

// NewClaim returns a new map representing the claims with the default values. The schema is detailed below.
//		claim["iis"] Issuer - string - identifies principal that issued the JWT;
//		claim["sub"] Subject - string - identifies the subject of the JWT;
//		claim["aud"] Audience - string - The "aud" (audience) claim identifies the recipients that the JWT is intended for. Each principal intended to process the JWT MUST identify itself with a value in the audience claim. If the principal processing the claim does not identify itself with a value in the aud claim when this claim is present, then the JWT MUST be rejected.
//		claim["exp"] Expiration time - time - The "exp" (expiration time) claim identifies the expiration time on or after which the JWT MUST NOT be accepted for processing.
//		claim["nbf"] Not before - time - Similarly, the not-before time claim identifies the time on which the JWT will start to be accepted for processing.
//		claim["iat"] Issued at - time - The "iat" (issued at) claim identifies the time at which the JWT was issued.
//		claim["jti"] JWT ID - string - case sensitive unique identifier of the token even among different issuers.
func NewClaim() *Claims {
	claimsMap := make(map[string]string)

	claims := &Claims{
		claimsMap: claimsMap,
	}

	claims.SetTime("iat", time.Now())

	return claims
}

// Set sets the claim in string form.
func (c *Claims) Set(key, value string) {
	c.claimsMap[key] = value
}

// Get returns the claim in string form.
func (c Claims) Get(key string) (string, error) {
	result, ok := c.claimsMap[key]
	if ok != true {
		return "", errors.New("claim doesn't exist")
	}
	return result, nil
}

// GetTime attempts to return a claim as a time.
func (c *Claims) GetTime(key string) (time.Time, error) {
	var err error
	var timeString string

	if timeString, err = c.Get(key); err != nil {
		return time.Unix(0, 0), errors.Wrap(err, "claim doesn't exist")
	}

	timeFloat, err := strconv.ParseFloat(timeString, 64)
	if err != nil {
		return time.Unix(0, 0), errors.Wrap(err, "claim isn't a valid float")
	}

	return time.Unix(int64(timeFloat), 0), nil
}

// SetTime sets the claim given to the specified time.
func (c *Claims) SetTime(key string, value time.Time) {
	c.Set(key, fmt.Sprintf("%d", value.Unix()))
}
