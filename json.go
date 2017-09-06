package node

type Owner struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
}
type Object struct {
	TotalCount        int     `json:"total_count"`
	IncompleteResults bool    `json:"incomplete_results"`
	Items             []Items `json:"items"`
}
type Items struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Owner           Owner  `json:"owner"`
	Private         bool   `json:"private"`
	HtmlUrl         string `json:"html_url"`
	Description     string `json:"description"`
	Fork            bool   `json:"fork"`
	Url             string `json:"url"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	PushedAt        string `json:"pushed_at"`
	Homepage        string `json:"homepage"`
	Size            int    `json:"size"`
	StargazersCount int    `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
	Language        string `json:"language"`
	ForksCount      int    `json:"forks_count"`
	OpenIssuesCount int    `json:"open_issues_count"`
	MasterBranch    string `json:"master_branch"`
	DefaultBranch   string `json:"default_branch"`
	Score           int    `json:"score"`
}
