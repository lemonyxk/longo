/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:34
**/

package mongo

import (
	"context"
	"strings"
	"time"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	config Config
}

func (c *Client) SetReadPreference(readPreference string) {
	c.config.ReadPreference = NewReadPreference(readPreference)
}

func (c *Client) SetReadConcern(readConcern string) {
	c.config.ReadConcern = NewReadConcern(readConcern)
}

func (c *Client) SetWriteConcern(writeConcern WriteConcern) {
	c.config.WriteConcern = NewWriteConcern(writeConcern)
}

func (c *Client) SetUrl(url string) {
	c.config.Url = url
}

func (c *Client) init(config *Config) {
	
	if config != nil {
		c.config = *config
	}
	
	if c.config.Url == "" {
		if len(c.config.Hosts) == 0 {
			c.config.Hosts = []string{"127.0.0.1:27017"}
		}
		var hostsString = strings.Join(c.config.Hosts, ",")
		if c.config.User == "" || c.config.Auth == "" {
			c.config.Url = "mongodb://" + hostsString
		} else {
			c.config.Url = "mongodb://" + c.config.User + ":" + c.config.Auth + "@" + hostsString
		}
	}
	
	if c.config.ReadPreference == nil {
		c.config.ReadPreference = ReadPreference.Primary
	}
	
	if c.config.ReadConcern == nil {
		c.config.ReadConcern = ReadConcern.Local
	}
	
	if c.config.WriteConcern == nil {
		c.config.WriteConcern = NewWriteConcern(WriteConcern{W: 1, J: false, Wtimeout: 3 * time.Second})
	}
}



func (c *Client) Connect(config *Config) (*Mgo, error) {
	
	if config == nil {
		config = &Config{}
	}
	
	c.init(config)
	
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.config.Url), &options.ClientOptions{
		ReadPreference: c.config.ReadPreference, // default is Primary
		ReadConcern:    c.config.ReadConcern,    // default is local
		WriteConcern:   c.config.WriteConcern,   // default is w:1 j:false wTimeout:when w > 1
	})
	
	if err != nil {
		return nil, err
	}
	
	return &Mgo{client: client, config: c.config}, nil
}