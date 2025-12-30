package auth

// maker is an interface for managing token
type Maker interface {
	//CreateToken creates a new token for a specific username and duration
	CreateToken(username, uid string) (string, error)

	//VerifyToken checks the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
