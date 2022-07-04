package tourcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidTourCode(t *testing.T) {
	assert.True(t, IsValidTourCode(PgaTour))
	assert.True(t, IsValidTourCode(KornFerryTour))
	assert.True(t, IsValidTourCode(PgaTourChampions))
	assert.True(t, IsValidTourCode(PgaTourCanada))
	assert.True(t, IsValidTourCode(PgaTourLatinAmerica))
}

func TestIsValidTourCode_NotValid(t *testing.T) {
	assert.False(t, IsValidTourCode("q"))
}

func TestIsValidTourCodeStr(t *testing.T) {
	assert.True(t, IsValidTourCodeStr(PgaTour))
	assert.True(t, IsValidTourCodeStr(KornFerryTour))
	assert.True(t, IsValidTourCodeStr(PgaTourChampions))
	assert.True(t, IsValidTourCodeStr(PgaTourCanada))
	assert.True(t, IsValidTourCodeStr(PgaTourLatinAmerica))
}

func TestIsValidTourCodeStr_NotValid(t *testing.T) {
	assert.False(t, IsValidTourCodeStr("q"))
}
