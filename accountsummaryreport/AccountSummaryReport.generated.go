package accountsummaryreport

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// AccountSummaryReport is the fix50sp2 AccountSummaryReport type, MsgType = CQ.
type AccountSummaryReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a AccountSummaryReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) AccountSummaryReport {
	return AccountSummaryReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m AccountSummaryReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a AccountSummaryReport initialized with the required fields for AccountSummaryReport.
func New(accountsummaryreportid field.AccountSummaryReportIDField, clearingbusinessdate field.ClearingBusinessDateField) (m AccountSummaryReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CQ"))
	m.Set(accountsummaryreportid)
	m.Set(clearingbusinessdate)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg AccountSummaryReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CQ", r
}

// SetCurrency sets Currency, Tag 15.
func (m AccountSummaryReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m AccountSummaryReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m AccountSummaryReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetClearingBusinessDate sets ClearingBusinessDate, Tag 715.
func (m AccountSummaryReport) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

// SetSettlSessID sets SettlSessID, Tag 716.
func (m AccountSummaryReport) SetSettlSessID(v enum.SettlSessID) {
	m.Set(field.NewSettlSessID(v))
}

// SetSettlSessSubID sets SettlSessSubID, Tag 717.
func (m AccountSummaryReport) SetSettlSessSubID(v string) {
	m.Set(field.NewSettlSessSubID(v))
}

// SetNoPosAmt sets NoPosAmt, Tag 753.
func (m AccountSummaryReport) SetNoPosAmt(f NoPosAmtRepeatingGroup) {
	m.SetGroup(f)
}

// SetMarginExcess sets MarginExcess, Tag 899.
func (m AccountSummaryReport) SetMarginExcess(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginExcess(value, scale))
}

// SetTotalNetValue sets TotalNetValue, Tag 900.
func (m AccountSummaryReport) SetTotalNetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalNetValue(value, scale))
}

// SetNoMarginAmt sets NoMarginAmt, Tag 1643.
func (m AccountSummaryReport) SetNoMarginAmt(f NoMarginAmtRepeatingGroup) {
	m.SetGroup(f)
}

// SetAccountSummaryReportID sets AccountSummaryReportID, Tag 1699.
func (m AccountSummaryReport) SetAccountSummaryReportID(v string) {
	m.Set(field.NewAccountSummaryReportID(v))
}

