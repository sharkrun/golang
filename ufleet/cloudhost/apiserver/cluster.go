package apiserver

import (
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"
	"ufleet/cloudhost/core"
	"github.com/emicklei/go-restful/log"
)

// return all clusters list

func GetClusters(w rest.ResponseWriter, r *rest.Request){
	log.Printf("GetClusters...")
	clusters, found := core.Getclusters()
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(ErrorResponse(404, "Not Found!"))
		return
	}
	w.WriteJson(clusters)
}

// GetCluster search a specific cluster use provided clusername and return
func GetCluster(w rest.ResponseWriter, r *rest.Request) {
	clusterName := r.PathParam("clustername")
	cluster, found := core.GetCluster(clusterName)
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(ErrorResponse(404, "Not Found!"))
		return
	}
	w.WriteJson(cluster)
}