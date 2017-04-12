package main

import (
	"fmt"
	"time"

	mc "git.intra.weibo.com/passport/gopkg/memcache"
	"github.com/piaoqingbin/gomemcache/memcache"
	fhttp "github.com/valyala/fasthttp"
	"zeus/app"
)

var (
	cache_idc1_hosts = []string{"127.0.0.1:11211", "127.0.0.1:11212"}
	cache_idc2_hosts = []string{"127.0.0.1:11213", "127.0.0.1:11214"}
	cache_idc3_hosts = []string{"127.0.0.1:11215", "127.0.0.1:11216"}
)
var (
	idc1_slave = map[string][]string{"idc2": cache_idc2_hosts, "idc3": cache_idc3_hosts}
	idc2_slave = map[string][]string{"idc1": cache_idc1_hosts, "idc3": cache_idc3_hosts}
	idc3_slave = map[string][]string{"idc1": cache_idc1_hosts, "idc2": cache_idc3_hosts}
)

func set(client *mc.ClusterClient, key string, value string, ttl int32) (error) {
	item := memcache.Item{Key: key, Value: []byte(value), Flags: 0, Expiration: ttl, }
	err := client.Set(&item)
	return err
}

func main() {
	fhttp.ListenAndServe(":8089", app.Route.HandleRequest)
	idc1_client := createMcCluster(cache_idc1_hosts, idc1_slave)
	idc2_client := createMcCluster(cache_idc2_hosts, idc2_slave)
	idc3_client := createMcCluster(cache_idc3_hosts, idc3_slave)

	err := set(idc1_client, "abc_h1", "abc=123", 90)
	set(idc1_client, "abc_a2", "abc=123", 90)
	set(idc1_client, "abc_a3", "abc=123", 90)
	set(idc1_client, "abc_a4", "abc=123", 90)
	set(idc1_client, "abc_a5", "abc=123", 90)
	set(idc1_client, "abc_a6", "abc=123", 90)
	set(idc2_client, "def", "def=123", 90)
	set(idc2_client, "def_a2", "def=123", 90)
	set(idc3_client, "ghi", "ghi=123", 90)
	set(idc2_client, "ghi_a2", "def=123", 90)
	item, _ := idc1_client.Get("def")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println("get", item.Key, string(item.Value))
	item, _ = idc1_client.Get("def_a2")
	fmt.Println("get", item.Key, string(item.Value))
	item, _ = idc1_client.Get("ghi")
	fmt.Println("get", item.Key, string(item.Value))
	item, _ = idc1_client.Get("ghi_a2")
	fmt.Println("get", item.Key, string(item.Value))
	time.Sleep(10 * time.Second)
}

func createMcCluster(master []string, slave map[string][]string) (client *mc.ClusterClient) {
	cfg := mc.DefaultClusterConfig()
	cfg.Master = mc.NewClusterGroup()
	cfg.Master.Servers = master

	for idc_idx, idc := range slave {
		cfg.Slave[idc_idx] = mc.NewClusterGroup()
		cfg.Slave[idc_idx].Servers = idc
	}

	client = mc.NewClusterClient(cfg)
	return
}
