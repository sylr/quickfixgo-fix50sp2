package partyactionrequest

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// PartyActionRequest is the fix50sp2 PartyActionRequest type, MsgType = DH.
type PartyActionRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a PartyActionRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) PartyActionRequest {
	return PartyActionRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m PartyActionRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a PartyActionRequest initialized with the required fields for PartyActionRequest.
func New(partyactionrequestid field.PartyActionRequestIDField, partyactiontype field.PartyActionTypeField) (m PartyActionRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("DH"))
	m.Set(partyactionrequestid)
	m.Set(partyactiontype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg PartyActionRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "DH", r
}

// SetText sets Text, Tag 58.
func (m PartyActionRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m PartyActionRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m PartyActionRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m PartyActionRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m PartyActionRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetMarketSegmentID sets MarketSegmentID, Tag 1300.
func (m PartyActionRequest) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

// SetMarketID sets MarketID, Tag 1301.
func (m PartyActionRequest) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

// SetInstrumentScopeSymbol sets InstrumentScopeSymbol, Tag 1536.
func (m PartyActionRequest) SetInstrumentScopeSymbol(v string) {
	m.Set(field.NewInstrumentScopeSymbol(v))
}

// SetInstrumentScopeSymbolSfx sets InstrumentScopeSymbolSfx, Tag 1537.
func (m PartyActionRequest) SetInstrumentScopeSymbolSfx(v string) {
	m.Set(field.NewInstrumentScopeSymbolSfx(v))
}

// SetInstrumentScopeSecurityID sets InstrumentScopeSecurityID, Tag 1538.
func (m PartyActionRequest) SetInstrumentScopeSecurityID(v string) {
	m.Set(field.NewInstrumentScopeSecurityID(v))
}

// SetInstrumentScopeSecurityIDSource sets InstrumentScopeSecurityIDSource, Tag 1539.
func (m PartyActionRequest) SetInstrumentScopeSecurityIDSource(v string) {
	m.Set(field.NewInstrumentScopeSecurityIDSource(v))
}

// SetNoInstrumentScopeSecurityAltID sets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m PartyActionRequest) SetNoInstrumentScopeSecurityAltID(f NoInstrumentScopeSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrumentScopeProduct sets InstrumentScopeProduct, Tag 1543.
func (m PartyActionRequest) SetInstrumentScopeProduct(v int) {
	m.Set(field.NewInstrumentScopeProduct(v))
}

// SetInstrumentScopeProductComplex sets InstrumentScopeProductComplex, Tag 1544.
func (m PartyActionRequest) SetInstrumentScopeProductComplex(v string) {
	m.Set(field.NewInstrumentScopeProductComplex(v))
}

// SetInstrumentScopeSecurityGroup sets InstrumentScopeSecurityGroup, Tag 1545.
func (m PartyActionRequest) SetInstrumentScopeSecurityGroup(v string) {
	m.Set(field.NewInstrumentScopeSecurityGroup(v))
}

// SetInstrumentScopeCFICode sets InstrumentScopeCFICode, Tag 1546.
func (m PartyActionRequest) SetInstrumentScopeCFICode(v string) {
	m.Set(field.NewInstrumentScopeCFICode(v))
}

// SetInstrumentScopeSecurityType sets InstrumentScopeSecurityType, Tag 1547.
func (m PartyActionRequest) SetInstrumentScopeSecurityType(v string) {
	m.Set(field.NewInstrumentScopeSecurityType(v))
}

// SetInstrumentScopeSecuritySubType sets InstrumentScopeSecuritySubType, Tag 1548.
func (m PartyActionRequest) SetInstrumentScopeSecuritySubType(v string) {
	m.Set(field.NewInstrumentScopeSecuritySubType(v))
}

// SetInstrumentScopeMaturityMonthYear sets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m PartyActionRequest) SetInstrumentScopeMaturityMonthYear(v string) {
	m.Set(field.NewInstrumentScopeMaturityMonthYear(v))
}

// SetInstrumentScopeMaturityTime sets InstrumentScopeMaturityTime, Tag 1550.
func (m PartyActionRequest) SetInstrumentScopeMaturityTime(v string) {
	m.Set(field.NewInstrumentScopeMaturityTime(v))
}

// SetInstrumentScopeRestructuringType sets InstrumentScopeRestructuringType, Tag 1551.
func (m PartyActionRequest) SetInstrumentScopeRestructuringType(v string) {
	m.Set(field.NewInstrumentScopeRestructuringType(v))
}

// SetInstrumentScopeSeniority sets InstrumentScopeSeniority, Tag 1552.
func (m PartyActionRequest) SetInstrumentScopeSeniority(v string) {
	m.Set(field.NewInstrumentScopeSeniority(v))
}

// SetInstrumentScopePutOrCall sets InstrumentScopePutOrCall, Tag 1553.
func (m PartyActionRequest) SetInstrumentScopePutOrCall(v int) {
	m.Set(field.NewInstrumentScopePutOrCall(v))
}

// SetInstrumentScopeFlexibleIndicator sets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m PartyActionRequest) SetInstrumentScopeFlexibleIndicator(v bool) {
	m.Set(field.NewInstrumentScopeFlexibleIndicator(v))
}

// SetInstrumentScopeCouponRate sets InstrumentScopeCouponRate, Tag 1555.
func (m PartyActionRequest) SetInstrumentScopeCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewInstrumentScopeCouponRate(value, scale))
}

// SetInstrumentScopeSecurityDesc sets InstrumentScopeSecurityDesc, Tag 1556.
func (m PartyActionRequest) SetInstrumentScopeSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeSecurityDesc(v))
}

// SetInstrumentScopeSettlType sets InstrumentScopeSettlType, Tag 1557.
func (m PartyActionRequest) SetInstrumentScopeSettlType(v string) {
	m.Set(field.NewInstrumentScopeSettlType(v))
}

