package instellar

type Cluster struct {
	Data struct {
		Attributes struct {
			Id           int    `json:"id"`
			CurrentState string `json:"current_state"`
			Slug         string `json:"slug"`
		} `json:"attributes"`
	} `json:"data"`
}
