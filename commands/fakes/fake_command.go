// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/mike-carey/cfquery/commands"
	"github.com/mike-carey/cfquery/query"
)

type FakeCommand struct {
	ExecuteStub        func([]string) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 []string
	}
	executeReturns struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	GroupByOptionsStub        func() []string
	groupByOptionsMutex       sync.RWMutex
	groupByOptionsArgsForCall []struct {
	}
	groupByOptionsReturns struct {
		result1 []string
	}
	groupByOptionsReturnsOnCall map[int]struct {
		result1 []string
	}
	RunStub        func(*commands.Options, query.Inquisitor) (interface{}, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 *commands.Options
		arg2 query.Inquisitor
	}
	runReturns struct {
		result1 interface{}
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	SortByOptionsStub        func() []string
	sortByOptionsMutex       sync.RWMutex
	sortByOptionsArgsForCall []struct {
	}
	sortByOptionsReturns struct {
		result1 []string
	}
	sortByOptionsReturnsOnCall map[int]struct {
		result1 []string
	}
	TargetOptionsStub        func() []string
	targetOptionsMutex       sync.RWMutex
	targetOptionsArgsForCall []struct {
	}
	targetOptionsReturns struct {
		result1 []string
	}
	targetOptionsReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommand) Execute(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("Execute", []interface{}{arg1Copy})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1
}

func (fake *FakeCommand) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeCommand) ExecuteCalls(stub func([]string) error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeCommand) ExecuteArgsForCall(i int) []string {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCommand) ExecuteReturns(result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommand) ExecuteReturnsOnCall(i int, result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommand) GroupByOptions() []string {
	fake.groupByOptionsMutex.Lock()
	ret, specificReturn := fake.groupByOptionsReturnsOnCall[len(fake.groupByOptionsArgsForCall)]
	fake.groupByOptionsArgsForCall = append(fake.groupByOptionsArgsForCall, struct {
	}{})
	fake.recordInvocation("GroupByOptions", []interface{}{})
	fake.groupByOptionsMutex.Unlock()
	if fake.GroupByOptionsStub != nil {
		return fake.GroupByOptionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.groupByOptionsReturns
	return fakeReturns.result1
}

func (fake *FakeCommand) GroupByOptionsCallCount() int {
	fake.groupByOptionsMutex.RLock()
	defer fake.groupByOptionsMutex.RUnlock()
	return len(fake.groupByOptionsArgsForCall)
}

func (fake *FakeCommand) GroupByOptionsCalls(stub func() []string) {
	fake.groupByOptionsMutex.Lock()
	defer fake.groupByOptionsMutex.Unlock()
	fake.GroupByOptionsStub = stub
}

func (fake *FakeCommand) GroupByOptionsReturns(result1 []string) {
	fake.groupByOptionsMutex.Lock()
	defer fake.groupByOptionsMutex.Unlock()
	fake.GroupByOptionsStub = nil
	fake.groupByOptionsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCommand) GroupByOptionsReturnsOnCall(i int, result1 []string) {
	fake.groupByOptionsMutex.Lock()
	defer fake.groupByOptionsMutex.Unlock()
	fake.GroupByOptionsStub = nil
	if fake.groupByOptionsReturnsOnCall == nil {
		fake.groupByOptionsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.groupByOptionsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCommand) Run(arg1 *commands.Options, arg2 query.Inquisitor) (interface{}, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 *commands.Options
		arg2 query.Inquisitor
	}{arg1, arg2})
	fake.recordInvocation("Run", []interface{}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCommand) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeCommand) RunCalls(stub func(*commands.Options, query.Inquisitor) (interface{}, error)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeCommand) RunArgsForCall(i int) (*commands.Options, query.Inquisitor) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCommand) RunReturns(result1 interface{}, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCommand) RunReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCommand) SortByOptions() []string {
	fake.sortByOptionsMutex.Lock()
	ret, specificReturn := fake.sortByOptionsReturnsOnCall[len(fake.sortByOptionsArgsForCall)]
	fake.sortByOptionsArgsForCall = append(fake.sortByOptionsArgsForCall, struct {
	}{})
	fake.recordInvocation("SortByOptions", []interface{}{})
	fake.sortByOptionsMutex.Unlock()
	if fake.SortByOptionsStub != nil {
		return fake.SortByOptionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sortByOptionsReturns
	return fakeReturns.result1
}

func (fake *FakeCommand) SortByOptionsCallCount() int {
	fake.sortByOptionsMutex.RLock()
	defer fake.sortByOptionsMutex.RUnlock()
	return len(fake.sortByOptionsArgsForCall)
}

func (fake *FakeCommand) SortByOptionsCalls(stub func() []string) {
	fake.sortByOptionsMutex.Lock()
	defer fake.sortByOptionsMutex.Unlock()
	fake.SortByOptionsStub = stub
}

func (fake *FakeCommand) SortByOptionsReturns(result1 []string) {
	fake.sortByOptionsMutex.Lock()
	defer fake.sortByOptionsMutex.Unlock()
	fake.SortByOptionsStub = nil
	fake.sortByOptionsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCommand) SortByOptionsReturnsOnCall(i int, result1 []string) {
	fake.sortByOptionsMutex.Lock()
	defer fake.sortByOptionsMutex.Unlock()
	fake.SortByOptionsStub = nil
	if fake.sortByOptionsReturnsOnCall == nil {
		fake.sortByOptionsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.sortByOptionsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCommand) TargetOptions() []string {
	fake.targetOptionsMutex.Lock()
	ret, specificReturn := fake.targetOptionsReturnsOnCall[len(fake.targetOptionsArgsForCall)]
	fake.targetOptionsArgsForCall = append(fake.targetOptionsArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetOptions", []interface{}{})
	fake.targetOptionsMutex.Unlock()
	if fake.TargetOptionsStub != nil {
		return fake.TargetOptionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetOptionsReturns
	return fakeReturns.result1
}

func (fake *FakeCommand) TargetOptionsCallCount() int {
	fake.targetOptionsMutex.RLock()
	defer fake.targetOptionsMutex.RUnlock()
	return len(fake.targetOptionsArgsForCall)
}

func (fake *FakeCommand) TargetOptionsCalls(stub func() []string) {
	fake.targetOptionsMutex.Lock()
	defer fake.targetOptionsMutex.Unlock()
	fake.TargetOptionsStub = stub
}

func (fake *FakeCommand) TargetOptionsReturns(result1 []string) {
	fake.targetOptionsMutex.Lock()
	defer fake.targetOptionsMutex.Unlock()
	fake.TargetOptionsStub = nil
	fake.targetOptionsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCommand) TargetOptionsReturnsOnCall(i int, result1 []string) {
	fake.targetOptionsMutex.Lock()
	defer fake.targetOptionsMutex.Unlock()
	fake.TargetOptionsStub = nil
	if fake.targetOptionsReturnsOnCall == nil {
		fake.targetOptionsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.targetOptionsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeCommand) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.groupByOptionsMutex.RLock()
	defer fake.groupByOptionsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.sortByOptionsMutex.RLock()
	defer fake.sortByOptionsMutex.RUnlock()
	fake.targetOptionsMutex.RLock()
	defer fake.targetOptionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCommand) recordInvocation(key string, args []interface{}) {
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

var _ commands.Command = new(FakeCommand)