// SetNoSettlementAmounts sets NoSettlementAmounts, Tag 1700.
func (m AccountSummaryReport) SetNoSettlementAmounts(f NoSettlementAmountsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoCollateralAmounts sets NoCollateralAmounts, Tag 1703.
func (m AccountSummaryReport) SetNoCollateralAmounts(f NoCollateralAmountsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoPayCollects sets NoPayCollects, Tag 1707.
func (m AccountSummaryReport) SetNoPayCollects(f NoPayCollectsRepeatingGroup) {
	m.SetGroup(f)
}

// GetCurrency gets Currency, Tag 15.
func (m AccountSummaryReport) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m AccountSummaryReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m AccountSummaryReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetClearingBusinessDate gets ClearingBusinessDate, Tag 715.
func (m AccountSummaryReport) GetClearingBusinessDate() (v string, err quickfix.MessageRejectError) {
	var f field.ClearingBusinessDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlSessID gets SettlSessID, Tag 716.
func (m AccountSummaryReport) GetSettlSessID() (v enum.SettlSessID, err quickfix.MessageRejectError) {
	var f field.SettlSessIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlSessSubID gets SettlSessSubID, Tag 717.
func (m AccountSummaryReport) GetSettlSessSubID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlSessSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPosAmt gets NoPosAmt, Tag 753.
func (m AccountSummaryReport) GetNoPosAmt() (NoPosAmtRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPosAmtRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMarginExcess gets MarginExcess, Tag 899.
func (m AccountSummaryReport) GetMarginExcess() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginExcessField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotalNetValue gets TotalNetValue, Tag 900.
func (m AccountSummaryReport) GetTotalNetValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalNetValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMarginAmt gets NoMarginAmt, Tag 1643.
func (m AccountSummaryReport) GetNoMarginAmt() (NoMarginAmtRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMarginAmtRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetAccountSummaryReportID gets AccountSummaryReportID, Tag 1699.
func (m AccountSummaryReport) GetAccountSummaryReportID() (v string, err quickfix.MessageRejectError) {
	var f field.AccountSummaryReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoSettlementAmounts gets NoSettlementAmounts, Tag 1700.
func (m AccountSummaryReport) GetNoSettlementAmounts() (NoSettlementAmountsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlementAmountsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoCollateralAmounts gets NoCollateralAmounts, Tag 1703.
func (m AccountSummaryReport) GetNoCollateralAmounts() (NoCollateralAmountsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoCollateralAmountsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoPayCollects gets NoPayCollects, Tag 1707.
func (m AccountSummaryReport) GetNoPayCollects() (NoPayCollectsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPayCollectsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m AccountSummaryReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m AccountSummaryReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m AccountSummaryReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715.
func (m AccountSummaryReport) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

// HasSettlSessID returns true if SettlSessID is present, Tag 716.
func (m AccountSummaryReport) HasSettlSessID() bool {
	return m.Has(tag.SettlSessID)
}

// HasSettlSessSubID returns true if SettlSessSubID is present, Tag 717.
func (m AccountSummaryReport) HasSettlSessSubID() bool {
	return m.Has(tag.SettlSessSubID)
}

// HasNoPosAmt returns true if NoPosAmt is present, Tag 753.
func (m AccountSummaryReport) HasNoPosAmt() bool {
	return m.Has(tag.NoPosAmt)
}

// HasMarginExcess returns true if MarginExcess is present, Tag 899.
func (m AccountSummaryReport) HasMarginExcess() bool {
	return m.Has(tag.MarginExcess)
}

// HasTotalNetValue returns true if TotalNetValue is present, Tag 900.
func (m AccountSummaryReport) HasTotalNetValue() bool {
	return m.Has(tag.TotalNetValue)
}

// HasNoMarginAmt returns true if NoMarginAmt is present, Tag 1643.
func (m AccountSummaryReport) HasNoMarginAmt() bool {
	return m.Has(tag.NoMarginAmt)
}

// HasAccountSummaryReportID returns true if AccountSummaryReportID is present, Tag 1699.
func (m AccountSummaryReport) HasAccountSummaryReportID() bool {
	return m.Has(tag.AccountSummaryReportID)
}

// HasNoSettlementAmounts returns true if NoSettlementAmounts is present, Tag 1700.
func (m AccountSummaryReport) HasNoSettlementAmounts() bool {
	return m.Has(tag.NoSettlementAmounts)
}

// HasNoCollateralAmounts returns true if NoCollateralAmounts is present, Tag 1703.
func (m AccountSummaryReport) HasNoCollateralAmounts() bool {
	return m.Has(tag.NoCollateralAmounts)
}

// HasNoPayCollects returns true if NoPayCollects is present, Tag 1707.
func (m AccountSummaryReport) HasNoPayCollects() bool {
	return m.Has(tag.NoPayCollects)
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

// NoPosAmt is a repeating group element, Tag 753.
type NoPosAmt struct {
	*quickfix.Group
}

// SetPosAmtType sets PosAmtType, Tag 707.
func (m NoPosAmt) SetPosAmtType(v enum.PosAmtType) {
	m.Set(field.NewPosAmtType(v))
}

// SetPosAmt sets PosAmt, Tag 708.
func (m NoPosAmt) SetPosAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewPosAmt(value, scale))
}

// SetPositionCurrency sets PositionCurrency, Tag 1055.
func (m NoPosAmt) SetPositionCurrency(v string) {
	m.Set(field.NewPositionCurrency(v))
}

// SetPosAmtReason sets PosAmtReason, Tag 1585.
func (m NoPosAmt) SetPosAmtReason(v enum.PosAmtReason) {
	m.Set(field.NewPosAmtReason(v))
}

// SetPosAmtStreamDesc sets PosAmtStreamDesc, Tag 2096.
func (m NoPosAmt) SetPosAmtStreamDesc(v string) {
	m.Set(field.NewPosAmtStreamDesc(v))
}

// SetPositionFXRate sets PositionFXRate, Tag 2097.
func (m NoPosAmt) SetPositionFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewPositionFXRate(value, scale))
}

// SetPositionFXRateCalc sets PositionFXRateCalc, Tag 2098.
func (m NoPosAmt) SetPositionFXRateCalc(v string) {
	m.Set(field.NewPositionFXRateCalc(v))
}

// SetPosAmtMarketSegmentID sets PosAmtMarketSegmentID, Tag 2099.
func (m NoPosAmt) SetPosAmtMarketSegmentID(v string) {
	m.Set(field.NewPosAmtMarketSegmentID(v))
}

// SetPosAmtMarketID sets PosAmtMarketID, Tag 2100.
func (m NoPosAmt) SetPosAmtMarketID(v string) {
	m.Set(field.NewPosAmtMarketID(v))
}

// SetPosAmtPrice sets PosAmtPrice, Tag 2876.
func (m NoPosAmt) SetPosAmtPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPosAmtPrice(value, scale))
}

// SetPosAmtPriceType sets PosAmtPriceType, Tag 2877.
func (m NoPosAmt) SetPosAmtPriceType(v int) {
	m.Set(field.NewPosAmtPriceType(v))
}

// GetPosAmtType gets PosAmtType, Tag 707.
func (m NoPosAmt) GetPosAmtType() (v enum.PosAmtType, err quickfix.MessageRejectError) {
	var f field.PosAmtTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPosAmt gets PosAmt, Tag 708.
func (m NoPosAmt) GetPosAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PosAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionCurrency gets PositionCurrency, Tag 1055.
func (m NoPosAmt) GetPositionCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.PositionCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPosAmtReason gets PosAmtReason, Tag 1585.
func (m NoPosAmt) GetPosAmtReason() (v enum.PosAmtReason, err quickfix.MessageRejectError) {
	var f field.PosAmtReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPosAmtStreamDesc gets PosAmtStreamDesc, Tag 2096.
func (m NoPosAmt) GetPosAmtStreamDesc() (v string, err quickfix.MessageRejectError) {
	var f field.PosAmtStreamDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionFXRate gets PositionFXRate, Tag 2097.
func (m NoPosAmt) GetPositionFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PositionFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionFXRateCalc gets PositionFXRateCalc, Tag 2098.
func (m NoPosAmt) GetPositionFXRateCalc() (v string, err quickfix.MessageRejectError) {
	var f field.PositionFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPosAmtMarketSegmentID gets PosAmtMarketSegmentID, Tag 2099.
func (m NoPosAmt) GetPosAmtMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.PosAmtMarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPosAmtMarketID gets PosAmtMarketID, Tag 2100.
func (m NoPosAmt) GetPosAmtMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.PosAmtMarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPosAmtPrice gets PosAmtPrice, Tag 2876.
func (m NoPosAmt) GetPosAmtPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PosAmtPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPosAmtPriceType gets PosAmtPriceType, Tag 2877.
func (m NoPosAmt) GetPosAmtPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.PosAmtPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPosAmtType returns true if PosAmtType is present, Tag 707.
func (m NoPosAmt) HasPosAmtType() bool {
	return m.Has(tag.PosAmtType)
}

// HasPosAmt returns true if PosAmt is present, Tag 708.
func (m NoPosAmt) HasPosAmt() bool {
	return m.Has(tag.PosAmt)
}

// HasPositionCurrency returns true if PositionCurrency is present, Tag 1055.
func (m NoPosAmt) HasPositionCurrency() bool {
	return m.Has(tag.PositionCurrency)
}

// HasPosAmtReason returns true if PosAmtReason is present, Tag 1585.
func (m NoPosAmt) HasPosAmtReason() bool {
	return m.Has(tag.PosAmtReason)
}

// HasPosAmtStreamDesc returns true if PosAmtStreamDesc is present, Tag 2096.
func (m NoPosAmt) HasPosAmtStreamDesc() bool {
	return m.Has(tag.PosAmtStreamDesc)
}

// HasPositionFXRate returns true if PositionFXRate is present, Tag 2097.
func (m NoPosAmt) HasPositionFXRate() bool {
	return m.Has(tag.PositionFXRate)
}

// HasPositionFXRateCalc returns true if PositionFXRateCalc is present, Tag 2098.
func (m NoPosAmt) HasPositionFXRateCalc() bool {
	return m.Has(tag.PositionFXRateCalc)
}

// HasPosAmtMarketSegmentID returns true if PosAmtMarketSegmentID is present, Tag 2099.
func (m NoPosAmt) HasPosAmtMarketSegmentID() bool {
	return m.Has(tag.PosAmtMarketSegmentID)
}

// HasPosAmtMarketID returns true if PosAmtMarketID is present, Tag 2100.
func (m NoPosAmt) HasPosAmtMarketID() bool {
	return m.Has(tag.PosAmtMarketID)
}

// HasPosAmtPrice returns true if PosAmtPrice is present, Tag 2876.
func (m NoPosAmt) HasPosAmtPrice() bool {
	return m.Has(tag.PosAmtPrice)
}

// HasPosAmtPriceType returns true if PosAmtPriceType is present, Tag 2877.
func (m NoPosAmt) HasPosAmtPriceType() bool {
	return m.Has(tag.PosAmtPriceType)
}

// NoPosAmtRepeatingGroup is a repeating group, Tag 753.
type NoPosAmtRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPosAmtRepeatingGroup returns an initialized, NoPosAmtRepeatingGroup.
func NewNoPosAmtRepeatingGroup() NoPosAmtRepeatingGroup {
	return NoPosAmtRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPosAmt,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PosAmtType),
				quickfix.GroupElement(tag.PosAmt),
				quickfix.GroupElement(tag.PositionCurrency),
				quickfix.GroupElement(tag.PosAmtReason),
				quickfix.GroupElement(tag.PosAmtStreamDesc),
				quickfix.GroupElement(tag.PositionFXRate),
				quickfix.GroupElement(tag.PositionFXRateCalc),
				quickfix.GroupElement(tag.PosAmtMarketSegmentID),
				quickfix.GroupElement(tag.PosAmtMarketID),
				quickfix.GroupElement(tag.PosAmtPrice),
				quickfix.GroupElement(tag.PosAmtPriceType),
			},
		),
	}
}

// Add create and append a new NoPosAmt to this group.
func (m NoPosAmtRepeatingGroup) Add() NoPosAmt {
	g := m.RepeatingGroup.Add()
	return NoPosAmt{g}
}

// Get returns the ith NoPosAmt in the NoPosAmtRepeatinGroup.
func (m NoPosAmtRepeatingGroup) Get(i int) NoPosAmt {
	return NoPosAmt{m.RepeatingGroup.Get(i)}
}

// NoMarginAmt is a repeating group element, Tag 1643.
type NoMarginAmt struct {
	*quickfix.Group
}

// SetMarginAmt sets MarginAmt, Tag 1645.
func (m NoMarginAmt) SetMarginAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginAmt(value, scale))
}

// SetMarginAmtType sets MarginAmtType, Tag 1644.
func (m NoMarginAmt) SetMarginAmtType(v enum.MarginAmtType) {
	m.Set(field.NewMarginAmtType(v))
}

// SetMarginAmtCcy sets MarginAmtCcy, Tag 1646.
func (m NoMarginAmt) SetMarginAmtCcy(v string) {
	m.Set(field.NewMarginAmtCcy(v))
}

// SetMarginAmountMarketSegmentID sets MarginAmountMarketSegmentID, Tag 1714.
func (m NoMarginAmt) SetMarginAmountMarketSegmentID(v string) {
	m.Set(field.NewMarginAmountMarketSegmentID(v))
}

// SetMarginAmountMarketID sets MarginAmountMarketID, Tag 1715.
func (m NoMarginAmt) SetMarginAmountMarketID(v string) {
	m.Set(field.NewMarginAmountMarketID(v))
}

// SetMarginAmtFXRate sets MarginAmtFXRate, Tag 2088.
func (m NoMarginAmt) SetMarginAmtFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginAmtFXRate(value, scale))
}

// SetMarginAmtFXRateCalc sets MarginAmtFXRateCalc, Tag 2089.
func (m NoMarginAmt) SetMarginAmtFXRateCalc(v string) {
	m.Set(field.NewMarginAmtFXRateCalc(v))
}

// SetMarginDirection sets MarginDirection, Tag 2851.
func (m NoMarginAmt) SetMarginDirection(v enum.MarginDirection) {
	m.Set(field.NewMarginDirection(v))
}

// GetMarginAmt gets MarginAmt, Tag 1645.
func (m NoMarginAmt) GetMarginAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginAmtType gets MarginAmtType, Tag 1644.
func (m NoMarginAmt) GetMarginAmtType() (v enum.MarginAmtType, err quickfix.MessageRejectError) {
	var f field.MarginAmtTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginAmtCcy gets MarginAmtCcy, Tag 1646.
func (m NoMarginAmt) GetMarginAmtCcy() (v string, err quickfix.MessageRejectError) {
	var f field.MarginAmtCcyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginAmountMarketSegmentID gets MarginAmountMarketSegmentID, Tag 1714.
func (m NoMarginAmt) GetMarginAmountMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarginAmountMarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginAmountMarketID gets MarginAmountMarketID, Tag 1715.
func (m NoMarginAmt) GetMarginAmountMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarginAmountMarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginAmtFXRate gets MarginAmtFXRate, Tag 2088.
func (m NoMarginAmt) GetMarginAmtFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginAmtFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginAmtFXRateCalc gets MarginAmtFXRateCalc, Tag 2089.
func (m NoMarginAmt) GetMarginAmtFXRateCalc() (v string, err quickfix.MessageRejectError) {
	var f field.MarginAmtFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginDirection gets MarginDirection, Tag 2851.
func (m NoMarginAmt) GetMarginDirection() (v enum.MarginDirection, err quickfix.MessageRejectError) {
	var f field.MarginDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMarginAmt returns true if MarginAmt is present, Tag 1645.
func (m NoMarginAmt) HasMarginAmt() bool {
	return m.Has(tag.MarginAmt)
}

// HasMarginAmtType returns true if MarginAmtType is present, Tag 1644.
func (m NoMarginAmt) HasMarginAmtType() bool {
	return m.Has(tag.MarginAmtType)
}

// HasMarginAmtCcy returns true if MarginAmtCcy is present, Tag 1646.
func (m NoMarginAmt) HasMarginAmtCcy() bool {
	return m.Has(tag.MarginAmtCcy)
}

// HasMarginAmountMarketSegmentID returns true if MarginAmountMarketSegmentID is present, Tag 1714.
func (m NoMarginAmt) HasMarginAmountMarketSegmentID() bool {
	return m.Has(tag.MarginAmountMarketSegmentID)
}

// HasMarginAmountMarketID returns true if MarginAmountMarketID is present, Tag 1715.
func (m NoMarginAmt) HasMarginAmountMarketID() bool {
	return m.Has(tag.MarginAmountMarketID)
}

// HasMarginAmtFXRate returns true if MarginAmtFXRate is present, Tag 2088.
func (m NoMarginAmt) HasMarginAmtFXRate() bool {
	return m.Has(tag.MarginAmtFXRate)
}

// HasMarginAmtFXRateCalc returns true if MarginAmtFXRateCalc is present, Tag 2089.
func (m NoMarginAmt) HasMarginAmtFXRateCalc() bool {
	return m.Has(tag.MarginAmtFXRateCalc)
}

// HasMarginDirection returns true if MarginDirection is present, Tag 2851.
func (m NoMarginAmt) HasMarginDirection() bool {
	return m.Has(tag.MarginDirection)
}

// NoMarginAmtRepeatingGroup is a repeating group, Tag 1643.
type NoMarginAmtRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMarginAmtRepeatingGroup returns an initialized, NoMarginAmtRepeatingGroup.
func NewNoMarginAmtRepeatingGroup() NoMarginAmtRepeatingGroup {
	return NoMarginAmtRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMarginAmt,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MarginAmt),
				quickfix.GroupElement(tag.MarginAmtType),
				quickfix.GroupElement(tag.MarginAmtCcy),
				quickfix.GroupElement(tag.MarginAmountMarketSegmentID),
				quickfix.GroupElement(tag.MarginAmountMarketID),
				quickfix.GroupElement(tag.MarginAmtFXRate),
				quickfix.GroupElement(tag.MarginAmtFXRateCalc),
				quickfix.GroupElement(tag.MarginDirection),
			},
		),
	}
}

// Add create and append a new NoMarginAmt to this group.
func (m NoMarginAmtRepeatingGroup) Add() NoMarginAmt {
	g := m.RepeatingGroup.Add()
	return NoMarginAmt{g}
}

// Get returns the ith NoMarginAmt in the NoMarginAmtRepeatinGroup.
func (m NoMarginAmtRepeatingGroup) Get(i int) NoMarginAmt {
	return NoMarginAmt{m.RepeatingGroup.Get(i)}
}

// NoSettlementAmounts is a repeating group element, Tag 1700.
type NoSettlementAmounts struct {
	*quickfix.Group
}

// SetSettlementAmount sets SettlementAmount, Tag 1701.
func (m NoSettlementAmounts) SetSettlementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlementAmount(value, scale))
}

