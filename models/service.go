package models

type Service struct {
	Id                  string `json:"id" yaml:"id" sql:"primary_key;column:id"`
	GroupId             int    `json:"group_id" yaml:"group_id" sql:"column:group_id"`
	Name                string `json:"name" yaml:"name" sql:"column:name"`
	Description         string `json:"description" yaml:"description" sql:"column:description"`
	IsEnabled           bool   `json:"is_enabled" yaml:"-" sql:"column:is_enabled"`
	IsOnline            bool   `json:"is_online" yaml:"-" sql:"-"`
	IsPublic            bool   `json:"is_public" yaml:"is_public" sql:"column:is_public"`
	Domain              string `json:"domain" yaml:"domain" sql:"column:domain"`
	Port                int    `json:"port" yaml:"port" sql:"column:port"`
	Interval            int    `json:"interval" yaml:"interval" sql:"column:interval"`
	Timeout             int    `json:"timeout" yaml:"timeout" sql:"column:timeout"`
	Type                string `json:"type" yaml:"type" sql:"column:type"`
	Method              string `json:"method" yaml:"method" sql:"column:method"`
	FollowRedirects     bool   `json:"follow_redirects" yaml:"follow_redirects" sql:"column:follow_redirects"`
	VerifySSL           bool   `json:"verify_ssl" yaml:"verify_ssl" sql:"column:verify_ssl"`
	CreatedAt           int64  `json:"created_at" yaml:"-" sql:"column:created_at"`
	UpdatedAt           int64  `json:"updated_at" yaml:"-" sql:"column:updated_at"`
	Latency             int64  `json:"latency" yaml:"-" sql:"-"`
	Ping                int64  `json:"ping" yaml:"-" sql:"-"`
	AvgResponse         int64  `json:"avg_response" yaml:"-" sql:"-"`
	LastCheckAt         int64  `json:"last_check" yaml:"-" sql:"column:last_check_at"`
	LastLatency         int64  `json:"-" yaml:"-" sql:"-"`
	LastResponse        string `json:"response" yaml:"-" sql:"-"`
	LastStatusCode      int    `json:"status_code" yaml:"-" sql:"-"`
	FailuresLast24Hours int    `json:"failures_24_hours" yaml:"-" sql:"-"`
	AllowNotifications  bool   `json:"allow_notifications" yaml:"allow_notifications" sql:"column:allow_notifications"`
	NotifyAfter         int64  `json:"notify_after" yaml:"notify_after" sql:"column:notify_after"`
	NotifyAllChanges    bool   `json:"notify_all_changes" yaml:"notify_all_changes" sql:"column:notify_all_changes"`
}
