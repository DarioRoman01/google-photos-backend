package models

// User represents a user account in the system.
type User struct {
	ID       int    `json:"id"`       // ID is unique identifier for the user.
	Username string `json:"username"` // Username is the user's username.
	Email    string `json:"email"`    // Email is the user's email address.
	Password string `json:"-"`        // Password is the user's password, hashed.
}

// UserLoginRegisters represents a user login or registration request.
type UserLoginRegister struct {
	Username string `json:"username"` // Username is the user's username.
	Email    string `json:"email"`    // Email is the user's email address.
	Password string `json:"password"` // Password is the user's password, plaintext.
}

// Image represents an image in the bucket.
type Image struct {
	ID     int    `json:"id"`      // ID is unique identifier for the image.
	Name   string `json:"name"`    // Name is the image's name.
	URL    string `json:"url"`     // URL is the image's URL.
	UserID int    `json:"user_id"` // UserID is the ID of the user who uploaded the image.
}

// Folder represents a folder in the bucket.
type Folder struct {
	ID     int    `json:"id"`      // ID is unique identifier for the folder.
	Name   string `json:"name"`    // Name is the folder's name.
	UserID int    `json:"user_id"` // UserID is the ID of the user who uploaded the image.
}
