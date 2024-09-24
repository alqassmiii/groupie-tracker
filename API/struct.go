package API


type Artist struct {
	Id           int     `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int     `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
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
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Data struct {
	A Artist
	R Relation
	L Location
	D Dates
}

type ISPI struct {
	Artests   string `json:"artists"`
	Dates     string `json:"dates"`
	Locations string `json:"locations"`
	Relations string `json:"relations"`
}
