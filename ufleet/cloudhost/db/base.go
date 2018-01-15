package db

import (
	"context"
	"github.com/coreos/etcd/client"
	"ufleet/cloudhost/utils"
	"fmt"
)

func Etcdhander() (client.KeysAPI, error){
	cfg := client.Config{
		Endpoints: []string{utils.EtcdEndPoint},
		Transport: client.DefaultTransport,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	kapi := client.NewKeysAPI(c)
	return kapi, nil
}
func EtcdGet(key string) (*client.Response, error) {
	kapi, err := Etcdhander()
	if err != nil {
		fmt.Println("not etcdhander")
		return nil, err
	}
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		fmt.Println("get nil")
		return nil, err
	}
	return resp, nil
}
// EtcdKeyCheck check if key exist in etcd
func EtcdKeyCheck(key string) (bool, error) {
	_, err := EtcdGet(key)
	if err != nil {
		return false, err
	}
	return true, nil
}