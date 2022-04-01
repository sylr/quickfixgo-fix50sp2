package marketdefinitionupdatereport

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"
)

// MarketDefinitionUpdateReport is the fix50sp2 MarketDefinitionUpdateReport type, MsgType = BV.
type MarketDefinitionUpdateReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a MarketDefinitionUpdateReport from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) MarketDefinitionUpdateReport {
	return MarketDefinitionUpdateReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m MarketDefinitionUpdateReport) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a MarketDefinitionUpdateReport initialized with the required fields for MarketDefinitionUpdateReport.
func New(marketreportid field.MarketReportIDField, marketid field.MarketIDField) (m MarketDefinitionUpdateReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BV"))
	m.Set(marketreportid)
	m.Set(marketid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg MarketDefinitionUpdateReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "BV", r
}

// SetCurrency sets Currency, Tag 15.
func (m MarketDefinitionUpdateReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetText sets Text, Tag 58.
func (m MarketDefinitionUpdateReport) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m MarketDefinitionUpdateReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m MarketDefinitionUpdateReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m MarketDefinitionUpdateReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetPriceType sets PriceType, Tag 423.
func (m MarketDefinitionUpdateReport) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m MarketDefinitionUpdateReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRoundLot sets RoundLot, Tag 561.
func (m MarketDefinitionUpdateReport) SetRoundLot(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundLot(value, scale))
}

// SetMinTradeVol sets MinTradeVol, Tag 562.
func (m MarketDefinitionUpdateReport) SetMinTradeVol(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinTradeVol(value, scale))
}

// SetExpirationCycle sets ExpirationCycle, Tag 827.
func (m MarketDefinitionUpdateReport) SetExpirationCycle(v enum.ExpirationCycle) {
	m.Set(field.NewExpirationCycle(v))
}

// SetMaxTradeVol sets MaxTradeVol, Tag 1140.
func (m MarketDefinitionUpdateReport) SetMaxTradeVol(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxTradeVol(value, scale))
}

// SetNoMDFeedTypes sets NoMDFeedTypes, Tag 1141.
func (m MarketDefinitionUpdateReport) SetNoMDFeedTypes(f NoMDFeedTypesRepeatingGroup) {
	m.SetGroup(f)
}

// SetMaxPriceVariation sets MaxPriceVariation, Tag 1143.
func (m MarketDefinitionUpdateReport) SetMaxPriceVariation(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxPriceVariation(value, scale))
}

// SetImpliedMarketIndicator sets ImpliedMarketIndicator, Tag 1144.
func (m MarketDefinitionUpdateReport) SetImpliedMarketIndicator(v enum.ImpliedMarketIndicator) {
	m.Set(field.NewImpliedMarketIndicator(v))
}

// SetLowLimitPrice sets LowLimitPrice, Tag 1148.
func (m MarketDefinitionUpdateReport) SetLowLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLowLimitPrice(value, scale))
}

// SetHighLimitPrice sets HighLimitPrice, Tag 1149.
func (m MarketDefinitionUpdateReport) SetHighLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewHighLimitPrice(value, scale))
}

// SetTradingReferencePrice sets TradingReferencePrice, Tag 1150.
func (m MarketDefinitionUpdateReport) SetTradingReferencePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewTradingReferencePrice(value, scale))
}

