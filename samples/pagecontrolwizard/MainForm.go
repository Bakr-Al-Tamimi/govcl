// 由GOVCL UI设计器自动生成，不要编辑
package main

import (
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Panel1       *vcl.TPanel
	Panel2       *vcl.TPanel
	PageControl1 *vcl.TPageControl
	TabSheet1    *vcl.TTabSheet
	TabSheet2    *vcl.TTabSheet
	TabSheet3    *vcl.TTabSheet
	TabSheet4    *vcl.TTabSheet
	TabSheet5    *vcl.TTabSheet
	TabSheet6    *vcl.TTabSheet
	TabSheet7    *vcl.TTabSheet
	Button1      *vcl.TButton
	Button2      *vcl.TButton
	Button3      *vcl.TButton
	Button4      *vcl.TButton
	Button5      *vcl.TButton
	Button6      *vcl.TButton
	Button7      *vcl.TButton
	Bevel1       *vcl.TBevel
	Button8      *vcl.TButton
	Button9      *vcl.TButton
	ActionList1  *vcl.TActionList
	ActPagePrev  *vcl.TAction
	ActPageNext  *vcl.TAction

	//::private::
	TMainFormFields
}

var MainForm *TMainForm

// 以字节形式加载
// vcl.Application.CreateForm(&MainForm)

