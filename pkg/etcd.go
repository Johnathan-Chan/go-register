package pkg

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go-register/utils"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type Etcd struct {
	Client *clientv3.Client
	ID     clientv3.LeaseID
	Lease  clientv3.Lease
	Config *EtcdConfig
	Instance map[string][]string
}

func NewEtcd(read ReadConfig) *Etcd {
	data := read.Read()
	config := &EtcdConfig{}
	err := utils.DataToConfig(&data, config)
	if err != nil {
		return nil
	}

	return &Etcd{
		Config: config,
		Instance: make(map[string][]string),
	}
}

func (this *Etcd) Register() error {
	cfg := clientv3.Config{
		Endpoints:   []string{this.Config.Address},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return err
	}
	this.Client = client

	key := fmt.Sprintf("%s/%s", this.Config.ServerName, uuid.NewV4().String())
	value := fmt.Sprintf("%s:%d", this.Config.ServerUrl, this.Config.Port)
	this.Lease = clientv3.NewLease(this.Client)
	leaseResp, err := this.Lease.Create(context.TODO(), this.Config.Ttl)
	if err != nil {
		return err
	}
	this.ID = clientv3.LeaseID(leaseResp.ID)
	_, err = this.Client.KV.Put(context.TODO(), key, value, clientv3.WithLease(this.ID))
	if err != nil {
		return err
	}

	leaseKeepAliveChan, err := this.Lease.KeepAlive(context.TODO(), this.ID)
	if err != nil {
		return err
	}

	go this.Healthy(leaseKeepAliveChan)

	return nil
}

func (this *Etcd) Healthy(lkap <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case leaseKeepAliveResponse := <-lkap:
			if leaseKeepAliveResponse != nil {
				fmt.Println("续租成功,leaseID :", leaseKeepAliveResponse.ID)
			} else {
				fmt.Println("续租失败")

			}

		}
		time.Sleep(time.Second * 1)
	}
}

func (this *Etcd) Discovery(serName string) ([]string, error) {
	getResponse, err := this.Client.KV.Get(context.TODO(), serName, clientv3.WithPrefix(), clientv3.WithLease(this.ID))
	if err != nil {
		return nil, err
	}

	for k, v := range getResponse.Kvs{
		fmt.Println(k, v)
	}

	return nil, nil
}

func (this *Etcd) Destory() error {
	_, err := this.Lease.Revoke(context.TODO(), this.ID)
	if err != nil {
		return err
	}
	return nil
}
