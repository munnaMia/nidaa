package domain

type Header struct {
	ALG string `json:"alg"`
	TYP string `json:"typ"`
}

type Payload struct {
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
	IAT   int64  `json:"iat"`
	EXP   int64  `json:"exp"`
}

type TokenService interface {
	GenerateToken(payload Payload) (string, error)
}
