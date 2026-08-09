package validate

import "github.com/munnaMia/nidaa/util/responder"

type PasswordRules struct {
	MaxLength      int
	MinLenght      int
	RequireUpper   bool
	RequireLower   bool
	RequireNumber  bool
	RequireSpecial bool
}

type Validator interface {
	String(max int, min int, field, s string) error                    // validate a string
	Email(e string) error                                              // validate a email based on regexp
	Password(p string, pRules PasswordRules) []responder.ValidationErr // ValidatePassword returns a list of missing requirements, or nil if valid.
}
