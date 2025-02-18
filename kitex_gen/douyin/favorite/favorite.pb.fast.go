// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package favorite

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	feed "simpledouyin/kitex_gen/douyin/feed"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *FavoriteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteRequest[number], err)
}

func (x *FavoriteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ActorId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteResponse[number], err)
}

func (x *FavoriteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *FavoriteListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteListRequest[number], err)
}

func (x *FavoriteListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ActorId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FavoriteListResponse[number], err)
}

func (x *FavoriteListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *FavoriteListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v feed.Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *IsFavoriteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_IsFavoriteRequest[number], err)
}

func (x *IsFavoriteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *IsFavoriteRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *IsFavoriteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_IsFavoriteResponse[number], err)
}

func (x *IsFavoriteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Result, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *CountFavoriteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountFavoriteRequest[number], err)
}

func (x *CountFavoriteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountFavoriteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountFavoriteResponse[number], err)
}

func (x *CountFavoriteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountFavoriteResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *CountFavoriteResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Count, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountUserFavoriteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountUserFavoriteRequest[number], err)
}

func (x *CountUserFavoriteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountUserFavoriteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountUserFavoriteResponse[number], err)
}

func (x *CountUserFavoriteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountUserFavoriteResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *CountUserFavoriteResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Count, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountUserTotalFavoritedRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountUserTotalFavoritedRequest[number], err)
}

func (x *CountUserTotalFavoritedRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ActorId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountUserTotalFavoritedRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountUserTotalFavoritedResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountUserTotalFavoritedResponse[number], err)
}

func (x *CountUserTotalFavoritedResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *CountUserTotalFavoritedResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *CountUserTotalFavoritedResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Count, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *FavoriteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FavoriteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ActorId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetActorId())
	return offset
}

func (x *FavoriteRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetVideoId())
	return offset
}

func (x *FavoriteRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetActionType())
	return offset
}

func (x *FavoriteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *FavoriteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *FavoriteResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *FavoriteListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *FavoriteListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ActorId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetActorId())
	return offset
}

func (x *FavoriteListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *FavoriteListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FavoriteListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *FavoriteListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *FavoriteListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.GetVideoList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetVideoList()[i])
	}
	return offset
}

func (x *IsFavoriteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *IsFavoriteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *IsFavoriteRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetVideoId())
	return offset
}

func (x *IsFavoriteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *IsFavoriteResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Result {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetResult())
	return offset
}

func (x *CountFavoriteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CountFavoriteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetVideoId())
	return offset
}

func (x *CountFavoriteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CountFavoriteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *CountFavoriteResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *CountFavoriteResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Count == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetCount())
	return offset
}

func (x *CountUserFavoriteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CountUserFavoriteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CountUserFavoriteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CountUserFavoriteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *CountUserFavoriteResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *CountUserFavoriteResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Count == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetCount())
	return offset
}

func (x *CountUserTotalFavoritedRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CountUserTotalFavoritedRequest) fastWriteField1(buf []byte) (offset int) {
	if x.ActorId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetActorId())
	return offset
}

func (x *CountUserTotalFavoritedRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *CountUserTotalFavoritedResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CountUserTotalFavoritedResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *CountUserTotalFavoritedResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *CountUserTotalFavoritedResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Count == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetCount())
	return offset
}

func (x *FavoriteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FavoriteRequest) sizeField1() (n int) {
	if x.ActorId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetActorId())
	return n
}

func (x *FavoriteRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetVideoId())
	return n
}

func (x *FavoriteRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetActionType())
	return n
}

func (x *FavoriteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *FavoriteResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *FavoriteResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *FavoriteListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *FavoriteListRequest) sizeField1() (n int) {
	if x.ActorId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetActorId())
	return n
}

func (x *FavoriteListRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetUserId())
	return n
}

func (x *FavoriteListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FavoriteListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *FavoriteListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *FavoriteListResponse) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.GetVideoList() {
		n += fastpb.SizeMessage(3, x.GetVideoList()[i])
	}
	return n
}

func (x *IsFavoriteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *IsFavoriteRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *IsFavoriteRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetVideoId())
	return n
}

func (x *IsFavoriteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *IsFavoriteResponse) sizeField1() (n int) {
	if !x.Result {
		return n
	}
	n += fastpb.SizeBool(1, x.GetResult())
	return n
}

func (x *CountFavoriteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CountFavoriteRequest) sizeField1() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetVideoId())
	return n
}

func (x *CountFavoriteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CountFavoriteResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *CountFavoriteResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *CountFavoriteResponse) sizeField3() (n int) {
	if x.Count == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetCount())
	return n
}

func (x *CountUserFavoriteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CountUserFavoriteRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetUserId())
	return n
}

func (x *CountUserFavoriteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CountUserFavoriteResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *CountUserFavoriteResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *CountUserFavoriteResponse) sizeField3() (n int) {
	if x.Count == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetCount())
	return n
}

func (x *CountUserTotalFavoritedRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CountUserTotalFavoritedRequest) sizeField1() (n int) {
	if x.ActorId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetActorId())
	return n
}

func (x *CountUserTotalFavoritedRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetUserId())
	return n
}

func (x *CountUserTotalFavoritedResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CountUserTotalFavoritedResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *CountUserTotalFavoritedResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *CountUserTotalFavoritedResponse) sizeField3() (n int) {
	if x.Count == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetCount())
	return n
}

var fieldIDToName_FavoriteRequest = map[int32]string{
	1: "ActorId",
	2: "VideoId",
	3: "ActionType",
}

var fieldIDToName_FavoriteResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_FavoriteListRequest = map[int32]string{
	1: "ActorId",
	2: "UserId",
}

var fieldIDToName_FavoriteListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
}

var fieldIDToName_IsFavoriteRequest = map[int32]string{
	1: "UserId",
	2: "VideoId",
}

var fieldIDToName_IsFavoriteResponse = map[int32]string{
	1: "Result",
}

var fieldIDToName_CountFavoriteRequest = map[int32]string{
	1: "VideoId",
}

var fieldIDToName_CountFavoriteResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Count",
}

var fieldIDToName_CountUserFavoriteRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_CountUserFavoriteResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Count",
}

var fieldIDToName_CountUserTotalFavoritedRequest = map[int32]string{
	1: "ActorId",
	2: "UserId",
}

var fieldIDToName_CountUserTotalFavoritedResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "Count",
}

var _ = feed.File_feed_proto
