package confirmationack

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// ConfirmationAck is the fix50sp2 ConfirmationAck type, MsgType = AU.
type ConfirmationAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a ConfirmationAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) ConfirmationAck {
	return ConfirmationAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m ConfirmationAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a ConfirmationAck initialized with the required fields for ConfirmationAck.
func New(confirmid field.ConfirmIDField, tradedate field.TradeDateField, transacttime field.TransactTimeField, affirmstatus field.AffirmStatusField) (m ConfirmationAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AU"))
	m.Set(confirmid)
	m.Set(tradedate)
	m.Set(transacttime)
	m.Set(affirmstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg ConfirmationAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "AU", r
}

// SetText sets Text, Tag 58.
func (m ConfirmationAck) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m ConfirmationAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetTradeDate sets TradeDate, Tag 75.
func (m ConfirmationAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m ConfirmationAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m ConfirmationAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetMatchStatus sets MatchStatus, Tag 573.
func (m ConfirmationAck) SetMatchStatus(v enum.MatchStatus) {
	m.Set(field.NewMatchStatus(v))
}

// SetConfirmID sets ConfirmID, Tag 664.
func (m ConfirmationAck) SetConfirmID(v string) {
	m.Set(field.NewConfirmID(v))
}

// SetConfirmRejReason sets ConfirmRejReason, Tag 774.
func (m ConfirmationAck) SetConfirmRejReason(v enum.ConfirmRejReason) {
	m.Set(field.NewConfirmRejReason(v))
}

// SetAffirmStatus sets AffirmStatus, Tag 940.
func (m ConfirmationAck) SetAffirmStatus(v enum.AffirmStatus) {
	m.Set(field.NewAffirmStatus(v))
}

// SetTradeConfirmationReferenceID sets TradeConfirmationReferenceID, Tag 2390.
func (m ConfirmationAck) SetTradeConfirmationReferenceID(v string) {
	m.Set(field.NewTradeConfirmationReferenceID(v))
}

// SetNoMatchExceptions sets NoMatchExceptions, Tag 2772.
func (m ConfirmationAck) SetNoMatchExceptions(f NoMatchExceptionsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoMatchingDataPoints sets NoMatchingDataPoints, Tag 2781.
func (m ConfirmationAck) SetNoMatchingDataPoints(f NoMatchingDataPointsRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m ConfirmationAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m ConfirmationAck) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeDate gets TradeDate, Tag 75.
func (m ConfirmationAck) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m ConfirmationAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m ConfirmationAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchStatus gets MatchStatus, Tag 573.
func (m ConfirmationAck) GetMatchStatus() (v enum.MatchStatus, err quickfix.MessageRejectError) {
	var f field.MatchStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetConfirmID gets ConfirmID, Tag 664.
func (m ConfirmationAck) GetConfirmID() (v string, err quickfix.MessageRejectError) {
	var f field.ConfirmIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetConfirmRejReason gets ConfirmRejReason, Tag 774.
func (m ConfirmationAck) GetConfirmRejReason() (v enum.ConfirmRejReason, err quickfix.MessageRejectError) {
	var f field.ConfirmRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAffirmStatus gets AffirmStatus, Tag 940.
func (m ConfirmationAck) GetAffirmStatus() (v enum.AffirmStatus, err quickfix.MessageRejectError) {
	var f field.AffirmStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeConfirmationReferenceID gets TradeConfirmationReferenceID, Tag 2390.
func (m ConfirmationAck) GetTradeConfirmationReferenceID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeConfirmationReferenceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMatchExceptions gets NoMatchExceptions, Tag 2772.
func (m ConfirmationAck) GetNoMatchExceptions() (NoMatchExceptionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMatchExceptionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoMatchingDataPoints gets NoMatchingDataPoints, Tag 2781.
func (m ConfirmationAck) GetNoMatchingDataPoints() (NoMatchingDataPointsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMatchingDataPointsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m ConfirmationAck) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m ConfirmationAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasTradeDate returns true if TradeDate is present, Tag 75.
func (m ConfirmationAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m ConfirmationAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m ConfirmationAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasMatchStatus returns true if MatchStatus is present, Tag 573.
func (m ConfirmationAck) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

// HasConfirmID returns true if ConfirmID is present, Tag 664.
func (m ConfirmationAck) HasConfirmID() bool {
	return m.Has(tag.ConfirmID)
}

// HasConfirmRejReason returns true if ConfirmRejReason is present, Tag 774.
func (m ConfirmationAck) HasConfirmRejReason() bool {
	return m.Has(tag.ConfirmRejReason)
}

// HasAffirmStatus returns true if AffirmStatus is present, Tag 940.
func (m ConfirmationAck) HasAffirmStatus() bool {
	return m.Has(tag.AffirmStatus)
}

// HasTradeConfirmationReferenceID returns true if TradeConfirmationReferenceID is present, Tag 2390.
func (m ConfirmationAck) HasTradeConfirmationReferenceID() bool {
	return m.Has(tag.TradeConfirmationReferenceID)
}

// HasNoMatchExceptions returns true if NoMatchExceptions is present, Tag 2772.
func (m ConfirmationAck) HasNoMatchExceptions() bool {
	return m.Has(tag.NoMatchExceptions)
}

// HasNoMatchingDataPoints returns true if NoMatchingDataPoints is present, Tag 2781.
func (m ConfirmationAck) HasNoMatchingDataPoints() bool {
	return m.Has(tag.NoMatchingDataPoints)
}

// NoMatchExceptions is a repeating group element, Tag 2772.
type NoMatchExceptions struct {
	*quickfix.Group
}

// SetMatchExceptionType sets MatchExceptionType, Tag 2773.
func (m NoMatchExceptions) SetMatchExceptionType(v enum.MatchExceptionType) {
	m.Set(field.NewMatchExceptionType(v))
}

// SetMatchExceptionElementType sets MatchExceptionElementType, Tag 2774.
func (m NoMatchExceptions) SetMatchExceptionElementType(v enum.MatchExceptionElementType) {
	m.Set(field.NewMatchExceptionElementType(v))
}

// SetMatchExceptionElementName sets MatchExceptionElementName, Tag 2775.
func (m NoMatchExceptions) SetMatchExceptionElementName(v string) {
	m.Set(field.NewMatchExceptionElementName(v))
}

// SetMatchExceptionAllocValue sets MatchExceptionAllocValue, Tag 2776.
func (m NoMatchExceptions) SetMatchExceptionAllocValue(v string) {
	m.Set(field.NewMatchExceptionAllocValue(v))
}

// SetMatchExceptionConfirmValue sets MatchExceptionConfirmValue, Tag 2777.
func (m NoMatchExceptions) SetMatchExceptionConfirmValue(v string) {
	m.Set(field.NewMatchExceptionConfirmValue(v))
}

// SetMatchExceptionToleranceValue sets MatchExceptionToleranceValue, Tag 2778.
func (m NoMatchExceptions) SetMatchExceptionToleranceValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewMatchExceptionToleranceValue(value, scale))
}

// SetMatchExceptionToleranceValueType sets MatchExceptionToleranceValueType, Tag 2779.
func (m NoMatchExceptions) SetMatchExceptionToleranceValueType(v enum.MatchExceptionToleranceValueType) {
	m.Set(field.NewMatchExceptionToleranceValueType(v))
}

// SetMatchExceptionText sets MatchExceptionText, Tag 2780.
func (m NoMatchExceptions) SetMatchExceptionText(v string) {
	m.Set(field.NewMatchExceptionText(v))
}

// SetEncodedMatchExceptionTextLen sets EncodedMatchExceptionTextLen, Tag 2797.
func (m NoMatchExceptions) SetEncodedMatchExceptionTextLen(v int) {
	m.Set(field.NewEncodedMatchExceptionTextLen(v))
}

// SetEncodedMatchExecptionText sets EncodedMatchExecptionText, Tag 2798.
func (m NoMatchExceptions) SetEncodedMatchExecptionText(v string) {
	m.Set(field.NewEncodedMatchExecptionText(v))
}

// GetMatchExceptionType gets MatchExceptionType, Tag 2773.
func (m NoMatchExceptions) GetMatchExceptionType() (v enum.MatchExceptionType, err quickfix.MessageRejectError) {
	var f field.MatchExceptionTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchExceptionElementType gets MatchExceptionElementType, Tag 2774.
func (m NoMatchExceptions) GetMatchExceptionElementType() (v enum.MatchExceptionElementType, err quickfix.MessageRejectError) {
	var f field.MatchExceptionElementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchExceptionElementName gets MatchExceptionElementName, Tag 2775.
func (m NoMatchExceptions) GetMatchExceptionElementName() (v string, err quickfix.MessageRejectError) {
	var f field.MatchExceptionElementNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchExceptionAllocValue gets MatchExceptionAllocValue, Tag 2776.
func (m NoMatchExceptions) GetMatchExceptionAllocValue() (v string, err quickfix.MessageRejectError) {
	var f field.MatchExceptionAllocValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchExceptionConfirmValue gets MatchExceptionConfirmValue, Tag 2777.
func (m NoMatchExceptions) GetMatchExceptionConfirmValue() (v string, err quickfix.MessageRejectError) {
	var f field.MatchExceptionConfirmValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchExceptionToleranceValue gets MatchExceptionToleranceValue, Tag 2778.
func (m NoMatchExceptions) GetMatchExceptionToleranceValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MatchExceptionToleranceValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchExceptionToleranceValueType gets MatchExceptionToleranceValueType, Tag 2779.
func (m NoMatchExceptions) GetMatchExceptionToleranceValueType() (v enum.MatchExceptionToleranceValueType, err quickfix.MessageRejectError) {
	var f field.MatchExceptionToleranceValueTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchExceptionText gets MatchExceptionText, Tag 2780.
func (m NoMatchExceptions) GetMatchExceptionText() (v string, err quickfix.MessageRejectError) {
	var f field.MatchExceptionTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedMatchExceptionTextLen gets EncodedMatchExceptionTextLen, Tag 2797.
func (m NoMatchExceptions) GetEncodedMatchExceptionTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedMatchExceptionTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedMatchExecptionText gets EncodedMatchExecptionText, Tag 2798.
func (m NoMatchExceptions) GetEncodedMatchExecptionText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedMatchExecptionTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMatchExceptionType returns true if MatchExceptionType is present, Tag 2773.
func (m NoMatchExceptions) HasMatchExceptionType() bool {
	return m.Has(tag.MatchExceptionType)
}

// HasMatchExceptionElementType returns true if MatchExceptionElementType is present, Tag 2774.
func (m NoMatchExceptions) HasMatchExceptionElementType() bool {
	return m.Has(tag.MatchExceptionElementType)
}

// HasMatchExceptionElementName returns true if MatchExceptionElementName is present, Tag 2775.
func (m NoMatchExceptions) HasMatchExceptionElementName() bool {
	return m.Has(tag.MatchExceptionElementName)
}

// HasMatchExceptionAllocValue returns true if MatchExceptionAllocValue is present, Tag 2776.
func (m NoMatchExceptions) HasMatchExceptionAllocValue() bool {
	return m.Has(tag.MatchExceptionAllocValue)
}

// HasMatchExceptionConfirmValue returns true if MatchExceptionConfirmValue is present, Tag 2777.
func (m NoMatchExceptions) HasMatchExceptionConfirmValue() bool {
	return m.Has(tag.MatchExceptionConfirmValue)
}

// HasMatchExceptionToleranceValue returns true if MatchExceptionToleranceValue is present, Tag 2778.
func (m NoMatchExceptions) HasMatchExceptionToleranceValue() bool {
	return m.Has(tag.MatchExceptionToleranceValue)
}

// HasMatchExceptionToleranceValueType returns true if MatchExceptionToleranceValueType is present, Tag 2779.
func (m NoMatchExceptions) HasMatchExceptionToleranceValueType() bool {
	return m.Has(tag.MatchExceptionToleranceValueType)
}

// HasMatchExceptionText returns true if MatchExceptionText is present, Tag 2780.
func (m NoMatchExceptions) HasMatchExceptionText() bool {
	return m.Has(tag.MatchExceptionText)
}

// HasEncodedMatchExceptionTextLen returns true if EncodedMatchExceptionTextLen is present, Tag 2797.
func (m NoMatchExceptions) HasEncodedMatchExceptionTextLen() bool {
	return m.Has(tag.EncodedMatchExceptionTextLen)
}

// HasEncodedMatchExecptionText returns true if EncodedMatchExecptionText is present, Tag 2798.
func (m NoMatchExceptions) HasEncodedMatchExecptionText() bool {
	return m.Has(tag.EncodedMatchExecptionText)
}

// NoMatchExceptionsRepeatingGroup is a repeating group, Tag 2772.
type NoMatchExceptionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMatchExceptionsRepeatingGroup returns an initialized, NoMatchExceptionsRepeatingGroup.
func NewNoMatchExceptionsRepeatingGroup() NoMatchExceptionsRepeatingGroup {
	return NoMatchExceptionsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMatchExceptions,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MatchExceptionType),
				quickfix.GroupElement(tag.MatchExceptionElementType),
				quickfix.GroupElement(tag.MatchExceptionElementName),
				quickfix.GroupElement(tag.MatchExceptionAllocValue),
				quickfix.GroupElement(tag.MatchExceptionConfirmValue),
				quickfix.GroupElement(tag.MatchExceptionToleranceValue),
				quickfix.GroupElement(tag.MatchExceptionToleranceValueType),
				quickfix.GroupElement(tag.MatchExceptionText),
				quickfix.GroupElement(tag.EncodedMatchExceptionTextLen),
				quickfix.GroupElement(tag.EncodedMatchExecptionText),
			},
		),
	}
}

// Add create and append a new NoMatchExceptions to this group.
func (m NoMatchExceptionsRepeatingGroup) Add() NoMatchExceptions {
	g := m.RepeatingGroup.Add()
	return NoMatchExceptions{g}
}

// Get returns the ith NoMatchExceptions in the NoMatchExceptionsRepeatinGroup.
func (m NoMatchExceptionsRepeatingGroup) Get(i int) NoMatchExceptions {
	return NoMatchExceptions{m.RepeatingGroup.Get(i)}
}

// NoMatchingDataPoints is a repeating group element, Tag 2781.
type NoMatchingDataPoints struct {
	*quickfix.Group
}

// SetMatchingDataPointIndicator sets MatchingDataPointIndicator, Tag 2782.
func (m NoMatchingDataPoints) SetMatchingDataPointIndicator(v enum.MatchingDataPointIndicator) {
	m.Set(field.NewMatchingDataPointIndicator(v))
}

// SetMatchingDataPointValue sets MatchingDataPointValue, Tag 2783.
func (m NoMatchingDataPoints) SetMatchingDataPointValue(v string) {
	m.Set(field.NewMatchingDataPointValue(v))
}

// SetMatchingDataPointType sets MatchingDataPointType, Tag 2784.
func (m NoMatchingDataPoints) SetMatchingDataPointType(v int) {
	m.Set(field.NewMatchingDataPointType(v))
}

// SetMatchingDataPointName sets MatchingDataPointName, Tag 2785.
func (m NoMatchingDataPoints) SetMatchingDataPointName(v string) {
	m.Set(field.NewMatchingDataPointName(v))
}

// GetMatchingDataPointIndicator gets MatchingDataPointIndicator, Tag 2782.
func (m NoMatchingDataPoints) GetMatchingDataPointIndicator() (v enum.MatchingDataPointIndicator, err quickfix.MessageRejectError) {
	var f field.MatchingDataPointIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchingDataPointValue gets MatchingDataPointValue, Tag 2783.
func (m NoMatchingDataPoints) GetMatchingDataPointValue() (v string, err quickfix.MessageRejectError) {
	var f field.MatchingDataPointValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchingDataPointType gets MatchingDataPointType, Tag 2784.
func (m NoMatchingDataPoints) GetMatchingDataPointType() (v int, err quickfix.MessageRejectError) {
	var f field.MatchingDataPointTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchingDataPointName gets MatchingDataPointName, Tag 2785.
func (m NoMatchingDataPoints) GetMatchingDataPointName() (v string, err quickfix.MessageRejectError) {
	var f field.MatchingDataPointNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMatchingDataPointIndicator returns true if MatchingDataPointIndicator is present, Tag 2782.
func (m NoMatchingDataPoints) HasMatchingDataPointIndicator() bool {
	return m.Has(tag.MatchingDataPointIndicator)
}

// HasMatchingDataPointValue returns true if MatchingDataPointValue is present, Tag 2783.
func (m NoMatchingDataPoints) HasMatchingDataPointValue() bool {
	return m.Has(tag.MatchingDataPointValue)
}

// HasMatchingDataPointType returns true if MatchingDataPointType is present, Tag 2784.
func (m NoMatchingDataPoints) HasMatchingDataPointType() bool {
	return m.Has(tag.MatchingDataPointType)
}

// HasMatchingDataPointName returns true if MatchingDataPointName is present, Tag 2785.
func (m NoMatchingDataPoints) HasMatchingDataPointName() bool {
	return m.Has(tag.MatchingDataPointName)
}

// NoMatchingDataPointsRepeatingGroup is a repeating group, Tag 2781.
type NoMatchingDataPointsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMatchingDataPointsRepeatingGroup returns an initialized, NoMatchingDataPointsRepeatingGroup.
func NewNoMatchingDataPointsRepeatingGroup() NoMatchingDataPointsRepeatingGroup {
	return NoMatchingDataPointsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMatchingDataPoints,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MatchingDataPointIndicator),
				quickfix.GroupElement(tag.MatchingDataPointValue),
				quickfix.GroupElement(tag.MatchingDataPointType),
				quickfix.GroupElement(tag.MatchingDataPointName),
			},
		),
	}
}

// Add create and append a new NoMatchingDataPoints to this group.
func (m NoMatchingDataPointsRepeatingGroup) Add() NoMatchingDataPoints {
	g := m.RepeatingGroup.Add()
	return NoMatchingDataPoints{g}
}

// Get returns the ith NoMatchingDataPoints in the NoMatchingDataPointsRepeatinGroup.
func (m NoMatchingDataPointsRepeatingGroup) Get(i int) NoMatchingDataPoints {
	return NoMatchingDataPoints{m.RepeatingGroup.Get(i)}
}
