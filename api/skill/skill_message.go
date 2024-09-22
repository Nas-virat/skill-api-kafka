package skill

type NameUpdateMessage struct {
	Key  string
	Name string
}

type DescriptionUpdateMessage struct {
	Key         string
	Description string
}

type LogoUpdateMessage struct {
	Key  string
	Logo string
}

type TagsUpdateMessage struct {
	Key  string
	Tags []string
}
