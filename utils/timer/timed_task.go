package timer

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type Timer interface {
	GetTasks() []timeTask
	AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error)
	AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error)
	FindCron(taskName string) (*timeTask, error)
	StartTask(taskName string) error
	StopTask(taskName string) error
	Remove(taskName string, id int) error
	Clear(taskName string) error
	Close()
}

// timer 定时任务管理
type timer struct {
	taskList []timeTask
	sync.Mutex
}

type timeTask struct {
	Id        cron.EntryID `json:"id"`
	Name      string       `json:"name"`
	Spec      string       `json:"spec"`
	EntriyLen int          `json:"entriyLen"`
	Running   bool         `json:"running"`
	Corn      *cron.Cron   `json:"cron"`
	CreatedAt time.Time    `json:"createdAt"`
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *timer) AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	for _, task := range t.taskList {
		if task.Name == taskName {
			return 0, fmt.Errorf("任务 %s 已经存在", taskName)
		}
	}
	timeTask := timeTask{Name: taskName, Spec: spec, Corn: cron.New(cron.WithSeconds()), CreatedAt: time.Now()}
	id, err := timeTask.Corn.AddFunc(spec, task)
	timeTask.Corn.Start()
	timeTask.Id = id
	timeTask.Running = true
	t.taskList = append(t.taskList, timeTask)
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *timer) AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	for _, task := range t.taskList {
		if task.Name == taskName {
			return 0, fmt.Errorf("任务 %s 已经存在", taskName)
		}
	}
	timeTask := timeTask{Name: taskName, Spec: spec, Corn: cron.New(cron.WithSeconds()), CreatedAt: time.Now()}
	id, err := timeTask.Corn.AddJob(spec, job)
	timeTask.Corn.Start()
	timeTask.Id = id
	timeTask.Running = true
	t.taskList = append(t.taskList, timeTask)
	return id, err
}

// FindCron 获取对应taskName的cron 可能会为空
func (t *timer) FindCron(taskName string) (*timeTask, error) {
	t.Lock()
	defer t.Unlock()
	for _, task := range t.taskList {
		if task.Name == taskName {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("任务 %s 不存在", taskName)
}

// StartTask 开始任务
func (t *timer) StartTask(taskName string) error {
	t.Lock()
	defer t.Unlock()
	for key, task := range t.taskList {
		if task.Name == taskName {
			task.Corn.Start()
			t.taskList[key].Running = true
			return nil
		}
	}
	return fmt.Errorf("任务 %s 不存在", taskName)
}

// StopTask 停止任务
func (t *timer) StopTask(taskName string) error {
	t.Lock()
	defer t.Unlock()
	for key, task := range t.taskList {
		if task.Name == taskName {
			task.Corn.Stop()
			t.taskList[key].Running = false
			return nil
		}
	}
	return fmt.Errorf("任务 %s 不存在", taskName)
}

// Remove 从taskName 删除指定任务
func (t *timer) Remove(taskName string, id int) error {
	t.Lock()
	defer t.Unlock()
	for _, task := range t.taskList {
		if task.Name == taskName {
			task.Corn.Remove(cron.EntryID(id))
			return nil
		}
	}
	return fmt.Errorf("任务 %s 不存在", taskName)
}

// Clear 清除任务
func (t *timer) Clear(taskName string) error {
	t.Lock()
	defer t.Unlock()
	for key, task := range t.taskList {
		if task.Name == taskName {
			task.Corn.Stop()
			t.taskList[key].Running = false
			t.taskList = append(t.taskList[:key], t.taskList[key+1:]...)
			return nil
		}
	}
	return fmt.Errorf("任务 %s 不存在", taskName)
}

// Close 释放资源
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for key, task := range t.taskList {
		task.Corn.Stop()
		t.taskList[key].Running = false
	}
}

// GetTasks 定时任务列表
func (t *timer) GetTasks() []timeTask {
	var tl []timeTask
	for _, ttsk := range t.taskList {
		if ttsk.Id > 0 {
			tl = append(tl, ttsk)
		}
	}
	return tl
}

func NewTimerTask() Timer {
	return &timer{taskList: []timeTask{}}
}