var (
	mainFormBytes = []byte{
		0x47, 0x4F, 0x56, 0x43, 0x4C, 0x46, 0x4F, 0x52, 0x4D, 0x5A, 0x01, 0x00,
		0x93, 0x48, 0x06, 0xDD, 0x28, 0x61, 0xEB, 0x1B, 0x9D, 0x7F, 0x00, 0xA4,
		0x52, 0x36, 0xDC, 0xA6, 0xFB, 0x64, 0xA3, 0x68, 0x26, 0x3E, 0xB8, 0x4B,
		0x2B, 0x51, 0xA0, 0x02, 0x34, 0x38, 0xDC, 0xA7, 0xFE, 0x45, 0x35, 0xC9,
		0x8F, 0x84, 0x60, 0x9B, 0xF8, 0xE6, 0xBC, 0x91, 0xFF, 0x4C, 0xC9, 0xE3,
		0x39, 0xC2, 0x85, 0x3E, 0x07, 0x26, 0x6A, 0x5C, 0x4D, 0x78, 0xEB, 0x02,
		0xBC, 0x88, 0x07, 0x87, 0x14, 0xD6, 0x81, 0x53, 0xC5, 0xE7, 0x27, 0x9B,
		0x96, 0xB8, 0x48, 0x4D, 0xEC, 0xD7, 0x9B, 0x9E, 0x62, 0xD3, 0x68, 0x86,
		0x60, 0xE0, 0x85, 0x94, 0x7E, 0x25, 0xDA, 0x67, 0xE0, 0xB9, 0x73, 0xF2,
		0xF0, 0x81, 0x79, 0xBC, 0xA5, 0xE5, 0x1E, 0x7D, 0x57, 0xD8, 0x89, 0xD4,
		0x6A, 0x5D, 0xEF, 0x0E, 0xAA, 0x2F, 0x81, 0x9A, 0xBD, 0xDA, 0x51, 0xD8,
		0xFA, 0x77, 0x8D, 0x16, 0xE4, 0xE3, 0x80, 0x60, 0x37, 0xA1, 0x8A, 0xBC,
		0x41, 0x8F, 0x34, 0x29, 0xA4, 0x04, 0xA1, 0xDB, 0xDF, 0xF0, 0x3A, 0xAC,
		0xAD, 0xE1, 0x4A, 0xE6, 0x9A, 0x64, 0x9E, 0xBE, 0xBF, 0x89, 0xC6, 0x22,
		0xA0, 0x7F, 0x5B, 0xA5, 0x4D, 0x30, 0x41, 0x5D, 0x13, 0x5E, 0xFB, 0xB7,
		0x09, 0xF6, 0x26, 0xA1, 0x99, 0x6C, 0x11, 0x4B, 0x15, 0x2C, 0x58, 0x99,
		0x5E, 0xCF, 0xA3, 0xB8, 0x4A, 0xCF, 0x90, 0x5C, 0xC7, 0x5E, 0x93, 0xA3,
		0x0D, 0xA1, 0x88, 0x4B, 0xF6, 0x65, 0xB8, 0x60, 0xD4, 0x9E, 0x30, 0x76,
		0x05, 0x20, 0x72, 0x78, 0x86, 0x02, 0x06, 0x37, 0xDD, 0x67, 0xA5, 0x81,
		0x5D, 0xE9, 0x9C, 0x44, 0x6C, 0x70, 0x09, 0xD6, 0xB1, 0x84, 0xCB, 0xE9,
		0x14, 0x98, 0x35, 0xCE, 0x48, 0x6E, 0x39, 0xEA, 0xE1, 0xDF, 0x4B, 0x17,
		0xD7, 0x27, 0xBE, 0xF5, 0xFC, 0x34, 0x16, 0x6B, 0x7C, 0x64, 0xE5, 0xB8,
		0x84, 0xED, 0x05, 0xDA, 0xD7, 0x07, 0x44, 0x82, 0x98, 0x95, 0x93, 0x1D,
		0xC7, 0x14, 0xD5, 0xEB, 0xC3, 0x87, 0x10, 0x35, 0xA2, 0x4F, 0xFD, 0x53,
		0x27, 0xCD, 0x0C, 0x24, 0xC8, 0xEC, 0xBA, 0xAC, 0x56, 0xEC, 0xF4, 0x61,
		0x07, 0x55, 0x31, 0x50, 0xF1, 0x99, 0xCE, 0x17, 0x38, 0x69, 0x78, 0x77,
		0x90, 0xD3, 0x2B, 0x61, 0x3D, 0xC9, 0xA3, 0x51, 0xFB, 0x76, 0x62, 0x68,
		0x10, 0xBE, 0x5A, 0x63, 0x16, 0x5E, 0x80, 0x48, 0x82, 0x07, 0xE1, 0x48,
		0x1F, 0xC0, 0xA1, 0x9A, 0x10, 0xD5, 0xCF, 0xCF, 0x81, 0xCC, 0xAB, 0x53,
		0xF1, 0x65, 0x29, 0x9A, 0x64, 0x92, 0x6F, 0x71, 0xA1, 0xCA, 0x31, 0xFB,
		0x77, 0xD2, 0x89, 0xDC, 0xC0, 0x2A, 0xEE, 0x66, 0x45, 0xE5, 0x4C, 0xF3,
		0xDE, 0xB1, 0x35, 0x23, 0x13, 0x84, 0x83, 0x67, 0x53, 0x97, 0x47, 0xC4,
		0xFE, 0x85, 0xB6, 0xBF, 0xA3, 0x88, 0x2B, 0xA9, 0x52, 0xD1, 0x66, 0x14,
		0xA1, 0xBF, 0xCB, 0xEF, 0x93, 0x3C, 0xBE, 0x82, 0x7B, 0x31, 0x0B, 0x46,
		0x6E, 0xCC, 0x44, 0x67, 0x33, 0x16, 0x5F, 0x83, 0xBD, 0x71, 0x7B, 0x01,
		0xB1, 0x4F, 0xFD, 0xFD, 0xFE, 0x97, 0x9E, 0xED, 0x0D, 0xFF, 0xB9, 0xCC,
		0xB8, 0x0C, 0x80, 0xB0, 0xA1, 0xD7, 0xA3, 0xCE, 0x62, 0x6F, 0xC9, 0x80,
		0xFD, 0xE0, 0xC4, 0xFE, 0x8F, 0x10, 0xBA, 0xEC, 0xCC, 0x75, 0x5D, 0x70,
		0x57, 0x06, 0x2D, 0x29, 0x09, 0x2C, 0xB0, 0xE4, 0x8D, 0x11, 0xA6, 0x2D,
		0xB9, 0x7B, 0xD5, 0xC7, 0xFB, 0x09, 0x40, 0xF9, 0x63, 0x2B, 0x57, 0x48,
		0x72, 0xE3, 0x2C, 0x91, 0x93, 0xD9, 0xDD, 0x01, 0x1B, 0x62, 0x4C, 0xCB,
		0x4D, 0x48, 0x2E, 0xC2, 0x78, 0x72, 0xA2, 0x19, 0x23, 0xE9, 0x36, 0x22,
		0x23, 0x55, 0xB8, 0xB3, 0x4A, 0x9C, 0x2D, 0xD2, 0x65, 0x2A, 0xF1, 0x20,
		0x08, 0x9E, 0xC6, 0xAC, 0x11, 0x58, 0x17, 0xA0, 0x8D, 0x69, 0xB3, 0xBB,
		0x38, 0xF9, 0xDE, 0x70, 0xDD, 0xA7, 0x53, 0xCA, 0xF0, 0xF6, 0x7A, 0x83,
		0xBB, 0x5E, 0xDA, 0xDC, 0x75, 0x76, 0xAC, 0xF9, 0xC9, 0x01, 0x0E, 0x82,
		0x0F, 0x9F, 0xE8, 0xC5, 0xD0, 0xDD, 0x59, 0x34, 0x85, 0x6B, 0x46, 0x67,
		0x43, 0x44, 0xBB, 0xFA, 0xE8, 0x58, 0xFD, 0xAA, 0x79, 0xFC, 0x98, 0xD9,
		0x14, 0x98, 0xFF, 0xD8, 0xA4, 0xF4, 0x26, 0x24, 0xD2, 0xE0, 0xA9, 0xDC,
		0x60, 0x9D, 0x23, 0x8D, 0x66, 0x26, 0xB3, 0x08, 0x2B, 0x54, 0x60, 0x00,
		0x92, 0xB4, 0x89, 0x76, 0x39, 0xAC, 0x07, 0xED, 0xBE, 0x5F, 0x42, 0x54,
		0xAF, 0xF8, 0x06, 0xF7, 0x60, 0x5B, 0x9C, 0x53, 0x13, 0xFA, 0x56, 0x91,
		0x09, 0x2A, 0xE0, 0x11, 0xDA, 0x0A, 0xE6, 0x5B, 0x07, 0xB0, 0x05, 0xE5,
		0xFD, 0x93, 0xAF, 0x80, 0x24, 0xC8, 0xA0, 0x52, 0xC0, 0xCF, 0xE0, 0x0E,
		0x77, 0x89, 0x64, 0xC3, 0x26, 0xE6, 0x86, 0x79, 0xDC, 0x27, 0xFD, 0x30,
		0x70, 0x59, 0xA5, 0xAD, 0x3E, 0xED, 0x83, 0x51, 0x10, 0x09, 0xDF, 0x72,
		0xC3, 0x13, 0xAE, 0x09, 0x5A, 0x75, 0x81, 0x0B, 0xD6, 0x5F, 0x7A, 0xCE,
		0x92, 0x94, 0x9E, 0x5B, 0xD7, 0x7F, 0xC1, 0x2C, 0x0D, 0x44, 0x44, 0x40,
		0xAE, 0xF3, 0xCF, 0xFE, 0x56, 0x36, 0x8A, 0x6E, 0x18, 0xD8, 0xC0, 0x88,
		0xFA, 0x75, 0xA9, 0xD1, 0xF2, 0x99, 0x51, 0xE4, 0x55, 0xB1, 0x02, 0xA7,
		0x62, 0xAA, 0xC1, 0xB1, 0x75, 0x1F, 0x1C, 0x70, 0xA7, 0x36, 0x79, 0xB0,
		0xFA, 0x86, 0x50, 0x66, 0x44, 0xC0, 0xBE, 0xDA, 0x98, 0xC6, 0xDD, 0x1D,
		0x78, 0x0B, 0x5F, 0x06, 0x89, 0xB9, 0x87, 0x64, 0xE5, 0x5A, 0x81, 0xE1,
		0x6E, 0x3D, 0x07, 0xC3, 0xD9, 0x44, 0x21, 0xB8, 0x81, 0x95, 0xB3, 0x0A,
		0xE8, 0x0F, 0xF5, 0x98, 0x11, 0x00}
)

var _ = vcl.RegisterFormResource(MainForm, &mainFormBytes)
