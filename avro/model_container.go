// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     model.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/container"
	"github.com/actgardner/gogen-avro/v9/vm"
)

func NewModelWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewModel()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ModelReader struct {
	r io.Reader
	p *vm.Program
}

func NewModelReader(r io.Reader) (*ModelReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewModel()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ModelReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ModelReader) Read() (Model, error) {
	t := NewModel()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
