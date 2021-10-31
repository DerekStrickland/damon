// Code generated by counterfeiter. DO NOT EDIT.
package componentfakes

import (
	"sync"

	"github.com/hashicorp/damon/component"
	"github.com/rivo/tview"
)

type FakeTextView struct {
	ClearStub        func()
	clearMutex       sync.RWMutex
	clearArgsForCall []struct {
	}
	GetTextStub        func() string
	getTextMutex       sync.RWMutex
	getTextArgsForCall []struct {
	}
	getTextReturns struct {
		result1 string
	}
	getTextReturnsOnCall map[int]struct {
		result1 string
	}
	ModifyPrimitiveStub        func(func(t *tview.TextView))
	modifyPrimitiveMutex       sync.RWMutex
	modifyPrimitiveArgsForCall []struct {
		arg1 func(t *tview.TextView)
	}
	PrimitiveStub        func() tview.Primitive
	primitiveMutex       sync.RWMutex
	primitiveArgsForCall []struct {
	}
	primitiveReturns struct {
		result1 tview.Primitive
	}
	primitiveReturnsOnCall map[int]struct {
		result1 tview.Primitive
	}
	SetTextStub        func(string)
	setTextMutex       sync.RWMutex
	setTextArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTextView) Clear() {
	fake.clearMutex.Lock()
	fake.clearArgsForCall = append(fake.clearArgsForCall, struct {
	}{})
	stub := fake.ClearStub
	fake.recordInvocation("Clear", []interface{}{})
	fake.clearMutex.Unlock()
	if stub != nil {
		fake.ClearStub()
	}
}

func (fake *FakeTextView) ClearCallCount() int {
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	return len(fake.clearArgsForCall)
}

func (fake *FakeTextView) ClearCalls(stub func()) {
	fake.clearMutex.Lock()
	defer fake.clearMutex.Unlock()
	fake.ClearStub = stub
}

func (fake *FakeTextView) GetText() string {
	fake.getTextMutex.Lock()
	ret, specificReturn := fake.getTextReturnsOnCall[len(fake.getTextArgsForCall)]
	fake.getTextArgsForCall = append(fake.getTextArgsForCall, struct {
	}{})
	stub := fake.GetTextStub
	fakeReturns := fake.getTextReturns
	fake.recordInvocation("GetText", []interface{}{})
	fake.getTextMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTextView) GetTextCallCount() int {
	fake.getTextMutex.RLock()
	defer fake.getTextMutex.RUnlock()
	return len(fake.getTextArgsForCall)
}

func (fake *FakeTextView) GetTextCalls(stub func() string) {
	fake.getTextMutex.Lock()
	defer fake.getTextMutex.Unlock()
	fake.GetTextStub = stub
}

func (fake *FakeTextView) GetTextReturns(result1 string) {
	fake.getTextMutex.Lock()
	defer fake.getTextMutex.Unlock()
	fake.GetTextStub = nil
	fake.getTextReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTextView) GetTextReturnsOnCall(i int, result1 string) {
	fake.getTextMutex.Lock()
	defer fake.getTextMutex.Unlock()
	fake.GetTextStub = nil
	if fake.getTextReturnsOnCall == nil {
		fake.getTextReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getTextReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeTextView) ModifyPrimitive(arg1 func(t *tview.TextView)) {
	fake.modifyPrimitiveMutex.Lock()
	fake.modifyPrimitiveArgsForCall = append(fake.modifyPrimitiveArgsForCall, struct {
		arg1 func(t *tview.TextView)
	}{arg1})
	stub := fake.ModifyPrimitiveStub
	fake.recordInvocation("ModifyPrimitive", []interface{}{arg1})
	fake.modifyPrimitiveMutex.Unlock()
	if stub != nil {
		fake.ModifyPrimitiveStub(arg1)
	}
}

func (fake *FakeTextView) ModifyPrimitiveCallCount() int {
	fake.modifyPrimitiveMutex.RLock()
	defer fake.modifyPrimitiveMutex.RUnlock()
	return len(fake.modifyPrimitiveArgsForCall)
}

func (fake *FakeTextView) ModifyPrimitiveCalls(stub func(func(t *tview.TextView))) {
	fake.modifyPrimitiveMutex.Lock()
	defer fake.modifyPrimitiveMutex.Unlock()
	fake.ModifyPrimitiveStub = stub
}

func (fake *FakeTextView) ModifyPrimitiveArgsForCall(i int) func(t *tview.TextView) {
	fake.modifyPrimitiveMutex.RLock()
	defer fake.modifyPrimitiveMutex.RUnlock()
	argsForCall := fake.modifyPrimitiveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTextView) Primitive() tview.Primitive {
	fake.primitiveMutex.Lock()
	ret, specificReturn := fake.primitiveReturnsOnCall[len(fake.primitiveArgsForCall)]
	fake.primitiveArgsForCall = append(fake.primitiveArgsForCall, struct {
	}{})
	stub := fake.PrimitiveStub
	fakeReturns := fake.primitiveReturns
	fake.recordInvocation("Primitive", []interface{}{})
	fake.primitiveMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTextView) PrimitiveCallCount() int {
	fake.primitiveMutex.RLock()
	defer fake.primitiveMutex.RUnlock()
	return len(fake.primitiveArgsForCall)
}

func (fake *FakeTextView) PrimitiveCalls(stub func() tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = stub
}

func (fake *FakeTextView) PrimitiveReturns(result1 tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = nil
	fake.primitiveReturns = struct {
		result1 tview.Primitive
	}{result1}
}

func (fake *FakeTextView) PrimitiveReturnsOnCall(i int, result1 tview.Primitive) {
	fake.primitiveMutex.Lock()
	defer fake.primitiveMutex.Unlock()
	fake.PrimitiveStub = nil
	if fake.primitiveReturnsOnCall == nil {
		fake.primitiveReturnsOnCall = make(map[int]struct {
			result1 tview.Primitive
		})
	}
	fake.primitiveReturnsOnCall[i] = struct {
		result1 tview.Primitive
	}{result1}
}

func (fake *FakeTextView) SetText(arg1 string) {
	fake.setTextMutex.Lock()
	fake.setTextArgsForCall = append(fake.setTextArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetTextStub
	fake.recordInvocation("SetText", []interface{}{arg1})
	fake.setTextMutex.Unlock()
	if stub != nil {
		fake.SetTextStub(arg1)
	}
}

func (fake *FakeTextView) SetTextCallCount() int {
	fake.setTextMutex.RLock()
	defer fake.setTextMutex.RUnlock()
	return len(fake.setTextArgsForCall)
}

func (fake *FakeTextView) SetTextCalls(stub func(string)) {
	fake.setTextMutex.Lock()
	defer fake.setTextMutex.Unlock()
	fake.SetTextStub = stub
}

func (fake *FakeTextView) SetTextArgsForCall(i int) string {
	fake.setTextMutex.RLock()
	defer fake.setTextMutex.RUnlock()
	argsForCall := fake.setTextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTextView) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearMutex.RLock()
	defer fake.clearMutex.RUnlock()
	fake.getTextMutex.RLock()
	defer fake.getTextMutex.RUnlock()
	fake.modifyPrimitiveMutex.RLock()
	defer fake.modifyPrimitiveMutex.RUnlock()
	fake.primitiveMutex.RLock()
	defer fake.primitiveMutex.RUnlock()
	fake.setTextMutex.RLock()
	defer fake.setTextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTextView) recordInvocation(key string, args []interface{}) {
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

var _ component.TextView = new(FakeTextView)
