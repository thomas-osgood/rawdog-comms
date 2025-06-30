package comms

import "bytes"

// structure of the metadata block that will
// be expected in a transmission.
type TcpHeader struct {
	Agentname string `json:"agentname" xml:"agentname"`
	Endpoint  int    `json:"endpoint" xml:"endpoint"`
	Addldata  string `json:"addldata" xml:"addldata"`
}

// structure designed to convey a status.
type TcpStatusMessage struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message,omitempty" xml:"message"`
}

// structure of a TCP transmission.
type TcpTransmission struct {
	// unsigned int representing the size of
	// the data/payload.
	DatSize uint64
	// data/payload value.
	Data *bytes.Buffer
	// unsigned short representing the size
	// of the metadata.
	MdSize uint16
	// metadata value.
	Metadata []byte
}
