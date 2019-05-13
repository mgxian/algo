package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	hm := newHashMap(10)
	v, ok := hm.get("mgxian")
	assert.Equal(t, false, ok)
	assert.Equal(t, nil, v)

	hm.save("mgxian", "will")

	v, ok = hm.get("mgxian")
	assert.Equal(t, true, ok)
	assert.Equal(t, "will", v)

	hm.delete("mgxian")

	v, ok = hm.get("mgxian")
	assert.Equal(t, false, ok)
	assert.Equal(t, nil, v)
}
