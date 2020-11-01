package layers

/*
	There is no import in this layer!!! this represents your core business logic,
	and therefore you can easily replace your frontend code (maybe you use rpc / grpc / websocket
	or any other future things?) and backend code (replacing database interface, 
	will be mentioned at end of file)
*/

// section 1: core data model

//Notepad ...
type Notepad struct {
	ID      string
	Content string
}

// ******************************************************************
// section 2: functions considering a data model

/*
	define a service about a model. no special meaning in using interface here,
	except it provides a short table_of_content-like meaning here.
*/
type NotepadService interface {
	ReadNotepad(id string) (*Notepad, error)
	WriteNotepad(notepad Notepad) error
	AppendToOtherNotepad(idFrom string, idTo string) error
}
type NotepadImpl struct {
	dbService DatabaseService
}

func (s NotepadImpl) ReadNotepad(id string) (*Notepad, error) {
	return s.dbService.ReadNotepad(id)
}
func (s NotepadImpl) WriteNotepad(notepad Notepad) error {
	return s.dbService.UpdateNotepad(notepad)
}
func (s NotepadImpl) AppendToOtherNotepad(idFrom string, idTo string) error {
	fromNotepad, err := s.ReadNotepad(idFrom)
	if err != nil {
		return err
	}
	toNotepad, err := s.ReadNotepad(idTo)
	if err != nil {
		return err
	}
	toNotepad.Content = toNotepad.Content + fromNotepad.Content
	if err := s.WriteNotepad(*toNotepad); err != nil {
		return err
	}
	return nil
}

// this line is just simple type checking to see if i fit the struct into the interface
var _ NotepadService = NotepadImpl{}

// ******************************************************************
// section 3: define the interface that u will need your database do

// no need to write real class, so you can replace the real class when you test
type DatabaseService interface {
	ReadNotepad(id string) (*Notepad, error)
	UpdateNotepad(notepad Notepad) error
}
