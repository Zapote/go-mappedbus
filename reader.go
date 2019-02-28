package mappedbus

type Reader interface {
	//Open file for read
	Open() error
	//Timeout sets the read timeout
	Timeout(t int)
	//Next
	Next() (bool, error)
	//Type reads the message type
	Type() error
	//Message reads incoming message
	Message(m *Message) error
	//Close the reader
	Close() error
}

// NewReader constructs a new Reader.
// @param fileName the name of the memory mapped file
// @param fileSize the maximum size of the file
// @param recordSize the maximum size of a record (excluding status flags and meta data)

func NewReader(fileName string, fileSize int64, recordSize int64) Reader {
	return reader{
		timeout:    0,
		fileName:   fileName,
		fileSize:   fileSize,
		recordSize: recordSize,
	}
}

type reader struct {
	timeout    int
	fileName   string
	fileSize   int64
	recordSize int64
}

func (r reader) Open() error {
	return nil
}

func (r reader) Timeout(t int) {
	r.timeout = t
}

func (r reader) Next() (bool, error) {
	return false, nil
}

func (r reader) Type() error {
	return nil
}

func (r reader) Message(m *Message) error {
	return nil
}

func (r reader) Close() error {
	return nil
}
