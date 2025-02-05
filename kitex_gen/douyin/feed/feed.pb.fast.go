// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package feed

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	user "simpledouyin/kitex_gen/douyin/user"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Video) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Video[number], err)
}

func (x *Video) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Video) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v user.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Author = &v
	return offset, nil
}

func (x *Video) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.FavoriteCount, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Video) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CommentCount, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Video) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.IsFavorite, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Video) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ListFeedRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListFeedRequest[number], err)
}

func (x *ListFeedRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.LatestTime = &tmp
	return offset, err
}

func (x *ListFeedRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadUint32(buf, _type)
	x.ActorId = &tmp
	return offset, err
}

func (x *ListFeedResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListFeedResponse[number], err)
}

func (x *ListFeedResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *ListFeedResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *ListFeedResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.NextTime = &tmp
	return offset, err
}

func (x *ListFeedResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *QueryVideosRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryVideosRequest[number], err)
}

func (x *QueryVideosRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ActorId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *QueryVideosRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint32
			v, offset, err = fastpb.ReadUint32(buf, _type)
			if err != nil {
				return offset, err
			}
			x.VideoIds = append(x.VideoIds, v)
			return offset, err
		})
	return offset, err
}

func (x *QueryVideosResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryVideosResponse[number], err)
}

func (x *QueryVideosResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *QueryVideosResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *QueryVideosResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *Video) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Video) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Video) fastWriteField2(buf []byte) (offset int) {
	if x.Author == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetAuthor())
	return offset
}

func (x *Video) fastWriteField3(buf []byte) (offset int) {
	if x.PlayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPlayUrl())
	return offset
}

func (x *Video) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCoverUrl())
	return offset
}

func (x *Video) fastWriteField5(buf []byte) (offset int) {
	if x.FavoriteCount == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 5, x.GetFavoriteCount())
	return offset
}

func (x *Video) fastWriteField6(buf []byte) (offset int) {
	if x.CommentCount == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 6, x.GetCommentCount())
	return offset
}

func (x *Video) fastWriteField7(buf []byte) (offset int) {
	if !x.IsFavorite {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 7, x.GetIsFavorite())
	return offset
}

func (x *Video) fastWriteField8(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetTitle())
	return offset
}

func (x *ListFeedRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListFeedRequest) fastWriteField1(buf []byte) (offset int) {
	if x.LatestTime == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetLatestTime())
	return offset
}

func (x *ListFeedRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ActorId == nil {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetActorId())
	return offset
}

func (x *ListFeedResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ListFeedResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *ListFeedResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *ListFeedResponse) fastWriteField3(buf []byte) (offset int) {
	if x.NextTime == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetNextTime())
	return offset
}

func (x *ListFeedResponse) fastWriteField4(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.GetVideoList() {
		offset += fastpb.WriteMessage(buf[offset:], 4, x.GetVideoList()[i])
	}
	return offset
}

func (x *QueryVideosRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *QueryVideosRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ActorId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetActorId())
	return offset
}

func (x *QueryVideosRequest) fastWriteField2(buf []byte) (offset int) {
	if len(x.VideoIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 2, len(x.GetVideoIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint32(buf[offset:], numTagOrKey, x.GetVideoIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *QueryVideosResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *QueryVideosResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *QueryVideosResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *QueryVideosResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.GetVideoList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetVideoList()[i])
	}
	return offset
}

func (x *Video) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *Video) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetId())
	return n
}

func (x *Video) sizeField2() (n int) {
	if x.Author == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetAuthor())
	return n
}

func (x *Video) sizeField3() (n int) {
	if x.PlayUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPlayUrl())
	return n
}

func (x *Video) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCoverUrl())
	return n
}

func (x *Video) sizeField5() (n int) {
	if x.FavoriteCount == 0 {
		return n
	}
	n += fastpb.SizeUint32(5, x.GetFavoriteCount())
	return n
}

func (x *Video) sizeField6() (n int) {
	if x.CommentCount == 0 {
		return n
	}
	n += fastpb.SizeUint32(6, x.GetCommentCount())
	return n
}

func (x *Video) sizeField7() (n int) {
	if !x.IsFavorite {
		return n
	}
	n += fastpb.SizeBool(7, x.GetIsFavorite())
	return n
}

func (x *Video) sizeField8() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetTitle())
	return n
}

func (x *ListFeedRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListFeedRequest) sizeField1() (n int) {
	if x.LatestTime == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetLatestTime())
	return n
}

func (x *ListFeedRequest) sizeField2() (n int) {
	if x.ActorId == nil {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetActorId())
	return n
}

func (x *ListFeedResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ListFeedResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *ListFeedResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *ListFeedResponse) sizeField3() (n int) {
	if x.NextTime == nil {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetNextTime())
	return n
}

func (x *ListFeedResponse) sizeField4() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.GetVideoList() {
		n += fastpb.SizeMessage(4, x.GetVideoList()[i])
	}
	return n
}

func (x *QueryVideosRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *QueryVideosRequest) sizeField1() (n int) {
	if x.ActorId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetActorId())
	return n
}

func (x *QueryVideosRequest) sizeField2() (n int) {
	if len(x.VideoIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(2, len(x.GetVideoIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint32(numTagOrKey, x.GetVideoIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *QueryVideosResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *QueryVideosResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *QueryVideosResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *QueryVideosResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.GetVideoList() {
		n += fastpb.SizeMessage(3, x.GetVideoList()[i])
	}
	return n
}

var fieldIDToName_Video = map[int32]string{
	1: "Id",
	2: "Author",
	3: "PlayUrl",
	4: "CoverUrl",
	5: "FavoriteCount",
	6: "CommentCount",
	7: "IsFavorite",
	8: "Title",
}

var fieldIDToName_ListFeedRequest = map[int32]string{
	1: "LatestTime",
	2: "ActorId",
}

var fieldIDToName_ListFeedResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "NextTime",
	4: "VideoList",
}

var fieldIDToName_QueryVideosRequest = map[int32]string{
	1: "ActorId",
	2: "VideoIds",
}

var fieldIDToName_QueryVideosResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
}

var _ = user.File_user_proto
