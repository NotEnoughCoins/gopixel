package structs

type Key struct {
	Success bool   `json:"success"`
	Cause   string `json:"cause"`
	Record  struct {
		Key              string `json:"key"`
		Owner            string `json:"owner"`
		Limit            int    `json:"limit"`
		QueriesInPastMin int    `json:"queriesInPastMin"`
		TotalQueries     int    `json:"totalQueries"`
	} `json:"record"`
}
