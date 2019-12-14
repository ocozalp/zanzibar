// Code generated by thriftrw v1.20.2. DO NOT EDIT.
// @generated

package serverless

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type Request struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// ToWire translates a Request struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Request) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.FirstName != nil {
		w, err = wire.NewValueString(*(v.FirstName)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.LastName != nil {
		w, err = wire.NewValueString(*(v.LastName)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Request struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Request struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Request
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Request) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.FirstName = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.LastName = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Request
// struct.
func (v *Request) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.FirstName != nil {
		fields[i] = fmt.Sprintf("FirstName: %v", *(v.FirstName))
		i++
	}
	if v.LastName != nil {
		fields[i] = fmt.Sprintf("LastName: %v", *(v.LastName))
		i++
	}

	return fmt.Sprintf("Request{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Request match the
// provided Request.
//
// This function performs a deep comparison.
func (v *Request) Equals(rhs *Request) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.FirstName, rhs.FirstName) {
		return false
	}
	if !_String_EqualsPtr(v.LastName, rhs.LastName) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Request.
func (v *Request) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.FirstName != nil {
		enc.AddString("firstName", *v.FirstName)
	}
	if v.LastName != nil {
		enc.AddString("lastName", *v.LastName)
	}
	return err
}

// GetFirstName returns the value of FirstName if it is set or its
// zero value if it is unset.
func (v *Request) GetFirstName() (o string) {
	if v != nil && v.FirstName != nil {
		return *v.FirstName
	}

	return
}

// IsSetFirstName returns true if FirstName is not nil.
func (v *Request) IsSetFirstName() bool {
	return v != nil && v.FirstName != nil
}

// GetLastName returns the value of LastName if it is set or its
// zero value if it is unset.
func (v *Request) GetLastName() (o string) {
	if v != nil && v.LastName != nil {
		return *v.LastName
	}

	return
}

// IsSetLastName returns true if LastName is not nil.
func (v *Request) IsSetLastName() bool {
	return v != nil && v.LastName != nil
}

type Response struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName1 *string `json:"lastName1,omitempty"`
}

