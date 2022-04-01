package quoteack

import (
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// QuoteAck is the fix50sp2 QuoteAck type, MsgType = CW.
type QuoteAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a QuoteAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) QuoteAck {
	return QuoteAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m QuoteAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a QuoteAck initialized with the required fields for QuoteAck.
func New(quoteackstatus field.QuoteAckStatusField) (m QuoteAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CW"))
	m.Set(quoteackstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg QuoteAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CW", r
}

// SetText sets Text, Tag 58.
func (m QuoteAck) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetQuoteID sets QuoteID, Tag 117.
func (m QuoteAck) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

// SetQuoteReqID sets QuoteReqID, Tag 131.
func (m QuoteAck) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

// SetQuoteCancelType sets QuoteCancelType, Tag 298.
func (m QuoteAck) SetQuoteCancelType(v enum.QuoteCancelType) {
	m.Set(field.NewQuoteCancelType(v))
}

// SetQuoteRejectReason sets QuoteRejectReason, Tag 300.
func (m QuoteAck) SetQuoteRejectReason(v enum.QuoteRejectReason) {
	m.Set(field.NewQuoteRejectReason(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m QuoteAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m QuoteAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m QuoteAck) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteType sets QuoteType, Tag 537.
func (m QuoteAck) SetQuoteType(v enum.QuoteType) {
	m.Set(field.NewQuoteType(v))
}

// SetQuoteMsgID sets QuoteMsgID, Tag 1166.
func (m QuoteAck) SetQuoteMsgID(v string) {
	m.Set(field.NewQuoteMsgID(v))
}

// SetRejectText sets RejectText, Tag 1328.
func (m QuoteAck) SetRejectText(v string) {
	m.Set(field.NewRejectText(v))
}

// SetEncodedRejectTextLen sets EncodedRejectTextLen, Tag 1664.
func (m QuoteAck) SetEncodedRejectTextLen(v int) {
	m.Set(field.NewEncodedRejectTextLen(v))
}

// SetEncodedRejectText sets EncodedRejectText, Tag 1665.
func (m QuoteAck) SetEncodedRejectText(v string) {
	m.Set(field.NewEncodedRejectText(v))
}

// SetSecondaryQuoteID sets SecondaryQuoteID, Tag 1751.
func (m QuoteAck) SetSecondaryQuoteID(v string) {
	m.Set(field.NewSecondaryQuoteID(v))
}

// SetQuoteAckStatus sets QuoteAckStatus, Tag 1865.
func (m QuoteAck) SetQuoteAckStatus(v enum.QuoteAckStatus) {
	m.Set(field.NewQuoteAckStatus(v))
}

// SetNoQuoteAttributes sets NoQuoteAttributes, Tag 2706.
func (m QuoteAck) SetNoQuoteAttributes(f NoQuoteAttributesRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m QuoteAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteID gets QuoteID, Tag 117.
func (m QuoteAck) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteReqID gets QuoteReqID, Tag 131.
func (m QuoteAck) GetQuoteReqID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteCancelType gets QuoteCancelType, Tag 298.
func (m QuoteAck) GetQuoteCancelType() (v enum.QuoteCancelType, err quickfix.MessageRejectError) {
	var f field.QuoteCancelTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteRejectReason gets QuoteRejectReason, Tag 300.
func (m QuoteAck) GetQuoteRejectReason() (v enum.QuoteRejectReason, err quickfix.MessageRejectError) {
	var f field.QuoteRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m QuoteAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m QuoteAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m QuoteAck) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteType gets QuoteType, Tag 537.
func (m QuoteAck) GetQuoteType() (v enum.QuoteType, err quickfix.MessageRejectError) {
	var f field.QuoteTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteMsgID gets QuoteMsgID, Tag 1166.
func (m QuoteAck) GetQuoteMsgID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteMsgIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRejectText gets RejectText, Tag 1328.
func (m QuoteAck) GetRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.RejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectTextLen gets EncodedRejectTextLen, Tag 1664.
func (m QuoteAck) GetEncodedRejectTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedRejectText gets EncodedRejectText, Tag 1665.
func (m QuoteAck) GetEncodedRejectText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedRejectTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryQuoteID gets SecondaryQuoteID, Tag 1751.
func (m QuoteAck) GetSecondaryQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryQuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteAckStatus gets QuoteAckStatus, Tag 1865.
func (m QuoteAck) GetQuoteAckStatus() (v enum.QuoteAckStatus, err quickfix.MessageRejectError) {
	var f field.QuoteAckStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoQuoteAttributes gets NoQuoteAttributes, Tag 2706.
func (m QuoteAck) GetNoQuoteAttributes() (NoQuoteAttributesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoQuoteAttributesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m QuoteAck) HasText() bool {
	return m.Has(tag.Text)
}

// HasQuoteID returns true if QuoteID is present, Tag 117.
func (m QuoteAck) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

// HasQuoteReqID returns true if QuoteReqID is present, Tag 131.
func (m QuoteAck) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

// HasQuoteCancelType returns true if QuoteCancelType is present, Tag 298.
func (m QuoteAck) HasQuoteCancelType() bool {
	return m.Has(tag.QuoteCancelType)
}

// HasQuoteRejectReason returns true if QuoteRejectReason is present, Tag 300.
func (m QuoteAck) HasQuoteRejectReason() bool {
	return m.Has(tag.QuoteRejectReason)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m QuoteAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m QuoteAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m QuoteAck) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasQuoteType returns true if QuoteType is present, Tag 537.
func (m QuoteAck) HasQuoteType() bool {
	return m.Has(tag.QuoteType)
}

// HasQuoteMsgID returns true if QuoteMsgID is present, Tag 1166.
func (m QuoteAck) HasQuoteMsgID() bool {
	return m.Has(tag.QuoteMsgID)
}

// HasRejectText returns true if RejectText is present, Tag 1328.
func (m QuoteAck) HasRejectText() bool {
	return m.Has(tag.RejectText)
}

// HasEncodedRejectTextLen returns true if EncodedRejectTextLen is present, Tag 1664.
func (m QuoteAck) HasEncodedRejectTextLen() bool {
	return m.Has(tag.EncodedRejectTextLen)
}

// HasEncodedRejectText returns true if EncodedRejectText is present, Tag 1665.
func (m QuoteAck) HasEncodedRejectText() bool {
	return m.Has(tag.EncodedRejectText)
}

// HasSecondaryQuoteID returns true if SecondaryQuoteID is present, Tag 1751.
func (m QuoteAck) HasSecondaryQuoteID() bool {
	return m.Has(tag.SecondaryQuoteID)
}

// HasQuoteAckStatus returns true if QuoteAckStatus is present, Tag 1865.
func (m QuoteAck) HasQuoteAckStatus() bool {
	return m.Has(tag.QuoteAckStatus)
}

// HasNoQuoteAttributes returns true if NoQuoteAttributes is present, Tag 2706.
func (m QuoteAck) HasNoQuoteAttributes() bool {
	return m.Has(tag.NoQuoteAttributes)
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetPartyRoleQualifier sets PartyRoleQualifier, Tag 2376.
func (m NoPartyIDs) SetPartyRoleQualifier(v int) {
	m.Set(field.NewPartyRoleQualifier(v))
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetPartyRoleQualifier gets PartyRoleQualifier, Tag 2376.
func (m NoPartyIDs) GetPartyRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.PartyRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// HasPartyRoleQualifier returns true if PartyRoleQualifier is present, Tag 2376.
func (m NoPartyIDs) HasPartyRoleQualifier() bool {
	return m.Has(tag.PartyRoleQualifier)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartySubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartySubID),
				quickfix.GroupElement(tag.PartySubIDType),
			},
		),
	}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyID),
				quickfix.GroupElement(tag.PartyIDSource),
				quickfix.GroupElement(tag.PartyRole),
				NewNoPartySubIDsRepeatingGroup(),
				quickfix.GroupElement(tag.PartyRoleQualifier),
			},
		),
	}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoQuoteAttributes is a repeating group element, Tag 2706.
type NoQuoteAttributes struct {
	*quickfix.Group
}

// SetQuoteAttributeType sets QuoteAttributeType, Tag 2707.
func (m NoQuoteAttributes) SetQuoteAttributeType(v enum.QuoteAttributeType) {
	m.Set(field.NewQuoteAttributeType(v))
}

// SetQuoteAttributeValue sets QuoteAttributeValue, Tag 2708.
func (m NoQuoteAttributes) SetQuoteAttributeValue(v string) {
	m.Set(field.NewQuoteAttributeValue(v))
}

// GetQuoteAttributeType gets QuoteAttributeType, Tag 2707.
func (m NoQuoteAttributes) GetQuoteAttributeType() (v enum.QuoteAttributeType, err quickfix.MessageRejectError) {
	var f field.QuoteAttributeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteAttributeValue gets QuoteAttributeValue, Tag 2708.
func (m NoQuoteAttributes) GetQuoteAttributeValue() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteAttributeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasQuoteAttributeType returns true if QuoteAttributeType is present, Tag 2707.
func (m NoQuoteAttributes) HasQuoteAttributeType() bool {
	return m.Has(tag.QuoteAttributeType)
}

// HasQuoteAttributeValue returns true if QuoteAttributeValue is present, Tag 2708.
func (m NoQuoteAttributes) HasQuoteAttributeValue() bool {
	return m.Has(tag.QuoteAttributeValue)
}

// NoQuoteAttributesRepeatingGroup is a repeating group, Tag 2706.
type NoQuoteAttributesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoQuoteAttributesRepeatingGroup returns an initialized, NoQuoteAttributesRepeatingGroup.
func NewNoQuoteAttributesRepeatingGroup() NoQuoteAttributesRepeatingGroup {
	return NoQuoteAttributesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoQuoteAttributes,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.QuoteAttributeType),
				quickfix.GroupElement(tag.QuoteAttributeValue),
			},
		),
	}
}

// Add create and append a new NoQuoteAttributes to this group.
func (m NoQuoteAttributesRepeatingGroup) Add() NoQuoteAttributes {
	g := m.RepeatingGroup.Add()
	return NoQuoteAttributes{g}
}

// Get returns the ith NoQuoteAttributes in the NoQuoteAttributesRepeatinGroup.
func (m NoQuoteAttributesRepeatingGroup) Get(i int) NoQuoteAttributes {
	return NoQuoteAttributes{m.RepeatingGroup.Get(i)}
}
