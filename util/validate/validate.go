package validate

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Pre-compile the regex pattern once at startup for high performance.
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type Validate struct {
}

func NewValidate() Validator {
	return &Validate{}
}

func (v *Validate) String(max int, min int, s string) error {
	str := strings.TrimSpace(s)
	length := utf8.RuneCountInString(str)

	if length < min {
		return fmt.Errorf("must be at least %d characters long", min)
	}
	if length > max && length > 0 {
		return fmt.Errorf("must be less then %d characters", max)
	}

	return nil
}

func (v *Validate) Email(e string) (error, bool) {
	email := strings.TrimSpace(e)

	if len(email) == 0 || len(email) > 254 {
		return fmt.Errorf("invalid email length: must be between 1 and 254 characters"), false
	}

	// match the regex with given string
	if ok := emailRegex.MatchString(email); !ok {
		return fmt.Errorf("invalid email input."), false
	}

	return nil, true
}

func (v *Validate) Password(p string, pRule PasswordRules) []error {
	issues := make([]error, 0)

	length := utf8.RuneCountInString(p)

	if length > pRule.MaxLength {
		issues = append(issues, fmt.Errorf("must be less then %d characters ", pRule.MaxLength))
	}
	if length < pRule.MinLenght {
		issues = append(issues, fmt.Errorf("must be more then %d characters ", pRule.MinLenght))
	}

	var hasLower, hasUpper, hasNumber, hasSpecial bool

	for _, char := range p {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			hasSpecial = true
		}
	}

	if pRule.RequireLower && !hasLower {
		issues = append(issues,fmt.Errorf( "must contain at least one lowercase letter"))
	}
	if pRule.RequireUpper && !hasUpper {
		issues = append(issues,fmt.Errorf( "must contain at least one uppercase letter"))
	}
	if pRule.RequireNumber && !hasNumber {
		issues = append(issues,fmt.Errorf( "must contain at least one number letter"))
	}
	if pRule.RequireSpecial && !hasSpecial {
		issues = append(issues,fmt.Errorf( "must contain at least one special letter"))
	}

	if len(issues) > 0 {
		return issues
	}

	return nil
}
