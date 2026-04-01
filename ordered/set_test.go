package ordered_test

import (
	"testing"

	"github.com/KrischanCS/go-toolbox/ordered"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	t.Parallel()

	// Arrange & Act
	s := ordered.NewSet[string](10)

	// Assert
	assert.NotNil(t, s)
	assert.Equal(t, 0, s.Len())
	assert.True(t, s.IsEmpty())
}

func TestSetOf(t *testing.T) {
	t.Parallel()

	// Arrange & Act
	s := ordered.SetOf("one", "two", "three")

	// Assert
	assert.NotNil(t, s)
	assert.Equal(t, 3, s.Len())
	assert.False(t, s.IsEmpty())
	assert.True(t, s.Contains("one"))
	assert.True(t, s.Contains("two"))
	assert.True(t, s.Contains("three"))
}

func TestSetOf_WithDuplicates(t *testing.T) {
	t.Parallel()

	// Arrange & Act
	s := ordered.SetOf("one", "two", "one", "three", "two")

	// Assert
	assert.Equal(t, 3, s.Len())
	assert.ElementsMatch(t, []string{"one", "two", "three"}, s.Values())
}

func TestSet_Add(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[string](10)

	// Act
	s.Add("one")
	s.Add("two")
	s.Add("three")

	// Assert
	assert.Equal(t, 3, s.Len())
	assert.True(t, s.Contains("one"))
	assert.True(t, s.Contains("two"))
	assert.True(t, s.Contains("three"))
}

func TestSet_Add_Duplicate(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[string](10)
	s.Add("one")

	// Act
	s.Add("one")

	// Assert
	assert.Equal(t, 1, s.Len())
}

func TestSet_Remove(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	ok := s.Remove("two")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, 2, s.Len())
	assert.True(t, s.Contains("one"))
	assert.False(t, s.Contains("two"))
	assert.True(t, s.Contains("three"))
}

func TestSet_Remove_NonExistent(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	ok := s.Remove("four")

	// Assert
	assert.False(t, ok)
	assert.Equal(t, 3, s.Len())
}

func TestSet_Remove_First(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	ok := s.Remove("one")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, []string{"two", "three"}, s.Values())
}

func TestSet_Remove_Last(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	ok := s.Remove("three")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, []string{"one", "two"}, s.Values())
}

func TestSet_Remove_Middle(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	ok := s.Remove("two")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, []string{"one", "three"}, s.Values())
}

func TestSet_Remove_OnlyElement(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one")

	// Act
	ok := s.Remove("one")

	// Assert
	assert.True(t, ok)
	assert.Equal(t, 0, s.Len())
	assert.True(t, s.IsEmpty())
}

func TestSet_Contains(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act & Assert
	assert.True(t, s.Contains("one"))
	assert.True(t, s.Contains("two"))
	assert.True(t, s.Contains("three"))
	assert.False(t, s.Contains("four"))
}

func TestSet_Len(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[string](10)

	// Assert initial
	assert.Equal(t, 0, s.Len())

	// Act & Assert after adding
	s.Add("one")
	assert.Equal(t, 1, s.Len())

	s.Add("two")
	assert.Equal(t, 2, s.Len())

	// Act & Assert after removing
	s.Remove("one")
	assert.Equal(t, 1, s.Len())
}

func TestSet_IsEmpty(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[string](10)

	// Assert initial
	assert.True(t, s.IsEmpty())

	// Act & Assert after adding
	s.Add("one")
	assert.False(t, s.IsEmpty())

	// Act & Assert after removing
	s.Remove("one")
	assert.True(t, s.IsEmpty())
}

func TestSet_Clear(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	s.Clear()

	// Assert
	assert.Equal(t, 0, s.Len())
	assert.True(t, s.IsEmpty())
	assert.False(t, s.Contains("one"))
}

func TestSet_Clone(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	clone := s.Clone()

	// Assert
	assert.Equal(t, s.Len(), clone.Len())
	assert.Equal(t, s.Values(), clone.Values())

	// Verify independence
	s.Add("four")
	assert.Equal(t, 4, s.Len())
	assert.Equal(t, 3, clone.Len())
}

func TestSet_Values(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	values := s.Values()

	// Assert - values should be in insertion order
	assert.Equal(t, []string{"one", "two", "three"}, values)
}

func TestSet_Values_Empty(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[string](10)

	// Act
	values := s.Values()

	// Assert
	assert.Empty(t, values)
}

func TestSet_All(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	var values []string
	for v := range s.All() {
		values = append(values, v)
	}

	// Assert - values should be in insertion order
	assert.Equal(t, []string{"one", "two", "three"}, values)
}

func TestSet_All_Empty(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[string](10)

	// Act
	var values []string
	for v := range s.All() {
		values = append(values, v)
	}

	// Assert
	assert.Empty(t, values)
}

func TestSet_All_EarlyBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three", "four", "five")

	// Act
	var values []string
	for v := range s.All() {
		values = append(values, v)
		if len(values) == 2 {
			break
		}
	}

	// Assert
	assert.Equal(t, []string{"one", "two"}, values)
}

func TestSet_String(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf("one", "two", "three")

	// Act
	str := s.String()

	// Assert
	assert.Equal(t, "(OrderedSet[string]: [one, two, three])", str)
}

func TestSet_String_Empty(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[string](10)

	// Act
	str := s.String()

	// Assert
	assert.Equal(t, "(OrderedSet[string]: <empty>)", str)
}

func TestSet_String_Int(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf(1, 2, 3)

	// Act
	str := s.String()

	// Assert
	assert.Equal(t, "(OrderedSet[int]: [1, 2, 3])", str)
}

func TestSet_MaintainsInsertionOrder(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.NewSet[int](10)

	// Act - add in specific order
	s.Add(5)
	s.Add(3)
	s.Add(8)
	s.Add(1)
	s.Add(9)

	// Assert - should maintain insertion order, not sorted order
	assert.Equal(t, []int{5, 3, 8, 1, 9}, s.Values())
}

func TestSet_MaintainsOrderAfterRemove(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf(1, 2, 3, 4, 5)

	// Act - remove middle element
	s.Remove(3)

	// Assert - order should be maintained
	assert.Equal(t, []int{1, 2, 4, 5}, s.Values())
}

func TestSet_AddAfterRemove(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf(1, 2, 3)
	s.Remove(2)

	// Act - add new element
	s.Add(4)

	// Assert - new element should be at the end
	assert.Equal(t, []int{1, 3, 4}, s.Values())
}

func TestSet_ReAddRemovedElement(t *testing.T) {
	t.Parallel()

	// Arrange
	s := ordered.SetOf(1, 2, 3)
	s.Remove(2)

	// Act - re-add removed element
	s.Add(2)

	// Assert - re-added element should be at the end
	assert.Equal(t, []int{1, 3, 2}, s.Values())
}
