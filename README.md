# sap-api-integrations-sales-quotation-creates
sap-api-integrations-sales-quotation-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 販売見積データを登録するマイクロサービスです。    
sap-api-integrations-sales-quotation-creates には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-sales-quotation-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SALES_QUOTATION_SRV_0001/overview  

## 動作環境  
sap-api-integrations-sales-quotation-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-sales-quotation-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-quotation-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SALES_QUOTATION_SRV_0001/overview  
* APIサービス名(=baseURL): API_SALES_QUOTATION_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-quotation-creates には、次の API をコールするためのリソースが含まれています。  

* A_SalesQuotation（販売見積)

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめ登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "SAPSalesQuotationCreates",
	"accepter": ["Header"],
	"sales_quotaion": "",
	"deleted": false
```
  
* 全データを登録する際のsample.jsonの記載例(2)  

全データを登録する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPSalesQuotationCreates",
	"accepter": ["All"],
	"sales_quotaion": "",
	"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostSalesQuotation(
	headerItem *requests.HeaderItem,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "HeaderItem":
			func() {
				c.HeaderItem(headerItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 販売見積 の ヘッダ/明細データ が登録された結果の JSON の例です。  
以下の項目のうち、"SalesInquiry" ～ "ToHeaderPartner" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-quotation-creates/SAP_API_Caller/caller.go#L59",
	"function": "sap-api-integrations-sales-quotation-creates/SAP_API_Caller.(*SAPAPICaller).HeaderItem",
	"level": "INFO",
	"message": {
		"SalesQuotation": "20000002",
		"SalesQuotationType": "QT",
		"SalesOrganization": "0001",
		"DistributionChannel": "01",
		"OrganizationDivision": "01",
		"SalesGroup": "",
		"SalesOffice": "",
		"SalesDistrict": "",
		"SoldToParty": "1",
		"PurchaseOrderByCustomer": "",
		"SalesQuotationDate": "2022-09-20",
		"TransactionCurrency": "",
		"SDDocumentReason": "",
		"PricingDate": "2022-09-21",
		"RequestedDeliveryDate": "2022-10-31",
		"ShippingCondition": "",
		"CompleteDeliveryIsDefined": null,
		"ShippingType": "",
		"HeaderBillingBlockReason": "",
		"DeliveryBlockReason": "",
		"BindingPeriodValidityStartDate": "2022-09-21",
		"BindingPeriodValidityEndDate": "2022-09-30",
		"HdrOrderProbabilityInPercent": "",
		"IncotermsClassification": "",
		"CustomerPaymentTerms": "",
		"CustomerPriceGroup": "",
		"PriceListType": "",
		"PaymentMethod": "",
		"ReferenceSDDocument": "",
		"ReferenceSDDocumentCategory": "",
		"SalesQuotationApprovalReason": "",
		"SalesDocApprovalStatus": "",
		"OverallSDProcessStatus": "",
		"TotalCreditCheckStatus": "",
		"OverallSDDocumentRejectionSts": "",
		"to_Item": {
			"results": [
				{
					"SalesQuotation": "",
					"SalesQuotationItem": "10",
					"HigherLevelItem": null,
					"SalesQuotationItemCategory": "",
					"SalesQuotationItemText": "",
					"PurchaseOrderByCustomer": "",
					"Material": "21",
					"MaterialByCustomer": "",
					"RequestedQuantity": "100",
					"RequestedQuantityUnit": "",
					"ItemOrderProbabilityInPercent": "",
					"ItemWeightUnit": "",
					"ItemVolumeUnit": "",
					"TransactionCurrency": "",
					"MaterialGroup": "",
					"MaterialPricingGroup": "",
					"Batch": "",
					"Plant": "",
					"IncotermsClassification": "",
					"CustomerPaymentTerms": "",
					"SalesDocumentRjcnReason": "",
					"WBSElement": "",
					"ProfitCenter": "",
					"ReferenceSDDocument": "",
					"ReferenceSDDocumentItem": "",
					"SDProcessStatus": ""
				},
				{
					"SalesQuotation": "",
					"SalesQuotationItem": "20",
					"HigherLevelItem": null,
					"SalesQuotationItemCategory": "",
					"SalesQuotationItemText": "",
					"PurchaseOrderByCustomer": "",
					"Material": "21",
					"MaterialByCustomer": "",
					"RequestedQuantity": "200",
					"RequestedQuantityUnit": "",
					"ItemOrderProbabilityInPercent": "",
					"ItemWeightUnit": "",
					"ItemVolumeUnit": "",
					"TransactionCurrency": "",
					"MaterialGroup": "",
					"MaterialPricingGroup": "",
					"Batch": "",
					"Plant": "",
					"IncotermsClassification": "",
					"CustomerPaymentTerms": "",
					"SalesDocumentRjcnReason": "",
					"WBSElement": "",
					"ProfitCenter": "",
					"ReferenceSDDocument": "",
					"ReferenceSDDocumentItem": "",
					"SDProcessStatus": ""
				}
			]
		}
	},
	"time": "2022-09-23T19:47:48+09:00"
}

```