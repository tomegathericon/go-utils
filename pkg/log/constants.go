package log

import "errors"

type LogFormat string

const LOGFMT LogFormat = "logfmt"
const JSON LogFormat = "json"

func GetFormat(format string) (LogFormat, error) {
	switch format {
	case "json", "JSON":
		return JSON, nil
	case "logfmt", "LOGFMT":
		return LOGFMT, nil
	default:
		return "", errors.New("unknown log format")
	}
}
