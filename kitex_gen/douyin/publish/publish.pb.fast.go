// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package publish

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	feed "simpledouyin/kitex_gen/douyin/feed"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CreateVideoRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateVideoRequest[number], err)
}

func (x *CreateVideoRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ActorId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CreateVideoRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *CreateVideoRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateVideoResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateVideoResponse[number], err)
}

func (x *CreateVideoResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CreateVideoResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ListVideoRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListVideoRequest[number], err)
}

func (x *ListVideoRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *ListVideoRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ActorId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *ListVideoResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListVideoResponse[number], err)
}

func (x *ListVideoResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *ListVideoResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *ListVideoResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v feed.Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *CountVideoRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountVideoRequest[number], err)
}

func (x *CountVideoRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountVideoResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountVideoResponse[number], err)
}

func (x *CountVideoResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountVideoResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *CountVideoResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Count, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CreateVideoRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CreateVideoRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ActorId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetActorId())
	return offset
}

func (x *CreateVideoRequest) fastWriteField2(buf []byte) (offset int) {
	if len(x.Data) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 2, x.GetData())
	return offset
}

func (x *CreateVideoRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetTitle())
	return offset
}

func (x *CreateVideoResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateVideoResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *CreateVideoResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *ListVideoRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListVideoRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *ListVideoRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ActorId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetActorId())
	return offset
}

func (x *ListVideoResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ListVideoResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *ListVideoResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *ListVideoResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.GetVideoList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetVideoList()[i])
	}
	return offset
}

func (x *CountVideoRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CountVideoRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CountVideoResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CountVideoResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *CountVideoResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *CountVideoResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Count == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetCount())
	return offset
}

func (x *CreateVideoRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CreateVideoRequest) sizeField1() (n int) {
	if x.ActorId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetActorId())
	return n
}

func (x *CreateVideoRequest) sizeField2() (n int) {
	if len(x.Data) == 0 {
		return n
	}
	n += fastpb.SizeBytes(2, x.GetData())
	return n
}

func (x *CreateVideoRequest) sizeField3() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetTitle())
	return n
}

func (x *CreateVideoResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateVideoResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *CreateVideoResponse) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *ListVideoRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListVideoRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *ListVideoRequest) sizeField2() (n int) {
	if x.ActorId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetActorId())
	return n
}

func (x *ListVideoResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ListVideoResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *ListVideoResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *ListVideoResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.GetVideoList() {
		n += fastpb.SizeMessage(3, x.GetVideoList()[i])
	}
	return n
}

func (x *CountVideoRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CountVideoRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *CountVideoResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CountVideoResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *CountVideoResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *CountVideoResponse) sizeField3() (n int) {
	if x.Count == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetCount())
	return n
}

var fieldIDToName_CreateVideoRequest = map[int32]string{
	1: "ActorId",
	2: "Data",
	3: "Title",
}

var fieldIDToName_CreateVideoResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_ListVideoRequest = map[int32]string{
	1: "UserId",
	2: "ActorId",
}

var fieldIDToName_ListVideoResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
}

var fieldIDToName_CountVideoRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_CountVideoResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Count",
}

var _ = feed.File_feed_proto
