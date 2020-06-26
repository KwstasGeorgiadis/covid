package mcsse

type ResponseCountry struct {
	Country   string      `json:"country"`
	County    interface{} `json:"county"`
	UpdatedAt string      `json:"updatedAt"`
	Stats     struct {
		Confirmed int `json:"confirmed"`
		Deaths    int `json:"deaths"`
		Recovered int `json:"recovered"`
	} `json:"stats"`
	Coordinates struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"coordinates"`
	Province string `json:"province"`
}

//CSSEResponse response that we will return on /api/csse/all
type CSSEResponse struct {
	Data []CSEECountryResponse `json:"data"`
}

//CSEECountryResponse response that we will return on /api/csse/{country}
type CSEECountryResponse struct {
	Country string          `json:"country"`
	Data    []CSEEProvision `json:"data"`
}

type CSEEProvision struct {
	County    string `json:"county"`
	Province  string `json:"province"`
	Cases     int    `json:"cases"`
	Deaths    int    `json:"deaths"`
	Recovered int    `json:"recovered"`
}
