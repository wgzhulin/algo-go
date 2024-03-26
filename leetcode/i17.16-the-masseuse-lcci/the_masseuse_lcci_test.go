package i1716themasseuselcci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMassage(t *testing.T) {
	type data struct {
		input  []int
		except int
	}

	testData := []data{
		{
			input:  []int{1,2,3,1},
			except: 4,
		},
		{
			input:  []int{2,7,9,3,1},
			except: 12,
		},
		{
			input:  []int{2,1,4,5,3,1,1,3},
			except: 12,
		},
	}


	for _, tdata := range testData {
		assert.Equal(t, tdata.except, massage(tdata.input))
	}
}