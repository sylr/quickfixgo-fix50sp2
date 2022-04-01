package tradematchreportack

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// TradeMatchReportAck is the fix50sp2 TradeMatchReportAck type, MsgType = DD.
type TradeMatchReportAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a TradeMatchReportAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) TradeMatchReportAck {
	return TradeMatchReportAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m TradeMatchReportAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a TradeMatchReportAck initialized with the required fields for TradeMatchReportAck.
func New(trdmatchid field.TrdMatchIDField, tradematchackstatus field.TradeMatchAckStatusField) (m TradeMatchReportAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("DD"))
	m.Set(trdmatchid)
	m.Set(tradematchackstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg TradeMatchReportAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "DD", r
}

// SetText sets Text, Tag 58.
func (m TradeMatchReportAck) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m TradeMatchReportAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m TradeMatchReportAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetTrdMatchID sets TrdMatchID, Tag 880.
func (m TradeMatchReportAck) SetTrdMatchID(v string) {
	m.Set(field.NewTrdMatchID(v))
}

// SetApplID sets ApplID, Tag 1180.
func (m TradeMatchReportAck) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181.
func (m TradeMatchReportAck) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetRejectText sets RejectText, Tag 1328.
func (m TradeMatchReportAck) SetRejectText(v string) {
	m.Set(field.NewRejectText(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350.
func (m TradeMatchReportAck) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352.
func (m TradeMatchReportAck) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// SetEncodedRejectTextLen sets EncodedRejectTextLen, Tag 1664.
func (m TradeMatchReportAck) SetEncodedRejectTextLen(v int) {
	m.Set(field.NewEncodedRejectTextLen(v))
}

// SetEncodedRejectText sets EncodedRejectText, Tag 1665.
func (m TradeMatchReportAck) SetEncodedRejectText(v string) {
	m.Set(field.NewEncodedRejectText(v))
}

// SetTradeMatchAckStatus sets TradeMatchAckStatus, Tag 1896.
func (m TradeMatchReportAck) SetTradeMatchAckStatus(v enum.TradeMatchAckStatus) {
	m.Set(field.NewTradeMatchAckStatus(v))
}

// SetTradeMatchRejectReason sets TradeMatchRejectReason, Tag 1897.
func (m TradeMatchReportAck) SetTradeMatchRejectReason(v enum.TradeMatchRejectReason) {
	m.Set(field.NewTradeMatchRejectReason(v))
}

// GetText gets Text, Tag 58.
func (m TradeMatchReportAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m TradeMatchReportAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m TradeMatchReportAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTrdMatchID gets TrdMatchID, Tag 880.
func (m TradeMatchReportAck) GetTrdMatchID() (v string, err quickfix.MessageRejectError) {
	var f field.TrdMatchIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180.
func (m TradeMatchReportAck) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181.
func (m TradeMatchReportAck) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRejectText gets RejectText, Tag 1328.
func (m TradeMatchReportAck) GetRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.RejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350.
func (m TradeMatchReportAck) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352.
func (m TradeMatchReportAck) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectTextLen gets EncodedRejectTextLen, Tag 1664.
func (m TradeMatchReportAck) GetEncodedRejectTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectText gets EncodedRejectText, Tag 1665.
func (m TradeMatchReportAck) GetEncodedRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeMatchAckStatus gets TradeMatchAckStatus, Tag 1896.
func (m TradeMatchReportAck) GetTradeMatchAckStatus() (v enum.TradeMatchAckStatus, err quickfix.MessageRejectError) {
	var f field.TradeMatchAckStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeMatchRejectReason gets TradeMatchRejectReason, Tag 1897.
func (m TradeMatchReportAck) GetTradeMatchRejectReason() (v enum.TradeMatchRejectReason, err quickfix.MessageRejectError) {
	var f field.TradeMatchRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m TradeMatchReportAck) HasText() bool {
	return m.Has(tag.Text)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m TradeMatchReportAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m TradeMatchReportAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasTrdMatchID returns true if TrdMatchID is present, Tag 880.
func (m TradeMatchReportAck) HasTrdMatchID() bool {
	return m.Has(tag.TrdMatchID)
}

// HasApplID returns true if ApplID is present, Tag 1180.
func (m TradeMatchReportAck) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181.
func (m TradeMatchReportAck) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasRejectText returns true if RejectText is present, Tag 1328.
func (m TradeMatchReportAck) HasRejectText() bool {
	return m.Has(tag.RejectText)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350.
func (m TradeMatchReportAck) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352.
func (m TradeMatchReportAck) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// HasEncodedRejectTextLen returns true if EncodedRejectTextLen is present, Tag 1664.
func (m TradeMatchReportAck) HasEncodedRejectTextLen() bool {
	return m.Has(tag.EncodedRejectTextLen)
}

// HasEncodedRejectText returns true if EncodedRejectText is present, Tag 1665.
func (m TradeMatchReportAck) HasEncodedRejectText() bool {
	return m.Has(tag.EncodedRejectText)
}

// HasTradeMatchAckStatus returns true if TradeMatchAckStatus is present, Tag 1896.
func (m TradeMatchReportAck) HasTradeMatchAckStatus() bool {
	return m.Has(tag.TradeMatchAckStatus)
}

// HasTradeMatchRejectReason returns true if TradeMatchRejectReason is present, Tag 1897.
func (m TradeMatchReportAck) HasTradeMatchRejectReason() bool {
	return m.Has(tag.TradeMatchRejectReason)
}
