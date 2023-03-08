package commonframe

import (
	"fmt"
)

type GetinfoStruct struct {
	OpName string
}

func (info *GetinfoStruct) GetfuncInfo() {
	fmt.Println("Function Name is: ", info.OpName)
}
