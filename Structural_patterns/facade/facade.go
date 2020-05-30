package facade

import (
	"fmt"
	"path/filepath"
	"strings"
)

type VideoFile struct {
	Filename string
	Format   string
	Buffer   []byte
}

func NewVideoFile(filename string) *VideoFile {
	format := filepath.Ext(filename)
	if len(format) <= 0 {
		return nil
	}

	format = format[1:] // remove dot
	return &VideoFile{
		Filename: filename,
		Format:   format,
		Buffer:   []byte("fake content"),
	}
}

func (f *VideoFile) Save(path string) {
	fmt.Printf("Saving video file into %s\n", path)
}

type Codec interface {
	Extract(file *VideoFile) []byte
	Compress(data []byte) []byte
}

type MP4CompressionCodec struct{}

func (codec *MP4CompressionCodec) Extract(file *VideoFile) []byte {
	fmt.Println("extract mp4 video")
	return file.Buffer
}

func (codec *MP4CompressionCodec) Compress(data []byte) []byte {
	// compress file data into mp4 format
	fmt.Println("compress data into mp4 format")
	return []byte("mp4 compress file")
}

func NewMP4CompressionCodec() *MP4CompressionCodec {
	return &MP4CompressionCodec{}
}

type AVICompressionCodec struct{}

func (codec *AVICompressionCodec) Extract(file *VideoFile) []byte {
	fmt.Println("extract avi video")
	return []byte("file")
}

func (codec *AVICompressionCodec) Compress(data []byte) []byte {
	fmt.Println("compress data into avi format")
	return []byte("avi compress file")
}

func NewAVICompressionCodec() *AVICompressionCodec {
	return &AVICompressionCodec{}
}

type CodecFactory struct{}

func (c *CodecFactory) GetCodec(format string) Codec {
	switch strings.ToLower(format) {
	case "mp4":
		return NewMP4CompressionCodec()
	case "avi":
		return NewAVICompressionCodec()
	default:
		return NewAVICompressionCodec()
	}
}

func NewCodecFactory() *CodecFactory {
	return &CodecFactory{}
}

type VideoFileReader struct{}

func (v *VideoFileReader) Read(file *VideoFile, codec Codec) []byte {
	return codec.Extract(file)
}

func (v *VideoFileReader) Convert(data []byte, codec Codec) []byte {
	codec.Compress(data)
	return []byte("result")
}

func NewVideoFileReader() *VideoFileReader {
	return &VideoFileReader{}
}

// 为了将框架的复杂性隐藏在一个简单接口背后，我们创建了一个外观类VideoConverter。它是在
// 功能性和简洁性之间做出的权衡。
type VideoConverter struct{}
func (v *VideoConverter) Convert(filename string, format string) *VideoFile {
	file := NewVideoFile(filename)

	codecFactory := NewCodecFactory()
	sourceCodec := codecFactory.GetCodec(file.Format)
	destinationCodec := codecFactory.GetCodec(format)

	reader := NewVideoFileReader()
	buffer := reader.Read(file, sourceCodec)
	result := reader.Convert(buffer, destinationCodec)

	return &VideoFile{
		Filename: strings.TrimSuffix(file.Filename, file.Format) + "." + format,
		Format:   format,
		Buffer:   result,
	}
}