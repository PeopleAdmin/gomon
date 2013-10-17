package gomon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {

	Describe("calculateStats", func() {
		var slice []float64

		Context("called with a slice of all zeros", func(){
			It("should return the appropriate stats", func(){
				slice = []float64{0.0}
				allZero, min, max, sum := calculateStats(slice)
				Expect(allZero).To(BeTrue())
				Expect(min).To(BeZero())
				Expect(max).To(BeZero())
				Expect(sum).To(BeZero())
			})
		})

		Context("with an empty slice", func(){
			It("should panic", func(){
				slice = []float64{}
				Expect(func(){ calculateStats(slice) }).To(Panic())
			})
		})

	})

})
