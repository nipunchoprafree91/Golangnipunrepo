package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

)

func main() {
	st := `{"Hostname":"62b35752b4e5","Domainname":"","User":"seluser","AttachStdin":false,"AttachStdout":false,"AttachStderr":false,"ExposedPorts":{"4444/tcp":{},"5900/tcp":{}},"Tty":false,"OpenStdin":false,"StdinOnce":false,"Env":["SCREEN_WIDTH=1500","SCREEN_HEIGHT=1280","PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin","DEBIAN_FRONTEND=noninteractive","DEBCONF_NONINTERACTIVE_SEEN=true","TZ=UTC","HOME=/home/seluser","LANG_WHICH=en","LANG_WHERE=US","ENCODING=UTF-8","LANGUAGE=en_US.UTF-8","LANG=en_US.UTF-8","SCREEN_DEPTH=24","DISPLAY=:99.0","START_XVFB=true","NODE_MAX_INSTANCES=1","NODE_MAX_SESSION=1","NODE_HOST=0.0.0.0","NODE_PORT=5555","NODE_REGISTER_CYCLE=5000","NODE_POLLING=5000","NODE_UNREGISTER_IF_STILL_DOWN_AFTER=60000","NODE_DOWN_POLLING_LIMIT=2","NODE_APPLICATION_NAME=","GRID_DEBUG=false","DBUS_SESSION_BUS_ADDRESS=/dev/null"],"Cmd":["/opt/bin/entry_point.sh"],"ArgsEscaped":true,"Image":"selenium/standalone-chrome-debug:3.141.59-oxygen","Volumes":null,"WorkingDir":"","Entrypoint":null,"OnBuild":null,"Labels":{"authors":"SeleniumHQ"}}
	`
	mapn := make(map[string]interface{})

	err := json.Unmarshal([]byte(st), &mapn)

	if err != nil {
		panic(err)
	}
	fmt.Println(mapn["ExposedPorts"])

	myst := "new:101010"
	splitstr := strings.Split(myst, ":")
	fmt.Println(splitstr)

	url1 := "helios-sandbox.cohesity.com/v2/data-protect/runs/OBJECT-24881%3A1613412057500779/progress"
	qs := url.QueryEscape(url1)
	fmt.Println(qs)

	u, err := url.Parse(qs)
	if err != nil {
		panic(err)
	}

	fmt.Println("scheme: ", u.Path)
}
