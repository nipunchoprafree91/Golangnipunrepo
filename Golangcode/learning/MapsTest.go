package main

import (
	"fmt"
	"encoding/json"
	"time"

)

type obj struct{
	objmain map[string]map[string]interface{}
}
func ret() (interface{}) {
	nest := make(map[int64]map[string]int64)

	var arr []interface{}

    nest[1] = make(map[string]int64)
	nest[1]["st"] = 20
	nest[1]["et"] = 10

	arr = append(arr,nest)
	fmt.Println(nest)
	fmt.Println(arr)

	return arr
}
func main() {
	// maps := make(map[string]string)
	// arrMaps := make([]map[string]string, 0)

	// maps["a"] = "1"
	// maps["b"] = "2"
	// maps["c"] = "3"
	// maps["d"] = "4"

	// fmt.Println(maps)
	// arrMaps = append(arrMaps, maps)
	// fmt.Println(arrMaps)

	// keys := make([]string, 0, len(maps))
	// for k := range maps {
	// 	keys = append(keys, k)
	// }
	// fmt.Println(keys)

	v := ret()
	fmt.Println("Inside Function:",v)

	Data := make(map[string]map[string]interface{})
	mapstr :=  `{ "VMware": {
		"vmware-hosts"              : "MxqCLP0mk",
                "vsphere-dashboard" : "kwtyVaAmz",
                "hosts"             : ["10.14.17.68","10.14.17.69", "10.14.17.70", "10.14.17.71"]
	}
	}`

	_ = json.Unmarshal([]byte(mapstr), &Data)

	fmt.Printf("Data is:%s\n\n",Data)

	for k,v :=range Data{
		fmt.Printf("Adapter : %s\n",k)
		keys := make([]string, 0, len(v))
		for key, _ := range v{

			keys = append(keys,key )

		}
		fmt.Println("All Keys: ",keys)

	}

	fmt.Printf("Time: %v",time.Now().Format("1602700200000"))
}