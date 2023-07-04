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

type Node struct {
	Data struct {
		Attributes struct {
			ID           int    `json:"id"`
			CurrentState string `json:"current_state"`
			Slug         string `json:"slug"`
			PublicIP     string `json:"public_ip"`
			ClusterID    int    `json:"cluster_id"`
		} `json:"attributes"`
	} `json:"data"`
}

type Storage struct {
	Data struct {
		Attributes struct {
			ID                        int    `json:"id"`
			CurrentState              string `json:"current_state"`
			Host                      string `json:"host"`
			Bucket                    string `json:"bucket"`
			Region                    string `json:"region"`
			CredentialAccessKeyID     string `json:"credential_access_key_id"`
			CredentialSecretAccessKey string `json:"credential_secret_access_key"`
		} `json:"attributes"`
	} `json:"data"`
}

type Component struct {
	Data struct {
		Attributes struct {
			ID           int      `json:"id"`
			CurrentState string   `json:"current_state"`
			Slug         string   `json:"slug"`
			Provider     string   `json:"provider"`
			Version      string   `json:"version"`
			ClusterIDS   []int    `json:"cluster_ids"`
			Channels     []string `json:"channels"`
			Credential   struct {
				Username string `json:"username"`
				Password string `json:"password"`
				Host     string `json:"host"`
				Port     int    `json:"port"`
				Database string `json:"database"`
			} `json:"credential"`
		} `json:"attributes"`
	} `json:"data"`
}
