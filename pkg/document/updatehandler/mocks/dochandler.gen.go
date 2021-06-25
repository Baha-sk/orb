// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/document"
	"github.com/trustbloc/sidetree-core-go/pkg/restapi/dochandler"
)

type Processor struct {
	NamespaceStub        func() string
	namespaceMutex       sync.RWMutex
	namespaceArgsForCall []struct {
	}
	namespaceReturns struct {
		result1 string
	}
	namespaceReturnsOnCall map[int]struct {
		result1 string
	}
	ProcessOperationStub        func([]byte, uint64) (*document.ResolutionResult, error)
	processOperationMutex       sync.RWMutex
	processOperationArgsForCall []struct {
		arg1 []byte
		arg2 uint64
	}
	processOperationReturns struct {
		result1 *document.ResolutionResult
		result2 error
	}
	processOperationReturnsOnCall map[int]struct {
		result1 *document.ResolutionResult
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Processor) Namespace() string {
	fake.namespaceMutex.Lock()
	ret, specificReturn := fake.namespaceReturnsOnCall[len(fake.namespaceArgsForCall)]
	fake.namespaceArgsForCall = append(fake.namespaceArgsForCall, struct {
	}{})
	fake.recordInvocation("Namespace", []interface{}{})
	fake.namespaceMutex.Unlock()
	if fake.NamespaceStub != nil {
		return fake.NamespaceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.namespaceReturns
	return fakeReturns.result1
}

func (fake *Processor) NamespaceCallCount() int {
	fake.namespaceMutex.RLock()
	defer fake.namespaceMutex.RUnlock()
	return len(fake.namespaceArgsForCall)
}

func (fake *Processor) NamespaceCalls(stub func() string) {
	fake.namespaceMutex.Lock()
	defer fake.namespaceMutex.Unlock()
	fake.NamespaceStub = stub
}

func (fake *Processor) NamespaceReturns(result1 string) {
	fake.namespaceMutex.Lock()
	defer fake.namespaceMutex.Unlock()
	fake.NamespaceStub = nil
	fake.namespaceReturns = struct {
		result1 string
	}{result1}
}

func (fake *Processor) NamespaceReturnsOnCall(i int, result1 string) {
	fake.namespaceMutex.Lock()
	defer fake.namespaceMutex.Unlock()
	fake.NamespaceStub = nil
	if fake.namespaceReturnsOnCall == nil {
		fake.namespaceReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.namespaceReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Processor) ProcessOperation(arg1 []byte, arg2 uint64) (*document.ResolutionResult, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.processOperationMutex.Lock()
	ret, specificReturn := fake.processOperationReturnsOnCall[len(fake.processOperationArgsForCall)]
	fake.processOperationArgsForCall = append(fake.processOperationArgsForCall, struct {
		arg1 []byte
		arg2 uint64
	}{arg1Copy, arg2})
	fake.recordInvocation("ProcessOperation", []interface{}{arg1Copy, arg2})
	fake.processOperationMutex.Unlock()
	if fake.ProcessOperationStub != nil {
		return fake.ProcessOperationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.processOperationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Processor) ProcessOperationCallCount() int {
	fake.processOperationMutex.RLock()
	defer fake.processOperationMutex.RUnlock()
	return len(fake.processOperationArgsForCall)
}

func (fake *Processor) ProcessOperationCalls(stub func([]byte, uint64) (*document.ResolutionResult, error)) {
	fake.processOperationMutex.Lock()
	defer fake.processOperationMutex.Unlock()
	fake.ProcessOperationStub = stub
}

func (fake *Processor) ProcessOperationArgsForCall(i int) ([]byte, uint64) {
	fake.processOperationMutex.RLock()
	defer fake.processOperationMutex.RUnlock()
	argsForCall := fake.processOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Processor) ProcessOperationReturns(result1 *document.ResolutionResult, result2 error) {
	fake.processOperationMutex.Lock()
	defer fake.processOperationMutex.Unlock()
	fake.ProcessOperationStub = nil
	fake.processOperationReturns = struct {
		result1 *document.ResolutionResult
		result2 error
	}{result1, result2}
}

func (fake *Processor) ProcessOperationReturnsOnCall(i int, result1 *document.ResolutionResult, result2 error) {
	fake.processOperationMutex.Lock()
	defer fake.processOperationMutex.Unlock()
	fake.ProcessOperationStub = nil
	if fake.processOperationReturnsOnCall == nil {
		fake.processOperationReturnsOnCall = make(map[int]struct {
			result1 *document.ResolutionResult
			result2 error
		})
	}
	fake.processOperationReturnsOnCall[i] = struct {
		result1 *document.ResolutionResult
		result2 error
	}{result1, result2}
}

func (fake *Processor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.namespaceMutex.RLock()
	defer fake.namespaceMutex.RUnlock()
	fake.processOperationMutex.RLock()
	defer fake.processOperationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Processor) recordInvocation(key string, args []interface{}) {
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

var _ dochandler.Processor = new(Processor)
