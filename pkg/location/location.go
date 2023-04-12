package location

// Post created by a user.
type Location struct {
	ID        uint   `json:"idlocation,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Tipo      string `json:"tipo,omitempty"`
	Dimension string `json:"dimension,omitempty"`
}