// SetNoRelatedPartyDetailID sets NoRelatedPartyDetailID, Tag 1562.
func (m PartyActionRequest) SetNoRelatedPartyDetailID(f NoRelatedPartyDetailIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrumentScopeSecurityExchange sets InstrumentScopeSecurityExchange, Tag 1616.
func (m PartyActionRequest) SetInstrumentScopeSecurityExchange(v string) {
	m.Set(field.NewInstrumentScopeSecurityExchange(v))
}

// SetInstrumentScopeEncodedSecurityDescLen sets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m PartyActionRequest) SetInstrumentScopeEncodedSecurityDescLen(v int) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDescLen(v))
}

// SetInstrumentScopeEncodedSecurityDesc sets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m PartyActionRequest) SetInstrumentScopeEncodedSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDesc(v))
}

// SetNoRequestingPartyIDs sets NoRequestingPartyIDs, Tag 1657.
func (m PartyActionRequest) SetNoRequestingPartyIDs(f NoRequestingPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetPartyActionRequestID sets PartyActionRequestID, Tag 2328.
func (m PartyActionRequest) SetPartyActionRequestID(v string) {
	m.Set(field.NewPartyActionRequestID(v))
}

// SetPartyActionType sets PartyActionType, Tag 2329.
func (m PartyActionRequest) SetPartyActionType(v enum.PartyActionType) {
	m.Set(field.NewPartyActionType(v))
}

// SetApplTestMessageIndicator sets ApplTestMessageIndicator, Tag 2330.
func (m PartyActionRequest) SetApplTestMessageIndicator(v bool) {
	m.Set(field.NewApplTestMessageIndicator(v))
}

// SetInstrumentScopeUPICode sets InstrumentScopeUPICode, Tag 2895.
func (m PartyActionRequest) SetInstrumentScopeUPICode(v string) {
	m.Set(field.NewInstrumentScopeUPICode(v))
}

// GetText gets Text, Tag 58.
func (m PartyActionRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m PartyActionRequest) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m PartyActionRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m PartyActionRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m PartyActionRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMarketSegmentID gets MarketSegmentID, Tag 1300.
func (m PartyActionRequest) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketID gets MarketID, Tag 1301.
func (m PartyActionRequest) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbol gets InstrumentScopeSymbol, Tag 1536.
func (m PartyActionRequest) GetInstrumentScopeSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbolSfx gets InstrumentScopeSymbolSfx, Tag 1537.
func (m PartyActionRequest) GetInstrumentScopeSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityID gets InstrumentScopeSecurityID, Tag 1538.
func (m PartyActionRequest) GetInstrumentScopeSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityIDSource gets InstrumentScopeSecurityIDSource, Tag 1539.
func (m PartyActionRequest) GetInstrumentScopeSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentScopeSecurityAltID gets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m PartyActionRequest) GetNoInstrumentScopeSecurityAltID() (NoInstrumentScopeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentScopeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrumentScopeProduct gets InstrumentScopeProduct, Tag 1543.
func (m PartyActionRequest) GetInstrumentScopeProduct() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeProductComplex gets InstrumentScopeProductComplex, Tag 1544.
func (m PartyActionRequest) GetInstrumentScopeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityGroup gets InstrumentScopeSecurityGroup, Tag 1545.
func (m PartyActionRequest) GetInstrumentScopeSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCFICode gets InstrumentScopeCFICode, Tag 1546.
func (m PartyActionRequest) GetInstrumentScopeCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityType gets InstrumentScopeSecurityType, Tag 1547.
func (m PartyActionRequest) GetInstrumentScopeSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecuritySubType gets InstrumentScopeSecuritySubType, Tag 1548.
func (m PartyActionRequest) GetInstrumentScopeSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityMonthYear gets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m PartyActionRequest) GetInstrumentScopeMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityTime gets InstrumentScopeMaturityTime, Tag 1550.
func (m PartyActionRequest) GetInstrumentScopeMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeRestructuringType gets InstrumentScopeRestructuringType, Tag 1551.
func (m PartyActionRequest) GetInstrumentScopeRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSeniority gets InstrumentScopeSeniority, Tag 1552.
func (m PartyActionRequest) GetInstrumentScopeSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopePutOrCall gets InstrumentScopePutOrCall, Tag 1553.
func (m PartyActionRequest) GetInstrumentScopePutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopePutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeFlexibleIndicator gets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m PartyActionRequest) GetInstrumentScopeFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeFlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCouponRate gets InstrumentScopeCouponRate, Tag 1555.
func (m PartyActionRequest) GetInstrumentScopeCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityDesc gets InstrumentScopeSecurityDesc, Tag 1556.
func (m PartyActionRequest) GetInstrumentScopeSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSettlType gets InstrumentScopeSettlType, Tag 1557.
func (m PartyActionRequest) GetInstrumentScopeSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedPartyDetailID gets NoRelatedPartyDetailID, Tag 1562.
func (m PartyActionRequest) GetNoRelatedPartyDetailID() (NoRelatedPartyDetailIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrumentScopeSecurityExchange gets InstrumentScopeSecurityExchange, Tag 1616.
func (m PartyActionRequest) GetInstrumentScopeSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDescLen gets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m PartyActionRequest) GetInstrumentScopeEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDesc gets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m PartyActionRequest) GetInstrumentScopeEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRequestingPartyIDs gets NoRequestingPartyIDs, Tag 1657.
func (m PartyActionRequest) GetNoRequestingPartyIDs() (NoRequestingPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestingPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetPartyActionRequestID gets PartyActionRequestID, Tag 2328.
func (m PartyActionRequest) GetPartyActionRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyActionRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyActionType gets PartyActionType, Tag 2329.
func (m PartyActionRequest) GetPartyActionType() (v enum.PartyActionType, err quickfix.MessageRejectError) {
	var f field.PartyActionTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplTestMessageIndicator gets ApplTestMessageIndicator, Tag 2330.
func (m PartyActionRequest) GetApplTestMessageIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplTestMessageIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeUPICode gets InstrumentScopeUPICode, Tag 2895.
func (m PartyActionRequest) GetInstrumentScopeUPICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeUPICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m PartyActionRequest) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m PartyActionRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m PartyActionRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m PartyActionRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m PartyActionRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300.
func (m PartyActionRequest) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

// HasMarketID returns true if MarketID is present, Tag 1301.
func (m PartyActionRequest) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

// HasInstrumentScopeSymbol returns true if InstrumentScopeSymbol is present, Tag 1536.
func (m PartyActionRequest) HasInstrumentScopeSymbol() bool {
	return m.Has(tag.InstrumentScopeSymbol)
}

// HasInstrumentScopeSymbolSfx returns true if InstrumentScopeSymbolSfx is present, Tag 1537.
func (m PartyActionRequest) HasInstrumentScopeSymbolSfx() bool {
	return m.Has(tag.InstrumentScopeSymbolSfx)
}

// HasInstrumentScopeSecurityID returns true if InstrumentScopeSecurityID is present, Tag 1538.
func (m PartyActionRequest) HasInstrumentScopeSecurityID() bool {
	return m.Has(tag.InstrumentScopeSecurityID)
}

// HasInstrumentScopeSecurityIDSource returns true if InstrumentScopeSecurityIDSource is present, Tag 1539.
func (m PartyActionRequest) HasInstrumentScopeSecurityIDSource() bool {
	return m.Has(tag.InstrumentScopeSecurityIDSource)
}

// HasNoInstrumentScopeSecurityAltID returns true if NoInstrumentScopeSecurityAltID is present, Tag 1540.
func (m PartyActionRequest) HasNoInstrumentScopeSecurityAltID() bool {
	return m.Has(tag.NoInstrumentScopeSecurityAltID)
}

// HasInstrumentScopeProduct returns true if InstrumentScopeProduct is present, Tag 1543.
func (m PartyActionRequest) HasInstrumentScopeProduct() bool {
	return m.Has(tag.InstrumentScopeProduct)
}

// HasInstrumentScopeProductComplex returns true if InstrumentScopeProductComplex is present, Tag 1544.
func (m PartyActionRequest) HasInstrumentScopeProductComplex() bool {
	return m.Has(tag.InstrumentScopeProductComplex)
}

// HasInstrumentScopeSecurityGroup returns true if InstrumentScopeSecurityGroup is present, Tag 1545.
func (m PartyActionRequest) HasInstrumentScopeSecurityGroup() bool {
	return m.Has(tag.InstrumentScopeSecurityGroup)
}

// HasInstrumentScopeCFICode returns true if InstrumentScopeCFICode is present, Tag 1546.
func (m PartyActionRequest) HasInstrumentScopeCFICode() bool {
	return m.Has(tag.InstrumentScopeCFICode)
}

// HasInstrumentScopeSecurityType returns true if InstrumentScopeSecurityType is present, Tag 1547.
func (m PartyActionRequest) HasInstrumentScopeSecurityType() bool {
	return m.Has(tag.InstrumentScopeSecurityType)
}

// HasInstrumentScopeSecuritySubType returns true if InstrumentScopeSecuritySubType is present, Tag 1548.
func (m PartyActionRequest) HasInstrumentScopeSecuritySubType() bool {
	return m.Has(tag.InstrumentScopeSecuritySubType)
}

// HasInstrumentScopeMaturityMonthYear returns true if InstrumentScopeMaturityMonthYear is present, Tag 1549.
func (m PartyActionRequest) HasInstrumentScopeMaturityMonthYear() bool {
	return m.Has(tag.InstrumentScopeMaturityMonthYear)
}

// HasInstrumentScopeMaturityTime returns true if InstrumentScopeMaturityTime is present, Tag 1550.
func (m PartyActionRequest) HasInstrumentScopeMaturityTime() bool {
	return m.Has(tag.InstrumentScopeMaturityTime)
}

// HasInstrumentScopeRestructuringType returns true if InstrumentScopeRestructuringType is present, Tag 1551.
func (m PartyActionRequest) HasInstrumentScopeRestructuringType() bool {
	return m.Has(tag.InstrumentScopeRestructuringType)
}

// HasInstrumentScopeSeniority returns true if InstrumentScopeSeniority is present, Tag 1552.
func (m PartyActionRequest) HasInstrumentScopeSeniority() bool {
	return m.Has(tag.InstrumentScopeSeniority)
}

// HasInstrumentScopePutOrCall returns true if InstrumentScopePutOrCall is present, Tag 1553.
func (m PartyActionRequest) HasInstrumentScopePutOrCall() bool {
	return m.Has(tag.InstrumentScopePutOrCall)
}

// HasInstrumentScopeFlexibleIndicator returns true if InstrumentScopeFlexibleIndicator is present, Tag 1554.
func (m PartyActionRequest) HasInstrumentScopeFlexibleIndicator() bool {
	return m.Has(tag.InstrumentScopeFlexibleIndicator)
}

// HasInstrumentScopeCouponRate returns true if InstrumentScopeCouponRate is present, Tag 1555.
func (m PartyActionRequest) HasInstrumentScopeCouponRate() bool {
	return m.Has(tag.InstrumentScopeCouponRate)
}

// HasInstrumentScopeSecurityDesc returns true if InstrumentScopeSecurityDesc is present, Tag 1556.
func (m PartyActionRequest) HasInstrumentScopeSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeSecurityDesc)
}

