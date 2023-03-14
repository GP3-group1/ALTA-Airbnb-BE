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

	// Select
	USER_SuccessReadBalance string = "succesfully read user's balance data"

	// Update
	USER_SuccessUpdateUserData string = "succesfully update user data"

	// Update Password
	USER_SuccessUpdatePasswordUserData string = "succesfully update user's password"

	// Delete
	USER_SuccessDelete string = "succesfully delete user"
)

// Validation
const (
	USER_EmptyCredentialError     string = "email and password must be filled"
	USER_EmptyUpdatePasswordError string = "old password and new password must be filled"
)
