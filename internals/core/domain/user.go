// Generate package
package domain

// User struct
type User struct {
	ID       int
	Email    string
	Password string
}

// Function instance from User
func NewPerson(id int, email, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

// Return the struc user
func (u *User) GetEmail() string {
	// Return email
	return u.Email
}
