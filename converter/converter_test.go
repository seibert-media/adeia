// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seibert-media/k8s-ingress/model"
)

func TestSyncer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}

var _ = Describe("Fetcher", func() {
	var (
		domainConverter *Converter
		Domains =  []model.Domain{"http://server1.com/", "http://server2.com/"}
	)

	BeforeEach(func() {
		domainConverter = &Converter{
			Domains : Domains,
			}
	})

	Describe("Convert", func() {
		It("returns no error", func() {
			_, err := domainConverter.Convert()
			Expect(err).To(BeNil())
		})
		It("returns correct count of ingress objects", func() {
			ingress, _ := domainConverter.Convert()
			Expect(ingress).ToNot(BeNil())
			Expect(ingress.Spec.Rules).To(HaveLen(2))
		})
		It("returns ingress objects with correct host", func() {
			ingress, _ := domainConverter.Convert()
			Expect(ingress.Spec.Rules[0].Host).To(Equal("http://server1.com/"))
			Expect(ingress.Spec.Rules[1].Host).To(Equal("http://server2.com/"))
		})
	})
})