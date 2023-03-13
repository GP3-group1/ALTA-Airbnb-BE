package consts

// Response Error
const (
	// Login
	USER_UserNotFound  string = "user not found"
	USER_WrongPassword string = "wrong password"

	// Register
	USER_EmailAlreadyUsed string = "email is already used"
)

// Bind Error
const (
	USER_ErrorBindUserData string = "error bind user data"
)

// Response Success
const (
	// Login
	USER_LoginSuccess string = "login succeed"

	// Register
	USER_RegisterSuccess string = "succesfully insert user data"

	// Select
	USER_SuccessReadUserData string = "succesfully read user data"

	// Update
	USER_SuccessUpdateUserData string = "succesfully update user data"

	// Delete
	USER_SuccessDelete string = "succesfully delete user"
)

// Validation
const (
	USER_EmptyCredentialError string = "email and password must be filled"
)
