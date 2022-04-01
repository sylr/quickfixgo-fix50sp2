package allocationinstructionalertrequestack

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// AllocationInstructionAlertRequestAck is the fix50sp2 AllocationInstructionAlertRequestAck type, MsgType = DV.
type AllocationInstructionAlertRequestAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a AllocationInstructionAlertRequestAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) AllocationInstructionAlertRequestAck {
	return AllocationInstructionAlertRequestAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m AllocationInstructionAlertRequestAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a AllocationInstructionAlertRequestAck initialized with the required fields for AllocationInstructionAlertRequestAck.
func New(allocrequestid field.AllocRequestIDField, allocrequeststatus field.AllocRequestStatusField) (m AllocationInstructionAlertRequestAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("DV"))
	m.Set(allocrequestid)
	m.Set(allocrequeststatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg AllocationInstructionAlertRequestAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "DV", r
}

// SetRejectText sets RejectText, Tag 1328.
func (m AllocationInstructionAlertRequestAck) SetRejectText(v string) {
	m.Set(field.NewRejectText(v))
}

// SetEncodedRejectTextLen sets EncodedRejectTextLen, Tag 1664.
func (m AllocationInstructionAlertRequestAck) SetEncodedRejectTextLen(v int) {
	m.Set(field.NewEncodedRejectTextLen(v))
}

// SetEncodedRejectText sets EncodedRejectText, Tag 1665.
func (m AllocationInstructionAlertRequestAck) SetEncodedRejectText(v string) {
	m.Set(field.NewEncodedRejectText(v))
}

// SetAllocRequestID sets AllocRequestID, Tag 2758.
func (m AllocationInstructionAlertRequestAck) SetAllocRequestID(v string) {
	m.Set(field.NewAllocRequestID(v))
}

// SetAllocRequestStatus sets AllocRequestStatus, Tag 2768.
func (m AllocationInstructionAlertRequestAck) SetAllocRequestStatus(v enum.AllocRequestStatus) {
	m.Set(field.NewAllocRequestStatus(v))
}

// GetRejectText gets RejectText, Tag 1328.
func (m AllocationInstructionAlertRequestAck) GetRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.RejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectTextLen gets EncodedRejectTextLen, Tag 1664.
func (m AllocationInstructionAlertRequestAck) GetEncodedRejectTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectText gets EncodedRejectText, Tag 1665.
func (m AllocationInstructionAlertRequestAck) GetEncodedRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocRequestID gets AllocRequestID, Tag 2758.
func (m AllocationInstructionAlertRequestAck) GetAllocRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocRequestStatus gets AllocRequestStatus, Tag 2768.
func (m AllocationInstructionAlertRequestAck) GetAllocRequestStatus() (v enum.AllocRequestStatus, err quickfix.MessageRejectError) {
	var f field.AllocRequestStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRejectText returns true if RejectText is present, Tag 1328.
func (m AllocationInstructionAlertRequestAck) HasRejectText() bool {
	return m.Has(tag.RejectText)
}

// HasEncodedRejectTextLen returns true if EncodedRejectTextLen is present, Tag 1664.
func (m AllocationInstructionAlertRequestAck) HasEncodedRejectTextLen() bool {
	return m.Has(tag.EncodedRejectTextLen)
}

// HasEncodedRejectText returns true if EncodedRejectText is present, Tag 1665.
func (m AllocationInstructionAlertRequestAck) HasEncodedRejectText() bool {
	return m.Has(tag.EncodedRejectText)
}

// HasAllocRequestID returns true if AllocRequestID is present, Tag 2758.
func (m AllocationInstructionAlertRequestAck) HasAllocRequestID() bool {
	return m.Has(tag.AllocRequestID)
}

// HasAllocRequestStatus returns true if AllocRequestStatus is present, Tag 2768.
func (m AllocationInstructionAlertRequestAck) HasAllocRequestStatus() bool {
	return m.Has(tag.AllocRequestStatus)
}
