package node

type Owner struct {
	Login             string    `orm:"column(login)"  json:"login"`
	Id                int       `orm:"column(id)"  json:"id"`
	AvatarUrl         string    `orm:"column(avatar_url)"  json:"avatar_url"`
	GravatarId        string    `orm:"column(gravatar_id)"  json:"gravatar_id"`
	Url               string    `orm:"column(url)"  json:"url"`
	ReceivedEventsUrl string    `orm:"column(received_events_url)"  json:"received_events_url"`
	Type              string    `orm:"column(type)"  json:"type"`
	CreateTime        time.Time `orm:"column(create_time);type(datetime);auto_now_add"    json:"create_time"`
	UpdateTime        time.Time `orm:"column(update_time);type(datetime);auto_now"    json:"update_time"`
}
type Items struct {
	Id              int       `orm:"column(id)"  json:"id"`
	Name            string    `orm:"column(name)"  json:"name"`
	FullName        string    `orm:"column(full_name)"  json:"full_name"`
	Owner           Owner     `orm:"column(owner)"  json:"owner"`
	Private         bool      `orm:"column(private)"  json:"private"`
	HtmlUrl         string    `orm:"column(html_url)"  json:"html_url"`
	Description     string    `orm:"column(description)"  json:"description"`
	Fork            bool      `orm:"column(fork)"  json:"fork"`
	Url             string    `orm:"column(url)"  json:"url"`
	CreatedAt       string    `orm:"column(created_at)"  json:"created_at"`
	UpdatedAt       string    `orm:"column(updated_at)"  json:"updated_at"`
	PushedAt        string    `orm:"column(pushed_at)"  json:"pushed_at"`
	Homepage        string    `orm:"column(homepage)"  json:"homepage"`
	Size            int       `orm:"column(size)"  json:"size"`
	StargazersCount int       `orm:"column(stargazers_count)"  json:"stargazers_count"`
	WatchersCount   int       `orm:"column(watchers_count)"  json:"watchers_count"`
	Language        string    `orm:"column(language)"  json:"language"`
	ForksCount      int       `orm:"column(forks_count)"  json:"forks_count"`
	OpenIssuesCount int       `orm:"column(open_issues_count)"  json:"open_issues_count"`
	MasterBranch    string    `orm:"column(master_branch)"  json:"master_branch"`
	DefaultBranch   string    `orm:"column(default_branch)"  json:"default_branch"`
	Score           int       `orm:"column(score)"  json:"score"`
	CreateTime      time.Time `orm:"column(create_time);type(datetime);auto_now_add"    json:"create_time"`
	UpdateTime      time.Time `orm:"column(update_time);type(datetime);auto_now"    json:"update_time"`
}
type Object struct {
	TotalCount        int       `orm:"column(total_count)"  json:"total_count"`
	IncompleteResults bool      `orm:"column(incomplete_results)"  json:"incomplete_results"`
	Items             []Items   `orm:"column(items)"  json:"items"`
	CreateTime        time.Time `orm:"column(create_time);type(datetime);auto_now_add"    json:"create_time"`
	UpdateTime        time.Time `orm:"column(update_time);type(datetime);auto_now"    json:"update_time"`
}
