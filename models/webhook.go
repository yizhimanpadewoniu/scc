package models

import "time"

// WebHook对象
type WebHook struct {
	WebHookId      int       `orm:"pk;auto;unique;column(web_hook_id)" json:"web_hook_id"`
	RepositoryName string    `orm:"size(255);column(repo_name)" json:"repository_name"`
	BranchName     string    `orm:"size(255);column(branch_name)" json:"branch_name"`
	Tag            string    `orm:"size(1000);column(tag)" json:"tag"`
	Shell          string    `orm:"size(1000);column(shell)" json:"shell"`
	Status         int       `orm:"type(int);column(status);default(0)" json:"status"`
	Key            string    `orm:"size(255);column(key);unique" json:"key"`
	Secure         string    `orm:"size(255);column(secure)" json:"secure"`
	LastExecTime   time.Time `orm:"type(datetime);column(last_exec_time);null" json:"last_exec_time"`
	CreateTime     time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"create_time"`
	HookType       string    `orm:"column(hook_type);size(50)" json:"hook_type"`
	CreateAt       int       `orm:"type(int);column(create_at)"`
}
