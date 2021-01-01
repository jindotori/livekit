// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/pion/rtp"
)

type FakeReceiver struct {
	GetBufferedPacketsStub        func(uint32, uint16, uint32, []uint16) []rtp.Packet
	getBufferedPacketsMutex       sync.RWMutex
	getBufferedPacketsArgsForCall []struct {
		arg1 uint32
		arg2 uint16
		arg3 uint32
		arg4 []uint16
	}
	getBufferedPacketsReturns struct {
		result1 []rtp.Packet
	}
	getBufferedPacketsReturnsOnCall map[int]struct {
		result1 []rtp.Packet
	}
	ReadRTPStub        func() (*rtp.Packet, error)
	readRTPMutex       sync.RWMutex
	readRTPArgsForCall []struct {
	}
	readRTPReturns struct {
		result1 *rtp.Packet
		result2 error
	}
	readRTPReturnsOnCall map[int]struct {
		result1 *rtp.Packet
		result2 error
	}
	StartStub        func()
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	TrackIdStub        func() string
	trackIdMutex       sync.RWMutex
	trackIdArgsForCall []struct {
	}
	trackIdReturns struct {
		result1 string
	}
	trackIdReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReceiver) GetBufferedPackets(arg1 uint32, arg2 uint16, arg3 uint32, arg4 []uint16) []rtp.Packet {
	var arg4Copy []uint16
	if arg4 != nil {
		arg4Copy = make([]uint16, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.getBufferedPacketsMutex.Lock()
	ret, specificReturn := fake.getBufferedPacketsReturnsOnCall[len(fake.getBufferedPacketsArgsForCall)]
	fake.getBufferedPacketsArgsForCall = append(fake.getBufferedPacketsArgsForCall, struct {
		arg1 uint32
		arg2 uint16
		arg3 uint32
		arg4 []uint16
	}{arg1, arg2, arg3, arg4Copy})
	stub := fake.GetBufferedPacketsStub
	fakeReturns := fake.getBufferedPacketsReturns
	fake.recordInvocation("GetBufferedPackets", []interface{}{arg1, arg2, arg3, arg4Copy})
	fake.getBufferedPacketsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReceiver) GetBufferedPacketsCallCount() int {
	fake.getBufferedPacketsMutex.RLock()
	defer fake.getBufferedPacketsMutex.RUnlock()
	return len(fake.getBufferedPacketsArgsForCall)
}

func (fake *FakeReceiver) GetBufferedPacketsCalls(stub func(uint32, uint16, uint32, []uint16) []rtp.Packet) {
	fake.getBufferedPacketsMutex.Lock()
	defer fake.getBufferedPacketsMutex.Unlock()
	fake.GetBufferedPacketsStub = stub
}

func (fake *FakeReceiver) GetBufferedPacketsArgsForCall(i int) (uint32, uint16, uint32, []uint16) {
	fake.getBufferedPacketsMutex.RLock()
	defer fake.getBufferedPacketsMutex.RUnlock()
	argsForCall := fake.getBufferedPacketsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeReceiver) GetBufferedPacketsReturns(result1 []rtp.Packet) {
	fake.getBufferedPacketsMutex.Lock()
	defer fake.getBufferedPacketsMutex.Unlock()
	fake.GetBufferedPacketsStub = nil
	fake.getBufferedPacketsReturns = struct {
		result1 []rtp.Packet
	}{result1}
}

func (fake *FakeReceiver) GetBufferedPacketsReturnsOnCall(i int, result1 []rtp.Packet) {
	fake.getBufferedPacketsMutex.Lock()
	defer fake.getBufferedPacketsMutex.Unlock()
	fake.GetBufferedPacketsStub = nil
	if fake.getBufferedPacketsReturnsOnCall == nil {
		fake.getBufferedPacketsReturnsOnCall = make(map[int]struct {
			result1 []rtp.Packet
		})
	}
	fake.getBufferedPacketsReturnsOnCall[i] = struct {
		result1 []rtp.Packet
	}{result1}
}

func (fake *FakeReceiver) ReadRTP() (*rtp.Packet, error) {
	fake.readRTPMutex.Lock()
	ret, specificReturn := fake.readRTPReturnsOnCall[len(fake.readRTPArgsForCall)]
	fake.readRTPArgsForCall = append(fake.readRTPArgsForCall, struct {
	}{})
	stub := fake.ReadRTPStub
	fakeReturns := fake.readRTPReturns
	fake.recordInvocation("ReadRTP", []interface{}{})
	fake.readRTPMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReceiver) ReadRTPCallCount() int {
	fake.readRTPMutex.RLock()
	defer fake.readRTPMutex.RUnlock()
	return len(fake.readRTPArgsForCall)
}

func (fake *FakeReceiver) ReadRTPCalls(stub func() (*rtp.Packet, error)) {
	fake.readRTPMutex.Lock()
	defer fake.readRTPMutex.Unlock()
	fake.ReadRTPStub = stub
}

func (fake *FakeReceiver) ReadRTPReturns(result1 *rtp.Packet, result2 error) {
	fake.readRTPMutex.Lock()
	defer fake.readRTPMutex.Unlock()
	fake.ReadRTPStub = nil
	fake.readRTPReturns = struct {
		result1 *rtp.Packet
		result2 error
	}{result1, result2}
}

func (fake *FakeReceiver) ReadRTPReturnsOnCall(i int, result1 *rtp.Packet, result2 error) {
	fake.readRTPMutex.Lock()
	defer fake.readRTPMutex.Unlock()
	fake.ReadRTPStub = nil
	if fake.readRTPReturnsOnCall == nil {
		fake.readRTPReturnsOnCall = make(map[int]struct {
			result1 *rtp.Packet
			result2 error
		})
	}
	fake.readRTPReturnsOnCall[i] = struct {
		result1 *rtp.Packet
		result2 error
	}{result1, result2}
}

func (fake *FakeReceiver) Start() {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		fake.StartStub()
	}
}

func (fake *FakeReceiver) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeReceiver) StartCalls(stub func()) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeReceiver) TrackId() string {
	fake.trackIdMutex.Lock()
	ret, specificReturn := fake.trackIdReturnsOnCall[len(fake.trackIdArgsForCall)]
	fake.trackIdArgsForCall = append(fake.trackIdArgsForCall, struct {
	}{})
	stub := fake.TrackIdStub
	fakeReturns := fake.trackIdReturns
	fake.recordInvocation("TrackId", []interface{}{})
	fake.trackIdMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReceiver) TrackIdCallCount() int {
	fake.trackIdMutex.RLock()
	defer fake.trackIdMutex.RUnlock()
	return len(fake.trackIdArgsForCall)
}

func (fake *FakeReceiver) TrackIdCalls(stub func() string) {
	fake.trackIdMutex.Lock()
	defer fake.trackIdMutex.Unlock()
	fake.TrackIdStub = stub
}

func (fake *FakeReceiver) TrackIdReturns(result1 string) {
	fake.trackIdMutex.Lock()
	defer fake.trackIdMutex.Unlock()
	fake.TrackIdStub = nil
	fake.trackIdReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeReceiver) TrackIdReturnsOnCall(i int, result1 string) {
	fake.trackIdMutex.Lock()
	defer fake.trackIdMutex.Unlock()
	fake.TrackIdStub = nil
	if fake.trackIdReturnsOnCall == nil {
		fake.trackIdReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.trackIdReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeReceiver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBufferedPacketsMutex.RLock()
	defer fake.getBufferedPacketsMutex.RUnlock()
	fake.readRTPMutex.RLock()
	defer fake.readRTPMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.trackIdMutex.RLock()
	defer fake.trackIdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReceiver) recordInvocation(key string, args []interface{}) {
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

var _ types.Receiver = new(FakeReceiver)
