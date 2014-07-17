package token

func NewJavaEnvironment() *Environment {
	return &Environment{}
}

type StartMultiCommentMatcher struct {
	MatchingString       string
	MatchingLen          int
	AllowedEnvs          []string
	MatchPriorChar       bool
	AllowedPriorChar     []string
	MatchFollowingChar   bool
	AllowedFollowingChar []string

	CurrentString string
	PriorChar     string
	FollowingChar string
}

func NewStartMultiCommentMatcher() *StartMultiCommentMatcher {
	return &StartMultiCommentMatcher{}
}

func (m *StartMultiCommentMatcher) Match(c string, env string) bool {
	if !AllowedEnv(env, m) {
		return false
	}

	AddChar(c, m)




}

func AllowedEnv(env string, m *ForwardSlashMatcher) bool {
	for _,s := range m.AllowedEnvs {
		if env == s {
			return true
		}
	} else {
		return false
	}
}

func AddChar(c string, m *ForwardSlashMatcher) {
	l := len(m.CurrentString)
	if l == m.MatchingLen && m.FollowingChar != "" {
		m.PriorChar := m.CurrentString[0]
		m.CurrentString = m.CurrentString[1:l] + m.FollowingChar
		m.FollowingChar = c
	} else if l == m.MatchingLen && m.FollowingChar == "" {
		m.FollowingChar = c
	} else {
		m.CurrentString = m.CurrentString + c
	}
}