// HasInstrumentScopeSettlType returns true if InstrumentScopeSettlType is present, Tag 1557.
func (m PartyActionRequest) HasInstrumentScopeSettlType() bool {
	return m.Has(tag.InstrumentScopeSettlType)
}

// HasNoRelatedPartyDetailID returns true if NoRelatedPartyDetailID is present, Tag 1562.
func (m PartyActionRequest) HasNoRelatedPartyDetailID() bool {
	return m.Has(tag.NoRelatedPartyDetailID)
}

// HasInstrumentScopeSecurityExchange returns true if InstrumentScopeSecurityExchange is present, Tag 1616.
func (m PartyActionRequest) HasInstrumentScopeSecurityExchange() bool {
	return m.Has(tag.InstrumentScopeSecurityExchange)
}

// HasInstrumentScopeEncodedSecurityDescLen returns true if InstrumentScopeEncodedSecurityDescLen is present, Tag 1620.
func (m PartyActionRequest) HasInstrumentScopeEncodedSecurityDescLen() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDescLen)
}

// HasInstrumentScopeEncodedSecurityDesc returns true if InstrumentScopeEncodedSecurityDesc is present, Tag 1621.
func (m PartyActionRequest) HasInstrumentScopeEncodedSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDesc)
}

// HasNoRequestingPartyIDs returns true if NoRequestingPartyIDs is present, Tag 1657.
func (m PartyActionRequest) HasNoRequestingPartyIDs() bool {
	return m.Has(tag.NoRequestingPartyIDs)
}

