package securitymassstatusrequest

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// SecurityMassStatusRequest is the fix50sp2 SecurityMassStatusRequest type, MsgType = CN.
type SecurityMassStatusRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a SecurityMassStatusRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) SecurityMassStatusRequest {
	return SecurityMassStatusRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m SecurityMassStatusRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a SecurityMassStatusRequest initialized with the required fields for SecurityMassStatusRequest.
func New(securitystatusreqid field.SecurityStatusReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m SecurityMassStatusRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CN"))
	m.Set(securitystatusreqid)
	m.Set(subscriptionrequesttype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg SecurityMassStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CN", r
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m SecurityMassStatusRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetSecurityStatusReqID sets SecurityStatusReqID, Tag 324.
func (m SecurityMassStatusRequest) SetSecurityStatusReqID(v string) {
	m.Set(field.NewSecurityStatusReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m SecurityMassStatusRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m SecurityMassStatusRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetMarketSegmentID sets MarketSegmentID, Tag 1300.
func (m SecurityMassStatusRequest) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

// SetMarketID sets MarketID, Tag 1301.
func (m SecurityMassStatusRequest) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

// SetSecurityListID sets SecurityListID, Tag 1465.
func (m SecurityMassStatusRequest) SetSecurityListID(v string) {
	m.Set(field.NewSecurityListID(v))
}

// SetInstrumentScopeSymbol sets InstrumentScopeSymbol, Tag 1536.
func (m SecurityMassStatusRequest) SetInstrumentScopeSymbol(v string) {
	m.Set(field.NewInstrumentScopeSymbol(v))
}

// SetInstrumentScopeSymbolSfx sets InstrumentScopeSymbolSfx, Tag 1537.
func (m SecurityMassStatusRequest) SetInstrumentScopeSymbolSfx(v string) {
	m.Set(field.NewInstrumentScopeSymbolSfx(v))
}

// SetInstrumentScopeSecurityID sets InstrumentScopeSecurityID, Tag 1538.
func (m SecurityMassStatusRequest) SetInstrumentScopeSecurityID(v string) {
	m.Set(field.NewInstrumentScopeSecurityID(v))
}

// SetInstrumentScopeSecurityIDSource sets InstrumentScopeSecurityIDSource, Tag 1539.
func (m SecurityMassStatusRequest) SetInstrumentScopeSecurityIDSource(v string) {
	m.Set(field.NewInstrumentScopeSecurityIDSource(v))
}

// SetNoInstrumentScopeSecurityAltID sets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m SecurityMassStatusRequest) SetNoInstrumentScopeSecurityAltID(f NoInstrumentScopeSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrumentScopeProduct sets InstrumentScopeProduct, Tag 1543.
func (m SecurityMassStatusRequest) SetInstrumentScopeProduct(v int) {
	m.Set(field.NewInstrumentScopeProduct(v))
}

// SetInstrumentScopeProductComplex sets InstrumentScopeProductComplex, Tag 1544.
func (m SecurityMassStatusRequest) SetInstrumentScopeProductComplex(v string) {
	m.Set(field.NewInstrumentScopeProductComplex(v))
}

// SetInstrumentScopeSecurityGroup sets InstrumentScopeSecurityGroup, Tag 1545.
func (m SecurityMassStatusRequest) SetInstrumentScopeSecurityGroup(v string) {
	m.Set(field.NewInstrumentScopeSecurityGroup(v))
}

// SetInstrumentScopeCFICode sets InstrumentScopeCFICode, Tag 1546.
func (m SecurityMassStatusRequest) SetInstrumentScopeCFICode(v string) {
	m.Set(field.NewInstrumentScopeCFICode(v))
}

// SetInstrumentScopeSecurityType sets InstrumentScopeSecurityType, Tag 1547.
func (m SecurityMassStatusRequest) SetInstrumentScopeSecurityType(v string) {
	m.Set(field.NewInstrumentScopeSecurityType(v))
}

// SetInstrumentScopeSecuritySubType sets InstrumentScopeSecuritySubType, Tag 1548.
func (m SecurityMassStatusRequest) SetInstrumentScopeSecuritySubType(v string) {
	m.Set(field.NewInstrumentScopeSecuritySubType(v))
}

// SetInstrumentScopeMaturityMonthYear sets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m SecurityMassStatusRequest) SetInstrumentScopeMaturityMonthYear(v string) {
	m.Set(field.NewInstrumentScopeMaturityMonthYear(v))
}

// SetInstrumentScopeMaturityTime sets InstrumentScopeMaturityTime, Tag 1550.
func (m SecurityMassStatusRequest) SetInstrumentScopeMaturityTime(v string) {
	m.Set(field.NewInstrumentScopeMaturityTime(v))
}

// SetInstrumentScopeRestructuringType sets InstrumentScopeRestructuringType, Tag 1551.
func (m SecurityMassStatusRequest) SetInstrumentScopeRestructuringType(v string) {
	m.Set(field.NewInstrumentScopeRestructuringType(v))
}

// SetInstrumentScopeSeniority sets InstrumentScopeSeniority, Tag 1552.
func (m SecurityMassStatusRequest) SetInstrumentScopeSeniority(v string) {
	m.Set(field.NewInstrumentScopeSeniority(v))
}

// SetInstrumentScopePutOrCall sets InstrumentScopePutOrCall, Tag 1553.
func (m SecurityMassStatusRequest) SetInstrumentScopePutOrCall(v int) {
	m.Set(field.NewInstrumentScopePutOrCall(v))
}

// SetInstrumentScopeFlexibleIndicator sets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m SecurityMassStatusRequest) SetInstrumentScopeFlexibleIndicator(v bool) {
	m.Set(field.NewInstrumentScopeFlexibleIndicator(v))
}

// SetInstrumentScopeCouponRate sets InstrumentScopeCouponRate, Tag 1555.
func (m SecurityMassStatusRequest) SetInstrumentScopeCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewInstrumentScopeCouponRate(value, scale))
}

// SetInstrumentScopeSecurityDesc sets InstrumentScopeSecurityDesc, Tag 1556.
func (m SecurityMassStatusRequest) SetInstrumentScopeSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeSecurityDesc(v))
}

