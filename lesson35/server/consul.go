package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
)

// consul 定义一个consul结构体，其内部有一个`*api.Client`字段。
type consul struct {
	client *api.Client
}

// NewConsul 连接至consul服务返回一个consul对象
func NewConsul() (*consul, error) {
	cfg := api.DefaultConfig()
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consul{c}, nil
}

// RegisterService 将gRPC服务注册到consul
func (c *consul) RegisterService(serviceName string, ip string, port int) error {
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID, 每个实例一个ID
		Name:    serviceName,                                    // 服务名称, 同一个服务多实例是相同的
		Tags:    []string{"yk", "hello"},                        // 为服务打标签
		Address: ip,
		Port:    port,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", ip, port), // 外网地址
			Timeout:                        "5s",
			Interval:                       "5s", // 间隔
			DeregisterCriticalServiceAfter: "5m", // 5m后注销掉不健康的服务节点
		},
	}
	return c.client.Agent().ServiceRegister(srv)
}

func (c *consul) Deregister(serviceName string, ip string, port int) error {
	return c.client.Agent().ServiceDeregister(
		fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID, 每个实例一个ID
	)
}

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
