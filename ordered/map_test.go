package ordered_test

import (
	"testing"

	"github.com/KrischanCS/go-toolbox/ordered"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	// Arrange & Act
	m := ordered.NewMap[string, int](10)

	// Assert
	assert.NotNil(t, m)
	assert.Equal(t, 0, m.Len())
}

func TestOrdered_Add(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)

	// Act
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	// Assert
	assert.Equal(t, 3, m.Len())
	assert.True(t, m.Contains("one"))
	assert.True(t, m.Contains("two"))
	assert.True(t, m.Contains("three"))
}

func TestOrdered_Add_UpdateExistingKey(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("key", 1)

	// Act
	m.Add("key", 2)

	// Assert
	assert.Equal(t, 1, m.Len())
	value, ok := m.Get("key")
	assert.True(t, ok)
	assert.Equal(t, 2, value)
}

func TestOrdered_Get(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)

	// Act & Assert
	value, ok := m.Get("one")
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	value, ok = m.Get("two")
	assert.True(t, ok)
	assert.Equal(t, 2, value)
}

func TestOrdered_Get_NotFound(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)

	// Act
	value, ok := m.Get("nonexistent")

	// Assert
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestOrdered_Delete(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	// Act
	value, ok := m.Delete("two")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, 2, value)
	assert.Equal(t, 2, m.Len())
	assert.False(t, m.Contains("two"))
}

func TestOrdered_Delete_First(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	// Act
	value, ok := m.Delete("one")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, 1, value)
	assert.Equal(t, 2, m.Len())

	// Verify order is preserved
	var keys []string
	for k := range m.All() {
		keys = append(keys, k)
	}
	assert.Equal(t, []string{"two", "three"}, keys)
}

func TestOrdered_Delete_Last(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	// Act
	value, ok := m.Delete("three")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, 3, value)
	assert.Equal(t, 2, m.Len())

	// Verify order is preserved
	var keys []string
	for k := range m.All() {
		keys = append(keys, k)
	}
	assert.Equal(t, []string{"one", "two"}, keys)
}

func TestOrdered_Delete_NotFound(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)

	// Act
	value, ok := m.Delete("nonexistent")

	// Assert
	assert.False(t, ok)
	assert.Equal(t, 0, value)
	assert.Equal(t, 1, m.Len())
}

func TestOrdered_Contains(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)

	// Act & Assert
	assert.True(t, m.Contains("one"))
	assert.False(t, m.Contains("two"))
}

func TestOrdered_Len(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)

	// Assert initial length
	assert.Equal(t, 0, m.Len())

	// Act & Assert after adding
	m.Add("one", 1)
	assert.Equal(t, 1, m.Len())

	m.Add("two", 2)
	assert.Equal(t, 2, m.Len())

	// Act & Assert after deleting
	m.Delete("one")
	assert.Equal(t, 1, m.Len())
}

func TestOrdered_All_PreservesInsertionOrder(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("first", 1)
	m.Add("second", 2)
	m.Add("third", 3)
	m.Add("fourth", 4)

	// Act
	var keys []string
	var values []int
	for k, v := range m.All() {
		keys = append(keys, k)
		values = append(values, v)
	}

	// Assert
	assert.Equal(t, []string{"first", "second", "third", "fourth"}, keys)
	assert.Equal(t, []int{1, 2, 3, 4}, values)
}

func TestOrdered_All_Empty(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)

	// Act
	var count int
	for range m.All() {
		count++
	}

	// Assert
	assert.Equal(t, 0, count)
}

func TestOrdered_All_EarlyBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	// Act
	var keys []string
	for k := range m.All() {
		keys = append(keys, k)
		if k == "two" {
			break
		}
	}

	// Assert
	assert.Equal(t, []string{"one", "two"}, keys)
}

func TestOrdered_UpdatePreservesOrder(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("first", 1)
	m.Add("second", 2)
	m.Add("third", 3)

	// Act - update middle element
	m.Add("second", 20)

	// Assert - order should be preserved
	var keys []string
	var values []int
	for k, v := range m.All() {
		keys = append(keys, k)
		values = append(values, v)
	}
	assert.Equal(t, []string{"first", "second", "third"}, keys)
	assert.Equal(t, []int{1, 20, 3}, values)
}

func TestOrdered_DeleteAndReAdd(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("first", 1)
	m.Add("second", 2)
	m.Add("third", 3)

	// Act - delete and re-add
	m.Delete("second")
	m.Add("second", 22)

	// Assert - re-added element should be at the end
	var keys []string
	for k := range m.All() {
		keys = append(keys, k)
	}
	assert.Equal(t, []string{"first", "third", "second"}, keys)
}

func TestOrdered_DeleteAll(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	// Act
	m.Delete("one")
	m.Delete("two")
	m.Delete("three")

	// Assert
	assert.Equal(t, 0, m.Len())

	var count int
	for range m.All() {
		count++
	}
	assert.Equal(t, 0, count)
}

func TestOrdered_AddAfterDeleteAll(t *testing.T) {
	t.Parallel()

	// Arrange
	m := ordered.NewMap[string, int](10)
	m.Add("one", 1)
	m.Delete("one")

	// Act
	m.Add("new", 100)

	// Assert
	assert.Equal(t, 1, m.Len())
	value, ok := m.Get("new")
	assert.True(t, ok)
	assert.Equal(t, 100, value)
}
