package gomon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {

	Describe("calculateStats", func() {
		var slice []float64

		Context("called with a slice of all zeros", func() {
			slice = []float64{0.0}
			It("should return the appropriate stats", func() {
				allZero, min, max, sum := calculateStats(slice)
				Expect(allZero).To(BeTrue())
				Expect(min).To(BeZero())
				Expect(max).To(BeZero())
				Expect(sum).To(BeZero())
			})
		})

		Context("with an empty slice", func() {
			It("should panic", func() {
				slice = []float64{}
				Expect(func() { calculateStats(slice) }).To(Panic())
			})
		})

	})

	Describe("deltaSinceLastCall", func() {
		var emitter func() float64
		emitter = Fib()

		Context("with a function that changes", func() {
			It("reflects the changes in successive calls", func() {
				delta := DeltaSinceLastCall(emitter)
				Expect(delta()).To(Equal(0.0))
				Expect(delta()).To(Equal(1.0))
				Expect(delta()).To(Equal(1.0))
				Expect(delta()).To(Equal(2.0))
			})
		})

		Context("with a function that does not change", func() {
			It("always shows zero delta", func() {
				delta := DeltaSinceLastCall(func() float64 { return 1.0 })
				Expect(delta()).To(Equal(0.0))
				Expect(delta()).To(Equal(0.0))
				Expect(delta()).To(Equal(0.0))
			})
		})

	})

})

// When called successively, returns the next term in the Fibonacci sequence.
func Fib() (generatorFunc func() float64) {
	a, b := 0.0, 1.0
	generatorFunc = func() float64 {
		a, b = b, a+b
		return a
	}
	return
}
