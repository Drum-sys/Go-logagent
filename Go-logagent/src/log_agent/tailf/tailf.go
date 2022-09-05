package tailf

import (
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"sync"
	"time"
)

type CollectPath struct {
	LogPath string `json:"log_path"`
	Topic   string `json:"topic"`
}

type TailObj struct {
	tail *tail.Tail
	conf CollectPath
	status int
	exitChan chan int //delete
}

type TextMsg struct {
	Msg string
	Topic string
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan chan *TextMsg
	lock sync.Mutex
}

var (
	tailObjMgr *TailObjMgr
)

const (
	StatusDelete = 2
	StatusNormal = 1
)

func UpdateConf(conf []CollectPath) (err error) {
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()
	for _, oneConf := range conf{
		var isRunning = false
		for _, obj := range tailObjMgr.tailObjs {
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				//obj.status = StatusPut
				break
			}
		}
		if isRunning {
			continue
		}
		CreateNewTask(oneConf)
	}

	var tailObjs []*TailObj

	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete
		for _, oneConf := range conf {
			if oneConf.LogPath == obj.conf.LogPath{
				obj.status = StatusNormal
				break
			}
		}
		if obj.status == StatusDelete {
			obj.exitChan <- 1
			continue
		}
		tailObjs = append(tailObjs, obj)
	}
	tailObjMgr.tailObjs = tailObjs
	return
}

func CreateNewTask(conf CollectPath) (err error){
	obj := &TailObj{
		conf: conf,
		exitChan: make(chan int, 1),
	}

	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
	})
	if errTail != nil {
		logs.Error("collect filename[%s], failed, err:%v", conf.LogPath, errTail)
		//fmt.Println("tail file err:", err)
		return
	}
	obj.tail = tails
	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

	go readFromTail(obj)
	return
}


func InitTail(conf []CollectPath, chanSize int) (err error) {
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	if len(conf) == 0 {
		logs.Error("invalid config fot log collect, conf:%v", conf)
	}

	for _, v := range conf {
		/*obj := &TailObj{
			conf: v,
		}

		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true,
			Follow: true,
			//Location: &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll: true,
		})
		if errTail != nil {
			logs.Error("invalid tail file failed, cong:%v, err:%v", conf)
			continue
		}
		obj.tail = tails
		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFromTail(obj)*/
		CreateNewTask(v)
	}

	return
}

func readFromTail(tailObj *TailObj)  {
	for true {
		select {
		case line, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename:%v", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{
				Msg: line.Text,
				Topic: tailObj.conf.Topic,
			}
			tailObjMgr.msgChan <- textMsg
		case <-tailObj.exitChan:
			logs.Warn("tail obj delete conf exit")
			return
		}
	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <- tailObjMgr.msgChan
	return msg
}
