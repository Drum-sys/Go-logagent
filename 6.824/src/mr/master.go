package mr

import (
	"fmt"
	"log"
	"sync"
)
import "net"
import "os"
import "net/rpc"
import "net/http"

type Task struct {
	MapId    int
	ReduceId int
	Filename string
}

type Master struct {
	State         int // 0 表示map， 1表示reduce， 2表示finish
	NumMapTask    int
	NumReduceTask int
	MapTask       chan Task
	ReduceTask    chan Task
	MapTaskFin    chan bool
	ReduceTaskFin chan bool
	SLock         sync.Mutex
	TaskFin        chan bool
	MapFin	chan bool
	ReduceFin chan bool
	CurrNumMapTask int
	CurrNumReduceTask int
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Master) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Master) GetTaskState(args *TaskStateRequest, reply *TaskStateResponse) error {
	//c.SLock.Lock()
	if c.CurrNumMapTask > 0 {
		//c.SLock.Lock()
		reply.State = c.State
		reply.CurrNumMapTask = c.CurrNumMapTask
		reply.CurrNumReduceTask = c.CurrNumReduceTask
		//c.SLock.Unlock()
	}else {
		//c.SLock.Lock()
		reply.State = 1
		reply.CurrNumReduceTask = c.CurrNumReduceTask
		//c.SLock.Unlock()
	}

	//c.SLock.Unlock()
	return nil
}

func (c *Master) GetMapTask(args *MapTaskRequest, reply *MapTaskResponse) error {

	if c.State == 0 && len(c.MapTask) != 0 {
		mapTask, ok := <-c.MapTask
		if ok {
			reply.MapTask = mapTask
		}
		c.SLock.Lock()
		c.CurrNumMapTask = c.CurrNumMapTask - 1
		c.SLock.Unlock()
	}

	reply.NumMapTask = c.NumMapTask
	reply.NumReduceTask = c.NumReduceTask
	return nil
}

func (c *Master) GetMapFinTask(args *MapTaskRequest, reply *MapTaskResponse) error {

	if c.State == 0 && len(c.MapTaskFin) <= c.NumMapTask {
		c.MapTaskFin <- true
		if len(c.MapTaskFin) == c.NumMapTask {
			fmt.Println("mapTask has finished")
			c.State = 1
		}
	}
	//reply.State = c.State
	return nil
}

func (c *Master) GetReduceTask(args *ReduceTaskRequest, reply *ReduceTaskResponse) error {
	if c.State == 1 && len(c.ReduceTask) != 0 {
		reduceTask, ok := <-c.ReduceTask
		if ok {
			reply.ReduceTask = reduceTask
		}
		c.SLock.Lock()
		c.CurrNumReduceTask = c.CurrNumReduceTask - 1
		c.SLock.Unlock()
	}

	reply.NumMapTask = c.NumMapTask
	reply.NumReduceTask = c.NumReduceTask
	reply.CurrNumMapTask = c.CurrNumMapTask
	reply.CurrNumReduceTask = c.CurrNumReduceTask
	return nil
}

func (c *Master) GetReduceFinTask(args *ReduceTaskRequest, reply *ReduceTaskResponse) error {
	if c.State == 1 && len(c.ReduceTaskFin) <= c.NumReduceTask {
		fmt.Println("one reduceTask ->>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		c.ReduceTaskFin <- true
		if len(c.ReduceTaskFin) == c.NumReduceTask {
			fmt.Println("reduceTask has finished")
			c.State = 2
			c.TaskFin <- true
		}
	}

	//reply.State = c.State
	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (c *Master) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrmaster.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Master) Done() bool {
	ret := false

	// Your code here.
	//if c.State == 2 {
	//	ret = true
	//}
	select {
	case <- c.TaskFin:
		ret = true
		return ret
	}

	//return ret
}

//
// create a Master.
// main/mrmaster.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeMaster(files []string, nReduce int) *Master {
	//m := Master{}

	// Your code here.
	c := Master{
		State:         0,
		NumReduceTask: nReduce,
		NumMapTask:    len(files),
		MapTask:       make(chan Task, len(files)),
		ReduceTask:    make(chan Task, nReduce),
		MapTaskFin:    make(chan bool, len(files)),
		ReduceTaskFin: make(chan bool, nReduce),
		MapFin:        make(chan bool),
		ReduceFin: make(chan bool),
		TaskFin: make(chan bool),
		CurrNumMapTask: len(files),
		CurrNumReduceTask: nReduce,

	}

	// 初始化Task, 将任务分发到worker中Map chan，等待worke取task
	for fileId, filename := range files {
		c.MapTask <- Task{
			MapId:    fileId,
			Filename: filename,
		}
	}

	for i := 0; i < nReduce; i++ {
		c.ReduceTask <- Task{
			ReduceId: i,
		}
	}

	c.server()
	return &c
}
