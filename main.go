package main

/**/
import "C"
import (
	"context"
	"log"
	"mosquitto_auth_plugin/mosq_err"
	pb "mosquitto_auth_plugin/mosquitto_auth"
	"time"

	"google.golang.org/grpc/connectivity"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientConn *grpc.ClientConn = nil

//export PluginInit
func PluginInit(addr *C.char) C.int {
	if clientConn != nil && clientConn.GetState() != connectivity.Shutdown {
		clientConn.Close()
	}
	var err error

	clientConn, err = grpc.Dial(C.GoString(addr), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil || clientConn == nil {
		log.Fatalf("[PluginInit] did not connect: %v", err)
		return mosq_err.MOSQ_ERR_UNKNOWN
	}

	return mosq_err.MOSQ_ERR_SUCCESS
}

//export PluginBasicAuth
func PluginBasicAuth(username *C.char, password *C.char, clientId *C.char, clientAddress *C.char) C.int {
	Username := C.GoString(username)
	Password := C.GoString(password)
	ClientId := C.GoString(clientId)
	ClientAddress := C.GoString(clientAddress)

	// if clientConn == nil {
	// 	log.Fatalf("[PluginBasicAuth] did not connect , username = %s", Username)
	// 	return mosq_err.MOSQ_ERR_UNKNOWN
	// }

	c := pb.NewGreeterClient(clientConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.BasicAuth(ctx, &pb.BasicAuthRequest{Username: Username, Password: Password, ClientId: ClientId, ClientAddress: ClientAddress})
	if err != nil {
		log.Fatalf("[PluginBasicAuth]could not greet: %v", err)
		return mosq_err.MOSQ_ERR_UNKNOWN
	}

	log.Printf("[PluginBasicAuth] Greeting: %d", r.GetCode())

	return C.int(r.GetCode())
}

//export PluginAclCheck
func PluginAclCheck(username *C.char, clientId *C.char, topic *C.char, access C.int, qos C.int, retain C.int) C.int {
	Username := C.GoString(username)
	ClientId := C.GoString(clientId)
	Topic := C.GoString(topic)
	Access := int32(access)
	Retain := int32(retain)

	// if clientConn == nil {
	// 	log.Fatalf("[PluginAclCheck] did not connect , username = %s", Username)
	// 	return mosq_err.MOSQ_ERR_UNKNOWN
	// }

	c := pb.NewGreeterClient(clientConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AclCheck(ctx, &pb.AclCheckRequest{Username: Username, ClientId: ClientId, Topic: Topic, Access: Access, Retain: Retain})
	if err != nil {
		log.Fatalf("[PluginAclCheck] could not greet: %v", err)
		return mosq_err.MOSQ_ERR_UNKNOWN
	}

	log.Printf("[PluginAclCheck] Greeting: %d", r.GetCode())

	return C.int(r.GetCode())
}

func main() {}
