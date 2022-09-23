package sap_api_input_reader

import (
	"sap-api-integrations-sales-quotation-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.SalesQuotation
	results := make([]requests.Item, 0, len(data.SalesQuotationItem))

	header := sdc.ConvertToHeader()

	for i := range data.SalesQuotationItem {
		results = append(results, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.SalesQuotation
	return &requests.Header{
		SalesQuotation:       data.SalesQuotation,
		SalesQuotationType:   data.SalesQuotationType,
		SalesOrganization:    data.SalesOrganization,
		DistributionChannel:  data.DistributionChannel,
		OrganizationDivision: data.OrganizationDivision,
		SalesGroup:           data.SalesGroup,
		SalesOffice:          data.SalesOffice,
		SalesDistrict:        data.SalesDistrict,
		SoldToParty:          data.SoldToParty,
		// CreationDate:                   data.CreationDate,
		// LastChangeDate:                 data.LastChangeDate,
		PurchaseOrderByCustomer: data.PurchaseOrderByCustomer,
		// CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
		SalesQuotationDate: data.SalesQuotationDate,
		// TotalNetAmount:                 data.TotalNetAmount,
		TransactionCurrency:            data.TransactionCurrency,
		SDDocumentReason:               data.SDDocumentReason,
		PricingDate:                    data.PricingDate,
		RequestedDeliveryDate:          data.RequestedDeliveryDate,
		ShippingCondition:              data.ShippingCondition,
		CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
		ShippingType:                   data.ShippingType,
		HeaderBillingBlockReason:       data.HeaderBillingBlockReason,
		DeliveryBlockReason:            data.DeliveryBlockReason,
		BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
		BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
		HdrOrderProbabilityInPercent:   data.HdrOrderProbabilityInPercent,
		// ExpectedOrderNetAmount:         data.ExpectedOrderNetAmount,
		IncotermsClassification: data.IncotermsClassification,
		CustomerPaymentTerms:    data.CustomerPaymentTerms,
		CustomerPriceGroup:      data.CustomerPriceGroup,
		PriceListType:           data.PriceListType,
		PaymentMethod:           data.PaymentMethod,
		// CustomerTaxClassification1:    data.CustomerTaxClassification1,
		ReferenceSDDocument:           data.ReferenceSDDocument,
		ReferenceSDDocumentCategory:   data.ReferenceSDDocumentCategory,
		SalesQuotationApprovalReason:  data.SalesQuotationApprovalReason,
		SalesDocApprovalStatus:        data.SalesDocApprovalStatus,
		OverallSDProcessStatus:        data.OverallSDProcessStatus,
		TotalCreditCheckStatus:        data.TotalCreditCheckStatus,
		OverallSDDocumentRejectionSts: data.OverallSDDocumentRejectionSts,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataSalesQuotation := sdc.SalesQuotation
	data := sdc.SalesQuotation.SalesQuotationItem[num]
	return &requests.Item{
		SalesQuotation:                dataSalesQuotation.SalesQuotation,
		SalesQuotationItem:            data.SalesQuotationItem,
		SalesQuotationItemCategory:    data.SalesQuotationItemCategory,
		SalesQuotationItemText:        data.SalesQuotationItemText,
		PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
		Material:                      data.Material,
		MaterialByCustomer:            data.MaterialByCustomer,
		RequestedQuantity:             data.RequestedQuantity,
		RequestedQuantityUnit:         data.RequestedQuantityUnit,
		ItemOrderProbabilityInPercent: data.ItemOrderProbabilityInPercent,
		// ItemGrossWeight:               data.ItemGrossWeight,
		// ItemNetWeight:             data.ItemNetWeight,
		ItemWeightUnit: data.ItemWeightUnit,
		// ItemVolume:                data.ItemVolume,
		ItemVolumeUnit:      data.ItemVolumeUnit,
		TransactionCurrency: data.TransactionCurrency,
		// NetAmount:                 data.NetAmount,
		MaterialGroup:           data.MaterialGroup,
		MaterialPricingGroup:    data.MaterialPricingGroup,
		Batch:                   data.Batch,
		Plant:                   data.Plant,
		IncotermsClassification: data.IncotermsClassification,
		CustomerPaymentTerms:    data.CustomerPaymentTerms,
		// ProductTaxClassification1: data.ProductTaxClassification1,
		SalesDocumentRjcnReason: data.SalesDocumentRjcnReason,
		WBSElement:              data.WBSElement,
		ProfitCenter:            data.ProfitCenter,
		ReferenceSDDocument:     data.ReferenceSDDocument,
		ReferenceSDDocumentItem: data.ReferenceSDDocumentItem,
		SDProcessStatus:         data.SDProcessStatus,
	}
}
