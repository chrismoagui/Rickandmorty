package character

type Character struct {
	ID      uint   `json:"idcharacter,omitempty"`
	Nombre  string `json:"nombre,omitempty"`
	Estado  string `json:"estado,omitempty"`
	Especie string `json:"especie,omitempty"`
}