// SetSettlementAmountCurrency sets SettlementAmountCurrency, Tag 1702.
func (m NoSettlementAmounts) SetSettlementAmountCurrency(v string) {
	m.Set(field.NewSettlementAmountCurrency(v))
}

// GetSettlementAmount gets SettlementAmount, Tag 1701.
func (m NoSettlementAmounts) GetSettlementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlementAmountCurrency gets SettlementAmountCurrency, Tag 1702.
func (m NoSettlementAmounts) GetSettlementAmountCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlementAmountCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSettlementAmount returns true if SettlementAmount is present, Tag 1701.
func (m NoSettlementAmounts) HasSettlementAmount() bool {
	return m.Has(tag.SettlementAmount)
}

// HasSettlementAmountCurrency returns true if SettlementAmountCurrency is present, Tag 1702.
func (m NoSettlementAmounts) HasSettlementAmountCurrency() bool {
	return m.Has(tag.SettlementAmountCurrency)
}

// NoSettlementAmountsRepeatingGroup is a repeating group, Tag 1700.
type NoSettlementAmountsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoSettlementAmountsRepeatingGroup returns an initialized, NoSettlementAmountsRepeatingGroup.
func NewNoSettlementAmountsRepeatingGroup() NoSettlementAmountsRepeatingGroup {
	return NoSettlementAmountsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoSettlementAmounts,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.SettlementAmount),
				quickfix.GroupElement(tag.SettlementAmountCurrency),
			},
		),
	}
}

