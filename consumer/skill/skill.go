package skill

type Skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name" default:""`
	Description string   `json:"description" default:""`
	Logo        string   `json:"logo" default:""`
	Tags        []string `json:"tags" default:"{}"`
}
