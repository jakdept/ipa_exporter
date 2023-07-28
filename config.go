package ipaexporter

import "time"

type Config struct {
	URL      string
	Username *string
	Password *string

	ForceTLS              bool
	InsecureSkipVerifyTLS bool
	ConnectionTimeout     time.Duration
}
