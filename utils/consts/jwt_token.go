package consts

// Key
const (
	JWT_Authorized string = "authorized"
	JWT_UserId string = "userId"
	JWT_Role string = "role"
	JWT_ExpiredTime string = "exp"
)

// Failed Response
const (
	JWT_InvalidJwtToken       string = "jwt token missing or invalid"
	JWT_FailedCastingJwtToken string = "failed to cast claims as jwt.MapClaims"
)
