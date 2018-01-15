package worker

import (
	"strings"
	"ufleet/cluster/interf"
	"ufleet/cluster/db"
	"log"
)

// GetCluster get a specfic cluster by clustername from etcd
func GetCluster(clusterName string) (map[string]string, bool) {
	v := map[string]string{}
	if !CheckCluster(clusterName) {
		return v, false
	}
	key := strings.Join([]string{interf.ClusterKey, clusterName}, interf.Sep)
	res, err := db.Get(key)
	if err != nil {
		return v, false
	}

	for key := range res {
		etcdPath := strings.Split(key, "/")
		etcdKey := etcdPath[len(etcdPath)-1]
		v[etcdKey] = res[key]
	}
	return v, true
}

func Getclusters() ([]interface{}, bool) {
	var v1 []interface{}
	values, err := db.Get(interf.ClusterKey)
	if err != nil {
		return v1, false
	}
	log.Println("values:", values)
	log.Println("v0:", v1)
	for k, v := range values {
		sp := strings.Split(k, "/")
		sl := len(sp)
		switch sp[sl -1] {
		case "cluster_info":
			v1 = append(v1, v)
		default:
			break

		}
	}
	return v1, true
}

// CheckCluster check if cluster exist in etcd
func CheckCluster(clusterName string) bool {
	key := strings.Join([]string{interf.ClusterKey, clusterName}, interf.Sep)
	exist, _ := db.IsExist(key)
	return exist
}

// CheckMaster check if master node exist in cluster
func CheckMaster(clusterName, hostip string) bool {
	key := strings.Join([]string{interf.ClusterKey, clusterName, interf.MasterKey, hostip}, interf.Sep)
	exist, _ := db.IsExist(key)
	return exist
}