// SetApplID sets ApplID, Tag 1180.
func (m MarketDefinitionUpdateReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181.
func (m MarketDefinitionUpdateReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetNoTickRules sets NoTickRules, Tag 1205.
func (m MarketDefinitionUpdateReport) SetNoTickRules(f NoTickRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoExecInstRules sets NoExecInstRules, Tag 1232.
func (m MarketDefinitionUpdateReport) SetNoExecInstRules(f NoExecInstRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoLotTypeRules sets NoLotTypeRules, Tag 1234.
func (m MarketDefinitionUpdateReport) SetNoLotTypeRules(f NoLotTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoMatchRules sets NoMatchRules, Tag 1235.
func (m MarketDefinitionUpdateReport) SetNoMatchRules(f NoMatchRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoOrdTypeRules sets NoOrdTypeRules, Tag 1237.
func (m MarketDefinitionUpdateReport) SetNoOrdTypeRules(f NoOrdTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoTimeInForceRules sets NoTimeInForceRules, Tag 1239.
func (m MarketDefinitionUpdateReport) SetNoTimeInForceRules(f NoTimeInForceRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetTradingCurrency sets TradingCurrency, Tag 1245.
func (m MarketDefinitionUpdateReport) SetTradingCurrency(v string) {
	m.Set(field.NewTradingCurrency(v))
}

// SetMarketSegmentID sets MarketSegmentID, Tag 1300.
func (m MarketDefinitionUpdateReport) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

// SetMarketID sets MarketID, Tag 1301.
func (m MarketDefinitionUpdateReport) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

// SetPriceLimitType sets PriceLimitType, Tag 1306.
func (m MarketDefinitionUpdateReport) SetPriceLimitType(v enum.PriceLimitType) {
	m.Set(field.NewPriceLimitType(v))
}

// SetParentMktSegmID sets ParentMktSegmID, Tag 1325.
func (m MarketDefinitionUpdateReport) SetParentMktSegmID(v string) {
	m.Set(field.NewParentMktSegmID(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350.
func (m MarketDefinitionUpdateReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352.
func (m MarketDefinitionUpdateReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// SetMultilegModel sets MultilegModel, Tag 1377.
func (m MarketDefinitionUpdateReport) SetMultilegModel(v enum.MultilegModel) {
	m.Set(field.NewMultilegModel(v))
}

// SetMultilegPriceMethod sets MultilegPriceMethod, Tag 1378.
func (m MarketDefinitionUpdateReport) SetMultilegPriceMethod(v enum.MultilegPriceMethod) {
	m.Set(field.NewMultilegPriceMethod(v))
}

// SetMarketReqID sets MarketReqID, Tag 1393.
func (m MarketDefinitionUpdateReport) SetMarketReqID(v string) {
	m.Set(field.NewMarketReqID(v))
}

// SetMarketReportID sets MarketReportID, Tag 1394.
func (m MarketDefinitionUpdateReport) SetMarketReportID(v string) {
	m.Set(field.NewMarketReportID(v))
}

// SetMarketUpdateAction sets MarketUpdateAction, Tag 1395.
func (m MarketDefinitionUpdateReport) SetMarketUpdateAction(v enum.MarketUpdateAction) {
	m.Set(field.NewMarketUpdateAction(v))
}

// SetMarketSegmentDesc sets MarketSegmentDesc, Tag 1396.
func (m MarketDefinitionUpdateReport) SetMarketSegmentDesc(v string) {
	m.Set(field.NewMarketSegmentDesc(v))
}

// SetEncodedMktSegmDescLen sets EncodedMktSegmDescLen, Tag 1397.
func (m MarketDefinitionUpdateReport) SetEncodedMktSegmDescLen(v int) {
	m.Set(field.NewEncodedMktSegmDescLen(v))
}

// SetEncodedMktSegmDesc sets EncodedMktSegmDesc, Tag 1398.
func (m MarketDefinitionUpdateReport) SetEncodedMktSegmDesc(v string) {
	m.Set(field.NewEncodedMktSegmDesc(v))
}

// SetNoInstrumentScopes sets NoInstrumentScopes, Tag 1656.
func (m MarketDefinitionUpdateReport) SetNoInstrumentScopes(f NoInstrumentScopesRepeatingGroup) {
	m.SetGroup(f)
}

// SetTradeVolType sets TradeVolType, Tag 1786.
func (m MarketDefinitionUpdateReport) SetTradeVolType(v enum.TradeVolType) {
	m.Set(field.NewTradeVolType(v))
}

// SetEffectiveBusinessDate sets EffectiveBusinessDate, Tag 2400.
func (m MarketDefinitionUpdateReport) SetEffectiveBusinessDate(v string) {
	m.Set(field.NewEffectiveBusinessDate(v))
}

// SetMarketSegmentStatus sets MarketSegmentStatus, Tag 2542.
func (m MarketDefinitionUpdateReport) SetMarketSegmentStatus(v enum.MarketSegmentStatus) {
	m.Set(field.NewMarketSegmentStatus(v))
}

// SetMarketSegmentType sets MarketSegmentType, Tag 2543.
func (m MarketDefinitionUpdateReport) SetMarketSegmentType(v enum.MarketSegmentType) {
	m.Set(field.NewMarketSegmentType(v))
}

// SetMarketSegmentSubType sets MarketSegmentSubType, Tag 2544.
func (m MarketDefinitionUpdateReport) SetMarketSegmentSubType(v enum.MarketSegmentSubType) {
	m.Set(field.NewMarketSegmentSubType(v))
}

// SetNoRelatedMarketSegments sets NoRelatedMarketSegments, Tag 2545.
func (m MarketDefinitionUpdateReport) SetNoRelatedMarketSegments(f NoRelatedMarketSegmentsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoAuctionTypeRules sets NoAuctionTypeRules, Tag 2548.
func (m MarketDefinitionUpdateReport) SetNoAuctionTypeRules(f NoAuctionTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoPriceRangeRules sets NoPriceRangeRules, Tag 2550.
func (m MarketDefinitionUpdateReport) SetNoPriceRangeRules(f NoPriceRangeRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetFastMarketPercentage sets FastMarketPercentage, Tag 2557.
func (m MarketDefinitionUpdateReport) SetFastMarketPercentage(value decimal.Decimal, scale int32) {
	m.Set(field.NewFastMarketPercentage(value, scale))
}

// SetNoQuoteSizeRules sets NoQuoteSizeRules, Tag 2558.
func (m MarketDefinitionUpdateReport) SetNoQuoteSizeRules(f NoQuoteSizeRulesRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteSideIndicator sets QuoteSideIndicator, Tag 2559.
func (m MarketDefinitionUpdateReport) SetQuoteSideIndicator(v bool) {
	m.Set(field.NewQuoteSideIndicator(v))
}

// SetNoFlexProductEligibilities sets NoFlexProductEligibilities, Tag 2560.
func (m MarketDefinitionUpdateReport) SetNoFlexProductEligibilities(f NoFlexProductEligibilitiesRepeatingGroup) {
	m.SetGroup(f)
}

// GetCurrency gets Currency, Tag 15.
func (m MarketDefinitionUpdateReport) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m MarketDefinitionUpdateReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m MarketDefinitionUpdateReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m MarketDefinitionUpdateReport) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m MarketDefinitionUpdateReport) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceType gets PriceType, Tag 423.
func (m MarketDefinitionUpdateReport) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m MarketDefinitionUpdateReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRoundLot gets RoundLot, Tag 561.
func (m MarketDefinitionUpdateReport) GetRoundLot() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundLotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinTradeVol gets MinTradeVol, Tag 562.
func (m MarketDefinitionUpdateReport) GetMinTradeVol() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinTradeVolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExpirationCycle gets ExpirationCycle, Tag 827.
func (m MarketDefinitionUpdateReport) GetExpirationCycle() (v enum.ExpirationCycle, err quickfix.MessageRejectError) {
	var f field.ExpirationCycleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaxTradeVol gets MaxTradeVol, Tag 1140.
func (m MarketDefinitionUpdateReport) GetMaxTradeVol() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxTradeVolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMDFeedTypes gets NoMDFeedTypes, Tag 1141.
func (m MarketDefinitionUpdateReport) GetNoMDFeedTypes() (NoMDFeedTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDFeedTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMaxPriceVariation gets MaxPriceVariation, Tag 1143.
func (m MarketDefinitionUpdateReport) GetMaxPriceVariation() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxPriceVariationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetImpliedMarketIndicator gets ImpliedMarketIndicator, Tag 1144.
func (m MarketDefinitionUpdateReport) GetImpliedMarketIndicator() (v enum.ImpliedMarketIndicator, err quickfix.MessageRejectError) {
	var f field.ImpliedMarketIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLowLimitPrice gets LowLimitPrice, Tag 1148.
func (m MarketDefinitionUpdateReport) GetLowLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LowLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetHighLimitPrice gets HighLimitPrice, Tag 1149.
func (m MarketDefinitionUpdateReport) GetHighLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.HighLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingReferencePrice gets TradingReferencePrice, Tag 1150.
func (m MarketDefinitionUpdateReport) GetTradingReferencePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TradingReferencePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180.
func (m MarketDefinitionUpdateReport) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181.
func (m MarketDefinitionUpdateReport) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTickRules gets NoTickRules, Tag 1205.
func (m MarketDefinitionUpdateReport) GetNoTickRules() (NoTickRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTickRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoExecInstRules gets NoExecInstRules, Tag 1232.
func (m MarketDefinitionUpdateReport) GetNoExecInstRules() (NoExecInstRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecInstRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoLotTypeRules gets NoLotTypeRules, Tag 1234.
func (m MarketDefinitionUpdateReport) GetNoLotTypeRules() (NoLotTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLotTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoMatchRules gets NoMatchRules, Tag 1235.
func (m MarketDefinitionUpdateReport) GetNoMatchRules() (NoMatchRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMatchRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoOrdTypeRules gets NoOrdTypeRules, Tag 1237.
func (m MarketDefinitionUpdateReport) GetNoOrdTypeRules() (NoOrdTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoTimeInForceRules gets NoTimeInForceRules, Tag 1239.
func (m MarketDefinitionUpdateReport) GetNoTimeInForceRules() (NoTimeInForceRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTimeInForceRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetTradingCurrency gets TradingCurrency, Tag 1245.
func (m MarketDefinitionUpdateReport) GetTradingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.TradingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentID gets MarketSegmentID, Tag 1300.
func (m MarketDefinitionUpdateReport) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketID gets MarketID, Tag 1301.
func (m MarketDefinitionUpdateReport) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceLimitType gets PriceLimitType, Tag 1306.
func (m MarketDefinitionUpdateReport) GetPriceLimitType() (v enum.PriceLimitType, err quickfix.MessageRejectError) {
	var f field.PriceLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetParentMktSegmID gets ParentMktSegmID, Tag 1325.
func (m MarketDefinitionUpdateReport) GetParentMktSegmID() (v string, err quickfix.MessageRejectError) {
	var f field.ParentMktSegmIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350.
func (m MarketDefinitionUpdateReport) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352.
func (m MarketDefinitionUpdateReport) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMultilegModel gets MultilegModel, Tag 1377.
func (m MarketDefinitionUpdateReport) GetMultilegModel() (v enum.MultilegModel, err quickfix.MessageRejectError) {
	var f field.MultilegModelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMultilegPriceMethod gets MultilegPriceMethod, Tag 1378.
func (m MarketDefinitionUpdateReport) GetMultilegPriceMethod() (v enum.MultilegPriceMethod, err quickfix.MessageRejectError) {
	var f field.MultilegPriceMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketReqID gets MarketReqID, Tag 1393.
func (m MarketDefinitionUpdateReport) GetMarketReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketReportID gets MarketReportID, Tag 1394.
func (m MarketDefinitionUpdateReport) GetMarketReportID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketUpdateAction gets MarketUpdateAction, Tag 1395.
func (m MarketDefinitionUpdateReport) GetMarketUpdateAction() (v enum.MarketUpdateAction, err quickfix.MessageRejectError) {
	var f field.MarketUpdateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentDesc gets MarketSegmentDesc, Tag 1396.
func (m MarketDefinitionUpdateReport) GetMarketSegmentDesc() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedMktSegmDescLen gets EncodedMktSegmDescLen, Tag 1397.
func (m MarketDefinitionUpdateReport) GetEncodedMktSegmDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedMktSegmDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedMktSegmDesc gets EncodedMktSegmDesc, Tag 1398.
func (m MarketDefinitionUpdateReport) GetEncodedMktSegmDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedMktSegmDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentScopes gets NoInstrumentScopes, Tag 1656.
func (m MarketDefinitionUpdateReport) GetNoInstrumentScopes() (NoInstrumentScopesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentScopesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetTradeVolType gets TradeVolType, Tag 1786.
func (m MarketDefinitionUpdateReport) GetTradeVolType() (v enum.TradeVolType, err quickfix.MessageRejectError) {
	var f field.TradeVolTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEffectiveBusinessDate gets EffectiveBusinessDate, Tag 2400.
func (m MarketDefinitionUpdateReport) GetEffectiveBusinessDate() (v string, err quickfix.MessageRejectError) {
	var f field.EffectiveBusinessDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentStatus gets MarketSegmentStatus, Tag 2542.
func (m MarketDefinitionUpdateReport) GetMarketSegmentStatus() (v enum.MarketSegmentStatus, err quickfix.MessageRejectError) {
	var f field.MarketSegmentStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentType gets MarketSegmentType, Tag 2543.
func (m MarketDefinitionUpdateReport) GetMarketSegmentType() (v enum.MarketSegmentType, err quickfix.MessageRejectError) {
	var f field.MarketSegmentTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentSubType gets MarketSegmentSubType, Tag 2544.
func (m MarketDefinitionUpdateReport) GetMarketSegmentSubType() (v enum.MarketSegmentSubType, err quickfix.MessageRejectError) {
	var f field.MarketSegmentSubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedMarketSegments gets NoRelatedMarketSegments, Tag 2545.
func (m MarketDefinitionUpdateReport) GetNoRelatedMarketSegments() (NoRelatedMarketSegmentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedMarketSegmentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoAuctionTypeRules gets NoAuctionTypeRules, Tag 2548.
func (m MarketDefinitionUpdateReport) GetNoAuctionTypeRules() (NoAuctionTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAuctionTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoPriceRangeRules gets NoPriceRangeRules, Tag 2550.
func (m MarketDefinitionUpdateReport) GetNoPriceRangeRules() (NoPriceRangeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPriceRangeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetFastMarketPercentage gets FastMarketPercentage, Tag 2557.
func (m MarketDefinitionUpdateReport) GetFastMarketPercentage() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FastMarketPercentageField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoQuoteSizeRules gets NoQuoteSizeRules, Tag 2558.
func (m MarketDefinitionUpdateReport) GetNoQuoteSizeRules() (NoQuoteSizeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoQuoteSizeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteSideIndicator gets QuoteSideIndicator, Tag 2559.
func (m MarketDefinitionUpdateReport) GetQuoteSideIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.QuoteSideIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoFlexProductEligibilities gets NoFlexProductEligibilities, Tag 2560.
func (m MarketDefinitionUpdateReport) GetNoFlexProductEligibilities() (NoFlexProductEligibilitiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoFlexProductEligibilitiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m MarketDefinitionUpdateReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasText returns true if Text is present, Tag 58.
func (m MarketDefinitionUpdateReport) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m MarketDefinitionUpdateReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m MarketDefinitionUpdateReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m MarketDefinitionUpdateReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasPriceType returns true if PriceType is present, Tag 423.
func (m MarketDefinitionUpdateReport) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m MarketDefinitionUpdateReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasRoundLot returns true if RoundLot is present, Tag 561.
func (m MarketDefinitionUpdateReport) HasRoundLot() bool {
	return m.Has(tag.RoundLot)
}

// HasMinTradeVol returns true if MinTradeVol is present, Tag 562.
func (m MarketDefinitionUpdateReport) HasMinTradeVol() bool {
	return m.Has(tag.MinTradeVol)
}

// HasExpirationCycle returns true if ExpirationCycle is present, Tag 827.
func (m MarketDefinitionUpdateReport) HasExpirationCycle() bool {
	return m.Has(tag.ExpirationCycle)
}

// HasMaxTradeVol returns true if MaxTradeVol is present, Tag 1140.
func (m MarketDefinitionUpdateReport) HasMaxTradeVol() bool {
	return m.Has(tag.MaxTradeVol)
}

// HasNoMDFeedTypes returns true if NoMDFeedTypes is present, Tag 1141.
func (m MarketDefinitionUpdateReport) HasNoMDFeedTypes() bool {
	return m.Has(tag.NoMDFeedTypes)
}

// HasMaxPriceVariation returns true if MaxPriceVariation is present, Tag 1143.
func (m MarketDefinitionUpdateReport) HasMaxPriceVariation() bool {
	return m.Has(tag.MaxPriceVariation)
}

// HasImpliedMarketIndicator returns true if ImpliedMarketIndicator is present, Tag 1144.
func (m MarketDefinitionUpdateReport) HasImpliedMarketIndicator() bool {
	return m.Has(tag.ImpliedMarketIndicator)
}

// HasLowLimitPrice returns true if LowLimitPrice is present, Tag 1148.
func (m MarketDefinitionUpdateReport) HasLowLimitPrice() bool {
	return m.Has(tag.LowLimitPrice)
}

// HasHighLimitPrice returns true if HighLimitPrice is present, Tag 1149.
func (m MarketDefinitionUpdateReport) HasHighLimitPrice() bool {
	return m.Has(tag.HighLimitPrice)
}

// HasTradingReferencePrice returns true if TradingReferencePrice is present, Tag 1150.
func (m MarketDefinitionUpdateReport) HasTradingReferencePrice() bool {
	return m.Has(tag.TradingReferencePrice)
}

// HasApplID returns true if ApplID is present, Tag 1180.
func (m MarketDefinitionUpdateReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181.
func (m MarketDefinitionUpdateReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasNoTickRules returns true if NoTickRules is present, Tag 1205.
func (m MarketDefinitionUpdateReport) HasNoTickRules() bool {
	return m.Has(tag.NoTickRules)
}

// HasNoExecInstRules returns true if NoExecInstRules is present, Tag 1232.
func (m MarketDefinitionUpdateReport) HasNoExecInstRules() bool {
	return m.Has(tag.NoExecInstRules)
}

// HasNoLotTypeRules returns true if NoLotTypeRules is present, Tag 1234.
func (m MarketDefinitionUpdateReport) HasNoLotTypeRules() bool {
	return m.Has(tag.NoLotTypeRules)
}

// HasNoMatchRules returns true if NoMatchRules is present, Tag 1235.
func (m MarketDefinitionUpdateReport) HasNoMatchRules() bool {
	return m.Has(tag.NoMatchRules)
}

// HasNoOrdTypeRules returns true if NoOrdTypeRules is present, Tag 1237.
func (m MarketDefinitionUpdateReport) HasNoOrdTypeRules() bool {
	return m.Has(tag.NoOrdTypeRules)
}

// HasNoTimeInForceRules returns true if NoTimeInForceRules is present, Tag 1239.
func (m MarketDefinitionUpdateReport) HasNoTimeInForceRules() bool {
	return m.Has(tag.NoTimeInForceRules)
}

// HasTradingCurrency returns true if TradingCurrency is present, Tag 1245.
func (m MarketDefinitionUpdateReport) HasTradingCurrency() bool {
	return m.Has(tag.TradingCurrency)
}

// HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300.
func (m MarketDefinitionUpdateReport) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

// HasMarketID returns true if MarketID is present, Tag 1301.
func (m MarketDefinitionUpdateReport) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

// HasPriceLimitType returns true if PriceLimitType is present, Tag 1306.
func (m MarketDefinitionUpdateReport) HasPriceLimitType() bool {
	return m.Has(tag.PriceLimitType)
}

// HasParentMktSegmID returns true if ParentMktSegmID is present, Tag 1325.
func (m MarketDefinitionUpdateReport) HasParentMktSegmID() bool {
	return m.Has(tag.ParentMktSegmID)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350.
func (m MarketDefinitionUpdateReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352.
func (m MarketDefinitionUpdateReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// HasMultilegModel returns true if MultilegModel is present, Tag 1377.
func (m MarketDefinitionUpdateReport) HasMultilegModel() bool {
	return m.Has(tag.MultilegModel)
}

// HasMultilegPriceMethod returns true if MultilegPriceMethod is present, Tag 1378.
func (m MarketDefinitionUpdateReport) HasMultilegPriceMethod() bool {
	return m.Has(tag.MultilegPriceMethod)
}

// HasMarketReqID returns true if MarketReqID is present, Tag 1393.
func (m MarketDefinitionUpdateReport) HasMarketReqID() bool {
	return m.Has(tag.MarketReqID)
}

// HasMarketReportID returns true if MarketReportID is present, Tag 1394.
func (m MarketDefinitionUpdateReport) HasMarketReportID() bool {
	return m.Has(tag.MarketReportID)
}

// HasMarketUpdateAction returns true if MarketUpdateAction is present, Tag 1395.
func (m MarketDefinitionUpdateReport) HasMarketUpdateAction() bool {
	return m.Has(tag.MarketUpdateAction)
}

// HasMarketSegmentDesc returns true if MarketSegmentDesc is present, Tag 1396.
func (m MarketDefinitionUpdateReport) HasMarketSegmentDesc() bool {
	return m.Has(tag.MarketSegmentDesc)
}

// HasEncodedMktSegmDescLen returns true if EncodedMktSegmDescLen is present, Tag 1397.
func (m MarketDefinitionUpdateReport) HasEncodedMktSegmDescLen() bool {
	return m.Has(tag.EncodedMktSegmDescLen)
}

// HasEncodedMktSegmDesc returns true if EncodedMktSegmDesc is present, Tag 1398.
func (m MarketDefinitionUpdateReport) HasEncodedMktSegmDesc() bool {
	return m.Has(tag.EncodedMktSegmDesc)
}

// HasNoInstrumentScopes returns true if NoInstrumentScopes is present, Tag 1656.
func (m MarketDefinitionUpdateReport) HasNoInstrumentScopes() bool {
	return m.Has(tag.NoInstrumentScopes)
}

// HasTradeVolType returns true if TradeVolType is present, Tag 1786.
func (m MarketDefinitionUpdateReport) HasTradeVolType() bool {
	return m.Has(tag.TradeVolType)
}

// HasEffectiveBusinessDate returns true if EffectiveBusinessDate is present, Tag 2400.
func (m MarketDefinitionUpdateReport) HasEffectiveBusinessDate() bool {
	return m.Has(tag.EffectiveBusinessDate)
}

// HasMarketSegmentStatus returns true if MarketSegmentStatus is present, Tag 2542.
func (m MarketDefinitionUpdateReport) HasMarketSegmentStatus() bool {
	return m.Has(tag.MarketSegmentStatus)
}

// HasMarketSegmentType returns true if MarketSegmentType is present, Tag 2543.
func (m MarketDefinitionUpdateReport) HasMarketSegmentType() bool {
	return m.Has(tag.MarketSegmentType)
}

// HasMarketSegmentSubType returns true if MarketSegmentSubType is present, Tag 2544.
func (m MarketDefinitionUpdateReport) HasMarketSegmentSubType() bool {
	return m.Has(tag.MarketSegmentSubType)
}

// HasNoRelatedMarketSegments returns true if NoRelatedMarketSegments is present, Tag 2545.
func (m MarketDefinitionUpdateReport) HasNoRelatedMarketSegments() bool {
	return m.Has(tag.NoRelatedMarketSegments)
}

// HasNoAuctionTypeRules returns true if NoAuctionTypeRules is present, Tag 2548.
func (m MarketDefinitionUpdateReport) HasNoAuctionTypeRules() bool {
	return m.Has(tag.NoAuctionTypeRules)
}

// HasNoPriceRangeRules returns true if NoPriceRangeRules is present, Tag 2550.
func (m MarketDefinitionUpdateReport) HasNoPriceRangeRules() bool {
	return m.Has(tag.NoPriceRangeRules)
}

// HasFastMarketPercentage returns true if FastMarketPercentage is present, Tag 2557.
func (m MarketDefinitionUpdateReport) HasFastMarketPercentage() bool {
	return m.Has(tag.FastMarketPercentage)
}

// HasNoQuoteSizeRules returns true if NoQuoteSizeRules is present, Tag 2558.
func (m MarketDefinitionUpdateReport) HasNoQuoteSizeRules() bool {
	return m.Has(tag.NoQuoteSizeRules)
}

// HasQuoteSideIndicator returns true if QuoteSideIndicator is present, Tag 2559.
func (m MarketDefinitionUpdateReport) HasQuoteSideIndicator() bool {
	return m.Has(tag.QuoteSideIndicator)
}

// HasNoFlexProductEligibilities returns true if NoFlexProductEligibilities is present, Tag 2560.
func (m MarketDefinitionUpdateReport) HasNoFlexProductEligibilities() bool {
	return m.Has(tag.NoFlexProductEligibilities)
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

// NoMDFeedTypes is a repeating group element, Tag 1141.
type NoMDFeedTypes struct {
	*quickfix.Group
}

// SetMDFeedType sets MDFeedType, Tag 1022.
func (m NoMDFeedTypes) SetMDFeedType(v string) {
	m.Set(field.NewMDFeedType(v))
}

// SetMarketDepth sets MarketDepth, Tag 264.
func (m NoMDFeedTypes) SetMarketDepth(v int) {
	m.Set(field.NewMarketDepth(v))
}

// SetMDBookType sets MDBookType, Tag 1021.
func (m NoMDFeedTypes) SetMDBookType(v enum.MDBookType) {
	m.Set(field.NewMDBookType(v))
}

// SetMDSubFeedType sets MDSubFeedType, Tag 1683.
func (m NoMDFeedTypes) SetMDSubFeedType(v string) {
	m.Set(field.NewMDSubFeedType(v))
}

// SetMarketDepthTimeInterval sets MarketDepthTimeInterval, Tag 2563.
func (m NoMDFeedTypes) SetMarketDepthTimeInterval(v int) {
	m.Set(field.NewMarketDepthTimeInterval(v))
}

// SetMarketDepthTimeIntervalUnit sets MarketDepthTimeIntervalUnit, Tag 2564.
func (m NoMDFeedTypes) SetMarketDepthTimeIntervalUnit(v int) {
	m.Set(field.NewMarketDepthTimeIntervalUnit(v))
}

// SetMDRecoveryTimeInterval sets MDRecoveryTimeInterval, Tag 2565.
func (m NoMDFeedTypes) SetMDRecoveryTimeInterval(v int) {
	m.Set(field.NewMDRecoveryTimeInterval(v))
}

// SetMDRecoveryTimeIntervalUnit sets MDRecoveryTimeIntervalUnit, Tag 2566.
func (m NoMDFeedTypes) SetMDRecoveryTimeIntervalUnit(v int) {
	m.Set(field.NewMDRecoveryTimeIntervalUnit(v))
}

// SetMDSubBookType sets MDSubBookType, Tag 1173.
func (m NoMDFeedTypes) SetMDSubBookType(v int) {
	m.Set(field.NewMDSubBookType(v))
}

// SetPrimaryServiceLocationID sets PrimaryServiceLocationID, Tag 2567.
func (m NoMDFeedTypes) SetPrimaryServiceLocationID(v string) {
	m.Set(field.NewPrimaryServiceLocationID(v))
}

// SetSecondaryServiceLocationID sets SecondaryServiceLocationID, Tag 2568.
func (m NoMDFeedTypes) SetSecondaryServiceLocationID(v string) {
	m.Set(field.NewSecondaryServiceLocationID(v))
}

// GetMDFeedType gets MDFeedType, Tag 1022.
func (m NoMDFeedTypes) GetMDFeedType() (v string, err quickfix.MessageRejectError) {
	var f field.MDFeedTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketDepth gets MarketDepth, Tag 264.
func (m NoMDFeedTypes) GetMarketDepth() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDBookType gets MDBookType, Tag 1021.
func (m NoMDFeedTypes) GetMDBookType() (v enum.MDBookType, err quickfix.MessageRejectError) {
	var f field.MDBookTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDSubFeedType gets MDSubFeedType, Tag 1683.
func (m NoMDFeedTypes) GetMDSubFeedType() (v string, err quickfix.MessageRejectError) {
	var f field.MDSubFeedTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketDepthTimeInterval gets MarketDepthTimeInterval, Tag 2563.
func (m NoMDFeedTypes) GetMarketDepthTimeInterval() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthTimeIntervalField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketDepthTimeIntervalUnit gets MarketDepthTimeIntervalUnit, Tag 2564.
func (m NoMDFeedTypes) GetMarketDepthTimeIntervalUnit() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthTimeIntervalUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDRecoveryTimeInterval gets MDRecoveryTimeInterval, Tag 2565.
func (m NoMDFeedTypes) GetMDRecoveryTimeInterval() (v int, err quickfix.MessageRejectError) {
	var f field.MDRecoveryTimeIntervalField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDRecoveryTimeIntervalUnit gets MDRecoveryTimeIntervalUnit, Tag 2566.
func (m NoMDFeedTypes) GetMDRecoveryTimeIntervalUnit() (v int, err quickfix.MessageRejectError) {
	var f field.MDRecoveryTimeIntervalUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDSubBookType gets MDSubBookType, Tag 1173.
func (m NoMDFeedTypes) GetMDSubBookType() (v int, err quickfix.MessageRejectError) {
	var f field.MDSubBookTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrimaryServiceLocationID gets PrimaryServiceLocationID, Tag 2567.
func (m NoMDFeedTypes) GetPrimaryServiceLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.PrimaryServiceLocationIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryServiceLocationID gets SecondaryServiceLocationID, Tag 2568.
func (m NoMDFeedTypes) GetSecondaryServiceLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryServiceLocationIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMDFeedType returns true if MDFeedType is present, Tag 1022.
func (m NoMDFeedTypes) HasMDFeedType() bool {
	return m.Has(tag.MDFeedType)
}

// HasMarketDepth returns true if MarketDepth is present, Tag 264.
func (m NoMDFeedTypes) HasMarketDepth() bool {
	return m.Has(tag.MarketDepth)
}

// HasMDBookType returns true if MDBookType is present, Tag 1021.
func (m NoMDFeedTypes) HasMDBookType() bool {
	return m.Has(tag.MDBookType)
}

// HasMDSubFeedType returns true if MDSubFeedType is present, Tag 1683.
func (m NoMDFeedTypes) HasMDSubFeedType() bool {
	return m.Has(tag.MDSubFeedType)
}

// HasMarketDepthTimeInterval returns true if MarketDepthTimeInterval is present, Tag 2563.
func (m NoMDFeedTypes) HasMarketDepthTimeInterval() bool {
	return m.Has(tag.MarketDepthTimeInterval)
}

// HasMarketDepthTimeIntervalUnit returns true if MarketDepthTimeIntervalUnit is present, Tag 2564.
func (m NoMDFeedTypes) HasMarketDepthTimeIntervalUnit() bool {
	return m.Has(tag.MarketDepthTimeIntervalUnit)
}

// HasMDRecoveryTimeInterval returns true if MDRecoveryTimeInterval is present, Tag 2565.
func (m NoMDFeedTypes) HasMDRecoveryTimeInterval() bool {
	return m.Has(tag.MDRecoveryTimeInterval)
}

// HasMDRecoveryTimeIntervalUnit returns true if MDRecoveryTimeIntervalUnit is present, Tag 2566.
func (m NoMDFeedTypes) HasMDRecoveryTimeIntervalUnit() bool {
	return m.Has(tag.MDRecoveryTimeIntervalUnit)
}

// HasMDSubBookType returns true if MDSubBookType is present, Tag 1173.
func (m NoMDFeedTypes) HasMDSubBookType() bool {
	return m.Has(tag.MDSubBookType)
}

// HasPrimaryServiceLocationID returns true if PrimaryServiceLocationID is present, Tag 2567.
func (m NoMDFeedTypes) HasPrimaryServiceLocationID() bool {
	return m.Has(tag.PrimaryServiceLocationID)
}

// HasSecondaryServiceLocationID returns true if SecondaryServiceLocationID is present, Tag 2568.
func (m NoMDFeedTypes) HasSecondaryServiceLocationID() bool {
	return m.Has(tag.SecondaryServiceLocationID)
}

// NoMDFeedTypesRepeatingGroup is a repeating group, Tag 1141.
type NoMDFeedTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMDFeedTypesRepeatingGroup returns an initialized, NoMDFeedTypesRepeatingGroup.
func NewNoMDFeedTypesRepeatingGroup() NoMDFeedTypesRepeatingGroup {
	return NoMDFeedTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMDFeedTypes,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MDFeedType),
				quickfix.GroupElement(tag.MarketDepth),
				quickfix.GroupElement(tag.MDBookType),
				quickfix.GroupElement(tag.MDSubFeedType),
				quickfix.GroupElement(tag.MarketDepthTimeInterval),
				quickfix.GroupElement(tag.MarketDepthTimeIntervalUnit),
				quickfix.GroupElement(tag.MDRecoveryTimeInterval),
				quickfix.GroupElement(tag.MDRecoveryTimeIntervalUnit),
				quickfix.GroupElement(tag.MDSubBookType),
				quickfix.GroupElement(tag.PrimaryServiceLocationID),
				quickfix.GroupElement(tag.SecondaryServiceLocationID),
			},
		),
	}
}

// Add create and append a new NoMDFeedTypes to this group.
func (m NoMDFeedTypesRepeatingGroup) Add() NoMDFeedTypes {
	g := m.RepeatingGroup.Add()
	return NoMDFeedTypes{g}
}

// Get returns the ith NoMDFeedTypes in the NoMDFeedTypesRepeatinGroup.
func (m NoMDFeedTypesRepeatingGroup) Get(i int) NoMDFeedTypes {
	return NoMDFeedTypes{m.RepeatingGroup.Get(i)}
}

// NoTickRules is a repeating group element, Tag 1205.
type NoTickRules struct {
	*quickfix.Group
}

// SetStartTickPriceRange sets StartTickPriceRange, Tag 1206.
func (m NoTickRules) SetStartTickPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartTickPriceRange(value, scale))
}

// SetEndTickPriceRange sets EndTickPriceRange, Tag 1207.
func (m NoTickRules) SetEndTickPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndTickPriceRange(value, scale))
}

// SetTickIncrement sets TickIncrement, Tag 1208.
func (m NoTickRules) SetTickIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewTickIncrement(value, scale))
}

// SetTickRuleType sets TickRuleType, Tag 1209.
func (m NoTickRules) SetTickRuleType(v enum.TickRuleType) {
	m.Set(field.NewTickRuleType(v))
}

// SetSettlPriceIncrement sets SettlPriceIncrement, Tag 1830.
func (m NoTickRules) SetSettlPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlPriceIncrement(value, scale))
}

// SetSettlPriceSecondaryIncrement sets SettlPriceSecondaryIncrement, Tag 1831.
func (m NoTickRules) SetSettlPriceSecondaryIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlPriceSecondaryIncrement(value, scale))
}

// SetTickRuleProductComplex sets TickRuleProductComplex, Tag 2571.
func (m NoTickRules) SetTickRuleProductComplex(v string) {
	m.Set(field.NewTickRuleProductComplex(v))
}

// GetStartTickPriceRange gets StartTickPriceRange, Tag 1206.
func (m NoTickRules) GetStartTickPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StartTickPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndTickPriceRange gets EndTickPriceRange, Tag 1207.
func (m NoTickRules) GetEndTickPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EndTickPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTickIncrement gets TickIncrement, Tag 1208.
func (m NoTickRules) GetTickIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TickIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTickRuleType gets TickRuleType, Tag 1209.
func (m NoTickRules) GetTickRuleType() (v enum.TickRuleType, err quickfix.MessageRejectError) {
	var f field.TickRuleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlPriceIncrement gets SettlPriceIncrement, Tag 1830.
func (m NoTickRules) GetSettlPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlPriceSecondaryIncrement gets SettlPriceSecondaryIncrement, Tag 1831.
func (m NoTickRules) GetSettlPriceSecondaryIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlPriceSecondaryIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTickRuleProductComplex gets TickRuleProductComplex, Tag 2571.
func (m NoTickRules) GetTickRuleProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.TickRuleProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasStartTickPriceRange returns true if StartTickPriceRange is present, Tag 1206.
func (m NoTickRules) HasStartTickPriceRange() bool {
	return m.Has(tag.StartTickPriceRange)
}

// HasEndTickPriceRange returns true if EndTickPriceRange is present, Tag 1207.
func (m NoTickRules) HasEndTickPriceRange() bool {
	return m.Has(tag.EndTickPriceRange)
}

// HasTickIncrement returns true if TickIncrement is present, Tag 1208.
func (m NoTickRules) HasTickIncrement() bool {
	return m.Has(tag.TickIncrement)
}

// HasTickRuleType returns true if TickRuleType is present, Tag 1209.
func (m NoTickRules) HasTickRuleType() bool {
	return m.Has(tag.TickRuleType)
}

// HasSettlPriceIncrement returns true if SettlPriceIncrement is present, Tag 1830.
func (m NoTickRules) HasSettlPriceIncrement() bool {
	return m.Has(tag.SettlPriceIncrement)
}

// HasSettlPriceSecondaryIncrement returns true if SettlPriceSecondaryIncrement is present, Tag 1831.
func (m NoTickRules) HasSettlPriceSecondaryIncrement() bool {
	return m.Has(tag.SettlPriceSecondaryIncrement)
}

// HasTickRuleProductComplex returns true if TickRuleProductComplex is present, Tag 2571.
func (m NoTickRules) HasTickRuleProductComplex() bool {
	return m.Has(tag.TickRuleProductComplex)
}

// NoTickRulesRepeatingGroup is a repeating group, Tag 1205.
type NoTickRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTickRulesRepeatingGroup returns an initialized, NoTickRulesRepeatingGroup.
func NewNoTickRulesRepeatingGroup() NoTickRulesRepeatingGroup {
	return NoTickRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoTickRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.StartTickPriceRange),
				quickfix.GroupElement(tag.EndTickPriceRange),
				quickfix.GroupElement(tag.TickIncrement),
				quickfix.GroupElement(tag.TickRuleType),
				quickfix.GroupElement(tag.SettlPriceIncrement),
				quickfix.GroupElement(tag.SettlPriceSecondaryIncrement),
				quickfix.GroupElement(tag.TickRuleProductComplex),
			},
		),
	}
}

// Add create and append a new NoTickRules to this group.
func (m NoTickRulesRepeatingGroup) Add() NoTickRules {
	g := m.RepeatingGroup.Add()
	return NoTickRules{g}
}

// Get returns the ith NoTickRules in the NoTickRulesRepeatinGroup.
func (m NoTickRulesRepeatingGroup) Get(i int) NoTickRules {
	return NoTickRules{m.RepeatingGroup.Get(i)}
}

// NoExecInstRules is a repeating group element, Tag 1232.
type NoExecInstRules struct {
	*quickfix.Group
}

// SetExecInstValue sets ExecInstValue, Tag 1308.
func (m NoExecInstRules) SetExecInstValue(v string) {
	m.Set(field.NewExecInstValue(v))
}

// GetExecInstValue gets ExecInstValue, Tag 1308.
func (m NoExecInstRules) GetExecInstValue() (v string, err quickfix.MessageRejectError) {
	var f field.ExecInstValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasExecInstValue returns true if ExecInstValue is present, Tag 1308.
func (m NoExecInstRules) HasExecInstValue() bool {
	return m.Has(tag.ExecInstValue)
}

// NoExecInstRulesRepeatingGroup is a repeating group, Tag 1232.
type NoExecInstRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoExecInstRulesRepeatingGroup returns an initialized, NoExecInstRulesRepeatingGroup.
func NewNoExecInstRulesRepeatingGroup() NoExecInstRulesRepeatingGroup {
	return NoExecInstRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoExecInstRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.ExecInstValue),
			},
		),
	}
}

// Add create and append a new NoExecInstRules to this group.
func (m NoExecInstRulesRepeatingGroup) Add() NoExecInstRules {
	g := m.RepeatingGroup.Add()
	return NoExecInstRules{g}
}

// Get returns the ith NoExecInstRules in the NoExecInstRulesRepeatinGroup.
func (m NoExecInstRulesRepeatingGroup) Get(i int) NoExecInstRules {
	return NoExecInstRules{m.RepeatingGroup.Get(i)}
}

// NoLotTypeRules is a repeating group element, Tag 1234.
type NoLotTypeRules struct {
	*quickfix.Group
}

// SetLotType sets LotType, Tag 1093.
func (m NoLotTypeRules) SetLotType(v enum.LotType) {
	m.Set(field.NewLotType(v))
}

// SetMinLotSize sets MinLotSize, Tag 1231.
func (m NoLotTypeRules) SetMinLotSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinLotSize(value, scale))
}

// GetLotType gets LotType, Tag 1093.
func (m NoLotTypeRules) GetLotType() (v enum.LotType, err quickfix.MessageRejectError) {
	var f field.LotTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinLotSize gets MinLotSize, Tag 1231.
func (m NoLotTypeRules) GetMinLotSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinLotSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLotType returns true if LotType is present, Tag 1093.
func (m NoLotTypeRules) HasLotType() bool {
	return m.Has(tag.LotType)
}

// HasMinLotSize returns true if MinLotSize is present, Tag 1231.
func (m NoLotTypeRules) HasMinLotSize() bool {
	return m.Has(tag.MinLotSize)
}

// NoLotTypeRulesRepeatingGroup is a repeating group, Tag 1234.
type NoLotTypeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLotTypeRulesRepeatingGroup returns an initialized, NoLotTypeRulesRepeatingGroup.
func NewNoLotTypeRulesRepeatingGroup() NoLotTypeRulesRepeatingGroup {
	return NoLotTypeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoLotTypeRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.LotType),
				quickfix.GroupElement(tag.MinLotSize),
			},
		),
	}
}

// Add create and append a new NoLotTypeRules to this group.
func (m NoLotTypeRulesRepeatingGroup) Add() NoLotTypeRules {
	g := m.RepeatingGroup.Add()
	return NoLotTypeRules{g}
}

// Get returns the ith NoLotTypeRules in the NoLotTypeRulesRepeatinGroup.
func (m NoLotTypeRulesRepeatingGroup) Get(i int) NoLotTypeRules {
	return NoLotTypeRules{m.RepeatingGroup.Get(i)}
}

// NoMatchRules is a repeating group element, Tag 1235.
type NoMatchRules struct {
	*quickfix.Group
}

// SetMatchAlgorithm sets MatchAlgorithm, Tag 1142.
func (m NoMatchRules) SetMatchAlgorithm(v string) {
	m.Set(field.NewMatchAlgorithm(v))
}

// SetMatchType sets MatchType, Tag 574.
func (m NoMatchRules) SetMatchType(v enum.MatchType) {
	m.Set(field.NewMatchType(v))
}

// SetMatchRuleProductComplex sets MatchRuleProductComplex, Tag 2569.
func (m NoMatchRules) SetMatchRuleProductComplex(v string) {
	m.Set(field.NewMatchRuleProductComplex(v))
}

// SetCustomerPriority sets CustomerPriority, Tag 2570.
func (m NoMatchRules) SetCustomerPriority(v enum.CustomerPriority) {
	m.Set(field.NewCustomerPriority(v))
}

// GetMatchAlgorithm gets MatchAlgorithm, Tag 1142.
func (m NoMatchRules) GetMatchAlgorithm() (v string, err quickfix.MessageRejectError) {
	var f field.MatchAlgorithmField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchType gets MatchType, Tag 574.
func (m NoMatchRules) GetMatchType() (v enum.MatchType, err quickfix.MessageRejectError) {
	var f field.MatchTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchRuleProductComplex gets MatchRuleProductComplex, Tag 2569.
func (m NoMatchRules) GetMatchRuleProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.MatchRuleProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCustomerPriority gets CustomerPriority, Tag 2570.
func (m NoMatchRules) GetCustomerPriority() (v enum.CustomerPriority, err quickfix.MessageRejectError) {
	var f field.CustomerPriorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMatchAlgorithm returns true if MatchAlgorithm is present, Tag 1142.
func (m NoMatchRules) HasMatchAlgorithm() bool {
	return m.Has(tag.MatchAlgorithm)
}

// HasMatchType returns true if MatchType is present, Tag 574.
func (m NoMatchRules) HasMatchType() bool {
	return m.Has(tag.MatchType)
}

// HasMatchRuleProductComplex returns true if MatchRuleProductComplex is present, Tag 2569.
func (m NoMatchRules) HasMatchRuleProductComplex() bool {
	return m.Has(tag.MatchRuleProductComplex)
}

// HasCustomerPriority returns true if CustomerPriority is present, Tag 2570.
func (m NoMatchRules) HasCustomerPriority() bool {
	return m.Has(tag.CustomerPriority)
}

// NoMatchRulesRepeatingGroup is a repeating group, Tag 1235.
type NoMatchRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMatchRulesRepeatingGroup returns an initialized, NoMatchRulesRepeatingGroup.
func NewNoMatchRulesRepeatingGroup() NoMatchRulesRepeatingGroup {
	return NoMatchRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoMatchRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MatchAlgorithm),
				quickfix.GroupElement(tag.MatchType),
				quickfix.GroupElement(tag.MatchRuleProductComplex),
				quickfix.GroupElement(tag.CustomerPriority),
			},
		),
	}
}

// Add create and append a new NoMatchRules to this group.
func (m NoMatchRulesRepeatingGroup) Add() NoMatchRules {
	g := m.RepeatingGroup.Add()
	return NoMatchRules{g}
}

// Get returns the ith NoMatchRules in the NoMatchRulesRepeatinGroup.
func (m NoMatchRulesRepeatingGroup) Get(i int) NoMatchRules {
	return NoMatchRules{m.RepeatingGroup.Get(i)}
}

// NoOrdTypeRules is a repeating group element, Tag 1237.
type NoOrdTypeRules struct {
	*quickfix.Group
}

// SetOrdType sets OrdType, Tag 40.
func (m NoOrdTypeRules) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

// GetOrdType gets OrdType, Tag 40.
func (m NoOrdTypeRules) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasOrdType returns true if OrdType is present, Tag 40.
func (m NoOrdTypeRules) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

// NoOrdTypeRulesRepeatingGroup is a repeating group, Tag 1237.
type NoOrdTypeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoOrdTypeRulesRepeatingGroup returns an initialized, NoOrdTypeRulesRepeatingGroup.
func NewNoOrdTypeRulesRepeatingGroup() NoOrdTypeRulesRepeatingGroup {
	return NoOrdTypeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoOrdTypeRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.OrdType),
			},
		),
	}
}

// Add create and append a new NoOrdTypeRules to this group.
func (m NoOrdTypeRulesRepeatingGroup) Add() NoOrdTypeRules {
	g := m.RepeatingGroup.Add()
	return NoOrdTypeRules{g}
}

// Get returns the ith NoOrdTypeRules in the NoOrdTypeRulesRepeatinGroup.
func (m NoOrdTypeRulesRepeatingGroup) Get(i int) NoOrdTypeRules {
	return NoOrdTypeRules{m.RepeatingGroup.Get(i)}
}

// NoTimeInForceRules is a repeating group element, Tag 1239.
type NoTimeInForceRules struct {
	*quickfix.Group
}

// SetTimeInForce sets TimeInForce, Tag 59.
func (m NoTimeInForceRules) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

// GetTimeInForce gets TimeInForce, Tag 59.
func (m NoTimeInForceRules) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTimeInForce returns true if TimeInForce is present, Tag 59.
func (m NoTimeInForceRules) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

// NoTimeInForceRulesRepeatingGroup is a repeating group, Tag 1239.
type NoTimeInForceRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTimeInForceRulesRepeatingGroup returns an initialized, NoTimeInForceRulesRepeatingGroup.
func NewNoTimeInForceRulesRepeatingGroup() NoTimeInForceRulesRepeatingGroup {
	return NoTimeInForceRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoTimeInForceRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.TimeInForce),
			},
		),
	}
}

// Add create and append a new NoTimeInForceRules to this group.
func (m NoTimeInForceRulesRepeatingGroup) Add() NoTimeInForceRules {
	g := m.RepeatingGroup.Add()
	return NoTimeInForceRules{g}
}

// Get returns the ith NoTimeInForceRules in the NoTimeInForceRulesRepeatinGroup.
func (m NoTimeInForceRulesRepeatingGroup) Get(i int) NoTimeInForceRules {
	return NoTimeInForceRules{m.RepeatingGroup.Get(i)}
}

// NoInstrumentScopes is a repeating group element, Tag 1656.
type NoInstrumentScopes struct {
	*quickfix.Group
}

// SetInstrumentScopeOperator sets InstrumentScopeOperator, Tag 1535.
func (m NoInstrumentScopes) SetInstrumentScopeOperator(v enum.InstrumentScopeOperator) {
	m.Set(field.NewInstrumentScopeOperator(v))
}

// SetInstrumentScopeSymbol sets InstrumentScopeSymbol, Tag 1536.
func (m NoInstrumentScopes) SetInstrumentScopeSymbol(v string) {
	m.Set(field.NewInstrumentScopeSymbol(v))
}

// SetInstrumentScopeSymbolSfx sets InstrumentScopeSymbolSfx, Tag 1537.
func (m NoInstrumentScopes) SetInstrumentScopeSymbolSfx(v string) {
	m.Set(field.NewInstrumentScopeSymbolSfx(v))
}

// SetInstrumentScopeSecurityID sets InstrumentScopeSecurityID, Tag 1538.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityID(v string) {
	m.Set(field.NewInstrumentScopeSecurityID(v))
}

// SetInstrumentScopeSecurityIDSource sets InstrumentScopeSecurityIDSource, Tag 1539.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityIDSource(v string) {
	m.Set(field.NewInstrumentScopeSecurityIDSource(v))
}

// SetNoInstrumentScopeSecurityAltID sets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m NoInstrumentScopes) SetNoInstrumentScopeSecurityAltID(f NoInstrumentScopeSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrumentScopeProduct sets InstrumentScopeProduct, Tag 1543.
func (m NoInstrumentScopes) SetInstrumentScopeProduct(v int) {
	m.Set(field.NewInstrumentScopeProduct(v))
}

// SetInstrumentScopeProductComplex sets InstrumentScopeProductComplex, Tag 1544.
func (m NoInstrumentScopes) SetInstrumentScopeProductComplex(v string) {
	m.Set(field.NewInstrumentScopeProductComplex(v))
}

// SetInstrumentScopeSecurityGroup sets InstrumentScopeSecurityGroup, Tag 1545.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityGroup(v string) {
	m.Set(field.NewInstrumentScopeSecurityGroup(v))
}

// SetInstrumentScopeCFICode sets InstrumentScopeCFICode, Tag 1546.
func (m NoInstrumentScopes) SetInstrumentScopeCFICode(v string) {
	m.Set(field.NewInstrumentScopeCFICode(v))
}

// SetInstrumentScopeSecurityType sets InstrumentScopeSecurityType, Tag 1547.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityType(v string) {
	m.Set(field.NewInstrumentScopeSecurityType(v))
}

// SetInstrumentScopeSecuritySubType sets InstrumentScopeSecuritySubType, Tag 1548.
func (m NoInstrumentScopes) SetInstrumentScopeSecuritySubType(v string) {
	m.Set(field.NewInstrumentScopeSecuritySubType(v))
}

// SetInstrumentScopeMaturityMonthYear sets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m NoInstrumentScopes) SetInstrumentScopeMaturityMonthYear(v string) {
	m.Set(field.NewInstrumentScopeMaturityMonthYear(v))
}

// SetInstrumentScopeMaturityTime sets InstrumentScopeMaturityTime, Tag 1550.
func (m NoInstrumentScopes) SetInstrumentScopeMaturityTime(v string) {
	m.Set(field.NewInstrumentScopeMaturityTime(v))
}

// SetInstrumentScopeRestructuringType sets InstrumentScopeRestructuringType, Tag 1551.
func (m NoInstrumentScopes) SetInstrumentScopeRestructuringType(v string) {
	m.Set(field.NewInstrumentScopeRestructuringType(v))
}

// SetInstrumentScopeSeniority sets InstrumentScopeSeniority, Tag 1552.
func (m NoInstrumentScopes) SetInstrumentScopeSeniority(v string) {
	m.Set(field.NewInstrumentScopeSeniority(v))
}

// SetInstrumentScopePutOrCall sets InstrumentScopePutOrCall, Tag 1553.
func (m NoInstrumentScopes) SetInstrumentScopePutOrCall(v int) {
	m.Set(field.NewInstrumentScopePutOrCall(v))
}

// SetInstrumentScopeFlexibleIndicator sets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m NoInstrumentScopes) SetInstrumentScopeFlexibleIndicator(v bool) {
	m.Set(field.NewInstrumentScopeFlexibleIndicator(v))
}

// SetInstrumentScopeCouponRate sets InstrumentScopeCouponRate, Tag 1555.
func (m NoInstrumentScopes) SetInstrumentScopeCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewInstrumentScopeCouponRate(value, scale))
}

// SetInstrumentScopeSecurityExchange sets InstrumentScopeSecurityExchange, Tag 1616.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityExchange(v string) {
	m.Set(field.NewInstrumentScopeSecurityExchange(v))
}

// SetInstrumentScopeSecurityDesc sets InstrumentScopeSecurityDesc, Tag 1556.
func (m NoInstrumentScopes) SetInstrumentScopeSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeSecurityDesc(v))
}

// SetInstrumentScopeEncodedSecurityDescLen sets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m NoInstrumentScopes) SetInstrumentScopeEncodedSecurityDescLen(v int) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDescLen(v))
}

// SetInstrumentScopeEncodedSecurityDesc sets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m NoInstrumentScopes) SetInstrumentScopeEncodedSecurityDesc(v string) {
	m.Set(field.NewInstrumentScopeEncodedSecurityDesc(v))
}

// SetInstrumentScopeSettlType sets InstrumentScopeSettlType, Tag 1557.
func (m NoInstrumentScopes) SetInstrumentScopeSettlType(v string) {
	m.Set(field.NewInstrumentScopeSettlType(v))
}

// SetInstrumentScopeUPICode sets InstrumentScopeUPICode, Tag 2895.
func (m NoInstrumentScopes) SetInstrumentScopeUPICode(v string) {
	m.Set(field.NewInstrumentScopeUPICode(v))
}

// GetInstrumentScopeOperator gets InstrumentScopeOperator, Tag 1535.
func (m NoInstrumentScopes) GetInstrumentScopeOperator() (v enum.InstrumentScopeOperator, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeOperatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbol gets InstrumentScopeSymbol, Tag 1536.
func (m NoInstrumentScopes) GetInstrumentScopeSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSymbolSfx gets InstrumentScopeSymbolSfx, Tag 1537.
func (m NoInstrumentScopes) GetInstrumentScopeSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityID gets InstrumentScopeSecurityID, Tag 1538.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityIDSource gets InstrumentScopeSecurityIDSource, Tag 1539.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentScopeSecurityAltID gets NoInstrumentScopeSecurityAltID, Tag 1540.
func (m NoInstrumentScopes) GetNoInstrumentScopeSecurityAltID() (NoInstrumentScopeSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentScopeSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrumentScopeProduct gets InstrumentScopeProduct, Tag 1543.
func (m NoInstrumentScopes) GetInstrumentScopeProduct() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeProductComplex gets InstrumentScopeProductComplex, Tag 1544.
func (m NoInstrumentScopes) GetInstrumentScopeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityGroup gets InstrumentScopeSecurityGroup, Tag 1545.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCFICode gets InstrumentScopeCFICode, Tag 1546.
func (m NoInstrumentScopes) GetInstrumentScopeCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityType gets InstrumentScopeSecurityType, Tag 1547.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecuritySubType gets InstrumentScopeSecuritySubType, Tag 1548.
func (m NoInstrumentScopes) GetInstrumentScopeSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityMonthYear gets InstrumentScopeMaturityMonthYear, Tag 1549.
func (m NoInstrumentScopes) GetInstrumentScopeMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeMaturityTime gets InstrumentScopeMaturityTime, Tag 1550.
func (m NoInstrumentScopes) GetInstrumentScopeMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeRestructuringType gets InstrumentScopeRestructuringType, Tag 1551.
func (m NoInstrumentScopes) GetInstrumentScopeRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSeniority gets InstrumentScopeSeniority, Tag 1552.
func (m NoInstrumentScopes) GetInstrumentScopeSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopePutOrCall gets InstrumentScopePutOrCall, Tag 1553.
func (m NoInstrumentScopes) GetInstrumentScopePutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopePutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeFlexibleIndicator gets InstrumentScopeFlexibleIndicator, Tag 1554.
func (m NoInstrumentScopes) GetInstrumentScopeFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeFlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeCouponRate gets InstrumentScopeCouponRate, Tag 1555.
func (m NoInstrumentScopes) GetInstrumentScopeCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityExchange gets InstrumentScopeSecurityExchange, Tag 1616.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSecurityDesc gets InstrumentScopeSecurityDesc, Tag 1556.
func (m NoInstrumentScopes) GetInstrumentScopeSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDescLen gets InstrumentScopeEncodedSecurityDescLen, Tag 1620.
func (m NoInstrumentScopes) GetInstrumentScopeEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeEncodedSecurityDesc gets InstrumentScopeEncodedSecurityDesc, Tag 1621.
func (m NoInstrumentScopes) GetInstrumentScopeEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeSettlType gets InstrumentScopeSettlType, Tag 1557.
func (m NoInstrumentScopes) GetInstrumentScopeSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentScopeUPICode gets InstrumentScopeUPICode, Tag 2895.
func (m NoInstrumentScopes) GetInstrumentScopeUPICode() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentScopeUPICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentScopeOperator returns true if InstrumentScopeOperator is present, Tag 1535.
func (m NoInstrumentScopes) HasInstrumentScopeOperator() bool {
	return m.Has(tag.InstrumentScopeOperator)
}

// HasInstrumentScopeSymbol returns true if InstrumentScopeSymbol is present, Tag 1536.
func (m NoInstrumentScopes) HasInstrumentScopeSymbol() bool {
	return m.Has(tag.InstrumentScopeSymbol)
}

// HasInstrumentScopeSymbolSfx returns true if InstrumentScopeSymbolSfx is present, Tag 1537.
func (m NoInstrumentScopes) HasInstrumentScopeSymbolSfx() bool {
	return m.Has(tag.InstrumentScopeSymbolSfx)
}

// HasInstrumentScopeSecurityID returns true if InstrumentScopeSecurityID is present, Tag 1538.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityID() bool {
	return m.Has(tag.InstrumentScopeSecurityID)
}

// HasInstrumentScopeSecurityIDSource returns true if InstrumentScopeSecurityIDSource is present, Tag 1539.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityIDSource() bool {
	return m.Has(tag.InstrumentScopeSecurityIDSource)
}

// HasNoInstrumentScopeSecurityAltID returns true if NoInstrumentScopeSecurityAltID is present, Tag 1540.
func (m NoInstrumentScopes) HasNoInstrumentScopeSecurityAltID() bool {
	return m.Has(tag.NoInstrumentScopeSecurityAltID)
}

// HasInstrumentScopeProduct returns true if InstrumentScopeProduct is present, Tag 1543.
func (m NoInstrumentScopes) HasInstrumentScopeProduct() bool {
	return m.Has(tag.InstrumentScopeProduct)
}

// HasInstrumentScopeProductComplex returns true if InstrumentScopeProductComplex is present, Tag 1544.
func (m NoInstrumentScopes) HasInstrumentScopeProductComplex() bool {
	return m.Has(tag.InstrumentScopeProductComplex)
}

// HasInstrumentScopeSecurityGroup returns true if InstrumentScopeSecurityGroup is present, Tag 1545.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityGroup() bool {
	return m.Has(tag.InstrumentScopeSecurityGroup)
}

// HasInstrumentScopeCFICode returns true if InstrumentScopeCFICode is present, Tag 1546.
func (m NoInstrumentScopes) HasInstrumentScopeCFICode() bool {
	return m.Has(tag.InstrumentScopeCFICode)
}

// HasInstrumentScopeSecurityType returns true if InstrumentScopeSecurityType is present, Tag 1547.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityType() bool {
	return m.Has(tag.InstrumentScopeSecurityType)
}

// HasInstrumentScopeSecuritySubType returns true if InstrumentScopeSecuritySubType is present, Tag 1548.
func (m NoInstrumentScopes) HasInstrumentScopeSecuritySubType() bool {
	return m.Has(tag.InstrumentScopeSecuritySubType)
}

// HasInstrumentScopeMaturityMonthYear returns true if InstrumentScopeMaturityMonthYear is present, Tag 1549.
func (m NoInstrumentScopes) HasInstrumentScopeMaturityMonthYear() bool {
	return m.Has(tag.InstrumentScopeMaturityMonthYear)
}

// HasInstrumentScopeMaturityTime returns true if InstrumentScopeMaturityTime is present, Tag 1550.
func (m NoInstrumentScopes) HasInstrumentScopeMaturityTime() bool {
	return m.Has(tag.InstrumentScopeMaturityTime)
}

// HasInstrumentScopeRestructuringType returns true if InstrumentScopeRestructuringType is present, Tag 1551.
func (m NoInstrumentScopes) HasInstrumentScopeRestructuringType() bool {
	return m.Has(tag.InstrumentScopeRestructuringType)
}

// HasInstrumentScopeSeniority returns true if InstrumentScopeSeniority is present, Tag 1552.
func (m NoInstrumentScopes) HasInstrumentScopeSeniority() bool {
	return m.Has(tag.InstrumentScopeSeniority)
}

// HasInstrumentScopePutOrCall returns true if InstrumentScopePutOrCall is present, Tag 1553.
func (m NoInstrumentScopes) HasInstrumentScopePutOrCall() bool {
	return m.Has(tag.InstrumentScopePutOrCall)
}

// HasInstrumentScopeFlexibleIndicator returns true if InstrumentScopeFlexibleIndicator is present, Tag 1554.
func (m NoInstrumentScopes) HasInstrumentScopeFlexibleIndicator() bool {
	return m.Has(tag.InstrumentScopeFlexibleIndicator)
}

// HasInstrumentScopeCouponRate returns true if InstrumentScopeCouponRate is present, Tag 1555.
func (m NoInstrumentScopes) HasInstrumentScopeCouponRate() bool {
	return m.Has(tag.InstrumentScopeCouponRate)
}

// HasInstrumentScopeSecurityExchange returns true if InstrumentScopeSecurityExchange is present, Tag 1616.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityExchange() bool {
	return m.Has(tag.InstrumentScopeSecurityExchange)
}

// HasInstrumentScopeSecurityDesc returns true if InstrumentScopeSecurityDesc is present, Tag 1556.
func (m NoInstrumentScopes) HasInstrumentScopeSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeSecurityDesc)
}

