package instellar

type Cluster struct {
	Data struct {
		Attributes struct {
			ID           int    `json:"id"`
			CurrentState string `json:"current_state"`
			Name         string `json:"name"`
			Slug         string `json:"slug"`
			Endpoint     string `json:"endpoint"`
			Provider     string `json:"provider"`
			Region       string `json:"region"`
		} `json:"attributes"`
	} `json:"data"`
}

type Uplink struct {
	Data struct {
		Attributes struct {
			ID             int     `json:"id"`
			CurrentState   string  `json:"current_state"`
			InstallationID int     `json:"installation_id"`
			ClusterID      int     `json:"cluster_id"`
			ChannelSlug    string  `json:"channel_slug"`
			DatabaseURL    *string `json:"database_url"`
		} `json:"attributes"`
	} `json:"data"`
}