// Add create and append a new NoSettlementAmounts to this group.
func (m NoSettlementAmountsRepeatingGroup) Add() NoSettlementAmounts {
	g := m.RepeatingGroup.Add()
	return NoSettlementAmounts{g}
}

// Get returns the ith NoSettlementAmounts in the NoSettlementAmountsRepeatinGroup.
func (m NoSettlementAmountsRepeatingGroup) Get(i int) NoSettlementAmounts {
	return NoSettlementAmounts{m.RepeatingGroup.Get(i)}
}

// NoCollateralAmounts is a repeating group element, Tag 1703.
type NoCollateralAmounts struct {
	*quickfix.Group
}

// SetCurrentCollateralAmount sets CurrentCollateralAmount, Tag 1704.
func (m NoCollateralAmounts) SetCurrentCollateralAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewCurrentCollateralAmount(value, scale))
}

// SetCollateralCurrency sets CollateralCurrency, Tag 1705.
func (m NoCollateralAmounts) SetCollateralCurrency(v string) {
	m.Set(field.NewCollateralCurrency(v))
}

// SetCollateralType sets CollateralType, Tag 1706.
func (m NoCollateralAmounts) SetCollateralType(v string) {
	m.Set(field.NewCollateralType(v))
}

// SetHaircutIndicator sets HaircutIndicator, Tag 1902.
func (m NoCollateralAmounts) SetHaircutIndicator(v bool) {
	m.Set(field.NewHaircutIndicator(v))
}

