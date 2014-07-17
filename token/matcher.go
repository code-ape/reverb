package token

import "regexp"


type Matcher struct {
	AllowedEnvs        []string

	CurrentMatchExp    string
	PriorMatchExp      bool
	FollowingMatchExp  bool

	MainMatcher      *regexp.Regexp
	PriorMatcher     *regexp.Regexp
	FollowingMatcher *regexp.Regexp

	CurrentString   string
	PriorString     string
	FollowingString string

	CurrentMatch   bool
	PriorMatch     bool
	FollowingMatch bool
}

func NewMatcher(ae []string, cme, pme, fme string) *Matcher {
	m := &Matcher{AllowedEnvs: ae, CurrentMatchExp: cme, PriorMatchExp pme,
									FollowingMatchExp: fme}
	m.CompileRegexp()
	return m
}

func (m *Matcher) Reset(prior_string string) {
	m.CurrentString   = ""
	m.PriorString     = ""
	m.FollowingString = ""
	m.CurrentMatch    = false
	m.PriorMatch      = false
	m.FollowingMatch  = false
	m.UpdatePriorMatch(prior_string)
}


func (m *Matcher) CompileRegexp() {
	m.MainMatcher      = regexp.Compile(m.MatchExp)
	m.PriorMatcher     = regexp.Compile(m.PriorMatchExp)
	m.FollowingMatcher = regexp.Compile(m.FollowingMatcher)
}

func (m *Matcher) Match(c string, env string) bool {
	if !m.AllowedEnv(env) || !m.PriorMatch {
		return false
	}

	if m.CurrentMatch && m.FollowingString == "" {
		temp_s  := m.CurrentString + c
		attempt := m.MainMatcher(temp_s)
		if attempt {
			m.CurrentString = temp_s
		} else {
			m.UpdateFollowingMatch(c)
		}
	} else if m.CurrentMatch && m.FollowingMatcher != "" {
		m.UpdateFollowingMatch(c)
	}

	return m.CheckMatches()
	
}


func (m *Matcher) AllowedEnv(env string) bool {
	for _,e := range m.AllowedEnvs {
		if env == e { return true }
	} else { return false }
}

func (m *Matcher) UpdateFollowingMatch(c string) {
	m.FollowingString += c
	m.FollowingMatch = m.FollowingMatcher(m.FollowingString)
}

func (m *Matcher) UpdatePriorMatch(c string) {
	m.PriorString += c
	m.PriorMatch = m.PriorMatcher(m.PriorString)
}

func (m *Matcher) CheckMatches(m *Matcher) {
	return m.CurrentMatch && m.PriorMatch && m.FollowingMatch
}

