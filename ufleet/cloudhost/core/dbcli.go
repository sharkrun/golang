package core

import (
	"strings"
	"ufleet/cloudhost/db"
	"ufleet/cloudhost/interf"
)

// GetCluster get a specfic cluster by clustername from etcd
func GetCluster(clusterName string) (map[string]string, bool) {
	v := map[string]string{}
	if !CheckCluster(clusterName) {
		return v, false
	}
	key := strings.Join([]string{interf.ClusterKey, clusterName}, interf.Sep)
	res, err := db.EtcdGet(key)
	if err != nil {
		return v, false
	}

	for _, n := range res.Node.Nodes {
		etcdPath := strings.Split(n.Key, "/")
		etcdKey := etcdPath[len(etcdPath)-1]
		v[etcdKey] = n.Value
	}
	return v, true
}

func Getclusters() (map[string](map[string]string), bool) {
	v := map[string](map[string]string){}
	res, err := db.EtcdGet(interf.ClusterKey)
	if err != nil {
		return v, false
	}
	for _, cluster := range res.Node.Nodes {
		etcdPath := strings.Split(cluster.Key, "/")
		etcdKey := etcdPath[len(etcdPath) - 1]
		v[etcdKey], _ = GetCluster(etcdKey)
	}
	return v, true
}
// CheckCluster check if cluster exist in etcd
func CheckCluster(clusterName string) bool {
	key := strings.Join([]string{interf.ClusterKey, clusterName}, interf.Sep)
	exist, _ := db.EtcdKeyCheck(key)
	return exist
}

