package Models

type State struct {
	Id    int    `json:"id" `
	State string `json:"State" validate:"required"`
}

func (b *State) TableName() string {
	return "state"
}
