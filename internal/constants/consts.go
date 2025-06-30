package constants

// helper constant defining a nullbyte
const NULL_BYTE string = "\x00"

// size of the byte block that will be
// used to read data.
const SZ_DATABUFF uint64 = (1 << 11)

// maximum size of the metadata block.
const SZ_METADATA_MAX uint16 = (1 << 16) - 1

// size of the data size block in the transmission.
const SZ_SIZEBLOCK_DAT uint64 = (1 << 3)

// size of the metadata size block in the transmission.
const SZ_SIZEBLOCK_MD uint64 = (1 << 1)
