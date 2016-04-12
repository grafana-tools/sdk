package grafana

import "testing"

func TestAddTags(t *testing.T) {
	b := NewBoard("Sample")

	b.AddTags("1", "2", "3")

	if len(b.Tags) != 3 {
		t.Errorf("len(tags) should be 3 but got %d", len(b.Tags))
	}
}

func TestBoardRemoveTags_Existent(t *testing.T) {
	b := NewBoard("Sample")
	b.AddTags("1", "2", "3", "4")

	b.RemoveTags("1", "2")

	if len(b.Tags) != 2 {
		t.Errorf("len(tags) should be 2 but got %d", len(b.Tags))
	}
}

func TestBoardRemoveTags_NonExistent(t *testing.T) {
	b := NewBoard("Sample")
	b.AddTags("1", "2")

	b.RemoveTags("3", "4")

	if len(b.Tags) != 2 {
		t.Errorf("len(tags) should be 2 but got %d", len(b.Tags))
	}
}

func TestBoardRemoveTags_WhenNoTags(t *testing.T) {
	b := NewBoard("Sample")

	b.RemoveTags("1", "2")

	if len(b.Tags) != 0 {
		t.Errorf("len(tags) should be 0 but got %d", len(b.Tags))
	}
}
