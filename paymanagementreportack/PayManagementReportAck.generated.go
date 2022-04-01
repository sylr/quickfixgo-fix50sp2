package paymanagementreportack

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PayManagementReportAck is the fix50sp2 PayManagementReportAck type, MsgType = EB.
type PayManagementReportAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PayManagementReportAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PayManagementReportAck {
	return PayManagementReportAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PayManagementReportAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PayManagementReportAck initialized with the required fields for PayManagementReportAck.
func New(payreportid field.PayReportIDField, payreportstatus field.PayReportStatusField) (m PayManagementReportAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("EB"))
	m.Set(payreportid)
	m.Set(payreportstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PayManagementReportAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "EB", r
}

// SetRejectText sets RejectText, Tag 1328.
func (m PayManagementReportAck) SetRejectText(v string) {
	m.Set(field.NewRejectText(v))
}

// SetEncodedRejectTextLen sets EncodedRejectTextLen, Tag 1664.
func (m PayManagementReportAck) SetEncodedRejectTextLen(v int) {
	m.Set(field.NewEncodedRejectTextLen(v))
}

// SetEncodedRejectText sets EncodedRejectText, Tag 1665.
func (m PayManagementReportAck) SetEncodedRejectText(v string) {
	m.Set(field.NewEncodedRejectText(v))
}

// SetPayReportID sets PayReportID, Tag 2799.
func (m PayManagementReportAck) SetPayReportID(v string) {
	m.Set(field.NewPayReportID(v))
}

// SetPayDisputeReason sets PayDisputeReason, Tag 2800.
func (m PayManagementReportAck) SetPayDisputeReason(v int) {
	m.Set(field.NewPayDisputeReason(v))
}

// SetPayReportStatus sets PayReportStatus, Tag 2806.
func (m PayManagementReportAck) SetPayReportStatus(v enum.PayReportStatus) {
	m.Set(field.NewPayReportStatus(v))
}

// GetRejectText gets RejectText, Tag 1328.
func (m PayManagementReportAck) GetRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.RejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectTextLen gets EncodedRejectTextLen, Tag 1664.
func (m PayManagementReportAck) GetEncodedRejectTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectText gets EncodedRejectText, Tag 1665.
func (m PayManagementReportAck) GetEncodedRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayReportID gets PayReportID, Tag 2799.
func (m PayManagementReportAck) GetPayReportID() (v string, err quickfix.MessageRejectError) {
	var f field.PayReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayDisputeReason gets PayDisputeReason, Tag 2800.
func (m PayManagementReportAck) GetPayDisputeReason() (v int, err quickfix.MessageRejectError) {
	var f field.PayDisputeReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayReportStatus gets PayReportStatus, Tag 2806.
func (m PayManagementReportAck) GetPayReportStatus() (v enum.PayReportStatus, err quickfix.MessageRejectError) {
	var f field.PayReportStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRejectText returns true if RejectText is present, Tag 1328.
func (m PayManagementReportAck) HasRejectText() bool {
	return m.Has(tag.RejectText)
}

// HasEncodedRejectTextLen returns true if EncodedRejectTextLen is present, Tag 1664.
func (m PayManagementReportAck) HasEncodedRejectTextLen() bool {
	return m.Has(tag.EncodedRejectTextLen)
}

// HasEncodedRejectText returns true if EncodedRejectText is present, Tag 1665.
func (m PayManagementReportAck) HasEncodedRejectText() bool {
	return m.Has(tag.EncodedRejectText)
}

// HasPayReportID returns true if PayReportID is present, Tag 2799.
func (m PayManagementReportAck) HasPayReportID() bool {
	return m.Has(tag.PayReportID)
}

// HasPayDisputeReason returns true if PayDisputeReason is present, Tag 2800.
func (m PayManagementReportAck) HasPayDisputeReason() bool {
	return m.Has(tag.PayDisputeReason)
}

// HasPayReportStatus returns true if PayReportStatus is present, Tag 2806.
func (m PayManagementReportAck) HasPayReportStatus() bool {
	return m.Has(tag.PayReportStatus)
}
