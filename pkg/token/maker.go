package token

import "time"

type Maker interface {
	CreateToken(userId int, email string, username string, age int, audience string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
