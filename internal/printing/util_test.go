package printing_test

import (
	"testing"

	xssfox "github.com/JGPatelOfficial/xssfox/internal/printing"
	"github.com/stretchr/testify/assert"
)

func TestCheckToShowPoC(t *testing.T) {
	t1g, t1r, t1v := xssfox.CheckToShowPoC("g")
	assert.Equal(t, t1g, true)
	assert.Equal(t, t1r, false)
	assert.Equal(t, t1v, false)

	t2g, t2r, t2v := xssfox.CheckToShowPoC("r")
	assert.Equal(t, t2g, false)
	assert.Equal(t, t2r, true)
	assert.Equal(t, t2v, false)

	t3g, t3r, t3v := xssfox.CheckToShowPoC("v")
	assert.Equal(t, t3g, false)
	assert.Equal(t, t3r, false)
	assert.Equal(t, t3v, true)

	t4g, t4r, t4v := xssfox.CheckToShowPoC("r,v")
	assert.Equal(t, t4g, false)
	assert.Equal(t, t4r, true)
	assert.Equal(t, t4v, true)

	t5g, t5r, t5v := xssfox.CheckToShowPoC("g,r,v")
	assert.Equal(t, t5g, true)
	assert.Equal(t, t5r, true)
	assert.Equal(t, t5v, true)
}
