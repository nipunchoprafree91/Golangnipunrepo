package commonframe

import (
	"fmt"
)

type Functionconstruct interface {
	COnstruct() ([]FunctionExecutor, error)
}

type FunctionExecutor interface {
	Execute() error
	GetOpName() string
	BackgroundExecute(chan bool, chan error)
	// To track
	GetSleepBeforeTestcase() *int64
	SetSleepBeforeTestcase(*int64)

	GetSleepAfterTestcase() *int64
	SetSleepAfterTestcase(*int64)

	GetMaxBackgroundOpIteration() *int64
	SetMaxBackgroundOpIteration(*int64)
}

func Functionexec() {
	fmt.Println("In function executor function inside executor intgerfaces")
}
