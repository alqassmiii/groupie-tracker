package server

type Artest struct {
	// Define the fields of the Artest struct
	ID           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Members      string `json:"members"`
	CreationDate string `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
}

type LocationResp struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Location struct {
	locations []string `json:"locations"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	Relation []string `json:"relation"`
}