// SetInstrumentScopeSettlType sets InstrumentScopeSettlType, Tag 1557.
func (m SecurityMassStatusRequest) SetInstrumentScopeSettlType(v string) {
	m.Set(field.NewInstrumentScopeSettlType(v))
}

// SetInstrumentScopeSecurityExchange sets InstrumentScopeSecurityExchange, Tag 1616.
func (m SecurityMassStatusRequest) SetInstrumentScopeSecurityExchange(v string) {
	m.Set(field.NewInstrumentScopeSecurityExchange(v))
}

// SetInstrumentScopeEncodedSecurityDescLen sets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m SecurityMassStatusRequest) SetInstrumentScopeEncodedSecurityDescLen(v int) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDescLen(v))
}

// SetInstrumentScopeEncodedSecurityDesc sets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m SecurityMassStatusRequest) SetInstrumentScopeEncodedSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDesc(v))
}

// SetInstrumentScopeUPICode sets InstrumentScopeUPICode, Tag 2895.
func (m SecurityMassStatusRequest) SetInstrumentScopeUPICode(v string) {
	m.Set(field.NewInstrumentScopeUPICode(v))
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m SecurityMassStatusRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityStatusReqID gets SecurityStatusReqID, Tag 324.
func (m SecurityMassStatusRequest) GetSecurityStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m SecurityMassStatusRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m SecurityMassStatusRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentID gets MarketSegmentID, Tag 1300.
func (m SecurityMassStatusRequest) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketID gets MarketID, Tag 1301.
func (m SecurityMassStatusRequest) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityListID gets SecurityListID, Tag 1465.
func (m SecurityMassStatusRequest) GetSecurityListID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbol gets InstrumentScopeSymbol, Tag 1536.
func (m SecurityMassStatusRequest) GetInstrumentScopeSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbolSfx gets InstrumentScopeSymbolSfx, Tag 1537.
func (m SecurityMassStatusRequest) GetInstrumentScopeSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityID gets InstrumentScopeSecurityID, Tag 1538.
func (m SecurityMassStatusRequest) GetInstrumentScopeSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityIDSource gets InstrumentScopeSecurityIDSource, Tag 1539.
func (m SecurityMassStatusRequest) GetInstrumentScopeSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentScopeSecurityAltID gets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m SecurityMassStatusRequest) GetNoInstrumentScopeSecurityAltID() (NoInstrumentScopeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentScopeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrumentScopeProduct gets InstrumentScopeProduct, Tag 1543.
func (m SecurityMassStatusRequest) GetInstrumentScopeProduct() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeProductComplex gets InstrumentScopeProductComplex, Tag 1544.
func (m SecurityMassStatusRequest) GetInstrumentScopeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityGroup gets InstrumentScopeSecurityGroup, Tag 1545.
func (m SecurityMassStatusRequest) GetInstrumentScopeSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCFICode gets InstrumentScopeCFICode, Tag 1546.
func (m SecurityMassStatusRequest) GetInstrumentScopeCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityType gets InstrumentScopeSecurityType, Tag 1547.
func (m SecurityMassStatusRequest) GetInstrumentScopeSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecuritySubType gets InstrumentScopeSecuritySubType, Tag 1548.
func (m SecurityMassStatusRequest) GetInstrumentScopeSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityMonthYear gets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m SecurityMassStatusRequest) GetInstrumentScopeMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityTime gets InstrumentScopeMaturityTime, Tag 1550.
func (m SecurityMassStatusRequest) GetInstrumentScopeMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeRestructuringType gets InstrumentScopeRestructuringType, Tag 1551.
func (m SecurityMassStatusRequest) GetInstrumentScopeRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSeniority gets InstrumentScopeSeniority, Tag 1552.
func (m SecurityMassStatusRequest) GetInstrumentScopeSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopePutOrCall gets InstrumentScopePutOrCall, Tag 1553.
func (m SecurityMassStatusRequest) GetInstrumentScopePutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopePutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeFlexibleIndicator gets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m SecurityMassStatusRequest) GetInstrumentScopeFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeFlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCouponRate gets InstrumentScopeCouponRate, Tag 1555.
func (m SecurityMassStatusRequest) GetInstrumentScopeCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityDesc gets InstrumentScopeSecurityDesc, Tag 1556.
func (m SecurityMassStatusRequest) GetInstrumentScopeSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSettlType gets InstrumentScopeSettlType, Tag 1557.
func (m SecurityMassStatusRequest) GetInstrumentScopeSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityExchange gets InstrumentScopeSecurityExchange, Tag 1616.
func (m SecurityMassStatusRequest) GetInstrumentScopeSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDescLen gets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m SecurityMassStatusRequest) GetInstrumentScopeEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDesc gets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m SecurityMassStatusRequest) GetInstrumentScopeEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeUPICode gets InstrumentScopeUPICode, Tag 2895.
func (m SecurityMassStatusRequest) GetInstrumentScopeUPICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeUPICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m SecurityMassStatusRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasSecurityStatusReqID returns true if SecurityStatusReqID is present, Tag 324.
func (m SecurityMassStatusRequest) HasSecurityStatusReqID() bool {
	return m.Has(tag.SecurityStatusReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m SecurityMassStatusRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m SecurityMassStatusRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300.
func (m SecurityMassStatusRequest) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

// HasMarketID returns true if MarketID is present, Tag 1301.
func (m SecurityMassStatusRequest) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

// HasSecurityListID returns true if SecurityListID is present, Tag 1465.
func (m SecurityMassStatusRequest) HasSecurityListID() bool {
	return m.Has(tag.SecurityListID)
}

// HasInstrumentScopeSymbol returns true if InstrumentScopeSymbol is present, Tag 1536.
func (m SecurityMassStatusRequest) HasInstrumentScopeSymbol() bool {
	return m.Has(tag.InstrumentScopeSymbol)
}

// HasInstrumentScopeSymbolSfx returns true if InstrumentScopeSymbolSfx is present, Tag 1537.
func (m SecurityMassStatusRequest) HasInstrumentScopeSymbolSfx() bool {
	return m.Has(tag.InstrumentScopeSymbolSfx)
}

// HasInstrumentScopeSecurityID returns true if InstrumentScopeSecurityID is present, Tag 1538.
func (m SecurityMassStatusRequest) HasInstrumentScopeSecurityID() bool {
	return m.Has(tag.InstrumentScopeSecurityID)
}

// HasInstrumentScopeSecurityIDSource returns true if InstrumentScopeSecurityIDSource is present, Tag 1539.
func (m SecurityMassStatusRequest) HasInstrumentScopeSecurityIDSource() bool {
	return m.Has(tag.InstrumentScopeSecurityIDSource)
}

// HasNoInstrumentScopeSecurityAltID returns true if NoInstrumentScopeSecurityAltID is present, Tag 1540.
func (m SecurityMassStatusRequest) HasNoInstrumentScopeSecurityAltID() bool {
	return m.Has(tag.NoInstrumentScopeSecurityAltID)
}

// HasInstrumentScopeProduct returns true if InstrumentScopeProduct is present, Tag 1543.
func (m SecurityMassStatusRequest) HasInstrumentScopeProduct() bool {
	return m.Has(tag.InstrumentScopeProduct)
}

// HasInstrumentScopeProductComplex returns true if InstrumentScopeProductComplex is present, Tag 1544.
func (m SecurityMassStatusRequest) HasInstrumentScopeProductComplex() bool {
	return m.Has(tag.InstrumentScopeProductComplex)
}

// HasInstrumentScopeSecurityGroup returns true if InstrumentScopeSecurityGroup is present, Tag 1545.
func (m SecurityMassStatusRequest) HasInstrumentScopeSecurityGroup() bool {
	return m.Has(tag.InstrumentScopeSecurityGroup)
}

// HasInstrumentScopeCFICode returns true if InstrumentScopeCFICode is present, Tag 1546.
func (m SecurityMassStatusRequest) HasInstrumentScopeCFICode() bool {
	return m.Has(tag.InstrumentScopeCFICode)
}

// HasInstrumentScopeSecurityType returns true if InstrumentScopeSecurityType is present, Tag 1547.
func (m SecurityMassStatusRequest) HasInstrumentScopeSecurityType() bool {
	return m.Has(tag.InstrumentScopeSecurityType)
}

// HasInstrumentScopeSecuritySubType returns true if InstrumentScopeSecuritySubType is present, Tag 1548.
func (m SecurityMassStatusRequest) HasInstrumentScopeSecuritySubType() bool {
	return m.Has(tag.InstrumentScopeSecuritySubType)
}

// HasInstrumentScopeMaturityMonthYear returns true if InstrumentScopeMaturityMonthYear is present, Tag 1549.
func (m SecurityMassStatusRequest) HasInstrumentScopeMaturityMonthYear() bool {
	return m.Has(tag.InstrumentScopeMaturityMonthYear)
}

// HasInstrumentScopeMaturityTime returns true if InstrumentScopeMaturityTime is present, Tag 1550.
func (m SecurityMassStatusRequest) HasInstrumentScopeMaturityTime() bool {
	return m.Has(tag.InstrumentScopeMaturityTime)
}

// HasInstrumentScopeRestructuringType returns true if InstrumentScopeRestructuringType is present, Tag 1551.
func (m SecurityMassStatusRequest) HasInstrumentScopeRestructuringType() bool {
	return m.Has(tag.InstrumentScopeRestructuringType)
}

// HasInstrumentScopeSeniority returns true if InstrumentScopeSeniority is present, Tag 1552.
func (m SecurityMassStatusRequest) HasInstrumentScopeSeniority() bool {
	return m.Has(tag.InstrumentScopeSeniority)
}

// HasInstrumentScopePutOrCall returns true if InstrumentScopePutOrCall is present, Tag 1553.
func (m SecurityMassStatusRequest) HasInstrumentScopePutOrCall() bool {
	return m.Has(tag.InstrumentScopePutOrCall)
}

// HasInstrumentScopeFlexibleIndicator returns true if InstrumentScopeFlexibleIndicator is present, Tag 1554.
func (m SecurityMassStatusRequest) HasInstrumentScopeFlexibleIndicator() bool {
	return m.Has(tag.InstrumentScopeFlexibleIndicator)
}

// HasInstrumentScopeCouponRate returns true if InstrumentScopeCouponRate is present, Tag 1555.
func (m SecurityMassStatusRequest) HasInstrumentScopeCouponRate() bool {
	return m.Has(tag.InstrumentScopeCouponRate)
}

// HasInstrumentScopeSecurityDesc returns true if InstrumentScopeSecurityDesc is present, Tag 1556.
func (m SecurityMassStatusRequest) HasInstrumentScopeSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeSecurityDesc)
}

