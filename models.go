package instellar

type Cluster struct {
	Data struct {
		Attributes struct {
			Id           int    `json:"id"`
			CurrentState string `json:"current_state"`
			Name         string `json:"name"`
			Slug         string `json:"slug"`
			Endpoint     string `json:"endpoint"`
			Provider     string `json:"provider"`
			Region       string `json:"region"`
		} `json:"attributes"`
	} `json:"data"`
}
