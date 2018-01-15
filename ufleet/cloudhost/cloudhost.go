package main

import (
	"log"
	"flag"
	"strconv"
	"os"
	"net/http"
	"ufleet/cloudhost/config"
	"path"
	"github.com/ant0ine/go-json-rest/rest"
	"ufleet/cloudhost/utils"
	"ufleet/cloudhost/apiserver"
)

func main() {
	etcdServer := flag.String("etcd", "http://" + config.ETCDHOST + ":" + strconv.Itoa(config.ETCDPORT), "etcd endpoint")
	serverPort := flag.String("server_port", strconv.Itoa(config.SERVER_PORT), "server port")

	// set log
	os.MkdirAll(path.Dir(config.LOGDIR), 0777)
	logPath := flag.String("log", config.LOGDIR, "log file path")
	logFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(config.GLogLevel)

	// parse params
	flag.Parse()

	// init log config and etcd server.
	utils.InitRunOption(*etcdServer)

	log.Printf("start apiserver...")

	// register route
	restapi := rest.NewApi()
	restapi.Use(rest.DefaultDevStack...)

	//restapi.Use(&rest.IfMiddleware{
	//	Condition: func(request *rest.Request) bool {
	//		return request.URL.Path != "/login"
	//	},

	//	IfTrue: jwt_middleware,
	//})

	router, err := rest.MakeRouter(
		rest.Get(apiserver.ClusterPath, apiserver.GetClusters),
		rest.Get(apiserver.ClusterPath+"/#clustername", apiserver.GetCluster),

		)

	if err != nil {
		log.Printf(err.Error())
	}
	restapi.SetApp(router)

	//restapi := rest.NewApi()
	//api, err := url.RegisterRoute(restapi)
	//if err != nil {
	//	panic(err)
	//	log.Printf(err.Error())
	//}


	log.Printf("Begin listening...")
	log.Fatal(http.ListenAndServe("localhost:" + *serverPort, restapi.MakeHandler()))
}