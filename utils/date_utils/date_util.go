package date_utils

import "time"

const dateLayout = "Mon Jan 2 2006 15:04:05 MST"

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	//now := time.Now().UTC()
	//now := time.Now()
	//user.DateCreated = now.Format("Mon Jan 2 2006 15:04:05 MST")
	return GetNow().Format(dateLayout)
}
