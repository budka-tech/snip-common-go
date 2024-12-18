package port

import (
	"fmt"
	"github.com/budka-tech/envo"
)

type Port = uint16

const (
	Redis      Port = 6379
	UsersGrpc       = 10100
	GsmGrpc         = 10200
	GsmHttp         = 10201
	PushGrpc        = 10300
	AssistGrpc      = 10400
	S3Grpc          = 10500
)

type ServiceName string

const (
	ServiceUsers  ServiceName = "snip-users-service"
	ServiceRedis              = "snip-redis-service"
	ServiceGsm                = "snip-gsm-service"
	ServicePush               = "snip-push-service"
	ServiceAssist             = "snip-assist-service"
	ServiceS3G                = "snip-s3-service"
)

var m = map[ServiceName]Port{
	ServiceUsers:  UsersGrpc,
	ServiceRedis:  Redis,
	ServiceGsm:    GsmGrpc,
	ServicePush:   PushGrpc,
	ServiceAssist: AssistGrpc,
	ServiceS3G:    S3Grpc,
}

func Format(port Port) string {
	return fmt.Sprintf("%d", port)
}

func FormatTCP(port Port) string {
	return fmt.Sprintf(":%d", port)
}

func FormatServiceTCP(env *envo.Env, name ServiceName) string {
	port := m[name]
	if env.IsLocal() {
		return FormatLocal(port)
	} else {
		return fmt.Sprintf("%v:%d", name, port)
	}
}

func FormatLocal(port Port) string {
	return fmt.Sprintf("localhost:%d", port)
}
