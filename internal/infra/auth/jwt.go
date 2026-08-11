package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/munnaMia/nidaa/internal/domain"
	"github.com/munnaMia/nidaa/util/conv"
)

type JWTService struct {
	secretKey string
}

// create a jwt service. take a secret key as parameter
func NewJWTService(scr string) domain.TokenService {
	return &JWTService{
		secretKey: scr,
	}
}

// create a jwt secret key and return error if anything gose wrong
func (jwtSvr *JWTService) GenerateToken(pl domain.Payload) (string, error) {
	header := domain.Header{
		ALG: "HS256",
		TYP: "JWT",
	}

	now := time.Now().Unix()

	if pl.IAT == 0 {
		pl.IAT = now
	}

	if pl.EXP == 0 {
		pl.EXP = now + 86400 // 24 hour of exp time.
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadBytes, err := json.Marshal(pl)
	if err != nil {
		return "", err
	}

	secretBytes := []byte(jwtSvr.secretKey)
	headerB64 := conv.B64UrlEncoding(headerBytes)
	payloadB64 := conv.B64UrlEncoding(payloadBytes)

	massageBytes := []byte(headerB64 + "." + payloadB64)

	hash := hmac.New(sha256.New, secretBytes)
	hash.Write(massageBytes)

	signature := hash.Sum(nil)

	signatureB64 := conv.B64UrlEncoding(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil
}

// validate a jwt token
func (jwtSvr *JWTService) ValidateToken(token string) (*domain.Payload, error) {
	tokenArr := strings.Split(token, ".")
	if len(tokenArr) != 3 {
		return nil, fmt.Errorf("invalide token format")
	}

	jwtHeader := tokenArr[0]
	jwtBody := tokenArr[1]
	jwtSignature := tokenArr[2]

	massage := jwtHeader + "." + jwtBody
	h := hmac.New(sha256.New, []byte(jwtSvr.secretKey))
	h.Write([]byte(massage))
	newHash := h.Sum(nil)
	newSignature := conv.B64UrlEncoding(newHash)

	// Securely compare signatures (prevents timing attacks)
	if subtle.ConstantTimeCompare([]byte(jwtSignature), []byte(newSignature)) != 1 {
		return nil, fmt.Errorf("invalid token signature")
	}

	jwtBodyBytes, err := conv.B64UrlDecoding(jwtBody)
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	var payload domain.Payload
	if err := json.Unmarshal(jwtBodyBytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse payload json: %w", err)
	}

	// Validate expiration timestamp
	if time.Now().Unix() > payload.EXP {
		return nil, fmt.Errorf("token has expired")
	}

	return &payload, nil
}
