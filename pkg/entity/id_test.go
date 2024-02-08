package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEntityUuid(t *testing.T) {
	id := NewID()

	assert.NotNil(t, id)
}

func TestParseId(t *testing.T) {
	idString := "cf8010b6-faef-4911-88ab-2117a14deed9"

	id, err := ParseID(idString)

	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, id.String(), idString)
}
