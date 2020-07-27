package customservice

const (
	CommandTyping       = "Typing"
	CommandCancelTyping = "CancelTyping"
)

type KF struct {
	KfAccount string `json:"kf_account"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
}