// HasPartyActionRequestID returns true if PartyActionRequestID is present, Tag 2328.
func (m PartyActionRequest) HasPartyActionRequestID() bool {
	return m.Has(tag.PartyActionRequestID)
}

// HasPartyActionType returns true if PartyActionType is present, Tag 2329.
func (m PartyActionRequest) HasPartyActionType() bool {
	return m.Has(tag.PartyActionType)
}

// HasApplTestMessageIndicator returns true if ApplTestMessageIndicator is present, Tag 2330.
func (m PartyActionRequest) HasApplTestMessageIndicator() bool {
	return m.Has(tag.ApplTestMessageIndicator)
}

// HasInstrumentScopeUPICode returns true if InstrumentScopeUPICode is present, Tag 2895.
func (m PartyActionRequest) HasInstrumentScopeUPICode() bool {
	return m.Has(tag.InstrumentScopeUPICode)
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

// NoInstrumentScopeSecurityAltID is a repeating group element, Tag 1540.
type NoInstrumentScopeSecurityAltID struct {
	*quickfix.Group
}

// SetInstrumentScopeSecurityAltID sets InstrumentScopeSecurityAltID, Tag 1541.
func (m NoInstrumentScopeSecurityAltID) SetInstrumentScopeSecurityAltID(v string) {
	m.Set(field.NewInstrumentScopeSecurityAltID(v))
}

// SetInstrumentScopeSecurityAltIDSource sets InstrumentScopeSecurityAltIDSource, Tag 1542.
func (m NoInstrumentScopeSecurityAltID) SetInstrumentScopeSecurityAltIDSource(v string) {
	m.Set(field.NewInstrumentScopeSecurityAltIDSource(v))
}

// GetInstrumentScopeSecurityAltID gets InstrumentScopeSecurityAltID, Tag 1541.
func (m NoInstrumentScopeSecurityAltID) GetInstrumentScopeSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityAltIDSource gets InstrumentScopeSecurityAltIDSource, Tag 1542.
func (m NoInstrumentScopeSecurityAltID) GetInstrumentScopeSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentScopeSecurityAltID returns true if InstrumentScopeSecurityAltID is present, Tag 1541.
func (m NoInstrumentScopeSecurityAltID) HasInstrumentScopeSecurityAltID() bool {
	return m.Has(tag.InstrumentScopeSecurityAltID)
}

// HasInstrumentScopeSecurityAltIDSource returns true if InstrumentScopeSecurityAltIDSource is present, Tag 1542.
func (m NoInstrumentScopeSecurityAltID) HasInstrumentScopeSecurityAltIDSource() bool {
	return m.Has(tag.InstrumentScopeSecurityAltIDSource)
}

// NoInstrumentScopeSecurityAltIDRepeatingGroup is a repeating group, Tag 1540.
type NoInstrumentScopeSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentScopeSecurityAltIDRepeatingGroup returns an initialized, NoInstrumentScopeSecurityAltIDRepeatingGroup.
func NewNoInstrumentScopeSecurityAltIDRepeatingGroup() NoInstrumentScopeSecurityAltIDRepeatingGroup {
	return NoInstrumentScopeSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoInstrumentScopeSecurityAltID,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.InstrumentScopeSecurityAltID),
				quickfix.GroupElement(tag.InstrumentScopeSecurityAltIDSource),
			},
		),
	}
}

// Add create and append a new NoInstrumentScopeSecurityAltID to this group.
func (m NoInstrumentScopeSecurityAltIDRepeatingGroup) Add() NoInstrumentScopeSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoInstrumentScopeSecurityAltID{g}
}

// Get returns the ith NoInstrumentScopeSecurityAltID in the NoInstrumentScopeSecurityAltIDRepeatinGroup.
func (m NoInstrumentScopeSecurityAltIDRepeatingGroup) Get(i int) NoInstrumentScopeSecurityAltID {
	return NoInstrumentScopeSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailID is a repeating group element, Tag 1562.
type NoRelatedPartyDetailID struct {
	*quickfix.Group
}

// SetRelatedPartyDetailID sets RelatedPartyDetailID, Tag 1563.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailID(v string) {
	m.Set(field.NewRelatedPartyDetailID(v))
}

// SetRelatedPartyDetailIDSource sets RelatedPartyDetailIDSource, Tag 1564.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailIDSource(v string) {
	m.Set(field.NewRelatedPartyDetailIDSource(v))
}

// SetRelatedPartyDetailRole sets RelatedPartyDetailRole, Tag 1565.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailRole(v int) {
	m.Set(field.NewRelatedPartyDetailRole(v))
}

