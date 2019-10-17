package storage

import (
	. "github.com/ngaut/gearmand/common"
)

/*定义队列的接口的*/
type JobQueue interface {
	Init() error
	AddJob(j *Job) error
	DoneJob(j *Job) error
	GetJobs() ([]*Job, error)
}
