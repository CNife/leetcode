package insert_delete_getrandom_o1_duplicates_allowed

import "testing"

func TestRandomizedCollection(t *testing.T) {
	assert := func(value bool) {
		if !value {
			t.FailNow()
		}
	}

	rc := Constructor()
	assert(rc.Insert(1))
	assert(!rc.Insert(1))
	assert(rc.Insert(2))

	value := rc.GetRandom()
	assert(value == 1 || value == 2)

	assert(rc.Remove(1))
	value = rc.GetRandom()
	assert(value == 1 || value == 2)
}
