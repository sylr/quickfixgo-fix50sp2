package marketdatareport

import (
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// MarketDataReport is the fix50sp2 MarketDataReport type, MsgType = DR.
type MarketDataReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a MarketDataReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) MarketDataReport {
	return MarketDataReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m MarketDataReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a MarketDataReport initialized with the required fields for MarketDataReport.
func New(mdreportevent field.MDReportEventField, mdreportcount field.MDReportCountField) (m MarketDataReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("DR"))
	m.Set(mdreportevent)
	m.Set(mdreportcount)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg MarketDataReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "DR", r
}

// SetTransactTime sets TransactTime, Tag 60.
func (m MarketDataReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetTotNumReports sets TotNumReports, Tag 911.
func (m MarketDataReport) SetTotNumReports(v int) {
	m.Set(field.NewTotNumReports(v))
}

// SetMDReportID sets MDReportID, Tag 963.
func (m MarketDataReport) SetMDReportID(v int) {
	m.Set(field.NewMDReportID(v))
}

// SetApplID sets ApplID, Tag 1180.
func (m MarketDataReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181.
func (m MarketDataReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350.
func (m MarketDataReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352.
func (m MarketDataReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// SetMDReportEvent sets MDReportEvent, Tag 2535.
func (m MarketDataReport) SetMDReportEvent(v enum.MDReportEvent) {
	m.Set(field.NewMDReportEvent(v))
}

// SetMDReportCount sets MDReportCount, Tag 2536.
func (m MarketDataReport) SetMDReportCount(v int) {
	m.Set(field.NewMDReportCount(v))
}

// SetTotNoMarketSegmentReports sets TotNoMarketSegmentReports, Tag 2537.
func (m MarketDataReport) SetTotNoMarketSegmentReports(v int) {
	m.Set(field.NewTotNoMarketSegmentReports(v))
}

// SetTotNoInstrumentReports sets TotNoInstrumentReports, Tag 2538.
func (m MarketDataReport) SetTotNoInstrumentReports(v int) {
	m.Set(field.NewTotNoInstrumentReports(v))
}

// SetTotNoPartyDetailReports sets TotNoPartyDetailReports, Tag 2539.
func (m MarketDataReport) SetTotNoPartyDetailReports(v int) {
	m.Set(field.NewTotNoPartyDetailReports(v))
}

// SetTotNoEntitlementReports sets TotNoEntitlementReports, Tag 2540.
func (m MarketDataReport) SetTotNoEntitlementReports(v int) {
	m.Set(field.NewTotNoEntitlementReports(v))
}

// SetTotNoRiskLimitReports sets TotNoRiskLimitReports, Tag 2541.
func (m MarketDataReport) SetTotNoRiskLimitReports(v int) {
	m.Set(field.NewTotNoRiskLimitReports(v))
}

// GetTransactTime gets TransactTime, Tag 60.
func (m MarketDataReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNumReports gets TotNumReports, Tag 911.
func (m MarketDataReport) GetTotNumReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNumReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReportID gets MDReportID, Tag 963.
func (m MarketDataReport) GetMDReportID() (v int, err quickfix.MessageRejectError) {
	var f field.MDReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180.
func (m MarketDataReport) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181.
func (m MarketDataReport) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350.
func (m MarketDataReport) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352.
func (m MarketDataReport) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReportEvent gets MDReportEvent, Tag 2535.
func (m MarketDataReport) GetMDReportEvent() (v enum.MDReportEvent, err quickfix.MessageRejectError) {
	var f field.MDReportEventField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReportCount gets MDReportCount, Tag 2536.
func (m MarketDataReport) GetMDReportCount() (v int, err quickfix.MessageRejectError) {
	var f field.MDReportCountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoMarketSegmentReports gets TotNoMarketSegmentReports, Tag 2537.
func (m MarketDataReport) GetTotNoMarketSegmentReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoMarketSegmentReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoInstrumentReports gets TotNoInstrumentReports, Tag 2538.
func (m MarketDataReport) GetTotNoInstrumentReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoInstrumentReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoPartyDetailReports gets TotNoPartyDetailReports, Tag 2539.
func (m MarketDataReport) GetTotNoPartyDetailReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoPartyDetailReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoEntitlementReports gets TotNoEntitlementReports, Tag 2540.
func (m MarketDataReport) GetTotNoEntitlementReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoEntitlementReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotNoRiskLimitReports gets TotNoRiskLimitReports, Tag 2541.
func (m MarketDataReport) GetTotNoRiskLimitReports() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoRiskLimitReportsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m MarketDataReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasTotNumReports returns true if TotNumReports is present, Tag 911.
func (m MarketDataReport) HasTotNumReports() bool {
	return m.Has(tag.TotNumReports)
}

// HasMDReportID returns true if MDReportID is present, Tag 963.
func (m MarketDataReport) HasMDReportID() bool {
	return m.Has(tag.MDReportID)
}

// HasApplID returns true if ApplID is present, Tag 1180.
func (m MarketDataReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181.
func (m MarketDataReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350.
func (m MarketDataReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352.
func (m MarketDataReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// HasMDReportEvent returns true if MDReportEvent is present, Tag 2535.
func (m MarketDataReport) HasMDReportEvent() bool {
	return m.Has(tag.MDReportEvent)
}

// HasMDReportCount returns true if MDReportCount is present, Tag 2536.
func (m MarketDataReport) HasMDReportCount() bool {
	return m.Has(tag.MDReportCount)
}

// HasTotNoMarketSegmentReports returns true if TotNoMarketSegmentReports is present, Tag 2537.
func (m MarketDataReport) HasTotNoMarketSegmentReports() bool {
	return m.Has(tag.TotNoMarketSegmentReports)
}

// HasTotNoInstrumentReports returns true if TotNoInstrumentReports is present, Tag 2538.
func (m MarketDataReport) HasTotNoInstrumentReports() bool {
	return m.Has(tag.TotNoInstrumentReports)
}

// HasTotNoPartyDetailReports returns true if TotNoPartyDetailReports is present, Tag 2539.
func (m MarketDataReport) HasTotNoPartyDetailReports() bool {
	return m.Has(tag.TotNoPartyDetailReports)
}

// HasTotNoEntitlementReports returns true if TotNoEntitlementReports is present, Tag 2540.
func (m MarketDataReport) HasTotNoEntitlementReports() bool {
	return m.Has(tag.TotNoEntitlementReports)
}

// HasTotNoRiskLimitReports returns true if TotNoRiskLimitReports is present, Tag 2541.
func (m MarketDataReport) HasTotNoRiskLimitReports() bool {
	return m.Has(tag.TotNoRiskLimitReports)
}
