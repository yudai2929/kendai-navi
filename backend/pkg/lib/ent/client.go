package ent

import (
	_ "github.com/lib/pq"
	"github.com/yudai2929/kendai-navi/backend/db/ent"
	"github.com/yudai2929/kendai-navi/backend/pkg/lib/errors"
)

type Client struct {
	*ent.Client
}

type SSLMode string

var (
	SLModeDisable SSLMode = "disable"
	SLModeRequire SSLMode = "require"
	SLModeVerify  SSLMode = "verify-ca"
)

type ClientParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  SSLMode
}

func NewClient(p ClientParams) (*Client, error) {
	dataSource := "host=" + p.Host + " port=" + p.Port + " user=" + p.User + " dbname=" + p.Database + " password=" + p.Password + " sslmode=" + string(p.SSLMode)
	client, err := ent.Open("postgres", dataSource)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return &Client{client}, nil
}
