package ethconvert

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFromEtherToWei(t *testing.T) {
	d, _ := decimal.NewFromString("1.23456789")
	wei, err := ToWei(d, "ether")
	assert.Nil(t, err)
	assert.Equal(t, "1234567890000000000", wei.String())
}

func TestInvalidUnit(t *testing.T) {
	_, err := ToWei(decimal.Zero, "invalid")
	assert.NotNil(t, err)
}

func TestFromWeiToWei(t *testing.T) {
	d, _ := decimal.NewFromString("1")
	wei, err := ToWei(d, "wei")
	assert.Nil(t, err)
	assert.Equal(t, "1", wei.String())
}

func TestFromWeiToEther(t *testing.T) {
	d, _ := decimal.NewFromString("1000000000000000001")
	ether, err := FromWei(d, "ether")
	t.Log(ether)
	assert.Nil(t, err)
	assert.Equal(t, "1.000000000000000001", ether.String())
}

func TestConvert(t *testing.T) {
	d := decimal.NewFromInt(50)
	t.Log(Convert(d, Ether, Gwei))
}
