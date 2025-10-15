package entity

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users = []User{
	{ID: 1, Username: "john_doe", Email: "johnDoe@example.com", Password: "password123"},
	{ID: 2, Username: "jane_smith", Email: "janeSmith@example.com", Password: "password456"},
	{ID: 3, Username: "alice_jones", Email: "aliceJones@example.com", Password: "password789"},
}
