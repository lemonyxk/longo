/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-28 15:34
**/

package longo

import (
	"context"
	"net/url"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	config Config
}

func (c *Client) SetReadPreference(readPreference string) {
	c.config.ReadPreference = NewReadPreference(readPreference)
}

func (c *Client) SetRegister(register *bsoncodec.Registry) {
	c.config.Register = register
}

func (c *Client) SetReadConcern(readConcern string) {
	c.config.ReadConcern = NewReadConcern(readConcern)
}

func (c *Client) SetWriteConcern(writeConcern WriteConcern) {
	c.config.WriteConcern = &writeConcern
}

func (c *Client) SetConnectTimeout(connectTimeout time.Duration) {
	c.config.ConnectTimeout = connectTimeout
}

func (c *Client) SetUrl(url string) {
	c.config.Url = url
}

func (c *Client) init(config *Config) {

	if config == nil {
		return
	}

	c.config = *config

	if c.config.Url == "" {
		if len(c.config.Hosts) == 0 {
			c.config.Hosts = []string{"127.0.0.1:27017"}
		}
		var hostsString = strings.Join(c.config.Hosts, ",")
		if c.config.User == "" || c.config.Pass == "" {
			c.config.Url = "mongodb://" + hostsString
		} else {
			c.config.Url = "mongodb://" + c.config.User + ":" + c.config.Pass + "@" + hostsString
		}
	}

	if c.config.ReadPreference == nil {
		c.config.ReadPreference = ReadPreference.Primary
	}

	if c.config.ReadConcern == nil {
		c.config.ReadConcern = ReadConcern.Local
	}

	if c.config.WriteConcern == nil {
		c.config.WriteConcern = &WriteConcern{W: -1, J: true, WTimeout: 3 * time.Second}
	}

	if c.config.ConnectTimeout == 0 {
		c.config.ConnectTimeout = 3 * time.Second
	}
}

func (c *Client) Connect(config *Config, opts ...*options.ClientOptions) (*Mgo, error) {

	if config == nil {
		config = &Config{}
	}

	c.init(config)

	if !strings.HasSuffix(c.config.Url, "/") {
		c.config.Url = c.config.Url + "/"
	}

	var u, err = url.Parse(c.config.Url)
	if err != nil {
		return nil, err
	}

	if c.config.TLS {
		var q = u.Query()
		q.Set("tls", "true")
		u.RawQuery = q.Encode()
	}

	var option = options.Client().ApplyURI(u.String())

	if config.Register != nil {
		option.SetRegistry(config.Register)
	}

	opts = append(opts, option)

	opts = append(opts, &options.ClientOptions{
		ReadPreference:         c.config.ReadPreference,                 // default is Primary
		ReadConcern:            c.config.ReadConcern,                    // default is local
		WriteConcern:           NewWriteConcern(*c.config.WriteConcern), // default is w:-1 j:false wTimeout:when w > 1
		ConnectTimeout:         &c.config.ConnectTimeout,
		SocketTimeout:          &c.config.ConnectTimeout,
		ServerSelectionTimeout: &c.config.ConnectTimeout,
	})

	client, err := mongo.Connect(context.Background(), opts...)

	if err != nil {
		return nil, err
	}

	return &Mgo{client: client, config: c.config}, nil
}
