package main

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "net/url"
	// "os"
	// "text/template"
	"time"
)

// GitHub API 响应数据结构
type Issue struct {
	Number    int
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
}

type Milestone struct {
	Number       int
	Title        string
	Description  string
	OpenIssues   int `json:"open_issues"`
	ClosedIssues int `json:"closed_issues"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// 模板数据
type ReportData struct {
	Repo       string
	Issues     []Issue
	Milestones []Milestone
	Users      map[string]*User
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. 获取仓库名（示例中硬编码为 "golang/go"）
	repo := "golang/go"

	// 2. 查询GitHub API（实际项目需替换为真实请求）
	issues := mockGetIssues(repo)
	milestones := mockGetMilestones(repo)
	users := collectUsers(issues)

	// 3. 准备模板数据
	data := ReportData{
		Repo:       repo,
		Issues:     issues,
		Milestones: milestones,
		Users:      users,
	}

	// 4. 根据URL参数选择输出格式
	format := r.URL.Query().Get("format")
	switch format {
	case "text":
		renderTextReport(w, data)
	case "html":
		renderHTMLReport(w, data)
	default:
		fmt.Fprintf(w, "请指定格式参数: ?format=text 或 ?format=html")
	}
}

// 渲染文本报告
func renderTextReport(w http.ResponseWriter, data ReportData) {
	tmpl := template.Must(template.ParseFiles("templates/report.txt"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 渲染HTML报告
func renderHTMLReport(w http.ResponseWriter, data ReportData) {
    // 定义模板函数
    funcMap := template.FuncMap{
        "add": func(a, b int) int { return a + b },
    }

    // 创建模板并注册函数
    tmpl := template.Must(
        template.New("report.html").
            Funcs(funcMap).
            ParseFiles("templates/report.html"),
    )

    // 执行模板
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// --- 以下为模拟数据函数（实际项目应调用GitHub API）---
func mockGetIssues(repo string) []Issue {
	return []Issue{
		{
			Number:    1,
			Title:     "Bug: panic in net/http",
			State:     "open",
			User:      &User{"alice", "https://github.com/alice"},
			CreatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			Number:    2,
			Title:     "Feature: add JSON support",
			State:     "closed",
			User:      &User{"bob", "https://github.com/bob"},
			CreatedAt: time.Now().Add(-48 * time.Hour),
		},
	}
}

func mockGetMilestones(repo string) []Milestone {
	return []Milestone{
		{
			Number:       1,
			Title:        "Go 1.20",
			Description:  "Next major release",
			OpenIssues:   12,
			ClosedIssues: 45,
		},
	}
}

func collectUsers(issues []Issue) map[string]*User {
	users := make(map[string]*User)
	for _, issue := range issues {
		users[issue.User.Login] = issue.User
	}
	return users
}