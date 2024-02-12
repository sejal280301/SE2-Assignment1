package decoder

import (
	"bytes"
	"os"
	"testing"

	"github.com/cloudflare/ebpf_exporter/v2/config"
	"github.com/cloudflare/ebpf_exporter/v2/kallsyms"
)

func TestKStackDecoder(t *testing.T) {
	cases := []struct {
		in  []byte
		out []byte
	}{
		{
			in: []byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			out: []byte(""),
		},
		{
			in: []byte{
				0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			out: []byte("one\none\ntwo"),
		},
		{
			in: []byte{
				0xab, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			out: []byte("three\ntwo\ntwo"),
		},
		{
			in: []byte{
				0xab, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			out: []byte("three\nzero"),
		},
		{
			in: []byte{
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0xac, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			out: []byte("??\nthree\n??"),
		},
	}

	fd, err := os.CreateTemp("", "kallsyms")
	if err != nil {
		t.Fatalf("Error creating temporary file for kallsyms: %v", err)
	}

	defer os.Remove(fd.Name())

	_, err = fd.WriteString("0000000000000004 T one\n0000000000000006 T two\n00000000000000aa T three\n")
	if err != nil {
		t.Fatalf("Error writing fake kallsyms data to %q: %v", fd.Name(), err)
	}

	decoder, err := kallsyms.NewDecoder(fd.Name())
	if err != nil {
		t.Fatalf("Error creating ksym decoder for %q: %v", fd.Name(), err)
	}

	_, err = fd.WriteString("0000000000000002 T zero\n")
	if err != nil {
		t.Fatalf("Error writing additional fake kallsyms data to %q: %v", fd.Name(), err)
	}

	d := KStack{decoder}

	for _, c := range cases {
		out, err := d.Decode(c.in, config.Decoder{})
		if err != nil {
			t.Errorf("Error decoding %#v: %s", c.in, err)
		}

		if !bytes.Equal(out, c.out) {
			t.Errorf("Expected %q, got %q", c.out, out)
		}
	}
}
