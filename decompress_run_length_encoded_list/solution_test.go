package decompress_run_length_encoded_list

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_DecompressRLElist(t *testing.T) {
	var input, output = []int{1, 2, 3, 4}, []int{2, 4, 4, 4}

	if reflect.DeepEqual(decompressRLElist(input), output) {
		t.Log("Pass")
	} else {
		t.Fail()
	}

	input, output = []int{1, 2, 3, 4, 4, 5, 6, 7}, []int{2, 4, 4, 4, 5, 5, 5, 5, 7, 7, 7, 7, 7, 7}
	assert.Equal(t, decompressRLElist(input), output)
}