// SetCollateralFXRate sets CollateralFXRate, Tag 2090.
func (m NoCollateralAmounts) SetCollateralFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCollateralFXRate(value, scale))
}

// SetCollateralFXRateCalc sets CollateralFXRateCalc, Tag 2091.
func (m NoCollateralAmounts) SetCollateralFXRateCalc(v string) {
	m.Set(field.NewCollateralFXRateCalc(v))
}

// SetCollateralAmountMarketSegmentID sets CollateralAmountMarketSegmentID, Tag 2092.
func (m NoCollateralAmounts) SetCollateralAmountMarketSegmentID(v string) {
	m.Set(field.NewCollateralAmountMarketSegmentID(v))
}

// SetCollateralAmountMarketID sets CollateralAmountMarketID, Tag 2093.
func (m NoCollateralAmounts) SetCollateralAmountMarketID(v string) {
	m.Set(field.NewCollateralAmountMarketID(v))
}

// SetCollateralPortfolioID sets CollateralPortfolioID, Tag 2350.
func (m NoCollateralAmounts) SetCollateralPortfolioID(v string) {
	m.Set(field.NewCollateralPortfolioID(v))
}

// SetCollateralAmountType sets CollateralAmountType, Tag 2632.
func (m NoCollateralAmounts) SetCollateralAmountType(v enum.CollateralAmountType) {
	m.Set(field.NewCollateralAmountType(v))
}

// SetCollateralPercentOverage sets CollateralPercentOverage, Tag 2690.
func (m NoCollateralAmounts) SetCollateralPercentOverage(value decimal.Decimal, scale int32) {
	m.Set(field.NewCollateralPercentOverage(value, scale))
}

// SetCollateralMarketPrice sets CollateralMarketPrice, Tag 2689.
func (m NoCollateralAmounts) SetCollateralMarketPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCollateralMarketPrice(value, scale))
}

// SetCollateralReinvestmentRate sets CollateralReinvestmentRate, Tag 2840.
func (m NoCollateralAmounts) SetCollateralReinvestmentRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCollateralReinvestmentRate(value, scale))
}

