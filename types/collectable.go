package types

type Collectable struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HasUpgrade  bool `json:"-"`
	UpgradesTo  *Collectable `json:"upgradesTo,omitempty"`
}
