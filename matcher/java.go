package token


func StartMultiCommentMatcher() *Matcher {
	kind        = "start_multi_comment"
	allowed_env = []string{"STRING", "CHAR", "SPACE"}
	allowed_prior_tokens = []string{""}

	current_matcher_exp   = "^\\/\\*$"
	following_matcher_exp = "^.$"

	m := NewMatcher(kind, allowed_env, allowed_prior_tokens, current_matcher_exp,
									following_matcher_exp)
	return &Matcher{}
}

func EndMultiCommentMatcher() *Matcher {
	kind        = "end_multi_comment"
	allowed_env = []string{"MULTI COMMENT"}
	allowed_prior_tokens = []string{""}

	current_matcher_exp   = "\\*\\/$"
	following_matcher_exp = "^.$"

	m := NewMatcher(kind, allowed_env, allowed_prior_tokens, current_matcher_exp,
									following_matcher_exp)
	return &Matcher{}
}