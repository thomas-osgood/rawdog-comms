package messages

// error to display when the size of the metdata
// is larger than the maximum allowed metadata length.
const ERR_MD_LARGE string = "metadata size of %d bytes is too large. max size is %d bytes"

// error to display if the context gets cancelled
// during a recv.
const ERR_READ_CANCELLED string = "read has been cancelled"

// error to display if the timeout is reached
// during a recv.
const ERR_READ_TIMEOUT string = "read transmission timeout"

// error to display if the context gets cancelled
// during a send.
const ERR_SEND_CANCELLED string = "send has been cancelled"

// error to display if the timeout is reached
// during a send.
const ERR_SEND_TIMEOUT string = "read transmission timeout"

// error to display if the provided timeout is
// an invalid value/less than one.
const ERR_TIMEOUT_LT_ONE string = "timeout must be > 0"