// HasInstrumentScopeEncodedSecurityDescLen returns true if InstrumentScopeEncodedSecurityDescLen is present, Tag 1620.
func (m NoInstrumentScopes) HasInstrumentScopeEncodedSecurityDescLen() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDescLen)
}

// HasInstrumentScopeEncodedSecurityDesc returns true if InstrumentScopeEncodedSecurityDesc is present, Tag 1621.
func (m NoInstrumentScopes) HasInstrumentScopeEncodedSecurityDesc() bool {
	return m.Has(tag.InstrumentScopeEncodedSecurityDesc)
}

// HasInstrumentScopeSettlType returns true if InstrumentScopeSettlType is present, Tag 1557.
func (m NoInstrumentScopes) HasInstrumentScopeSettlType() bool {
	return m.Has(tag.InstrumentScopeSettlType)
}

// HasInstrumentScopeUPICode returns true if InstrumentScopeUPICode is present, Tag 2895.
func (m NoInstrumentScopes) HasInstrumentScopeUPICode() bool {
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

// NoInstrumentScopesRepeatingGroup is a repeating group, Tag 1656.
type NoInstrumentScopesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentScopesRepeatingGroup returns an initialized, NoInstrumentScopesRepeatingGroup.
func NewNoInstrumentScopesRepeatingGroup() NoInstrumentScopesRepeatingGroup {
	return NoInstrumentScopesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoInstrumentScopes,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.InstrumentScopeOperator),
				quickfix.GroupElement(tag.InstrumentScopeSymbol),
				quickfix.GroupElement(tag.InstrumentScopeSymbolSfx),
				quickfix.GroupElement(tag.InstrumentScopeSecurityID),
				quickfix.GroupElement(tag.InstrumentScopeSecurityIDSource),
				NewNoInstrumentScopeSecurityAltIDRepeatingGroup(),
				quickfix.GroupElement(tag.InstrumentScopeProduct),
				quickfix.GroupElement(tag.InstrumentScopeProductComplex),
				quickfix.GroupElement(tag.InstrumentScopeSecurityGroup),
				quickfix.GroupElement(tag.InstrumentScopeCFICode),
				quickfix.GroupElement(tag.InstrumentScopeSecurityType),
				quickfix.GroupElement(tag.InstrumentScopeSecuritySubType),
				quickfix.GroupElement(tag.InstrumentScopeMaturityMonthYear),
				quickfix.GroupElement(tag.InstrumentScopeMaturityTime),
				quickfix.GroupElement(tag.InstrumentScopeRestructuringType),
				quickfix.GroupElement(tag.InstrumentScopeSeniority),
				quickfix.GroupElement(tag.InstrumentScopePutOrCall),
				quickfix.GroupElement(tag.InstrumentScopeFlexibleIndicator),
				quickfix.GroupElement(tag.InstrumentScopeCouponRate),
				quickfix.GroupElement(tag.InstrumentScopeSecurityExchange),
				quickfix.GroupElement(tag.InstrumentScopeSecurityDesc),
				quickfix.GroupElement(tag.InstrumentScopeEncodedSecurityDescLen),
				quickfix.GroupElement(tag.InstrumentScopeEncodedSecurityDesc),
				quickfix.GroupElement(tag.InstrumentScopeSettlType),
				quickfix.GroupElement(tag.InstrumentScopeUPICode),
			},
		),
	}
}

