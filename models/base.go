package models

import (
	"errors"
	"spos/auth/constant"
	"time"

	"github.com/s-pos/go-utils/config"
)

var (
	timezone = config.Timezone()
	layout   = "2006-01-02T15:04:05"
)

func convertTimezone(t time.Time) time.Time {
	return t.UTC()
}

func stringToTime(t string) (time.Time, error) {
	timeFormat, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, errors.New(constant.Message[constant.ErrorParseDate])
	}

	return timeFormat.UTC().In(timezone), nil
}

func timeToString(t time.Time) string {
	return convertTimezone(t).Format(layout)
}
