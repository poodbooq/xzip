package main

import (
	"bytes"
	"log"
	"os"
)

/*
ZIP format
Byte order: Little-endian

Overall zipfile format:
[Local file header + Compressed data [+ Extended local header]?]*
[Central directory]*
[End of central directory record]

Local file header:
Offset   Length   Contents
  0      4 bytes  Local file header signature (0x04034b50)
  4      2 bytes  Version needed to extract
  6      2 bytes  General purpose bit flag
  8      2 bytes  Compression method
 10      2 bytes  Last mod file time
 12      2 bytes  Last mod file date
 14      4 bytes  CRC-32
 18      4 bytes  Compressed size (n)
 22      4 bytes  Uncompressed size
 26      2 bytes  Filename length (f)
 28      2 bytes  Extra field length (e)
        (f)bytes  Filename
        (e)bytes  Extra field
        (n)bytes  Compressed data

Extended local header:
Offset   Length   Contents
  0      4 bytes  Extended Local file header signature (0x08074b50)
  4      4 bytes  CRC-32
  8      4 bytes  Compressed size
 12      4 bytes  Uncompressed size

Central directory:
Offset   Length   Contents
  0      4 bytes  Central file header signature (0x02014b50)
  4      2 bytes  Version made by
  6      2 bytes  Version needed to extract
  8      2 bytes  General purpose bit flag
 10      2 bytes  Compression method
 12      2 bytes  Last mod file time
 14      2 bytes  Last mod file date
 16      4 bytes  CRC-32
 20      4 bytes  Compressed size
 24      4 bytes  Uncompressed size
 28      2 bytes  Filename length (f)
 30      2 bytes  Extra field length (e)
 32      2 bytes  File comment length (c)
 34      2 bytes  Disk number start
 36      2 bytes  Internal file attributes
 38      4 bytes  External file attributes
 42      4 bytes  Relative offset of local header
 46     (f)bytes  Filename
        (e)bytes  Extra field
        (c)bytes  File comment

End of central directory record:
Offset   Length   Contents
  0      4 bytes  End of central dir signature (0x06054b50)
  4      2 bytes  Number of this disk
  6      2 bytes  Number of the disk with the start of the central directory
  8      2 bytes  Total number of entries in the central dir on this disk
 10      2 bytes  Total number of entries in the central dir
 12      4 bytes  Size of the central directory
 16      4 bytes  Offset of start of central directory with respect to the starting disk number
 20      2 bytes  zipfile comment length (c)
 22     (c)bytes  zipfile comment




compression method: (2 bytes)
          0 - The file is stored (no compression)
          1 - The file is Shrunk
          2 - The file is Reduced with compression factor 1
          3 - The file is Reduced with compression factor 2
          4 - The file is Reduced with compression factor 3
          5 - The file is Reduced with compression factor 4
          6 - The file is Imploded
          7 - Reserved for Tokenizing compression algorithm
          8 - The file is Deflated

*/

const (
	// Local File Headers
	localFileHeaderSignature = 0x04034b50

	// Extended Local Headers
	extLocalFileHeaderSignature = 0x08074b50

	// Central Directory Headers
	centralFileHeaderSignature = 0x02014b50

	// End Of Central Directory Record Headers
	endOfCentralDirSignature = 0x06054b50
)

type compression uint8

const (
	_ compression = iota
	noCompression
	fileIsShrunk
	factor1
	factor2
	factor3
	factor4
	imploded
	tokenizing
	deflated
)

// ZipStructure ...
type ZipStructure struct {
	localFileHeader       []*localFileHeader
	extendedLocalHeader   []*extendedLocalHeader
	centralDirectory      []*centralDirectory
	endOfCentralDirRecord *endOfCentralDirRecord
}

type localFileHeader struct {
	signature        uint16
	version          uint8
	bitFlag          uint8
	compression      compression
	lastModFileTime  uint8
	lastModFileDate  uint8
	crc32            uint16
	compressedSize   uint16
	uncompressedSize uint16
	filenameLen      uint8
	extraFieldLen    uint8
	filename         []byte
	extraField       []byte
	compressedData   []byte
}
type extendedLocalHeader struct {
	signature        uint16
	crc32            uint16
	compressedSize   uint16
	uncompressedSize uint16
}
type centralDirectory struct {
	signature              uint16
	versionMadeBy          uint8
	versionNeededToExtract uint8
	bitFlag                uint8
	compression            compression
	lastModFileTime        uint8
	lastModFileDate        uint8
	crc32                  uint16
	compressedSize         uint16
	uncompressedSize       uint16
	filenameLen            uint8
	extraFieldLen          uint8
	fileCommentLen         uint8
	diskNumStart           uint8
	internalFileAttrs      uint8
	externalFileAttrs      uint16
	offsetOfLocalHeader    uint16
	filename               []byte
	extraField             []byte
	fileComment            []byte
}
type endOfCentralDirRecord struct {
	signature                           uint16
	numberOfDisk                        uint8
	numberOfDiskWithStartOfCentralDir   uint8
	totalNumOfEntriesInCentralDirOnDisk uint8
	totalNumOfEntriesInCentralDir       uint8
	sizeOfCentralDir                    uint16
	offsetOfStartCentralDir             uint16
	zipFileCommentLen                   uint8
	zipFileComment                      []byte
}

func main() {
	f, err := os.Create("test333.zip")
	if err != nil {
		log.Fatal(err)
	}
	var buf = new(bytes.Buffer)

	// Write local file header signature 4 bytes
	buf.WriteByte(0x50)
	buf.WriteByte(0x4b)
	buf.WriteByte(0x03)
	buf.WriteByte(0x04)

	//Write version 2 bytes
	buf.WriteByte(0x00)
	buf.WriteByte(0x00)

	//General purpose bit flag 2 bytes
	buf.WriteByte(0x00)
	buf.WriteByte(0x00)

	//Compression method 2 bytes
	buf.WriteByte(0x00)
	buf.WriteByte(0x00)

	//Last mod file time 2 bytes
	buf.WriteByte(0x00)
	buf.WriteByte(0x00)

	// Last mod file date 2 bytes
	buf.WriteByte(0x00)
	buf.WriteByte(0x00)

	f.Write(buf.Bytes())
	f.Close()
}