// SetRelatedPartyDetailRoleQualifier sets RelatedPartyDetailRoleQualifier, Tag 1675.
func (m NoRelatedPartyDetailID) SetRelatedPartyDetailRoleQualifier(v int) {
	m.Set(field.NewRelatedPartyDetailRoleQualifier(v))
}

// SetNoRelatedPartyDetailSubIDs sets NoRelatedPartyDetailSubIDs, Tag 1566.
func (m NoRelatedPartyDetailID) SetNoRelatedPartyDetailSubIDs(f NoRelatedPartyDetailSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoRelatedPartyDetailAltID sets NoRelatedPartyDetailAltID, Tag 1569.
func (m NoRelatedPartyDetailID) SetNoRelatedPartyDetailAltID(f NoRelatedPartyDetailAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoPartyRelationships sets NoPartyRelationships, Tag 1514.
func (m NoRelatedPartyDetailID) SetNoPartyRelationships(f NoPartyRelationshipsRepeatingGroup) {
	m.SetGroup(f)
}

// GetRelatedPartyDetailID gets RelatedPartyDetailID, Tag 1563.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailIDSource gets RelatedPartyDetailIDSource, Tag 1564.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailRole gets RelatedPartyDetailRole, Tag 1565.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailRole() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailRoleQualifier gets RelatedPartyDetailRoleQualifier, Tag 1675.
func (m NoRelatedPartyDetailID) GetRelatedPartyDetailRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedPartyDetailSubIDs gets NoRelatedPartyDetailSubIDs, Tag 1566.
func (m NoRelatedPartyDetailID) GetNoRelatedPartyDetailSubIDs() (NoRelatedPartyDetailSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoRelatedPartyDetailAltID gets NoRelatedPartyDetailAltID, Tag 1569.
func (m NoRelatedPartyDetailID) GetNoRelatedPartyDetailAltID() (NoRelatedPartyDetailAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoPartyRelationships gets NoPartyRelationships, Tag 1514.
func (m NoRelatedPartyDetailID) GetNoPartyRelationships() (NoPartyRelationshipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyRelationshipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasRelatedPartyDetailID returns true if RelatedPartyDetailID is present, Tag 1563.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailID() bool {
	return m.Has(tag.RelatedPartyDetailID)
}

// HasRelatedPartyDetailIDSource returns true if RelatedPartyDetailIDSource is present, Tag 1564.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailIDSource() bool {
	return m.Has(tag.RelatedPartyDetailIDSource)
}

// HasRelatedPartyDetailRole returns true if RelatedPartyDetailRole is present, Tag 1565.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailRole() bool {
	return m.Has(tag.RelatedPartyDetailRole)
}

// HasRelatedPartyDetailRoleQualifier returns true if RelatedPartyDetailRoleQualifier is present, Tag 1675.
func (m NoRelatedPartyDetailID) HasRelatedPartyDetailRoleQualifier() bool {
	return m.Has(tag.RelatedPartyDetailRoleQualifier)
}

// HasNoRelatedPartyDetailSubIDs returns true if NoRelatedPartyDetailSubIDs is present, Tag 1566.
func (m NoRelatedPartyDetailID) HasNoRelatedPartyDetailSubIDs() bool {
	return m.Has(tag.NoRelatedPartyDetailSubIDs)
}

// HasNoRelatedPartyDetailAltID returns true if NoRelatedPartyDetailAltID is present, Tag 1569.
func (m NoRelatedPartyDetailID) HasNoRelatedPartyDetailAltID() bool {
	return m.Has(tag.NoRelatedPartyDetailAltID)
}

// HasNoPartyRelationships returns true if NoPartyRelationships is present, Tag 1514.
func (m NoRelatedPartyDetailID) HasNoPartyRelationships() bool {
	return m.Has(tag.NoPartyRelationships)
}

// NoRelatedPartyDetailSubIDs is a repeating group element, Tag 1566.
type NoRelatedPartyDetailSubIDs struct {
	*quickfix.Group
}

// SetRelatedPartyDetailSubID sets RelatedPartyDetailSubID, Tag 1567.
func (m NoRelatedPartyDetailSubIDs) SetRelatedPartyDetailSubID(v string) {
	m.Set(field.NewRelatedPartyDetailSubID(v))
}

// SetRelatedPartyDetailSubIDType sets RelatedPartyDetailSubIDType, Tag 1568.
func (m NoRelatedPartyDetailSubIDs) SetRelatedPartyDetailSubIDType(v int) {
	m.Set(field.NewRelatedPartyDetailSubIDType(v))
}

// GetRelatedPartyDetailSubID gets RelatedPartyDetailSubID, Tag 1567.
func (m NoRelatedPartyDetailSubIDs) GetRelatedPartyDetailSubID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailSubIDType gets RelatedPartyDetailSubIDType, Tag 1568.
func (m NoRelatedPartyDetailSubIDs) GetRelatedPartyDetailSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRelatedPartyDetailSubID returns true if RelatedPartyDetailSubID is present, Tag 1567.
func (m NoRelatedPartyDetailSubIDs) HasRelatedPartyDetailSubID() bool {
	return m.Has(tag.RelatedPartyDetailSubID)
}

// HasRelatedPartyDetailSubIDType returns true if RelatedPartyDetailSubIDType is present, Tag 1568.
func (m NoRelatedPartyDetailSubIDs) HasRelatedPartyDetailSubIDType() bool {
	return m.Has(tag.RelatedPartyDetailSubIDType)
}

// NoRelatedPartyDetailSubIDsRepeatingGroup is a repeating group, Tag 1566.
type NoRelatedPartyDetailSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailSubIDsRepeatingGroup returns an initialized, NoRelatedPartyDetailSubIDsRepeatingGroup.
func NewNoRelatedPartyDetailSubIDsRepeatingGroup() NoRelatedPartyDetailSubIDsRepeatingGroup {
	return NoRelatedPartyDetailSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailSubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailSubID),
				quickfix.GroupElement(tag.RelatedPartyDetailSubIDType),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailSubIDs to this group.
func (m NoRelatedPartyDetailSubIDsRepeatingGroup) Add() NoRelatedPartyDetailSubIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailSubIDs{g}
}

// Get returns the ith NoRelatedPartyDetailSubIDs in the NoRelatedPartyDetailSubIDsRepeatinGroup.
func (m NoRelatedPartyDetailSubIDsRepeatingGroup) Get(i int) NoRelatedPartyDetailSubIDs {
	return NoRelatedPartyDetailSubIDs{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailAltID is a repeating group element, Tag 1569.
type NoRelatedPartyDetailAltID struct {
	*quickfix.Group
}

// SetRelatedPartyDetailAltID sets RelatedPartyDetailAltID, Tag 1570.
func (m NoRelatedPartyDetailAltID) SetRelatedPartyDetailAltID(v string) {
	m.Set(field.NewRelatedPartyDetailAltID(v))
}

// SetRelatedPartyDetailAltIDSource sets RelatedPartyDetailAltIDSource, Tag 1571.
func (m NoRelatedPartyDetailAltID) SetRelatedPartyDetailAltIDSource(v string) {
	m.Set(field.NewRelatedPartyDetailAltIDSource(v))
}

// SetNoRelatedPartyDetailAltSubIDs sets NoRelatedPartyDetailAltSubIDs, Tag 1572.
func (m NoRelatedPartyDetailAltID) SetNoRelatedPartyDetailAltSubIDs(f NoRelatedPartyDetailAltSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetRelatedPartyDetailAltID gets RelatedPartyDetailAltID, Tag 1570.
func (m NoRelatedPartyDetailAltID) GetRelatedPartyDetailAltID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailAltIDSource gets RelatedPartyDetailAltIDSource, Tag 1571.
func (m NoRelatedPartyDetailAltID) GetRelatedPartyDetailAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedPartyDetailAltSubIDs gets NoRelatedPartyDetailAltSubIDs, Tag 1572.
func (m NoRelatedPartyDetailAltID) GetNoRelatedPartyDetailAltSubIDs() (NoRelatedPartyDetailAltSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyDetailAltSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasRelatedPartyDetailAltID returns true if RelatedPartyDetailAltID is present, Tag 1570.
func (m NoRelatedPartyDetailAltID) HasRelatedPartyDetailAltID() bool {
	return m.Has(tag.RelatedPartyDetailAltID)
}

// HasRelatedPartyDetailAltIDSource returns true if RelatedPartyDetailAltIDSource is present, Tag 1571.
func (m NoRelatedPartyDetailAltID) HasRelatedPartyDetailAltIDSource() bool {
	return m.Has(tag.RelatedPartyDetailAltIDSource)
}

// HasNoRelatedPartyDetailAltSubIDs returns true if NoRelatedPartyDetailAltSubIDs is present, Tag 1572.
func (m NoRelatedPartyDetailAltID) HasNoRelatedPartyDetailAltSubIDs() bool {
	return m.Has(tag.NoRelatedPartyDetailAltSubIDs)
}

// NoRelatedPartyDetailAltSubIDs is a repeating group element, Tag 1572.
type NoRelatedPartyDetailAltSubIDs struct {
	*quickfix.Group
}

// SetRelatedPartyDetailAltSubID sets RelatedPartyDetailAltSubID, Tag 1573.
func (m NoRelatedPartyDetailAltSubIDs) SetRelatedPartyDetailAltSubID(v string) {
	m.Set(field.NewRelatedPartyDetailAltSubID(v))
}

// SetRelatedPartyDetailAltSubIDType sets RelatedPartyDetailAltSubIDType, Tag 1574.
func (m NoRelatedPartyDetailAltSubIDs) SetRelatedPartyDetailAltSubIDType(v int) {
	m.Set(field.NewRelatedPartyDetailAltSubIDType(v))
}

// GetRelatedPartyDetailAltSubID gets RelatedPartyDetailAltSubID, Tag 1573.
func (m NoRelatedPartyDetailAltSubIDs) GetRelatedPartyDetailAltSubID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRelatedPartyDetailAltSubIDType gets RelatedPartyDetailAltSubIDType, Tag 1574.
func (m NoRelatedPartyDetailAltSubIDs) GetRelatedPartyDetailAltSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyDetailAltSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRelatedPartyDetailAltSubID returns true if RelatedPartyDetailAltSubID is present, Tag 1573.
func (m NoRelatedPartyDetailAltSubIDs) HasRelatedPartyDetailAltSubID() bool {
	return m.Has(tag.RelatedPartyDetailAltSubID)
}

// HasRelatedPartyDetailAltSubIDType returns true if RelatedPartyDetailAltSubIDType is present, Tag 1574.
func (m NoRelatedPartyDetailAltSubIDs) HasRelatedPartyDetailAltSubIDType() bool {
	return m.Has(tag.RelatedPartyDetailAltSubIDType)
}

// NoRelatedPartyDetailAltSubIDsRepeatingGroup is a repeating group, Tag 1572.
type NoRelatedPartyDetailAltSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailAltSubIDsRepeatingGroup returns an initialized, NoRelatedPartyDetailAltSubIDsRepeatingGroup.
func NewNoRelatedPartyDetailAltSubIDsRepeatingGroup() NoRelatedPartyDetailAltSubIDsRepeatingGroup {
	return NoRelatedPartyDetailAltSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailAltSubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailAltSubID),
				quickfix.GroupElement(tag.RelatedPartyDetailAltSubIDType),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailAltSubIDs to this group.
func (m NoRelatedPartyDetailAltSubIDsRepeatingGroup) Add() NoRelatedPartyDetailAltSubIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailAltSubIDs{g}
}

// Get returns the ith NoRelatedPartyDetailAltSubIDs in the NoRelatedPartyDetailAltSubIDsRepeatinGroup.
func (m NoRelatedPartyDetailAltSubIDsRepeatingGroup) Get(i int) NoRelatedPartyDetailAltSubIDs {
	return NoRelatedPartyDetailAltSubIDs{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailAltIDRepeatingGroup is a repeating group, Tag 1569.
type NoRelatedPartyDetailAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailAltIDRepeatingGroup returns an initialized, NoRelatedPartyDetailAltIDRepeatingGroup.
func NewNoRelatedPartyDetailAltIDRepeatingGroup() NoRelatedPartyDetailAltIDRepeatingGroup {
	return NoRelatedPartyDetailAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailAltID,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailAltID),
				quickfix.GroupElement(tag.RelatedPartyDetailAltIDSource),
				NewNoRelatedPartyDetailAltSubIDsRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailAltID to this group.
func (m NoRelatedPartyDetailAltIDRepeatingGroup) Add() NoRelatedPartyDetailAltID {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailAltID{g}
}

// Get returns the ith NoRelatedPartyDetailAltID in the NoRelatedPartyDetailAltIDRepeatinGroup.
func (m NoRelatedPartyDetailAltIDRepeatingGroup) Get(i int) NoRelatedPartyDetailAltID {
	return NoRelatedPartyDetailAltID{m.RepeatingGroup.Get(i)}
}

// NoPartyRelationships is a repeating group element, Tag 1514.
type NoPartyRelationships struct {
	*quickfix.Group
}

// SetPartyRelationship sets PartyRelationship, Tag 1515.
func (m NoPartyRelationships) SetPartyRelationship(v enum.PartyRelationship) {
	m.Set(field.NewPartyRelationship(v))
}

// GetPartyRelationship gets PartyRelationship, Tag 1515.
func (m NoPartyRelationships) GetPartyRelationship() (v enum.PartyRelationship, err quickfix.MessageRejectError) {
	var f field.PartyRelationshipField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyRelationship returns true if PartyRelationship is present, Tag 1515.
func (m NoPartyRelationships) HasPartyRelationship() bool {
	return m.Has(tag.PartyRelationship)
}

// NoPartyRelationshipsRepeatingGroup is a repeating group, Tag 1514.
type NoPartyRelationshipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyRelationshipsRepeatingGroup returns an initialized, NoPartyRelationshipsRepeatingGroup.
func NewNoPartyRelationshipsRepeatingGroup() NoPartyRelationshipsRepeatingGroup {
	return NoPartyRelationshipsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPartyRelationships,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PartyRelationship),
			},
		),
	}
}

// Add create and append a new NoPartyRelationships to this group.
func (m NoPartyRelationshipsRepeatingGroup) Add() NoPartyRelationships {
	g := m.RepeatingGroup.Add()
	return NoPartyRelationships{g}
}

// Get returns the ith NoPartyRelationships in the NoPartyRelationshipsRepeatinGroup.
func (m NoPartyRelationshipsRepeatingGroup) Get(i int) NoPartyRelationships {
	return NoPartyRelationships{m.RepeatingGroup.Get(i)}
}

// NoRelatedPartyDetailIDRepeatingGroup is a repeating group, Tag 1562.
type NoRelatedPartyDetailIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedPartyDetailIDRepeatingGroup returns an initialized, NoRelatedPartyDetailIDRepeatingGroup.
func NewNoRelatedPartyDetailIDRepeatingGroup() NoRelatedPartyDetailIDRepeatingGroup {
	return NoRelatedPartyDetailIDRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedPartyDetailID,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedPartyDetailID),
				quickfix.GroupElement(tag.RelatedPartyDetailIDSource),
				quickfix.GroupElement(tag.RelatedPartyDetailRole),
				quickfix.GroupElement(tag.RelatedPartyDetailRoleQualifier),
				NewNoRelatedPartyDetailSubIDsRepeatingGroup(),
				NewNoRelatedPartyDetailAltIDRepeatingGroup(),
				NewNoPartyRelationshipsRepeatingGroup(),
			},
		),
	}
}

