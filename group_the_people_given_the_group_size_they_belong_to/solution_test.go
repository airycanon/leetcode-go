package group_the_people_given_the_group_size_they_belong_to

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GroupThePeople(t *testing.T) {
	groupSizes := []int{2, 1, 3, 3, 3, 2}
	assert.ElementsMatch(t, groupThePeople(groupSizes), [][]int{[]int{1}, []int{0, 5}, []int{2, 3, 4}})

	groupSizes = []int{3, 3, 3, 3, 3, 1, 3}
	assert.ElementsMatch(t, groupThePeople(groupSizes), [][]int{[]int{0, 1, 2}, []int{5}, []int{3, 4, 6}})
}
