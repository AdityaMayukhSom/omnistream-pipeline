package apiConstant

const (
	CookieNameAccessToken  = "x-access-token-pixelated-pipeline"
	CookieNameRefreshToken = "x-refresh-token-pixelated-pipeline"
)

const (
	CookiePathAccessToken  = "/"
	CookiePathRefreshToken = "/api/v1/auth/"
)

const (
	DefaultAuthenticatedRoute   = "/home"
	DefaultUnauthenticatedRoute = "/login"
)

const (
	ContextAttributeKeyUsername = "username"
	ContextAttributeKeyName     = "name"
)

const (
	RESPONSE_TYPE_ALREADY_AUTHENTICATED string = "already_authenticated"
	RESPOSNE_TYPE_NEWLY_AUTHENTICATED   string = "newly_authenticated"
	RESPONSE_TYPE_SUCCESSFUL_LOGOUT     string = "logged_out_successfully"
	RESPONSE_TYPE_SUCCESSFUL_REGISTERED string = "user_successfully_registered"
)

const (
	ERROR_TYPE_INVALID_BODY              string = "could_not_parse_request_body"
	ERROR_TYPE_INVALID_LOGIN_CREDENTIALS string = "invalid_login_credentials"
)