// Add create and append a new NoRelatedPartyDetailID to this group.
func (m NoRelatedPartyDetailIDRepeatingGroup) Add() NoRelatedPartyDetailID {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyDetailID{g}
}

// Get returns the ith NoRelatedPartyDetailID in the NoRelatedPartyDetailIDRepeatinGroup.
func (m NoRelatedPartyDetailIDRepeatingGroup) Get(i int) NoRelatedPartyDetailID {
	return NoRelatedPartyDetailID{m.RepeatingGroup.Get(i)}
}

// NoRequestingPartyIDs is a repeating group element, Tag 1657.
type NoRequestingPartyIDs struct {
	*quickfix.Group
}

// SetRequestingPartyID sets RequestingPartyID, Tag 1658.
func (m NoRequestingPartyIDs) SetRequestingPartyID(v string) {
	m.Set(field.NewRequestingPartyID(v))
}

// SetRequestingPartyIDSource sets RequestingPartyIDSource, Tag 1659.
func (m NoRequestingPartyIDs) SetRequestingPartyIDSource(v string) {
	m.Set(field.NewRequestingPartyIDSource(v))
}

// SetRequestingPartyRole sets RequestingPartyRole, Tag 1660.
func (m NoRequestingPartyIDs) SetRequestingPartyRole(v int) {
	m.Set(field.NewRequestingPartyRole(v))
}

