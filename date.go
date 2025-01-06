package lib

import "time"

func DateTime() string {
	now := time.Now()

	return now.Format("2006-01-31 15:04:05")
}

func Date() string {
	now := time.Now()
	return now.Format("2006-01-31")
}

func Time() string {
	now := time.Now()
	return now.Format("15:04:05")
}

func TimeShort() string {
	now := time.Now()
	return now.Format("15:04")
}