// SetNoCollateralReinvestments sets NoCollateralReinvestments, Tag 2845.
func (m NoCollateralAmounts) SetNoCollateralReinvestments(f NoCollateralReinvestmentsRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingRefID sets UnderlyingRefID, Tag 2841.
func (m NoCollateralAmounts) SetUnderlyingRefID(v string) {
	m.Set(field.NewUnderlyingRefID(v))
}

// GetCurrentCollateralAmount gets CurrentCollateralAmount, Tag 1704.
func (m NoCollateralAmounts) GetCurrentCollateralAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CurrentCollateralAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralCurrency gets CollateralCurrency, Tag 1705.
func (m NoCollateralAmounts) GetCollateralCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CollateralCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralType gets CollateralType, Tag 1706.
func (m NoCollateralAmounts) GetCollateralType() (v string, err quickfix.MessageRejectError) {
	var f field.CollateralTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetHaircutIndicator gets HaircutIndicator, Tag 1902.
func (m NoCollateralAmounts) GetHaircutIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.HaircutIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralFXRate gets CollateralFXRate, Tag 2090.
func (m NoCollateralAmounts) GetCollateralFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CollateralFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralFXRateCalc gets CollateralFXRateCalc, Tag 2091.
func (m NoCollateralAmounts) GetCollateralFXRateCalc() (v string, err quickfix.MessageRejectError) {
	var f field.CollateralFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralAmountMarketSegmentID gets CollateralAmountMarketSegmentID, Tag 2092.
func (m NoCollateralAmounts) GetCollateralAmountMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.CollateralAmountMarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralAmountMarketID gets CollateralAmountMarketID, Tag 2093.
func (m NoCollateralAmounts) GetCollateralAmountMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.CollateralAmountMarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralPortfolioID gets CollateralPortfolioID, Tag 2350.
func (m NoCollateralAmounts) GetCollateralPortfolioID() (v string, err quickfix.MessageRejectError) {
	var f field.CollateralPortfolioIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralAmountType gets CollateralAmountType, Tag 2632.
func (m NoCollateralAmounts) GetCollateralAmountType() (v enum.CollateralAmountType, err quickfix.MessageRejectError) {
	var f field.CollateralAmountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralPercentOverage gets CollateralPercentOverage, Tag 2690.
func (m NoCollateralAmounts) GetCollateralPercentOverage() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CollateralPercentOverageField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralMarketPrice gets CollateralMarketPrice, Tag 2689.
func (m NoCollateralAmounts) GetCollateralMarketPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CollateralMarketPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralReinvestmentRate gets CollateralReinvestmentRate, Tag 2840.
func (m NoCollateralAmounts) GetCollateralReinvestmentRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CollateralReinvestmentRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoCollateralReinvestments gets NoCollateralReinvestments, Tag 2845.
func (m NoCollateralAmounts) GetNoCollateralReinvestments() (NoCollateralReinvestmentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoCollateralReinvestmentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingRefID gets UnderlyingRefID, Tag 2841.
func (m NoCollateralAmounts) GetUnderlyingRefID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasCurrentCollateralAmount returns true if CurrentCollateralAmount is present, Tag 1704.
func (m NoCollateralAmounts) HasCurrentCollateralAmount() bool {
	return m.Has(tag.CurrentCollateralAmount)
}

// HasCollateralCurrency returns true if CollateralCurrency is present, Tag 1705.
func (m NoCollateralAmounts) HasCollateralCurrency() bool {
	return m.Has(tag.CollateralCurrency)
}

// HasCollateralType returns true if CollateralType is present, Tag 1706.
func (m NoCollateralAmounts) HasCollateralType() bool {
	return m.Has(tag.CollateralType)
}

// HasHaircutIndicator returns true if HaircutIndicator is present, Tag 1902.
func (m NoCollateralAmounts) HasHaircutIndicator() bool {
	return m.Has(tag.HaircutIndicator)
}

// HasCollateralFXRate returns true if CollateralFXRate is present, Tag 2090.
func (m NoCollateralAmounts) HasCollateralFXRate() bool {
	return m.Has(tag.CollateralFXRate)
}

// HasCollateralFXRateCalc returns true if CollateralFXRateCalc is present, Tag 2091.
func (m NoCollateralAmounts) HasCollateralFXRateCalc() bool {
	return m.Has(tag.CollateralFXRateCalc)
}

// HasCollateralAmountMarketSegmentID returns true if CollateralAmountMarketSegmentID is present, Tag 2092.
func (m NoCollateralAmounts) HasCollateralAmountMarketSegmentID() bool {
	return m.Has(tag.CollateralAmountMarketSegmentID)
}

// HasCollateralAmountMarketID returns true if CollateralAmountMarketID is present, Tag 2093.
func (m NoCollateralAmounts) HasCollateralAmountMarketID() bool {
	return m.Has(tag.CollateralAmountMarketID)
}

// HasCollateralPortfolioID returns true if CollateralPortfolioID is present, Tag 2350.
func (m NoCollateralAmounts) HasCollateralPortfolioID() bool {
	return m.Has(tag.CollateralPortfolioID)
}

// HasCollateralAmountType returns true if CollateralAmountType is present, Tag 2632.
func (m NoCollateralAmounts) HasCollateralAmountType() bool {
	return m.Has(tag.CollateralAmountType)
}

// HasCollateralPercentOverage returns true if CollateralPercentOverage is present, Tag 2690.
func (m NoCollateralAmounts) HasCollateralPercentOverage() bool {
	return m.Has(tag.CollateralPercentOverage)
}

// HasCollateralMarketPrice returns true if CollateralMarketPrice is present, Tag 2689.
func (m NoCollateralAmounts) HasCollateralMarketPrice() bool {
	return m.Has(tag.CollateralMarketPrice)
}

// HasCollateralReinvestmentRate returns true if CollateralReinvestmentRate is present, Tag 2840.
func (m NoCollateralAmounts) HasCollateralReinvestmentRate() bool {
	return m.Has(tag.CollateralReinvestmentRate)
}

// HasNoCollateralReinvestments returns true if NoCollateralReinvestments is present, Tag 2845.
func (m NoCollateralAmounts) HasNoCollateralReinvestments() bool {
	return m.Has(tag.NoCollateralReinvestments)
}

// HasUnderlyingRefID returns true if UnderlyingRefID is present, Tag 2841.
func (m NoCollateralAmounts) HasUnderlyingRefID() bool {
	return m.Has(tag.UnderlyingRefID)
}

// NoCollateralReinvestments is a repeating group element, Tag 2845.
type NoCollateralReinvestments struct {
	*quickfix.Group
}

// SetCollateralReinvestmentType sets CollateralReinvestmentType, Tag 2844.
func (m NoCollateralReinvestments) SetCollateralReinvestmentType(v enum.CollateralReinvestmentType) {
	m.Set(field.NewCollateralReinvestmentType(v))
}

// SetCollateralReinvestmentAmount sets CollateralReinvestmentAmount, Tag 2842.
func (m NoCollateralReinvestments) SetCollateralReinvestmentAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewCollateralReinvestmentAmount(value, scale))
}

// SetCollateralReinvestmentCurrency sets CollateralReinvestmentCurrency, Tag 2843.
func (m NoCollateralReinvestments) SetCollateralReinvestmentCurrency(v string) {
	m.Set(field.NewCollateralReinvestmentCurrency(v))
}

// GetCollateralReinvestmentType gets CollateralReinvestmentType, Tag 2844.
func (m NoCollateralReinvestments) GetCollateralReinvestmentType() (v enum.CollateralReinvestmentType, err quickfix.MessageRejectError) {
	var f field.CollateralReinvestmentTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralReinvestmentAmount gets CollateralReinvestmentAmount, Tag 2842.
func (m NoCollateralReinvestments) GetCollateralReinvestmentAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CollateralReinvestmentAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollateralReinvestmentCurrency gets CollateralReinvestmentCurrency, Tag 2843.
func (m NoCollateralReinvestments) GetCollateralReinvestmentCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CollateralReinvestmentCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasCollateralReinvestmentType returns true if CollateralReinvestmentType is present, Tag 2844.
func (m NoCollateralReinvestments) HasCollateralReinvestmentType() bool {
	return m.Has(tag.CollateralReinvestmentType)
}

// HasCollateralReinvestmentAmount returns true if CollateralReinvestmentAmount is present, Tag 2842.
func (m NoCollateralReinvestments) HasCollateralReinvestmentAmount() bool {
	return m.Has(tag.CollateralReinvestmentAmount)
}

// HasCollateralReinvestmentCurrency returns true if CollateralReinvestmentCurrency is present, Tag 2843.
func (m NoCollateralReinvestments) HasCollateralReinvestmentCurrency() bool {
	return m.Has(tag.CollateralReinvestmentCurrency)
}

// NoCollateralReinvestmentsRepeatingGroup is a repeating group, Tag 2845.
type NoCollateralReinvestmentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoCollateralReinvestmentsRepeatingGroup returns an initialized, NoCollateralReinvestmentsRepeatingGroup.
func NewNoCollateralReinvestmentsRepeatingGroup() NoCollateralReinvestmentsRepeatingGroup {
	return NoCollateralReinvestmentsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoCollateralReinvestments,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.CollateralReinvestmentType),
				quickfix.GroupElement(tag.CollateralReinvestmentAmount),
				quickfix.GroupElement(tag.CollateralReinvestmentCurrency),
			},
		),
	}
}

// Add create and append a new NoCollateralReinvestments to this group.
func (m NoCollateralReinvestmentsRepeatingGroup) Add() NoCollateralReinvestments {
	g := m.RepeatingGroup.Add()
	return NoCollateralReinvestments{g}
}

// Get returns the ith NoCollateralReinvestments in the NoCollateralReinvestmentsRepeatinGroup.
func (m NoCollateralReinvestmentsRepeatingGroup) Get(i int) NoCollateralReinvestments {
	return NoCollateralReinvestments{m.RepeatingGroup.Get(i)}
}

// NoCollateralAmountsRepeatingGroup is a repeating group, Tag 1703.
type NoCollateralAmountsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoCollateralAmountsRepeatingGroup returns an initialized, NoCollateralAmountsRepeatingGroup.
func NewNoCollateralAmountsRepeatingGroup() NoCollateralAmountsRepeatingGroup {
	return NoCollateralAmountsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoCollateralAmounts,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.CurrentCollateralAmount),
				quickfix.GroupElement(tag.CollateralCurrency),
				quickfix.GroupElement(tag.CollateralType),
				quickfix.GroupElement(tag.HaircutIndicator),
				quickfix.GroupElement(tag.CollateralFXRate),
				quickfix.GroupElement(tag.CollateralFXRateCalc),
				quickfix.GroupElement(tag.CollateralAmountMarketSegmentID),
				quickfix.GroupElement(tag.CollateralAmountMarketID),
				quickfix.GroupElement(tag.CollateralPortfolioID),
				quickfix.GroupElement(tag.CollateralAmountType),
				quickfix.GroupElement(tag.CollateralPercentOverage),
				quickfix.GroupElement(tag.CollateralMarketPrice),
				quickfix.GroupElement(tag.CollateralReinvestmentRate),
				NewNoCollateralReinvestmentsRepeatingGroup(),
				quickfix.GroupElement(tag.UnderlyingRefID),
			},
		),
	}
}

