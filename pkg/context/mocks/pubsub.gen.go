// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/trustbloc/orb/pkg/pubsub/spi"
)

type PubSub struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	PublishWithOptsStub        func(string, *message.Message, ...spi.Option) error
	publishWithOptsMutex       sync.RWMutex
	publishWithOptsArgsForCall []struct {
		arg1 string
		arg2 *message.Message
		arg3 []spi.Option
	}
	publishWithOptsReturns struct {
		result1 error
	}
	publishWithOptsReturnsOnCall map[int]struct {
		result1 error
	}
	SubscribeWithOptsStub        func(context.Context, string, ...spi.Option) (<-chan *message.Message, error)
	subscribeWithOptsMutex       sync.RWMutex
	subscribeWithOptsArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []spi.Option
	}
	subscribeWithOptsReturns struct {
		result1 <-chan *message.Message
		result2 error
	}
	subscribeWithOptsReturnsOnCall map[int]struct {
		result1 <-chan *message.Message
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PubSub) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *PubSub) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *PubSub) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *PubSub) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *PubSub) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PubSub) PublishWithOpts(arg1 string, arg2 *message.Message, arg3 ...spi.Option) error {
	fake.publishWithOptsMutex.Lock()
	ret, specificReturn := fake.publishWithOptsReturnsOnCall[len(fake.publishWithOptsArgsForCall)]
	fake.publishWithOptsArgsForCall = append(fake.publishWithOptsArgsForCall, struct {
		arg1 string
		arg2 *message.Message
		arg3 []spi.Option
	}{arg1, arg2, arg3})
	stub := fake.PublishWithOptsStub
	fakeReturns := fake.publishWithOptsReturns
	fake.recordInvocation("PublishWithOpts", []interface{}{arg1, arg2, arg3})
	fake.publishWithOptsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *PubSub) PublishWithOptsCallCount() int {
	fake.publishWithOptsMutex.RLock()
	defer fake.publishWithOptsMutex.RUnlock()
	return len(fake.publishWithOptsArgsForCall)
}

func (fake *PubSub) PublishWithOptsCalls(stub func(string, *message.Message, ...spi.Option) error) {
	fake.publishWithOptsMutex.Lock()
	defer fake.publishWithOptsMutex.Unlock()
	fake.PublishWithOptsStub = stub
}

func (fake *PubSub) PublishWithOptsArgsForCall(i int) (string, *message.Message, []spi.Option) {
	fake.publishWithOptsMutex.RLock()
	defer fake.publishWithOptsMutex.RUnlock()
	argsForCall := fake.publishWithOptsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PubSub) PublishWithOptsReturns(result1 error) {
	fake.publishWithOptsMutex.Lock()
	defer fake.publishWithOptsMutex.Unlock()
	fake.PublishWithOptsStub = nil
	fake.publishWithOptsReturns = struct {
		result1 error
	}{result1}
}

func (fake *PubSub) PublishWithOptsReturnsOnCall(i int, result1 error) {
	fake.publishWithOptsMutex.Lock()
	defer fake.publishWithOptsMutex.Unlock()
	fake.PublishWithOptsStub = nil
	if fake.publishWithOptsReturnsOnCall == nil {
		fake.publishWithOptsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishWithOptsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PubSub) SubscribeWithOpts(arg1 context.Context, arg2 string, arg3 ...spi.Option) (<-chan *message.Message, error) {
	fake.subscribeWithOptsMutex.Lock()
	ret, specificReturn := fake.subscribeWithOptsReturnsOnCall[len(fake.subscribeWithOptsArgsForCall)]
	fake.subscribeWithOptsArgsForCall = append(fake.subscribeWithOptsArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []spi.Option
	}{arg1, arg2, arg3})
	stub := fake.SubscribeWithOptsStub
	fakeReturns := fake.subscribeWithOptsReturns
	fake.recordInvocation("SubscribeWithOpts", []interface{}{arg1, arg2, arg3})
	fake.subscribeWithOptsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PubSub) SubscribeWithOptsCallCount() int {
	fake.subscribeWithOptsMutex.RLock()
	defer fake.subscribeWithOptsMutex.RUnlock()
	return len(fake.subscribeWithOptsArgsForCall)
}

func (fake *PubSub) SubscribeWithOptsCalls(stub func(context.Context, string, ...spi.Option) (<-chan *message.Message, error)) {
	fake.subscribeWithOptsMutex.Lock()
	defer fake.subscribeWithOptsMutex.Unlock()
	fake.SubscribeWithOptsStub = stub
}

func (fake *PubSub) SubscribeWithOptsArgsForCall(i int) (context.Context, string, []spi.Option) {
	fake.subscribeWithOptsMutex.RLock()
	defer fake.subscribeWithOptsMutex.RUnlock()
	argsForCall := fake.subscribeWithOptsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PubSub) SubscribeWithOptsReturns(result1 <-chan *message.Message, result2 error) {
	fake.subscribeWithOptsMutex.Lock()
	defer fake.subscribeWithOptsMutex.Unlock()
	fake.SubscribeWithOptsStub = nil
	fake.subscribeWithOptsReturns = struct {
		result1 <-chan *message.Message
		result2 error
	}{result1, result2}
}

func (fake *PubSub) SubscribeWithOptsReturnsOnCall(i int, result1 <-chan *message.Message, result2 error) {
	fake.subscribeWithOptsMutex.Lock()
	defer fake.subscribeWithOptsMutex.Unlock()
	fake.SubscribeWithOptsStub = nil
	if fake.subscribeWithOptsReturnsOnCall == nil {
		fake.subscribeWithOptsReturnsOnCall = make(map[int]struct {
			result1 <-chan *message.Message
			result2 error
		})
	}
	fake.subscribeWithOptsReturnsOnCall[i] = struct {
		result1 <-chan *message.Message
		result2 error
	}{result1, result2}
}

func (fake *PubSub) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.publishWithOptsMutex.RLock()
	defer fake.publishWithOptsMutex.RUnlock()
	fake.subscribeWithOptsMutex.RLock()
	defer fake.subscribeWithOptsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PubSub) recordInvocation(key string, args []interface{}) {
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
