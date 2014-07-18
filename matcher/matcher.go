package token

import "regexp"


type Matcher struct {
	Kind string

	AllowedEnvs        []string
	AllowedPriorTokens []string

	CurrentMatchExp    string
	FollowingMatchExp  bool

	MainMatcher      *regexp.Regexp
	FollowingMatcher *regexp.Regexp

	Environment     string
	PriorKind       string
	CurrentString   string
	FollowingString string

	Active         bool
	EnvMatch       bool
	CurrentMatch   bool
	PriorMatch     bool
	FollowingMatch bool
}

func NewMatcher(k string, ae, apt []string, cme, fme string) *Matcher {
	m := &Matcher{Kind: k, AllowedEnvs: ae, AllowedPriorTokens: apt, CurrentMatchExp: cme,
									FollowingMatchExp: fme}
	m.CompileRegexp()
	return m
}

func (m *Matcher) Reset(prior_token_kind string, env string) {
	m.PriorKind       = prior_token_kind
	m.CurrentString   = ""
	m.FollowingString = ""
	m.Environment     = env

	m.CurrentMatch    = false
	m.FollowingMatch  = false

	m.EnvMatch        = m.AllowedEnv(env)
	m.PriorMatch      = m.AllowedPriorToken(prior_token_kind)
	m.Active          = m.EnvMatch && m.PriorMatch
}


func (m *Matcher) CompileRegexp() {
	m.MainMatcher      = regexp.Compile(m.MatchExp)
	m.FollowingMatcher = regexp.Compile(m.FollowingMatcher)
}

func (m *Matcher) Match(c string, env string) bool {
	if !m.Active {
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

func (m *Matcher) AllowedPriorToken(token_kind string) bool {
	for _,t := range m.AllowedPriorTokens {
		if token_kind == t { return true }
	} else { return false }
}

func (m *Matcher) UpdateFollowingMatch(c string) {
	m.FollowingString += c
	m.FollowingMatch = m.FollowingMatcher(m.FollowingString)
}


func (m *Matcher) CheckMatches(m *Matcher) {
	return m.CurrentMatch && m.FollowingMatch
}