// Add create and append a new NoInstrumentScopes to this group.
func (m NoInstrumentScopesRepeatingGroup) Add() NoInstrumentScopes {
	g := m.RepeatingGroup.Add()
	return NoInstrumentScopes{g}
}

// Get returns the ith NoInstrumentScopes in the NoInstrumentScopesRepeatinGroup.
func (m NoInstrumentScopesRepeatingGroup) Get(i int) NoInstrumentScopes {
	return NoInstrumentScopes{m.RepeatingGroup.Get(i)}
}

// NoRelatedMarketSegments is a repeating group element, Tag 2545.
type NoRelatedMarketSegments struct {
	*quickfix.Group
}

// SetRelatedMarketSegmentID sets RelatedMarketSegmentID, Tag 2546.
func (m NoRelatedMarketSegments) SetRelatedMarketSegmentID(v string) {
	m.Set(field.NewRelatedMarketSegmentID(v))
}

// SetMarketSegmentRelationship sets MarketSegmentRelationship, Tag 2547.
func (m NoRelatedMarketSegments) SetMarketSegmentRelationship(v enum.MarketSegmentRelationship) {
	m.Set(field.NewMarketSegmentRelationship(v))
}

// GetRelatedMarketSegmentID gets RelatedMarketSegmentID, Tag 2546.
func (m NoRelatedMarketSegments) GetRelatedMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedMarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketSegmentRelationship gets MarketSegmentRelationship, Tag 2547.
func (m NoRelatedMarketSegments) GetMarketSegmentRelationship() (v enum.MarketSegmentRelationship, err quickfix.MessageRejectError) {
	var f field.MarketSegmentRelationshipField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRelatedMarketSegmentID returns true if RelatedMarketSegmentID is present, Tag 2546.
func (m NoRelatedMarketSegments) HasRelatedMarketSegmentID() bool {
	return m.Has(tag.RelatedMarketSegmentID)
}

// HasMarketSegmentRelationship returns true if MarketSegmentRelationship is present, Tag 2547.
func (m NoRelatedMarketSegments) HasMarketSegmentRelationship() bool {
	return m.Has(tag.MarketSegmentRelationship)
}

// NoRelatedMarketSegmentsRepeatingGroup is a repeating group, Tag 2545.
type NoRelatedMarketSegmentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedMarketSegmentsRepeatingGroup returns an initialized, NoRelatedMarketSegmentsRepeatingGroup.
func NewNoRelatedMarketSegmentsRepeatingGroup() NoRelatedMarketSegmentsRepeatingGroup {
	return NoRelatedMarketSegmentsRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoRelatedMarketSegments,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.RelatedMarketSegmentID),
				quickfix.GroupElement(tag.MarketSegmentRelationship),
			},
		),
	}
}

