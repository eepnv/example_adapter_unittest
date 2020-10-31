package layers

import (
	"testing"
)

// a FAKE ONE!!!
type FakeDatabaseImpl struct{}

func (f FakeDatabaseImpl) ReadNotepad(id string) (*Notepad, error) {
	notepad := Notepad{
		ID:      id,
		Content: "content",
	}
	return &notepad, nil
}
func (f FakeDatabaseImpl) UpdateNotepad(notepad Notepad) error {
	return nil
}

func TestRead(t *testing.T) {
	notepadService := NotepadImpl{
		dbService: FakeDatabaseImpl{},
	}
	val, err := notepadService.ReadNotepad("1")
	if err != nil {
		t.Fatal(err)
	}
	if val.Content != "content" {
		t.Fatal("unexpected!")
	}
}
