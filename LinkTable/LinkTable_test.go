package LinkTable

import "testing"

func TestLinkTable(t *testing.T) {
	nums := []int{2, 1, 2, 2, 5, 3, 5, 7}
	lt := NewLinkTable()
	if lt.ToString() != "LinkTable: [  ]" {
		t.Error(lt.ToString())
	}

	// Insert
	for _, num := range nums {
		lt.InsertVal(num)
	}
	if lt.ToString() != "LinkTable: [ 2, 1, 2, 2, 5, 3, 5, 7 ]" {
		t.Error(lt.ToString())
	}
	if lt.Size() != 8 {
		t.Error(lt.Size())

	}

	// Delete
	head, _ := lt.Head()
	err := lt.Delete(head)
	if (err != nil) || (lt.ToString() != "LinkTable: [ 1, 2, 2, 5, 3, 5, 7 ]") {
		t.Error(lt.ToString())
	}
	err = lt.DeleteFirst(5)
	if (err != nil) || (lt.ToString() != "LinkTable: [ 1, 2, 2, 3, 5, 7 ]") {
		t.Error(lt.ToString())
	}

	// Find
	node1, _ := lt.FindFirst(2)
	node2, _ := lt.Head()
	if node1 != node2.next {
		t.Error(node1)
	}

	// Clear
	lt.Clear()
	if lt.Size() != 0 {
		t.Error(lt.Size())
	}
}