// Add create and append a new NoRelatedMarketSegments to this group.
func (m NoRelatedMarketSegmentsRepeatingGroup) Add() NoRelatedMarketSegments {
	g := m.RepeatingGroup.Add()
	return NoRelatedMarketSegments{g}
}

// Get returns the ith NoRelatedMarketSegments in the NoRelatedMarketSegmentsRepeatinGroup.
func (m NoRelatedMarketSegmentsRepeatingGroup) Get(i int) NoRelatedMarketSegments {
	return NoRelatedMarketSegments{m.RepeatingGroup.Get(i)}
}

// NoAuctionTypeRules is a repeating group element, Tag 2548.
type NoAuctionTypeRules struct {
	*quickfix.Group
}

// SetAuctionType sets AuctionType, Tag 1803.
func (m NoAuctionTypeRules) SetAuctionType(v enum.AuctionType) {
	m.Set(field.NewAuctionType(v))
}

// SetAuctionTypeProductComplex sets AuctionTypeProductComplex, Tag 2549.
func (m NoAuctionTypeRules) SetAuctionTypeProductComplex(v string) {
	m.Set(field.NewAuctionTypeProductComplex(v))
}

// GetAuctionType gets AuctionType, Tag 1803.
func (m NoAuctionTypeRules) GetAuctionType() (v enum.AuctionType, err quickfix.MessageRejectError) {
	var f field.AuctionTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAuctionTypeProductComplex gets AuctionTypeProductComplex, Tag 2549.
func (m NoAuctionTypeRules) GetAuctionTypeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.AuctionTypeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAuctionType returns true if AuctionType is present, Tag 1803.
func (m NoAuctionTypeRules) HasAuctionType() bool {
	return m.Has(tag.AuctionType)
}

// HasAuctionTypeProductComplex returns true if AuctionTypeProductComplex is present, Tag 2549.
func (m NoAuctionTypeRules) HasAuctionTypeProductComplex() bool {
	return m.Has(tag.AuctionTypeProductComplex)
}

// NoAuctionTypeRulesRepeatingGroup is a repeating group, Tag 2548.
type NoAuctionTypeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoAuctionTypeRulesRepeatingGroup returns an initialized, NoAuctionTypeRulesRepeatingGroup.
func NewNoAuctionTypeRulesRepeatingGroup() NoAuctionTypeRulesRepeatingGroup {
	return NoAuctionTypeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoAuctionTypeRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.AuctionType),
				quickfix.GroupElement(tag.AuctionTypeProductComplex),
			},
		),
	}
}

