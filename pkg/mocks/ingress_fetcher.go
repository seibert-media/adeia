// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/seibert-media/k8s-ingress/pkg/domain"
)

type IngressFetcher struct {
	FetchStub        func() ([]domain.Domain, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct{}
	fetchReturns     struct {
		result1 []domain.Domain
		result2 error
	}
	fetchReturnsOnCall map[int]struct {
		result1 []domain.Domain
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IngressFetcher) Fetch() ([]domain.Domain, error) {
	fake.fetchMutex.Lock()
	ret, specificReturn := fake.fetchReturnsOnCall[len(fake.fetchArgsForCall)]
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct{}{})
	fake.recordInvocation("Fetch", []interface{}{})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchReturns.result1, fake.fetchReturns.result2
}

func (fake *IngressFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *IngressFetcher) FetchReturns(result1 []domain.Domain, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 []domain.Domain
		result2 error
	}{result1, result2}
}

func (fake *IngressFetcher) FetchReturnsOnCall(i int, result1 []domain.Domain, result2 error) {
	fake.FetchStub = nil
	if fake.fetchReturnsOnCall == nil {
		fake.fetchReturnsOnCall = make(map[int]struct {
			result1 []domain.Domain
			result2 error
		})
	}
	fake.fetchReturnsOnCall[i] = struct {
		result1 []domain.Domain
		result2 error
	}{result1, result2}
}

func (fake *IngressFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IngressFetcher) recordInvocation(key string, args []interface{}) {
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