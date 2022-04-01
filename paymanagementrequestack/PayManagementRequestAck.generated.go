package paymanagementrequestack

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PayManagementRequestAck is the fix50sp2 PayManagementRequestAck type, MsgType = DZ.
type PayManagementRequestAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PayManagementRequestAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PayManagementRequestAck {
	return PayManagementRequestAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PayManagementRequestAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PayManagementRequestAck initialized with the required fields for PayManagementRequestAck.
func New(payrequestid field.PayRequestIDField, payrequeststatus field.PayRequestStatusField) (m PayManagementRequestAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("DZ"))
	m.Set(payrequestid)
	m.Set(payrequeststatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PayManagementRequestAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "DZ", r
}

// SetPayRequestID sets PayRequestID, Tag 2812.
func (m PayManagementRequestAck) SetPayRequestID(v string) {
	m.Set(field.NewPayRequestID(v))
}

// SetPayRequestStatus sets PayRequestStatus, Tag 2813.
func (m PayManagementRequestAck) SetPayRequestStatus(v enum.PayRequestStatus) {
	m.Set(field.NewPayRequestStatus(v))
}

// GetPayRequestID gets PayRequestID, Tag 2812.
func (m PayManagementRequestAck) GetPayRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.PayRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayRequestStatus gets PayRequestStatus, Tag 2813.
func (m PayManagementRequestAck) GetPayRequestStatus() (v enum.PayRequestStatus, err quickfix.MessageRejectError) {
	var f field.PayRequestStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPayRequestID returns true if PayRequestID is present, Tag 2812.
func (m PayManagementRequestAck) HasPayRequestID() bool {
	return m.Has(tag.PayRequestID)
}

// HasPayRequestStatus returns true if PayRequestStatus is present, Tag 2813.
func (m PayManagementRequestAck) HasPayRequestStatus() bool {
	return m.Has(tag.PayRequestStatus)
}
