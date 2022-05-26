/*
** description("").
** copyright('open-im,www.open-im.io').
** author("fg,Gordon@tuoyun.net").
** time(2021/9/8 14:35).
 */
package main

import (
	"flag"
	"fmt"
	"open_im_sdk/sdk_struct"
	"open_im_sdk/ws_wrapper/utils"
	"open_im_sdk/ws_wrapper/ws_local_server"
	"os"
	"runtime"
	"sync"
)

func main() {
<<<<<<< HEAD
	APIADDR := "http://172.20.10.5:10000"
	WSADDR := "ws://172.20.10.5:17778"

	sqliteDir := flag.String("sdk_db_dir", "./", "openIMSDK initialization path")
	sdkWsIp := flag.String("sdk_ws_ip", "", "openIM ws listening ip")
	sdkWsPort := flag.Int("sdk_ws_port", 7799, "openIM ws listening port")
	openIMApiAddress := flag.String("api_address", APIADDR, "openIM api address")
	openIMWsAddress := flag.String("ws_address", WSADDR, "openIM ws address")

=======
	var sdkWsPort, openIMApiPort, openIMWsPort *int
	//var openIMWsAddress, openIMApiAddress *string
	APIADDR := "http://43.128.5.63:10000"
	WSADDR := "ws://43.128.5.63:17778"
	sdkWsPort = flag.Int("sdk_ws_port", 30000, "openIM ws listening port")
	//openIMApiAddress = flag.String("openIMApiAddress", "", "openIM api listening port")
	//openIMWsAddress = flag.String("openIMWsAddress", "", "openIM ws listening port")
>>>>>>> a974ea82667b9d6c70ce7ff2122073e4dded5c1d
	flag.Parse()
	sysType := runtime.GOOS
	switch sysType {
	case "darwin":
<<<<<<< HEAD
		ws_local_server.InitServer(&sdk_struct.IMConfig{ApiAddr: *openIMApiAddress,
			WsAddr: *openIMWsAddress, Platform: utils.OSXPlatformID, DataDir: *sqliteDir})
	case "linux":
		ws_local_server.InitServer(&sdk_struct.IMConfig{ApiAddr: *openIMApiAddress,
			WsAddr: *openIMWsAddress, Platform: utils.WebPlatformID, DataDir: *sqliteDir})
	case "windows":
		ws_local_server.InitServer(&sdk_struct.IMConfig{ApiAddr: *openIMApiAddress,
			WsAddr: *openIMWsAddress, Platform: utils.WebPlatformID, DataDir: *sqliteDir})
=======
		ws_local_server.InitServer(&sdk_struct.IMConfig{ApiAddr: APIADDR,
			WsAddr: WSADDR, Platform: utils.WebPlatformID, DataDir: "./"})
	case "linux":
		//sdkDBDir:= flag.String("sdk_db_dir","","openIMSDK initialization path")
		ws_local_server.InitServer(&sdk_struct.IMConfig{ApiAddr: "http://" + utils.ServerIP + ":" + utils.IntToString(*openIMApiPort),
			WsAddr: "ws://" + utils.ServerIP + ":" + utils.IntToString(*openIMWsPort), Platform: utils.WebPlatformID, DataDir: "../db/sdk/"})

	case "windows":
		ws_local_server.InitServer(&sdk_struct.IMConfig{ApiAddr: APIADDR,
			WsAddr: WSADDR, Platform: utils.WindowsPlatformID, DataDir: "./", LogLevel: 6})
>>>>>>> a974ea82667b9d6c70ce7ff2122073e4dded5c1d
	default:
		fmt.Println("this os not support", sysType)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("ws server is starting")
<<<<<<< HEAD
	ws_local_server.WS.OnInit(*sdkWsPort, *sdkWsIp)
	if ws_local_server.WS.Run()!=nil {
		os.Exit(-10);
	}
=======
	ws_local_server.WS.OnInit(*sdkWsPort)
	//go func() {
	//	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	//}()

	ws_local_server.WS.Run()

	fmt.Println("run success")
>>>>>>> a974ea82667b9d6c70ce7ff2122073e4dded5c1d
	wg.Wait()

}
