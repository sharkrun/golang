package utils

var EtcdEndPoint string

// InitRunOption init run config
func InitRunOption(etcd string) {
	EtcdEndPoint = etcd
}