// SetNoRequestingPartySubIDs sets NoRequestingPartySubIDs, Tag 1661.
func (m NoRequestingPartyIDs) SetNoRequestingPartySubIDs(f NoRequestingPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRequestingPartyRoleQualifier sets RequestingPartyRoleQualifier, Tag 2338.
func (m NoRequestingPartyIDs) SetRequestingPartyRoleQualifier(v int) {
	m.Set(field.NewRequestingPartyRoleQualifier(v))
}

// GetRequestingPartyID gets RequestingPartyID, Tag 1658.
func (m NoRequestingPartyIDs) GetRequestingPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.RequestingPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestingPartyIDSource gets RequestingPartyIDSource, Tag 1659.
func (m NoRequestingPartyIDs) GetRequestingPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RequestingPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestingPartyRole gets RequestingPartyRole, Tag 1660.
func (m NoRequestingPartyIDs) GetRequestingPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.RequestingPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRequestingPartySubIDs gets NoRequestingPartySubIDs, Tag 1661.
func (m NoRequestingPartyIDs) GetNoRequestingPartySubIDs() (NoRequestingPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestingPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRequestingPartyRoleQualifier gets RequestingPartyRoleQualifier, Tag 2338.
func (m NoRequestingPartyIDs) GetRequestingPartyRoleQualifier() (v int, err quickfix.MessageRejectError) {
	var f field.RequestingPartyRoleQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRequestingPartyID returns true if RequestingPartyID is present, Tag 1658.
func (m NoRequestingPartyIDs) HasRequestingPartyID() bool {
	return m.Has(tag.RequestingPartyID)
}

// HasRequestingPartyIDSource returns true if RequestingPartyIDSource is present, Tag 1659.
func (m NoRequestingPartyIDs) HasRequestingPartyIDSource() bool {
	return m.Has(tag.RequestingPartyIDSource)
}

// HasRequestingPartyRole returns true if RequestingPartyRole is present, Tag 1660.
func (m NoRequestingPartyIDs) HasRequestingPartyRole() bool {
	return m.Has(tag.RequestingPartyRole)
}

// HasNoRequestingPartySubIDs returns true if NoRequestingPartySubIDs is present, Tag 1661.
func (m NoRequestingPartyIDs) HasNoRequestingPartySubIDs() bool {
	return m.Has(tag.NoRequestingPartySubIDs)
}

// HasRequestingPartyRoleQualifier returns true if RequestingPartyRoleQualifier is present, Tag 2338.
func (m NoRequestingPartyIDs) HasRequestingPartyRoleQualifier() bool {
	return m.Has(tag.RequestingPartyRoleQualifier)
}

// NoRequestingPartySubIDs is a repeating group element, Tag 1661.
type NoRequestingPartySubIDs struct {
	*quickfix.Group
}

// SetRequestingPartySubID sets RequestingPartySubID, Tag 1662.
func (m NoRequestingPartySubIDs) SetRequestingPartySubID(v string) {
	m.Set(field.NewRequestingPartySubID(v))
}

// SetRequestingPartySubIDType sets RequestingPartySubIDType, Tag 1663.
func (m NoRequestingPartySubIDs) SetRequestingPartySubIDType(v int) {
	m.Set(field.NewRequestingPartySubIDType(v))
}

// GetRequestingPartySubID gets RequestingPartySubID, Tag 1662.
func (m NoRequestingPartySubIDs) GetRequestingPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.RequestingPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRequestingPartySubIDType gets RequestingPartySubIDType, Tag 1663.
func (m NoRequestingPartySubIDs) GetRequestingPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RequestingPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRequestingPartySubID returns true if RequestingPartySubID is present, Tag 1662.
func (m NoRequestingPartySubIDs) HasRequestingPartySubID() bool {
	return m.Has(tag.RequestingPartySubID)
}

// HasRequestingPartySubIDType returns true if RequestingPartySubIDType is present, Tag 1663.
func (m NoRequestingPartySubIDs) HasRequestingPartySubIDType() bool {
	return m.Has(tag.RequestingPartySubIDType)
}

// NoRequestingPartySubIDsRepeatingGroup is a repeating group, Tag 1661.
type NoRequestingPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRequestingPartySubIDsRepeatingGroup returns an initialized, NoRequestingPartySubIDsRepeatingGroup.
func NewNoRequestingPartySubIDsRepeatingGroup() NoRequestingPartySubIDsRepeatingGroup {
	return NoRequestingPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRequestingPartySubIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RequestingPartySubID),
				quickfix.GroupElement(tag.RequestingPartySubIDType),
			},
		),
	}
}