// Add create and append a new NoCollateralAmounts to this group.
func (m NoCollateralAmountsRepeatingGroup) Add() NoCollateralAmounts {
	g := m.RepeatingGroup.Add()
	return NoCollateralAmounts{g}
}

// Get returns the ith NoCollateralAmounts in the NoCollateralAmountsRepeatinGroup.
func (m NoCollateralAmountsRepeatingGroup) Get(i int) NoCollateralAmounts {
	return NoCollateralAmounts{m.RepeatingGroup.Get(i)}
}

// NoPayCollects is a repeating group element, Tag 1707.
type NoPayCollects struct {
	*quickfix.Group
}

// SetPayCollectType sets PayCollectType, Tag 1708.
func (m NoPayCollects) SetPayCollectType(v string) {
	m.Set(field.NewPayCollectType(v))
}

// SetPayCollectCurrency sets PayCollectCurrency, Tag 1709.
func (m NoPayCollects) SetPayCollectCurrency(v string) {
	m.Set(field.NewPayCollectCurrency(v))
}

// SetPayAmount sets PayAmount, Tag 1710.
func (m NoPayCollects) SetPayAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewPayAmount(value, scale))
}

// SetCollectAmount sets CollectAmount, Tag 1711.
func (m NoPayCollects) SetCollectAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewCollectAmount(value, scale))
}

