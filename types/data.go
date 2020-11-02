package types

type Data struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Code        string `json:"Code"`
	ProducedAt  string `json:"ProducedAt"`
	CreatedAt   string `json:"CreatedAt"`
}

type ParamData struct {
	Code string `json:"Code"`
}
