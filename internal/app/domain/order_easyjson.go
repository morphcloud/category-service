// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain(in *jlexer.Lexer, out *OrderList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(OrderList, 0, 0)
			} else {
				*out = OrderList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Order
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain(out *jwriter.Writer, in OrderList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v OrderList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain(l, v)
}
func easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain1(in *jlexer.Lexer, out *Order) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "restaurantId":
			out.RestaurantID = int64(in.Int64())
		case "consumerId":
			out.ConsumerID = int64(in.Int64())
		case "deliveryId":
			out.DeliveryID = int64(in.Int64())
		case "paymentId":
			out.PaymentID = int64(in.Int64())
		case "statusId":
			out.StatusID = int32(in.Int32())
		case "createdAt":
			easyjson120d1ca2DecodeGoMongodbOrgMongoDriverBsonPrimitive(in, &out.CreatedAt)
		case "updatedAt":
			easyjson120d1ca2DecodeGoMongodbOrgMongoDriverBsonPrimitive(in, &out.UpdatedAt)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain1(out *jwriter.Writer, in Order) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"restaurantId\":"
		out.RawString(prefix)
		out.Int64(int64(in.RestaurantID))
	}
	{
		const prefix string = ",\"consumerId\":"
		out.RawString(prefix)
		out.Int64(int64(in.ConsumerID))
	}
	{
		const prefix string = ",\"deliveryId\":"
		out.RawString(prefix)
		out.Int64(int64(in.DeliveryID))
	}
	{
		const prefix string = ",\"paymentId\":"
		out.RawString(prefix)
		out.Int64(int64(in.PaymentID))
	}
	{
		const prefix string = ",\"statusId\":"
		out.RawString(prefix)
		out.Int32(int32(in.StatusID))
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		easyjson120d1ca2EncodeGoMongodbOrgMongoDriverBsonPrimitive(out, in.CreatedAt)
	}
	{
		const prefix string = ",\"updatedAt\":"
		out.RawString(prefix)
		easyjson120d1ca2EncodeGoMongodbOrgMongoDriverBsonPrimitive(out, in.UpdatedAt)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Order) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Order) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Order) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Order) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain1(l, v)
}
func easyjson120d1ca2DecodeGoMongodbOrgMongoDriverBsonPrimitive(in *jlexer.Lexer, out *primitive.Timestamp) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "T":
			out.T = uint32(in.Uint32())
		case "I":
			out.I = uint32(in.Uint32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson120d1ca2EncodeGoMongodbOrgMongoDriverBsonPrimitive(out *jwriter.Writer, in primitive.Timestamp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.T))
	}
	{
		const prefix string = ",\"I\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.I))
	}
	out.RawByte('}')
}
func easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain2(in *jlexer.Lexer, out *LineItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "menuItemId":
			out.MenuItemID = int64(in.Int64())
		case "orderId":
			out.OrderID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "price":
			out.Price = float32(in.Float32())
		case "quantity":
			out.Quantity = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain2(out *jwriter.Writer, in LineItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"menuItemId\":"
		out.RawString(prefix)
		out.Int64(int64(in.MenuItemID))
	}
	{
		const prefix string = ",\"orderId\":"
		out.RawString(prefix)
		out.Int64(int64(in.OrderID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Float32(float32(in.Price))
	}
	{
		const prefix string = ",\"quantity\":"
		out.RawString(prefix)
		out.Int32(int32(in.Quantity))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LineItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LineItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeGithubComMorphcloudOrderServiceInternalAppDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LineItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LineItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeGithubComMorphcloudOrderServiceInternalAppDomain2(l, v)
}