package structs

type SkyblockNews struct {
	Success bool `json:"success"`
	Items   []struct {
		Item struct {
			Material string `json:"material"`
		} `json:"item"`
		Link  string `json:"link"`
		Text  string `json:"text"`
		Title string `json:"title"`
	} `json:"items"`
}
