package models

type User struct {
    UserID       int    `json:"userID"`
    Username     string `json:"username"`
    Email        string `json:"email"`
    PasswordHash string `json:"-"`
    // Other user-related fields
}