// Add create and append a new NoAuctionTypeRules to this group.
func (m NoAuctionTypeRulesRepeatingGroup) Add() NoAuctionTypeRules {
	g := m.RepeatingGroup.Add()
	return NoAuctionTypeRules{g}
}

// Get returns the ith NoAuctionTypeRules in the NoAuctionTypeRulesRepeatinGroup.
func (m NoAuctionTypeRulesRepeatingGroup) Get(i int) NoAuctionTypeRules {
	return NoAuctionTypeRules{m.RepeatingGroup.Get(i)}
}

// NoPriceRangeRules is a repeating group element, Tag 2550.
type NoPriceRangeRules struct {
	*quickfix.Group
}

// SetStartPriceRange sets StartPriceRange, Tag 2551.
func (m NoPriceRangeRules) SetStartPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartPriceRange(value, scale))
}

// SetEndPriceRange sets EndPriceRange, Tag 2552.
func (m NoPriceRangeRules) SetEndPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndPriceRange(value, scale))
}

// SetPriceRangeValue sets PriceRangeValue, Tag 2553.
func (m NoPriceRangeRules) SetPriceRangeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceRangeValue(value, scale))
}

// SetPriceRangePercentage sets PriceRangePercentage, Tag 2554.
func (m NoPriceRangeRules) SetPriceRangePercentage(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceRangePercentage(value, scale))
}