// HasInstrumentScopeSettlType returns true if InstrumentScopeSettlType is present, Tag 1557.
func (m SecurityMassStatusRequest) HasInstrumentScopeSettlType() bool {
	return m.Has(tag.InstrumentScopeSettlType)
}

// HasInstrumentScopeSecurityExchange returns true if InstrumentScopeSecurityExchange is present, Tag 1616.
func (m SecurityMassStatusRequest) HasInstrumentScopeSecurityExchange() bool {
	return m.Has(tag.InstrumentScopeSecurityExchange)
}

// HasInstrumentScopeEncodedSecurityDescLen returns true if InstrumentScopeEncodedSecurityDescLen is present, Tag 1620.
func (m SecurityMassStatusRequest) HasInstrumentScopeEncodedSecurityDescLen() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDescLen)
}

// HasInstrumentScopeEncodedSecurityDesc returns true if InstrumentScopeEncodedSecurityDesc is present, Tag 1621.
func (m SecurityMassStatusRequest) HasInstrumentScopeEncodedSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDesc)
}

// HasInstrumentScopeUPICode returns true if InstrumentScopeUPICode is present, Tag 2895.
func (m SecurityMassStatusRequest) HasInstrumentScopeUPICode() bool {
	return m.Has(tag.InstrumentScopeUPICode)
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
