package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

type TaskStateRequest struct {
	TaskQueueId int //去第几个队列里取任务
	//TaskQueueTaskId int // 队列里的第几个任务
}

type TaskStateResponse struct {
	State int
	CurrNumMapTask int
	CurrNumReduceTask int
	MapFin	chan bool
	ReduceFin	chan bool
	TaskFin        chan bool
}

type MapTaskRequest struct {

}

type MapTaskResponse struct {
	//State int
	MapTask Task
	NumMapTask int
	NumReduceTask int
}

type ReduceTaskRequest struct {
	TaskQueueId int //去第几个队列里取任务
	//TaskQueueTaskId int // 队列里的第几个任务
}

type ReduceTaskResponse struct {
	//State int
	ReduceTask Task
	NumMapTask int
	NumReduceTask int
	CurrNumMapTask int
	CurrNumReduceTask int
}
// Add your RPC definitions here.


// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the master.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func masterSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
