// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/seibert-media/k8s-ingress/mocks"
	"testing"
)

func TestSyncer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Syncer Suite")
}

var _ = Describe("Syncer", func() {
	var (
		fetcher *mocks.IngressFetcher
		applier *mocks.DomainApplier
		syncer  *Syncer
	)

	BeforeEach(func() {
		fetcher = &mocks.IngressFetcher{}
		applier = &mocks.DomainApplier{}
		syncer = &Syncer{
			Fetcher: fetcher,
			Applier: applier,
		}
	})

	Describe("Sync", func() {
		It("calls ingress fetcher", func() {
			Expect(fetcher.FetchCallCount()).To(Equal(0))
			Expect(syncer.Sync()).To(BeNil())
			Expect(fetcher.FetchCallCount()).To(Equal(1))
		})
		It("return error when fetch fails", func() {
			fetcher.FetchReturns(nil, errors.New("Failed"))
			Expect(syncer.Sync()).NotTo(BeNil())
		})
		It("calls applier", func() {
			Expect(applier.ApplyCallCount()).To(Equal(0))
			syncer.Sync()
			Expect(applier.ApplyCallCount()).To(Equal(1))
		})
		It("does not apply if fetch fails", func() {
			fetcher.FetchReturns(nil, errors.New("Failed"))
			Expect(applier.ApplyCallCount()).To(Equal(0))
			syncer.Sync()
			Expect(applier.ApplyCallCount()).To(Equal(0))
		})
	})
})