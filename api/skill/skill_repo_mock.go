package skill

type mockRepo struct {
	SkillRepo
	err    error
	skill  Skill
	skills []Skill
}

func (m *mockRepo) GetSkillByKey(key string) (*Skill, error) {
	return &m.skill, m.err
}
func (m *mockRepo) GetSkills() ([]Skill, error) {
	return m.skills, m.err
}
func (m *mockRepo) CreateSkill(skill Skill) (*Skill, error) {
	return &m.skill, m.err
}
func (m *mockRepo) UpdateSkill(key string, skill Skill) (*Skill, error) {
	return &m.skill, m.err
}
func (m *mockRepo) UpdateSkillNameByKey(key string, name string) (*Skill, error) {
	return &m.skill, m.err
}
func (m *mockRepo) UpdateSkillDescriptionByKey(key string, description string) (*Skill, error) {
	return &m.skill, m.err
}
func (m *mockRepo) UpdateSkillLogoByKey(key string, logo string) (*Skill, error) {
	return &m.skill, m.err
}
func (m *mockRepo) UpdateSkillTagsByKey(key string, tags []string) (*Skill, error) {
	return &m.skill, m.err
}
func (m *mockRepo) DeleteSkillByKey(key string) error {
	return m.err
}
