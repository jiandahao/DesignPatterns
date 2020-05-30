package decorator

import (
	"fmt"
	"strings"
)

//  The base Component interface defines operations that can be altered by
//  decorators.
type BaseComponent interface {
	ReadData() []byte
	WriteData(data []byte)
}

//  Concrete Components provide default implementations of the operations. There
//  might be several variations of these classes.
type Buffer struct {
	buffer []byte
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (m *Buffer) ReadData() []byte {
	fmt.Println("Reading data from memory")
	return m.buffer
}

func (m *Buffer) WriteData(data []byte) {
	fmt.Println("Writing data to memory")
	m.buffer = append(m.buffer, data...)
}

type File struct {
	filename string
}

func NewFile() *File {
	return &File{}
}

func (f *File) ReadData() []byte {
	//ioutil.ReadFile(f.filename)
	fmt.Println("Reading data from file")
	return []byte("file content")
}

func (f *File) WriteData(data []byte) {
	fmt.Println("Writing data to file")
	return
}

type EncryptionDecorator struct {
	component BaseComponent
}

func NewEncryptionDecorator(component BaseComponent) *EncryptionDecorator {
	return &EncryptionDecorator{component: component}
}

func (d *EncryptionDecorator) ReadData() []byte {
	data := d.component.ReadData()
	fmt.Println("Decrypt data after reading data")
	return data
}

func (d *EncryptionDecorator) WriteData(data []byte) {
	fmt.Println("Encrypt data before writing data")
	data = []byte(strings.ToLower(string(data)))
	d.component.WriteData(data)
}

type ValidateDecorator struct {
	component BaseComponent
}

func NewValidateDecorator(component BaseComponent) *ValidateDecorator {
	return &ValidateDecorator{component: component}
}

func (v *ValidateDecorator) ReadData() []byte {
	data := v.component.ReadData()
	if len(data) <= 0 {
		fmt.Println("No data found")
		return []byte("")
	}
	fmt.Println("validate data after reading data")
	return data
}

func (v *ValidateDecorator) WriteData(data []byte) {
	fmt.Println("Validate data before writing data")
	if len(data) <= 0 {
		fmt.Println("No data is going to write")
		return
	}
	v.component.WriteData(data)
}
