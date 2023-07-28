package ipaexporter

import (
	"crypto/tls"
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

type conn struct {
	c      *ldap.Conn
	authed bool
}

type LdapConnection interface {
}

func TLSConnection(cfg *Config) (*conn, error) {

	var dialOpts []ldap.DialOpt

	if cfg.InsecureSkipVerifyTLS {
		dialOpts = append(dialOpts,
			ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		)
	}

	if cfg.ConnectionTimeout > 0 {
		ldap.DefaultTimeout = cfg.ConnectionTimeout
	}

	ldapConnection, err := ldap.DialURL(cfg.URL, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("connection problem: %w", err)
	}

	conn := &conn{c: ldapConnection}

	return conn, nil
}

func (conn conn) Hangup() error {
	return conn.c.Close()
}
