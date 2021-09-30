package entity

import "time"

type PgStatActivity struct {
	tableName        struct{}  `pg:"pg_stat_activity" json:"-"`
	Datid            int       `json:"datid"`
	Datname          string    `json:"datname"`
	Pid              int       `json:"pid"`
	Leader_pid       int       `json:"leader_pid"`
	Usesysid         int       `json:"usesysid"`
	Usename          string    `json:"usename"`
	Application_name string    `json:"application_name"`
	ClientAddr       string    `json:"client_addr"`
	ClientHostname   string    `json:"client_hostname"`
	ClientPort       int       `json:"client_port"`
	BackendStart     time.Time `json:"backend_start"`
	XactStart        time.Time `json:"xact_start"`
	QueryStart       time.Time `json:"query_start"`
	StateChange      time.Time `json:"state_change"`
	WaitEvent_type   string    `json:"wait_event_type"`
	WaitEvent        string    `json:"wait_event"`
	State            string    `json:"state"`
	BackendXid       string    `json:"backend_xid"`
	BackendXmin      string    `json:"backend_xmin"`
	Query            string    `json:"query"`
	BackendType      string    `json:"backend_type"`
	Runtime          time.Time `json:"runtime"`
}

type PgStatActivityResponse struct {
	Count int              `json:"count"`
	Data  []PgStatActivity `json:"data"`
}
