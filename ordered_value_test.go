package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedValue_Required(t *testing.T) {
	assert.Error(t, Ordered("name", "").Required().Err())
	assert.Error(t, Ordered("name", 0).Required().Err())
	assert.Error(t, Ordered("name", 0).Gt(1).Required().Err())
	assert.Nil(t, Ordered("name", "aha").Required().Err())
}

func TestOrderedValue_Gt(t *testing.T) {
	assert.Error(t, Ordered("age", 1).Gt(2).Err())
	assert.Error(t, Ordered("age", 0).Required().Gt(2).Err())
	assert.Nil(t, Ordered("age", 2).Gt(1).Err())
}

func TestOrderedValue_Gte(t *testing.T) {
	assert.Error(t, Ordered("age", 1).Gte(2).Err())
	assert.Error(t, Ordered("age", 0).Required().Gte(1).Err())
	assert.Nil(t, Ordered("age", 2).Gte(1).Err())
	assert.Nil(t, Ordered("age", 1).Gte(1).Err())
}

func TestOrderedValue_Lt(t *testing.T) {
	assert.Error(t, Ordered("age", 2).Lt(1).Err())
	assert.Error(t, Ordered("age", 0).Required().Lt(1).Err())
	assert.Nil(t, Ordered("age", 1).Lt(2).Err())
}

func TestOrderedValue_Lte(t *testing.T) {
	assert.Error(t, Ordered("age", 2).Lte(1).Err())
	assert.Error(t, Ordered("age", 0).Required().Lte(1).Err())
	assert.Nil(t, Ordered("age", 1).Lte(2).Err())
	assert.Nil(t, Ordered("age", 1).Lte(1).Err())
}

func TestOrderedValue_Include(t *testing.T) {
	assert.Error(t, Ordered("age", 2).IncludeBy(1, 3, 5).Err())
	assert.Error(t, Ordered("age", 0).Required().IncludeBy(1, 3, 5).Err())
	assert.Nil(t, Ordered("age", 3).IncludeBy(1, 3, 5).Err())
}

func TestOrderedValue_Exclude(t *testing.T) {
	assert.Nil(t, Ordered("age", 2).ExcludeBy(1, 3, 5).Err())
	assert.Error(t, Ordered("age", 0).Required().ExcludeBy(1, 3, 5).Err())
	assert.Error(t, Ordered("age", 3).ExcludeBy(1, 3, 5).Err())
}