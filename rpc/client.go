package rpc

import (
	"flag"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
	AthenaCli  client.XClient
)

func InitRpc() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	AthenaCli = client.NewXClient("AthenaImplement", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	//defer xclient.Close()
}
