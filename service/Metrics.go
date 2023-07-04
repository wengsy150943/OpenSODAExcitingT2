package service

var Special_Metric = map[string]bool{
	"issue_response_time":                true,
	"issue_resolution_duration":          true,
	"change_request_response_time":       true,
	"change_request_resolution_duration": true,
	"change_request_age":                 true,
}

var Metrics = [...]string{
	"openrank",
	"activity",
	"attention",
	"active_dates_and_times",
	"stars",
	"technical_fork",
	"participants",
	"new_contributors",
	"new_contributors_detail",
	"inactive_contributors",
	"bus_factor",
	"bus_factor_detail",
	"issues_new",
	"issues_closed",
	"issue_comments",
	"issue_response_time",
	"issue_resolution_duration",
	"issue_age",
	"code_change_lines_add",
	"code_change_lines_remove",
	"code_change_lines_sum",
	"change_requests",
	"change_requests_accepted",
	"change_requests_reviews",
	"change_request_response_time",
	"change_request_resolution_duration",
	"change_request_age",
	"activity_details",
}

const MetricNum = len(Metrics)
