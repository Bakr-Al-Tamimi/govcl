// 由res2go自动生成，不要编辑。
package main

import (
	"github.com/ying32/govcl/vcl"
)

type TForm11 struct {
	*vcl.TForm
	ListView1    *vcl.TListView
	Button1      *vcl.TButton
	ImageList1   *vcl.TImageList
	ColorDialog1 *vcl.TColorDialog

	//::private::
	TForm11Fields
}

var Form11 *TForm11

// 以字节形式加载
// vcl.Application.CreateForm(&Form11)

var form11Bytes = []byte("\x54\x50\x46\x30\x07\x54\x46\x6F\x72\x6D\x31\x31\x06\x46\x6F\x72\x6D\x31\x31\x04\x4C\x65\x66\x74\x03\x6B\x02\x03\x54\x6F\x70\x03\x8F\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x46\x6F\x72\x6D\x31\x31\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x41\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xA6\x02\x05\x43\x6F\x6C\x6F\x72\x07\x09\x63\x6C\x42\x74\x6E\x46\x61\x63\x65\x0C\x46\x6F\x6E\x74\x2E\x43\x68\x61\x72\x73\x65\x74\x07\x0F\x44\x45\x46\x41\x55\x4C\x54\x5F\x43\x48\x41\x52\x53\x45\x54\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF5\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x06\x54\x61\x68\x6F\x6D\x61\x0A\x46\x6F\x6E\x74\x2E\x53\x74\x79\x6C\x65\x0B\x00\x0E\x4F\x6C\x64\x43\x72\x65\x61\x74\x65\x4F\x72\x64\x65\x72\x08\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0A\x70\x6F\x44\x65\x73\x69\x67\x6E\x65\x64\x08\x4F\x6E\x43\x72\x65\x61\x74\x65\x07\x0A\x46\x6F\x72\x6D\x43\x72\x65\x61\x74\x65\x0D\x50\x69\x78\x65\x6C\x73\x50\x65\x72\x49\x6E\x63\x68\x02\x60\x0A\x54\x65\x78\x74\x48\x65\x69\x67\x68\x74\x02\x0D\x00\x09\x54\x4C\x69\x73\x74\x56\x69\x65\x77\x09\x4C\x69\x73\x74\x56\x69\x65\x77\x31\x04\x4C\x65\x66\x74\x02\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xA9\x02\x06\x48\x65\x69\x67\x68\x74\x03\x21\x01\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x02\x00\x00\x00\x0D\x54\xF0\x79\x05\x57\x69\x64\x74\x68\x02\x3C\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x02\x00\x00\x00\xEF\x53\xC1\x89\x05\x57\x69\x64\x74\x68\x02\x50\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x02\x00\x00\x00\x9C\x98\x72\x82\x05\x57\x69\x64\x74\x68\x02\x64\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x12\x01\x00\x00\x00\x3C\x50\x05\x57\x69\x64\x74\x68\x02\x78\x00\x00\x09\x47\x72\x69\x64\x4C\x69\x6E\x65\x73\x09\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x09\x52\x6F\x77\x53\x65\x6C\x65\x63\x74\x09\x0B\x53\x6D\x61\x6C\x6C\x49\x6D\x61\x67\x65\x73\x07\x0A\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x31\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x09\x56\x69\x65\x77\x53\x74\x79\x6C\x65\x07\x08\x76\x73\x52\x65\x70\x6F\x72\x74\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x00\x03\x54\x6F\x70\x03\x27\x01\x05\x57\x69\x64\x74\x68\x02\x4B\x06\x48\x65\x69\x67\x68\x74\x02\x19\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x42\x75\x74\x74\x6F\x6E\x31\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x0A\x54\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x0A\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x31\x06\x48\x65\x69\x67\x68\x74\x02\x18\x05\x57\x69\x64\x74\x68\x02\x18\x04\x4C\x65\x66\x74\x03\x50\x01\x03\x54\x6F\x70\x03\xA8\x00\x00\x00\x0C\x54\x43\x6F\x6C\x6F\x72\x44\x69\x61\x6C\x6F\x67\x0C\x43\x6F\x6C\x6F\x72\x44\x69\x61\x6C\x6F\x67\x31\x04\x4C\x65\x66\x74\x03\x58\x01\x03\x54\x6F\x70\x03\xB0\x00\x00\x00\x00")

var _ = vcl.RegisterFormResource(Form11, &form11Bytes)
