package mappedbus

type Writer interface {
	//Open file for write
	Open() error
	//Write message to bus
	Write(m Message) error
	//WriteBytes to bus. Starting from off and writes to n.
	WriteBytes(data []byte, off int64, n int64)
	//Close the writer
	Close() error
}

//Message for writing on the bus
type Message struct {
}

type writer struct {
}

func (w *writer) Open() error {
	return nil
}

func (w *writer) Write(m Message) error {
	return nil
}

func (w *writer) WriteBytes(data []byte, off int64, n int64) error {
	return nil
}

func (w *writer) Close() error {
	return nil
}
