# Json2Go
nodejs编写，一个可以把 Json 自动转变为golang 结构体的小工具，可自行修改支持各种 tag 如 orm，json 

# 使用
0.准备环境 nodejs，go
1.拷贝json到 `source.json`文件<br>
2.运行json文件 `./json`<br>
3.结果在 `json.go`<br>


## 例子
输入
```
{
    "total_count": 40,
    "incomplete_results": false,
    "items": [
      {
        "id": 3081286,
        "name": "Tetris",
        "full_name": "dtrupenn/Tetris",
        "owner": {
          "login": "dtrupenn",
          "id": 872147,
          "avatar_url": "https://secure.gravatar.com/avatar/e7956084e75f239de85d3a31bc172ace?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-user-420.png",
          "gravatar_id": "",
          "url": "https://api.github.com/users/dtrupenn",
          "received_events_url": "https://api.github.com/users/dtrupenn/received_events",
          "type": "User"
        },
        "private": false,
        "html_url": "https://github.com/dtrupenn/Tetris",
        "description": "A C implementation of Tetris using Pennsim through LC4",
        "fork": false,
        "url": "https://api.github.com/repos/dtrupenn/Tetris",
        "created_at": "2012-01-01T00:31:50Z",
        "updated_at": "2013-01-05T17:58:47Z",
        "pushed_at": "2012-01-01T00:37:02Z",
        "homepage": "",
        "size": 524,
        "stargazers_count": 1,
        "watchers_count": 1,
        "language": "Assembly",
        "forks_count": 0,
        "open_issues_count": 0,
        "master_branch": "master",
        "default_branch": "master",
        "score": 10.309712
      }
    ]
  }
```

输出
```
// json 版
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

// json orm 版
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
```
