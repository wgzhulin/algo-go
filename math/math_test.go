package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	type data struct {
		input1 int
		input2 int
		except int
	}

	testData := []data{
		{
			input1: 3,
			input2: 5,
			except: 1,
		},
		{
			input1: 8,
			input2: 2,
			except: 2,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, gcd(tdata.input1, tdata.input2))
	}
}


func TestLCM(t *testing.T) {
	type data struct {
		input1 int
		input2 int
		except int
	}

	testData := []data{
		{
			input1: 3,
			input2: 5,
			except: 15,
		},
		{
			input1: 8,
			input2: 2,
			except: 8,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, lcm(tdata.input1, tdata.input2))
	}
}