// SetPriceRangeRuleID sets PriceRangeRuleID, Tag 2556.
func (m NoPriceRangeRules) SetPriceRangeRuleID(v string) {
	m.Set(field.NewPriceRangeRuleID(v))
}

// SetPriceRangeProductComplex sets PriceRangeProductComplex, Tag 2555.
func (m NoPriceRangeRules) SetPriceRangeProductComplex(v string) {
	m.Set(field.NewPriceRangeProductComplex(v))
}

// GetStartPriceRange gets StartPriceRange, Tag 2551.
func (m NoPriceRangeRules) GetStartPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StartPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndPriceRange gets EndPriceRange, Tag 2552.
func (m NoPriceRangeRules) GetEndPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EndPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceRangeValue gets PriceRangeValue, Tag 2553.
func (m NoPriceRangeRules) GetPriceRangeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceRangeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceRangePercentage gets PriceRangePercentage, Tag 2554.
func (m NoPriceRangeRules) GetPriceRangePercentage() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceRangePercentageField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceRangeRuleID gets PriceRangeRuleID, Tag 2556.
func (m NoPriceRangeRules) GetPriceRangeRuleID() (v string, err quickfix.MessageRejectError) {
	var f field.PriceRangeRuleIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceRangeProductComplex gets PriceRangeProductComplex, Tag 2555.
func (m NoPriceRangeRules) GetPriceRangeProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.PriceRangeProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasStartPriceRange returns true if StartPriceRange is present, Tag 2551.
func (m NoPriceRangeRules) HasStartPriceRange() bool {
	return m.Has(tag.StartPriceRange)
}

// HasEndPriceRange returns true if EndPriceRange is present, Tag 2552.
func (m NoPriceRangeRules) HasEndPriceRange() bool {
	return m.Has(tag.EndPriceRange)
}

// HasPriceRangeValue returns true if PriceRangeValue is present, Tag 2553.
func (m NoPriceRangeRules) HasPriceRangeValue() bool {
	return m.Has(tag.PriceRangeValue)
}

// HasPriceRangePercentage returns true if PriceRangePercentage is present, Tag 2554.
func (m NoPriceRangeRules) HasPriceRangePercentage() bool {
	return m.Has(tag.PriceRangePercentage)
}

// HasPriceRangeRuleID returns true if PriceRangeRuleID is present, Tag 2556.
func (m NoPriceRangeRules) HasPriceRangeRuleID() bool {
	return m.Has(tag.PriceRangeRuleID)
}

// HasPriceRangeProductComplex returns true if PriceRangeProductComplex is present, Tag 2555.
func (m NoPriceRangeRules) HasPriceRangeProductComplex() bool {
	return m.Has(tag.PriceRangeProductComplex)
}

// NoPriceRangeRulesRepeatingGroup is a repeating group, Tag 2550.
type NoPriceRangeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPriceRangeRulesRepeatingGroup returns an initialized, NoPriceRangeRulesRepeatingGroup.
func NewNoPriceRangeRulesRepeatingGroup() NoPriceRangeRulesRepeatingGroup {
	return NoPriceRangeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoPriceRangeRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.StartPriceRange),
				quickfix.GroupElement(tag.EndPriceRange),
				quickfix.GroupElement(tag.PriceRangeValue),
				quickfix.GroupElement(tag.PriceRangePercentage),
				quickfix.GroupElement(tag.PriceRangeRuleID),
				quickfix.GroupElement(tag.PriceRangeProductComplex),
			},
		),
	}
}

