package responses

type InfoReply struct {
	Error    *Error `json:"error"`
	Response Info   `json:"response"`
}

type Info struct {
	Error     *Error `json:"error"`
	TwoFactor int    `json:"2fa_required,omitempty"`
	Country   string `json:"country,omitempty"`
	Secure    int    `json:"https_required,omitempty"`

	Intro             int               `json:"intro,omitempty"`
	CommunityComments bool              `json:"community_comments,omitempty"`
	LinkRedirects     map[string]string `json:"link_redirects,omitempty"`
	Language          int               `json:"lang,omitempty"`
	NoWallReplies     int               `json:"no_wall_replies,omitempty"`
	OwnerPosts        int               `json:"own_posts_default,omitempty"`

	VKPayEndpoint string   `json:"vk_pay_endpoint_v2,omitempty"`
	Translation   []string `json:"messages_translation_language_pairs,omitempty"`
	ObsceneFilter bool     `json:"obscene_text_filter,omitempty"`
}
