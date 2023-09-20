package middleware

import (
	"context"
	"testing"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoConnect = mongo.Connect
var mongoPing = (*mongo.Client).Ping
var goDotEnvVariable = godotenv.Getenv

func TestInit_a8a462ccb8(t *testing.T) {
	t.Run("Test with valid connection string", func(t *testing.T) {
		goDotEnvVariable = func(v string) string {
			return "localhost"
		}
		mongoConnect = func(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
			return &mongo.Client{}, nil
		}
		mongoPing = func(client *mongo.Client, ctx context.Context, rp *readpref.ReadPref) error {
			return nil
		}
		init()
		t.Log("Test with valid connection string passed")
	})

	t.Run("Test with invalid connection string", func(t *testing.T) {
		goDotEnvVariable = func(v string) string {
			return "invalidhost"
		}
		mongoConnect = func(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
			return nil, mongo.ErrClientDisconnected
		}
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			} else {
				t.Log("Recovered from panic. Test with invalid connection string passed")
			}
		}()
		init()
	})

	t.Run("Test with valid connection string but server not responding", func(t *testing.T) {
		goDotEnvVariable = func(v string) string {
			return "localhost"
		}
		mongoConnect = func(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
			return &mongo.Client{}, nil
		}
		mongoPing = func(client *mongo.Client, ctx context.Context, rp *readpref.ReadPref) error {
			return mongo.ErrClientDisconnected
		}
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			} else {
				t.Log("Recovered from panic. Test with valid connection string but server not responding passed")
			}
		}()
		init()
	})
}