// ToWire translates a Response struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Response) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.FirstName != nil {
		w, err = wire.NewValueString(*(v.FirstName)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.LastName1 != nil {
		w, err = wire.NewValueString(*(v.LastName1)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Response struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Response struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Response
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Response) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.FirstName = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.LastName1 = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Response
// struct.
func (v *Response) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.FirstName != nil {
		fields[i] = fmt.Sprintf("FirstName: %v", *(v.FirstName))
		i++
	}
	if v.LastName1 != nil {
		fields[i] = fmt.Sprintf("LastName1: %v", *(v.LastName1))
		i++
	}

	return fmt.Sprintf("Response{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Response match the
// provided Response.
//
// This function performs a deep comparison.
func (v *Response) Equals(rhs *Response) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.FirstName, rhs.FirstName) {
		return false
	}
	if !_String_EqualsPtr(v.LastName1, rhs.LastName1) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Response.
func (v *Response) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.FirstName != nil {
		enc.AddString("firstName", *v.FirstName)
	}
	if v.LastName1 != nil {
		enc.AddString("lastName1", *v.LastName1)
	}
	return err
}

// GetFirstName returns the value of FirstName if it is set or its
// zero value if it is unset.
func (v *Response) GetFirstName() (o string) {
	if v != nil && v.FirstName != nil {
		return *v.FirstName
	}

	return
}

// IsSetFirstName returns true if FirstName is not nil.
func (v *Response) IsSetFirstName() bool {
	return v != nil && v.FirstName != nil
}

// GetLastName1 returns the value of LastName1 if it is set or its
// zero value if it is unset.
func (v *Response) GetLastName1() (o string) {
	if v != nil && v.LastName1 != nil {
		return *v.LastName1
	}

	return
}

// IsSetLastName1 returns true if LastName1 is not nil.
func (v *Response) IsSetLastName1() bool {
	return v != nil && v.LastName1 != nil
}

// Serverless_Beta_Args represents the arguments for the Serverless.beta function.
//
// The arguments for beta are sent and received over the wire as this struct.
type Serverless_Beta_Args struct {
	Request *Request `json:"request,omitempty"`
}

// ToWire translates a Serverless_Beta_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Serverless_Beta_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Request_Read(w wire.Value) (*Request, error) {
	var v Request
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Serverless_Beta_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Serverless_Beta_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Serverless_Beta_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Serverless_Beta_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _Request_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Serverless_Beta_Args
// struct.
func (v *Serverless_Beta_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Serverless_Beta_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Serverless_Beta_Args match the
// provided Serverless_Beta_Args.
//
// This function performs a deep comparison.
func (v *Serverless_Beta_Args) Equals(rhs *Serverless_Beta_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Serverless_Beta_Args.
func (v *Serverless_Beta_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Serverless_Beta_Args) GetRequest() (o *Request) {
	if v != nil && v.Request != nil {
		return v.Request
	}

	return
}

// IsSetRequest returns true if Request is not nil.
func (v *Serverless_Beta_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "beta" for this struct.
func (v *Serverless_Beta_Args) MethodName() string {
	return "beta"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Serverless_Beta_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Serverless_Beta_Helper provides functions that aid in handling the
// parameters and return values of the Serverless.beta
// function.
var Serverless_Beta_Helper = struct {
	// Args accepts the parameters of beta in-order and returns
	// the arguments struct for the function.
	Args func(
		request *Request,
	) *Serverless_Beta_Args

	// IsException returns true if the given error can be thrown
	// by beta.
	//
	// An error can be thrown by beta only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for beta
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// beta into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by beta
	//
	//   value, err := beta(args)
	//   result, err := Serverless_Beta_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from beta: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*Response, error) (*Serverless_Beta_Result, error)

	// UnwrapResponse takes the result struct for beta
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if beta threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Serverless_Beta_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Serverless_Beta_Result) (*Response, error)
}{}

func init() {
	Serverless_Beta_Helper.Args = func(
		request *Request,
	) *Serverless_Beta_Args {
		return &Serverless_Beta_Args{
			Request: request,
		}
	}

	Serverless_Beta_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Serverless_Beta_Helper.WrapResponse = func(success *Response, err error) (*Serverless_Beta_Result, error) {
		if err == nil {
			return &Serverless_Beta_Result{Success: success}, nil
		}

		return nil, err
	}
	Serverless_Beta_Helper.UnwrapResponse = func(result *Serverless_Beta_Result) (success *Response, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Serverless_Beta_Result represents the result of a Serverless.beta function call.
//
// The result of a beta execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Serverless_Beta_Result struct {
	// Value returned by beta after a successful execution.
	Success *Response `json:"success,omitempty"`
}

// ToWire translates a Serverless_Beta_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Serverless_Beta_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Serverless_Beta_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Response_Read(w wire.Value) (*Response, error) {
	var v Response
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Serverless_Beta_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Serverless_Beta_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Serverless_Beta_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Serverless_Beta_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _Response_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Serverless_Beta_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Serverless_Beta_Result
// struct.
func (v *Serverless_Beta_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Serverless_Beta_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Serverless_Beta_Result match the
// provided Serverless_Beta_Result.
//
// This function performs a deep comparison.
func (v *Serverless_Beta_Result) Equals(rhs *Serverless_Beta_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Serverless_Beta_Result.
func (v *Serverless_Beta_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Serverless_Beta_Result) GetSuccess() (o *Response) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Serverless_Beta_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "beta" for this struct.
func (v *Serverless_Beta_Result) MethodName() string {
	return "beta"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Serverless_Beta_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

// Serverless_ServerlessArgWithHeaders_Args represents the arguments for the Serverless.serverlessArgWithHeaders function.
//
// The arguments for serverlessArgWithHeaders are sent and received over the wire as this struct.
type Serverless_ServerlessArgWithHeaders_Args struct {
	Name     string  `json:"-"`
	UserUUID *string `json:"-"`
}

// ToWire translates a Serverless_ServerlessArgWithHeaders_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Serverless_ServerlessArgWithHeaders_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.UserUUID != nil {
		w, err = wire.NewValueString(*(v.UserUUID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Serverless_ServerlessArgWithHeaders_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Serverless_ServerlessArgWithHeaders_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Serverless_ServerlessArgWithHeaders_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Serverless_ServerlessArgWithHeaders_Args) FromWire(w wire.Value) error {
	var err error

	nameIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.UserUUID = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !nameIsSet {
		return errors.New("field Name of Serverless_ServerlessArgWithHeaders_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Serverless_ServerlessArgWithHeaders_Args
// struct.
func (v *Serverless_ServerlessArgWithHeaders_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	if v.UserUUID != nil {
		fields[i] = fmt.Sprintf("UserUUID: %v", *(v.UserUUID))
		i++
	}

	return fmt.Sprintf("Serverless_ServerlessArgWithHeaders_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Serverless_ServerlessArgWithHeaders_Args match the
// provided Serverless_ServerlessArgWithHeaders_Args.
//
// This function performs a deep comparison.
func (v *Serverless_ServerlessArgWithHeaders_Args) Equals(rhs *Serverless_ServerlessArgWithHeaders_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Name == rhs.Name) {
		return false
	}
	if !_String_EqualsPtr(v.UserUUID, rhs.UserUUID) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Serverless_ServerlessArgWithHeaders_Args.
func (v *Serverless_ServerlessArgWithHeaders_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("name", v.Name)
	if v.UserUUID != nil {
		enc.AddString("userUUID", *v.UserUUID)
	}
	return err
}

// GetName returns the value of Name if it is set or its
// zero value if it is unset.
func (v *Serverless_ServerlessArgWithHeaders_Args) GetName() (o string) {
	if v != nil {
		o = v.Name
	}
	return
}

// GetUserUUID returns the value of UserUUID if it is set or its
// zero value if it is unset.
func (v *Serverless_ServerlessArgWithHeaders_Args) GetUserUUID() (o string) {
	if v != nil && v.UserUUID != nil {
		return *v.UserUUID
	}

	return
}

// IsSetUserUUID returns true if UserUUID is not nil.
func (v *Serverless_ServerlessArgWithHeaders_Args) IsSetUserUUID() bool {
	return v != nil && v.UserUUID != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "serverlessArgWithHeaders" for this struct.
func (v *Serverless_ServerlessArgWithHeaders_Args) MethodName() string {
	return "serverlessArgWithHeaders"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Serverless_ServerlessArgWithHeaders_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Serverless_ServerlessArgWithHeaders_Helper provides functions that aid in handling the
// parameters and return values of the Serverless.serverlessArgWithHeaders
// function.
var Serverless_ServerlessArgWithHeaders_Helper = struct {
	// Args accepts the parameters of serverlessArgWithHeaders in-order and returns
	// the arguments struct for the function.
	Args func(
		name string,
		userUUID *string,
	) *Serverless_ServerlessArgWithHeaders_Args

	// IsException returns true if the given error can be thrown
	// by serverlessArgWithHeaders.
	//
	// An error can be thrown by serverlessArgWithHeaders only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for serverlessArgWithHeaders
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// serverlessArgWithHeaders into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by serverlessArgWithHeaders
	//
	//   value, err := serverlessArgWithHeaders(args)
	//   result, err := Serverless_ServerlessArgWithHeaders_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from serverlessArgWithHeaders: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*Response, error) (*Serverless_ServerlessArgWithHeaders_Result, error)

	// UnwrapResponse takes the result struct for serverlessArgWithHeaders
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if serverlessArgWithHeaders threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Serverless_ServerlessArgWithHeaders_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Serverless_ServerlessArgWithHeaders_Result) (*Response, error)
}{}

func init() {
	Serverless_ServerlessArgWithHeaders_Helper.Args = func(
		name string,
		userUUID *string,
	) *Serverless_ServerlessArgWithHeaders_Args {
		return &Serverless_ServerlessArgWithHeaders_Args{
			Name:     name,
			UserUUID: userUUID,
		}
	}

	Serverless_ServerlessArgWithHeaders_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Serverless_ServerlessArgWithHeaders_Helper.WrapResponse = func(success *Response, err error) (*Serverless_ServerlessArgWithHeaders_Result, error) {
		if err == nil {
			return &Serverless_ServerlessArgWithHeaders_Result{Success: success}, nil
		}

		return nil, err
	}
	Serverless_ServerlessArgWithHeaders_Helper.UnwrapResponse = func(result *Serverless_ServerlessArgWithHeaders_Result) (success *Response, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Serverless_ServerlessArgWithHeaders_Result represents the result of a Serverless.serverlessArgWithHeaders function call.
//
// The result of a serverlessArgWithHeaders execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Serverless_ServerlessArgWithHeaders_Result struct {
	// Value returned by serverlessArgWithHeaders after a successful execution.
	Success *Response `json:"success,omitempty"`
}

// ToWire translates a Serverless_ServerlessArgWithHeaders_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Serverless_ServerlessArgWithHeaders_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Serverless_ServerlessArgWithHeaders_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Serverless_ServerlessArgWithHeaders_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Serverless_ServerlessArgWithHeaders_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Serverless_ServerlessArgWithHeaders_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Serverless_ServerlessArgWithHeaders_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _Response_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Serverless_ServerlessArgWithHeaders_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Serverless_ServerlessArgWithHeaders_Result
// struct.
func (v *Serverless_ServerlessArgWithHeaders_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Serverless_ServerlessArgWithHeaders_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Serverless_ServerlessArgWithHeaders_Result match the
// provided Serverless_ServerlessArgWithHeaders_Result.
//
// This function performs a deep comparison.
func (v *Serverless_ServerlessArgWithHeaders_Result) Equals(rhs *Serverless_ServerlessArgWithHeaders_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Serverless_ServerlessArgWithHeaders_Result.
func (v *Serverless_ServerlessArgWithHeaders_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Serverless_ServerlessArgWithHeaders_Result) GetSuccess() (o *Response) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Serverless_ServerlessArgWithHeaders_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "serverlessArgWithHeaders" for this struct.
func (v *Serverless_ServerlessArgWithHeaders_Result) MethodName() string {
	return "serverlessArgWithHeaders"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Serverless_ServerlessArgWithHeaders_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
