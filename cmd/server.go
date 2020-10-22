package main

import (
	"flag"
	"fmt"
	"github.com/dbaratey/otus_go_hw/sysmon"
	"github.com/ghodss/yaml"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"time"
)

var (
	port    = flag.Int("p", 10000, "The server port")
	avgTime = flag.Int("at", 10, "The AVG time")
	delay   = flag.Int("d", 3, "The delay")
	c       = sysmon.NewCore(10)
	confPth = flag.String("conf", "conf.yml", "Config file")
	conf = configuration{}
)

type configuration struct {
	Metriks []string `yaml:"metriks"`
}

func main() {
	flag.Parse()


	// Config
	buf, err := ioutil.ReadFile(*confPth)
	if err != nil {
		log.Fatal("Can't read conf file ", err.Error())
	}


	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		log.Fatal("Неверный конфигурационный файл")
	}

	for _,m := range conf.Metriks{
		switch m {
		case "cpu":
			w := sysmon.CpuWatcher{}
			go c.Append(&w)
		}
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gServer := grpc.NewServer([]grpc.ServerOption{}...)
	sysmon.RegisterSysmonAgentServer(gServer, &agentServer{})
	gServer.Serve(lis)
}

type agentServer struct {
	sysmon.UnimplementedSysmonAgentServer
}

func (agentServer) GetStats(req *sysmon.NullReq, server sysmon.SysmonAgent_GetStatsServer) error {
	t := time.Tick(time.Second * time.Duration(*delay))
	for {
		_ = <-t
		for _,m := range conf.Metriks{
			switch m {
			case "cpu":
				err := server.Send(&sysmon.SysMonInfo{
					AvgTime:              int32(*avgTime),
					Key:                  "cpu:user",
					Val:                  c.Avg("cpu:user"),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
				if err != nil {
					return err
				}
				//
				err = server.Send(&sysmon.SysMonInfo{
					AvgTime:              int32(*avgTime),
					Key:                  "cpu:idle",
					Val:                  c.Avg("cpu:idle"),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
				if err != nil {
					return err
				}
				err = server.Send(&sysmon.SysMonInfo{
					AvgTime:              int32(*avgTime),
					Key:                  "cpu:system",
					Val:                  c.Avg("cpu:system"),
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				})
				if err != nil {
					return err
				}
			}
		}
	}
}
