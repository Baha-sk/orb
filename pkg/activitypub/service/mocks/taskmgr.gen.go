// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"
	"time"
)

type TaskManager struct {
	RegisterTaskStub        func(taskType string, interval, maxRunTime time.Duration, task func())
	registerTaskMutex       sync.RWMutex
	registerTaskArgsForCall []struct {
		taskType   string
		interval   time.Duration
		maxRunTime time.Duration
		task       func()
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TaskManager) RegisterTask(taskType string, interval time.Duration, maxRunTime time.Duration, task func()) {
	fake.registerTaskMutex.Lock()
	fake.registerTaskArgsForCall = append(fake.registerTaskArgsForCall, struct {
		taskType   string
		interval   time.Duration
		maxRunTime time.Duration
		task       func()
	}{taskType, interval, maxRunTime, task})
	fake.recordInvocation("RegisterTask", []interface{}{taskType, interval, maxRunTime, task})
	fake.registerTaskMutex.Unlock()
	if fake.RegisterTaskStub != nil {
		fake.RegisterTaskStub(taskType, interval, maxRunTime, task)
	}
}

func (fake *TaskManager) RegisterTaskCallCount() int {
	fake.registerTaskMutex.RLock()
	defer fake.registerTaskMutex.RUnlock()
	return len(fake.registerTaskArgsForCall)
}

func (fake *TaskManager) RegisterTaskArgsForCall(i int) (string, time.Duration, time.Duration, func()) {
	fake.registerTaskMutex.RLock()
	defer fake.registerTaskMutex.RUnlock()
	return fake.registerTaskArgsForCall[i].taskType, fake.registerTaskArgsForCall[i].interval, fake.registerTaskArgsForCall[i].maxRunTime, fake.registerTaskArgsForCall[i].task
}

func (fake *TaskManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerTaskMutex.RLock()
	defer fake.registerTaskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TaskManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
