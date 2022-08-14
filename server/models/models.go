package models

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt/v4"
)

// TokenType is the type of token
type TokenType int

const (
	// AccessToken is the type of access token
	AccessToken TokenType = iota
	// VerifyToken is the type of verification token
	VerifyToken
	// ChangePasswordToken is the type of change password token
	ChangePasswordToken
)

// String returns the string representation of the token type
func (t TokenType) String() string {
	switch t {
	case AccessToken:
		return "access"
	case VerifyToken:
		return "verification"
	case ChangePasswordToken:
		return "changepassword"
	default:
		return "unknown"
	}
}

// User represents a user account in the system.
type User struct {
	ID         string `json:"id"`          // ID is unique identifier for the user.
	Username   string `json:"username"`    // Username is the user's username.
	Email      string `json:"email"`       // Email is the user's email address.
	Password   string `json:"-"`           // Password is the user's password, hashed.
	IsVerified bool   `json:"is_verified"` // IsVerified is true if the user has verified their email address.
}

// UserLoginRegisters represents a user login or registration request.
type UserLoginRegister struct {
	Username string `json:"username"` // Username is the user's username.
	Email    string `json:"email"`    // Email is the user's email address.
	Password string `json:"password"` // Password is the user's password, plaintext.
}

// Image represents an image in the bucket.
type Image struct {
	ID       string `json:"id"`        // ID is unique identifier for the image.
	Name     string `json:"name"`      // Name is the image's name.
	URL      string `json:"url"`       // URL is the image's URL.
	UserID   string `json:"user_id"`   // UserID is the ID of the user who uploaded the image.
	FolderID string `json:"folder_id"` // FolderID is the ID of the folder the image is in.
}

// Folder represents a folder in the bucket.
type Folder struct {
	ID     string  `json:"id"`      // ID is unique identifier for the folder.
	Name   string  `json:"name"`    // Name is the folder's name.
	UserID string  `json:"user_id"` // UserID is the ID of the user who uploaded the image.
	Images []Image `json:"images"`  // Images is the images in the folder.
}

// Claims represents the claims in a JWT.
type Claims struct {
	Username             string `json:"username"` // Username is the user's username.
	UserID               string `json:"user_id"`  // UserID is the ID of the user who uploaded the image.
	Type                 string `json:"type"`     // Type is the type of token.
	jwt.RegisteredClaims        // RegisteredClaims are the registered claims in a JWT.
}

// UploadRequest represents a request to upload an image.
type UploadRequest struct {
	FolderID   string         // FolderID is the ID of the folder the image is in.
	FolderName string         // FolderName is the name of the folder to upload to.
	Filename   string         // Filename is the name of the file to upload.
	Username   string         // Username is the user's username.
	UserID     string         // UserID is the ID of the user who uploaded the image.
	File       multipart.File // File is the file to upload.
}