// SetPayCollectMarketSegmentID sets PayCollectMarketSegmentID, Tag 1712.
func (m NoPayCollects) SetPayCollectMarketSegmentID(v string) {
	m.Set(field.NewPayCollectMarketSegmentID(v))
}

// SetPayCollectMarketID sets PayCollectMarketID, Tag 1713.
func (m NoPayCollects) SetPayCollectMarketID(v string) {
	m.Set(field.NewPayCollectMarketID(v))
}

// SetPayCollectFXRate sets PayCollectFXRate, Tag 2094.
func (m NoPayCollects) SetPayCollectFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewPayCollectFXRate(value, scale))
}

// SetPayCollectFXRateCalc sets PayCollectFXRateCalc, Tag 2095.
func (m NoPayCollects) SetPayCollectFXRateCalc(v string) {
	m.Set(field.NewPayCollectFXRateCalc(v))
}

// GetPayCollectType gets PayCollectType, Tag 1708.
func (m NoPayCollects) GetPayCollectType() (v string, err quickfix.MessageRejectError) {
	var f field.PayCollectTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayCollectCurrency gets PayCollectCurrency, Tag 1709.
func (m NoPayCollects) GetPayCollectCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.PayCollectCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayAmount gets PayAmount, Tag 1710.
func (m NoPayCollects) GetPayAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PayAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollectAmount gets CollectAmount, Tag 1711.
func (m NoPayCollects) GetCollectAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CollectAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayCollectMarketSegmentID gets PayCollectMarketSegmentID, Tag 1712.
func (m NoPayCollects) GetPayCollectMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.PayCollectMarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayCollectMarketID gets PayCollectMarketID, Tag 1713.
func (m NoPayCollects) GetPayCollectMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.PayCollectMarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayCollectFXRate gets PayCollectFXRate, Tag 2094.
func (m NoPayCollects) GetPayCollectFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PayCollectFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPayCollectFXRateCalc gets PayCollectFXRateCalc, Tag 2095.
func (m NoPayCollects) GetPayCollectFXRateCalc() (v string, err quickfix.MessageRejectError) {
	var f field.PayCollectFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPayCollectType returns true if PayCollectType is present, Tag 1708.
func (m NoPayCollects) HasPayCollectType() bool {
	return m.Has(tag.PayCollectType)
}

// HasPayCollectCurrency returns true if PayCollectCurrency is present, Tag 1709.
func (m NoPayCollects) HasPayCollectCurrency() bool {
	return m.Has(tag.PayCollectCurrency)
}

// HasPayAmount returns true if PayAmount is present, Tag 1710.
func (m NoPayCollects) HasPayAmount() bool {
	return m.Has(tag.PayAmount)
}

// HasCollectAmount returns true if CollectAmount is present, Tag 1711.
func (m NoPayCollects) HasCollectAmount() bool {
	return m.Has(tag.CollectAmount)
}

// HasPayCollectMarketSegmentID returns true if PayCollectMarketSegmentID is present, Tag 1712.
func (m NoPayCollects) HasPayCollectMarketSegmentID() bool {
	return m.Has(tag.PayCollectMarketSegmentID)
}

// HasPayCollectMarketID returns true if PayCollectMarketID is present, Tag 1713.
func (m NoPayCollects) HasPayCollectMarketID() bool {
	return m.Has(tag.PayCollectMarketID)
}

// HasPayCollectFXRate returns true if PayCollectFXRate is present, Tag 2094.
func (m NoPayCollects) HasPayCollectFXRate() bool {
	return m.Has(tag.PayCollectFXRate)
}

// HasPayCollectFXRateCalc returns true if PayCollectFXRateCalc is present, Tag 2095.
func (m NoPayCollects) HasPayCollectFXRateCalc() bool {
	return m.Has(tag.PayCollectFXRateCalc)
}

// NoPayCollectsRepeatingGroup is a repeating group, Tag 1707.
type NoPayCollectsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPayCollectsRepeatingGroup returns an initialized, NoPayCollectsRepeatingGroup.
func NewNoPayCollectsRepeatingGroup() NoPayCollectsRepeatingGroup {
	return NoPayCollectsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPayCollects,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.PayCollectType),
				quickfix.GroupElement(tag.PayCollectCurrency),
				quickfix.GroupElement(tag.PayAmount),
				quickfix.GroupElement(tag.CollectAmount),
				quickfix.GroupElement(tag.PayCollectMarketSegmentID),
				quickfix.GroupElement(tag.PayCollectMarketID),
				quickfix.GroupElement(tag.PayCollectFXRate),
				quickfix.GroupElement(tag.PayCollectFXRateCalc),
			},
		),
	}
}

// Add create and append a new NoPayCollects to this group.
func (m NoPayCollectsRepeatingGroup) Add() NoPayCollects {
	g := m.RepeatingGroup.Add()
	return NoPayCollects{g}
}

// Get returns the ith NoPayCollects in the NoPayCollectsRepeatinGroup.
func (m NoPayCollectsRepeatingGroup) Get(i int) NoPayCollects {
	return NoPayCollects{m.RepeatingGroup.Get(i)}
}
