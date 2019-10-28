package server

import (
	"bytes"
	"encoding/json"
	. "github.com/ngaut/gearmand/common"
	"net"
)

/*worker的状态的*/
const (
	wsRunning         = 1
	wsSleep           = 2
	wsPrepareForSleep = 3
)

func status2str(status int) string {
	switch status {
	case wsRunning:
		return "running"
	case wsSleep:
		return "sleep"
	case wsPrepareForSleep:
		return "prepareForSleep"
	}

	return "unknown"
}

type Worker struct {
	net.Conn
	Session

	workerId    string
	status      int
	runningJobs map[string]*Job
	canDo       map[string]bool
}

func (self *Worker) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	enc := json.NewEncoder(b)
	m := make(map[string]interface{})
	m["sessionId"] = self.SessionId
	m["Id"] = self.workerId
	m["status"] = status2str(self.status)

	/*worker注册的方法*/
	canDoSlice := make([]string, 0, len(self.canDo))
	for k, _ := range self.canDo {
		canDoSlice = append(canDoSlice, k)
	}
	m["canDo"] = canDoSlice

	/*正在处理的任务*/
	jobSlice := make([]string, 0, len(self.canDo))
	for k, _ := range self.runningJobs {
		jobSlice = append(jobSlice, k)
	}
	m["runningJobs"] = jobSlice

	if err := enc.Encode(m); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