// Add create and append a new NoPriceRangeRules to this group.
func (m NoPriceRangeRulesRepeatingGroup) Add() NoPriceRangeRules {
	g := m.RepeatingGroup.Add()
	return NoPriceRangeRules{g}
}

// Get returns the ith NoPriceRangeRules in the NoPriceRangeRulesRepeatinGroup.
func (m NoPriceRangeRulesRepeatingGroup) Get(i int) NoPriceRangeRules {
	return NoPriceRangeRules{m.RepeatingGroup.Get(i)}
}

// NoQuoteSizeRules is a repeating group element, Tag 2558.
type NoQuoteSizeRules struct {
	*quickfix.Group
}

// SetMinBidSize sets MinBidSize, Tag 647.
func (m NoQuoteSizeRules) SetMinBidSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinBidSize(value, scale))
}

// SetMinOfferSize sets MinOfferSize, Tag 648.
func (m NoQuoteSizeRules) SetMinOfferSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinOfferSize(value, scale))
}

// SetFastMarketIndicator sets FastMarketIndicator, Tag 2447.
func (m NoQuoteSizeRules) SetFastMarketIndicator(v bool) {
	m.Set(field.NewFastMarketIndicator(v))
}

// GetMinBidSize gets MinBidSize, Tag 647.
func (m NoQuoteSizeRules) GetMinBidSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinBidSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinOfferSize gets MinOfferSize, Tag 648.
func (m NoQuoteSizeRules) GetMinOfferSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinOfferSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFastMarketIndicator gets FastMarketIndicator, Tag 2447.
func (m NoQuoteSizeRules) GetFastMarketIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FastMarketIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMinBidSize returns true if MinBidSize is present, Tag 647.
func (m NoQuoteSizeRules) HasMinBidSize() bool {
	return m.Has(tag.MinBidSize)
}

// HasMinOfferSize returns true if MinOfferSize is present, Tag 648.
func (m NoQuoteSizeRules) HasMinOfferSize() bool {
	return m.Has(tag.MinOfferSize)
}

// HasFastMarketIndicator returns true if FastMarketIndicator is present, Tag 2447.
func (m NoQuoteSizeRules) HasFastMarketIndicator() bool {
	return m.Has(tag.FastMarketIndicator)
}

// NoQuoteSizeRulesRepeatingGroup is a repeating group, Tag 2558.
type NoQuoteSizeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoQuoteSizeRulesRepeatingGroup returns an initialized, NoQuoteSizeRulesRepeatingGroup.
func NewNoQuoteSizeRulesRepeatingGroup() NoQuoteSizeRulesRepeatingGroup {
	return NoQuoteSizeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoQuoteSizeRules,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.MinBidSize),
				quickfix.GroupElement(tag.MinOfferSize),
				quickfix.GroupElement(tag.FastMarketIndicator),
			},
		),
	}
}

// Add create and append a new NoQuoteSizeRules to this group.
func (m NoQuoteSizeRulesRepeatingGroup) Add() NoQuoteSizeRules {
	g := m.RepeatingGroup.Add()
	return NoQuoteSizeRules{g}
}

// Get returns the ith NoQuoteSizeRules in the NoQuoteSizeRulesRepeatinGroup.
func (m NoQuoteSizeRulesRepeatingGroup) Get(i int) NoQuoteSizeRules {
	return NoQuoteSizeRules{m.RepeatingGroup.Get(i)}
}

// NoFlexProductEligibilities is a repeating group element, Tag 2560.
type NoFlexProductEligibilities struct {
	*quickfix.Group
}

// SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242.
func (m NoFlexProductEligibilities) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

// SetFlexProductEligibilityComplex sets FlexProductEligibilityComplex, Tag 2561.
func (m NoFlexProductEligibilities) SetFlexProductEligibilityComplex(v string) {
	m.Set(field.NewFlexProductEligibilityComplex(v))
}

// GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242.
func (m NoFlexProductEligibilities) GetFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexProductEligibilityComplex gets FlexProductEligibilityComplex, Tag 2561.
func (m NoFlexProductEligibilities) GetFlexProductEligibilityComplex() (v string, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242.
func (m NoFlexProductEligibilities) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

// HasFlexProductEligibilityComplex returns true if FlexProductEligibilityComplex is present, Tag 2561.
func (m NoFlexProductEligibilities) HasFlexProductEligibilityComplex() bool {
	return m.Has(tag.FlexProductEligibilityComplex)
}

// NoFlexProductEligibilitiesRepeatingGroup is a repeating group, Tag 2560.
type NoFlexProductEligibilitiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoFlexProductEligibilitiesRepeatingGroup returns an initialized, NoFlexProductEligibilitiesRepeatingGroup.
func NewNoFlexProductEligibilitiesRepeatingGroup() NoFlexProductEligibilitiesRepeatingGroup {
	return NoFlexProductEligibilitiesRepeatingGroup{
		quickfix.NewRepeatingGroup(
			tag.NoFlexProductEligibilities,
			quickfix.GroupTemplate{
				quickfix.GroupElement(tag.FlexProductEligibilityIndicator),
				quickfix.GroupElement(tag.FlexProductEligibilityComplex),
			},
		),
	}
}

// Add create and append a new NoFlexProductEligibilities to this group.
func (m NoFlexProductEligibilitiesRepeatingGroup) Add() NoFlexProductEligibilities {
	g := m.RepeatingGroup.Add()
	return NoFlexProductEligibilities{g}
}

// Get returns the ith NoFlexProductEligibilities in the NoFlexProductEligibilitiesRepeatinGroup.
func (m NoFlexProductEligibilitiesRepeatingGroup) Get(i int) NoFlexProductEligibilities {
	return NoFlexProductEligibilities{m.RepeatingGroup.Get(i)}
}
