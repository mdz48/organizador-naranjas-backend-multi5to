package entities

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Id_jefe  *int   `json:"id_jefe,omitempty"`
}
