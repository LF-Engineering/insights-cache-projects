package projectscache

type ProjectsData struct {
	Projects []Project `json:"projects"`
	Total    int       `json:"total"`
}

type Project struct {
	ID                    string       `json:"id" `
	Name                  string       `json:"name"`
	Description           string       `json:"description"`
	Slug                  string       `json:"slug"`
	ProjectType           string       `json:"projectType"`
	Metadata              Metadata     `json:"metadata"`
	Metrics               Metrics      `json:"metrics"`
	MetricsFloat          MetricsFloat `json:"metrics_floats"`
	ConfiguredDataSources []string
	SubProjects           []SubProject
}

type Metadata struct {
	Tags []struct {
		Type  string `json:"type"`
		Label string `json:"label"`
	} `json:"tags"`
	Brand struct {
		Color  string `json:"color"`
		Assets struct {
			Logo string `json:"logo"`
		} `json:"assets"`
	} `json:"brand"`
}

type Metrics struct {
	Affiliations  string `json:"affiliations"`
	Commits       string `json:"commits"`
	Contributors  string `json:"contributors"`
	Contributions string `json:"contributions"`
	LinesOfCode   string `json:"linesOfCode"`
	Organizations string `json:"organizations"`
	Repositories  string `json:"repositories"`
}

type MetricsFloat struct {
	Affiliations  float64 `json:"affiliations"`
	Commits       float64 `json:"commits"`
	Contributors  float64 `json:"contributors"`
	Contributions float64 `json:"contributions"`
	LinesOfCode   float64 `json:"linesOfCode"`
	Organizations float64 `json:"organizations"`
	Repositories  float64 `json:"repositories"`
}

type SubProject struct {
	ID       string   `json:"id"`
	Metadata Metadata `json:"metadata"`
	Name     string   `json:"name"`
	Slug     string   `json:"slug"`
}
