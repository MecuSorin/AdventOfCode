/* Author: Mecu Sorin       Phone: 0040747020102 */

package Day09

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Santa's travel", func() {

	Describe("Given puzzle.input as puzzle data", func() {
		BeforeEach(func() {
			isABetterSolutionBasedOnDistance = isMax
		})
		AfterEach(func() {
			isABetterSolutionBasedOnDistance = isMax
		})

		It("Puzzle 1 the shortest distance route", func() {
			cities := loadCitiesFromFile("puzzle.input")
			solution, ok := findShortestDistance(cities)
			Expect(ok).To(BeTrue())
			Expect(len(solution.visited)).To(Equal(len(cities)))
			Expect(solution.distance).To(Equal(uint(117)))
		})

		It("Puzzle 2 the longest distance route", func() {
			cities := loadCitiesFromFile("puzzle.input")
			isABetterSolutionBasedOnDistance = isMin
			solution, ok := findShortestDistance(cities)
			Expect(ok).To(BeTrue())
			Expect(len(solution.visited)).To(Equal(len(cities)))
			Expect(solution.distance).To(Equal(uint(909)))
		})
	})

})
