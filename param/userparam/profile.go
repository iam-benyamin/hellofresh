package userparam

type ProfileRequest struct {
	UserID string
}

type ProfileResponse struct {
	ID        string
	FirstName string
	LastName  string
}
