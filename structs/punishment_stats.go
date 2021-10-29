package structs

type PunishmentStats struct {
	Success              bool `json:"success"`
	WatchdogLastMinute   int  `json:"watchdog_lastMinute"`
	StaffRollingDaily    int  `json:"staff_rollingDaily"`
	WatchdogTotal        int  `json:"watchdog_total"`
	WatchdogRollingDaily int  `json:"watchdog_rollingDaily"`
	StaffTotal           int  `json:"staff_total"`
}
