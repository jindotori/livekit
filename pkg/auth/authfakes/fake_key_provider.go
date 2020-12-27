// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/auth"
)

type FakeKeyProvider struct {
	GetSecretStub        func(string) string
	getSecretMutex       sync.RWMutex
	getSecretArgsForCall []struct {
		arg1 string
	}
	getSecretReturns struct {
		result1 string
	}
	getSecretReturnsOnCall map[int]struct {
		result1 string
	}
	KeyCountStub        func() int
	keyCountMutex       sync.RWMutex
	keyCountArgsForCall []struct {
	}
	keyCountReturns struct {
		result1 int
	}
	keyCountReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKeyProvider) GetSecret(arg1 string) string {
	fake.getSecretMutex.Lock()
	ret, specificReturn := fake.getSecretReturnsOnCall[len(fake.getSecretArgsForCall)]
	fake.getSecretArgsForCall = append(fake.getSecretArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetSecretStub
	fakeReturns := fake.getSecretReturns
	fake.recordInvocation("GetSecret", []interface{}{arg1})
	fake.getSecretMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKeyProvider) GetSecretCallCount() int {
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	return len(fake.getSecretArgsForCall)
}

func (fake *FakeKeyProvider) GetSecretCalls(stub func(string) string) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = stub
}

func (fake *FakeKeyProvider) GetSecretArgsForCall(i int) string {
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	argsForCall := fake.getSecretArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKeyProvider) GetSecretReturns(result1 string) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = nil
	fake.getSecretReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeKeyProvider) GetSecretReturnsOnCall(i int, result1 string) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = nil
	if fake.getSecretReturnsOnCall == nil {
		fake.getSecretReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getSecretReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeKeyProvider) NumKeys() int {
	fake.keyCountMutex.Lock()
	ret, specificReturn := fake.keyCountReturnsOnCall[len(fake.keyCountArgsForCall)]
	fake.keyCountArgsForCall = append(fake.keyCountArgsForCall, struct {
	}{})
	stub := fake.KeyCountStub
	fakeReturns := fake.keyCountReturns
	fake.recordInvocation("NumKeys", []interface{}{})
	fake.keyCountMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKeyProvider) KeyCountCallCount() int {
	fake.keyCountMutex.RLock()
	defer fake.keyCountMutex.RUnlock()
	return len(fake.keyCountArgsForCall)
}

func (fake *FakeKeyProvider) KeyCountCalls(stub func() int) {
	fake.keyCountMutex.Lock()
	defer fake.keyCountMutex.Unlock()
	fake.KeyCountStub = stub
}

func (fake *FakeKeyProvider) KeyCountReturns(result1 int) {
	fake.keyCountMutex.Lock()
	defer fake.keyCountMutex.Unlock()
	fake.KeyCountStub = nil
	fake.keyCountReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeKeyProvider) KeyCountReturnsOnCall(i int, result1 int) {
	fake.keyCountMutex.Lock()
	defer fake.keyCountMutex.Unlock()
	fake.KeyCountStub = nil
	if fake.keyCountReturnsOnCall == nil {
		fake.keyCountReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.keyCountReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeKeyProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	fake.keyCountMutex.RLock()
	defer fake.keyCountMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKeyProvider) recordInvocation(key string, args []interface{}) {
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

var _ auth.KeyProvider = new(FakeKeyProvider)