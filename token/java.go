package token


func NewMultiCommentStartMatcher() *Matcher {
	allowed_env = []string{"STRING", "CHAR"}
	current_matcher_exp   = ""
	prior_matcher_exp     = ""
	following_matcher_exp = ""

	m := NewMatcher(allowed_env, current_matcher_exp, prior_matcher_exp,
									following_matcher_exp)
	return &Matcher{}
}