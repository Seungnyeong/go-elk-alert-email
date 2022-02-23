package domain

type User struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	IsSuperUser bool   `json:"is_superuser"`
	IsActive    bool   `json:"is_active"`
}