// Add create and append a new NoRequestingPartySubIDs to this group.
func (m NoRequestingPartySubIDsRepeatingGroup) Add() NoRequestingPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoRequestingPartySubIDs{g}
}

// Get returns the ith NoRequestingPartySubIDs in the NoRequestingPartySubIDsRepeatinGroup.
func (m NoRequestingPartySubIDsRepeatingGroup) Get(i int) NoRequestingPartySubIDs {
	return NoRequestingPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoRequestingPartyIDsRepeatingGroup is a repeating group, Tag 1657.
type NoRequestingPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRequestingPartyIDsRepeatingGroup returns an initialized, NoRequestingPartyIDsRepeatingGroup.
func NewNoRequestingPartyIDsRepeatingGroup() NoRequestingPartyIDsRepeatingGroup {
	return NoRequestingPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRequestingPartyIDs,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RequestingPartyID),
				quickfix.GroupElement(tag.RequestingPartyIDSource),
				quickfix.GroupElement(tag.RequestingPartyRole),
				NewNoRequestingPartySubIDsRepeatingGroup(),
				quickfix.GroupElement(tag.RequestingPartyRoleQualifier),
			},
		),
	}
}

// Add create and append a new NoRequestingPartyIDs to this group.
func (m NoRequestingPartyIDsRepeatingGroup) Add() NoRequestingPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoRequestingPartyIDs{g}
}

// Get returns the ith NoRequestingPartyIDs in the NoRequestingPartyIDsRepeatinGroup.
func (m NoRequestingPartyIDsRepeatingGroup) Get(i int) NoRequestingPartyIDs {
	return NoRequestingPartyIDs{m.RepeatingGroup.Get(i)}
}
