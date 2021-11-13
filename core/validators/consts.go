package validators

const (
	emailValidateRegex = `\b([A-Za-z0-9_%+-]+[A-Za-z0-9_'\.%+-]*?@[A-Za-z0-9\.-]+\.[A-Za-z]{2,3})\b`

	defaultLocale = "en"

	emailAddressFnName = "EmailAddress"
	maxFnName          = "max"
	minFnName          = "min"
	requiredFnName     = "required"
)
