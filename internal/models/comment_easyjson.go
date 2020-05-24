// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE9abebc9Decode20201ColorNoiseInternalModels(in *jlexer.Lexer, out *ResponseComment) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint(in.Uint())
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(ResponseUser)
				}
				easyjsonE9abebc9Decode20201ColorNoiseInternalModels1(in, out.User)
			}
		case "pin_id":
			out.PindId = uint(in.Uint())
		case "created_at":
			if in.IsNull() {
				in.Skip()
				out.CreatedAt = nil
			} else {
				if out.CreatedAt == nil {
					out.CreatedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreatedAt).UnmarshalJSON(data))
				}
			}
		case "comment":
			out.Text = string(in.String())
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
func easyjsonE9abebc9Encode20201ColorNoiseInternalModels(out *jwriter.Writer, in ResponseComment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	if in.User != nil {
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonE9abebc9Encode20201ColorNoiseInternalModels1(out, *in.User)
	}
	if in.PindId != 0 {
		const prefix string = ",\"pin_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint(uint(in.PindId))
	}
	if in.CreatedAt != nil {
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.CreatedAt).MarshalJSON())
	}
	if in.Text != "" {
		const prefix string = ",\"comment\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseComment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseComment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseComment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseComment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels(l, v)
}
func easyjsonE9abebc9Decode20201ColorNoiseInternalModels1(in *jlexer.Lexer, out *ResponseUser) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint(in.Uint())
		case "email":
			out.Email = string(in.String())
		case "login":
			out.Login = string(in.String())
		case "about":
			out.About = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "subscriptions":
			out.Subscriptions = int(in.Int())
		case "subscribers":
			out.Subscribers = int(in.Int())
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
func easyjsonE9abebc9Encode20201ColorNoiseInternalModels1(out *jwriter.Writer, in ResponseUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix)
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	if in.Avatar != "" {
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"subscriptions\":"
		out.RawString(prefix)
		out.Int(int(in.Subscriptions))
	}
	{
		const prefix string = ",\"subscribers\":"
		out.RawString(prefix)
		out.Int(int(in.Subscribers))
	}
	out.RawByte('}')
}
func easyjsonE9abebc9Decode20201ColorNoiseInternalModels2(in *jlexer.Lexer, out *InputComment) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "pin_id":
			out.PinId = uint(in.Uint())
		case "comment":
			out.Text = string(in.String())
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
func easyjsonE9abebc9Encode20201ColorNoiseInternalModels2(out *jwriter.Writer, in InputComment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pin_id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.PinId))
	}
	{
		const prefix string = ",\"comment\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InputComment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InputComment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InputComment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InputComment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels2(l, v)
}
func easyjsonE9abebc9Decode20201ColorNoiseInternalModels3(in *jlexer.Lexer, out *DataBaseComment) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = uint(in.Uint())
		case "UserId":
			out.UserId = uint(in.Uint())
		case "PinId":
			out.PinId = uint(in.Uint())
		case "Text":
			out.Text = string(in.String())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
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
func easyjsonE9abebc9Encode20201ColorNoiseInternalModels3(out *jwriter.Writer, in DataBaseComment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix)
		out.Uint(uint(in.UserId))
	}
	{
		const prefix string = ",\"PinId\":"
		out.RawString(prefix)
		out.Uint(uint(in.PinId))
	}
	{
		const prefix string = ",\"Text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DataBaseComment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DataBaseComment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DataBaseComment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DataBaseComment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels3(l, v)
}
func easyjsonE9abebc9Decode20201ColorNoiseInternalModels4(in *jlexer.Lexer, out *Comment) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = uint(in.Uint())
		case "User":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(User)
				}
				easyjsonE9abebc9Decode20201ColorNoiseInternalModels5(in, out.User)
			}
		case "PinId":
			out.PinId = uint(in.Uint())
		case "Text":
			out.Text = string(in.String())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
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
func easyjsonE9abebc9Encode20201ColorNoiseInternalModels4(out *jwriter.Writer, in Comment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"User\":"
		out.RawString(prefix)
		if in.User == nil {
			out.RawString("null")
		} else {
			easyjsonE9abebc9Encode20201ColorNoiseInternalModels5(out, *in.User)
		}
	}
	{
		const prefix string = ",\"PinId\":"
		out.RawString(prefix)
		out.Uint(uint(in.PinId))
	}
	{
		const prefix string = ",\"Text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9abebc9Encode20201ColorNoiseInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9abebc9Decode20201ColorNoiseInternalModels4(l, v)
}
func easyjsonE9abebc9Decode20201ColorNoiseInternalModels5(in *jlexer.Lexer, out *User) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Id":
			out.Id = uint(in.Uint())
		case "Email":
			out.Email = string(in.String())
		case "Login":
			out.Login = string(in.String())
		case "EncryptedPassword":
			out.EncryptedPassword = string(in.String())
		case "About":
			out.About = string(in.String())
		case "Avatar":
			out.Avatar = string(in.String())
		case "Subscriptions":
			out.Subscriptions = int(in.Int())
		case "Subscribers":
			out.Subscribers = int(in.Int())
		case "Preferences":
			if in.IsNull() {
				in.Skip()
				out.Preferences = nil
			} else {
				in.Delim('[')
				if out.Preferences == nil {
					if !in.IsDelim(']') {
						out.Preferences = make([]string, 0, 4)
					} else {
						out.Preferences = []string{}
					}
				} else {
					out.Preferences = (out.Preferences)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Preferences = append(out.Preferences, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "Tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Tags = append(out.Tags, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonE9abebc9Encode20201ColorNoiseInternalModels5(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"Login\":"
		out.RawString(prefix)
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"EncryptedPassword\":"
		out.RawString(prefix)
		out.String(string(in.EncryptedPassword))
	}
	{
		const prefix string = ",\"About\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"Avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"Subscriptions\":"
		out.RawString(prefix)
		out.Int(int(in.Subscriptions))
	}
	{
		const prefix string = ",\"Subscribers\":"
		out.RawString(prefix)
		out.Int(int(in.Subscribers))
	}
	{
		const prefix string = ",\"Preferences\":"
		out.RawString(prefix)
		if in.Preferences == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Preferences {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"Tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tags {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
