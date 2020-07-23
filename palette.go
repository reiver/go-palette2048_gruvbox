package palette2048_gruvbox

import (
	"github.com/reiver/go-palette2048"
)

// Palette provides the gruvbox palette.
var Palette [palette2048.ByteSize]uint8 = [palette2048.ByteSize]uint8{
	0x1d,0x20,0x21, 0xff, // D0_h
	0x28,0x28,0x28, 0xff, // D0
	0x32,0x30,0x2f, 0xff, // D0_s
	0x3c,0x38,0x36, 0xff, // D1
	0x50,0x49,0x45, 0xff, // D2
	0x66,0x5c,0x54, 0xff, // D3
	0x7c,0x6f,0x64, 0xff, // D4
	0x92,0x83,0x74, 0xff, // gray
	0xa8,0x99,0x84, 0xff, // L4
	0xbd,0xae,0x93, 0xff, // L3
	0xd5,0xc4,0xa1, 0xff, // L2
	0xeb,0xdb,0xb2, 0xff, // L1
	0xf2,0xe5,0xbc, 0xff, // L0_s
	0xf9,0xf5,0xd7, 0xff, // L0_h
	0xfb,0xf1,0xc7, 0xff, // L0

	0x9d,0x00,0x06, 0xff, // faded red
	0xcc,0x24,0x1d, 0xff, // neutral red
	0xfb,0x49,0x34, 0xff, // bright red

	0x79,0x74,0x0e, 0xff, // faded green
	0x98,0x97,0x1a, 0xff, // neutral green
	0xb8,0xbb,0x26, 0xff, // bright green

	0xb5,0x76,0x14, 0xff, // faded yellow
	0xd7,0x99,0x21, 0xff, // neutral yellow
	0xfa,0xbd,0x2f, 0xff, // bright yellow

	0x07,0x66,0x78, 0xff, // faded blue
	0x45,0x85,0x88, 0xff, // neutral blue
	0x83,0xa5,0x98, 0xff, // bright blue

	0x8f,0x3f,0x71, 0xff, // faded purple
	0xb1,0x62,0x86, 0xff, // neutral purple
	0xd3,0x86,0x9b, 0xff, // bright purple

	0x42,0x7b,0x58, 0xff, // faded aqua
	0x68,0x9d,0x6a, 0xff, // neutral aqua
	0x8e,0xc0,0x7c, 0xff, // bright aqua

	0xaf,0x3a,0x03, 0xff, // faded orange
	0xd6,0x5d,0x0e, 0xff, // neutral orange
	0xfe,0x80,0x19, 0xff, // bright orange
}