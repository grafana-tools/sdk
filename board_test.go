package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

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

func TestBoardHasTag_TagExists(t *testing.T) {
	b := NewBoard("Sample")
	b.AddTags("1", "2", "3")

	if !b.HasTag("2") {
		t.Error("tag exists but not found")
	}
}

func TestBoardHasTag_TagNotExists(t *testing.T) {
	b := NewBoard("Sample")
	b.AddTags("1", "2")

	if b.HasTag("3") {
		t.Error("tag not exists but found")
	}
}
