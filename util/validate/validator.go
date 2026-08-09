package validate

type PasswordRules struct {
	MaxLength      int
	MinLenght      int
	RequireUpper   bool
	RequireLower   bool
	RequireNumber  bool
	RequireSpecial bool
}

type Validator interface {
	String(max int, min int, s string) error         // validate a string
	Email(e string) (error, bool)                    // validate a email based on regexp
	Password(p string, pRules PasswordRules) []error // ValidatePassword returns a list of missing requirements, or nil if valid.
}
