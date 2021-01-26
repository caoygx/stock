package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"stock/model"
	"stock/util"
	"strconv"
	"strings"
)

type ResponseList struct {
	Rc   int `json:"rc"`
	Rt   int `json:"rt"`
	Svr  int `json:"svr"`
	Lt   int `json:"lt"`
	Full int `json:"full"`
	Data struct {
		Total int `json:"total"`
		Diff  []struct {
			F1   int     `json:"f1"`
			F2   float64 `json:"f2"`
			F3   float64 `json:"f3"`
			F4   float64 `json:"f4"`
			F5   int     `json:"f5"`
			F6   float64 `json:"f6"`
			F7   float64 `json:"f7"`
			F8   float64 `json:"f8"`
			F9   float64 `json:"f9"`
			F10  float64 `json:"f10"`
			F11  float64 `json:"f11"`
			F12  string  `json:"f12"`
			F13  int     `json:"f13"`
			F14  string  `json:"f14"`
			F15  float64 `json:"f15"`
			F16  float64 `json:"f16"`
			F17  float64 `json:"f17"`
			F18  float64 `json:"f18"`
			F20  int64   `json:"f20"`
			F21  int     `json:"f21"`
			F22  float64 `json:"f22"`
			F23  float64 `json:"f23"`
			F24  float64 `json:"f24"`
			F25  float64 `json:"f25"`
			F62  float64 `json:"f62"`
			F115 float64 `json:"f115"`
			F128 string  `json:"f128"`
			F140 string  `json:"f140"`
			F141 string  `json:"f141"`
			F136 string  `json:"f136"`
			F152 int     `json:"f152"`
		} `json:"diff"`
	} `json:"data"`
}

type ResponseDetail struct {
	Rc   int `json:"rc"`
	Rt   int `json:"rt"`
	Svr  int `json:"svr"`
	Lt   int `json:"lt"`
	Full int `json:"full"`
	Data struct {
		F43  int     `json:"f43"`
		F44  int     `json:"f44"`
		F45  int     `json:"f45"`
		F46  int     `json:"f46"`
		F47  int     `json:"f47"`
		F48  float64 `json:"f48"`
		F49  int     `json:"f49"`
		F50  int     `json:"f50"`
		F57  string  `json:"f57"`
		F58  string  `json:"f58"`
		F59  int     `json:"f59"`
		F60  int     `json:"f60"`
		F113 string  `json:"f113"`
		F114 string  `json:"f114"`
		F115 string  `json:"f115"`
		F116 float64 `json:"f116"`
		F117 float64 `json:"f117"`
		F127 string  `json:"f127"`
		F130 string  `json:"f130"`
		F131 string  `json:"f131"`
		F132 string  `json:"f132"`
		F133 string  `json:"f133"`
		F135 float64 `json:"f135"`
		F136 float64 `json:"f136"`
		F137 float64 `json:"f137"`
		F138 float64 `json:"f138"`
		F139 string  `json:"f139"`
		F140 float64 `json:"f140"`
		F141 float64 `json:"f141"`
		F142 float64 `json:"f142"`
		F143 float64 `json:"f143"`
		F144 float64 `json:"f144"`
		F145 float64 `json:"f145"`
		F146 float64 `json:"f146"`
		F147 float64 `json:"f147"`
		F148 float64 `json:"f148"`
		F149 float64 `json:"f149"`
		F152 int     `json:"f152"`
		F161 int     `json:"f161"`
		F162 int     `json:"f162"`
		F164 int     `json:"f164"`
		F165 int     `json:"f165"`
		F167 int     `json:"f167"`
		F168 int     `json:"f168"`
		F169 int     `json:"f169"`
		F170 int     `json:"f170"`
		F171 int     `json:"f171"`
		F177 int     `json:"f177"`
		F178 string  `json:"f178"`
		F198 string  `json:"f198"`
		F199 int     `json:"f199"`
		F31  int     `json:"f31"`
		F32  int     `json:"f32"`
		F33  int     `json:"f33"`
		F34  int     `json:"f34"`
		F35  int     `json:"f35"`
		F36  int     `json:"f36"`
		F37  int     `json:"f37"`
		F38  int     `json:"f38"`
		F39  int     `json:"f39"`
		F40  int     `json:"f40"`
		F19  int     `json:"f19"`
		F20  int     `json:"f20"`
		F17  int     `json:"f17"`
		F18  int     `json:"f18"`
		F15  int     `json:"f15"`
		F16  int     `json:"f16"`
		F13  int     `json:"f13"`
		F14  int     `json:"f14"`
		F11  int     `json:"f11"`
		F12  int     `json:"f12"`
	} `json:"data"`
}
type ResponseProductInfo struct {
	Result struct {
		TiCaiXiangQingList []struct {
			SecurityCode       string `json:"SecurityCode"`
			KeyWord            string `json:"KeyWord"`
			MainPoint          string `json:"MainPoint"`
			MainPointCon       string `json:"MainPointCon"`
			Classification     string `json:"Classification"`
			ClassificationName string `json:"ClassificationName"`
			IsPoint            string `json:"IsPoint"`
		} `json:"TiCaiXiangQingList"`
		GuBenGuDong struct {
			ReportDate                              string `json:"ReportDate"`
			SecurityCode                            string `json:"SecurityCode"`
			TotalEquity                             string `json:"TotalEquity"`
			TotalEquityPercent                      string `json:"TotalEquityPercent"`
			Ashares                                 string `json:"Ashares"`
			ASharesPercent                          string `json:"ASharesPercent"`
			ASharesPercentChart                     string `json:"ASharesPercentChart"`
			RestrictedCirculationShares             string `json:"RestrictedCirculationShares"`
			RestrictedCirculationSharesPercent      string `json:"RestrictedCirculationSharesPercent"`
			RestrictedCirculationSharesPercentChart string `json:"RestrictedCirculationSharesPercentChart"`
		} `json:"GuBenGuDong"`
		GuDongRenShuUnit string `json:"GuDongRenShuUnit"`
		GuDongRenShuList []struct {
			SecurityCode            string `json:"SecurityCode"`
			ChangeDate              string `json:"ChangeDate"`
			TotalSh                 string `json:"TotalSh"`
			TotalShChart            string `json:"TotalShChart"`
			ChangeWithLastTermSh    string `json:"ChangeWithLastTermSh"`
			AvgShare                string `json:"AvgShare"`
			ChangeWithLastTermShare string `json:"ChangeWithLastTermShare"`
			StockPrice              string `json:"StockPrice"`
			StockConvergenceRate    string `json:"StockConvergenceRate"`
			SumLishold              string `json:"SumLishold"`
			SumCirLishold           string `json:"SumCirLishold"`
		} `json:"GuDongRenShuList"`
		GuDongTongJi struct {
			SecurityCode              string `json:"SecurityCode"`
			ReportDate                string `json:"ReportDate"`
			ActualController          string `json:"ActualController"`
			ActualControllerRate      string `json:"ActualControllerRate"`
			ShareholderController     string `json:"ShareholderController"`
			ShareholderControllerRate string `json:"ShareholderControllerRate"`
			ShareHolder               string `json:"ShareHolder"`
			ShareHolderRate           string `json:"ShareHolderRate"`
			CirShareHolder            string `json:"CirShareHolder"`
			CirShareHolderRate        string `json:"CirShareHolderRate"`
			OrgHolderNum              string `json:"OrgHolderNum"`
			OrgHolderRate             string `json:"OrgHolderRate"`
			OrgHolderNumPre           string `json:"OrgHolderNumPre"`
			OrgHolderRatePre          string `json:"OrgHolderRatePre"`
		} `json:"GuDongTongJi"`
		ZhuYingGouChengList []struct {
			SecurityCode         string `json:"SecurityCode"`
			ReportType           string `json:"ReportType"`
			ReportDate           string `json:"ReportDate"`
			MainForm             string `json:"MainForm"`
			MainIncome           string `json:"MainIncome"`
			MainIncomeRatio      string `json:"MainIncomeRatio"`
			MainIncomeRatioChart string `json:"MainIncomeRatioChart"`
			Flag                 string `json:"Flag"`
		} `json:"ZhuYingGouChengList"`
		FenHongSongZhuanList []struct {
			SecurityCode   string `json:"SecurityCode"`
			NoticeDate     string `json:"NoticeDate"`
			AssignDscrpt   string `json:"AssignDscrpt"`
			ExDividendDate string `json:"ExDividendDate"`
			RightRegDate   string `json:"RightRegDate"`
		} `json:"FenHongSongZhuanList"`
		JingliRunUnit    string `json:"JingliRunUnit"`
		YingYeShouRuUnit string `json:"YingYeShouRuUnit"`
		DanJiCaiWuList   []struct {
			SecurityCode     string `json:"SecurityCode"`
			ReportDate       string `json:"ReportDate"`
			NetProfit        string `json:"NetProfit"`
			NetProfitChart   string `json:"NetProfitChart"`
			EPS              string `json:"EPS"`
			EPSChart         string `json:"EPSChart"`
			TotalIncome      string `json:"TotalIncome"`
			TotalIncomeChart string `json:"TotalIncomeChart"`
			Year             string `json:"Year"`
		} `json:"DanJiCaiWuList"`
		GuQuanZhiYaTuJieList []struct {
			TradeDate     string `json:"TradeDate"`
			AmtShareRatio string `json:"AmtShareRatio"`
		} `json:"GuQuanZhiYaTuJieList"`
		GuQuanZhiYaTongJiList []struct {
			SecurityCode      string `json:"SecurityCode"`
			SecurityShortName string `json:"SecurityShortName"`
			ShareHolderCode   string `json:"ShareHolderCode"`
			ShareHolderName   string `json:"ShareHolderName"`
			BBallowance       string `json:"BBallowance"`
			HolderRatio       string `json:"HolderRatio"`
			TotalShareRatio   string `json:"TotalShareRatio"`
			NoticeDate        string `json:"NoticeDate"`
			TotalCount        string `json:"TotalCount"`
		} `json:"GuQuanZhiYaTongJiList"`
		ShangYuQingKuang struct {
			TuJieDataList []struct {
				BaoGaoQi       string      `json:"BaoGaoQi"`
				ShangYuZhi     string      `json:"ShangYuZhi"`
				ShangYuZhiText interface{} `json:"ShangYuZhiText"`
				ShangYuZhiUnit string      `json:"ShangYuZhiUnit"`
			} `json:"TuJieDataList"`
			Data struct {
				Code                    string      `json:"Code"`
				BaoGaoQi                string      `json:"BaoGaoQi"`
				BaoGaoQiLeiXing         string      `json:"BaoGaoQiLeiXing"`
				ShangYu                 string      `json:"ShangYu"`
				ShangYuTongBi           string      `json:"ShangYuTongBi"`
				ShangYuJingZiChan       string      `json:"ShangYuJingZiChan"`
				ShangYuJingZiChanTongBi string      `json:"ShangYuJingZiChanTongBi"`
				GuiMuJingLiRun          string      `json:"GuiMuJingLiRun"`
				GuiMuJingLiRunTongBi    string      `json:"GuiMuJingLiRunTongBi"`
				ShiFouXianShi           interface{} `json:"ShiFouXianShi"`
			} `json:"Data"`
		} `json:"ShangYuQingKuang"`
		YanFaTouRuList []interface{} `json:"YanFaTouRuList"`
	} `json:"Result"`
	Status    int         `json:"Status"`
	Message   interface{} `json:"Message"`
	OtherInfo struct {
	} `json:"OtherInfo"`
}

type ResponseProfit struct {
	Result struct {
		ZuiXinZhiBiao struct {
			ReportDate               string `json:"ReportDate"`
			SecurityCode             string `json:"SecurityCode"`
			CodeType                 int    `json:"CodeType"`
			ReportType               string `json:"ReportType"`
			PE                       string `json:"PE"`
			PB                       string `json:"PB"`
			EPS                      string `json:"EPS"`
			BPS                      string `json:"BPS"`
			TotalIncome              string `json:"TotalIncome"`
			TotalIncomeYOY           string `json:"TotalIncomeYOY"`
			NetProfit                string `json:"NetProfit"`
			NetProfitYOY             string `json:"NetProfitYOY"`
			GrossMargin              string `json:"GrossMargin"`
			NetMargin                string `json:"NetMargin"`
			WeightedYieldOnNetAssets string `json:"WeightedYieldOnNetAssets"`
			AssetLiabilityRatio      string `json:"AssetLiabilityRatio"`
			TotalShareCapital        string `json:"TotalShareCapital"`
			TotalMarketValue         string `json:"TotalMarketValue"`
			ShareA                   string `json:"ShareA"`
			MaketValueA              string `json:"MaketValueA"`
			TradableShareA           string `json:"TradableShareA"`
			TradableMarketValueA     string `json:"TradableMarketValueA"`
			ShareB                   string `json:"ShareB"`
			MaketValueB              string `json:"MaketValueB"`
			TradableShareB           string `json:"TradableShareB"`
			TradableMarketValueB     string `json:"TradableMarketValueB"`
			CapitalAdequacyRatio     string `json:"CapitalAdequacyRatio"`
			NPLRatio                 string `json:"NPLRatio"`
			NPLProvisionCoverage     string `json:"NPLProvisionCoverage"`
			Commreve                 string `json:"Commreve"`
			CommreveYOY              string `json:"CommreveYOY"`
			Premiumearned            string `json:"Premiumearned"`
			Comexpend                string `json:"Comexpend"`
			SurrenderRate            string `json:"SurrenderRate"`
			SolvencyRatio            string `json:"SolvencyRatio"`
			SecurityCodeType         int    `json:"SecurityCodeType"`
			CompanyType              string `json:"CompanyType"`
			ListingState             string `json:"ListingState"`
			CodeTypeName             string `json:"CodeTypeName"`
			CDRNum                   string `json:"CDRNum"`
			CDRZs                    string `json:"CDRZs"`
			IsProfit                 string `json:"IsProfit"`
			IsVotDif                 string `json:"IsVotDif"`
			VotRatio                 string `json:"VotRatio"`
			SecurityShortName        string `json:"SecurityShortName"`
			AmtShareRatio            string `json:"AmtShareRatio"`
			Goodwill                 string `json:"Goodwill"`
			Rdexp                    string `json:"Rdexp"`
			RdexpRatio               string `json:"RdexpRatio"`
			Researchstaff            string `json:"Researchstaff"`
			ResearchstaffRatio       string `json:"ResearchstaffRatio"`
		} `json:"ZuiXinZhiBiao"`
		SuoShuBanKuaiList []struct {
			SecurityCode string `json:"SecurityCode"`
			TypeCode     string `json:"TypeCode"`
			TypeName     string `json:"TypeName"`
			Reason       string `json:"Reason"`
			Accuracy     string `json:"Accuracy"`
		} `json:"SuoShuBanKuaiList"`
		XiangGuanGeGuList interface{} `json:"XiangGuanGeGuList"`
		DaShiTiXingList   []struct {
			NoticeDate string `json:"NoticeDate"`
			Content    string `json:"Content"`
			EventType  string `json:"EventType"`
		} `json:"DaShiTiXingList"`
	} `json:"Result"`
	Status    int         `json:"Status"`
	Message   interface{} `json:"Message"`
	OtherInfo struct {
	} `json:"OtherInfo"`
}

type FinancialReport struct {
	ReportDate            string   `json:"ReportDate"`
	Title                 []string `json:"Title"`
	Totaloperatereve      []string `json:"Totaloperatereve"`
	Operatereve           []string `json:"Operatereve"`
	Intreve               []string `json:"Intreve"`
	Premiumearned         []string `json:"Premiumearned"`
	Commreve              []string `json:"Commreve"`
	Otherreve             []string `json:"Otherreve"`
	Totaloperatereveother []string `json:"Totaloperatereveother"`
	Totaloperateexp       []string `json:"Totaloperateexp"`
	Operateexp            []string `json:"Operateexp"`
	Intexp                []string `json:"Intexp"`
	Commexp               []string `json:"Commexp"`
	Rdexp                 []string `json:"Rdexp"`
	Surrenderpremium      []string `json:"Surrenderpremium"`
	Netindemnityexp       []string `json:"Netindemnityexp"`
	Netcontactreserve     []string `json:"Netcontactreserve"`
	Policydiviexp         []string `json:"Policydiviexp"`
	Riexp                 []string `json:"Riexp"`
	Otherexp              []string `json:"Otherexp"`
	Operatetax            []string `json:"Operatetax"`
	Saleexp               []string `json:"Saleexp"`
	Manageexp             []string `json:"Manageexp"`
	Financeexp            []string `json:"Financeexp"`
	Ofwintexp             []string `json:"Ofwintexp"`
	Ofwintreve            []string `json:"Ofwintreve"`
	Assetdevalueloss      []string `json:"Assetdevalueloss"`
	Creddevalueloss       []string `json:"Creddevalueloss"`
	Totaloperateexpother  []string `json:"Totaloperateexpother"`
	Fvalueincome          []string `json:"Fvalueincome"`
	Investincome          []string `json:"Investincome"`
	Investjointincome     []string `json:"Investjointincome"`
	Netexhedgincome       []string `json:"Netexhedgincome"`
	Exchangeincome        []string `json:"Exchangeincome"`
	Adisposalincome       []string `json:"Adisposalincome"`
	Miotherincome         []string `json:"Miotherincome"`
	Operateprofitother    []string `json:"Operateprofitother"`
	Operateprofitbalance  []string `json:"Operateprofitbalance"`
	Operateprofit         []string `json:"Operateprofit"`
	Nonoperatereve        []string `json:"Nonoperatereve"`
	Nonlassetreve         []string `json:"Nonlassetreve"`
	Nonoperateexp         []string `json:"Nonoperateexp"`
	Nonlassetnetloss      []string `json:"Nonlassetnetloss"`
	Sumprofitbalance      []string `json:"Sumprofitbalance"`
	Sumprofit             []string `json:"Sumprofit"`
	Incometax             []string `json:"Incometax"`
	Sumprofitother        []string `json:"Sumprofitother"`
	Unconfirminvloss      []string `json:"Unconfirminvloss"`
	Netprofit             []string `json:"Netprofit"`
	Combinednetprofitb    []string `json:"Combinednetprofitb"`
	Continuousonprofit    []string `json:"Continuousonprofit"`
	Terminationonprofit   []string `json:"Terminationonprofit"`
	Parentnetprofit       []string `json:"Parentnetprofit"`
	Minorityincome        []string `json:"Minorityincome"`
	Kcfjcxsyjlr           []string `json:"Kcfjcxsyjlr"`
	Netprofitother2       []string `json:"Netprofitother2"`
	Netprofitbalance      []string `json:"Netprofitbalance"`
	Basiceps              []string `json:"Basiceps"`
	Dilutedeps            []string `json:"Dilutedeps"`
	Othercincome          []string `json:"Othercincome"`
	Parentothercincome    []string `json:"Parentothercincome"`
	Minorityothercincome  []string `json:"Minorityothercincome"`
	Sumcincome            []string `json:"Sumcincome"`
	Parentcincome         []string `json:"Parentcincome"`
	Minoritycincome       []string `json:"Minoritycincome"`
	Auditopinionsdomestic []string `json:"Auditopinionsdomestic"`
}

type ResponseFinancialReport struct {
	Result struct {
		LiRunBiaoListQuanShang interface{}       `json:"LiRunBiaoList_QuanShang"`
		LiRunBiaoListBaoXian   interface{}       `json:"LiRunBiaoList_BaoXian"`
		LiRunBiaoListQiYe      []FinancialReport `json:"LiRunBiaoList_QiYe"`
		LiRunBiaoListYinHang   interface{}       `json:"LiRunBiaoList_YinHang"`
	} `json:"Result"`
	Status    int    `json:"Status"`
	Message   string `json:"Message"`
	OtherInfo struct {
	} `json:"OtherInfo"`
}

type CompanyInfo struct {
	SecurityCode      string `json:"SecurityCode"`
	CompanyCode       string `json:"CompanyCode"`
	CompanyName       string `json:"CompanyName"`
	PreviousName      string `json:"PreviousName"`
	Provice           string `json:"Provice"`
	Industry          string `json:"Industry"`
	Block             string `json:"Block"`
	Chairman          string `json:"Chairman"`
	Website           string `json:"Website"`
	RegisteredAddress string `json:"RegisteredAddress"`
	OfficeAddress     string `json:"OfficeAddress"`
	CompRofile        string `json:"CompRofile"`
	MainBusiness      string `json:"MainBusiness"`
	SecurityCodeA     string `json:"SecurityCodeA"`
	SecurityNameA     string `json:"SecurityNameA"`
	SecurityCodeB     string `json:"SecurityCodeB"`
	SecurityNameB     string `json:"SecurityNameB"`
	SecurityCodeH     string `json:"SecurityCodeH"`
	SecurityNameH     string `json:"SecurityNameH"`
	Representative    string `json:"Representative"`
	GeneralManager    string `json:"GeneralManager"`
	Secretaries       string `json:"Secretaries"`
	FoundDate         string `json:"FoundDate"`
	RegisteredCapital string `json:"RegisteredCapital"`
	Currency          string `json:"Currency"`
	Employees         string `json:"Employees"`
	Managers          string `json:"Managers"`
	Phone             string `json:"Phone"`
	Email             string `json:"Email"`
	SecurityCodeType  string `json:"SecurityCodeType"`
	CodeType          string `json:"CodeType"`
	IsInnovation      int    `json:"IsInnovation"`
}

type ResponseCompanyInfo struct {
	Result struct {
		JiBenZiLiao CompanyInfo `json:"JiBenZiLiao"`
	} `json:"Result"`
	Status    int    `json:"Status"`
	Message   string `json:"Message"`
	OtherInfo struct {
	} `json:"OtherInfo"`
}

func getAllStock(page int) {
	url := "http://63.push2.eastmoney.com/api/qt/clist/get?cb=jQuery1124009397379737638278_1595857775509&pn=" + strconv.Itoa(page) + "&pz=20&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&_=1595857775516"
	r := util.FileGetContent(url)
	content := string(r)
	content = strings.Replace(content, "jQuery1124009397379737638278_1595857775509(", "", -1)
	content = strings.Replace(content, ");", "", -1)
	res := ResponseList{}
	json.Unmarshal([]byte(content), &res)
	for _, v := range res.Data.Diff {

		row := model.Piao{}
		model.DB.Where("code=?", v.F12).Find(&row)
		row.Code = v.F12
		row.Name = v.F14
		row.Price = v.F2
		if row.ID > 0 { //更新
			row.Update()
		} else { //新增
			row.Save()
		}

		//row.v.F116
		//v.F117

	}
	fmt.Print(res)
}

func getDetail(row *model.Piao) *model.Piao {

	var market string
	prefix := row.Code[0:2]
	if prefix == "60" || prefix == "68" {
		market = "1"
	} else if prefix == "00" || prefix == "30" {
		market = "0"
	}
	url := "https://push2.eastmoney.com/api/qt/stock/get?secid=" + market + "." + row.Code + "&ut=f057cbcbce2a86e2866ab8877db1d059&forcect=1&fields=f13,f19,f20,f23,f24,f25,f26,f27,f28,f29,f30,f43,f44,f45,f46,f47,f48,f49,f50,f57,f58,f59,f60,f113,f114,f115,f116,f117,f127,f130,f131,f132,f133,f135,f136,f137,f138,f139,f140,f141,f142,f143,f144,f145,f146,f147,f148,f149,f152,f161,f162,f164,f165,f167,f168,f169,f170,f171,f177,f178,f198,f199,f530,f531&invt=2&cb=jQuery34101417743628563637_1595924232053&_=1595924232055"
	r := util.FileGetContent(url)
	content := string(r)
	content = strings.Replace(content, "jQuery34101417743628563637_1595924232053(", "", -1)
	content = strings.Replace(content, ");", "", -1)
	res := ResponseDetail{}
	json.Unmarshal([]byte(content), &res)

	//row := model.Piao{}
	//row.Code = "605188"
	//row.Code = code
	//fmt.Println(res.Data)
	row.MarketCap = res.Data.F116
	row.Industry = res.Data.F127
	//row.PE = res.Data.F162
	//row.Description =
	row.DetailJson = content
	//fmt.Println("content:" + content)
	//os.Exit(1)

	return row
	//row.UpdateAll()

	// for _, v := range res.Data.Diff {
	// 	row := model.Piao{}
	// 	row.Code = v.F12
	// 	row.Name = v.F14
	// 	row.Price = v.F2
	// 	row.Save(row)
	// }
	// fmt.Print(res)
}

func getGanlian() {
	//url := "https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart1Get?fc=00217402&color=w"
	//resData := `{"Result":{"ZuiXinZhiBiao":{"ReportDate":"2020-03-31","SecurityCode":"002174.SZ","CodeType":1,"ReportType":"0","PE":"13.21","PB":"3.79","EPS":"0.4028","BPS":"5.6168","TotalIncome":"12.11亿","TotalIncomeYOY":"45.45%","NetProfit":"3.631亿","NetProfitYOY":"110.14%","GrossMargin":"46.19%","NetMargin":"29.97%","WeightedYieldOnNetAssets":"7.14%","AssetLiabilityRatio":"42.79%","TotalShareCapital":"9.014亿","TotalMarketValue":"191.8亿","ShareA":"9.014亿","MaketValueA":"191.8亿","TradableShareA":"6.752亿","TradableMarketValueA":"143.7亿","ShareB":"--","MaketValueB":"--","TradableShareB":"--","TradableMarketValueB":"--","CapitalAdequacyRatio":"--","NPLRatio":"--","NPLProvisionCoverage":"--","Commreve":"--","CommreveYOY":"--","Premiumearned":"--","Comexpend":"--","SurrenderRate":"--","SolvencyRatio":"--","SecurityCodeType":1,"CompanyType":"4","ListingState":"0","CodeTypeName":"A股","CDRNum":"--","CDRZs":"--","IsProfit":"尚未盈利","IsVotDif":"无差异","VotRatio":"--","SecurityShortName":"游族网络","AmtShareRatio":"20.75%","Goodwill":"4.52亿","Rdexp":"--","RdexpRatio":"--","Researchstaff":"--","ResearchstaffRatio":"--"},"SuoShuBanKuaiList":[{"SecurityCode":"002174.SZ","TypeCode":"923","TypeName":"字节概念","Reason":"2020年5月29日回复称公司与字节跳动建立长期深度渠道合作关系,公司游戏通过头条,抖音,TikTok等实现新用户的引流,为公司产品提供精准化营销。","Accuracy":"1"},{"SecurityCode":"002174.SZ","TypeCode":"903","TypeName":"云游戏","Reason":"与华为共同推进云游戏解决方案并发布云游戏产品,包括ARM安卓云游戏和PC云游戏等。现已加入云游戏产业联盟(CGIA)","Accuracy":"1"},{"SecurityCode":"002174.SZ","TypeCode":"854","TypeName":"华为概念","Reason":"2019年年报披露:公司于2019年11月与华为签署合作框架协议,联手开展云游戏合作,共同开拓云游戏行业市场。","Accuracy":"1"},{"SecurityCode":"002174.SZ","TypeCode":"853","TypeName":"电子竞技","Reason":"主营移动网络游戏的研发和运营。公司副总裁谭雁峰表示,公司看好电竞领域,基于《女神联盟》大的IP已开发出相关产品,并在电竞相关产业链层面有一些进入。","Accuracy":"1"},{"SecurityCode":"002174.SZ","TypeCode":"801","TypeName":"增强现实","Reason":"2020年5月18日回复称公司已联手号百控股、中国电信共建“5G-VR小镇”。","Accuracy":"1"},{"SecurityCode":"002174.SZ","TypeCode":"642","TypeName":"手游概念","Reason":" 公司主营网页网络游戏、移动网络游戏的研发和运营。2016年3月公告,公司全资子公司YOUSU GmbH于2016年3月22日在德国与Bigpoint Investments GmbH签订《股权购买协议》,拟购买其持有的 Bigpoint HoldCo GmbH(以下简称“BP”)100%股权,拟投资金额不超过8000万欧元(约5.8亿人民币)。BP是立足德国,欧洲领先的电脑网络游戏和手机游戏的开发商、内容提供商和运营商。游戏均免费提供给用户,收入来自于用户的游戏内消费(购买虚拟道具)。","Accuracy":"1"},{"SecurityCode":"002174.SZ","TypeCode":"634","TypeName":"大数据","Reason":"公司16年12月29日公告中披露,本次非公开发行募集资金拟用于网络游戏研发及发行运营建设项目、大数据分析与运营建设项目以及偿还银行贷款。其中,大数据分析与运营建设项目拟在3年建设期内投入18,622.00万元用于包括机柜租赁、机房维护、系统集成以及带宽费用在内的IDC机房费用。","Accuracy":"1"},{"SecurityCode":"002174.SZ","TypeCode":"509","TypeName":"网络游戏","Reason":"公司主营网页网络游戏、移动网络游戏的研发和运营。2016年2月份,公司拟向包含广发资管在内的不超10名特定投资者非公开发行不超1亿股,募集资金不超44.5亿元。募资投向:拟19.4亿元收购北京青果灵动科技97%股权,拟9.7亿元投入网络游戏研发及发行运营建设项目,拟11.9亿元投入大数据分析与运营建设项目,拟3.5亿元偿还银行贷款。青果灵动是中国领先的3D跨平台游戏企业。","Accuracy":"1"}],"XiangGuanGeGuList":null,"DaShiTiXingList":[{"NoticeDate":"2020-08-26","Content":"披露2020年中报","EventType":"报表披露"},{"NoticeDate":"2020-07-17","Content":"预计2020年1-6月业绩略增,归属净利润约4.885亿元至5.088亿元,同比增长20%至25%","EventType":"业绩预告"},{"NoticeDate":"2020-07-07","Content":"林奇(董事、高管)减持4446万股，交易均价为20.37元，变动途径：协议转让","EventType":"高管增减持"},{"NoticeDate":"2020-06-23","Content":"太平基金管理有限公司,太平基金-中国银行-太平基金-太平人寿-盛世锐进1号资产管理计划于2020-05-29减持1.01万股,交易均价19.23元","EventType":"股东增减持"},{"NoticeDate":"2020-04-30","Content":"2020年一季报归属净利润3.631亿元,同比增长110.14%,基本每股收益0.41元","EventType":"报表披露"}]},"Status":0,"Message":null,"OtherInfo":{}}`
}

func getCode(code string) string {
	prefix := code[0:2]
	var requestCode string
	if prefix == "60" || prefix == "68" {
		requestCode = code + "01"
	} else if prefix == "00" || prefix == "30" {
		requestCode = code + "02"
	}
	return requestCode
}

func getProduct(row *model.Piao) *model.Piao {
	code := row.Code
	requestCode := getCode(code)
	url := "https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=" + requestCode + "&color=w"
	r := util.FileGetContent(url)
	content := string(r)
	//content = strings.Replace(content, "jQuery34101417743628563637_1595924232053(", "", -1)
	//content = strings.Replace(content, ");", "", -1)
	res := ResponseProductInfo{}
	json.Unmarshal([]byte(content), &res)
	//fmt.Print(res.Result.DanJiCaiWuList)      //所有报表
	//fmt.Print(res.Result.ZhuYingGouChengList) //主营
	var arrZhuying []interface{}
	for _, v := range res.Result.ZhuYingGouChengList {

		if v.ReportType == "3" {
			row := make(map[string]interface{})
			row["MainForm"] = v.MainForm
			row["MainIncome"] = v.MainIncome
			row["MainIncomeRatio"] = v.MainIncomeRatio
			//zhuying += v.MainForm + ":" + v.MainIncome + ":" + v.MainIncomeRatio + "\n"
			arrZhuying = append(arrZhuying, row)
		}

		//fmt.Println(zhuying)

	}
	//row := model.Piao{}
	//row.Code = code
	zhuying, _ := json.Marshal(arrZhuying)
	row.Zhuying = string(zhuying)
	row.ProductJson = content
	return row
	//row.UpdateAll()
	//row.Description =
	//row.UpdateZhuying(row)
	//fmt.Println(row.Zhuying)

	//https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart1Get?fc=00217402&color=w

}

func getProfit(row *model.Piao) *model.Piao {
	code := row.Code
	requestCode := getCode(code)
	url := "https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart1Get?fc=" + requestCode + "&color=w"
	r := util.FileGetContent(url)
	content := string(r)

	//content = strings.Replace(content, "jQuery34101417743628563637_1595924232053(", "", -1)
	//content = strings.Replace(content, ");", "", -1)
	res := ResponseProfit{}
	json.Unmarshal([]byte(content), &res)
	//row := model.Piao{}
	//row.Code = code

	netProfit := res.Result.ZuiXinZhiBiao.NetProfit
	if strings.Contains(netProfit, "万") {
		num, _ := strconv.ParseFloat(strings.Replace(netProfit, "万", "", -1), 64)
		row.NetProfit = strconv.FormatFloat(num*10000, 'f', 2, 64)
	} else if strings.Contains(netProfit, "亿") {
		num, _ := strconv.ParseFloat(strings.Replace(netProfit, "亿", "", -1), 64)
		row.NetProfit = strconv.FormatFloat(num*100000000, 'f', 0, 64)
	}
	row.PE = res.Result.ZuiXinZhiBiao.PE

	row.NetprofitJson = content
	fmt.Println(row)
	//os.Exit(1)
	return row
	//os.Exit(1)
	//	row.NetProfit = res.Result.ZuiXinZhiBiao.NetProfit
	//row.UpdateAll()
	//os.Exit(1)
	//fmt.Print(res.Result.ZuiXinZhiBiao.NetProfit) //利润
}

func get1() {
	//https://emh5.eastmoney.com/api/CaoPanBiDu/GetGongSiGuiMoListGet?fc=00217402&color=w&RankType=1
	//{"Result":{"GongSiGuiMoList":[{"SecurityCode":"002174.SZ","Scode":"","SShortName":"行业均值","RankType":"1","ReportDate":"2020-03-31","PublishCode":"017017001","PublishName":"互联网服务","MarketValue":"164.3亿元","MarketValueChart":"164.35","TradableMarketValue":"--","TradableMarketValueChart":"","TotalIncome":"--","TotalIncomeChart":"","NetProfit":"--","NetProfitChart":"","Rank":"0","TotalCount":"65","publishNum":""},{"SecurityCode":"002174.SZ","Scode":"300413.SZ","SShortName":"芒果超媒","RankType":"1","ReportDate":"2020-03-31","PublishCode":"017017001","PublishName":"互联网服务","MarketValue":"1163亿元","MarketValueChart":"1163.30","TradableMarketValue":"--","TradableMarketValueChart":"","TotalIncome":"--","TotalIncomeChart":"","NetProfit":"--","NetProfitChart":"","Rank":"1","TotalCount":"65","publishNum":""},{"SecurityCode":"002174.SZ","Scode":"002602.SZ","SShortName":"世纪华通","RankType":"1","ReportDate":"2020-03-31","PublishCode":"017017001","PublishName":"互联网服务","MarketValue":"927.1亿元","MarketValueChart":"927.10","TradableMarketValue":"--","TradableMarketValueChart":"","TotalIncome":"--","TotalIncomeChart":"","NetProfit":"--","NetProfitChart":"","Rank":"2","TotalCount":"65","publishNum":""},{"SecurityCode":"002174.SZ","Scode":"002555.SZ","SShortName":"三七互娱","RankType":"1","ReportDate":"2020-03-31","PublishCode":"017017001","PublishName":"互联网服务","MarketValue":"925.8亿元","MarketValueChart":"925.80","TradableMarketValue":"--","TradableMarketValueChart":"","TotalIncome":"--","TotalIncomeChart":"","NetProfit":"--","NetProfitChart":"","Rank":"3","TotalCount":"65","publishNum":""},{"SecurityCode":"002174.SZ","Scode":"002174.SZ","SShortName":"游族网络","RankType":"1","ReportDate":"2020-03-31","PublishCode":"017017001","PublishName":"互联网服务","MarketValue":"191.8亿元","MarketValueChart":"191.81","TradableMarketValue":"--","TradableMarketValueChart":"","TotalIncome":"--","TotalIncomeChart":"","NetProfit":"--","NetProfitChart":"","Rank":"14","TotalCount":"65","publishNum":""}]},"Status":0,"Message":"","OtherInfo":{}}
}
func generateURL() {

	fields := ""

	for i := 1; i < 531; i++ {
		fields += "f" + strconv.Itoa(i) + ","
	}
	fields = strings.TrimRight(fields, ",")

	//fields := "f13,f19,f20,f23,f24,f25,f26,f27,f28,f29,f30,f43,f44,f45,f46,f47,f48,f49,f50,f57,f58,f59,f60,f61,f62,f63,f64,f113,f114,f115,f116,f117,f127,f130,f131,f132,f133,f135,f136,f137,f138,f139,f140,f141,f142,f143,f144,f145,f146,f147,f148,f149,f152,f161,f162,f164,f165,f167,f168,f169,f170,f171,f177,f178,f198,f199,f530,f531"
	//code := "605188"
	//url := "https://push2.eastmoney.com/api/qt/stock/get?secid=0." + code + "&ut=f057cbcbce2a86e2866ab8877db1d059&forcect=1&fields=" + fields + "&invt=2&cb=jQuery34101417743628563637_1595924232053&_=1595924232055"
	url := "https://push2.eastmoney.com/api/qt/stock/get?secid=0.002174&ut=f057cbcbce2a86e2866ab8877db1d059&forcect=1&fields=" + fields + "&invt=2&cb=jQuery34101417743628563637_1595924232053&_=1595924232055"
	fmt.Print(url)
}

func getFinancialReportOfQuarterly(row *model.Piao) {
	code := row.Code
	requestCode := getCode(code)
	url := "https://emh5.eastmoney.com/api/CaiWuFenXi/GetLiRunBiaoList"
	//requestCode = "00065102"
	params := `{"fc":"` + requestCode + `","color":"w","corpType":"4","endDate":"","latestCount":5,"reportDateType":0,"reportType":2}`
	r, _ := util.Post(url, params, "post")
	content := string(r)
	fmt.Println(content)
	fmt.Println(code)

	row.FinancialReportJson = content
	fmt.Println(row)
	//return row

	res := ResponseFinancialReport{}
	json.Unmarshal([]byte(content), &res)
	//fmt.Println(res.Result.LiRunBiaoListQiYe)

	for _, v := range res.Result.LiRunBiaoListQiYe {
		fmt.Println(v.ReportDate)
	}

	//os.Exit(1)

	//netProfit := res.Result.ZuiXinZhiBiao.NetProfit

}

//获取公司人数
func getCompanyInfo(row *model.Piao) *model.Piao { //
	code := row.Code
	requestCode := getCode(code)
	//code := "SZ300059"
	params := `{fc: "` + requestCode + `", color: "w", SecurityCode: "SZ300059"}`
	url := "https://emh5.eastmoney.com/api/GongSiGaiKuang/GetJiBenZiLiao"

	r, _ := util.Post(url, params, "post")
	content := string(r)
	fmt.Println(content)
	fmt.Println(code)
	res := ResponseCompanyInfo{}
	json.Unmarshal([]byte(content), &res)

	row.Employees = res.Result.JiBenZiLiao.Employees
	JiBenZiLiao, _ := json.Marshal(res.Result.JiBenZiLiao)

	row.CompanyInfoJson = string(JiBenZiLiao)
	fmt.Println(row)
	//os.Exit(1)
	//fmt.Println(res)
	return row

}

func getFullInfo() {
	piao := model.Piao{}
	r := piao.GetAll()
	for _, v := range r {
		//fmt.Println(k, v)
		/*getDetail(&v)
		getProduct(&v)
		getProfit(&v)
		getFinancialReportOfQuarterly(&v)*/
		//os.Exit(1)

		getCompanyInfo(&v)

		v.UpdateAll()

		//os.Exit(1)
		//break
		//time.Sleep(time.Second * 1)

		//break

	}
}

func getFund() {

	var page, url string
	var r []byte
	for i := 1; i <= 49; i++ {
		page = strconv.Itoa(i)
		url = "https://fundmobapi.eastmoney.com/FundMNewApi/FundMNRankNewList?callback=jQuery3110662759194538407_1603873115906&fundtype=25&SortColumn=RZDF&Sort=desc&pageIndex=" + page + "&pagesize=30&companyid=&deviceid=Wap&plat=Wap&product=EFund&version=2.0.0&Uid=&_=1603873115914"

		r = util.FileGetContent(url)
		content := string(r)

		//content = `jQuery3110662759194538407_1603873115906({"Datas":[{"FCODE"]})`
		reg := regexp.MustCompile(`.+\((.+)\)`)
		if reg == nil {
			fmt.Println("MustCompile err")
			return
		}
		//fmt.Println(content)
		result := reg.FindAllStringSubmatch(content, -1)
		fmt.Println(result)
		for _, text := range result {
			fmt.Println("text[1] = ", text[0])
		}

		os.Exit(1)

	}

}

func main() {

	//getFund()

	// hangye = model.Hangye{}
	// r := hangye.GetAll()
	// for k, v := range r {

	// }
	/*str := "60122112212"
	content := str[0:2]
	fmt.Print(content)
	*/

	//获取现有股票列表
	//getFullInfo()

	//url = "https://push2.eastmoney.com/api/qt/stock/get?secid=0." + code + "&ut=f057cbcbce2a86e2866ab8877db1d059&forcect=1&fields=" + fields + "&invt=2&cb=jQuery34101417743628563637_1595924232053&_=1595924232055"
	// code := "sh601006"
	// url := "http://hq.sinajs.cn/list=" + code
	// //fmt.Println(url)
	// r := util.FileGetContent(url)
	// fmt.Print(string(r))

	//增加新票和更新最新价格
	for page := 1; page <= 206; page++ {
		getAllStock(page)
		//time.Sleep(time.Second * 3)
	}

}

//https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart1Get?fc=00065102&color=w
//{"Result":{"ZuiXinZhiBiao":{"ReportDate":"2020-06-30","SecurityCode":"000651.SZ","CodeType":1,"ReportType":"0","PE":"25.69","PB":"2.88","EPS":"1.0576","BPS":"18.8859","TotalIncome":"706亿","TotalIncomeYOY":"-28.21%","NetProfit":"63.62亿","NetProfitYOY":"-53.73%","GrossMargin":"21.11%","NetMargin":"9.24%","WeightedYieldOnNetAssets":"5.57%","AssetLiabilityRatio":"59.27%","TotalShareCapital":"60.16亿","TotalMarketValue":"3268亿","ShareA":"60.16亿","MaketValueA":"3268亿","TradableShareA":"59.7亿","TradableMarketValueA":"3244亿","ShareB":"--","MaketValueB":"--","TradableShareB":"--","TradableMarketValueB":"--","CapitalAdequacyRatio":"--","NPLRatio":"--","NPLProvisionCoverage":"--","Commreve":"--","CommreveYOY":"--","Premiumearned":"--","Comexpend":"--","SurrenderRate":"--","SolvencyRatio":"--","SecurityCodeType":1,"CompanyType":"4","ListingState":"0","CodeTypeName":"A股","CDRNum":"--","CDRZs":"--","IsProfit":"尚未盈利","IsVotDif":"无差异","VotRatio":"--","SecurityShortName":"格力电器","AmtShareRatio":"15.74%","Goodwill":"3.26亿","Rdexp":"--","RdexpRatio":"--","Researchstaff":"--","ResearchstaffRatio":"--"},"SuoShuBanKuaiList":[{"SecurityCode":"000651.SZ","TypeCode":"911","TypeName":"口罩","Reason":"2020年3月13日,公司互动平台回应称近期成立珠海格健医疗科技有限公司,生产口罩、测温仪、护目镜等紧缺防疫物资,上述产品已经有产出,目前以满足国内需求为先。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"825","TypeName":"新零售","Reason":"2020年6月4日回复称公司现已开启“新零售”模式的探索,将线上销售与线下门店体验服务两种渠道彼此互补深度融合,构建稳固的线上、线下业务布局,打造全渠道销售平台。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"811","TypeName":"超级品牌","Reason":" 公司是目前全球最大的集研发、生产、销售、服务于一体的专业化空调企业,公司旗下的“格力”品牌空调,是中国空调业唯一的“世界名牌”产品,业务遍及全球100多个国家和地区。在全球拥有珠海、重庆、合肥、巴西、巴基斯坦、越南6大生产基地,5万多名员工,至今已开发出包括家用空调、商用空调在内的20大类、400个系列、7000多个品种规格的产品。力争2014年实现销售收入1400亿元;并力求未来年均销售收入增长200亿元,早日实现五年再造一个“格力”的目标。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"900","TypeName":"新能源车","Reason":"2017年2月20日公告,公司为了切入新能源汽车产业链、储能以及电池制造装备领域,打造公司新的产业增长点,公司拟与珠海银隆新能源有限公司签订《合作协议》,双方及其子、分公司利用各自产业优势,在智能装备、模具、铸造、汽车空调、电机电控、新能源汽车、储能等领域进行合作。在同等条件下,一方优先采购对方产品,购买对方服务。以一个年度为一个周期,甲乙双方相互的优先采购和总金额不超过人民币200亿元。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"683","TypeName":"国企改革","Reason":" 公司最终控制人为珠海市国有资产监督管理委员会。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"680","TypeName":"智能家居","Reason":"以产品多元化及智能家居为切入点,格力探索研究家用消费类电子、通信及工控芯片的发展方向,研发自主知识产权的芯片。2016年,重点在智能交互、智能连接、智能云平台方面进行了技术研发,将公司产品接入到智能家居系统中,再通过无线模块、“格力+”APP、云平台开放接口供第三方厂家接入,形成格力智能家居生态系统。目前已在画时代空调、金贝空调、玫瑰空调、除湿机、电饭煲、空气净化器、净水机、智能油烟机、消毒柜插件、洗衣机等公司重点产品上全部实现无线连接,并可以通过“格力+”APP全部进行远程控制、故障报警以及维修的服务。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"677","TypeName":"粤港自贸","Reason":"公司注册地址位于广东省珠海市，产业覆盖空调、生活电器、高端装备、通信设备等四大领域。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"588","TypeName":"太阳能","Reason":"2020年2月12日回复称公司合金软磁粉芯制成的电感元件应用于光伏发电系统中的光伏发电逆变器上。","Accuracy":"1"},{"SecurityCode":"000651.SZ","TypeCode":"574","TypeName":"锂电池","Reason":"2017年2月20日公告,公司为了切入新能源汽车产业链、储能以及电池制造装备领域,打造公司新的产业增长点,公司拟与珠海银隆新能源有限公司签订《合作协议》,双方及其子、分公司利用各自产业优势,在智能装备、模具、铸造、汽车空调、电机电控、新能源汽车、储能等领域进行合作。在同等条件下,一方优先采购对方产品,购买对方服务。以一个年度为一个周期,甲乙双方相互的优先采购和总金额不超过人民币200亿元。","Accuracy":"1"}],"XiangGuanGeGuList":null,"DaShiTiXingList":[{"NoticeDate":"2020-08-31","Content":"2020年中报归属净利润63.62亿元,同比下降53.73%,基本每股收益1.06元","EventType":"报表披露"},{"NoticeDate":"2020-08-31","Content":"2020中期分配10派10元(含税)(董事会决议通过)","EventType":"分红送转"},{"NoticeDate":"2020-07-10","Content":"河北京海担保投资有限公司于2020-07-09减持823.2万股,变动数量占流通股比例0.14%,交易均价58.92元","EventType":"股东增减持"},{"NoticeDate":"2020-04-30","Content":"2020年一季报归属净利润15.58亿元,同比下降72.53%,基本每股收益0.26元","EventType":"报表披露"},{"NoticeDate":"2020-04-15","Content":"2019年1-12月快报披露,营业总收入2005亿元,归属净利润246.7亿元,基本每股收益4.1元","EventType":"业绩快报"}]},"Status":0,"Message":null,"OtherInfo":{}}

//https://emh5.eastmoney.com/api/CaoPanBiDu/GetCaoPanBiDuPart2Get?fc=00065102&color=w
//{"Result":{"TiCaiXiangQingList":[{"SecurityCode":"000651.SZ","KeyWord":"家电行业","MainPoint":"4","MainPointCon":"","Classification":"004","ClassificationName":"行业背景","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"技术创新","MainPoint":"6","MainPointCon":"","Classification":"005","ClassificationName":"核心竞争力","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"企业文化与管理创新","MainPoint":"7","MainPointCon":"","Classification":"005","ClassificationName":"核心竞争力","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"智能家居","MainPoint":"10","MainPointCon":"","Classification":"005","ClassificationName":"核心竞争力","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"电商平台上线","MainPoint":"13","MainPointCon":"","Classification":"005","ClassificationName":"核心竞争力","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"全球最大空调企业","MainPoint":"15","MainPointCon":"","Classification":"005","ClassificationName":"核心竞争力","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"研发实力强","MainPoint":"16","MainPointCon":"","Classification":"005","ClassificationName":"核心竞争力","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"节能国际领先","MainPoint":"17","MainPointCon":"","Classification":"005","ClassificationName":"核心竞争力","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"拟30亿参与闻泰科技收购安世集团项目","MainPoint":"23","MainPointCon":"","Classification":"006002","ClassificationName":"对外扩张","IsPoint":"1"},{"SecurityCode":"000651.SZ","KeyWord":"控股股东拟通过公开征集受让方的方式协议转让公司15%股权","MainPoint":"34","MainPointCon":"","Classification":"007005","ClassificationName":"其他变动","IsPoint":"1"}],"GuBenGuDong":{"ReportDate":"2020-06-30","SecurityCode":"000651.SZ","TotalEquity":"60.16亿","TotalEquityPercent":"100.00%","Ashares":"59.70亿","ASharesPercent":"99.24%","ASharesPercentChart":"99.2419728221758","RestrictedCirculationShares":"4560.09万","RestrictedCirculationSharesPercent":"0.76%","RestrictedCirculationSharesPercentChart":"0.758027177824161"},"GuDongRenShuUnit":"万","GuDongRenShuList":[{"SecurityCode":"000651.SZ","ChangeDate":"2020-06-30","TotalSh":"55.39万","TotalShChart":"553857","ChangeWithLastTermSh":"11.61%","AvgShare":"1.078万","ChangeWithLastTermShare":"-10.40%","StockPrice":"56.57","StockConvergenceRate":"非常集中","SumLishold":"50.14","SumCirLishold":"49.95"},{"SecurityCode":"000651.SZ","ChangeDate":"2020-03-31","TotalSh":"49.63万","TotalShChart":"496265","ChangeWithLastTermSh":"61.01%","AvgShare":"1.203万","ChangeWithLastTermShare":"-37.89%","StockPrice":"52.2","StockConvergenceRate":"非常集中","SumLishold":"50.03","SumCirLishold":"49.79"},{"SecurityCode":"000651.SZ","ChangeDate":"2019-12-31","TotalSh":"30.82万","TotalShChart":"308228","ChangeWithLastTermSh":"-7.57%","AvgShare":"1.937万","ChangeWithLastTermShare":"8.19%","StockPrice":"65.58","StockConvergenceRate":"非常集中","SumLishold":"50.15","SumCirLishold":"49.88"},{"SecurityCode":"000651.SZ","ChangeDate":"2019-09-30","TotalSh":"33.35万","TotalShChart":"333472","ChangeWithLastTermSh":"-16.37%","AvgShare":"1.79万","ChangeWithLastTermShare":"19.58%","StockPrice":"57.3","StockConvergenceRate":"非常集中","SumLishold":"47.83","SumCirLishold":"47.65"},{"SecurityCode":"000651.SZ","ChangeDate":"2019-06-30","TotalSh":"39.88万","TotalShChart":"398767","ChangeWithLastTermSh":"-8.97%","AvgShare":"1.497万","ChangeWithLastTermShare":"9.84%","StockPrice":"55","StockConvergenceRate":"非常集中","SumLishold":"46.63","SumCirLishold":"46.41"},{"SecurityCode":"000651.SZ","ChangeDate":"2019-03-31","TotalSh":"43.81万","TotalShChart":"438056","ChangeWithLastTermSh":"-20.58%","AvgShare":"1.363万","ChangeWithLastTermShare":"25.92%","StockPrice":"47.21","StockConvergenceRate":"非常集中","SumLishold":"48.21","SumCirLishold":"48.06"},{"SecurityCode":"000651.SZ","ChangeDate":"2018-12-31","TotalSh":"55.16万","TotalShChart":"551582","ChangeWithLastTermSh":"10.24%","AvgShare":"1.082万","ChangeWithLastTermShare":"-9.28%","StockPrice":"35.69","StockConvergenceRate":"非常集中","SumLishold":"44.98","SumCirLishold":"44.98"},{"SecurityCode":"000651.SZ","ChangeDate":"2018-09-30","TotalSh":"50.03万","TotalShChart":"500334","ChangeWithLastTermSh":"8.78%","AvgShare":"1.193万","ChangeWithLastTermShare":"-8.07%","StockPrice":"40.2","StockConvergenceRate":"非常集中","SumLishold":"44.87","SumCirLishold":"44.86"},{"SecurityCode":"000651.SZ","ChangeDate":"2018-06-30","TotalSh":"46万","TotalShChart":"459959","ChangeWithLastTermSh":"5.69%","AvgShare":"1.298万","ChangeWithLastTermShare":"-5.38%","StockPrice":"47.15","StockConvergenceRate":"非常集中","SumLishold":"45.57","SumCirLishold":"45.57"},{"SecurityCode":"000651.SZ","ChangeDate":"2018-03-31","TotalSh":"43.52万","TotalShChart":"435187","ChangeWithLastTermSh":"21.77%","AvgShare":"1.372万","ChangeWithLastTermShare":"-17.88%","StockPrice":"46.9","StockConvergenceRate":"非常集中","SumLishold":"46.49","SumCirLishold":"46.49"},{"SecurityCode":"000651.SZ","ChangeDate":"2017-12-31","TotalSh":"35.74万","TotalShChart":"357373","ChangeWithLastTermSh":"-1.70%","AvgShare":"1.67万","ChangeWithLastTermShare":"1.71%","StockPrice":"43.7","StockConvergenceRate":"非常集中","SumLishold":"46.09","SumCirLishold":"46.09"},{"SecurityCode":"000651.SZ","ChangeDate":"2017-09-30","TotalSh":"36.36万","TotalShChart":"363550","ChangeWithLastTermSh":"-1.76%","AvgShare":"1.642万","ChangeWithLastTermShare":"1.79%","StockPrice":"37.9","StockConvergenceRate":"非常集中","SumLishold":"44.76","SumCirLishold":"44.7"}],"GuDongTongJi":{"SecurityCode":"000651.SZ","ReportDate":"","ActualController":"无","ActualControllerRate":"--","ShareholderController":"无","ShareholderControllerRate":"--","ShareHolder":"合计持股","ShareHolderRate":"38.65%","CirShareHolder":"QFII加仓","CirShareHolderRate":"0.61%","OrgHolderNum":"1211","OrgHolderRate":"57.93%","OrgHolderNumPre":"552","OrgHolderRatePre":"56.13%"},"ZhuYingGouChengList":[{"SecurityCode":"000651.SZ","ReportType":"1","ReportDate":"2020-06-30","MainForm":"制造业","MainIncome":"497.1亿","MainIncomeRatio":"71.53%","MainIncomeRatioChart":"71.5254576335575","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"1","ReportDate":"2020-06-30","MainForm":"其他(补充)","MainIncome":"197.9亿","MainIncomeRatio":"28.47%","MainIncomeRatioChart":"28.4745423664425","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"2","ReportDate":"2020-06-30","MainForm":"内销","MainIncome":"378.2亿","MainIncomeRatio":"54.41%","MainIncomeRatioChart":"54.4133576643572","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"2","ReportDate":"2020-06-30","MainForm":"其他(补充)","MainIncome":"197.9亿","MainIncomeRatio":"28.47%","MainIncomeRatioChart":"28.4745423664425","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"2","ReportDate":"2020-06-30","MainForm":"外销","MainIncome":"118.9亿","MainIncomeRatio":"17.11%","MainIncomeRatioChart":"17.1120999692003","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"3","ReportDate":"2020-06-30","MainForm":"空调","MainIncome":"413.3亿","MainIncomeRatio":"59.47%","MainIncomeRatioChart":"59.4703629913621","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"3","ReportDate":"2020-06-30","MainForm":"其他(补充)","MainIncome":"197.9亿","MainIncomeRatio":"28.47%","MainIncomeRatioChart":"28.4745423664425","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"3","ReportDate":"2020-06-30","MainForm":"生活电器","MainIncome":"22.19亿","MainIncomeRatio":"3.19%","MainIncomeRatioChart":"3.19225347606455","Flag":"0"},{"SecurityCode":"000651.SZ","ReportType":"3","ReportDate":"2020-06-30","MainForm":"其他","MainIncome":"61.6亿","MainIncomeRatio":"8.86%","MainIncomeRatioChart":"8.86284116613083","Flag":"0"}],"FenHongSongZhuanList":[{"SecurityCode":"000651.SZ","NoticeDate":"2020-08-31","AssignDscrpt":"10派10元","ExDividendDate":"--","RightRegDate":"--"},{"SecurityCode":"000651.SZ","NoticeDate":"2020-06-04","AssignDscrpt":"10派12元","ExDividendDate":"2020-06-11","RightRegDate":"2020-06-10"},{"SecurityCode":"000651.SZ","NoticeDate":"2019-07-31","AssignDscrpt":"10派15元","ExDividendDate":"2019-08-06","RightRegDate":"2019-08-05"},{"SecurityCode":"000651.SZ","NoticeDate":"2019-02-19","AssignDscrpt":"10派6元","ExDividendDate":"2019-02-25","RightRegDate":"2019-02-22"},{"SecurityCode":"000651.SZ","NoticeDate":"2017-06-29","AssignDscrpt":"10派18元","ExDividendDate":"2017-07-05","RightRegDate":"2017-07-04"},{"SecurityCode":"000651.SZ","NoticeDate":"2016-07-01","AssignDscrpt":"10派15元","ExDividendDate":"2016-07-07","RightRegDate":"2016-07-06"},{"SecurityCode":"000651.SZ","NoticeDate":"2015-06-25","AssignDscrpt":"10转10派30元","ExDividendDate":"2015-07-03","RightRegDate":"2015-07-02"},{"SecurityCode":"000651.SZ","NoticeDate":"2014-05-28","AssignDscrpt":"10派15元","ExDividendDate":"2014-06-06","RightRegDate":"2014-06-05"},{"SecurityCode":"000651.SZ","NoticeDate":"2013-07-03","AssignDscrpt":"10派10元","ExDividendDate":"2013-07-11","RightRegDate":"2013-07-10"},{"SecurityCode":"000651.SZ","NoticeDate":"2012-06-29","AssignDscrpt":"10派5元","ExDividendDate":"2012-07-06","RightRegDate":"2012-07-05"}],"JingliRunUnit":"亿","YingYeShouRuUnit":"亿","DanJiCaiWuList":[{"SecurityCode":"000651.SZ","ReportDate":"2020-06-30","NetProfit":"48.04亿","NetProfitChart":"48.04","EPS":"0.8","EPSChart":"0.8","TotalIncome":"496.9亿","TotalIncomeChart":"496.93","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2020-03-31","NetProfit":"15.58亿","NetProfitChart":"15.58","EPS":"0.26","EPSChart":"0.26","TotalIncome":"209.1亿","TotalIncomeChart":"209.09","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2019-12-31","NetProfit":"25.79亿","NetProfitChart":"25.79","EPS":"0.43","EPSChart":"0.43","TotalIncome":"438.3亿","TotalIncomeChart":"438.32","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2019-09-30","NetProfit":"83.67亿","NetProfitChart":"83.67","EPS":"1.39","EPSChart":"1.39","TotalIncome":"583.4亿","TotalIncomeChart":"583.35","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2019-06-30","NetProfit":"80.78亿","NetProfitChart":"80.78","EPS":"1.35","EPSChart":"1.35","TotalIncome":"573.3亿","TotalIncomeChart":"573.35","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2019-03-31","NetProfit":"56.72亿","NetProfitChart":"56.72","EPS":"0.94","EPSChart":"0.94","TotalIncome":"410.1亿","TotalIncomeChart":"410.06","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2018-12-31","NetProfit":"50.84亿","NetProfitChart":"50.84","EPS":"0.85","EPSChart":"0.85","TotalIncome":"499.7亿","TotalIncomeChart":"499.74","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2018-09-30","NetProfit":"83.12亿","NetProfitChart":"83.12","EPS":"1.38","EPSChart":"1.38","TotalIncome":"580.5亿","TotalIncomeChart":"580.45","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2018-06-30","NetProfit":"72.25亿","NetProfitChart":"72.25","EPS":"1.2","EPSChart":"1.2","TotalIncome":"519.8亿","TotalIncomeChart":"519.79","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2018-03-31","NetProfit":"55.82亿","NetProfitChart":"55.82","EPS":"0.93","EPSChart":"0.93","TotalIncome":"400.3亿","TotalIncomeChart":"400.25","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2017-12-31","NetProfit":"69.4亿","NetProfitChart":"69.40","EPS":"1.15","EPSChart":"1.15","TotalIncome":"379.9亿","TotalIncomeChart":"379.93","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2017-09-30","NetProfit":"60.08亿","NetProfitChart":"60.08","EPS":"1","EPSChart":"1.0","TotalIncome":"420.1亿","TotalIncomeChart":"420.06","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2017-06-30","NetProfit":"54.38亿","NetProfitChart":"54.38","EPS":"0.9","EPSChart":"0.9","TotalIncome":"399.9亿","TotalIncomeChart":"399.85","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2017-03-31","NetProfit":"40.15亿","NetProfitChart":"40.15","EPS":"0.67","EPSChart":"0.67","TotalIncome":"300.4亿","TotalIncomeChart":"300.35","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2016-12-31","NetProfit":"42.35亿","NetProfitChart":"42.35","EPS":"0.7","EPSChart":"0.7","TotalIncome":"263.4亿","TotalIncomeChart":"263.42","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2016-09-30","NetProfit":"48.27亿","NetProfitChart":"48.27","EPS":"0.81","EPSChart":"0.81","TotalIncome":"336.8亿","TotalIncomeChart":"336.79","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2016-06-30","NetProfit":"32.43亿","NetProfitChart":"32.43","EPS":"0.53","EPSChart":"0.53","TotalIncome":"249.9亿","TotalIncomeChart":"249.93","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2016-03-31","NetProfit":"31.6亿","NetProfitChart":"31.60","EPS":"0.53","EPSChart":"0.53","TotalIncome":"251亿","TotalIncomeChart":"250.98","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2015-12-31","NetProfit":"25.8亿","NetProfitChart":"25.80","EPS":"0.43","EPSChart":"0.43","TotalIncome":"168.9亿","TotalIncomeChart":"168.89","Year":"2020"},{"SecurityCode":"000651.SZ","ReportDate":"2015-09-30","NetProfit":"42.32亿","NetProfitChart":"42.32","EPS":"0.7","EPSChart":"0.7","TotalIncome":"320.6亿","TotalIncomeChart":"320.65","Year":"2020"}],"GuQuanZhiYaTuJieList":[{"TradeDate":"2020-09-18","AmtShareRatio":"15.74"},{"TradeDate":"2020-09-11","AmtShareRatio":"15.74"},{"TradeDate":"2020-09-04","AmtShareRatio":"15.74"},{"TradeDate":"2020-08-28","AmtShareRatio":"15.74"},{"TradeDate":"2020-08-21","AmtShareRatio":"15.74"},{"TradeDate":"2020-08-14","AmtShareRatio":"15.74"},{"TradeDate":"2020-08-07","AmtShareRatio":"15.74"},{"TradeDate":"2020-07-31","AmtShareRatio":"15.74"},{"TradeDate":"2020-07-24","AmtShareRatio":"15.74"},{"TradeDate":"2020-07-17","AmtShareRatio":"15.74"},{"TradeDate":"2020-07-10","AmtShareRatio":"15.73"},{"TradeDate":"2020-07-03","AmtShareRatio":"15.73"},{"TradeDate":"2020-06-24","AmtShareRatio":"15.73"},{"TradeDate":"2020-06-19","AmtShareRatio":"15.74"},{"TradeDate":"2020-06-12","AmtShareRatio":"15.74"},{"TradeDate":"2020-06-05","AmtShareRatio":"15.74"},{"TradeDate":"2020-05-29","AmtShareRatio":"15.74"},{"TradeDate":"2020-05-22","AmtShareRatio":"15.74"},{"TradeDate":"2020-05-15","AmtShareRatio":"15.74"},{"TradeDate":"2020-05-08","AmtShareRatio":"15.74"},{"TradeDate":"2020-04-30","AmtShareRatio":"15.74"},{"TradeDate":"2020-04-24","AmtShareRatio":"15.74"},{"TradeDate":"2020-04-17","AmtShareRatio":"15.74"},{"TradeDate":"2020-04-10","AmtShareRatio":"15.74"},{"TradeDate":"2020-04-03","AmtShareRatio":"15.74"},{"TradeDate":"2020-03-27","AmtShareRatio":"15.74"},{"TradeDate":"2020-03-20","AmtShareRatio":"15.74"},{"TradeDate":"2020-03-13","AmtShareRatio":"15.74"},{"TradeDate":"2020-03-06","AmtShareRatio":"15.57"},{"TradeDate":"2020-02-28","AmtShareRatio":"15.74"},{"TradeDate":"2020-02-21","AmtShareRatio":"15.74"},{"TradeDate":"2020-02-14","AmtShareRatio":"15.74"},{"TradeDate":"2020-02-07","AmtShareRatio":"0.74"},{"TradeDate":"2020-01-17","AmtShareRatio":"0.74"},{"TradeDate":"2020-01-10","AmtShareRatio":"0.74"},{"TradeDate":"2020-01-03","AmtShareRatio":"0.74"},{"TradeDate":"2019-12-27","AmtShareRatio":"0.74"},{"TradeDate":"2019-12-20","AmtShareRatio":"0.74"},{"TradeDate":"2019-12-13","AmtShareRatio":"0.74"},{"TradeDate":"2019-12-06","AmtShareRatio":"0.74"},{"TradeDate":"2019-11-29","AmtShareRatio":"0.74"},{"TradeDate":"2019-11-22","AmtShareRatio":"0.74"},{"TradeDate":"2019-11-15","AmtShareRatio":"0.74"},{"TradeDate":"2019-11-08","AmtShareRatio":"0.74"},{"TradeDate":"2019-11-01","AmtShareRatio":"0.74"},{"TradeDate":"2019-10-25","AmtShareRatio":"0.74"},{"TradeDate":"2019-10-18","AmtShareRatio":"0.68"},{"TradeDate":"2019-10-11","AmtShareRatio":"0.70"},{"TradeDate":"2019-09-30","AmtShareRatio":"0.73"},{"TradeDate":"2019-09-27","AmtShareRatio":"0.73"}],"GuQuanZhiYaTongJiList":[{"SecurityCode":"000651","SecurityShortName":"格力电器","ShareHolderCode":"81040911","ShareHolderName":"珠海明骏投资合伙企业(有限合伙)","BBallowance":"9.02亿","HolderRatio":"100.00%","TotalShareRatio":"15.00%","NoticeDate":"2020-02-17","TotalCount":"1"}],"ShangYuQingKuang":{"TuJieDataList":[{"BaoGaoQi":"2020","ShangYuZhi":"325919390.58","ShangYuZhiText":null,"ShangYuZhiUnit":null,"JingZiChanJiaZong":null,"ShangYuJingZiChan":"0.29%"},{"BaoGaoQi":"2019","ShangYuZhi":"325919390.58","ShangYuZhiText":null,"ShangYuZhiUnit":null,"JingZiChanJiaZong":null,"ShangYuJingZiChan":"0.30%"},{"BaoGaoQi":"2018","ShangYuZhi":"51804350.47","ShangYuZhiText":null,"ShangYuZhiUnit":null,"JingZiChanJiaZong":null,"ShangYuJingZiChan":"0.06%"},{"BaoGaoQi":"2017","ShangYuZhi":"","ShangYuZhiText":null,"ShangYuZhiUnit":null,"JingZiChanJiaZong":null,"ShangYuJingZiChan":"--"},{"BaoGaoQi":"2016","ShangYuZhi":"","ShangYuZhiText":null,"ShangYuZhiUnit":null,"JingZiChanJiaZong":null,"ShangYuJingZiChan":"--"}],"Data":{"Code":"000651.SZ","BaoGaoQi":"2020中报","BaoGaoQiLeiXing":"Q6","ShangYu":"3.26亿","ShangYuTongBi":"-9.74%","ShangYuJingZiChan":"0.29%","ShangYuJingZiChanTongBi":"-26.56%","GuiMuJingLiRun":"63.62亿","GuiMuJingLiRunTongBi":"-53.73%","ShiFouXianShi":null}},"YanFaTouRuList":[]},"Status":0,"Message":null,"OtherInfo":{}}

//https://emh5.eastmoney.com/api/CaiWuFenXi/GetLiRunBiaoList
//{"Result":{"LiRunBiaoList_QuanShang":null,"LiRunBiaoList_BaoXian":null,"LiRunBiaoList_QiYe":[{"ReportDate":"2020-06-30","Title":["2020二季度","同比"],"Totaloperatereve":["496.9亿","137.67%"],"Operatereve":["491.1亿","140.77%"],"Intreve":["5.862亿","14.56%"],"Premiumearned":["--","--"],"Commreve":["1.476万","-98.95%"],"Otherreve":["--","--"],"Totaloperatereveother":["--","--"],"Totaloperateexp":["443.1亿","133.89%"],"Operateexp":["380亿","125.80%"],"Intexp":["810.3万","-7.43%"],"Commexp":["18.84万","213.91%"],"Rdexp":["15.7亿","72.99%"],"Surrenderpremium":["--","--"],"Netindemnityexp":["--","--"],"Netcontactreserve":["--","--"],"Policydiviexp":["--","--"],"Riexp":["--","--"],"Otherexp":["--","--"],"Operatetax":["1.901亿","30.02%"],"Saleexp":["43.45亿","379.14%"],"Manageexp":["8.79亿","36.62%"],"Financeexp":["-6.833亿","-37.33%"],"Ofwintexp":["--","--"],"Ofwintreve":["--","--"],"Assetdevalueloss":["--","--"],"Creddevalueloss":["--","--"],"Totaloperateexpother":["--","--"],"Fvalueincome":["9112万","147.88%"],"Investincome":["4112万","-34.57%"],"Investjointincome":["965.1万","358.70%"],"Netexhedgincome":["--","--"],"Exchangeincome":["--","--"],"Adisposalincome":["--","--"],"Miotherincome":["--","--"],"Operateprofitother":["--","--"],"Operateprofitbalance":["--","--"],"Operateprofit":["57.56亿","215.23%"],"Nonoperatereve":["5157万","-31.46%"],"Nonlassetreve":["--","--"],"Nonoperateexp":["781.9万","56.29%"],"Nonlassetnetloss":["--","--"],"Sumprofitbalance":["--","--"],"Sumprofit":["58亿","205.86%"],"Incometax":["9.522亿","199.12%"],"Sumprofitother":["--","--"],"Unconfirminvloss":["--","--"],"Netprofit":["48.47亿","207.22%"],"Combinednetprofitb":["--","--"],"Continuousonprofit":["--","--"],"Terminationonprofit":["--","--"],"Parentnetprofit":["48.04亿","208.35%"],"Minorityincome":["4329万","118.22%"],"Kcfjcxsyjlr":["45.67亿","219.46%"],"Netprofitother2":["--","--"],"Netprofitbalance":["--","--"],"Basiceps":["0.8000","207.69%"],"Dilutedeps":["0.8000","207.69%"],"Othercincome":["30.83亿","150.12%"],"Parentothercincome":["30.84亿","150.26%"],"Minorityothercincome":["-22.06万","-135.67%"],"Sumcincome":["79.31亿","182.17%"],"Parentcincome":["78.88亿","182.70%"],"Minoritycincome":["4307万","110.54%"],"Auditopinionsdomestic":["--","--"]},{"ReportDate":"2020-03-31","Title":["2020一季度","同比"],"Totaloperatereve":["209.1亿","-52.30%"],"Operatereve":["204亿","-52.69%"],"Intreve":["5.117亿","-28.70%"],"Premiumearned":["--","--"],"Commreve":["140.8万","2794.73%"],"Otherreve":["--","--"],"Totaloperatereveother":["--","--"],"Totaloperateexp":["189.4亿","-52.66%"],"Operateexp":["168.3亿","-52.23%"],"Intexp":["875.4万","-65.27%"],"Commexp":["6万","-28.43%"],"Rdexp":["9.075亿","-28.06%"],"Surrenderpremium":["--","--"],"Netindemnityexp":["--","--"],"Netcontactreserve":["--","--"],"Policydiviexp":["--","--"],"Riexp":["--","--"],"Otherexp":["--","--"],"Operatetax":["1.462亿","-54.56%"],"Saleexp":["9.068亿","-71.18%"],"Manageexp":["6.434亿","-28.47%"],"Financeexp":["-4.975亿","42.10%"],"Ofwintexp":["--","--"],"Ofwintreve":["--","--"],"Assetdevalueloss":["--","--"],"Creddevalueloss":["--","--"],"Totaloperateexpother":["--","--"],"Fvalueincome":["-1.903亿","-152.15%"],"Investincome":["6285万","127.78%"],"Investjointincome":["210.4万","-51.20%"],"Netexhedgincome":["--","--"],"Exchangeincome":["--","--"],"Adisposalincome":["--","--"],"Miotherincome":["--","--"],"Operateprofitother":["--","--"],"Operateprofitbalance":["--","--"],"Operateprofit":["18.26亿","-44.42%"],"Nonoperatereve":["7524万","2431.89%"],"Nonlassetreve":["--","--"],"Nonoperateexp":["500.3万","-99.14%"],"Nonlassetnetloss":["--","--"],"Sumprofitbalance":["--","--"],"Sumprofit":["18.96亿","-29.95%"],"Incometax":["3.183亿","226.69%"],"Sumprofitother":["--","--"],"Unconfirminvloss":["--","--"],"Netprofit":["15.78亿","-39.54%"],"Combinednetprofitb":["--","--"],"Continuousonprofit":["--","--"],"Terminationonprofit":["--","--"],"Parentnetprofit":["15.58亿","-39.59%"],"Minorityincome":["1984万","-34.80%"],"Kcfjcxsyjlr":["14.3亿","-45.14%"],"Netprofitother2":["--","--"],"Netprofitbalance":["--","--"],"Basiceps":["0.2600","-39.53%"],"Dilutedeps":["0.2600","-39.53%"],"Othercincome":["12.33亿","-82.13%"],"Parentothercincome":["12.32亿","-82.14%"],"Minorityothercincome":["61.84万","1005.88%"],"Sumcincome":["28.11亿","-70.44%"],"Parentcincome":["27.9亿","-70.56%"],"Minoritycincome":["2046万","-32.62%"],"Auditopinionsdomestic":["--","--"]},{"ReportDate":"2019-12-31","Title":["2019四季度","同比"],"Totaloperatereve":["438.3亿","-24.86%"],"Operatereve":["431.1亿","-25.33%"],"Intreve":["7.177亿","21.00%"],"Premiumearned":["--","--"],"Commreve":["4.864万","-81.06%"],"Otherreve":["--","--"],"Totaloperatereveother":["--","--"],"Totaloperateexp":["400.2亿","-17.56%"],"Operateexp":["352.2亿","-14.43%"],"Intexp":["2520万","-32.62%"],"Commexp":["8.384万","-21.78%"],"Rdexp":["12.61亿","-21.59%"],"Surrenderpremium":["--","--"],"Netindemnityexp":["--","--"],"Netcontactreserve":["--","--"],"Policydiviexp":["--","--"],"Riexp":["--","--"],"Otherexp":["--","--"],"Operatetax":["3.218亿","-20.42%"],"Saleexp":["31.46亿","-33.77%"],"Manageexp":["8.995亿","-12.67%"],"Financeexp":["-8.593亿","-91.05%"],"Ofwintexp":["--","--"],"Ofwintreve":["--","--"],"Assetdevalueloss":["--","--"],"Creddevalueloss":["--","--"],"Totaloperateexpother":["--","--"],"Fvalueincome":["3.649亿","5682.89%"],"Investincome":["-2.263亿","-71.20%"],"Investjointincome":["431.1万","130.46%"],"Netexhedgincome":["--","--"],"Exchangeincome":["--","--"],"Adisposalincome":["--","--"],"Miotherincome":["--","--"],"Operateprofitother":["--","--"],"Operateprofitbalance":["--","--"],"Operateprofit":["32.85亿","-66.90%"],"Nonoperatereve":["297.2万","-99.04%"],"Nonlassetreve":["--","--"],"Nonoperateexp":["5.811亿","12178.33%"],"Nonlassetnetloss":["--","--"],"Sumprofitbalance":["--","--"],"Sumprofit":["27.07亿","-73.54%"],"Incometax":["9744万","-94.68%"],"Sumprofitother":["--","--"],"Unconfirminvloss":["--","--"],"Netprofit":["26.1亿","-68.94%"],"Combinednetprofitb":["--","--"],"Continuousonprofit":["--","--"],"Terminationonprofit":["--","--"],"Parentnetprofit":["25.79亿","-69.18%"],"Minorityincome":["3043万","-12.19%"],"Kcfjcxsyjlr":["26.06亿","-68.11%"],"Netprofitother2":["--","--"],"Netprofitbalance":["--","--"],"Basiceps":["0.4300","-69.06%"],"Dilutedeps":["0.4300","-69.06%"],"Othercincome":["69亿","41245.60%"],"Parentothercincome":["69亿","41572.67%"],"Minorityothercincome":["-6.827万","48.32%"],"Sumcincome":["95.09亿","13.41%"],"Parentcincome":["94.79亿","13.51%"],"Minoritycincome":["3036万","-12.05%"],"Auditopinionsdomestic":["--","--"]},{"ReportDate":"2019-09-30","Title":["2019三季度","同比"],"Totaloperatereve":["583.4亿","1.74%"],"Operatereve":["577.4亿","1.75%"],"Intreve":["5.932亿","1.75%"],"Premiumearned":["--","--"],"Commreve":["25.69万","-90.52%"],"Otherreve":["--","--"],"Totaloperatereveother":["--","--"],"Totaloperateexp":["485.4亿","2.39%"],"Operateexp":["411.6亿","5.63%"],"Intexp":["3740万","146.71%"],"Commexp":["10.72万","-53.60%"],"Rdexp":["16.09亿","-9.91%"],"Surrenderpremium":["--","--"],"Netindemnityexp":["--","--"],"Netcontactreserve":["--","--"],"Policydiviexp":["--","--"],"Riexp":["--","--"],"Otherexp":["--","--"],"Operatetax":["4.044亿","-20.90%"],"Saleexp":["47.51亿","-23.89%"],"Manageexp":["10.3亿","4.88%"],"Financeexp":["-4.498亿","58.91%"],"Ofwintexp":["--","--"],"Ofwintreve":["--","--"],"Assetdevalueloss":["--","--"],"Creddevalueloss":["--","--"],"Totaloperateexpother":["--","--"],"Fvalueincome":["-653.6万","97.24%"],"Investincome":["-1.322亿","-54.50%"],"Investjointincome":["-1415万","-1694.68%"],"Netexhedgincome":["--","--"],"Exchangeincome":["--","--"],"Adisposalincome":["--","--"],"Miotherincome":["--","--"],"Operateprofitother":["--","--"],"Operateprofitbalance":["--","--"],"Operateprofit":["99.26亿","1.42%"],"Nonoperatereve":["3.111亿","1308.98%"],"Nonlassetreve":["--","--"],"Nonoperateexp":["473.2万","-41.51%"],"Nonlassetnetloss":["--","--"],"Sumprofitbalance":["--","--"],"Sumprofit":["102.3亿","4.40%"],"Incometax":["18.3亿","8.73%"],"Sumprofitother":["--","--"],"Unconfirminvloss":["--","--"],"Netprofit":["84.02亿","3.50%"],"Combinednetprofitb":["--","--"],"Continuousonprofit":["--","--"],"Terminationonprofit":["--","--"],"Parentnetprofit":["83.67亿","3.58%"],"Minorityincome":["3466万","-12.68%"],"Kcfjcxsyjlr":["81.72亿","-1.30%"],"Netprofitother2":["--","--"],"Netprofitbalance":["--","--"],"Basiceps":["1.3900","2.96%"],"Dilutedeps":["1.3900","2.96%"],"Othercincome":["-1677万","83.76%"],"Parentothercincome":["-1664万","83.85%"],"Minorityothercincome":["-13.21万","44.07%"],"Sumcincome":["83.85亿","4.62%"],"Parentcincome":["83.51亿","4.71%"],"Minoritycincome":["3452万","-12.49%"],"Auditopinionsdomestic":["--","--"]},{"ReportDate":"2019-06-30","Title":["2019二季度","同比"],"Totaloperatereve":["573.3亿","39.82%"],"Operatereve":["567.5亿","39.96%"],"Intreve":["5.83亿","27.39%"],"Premiumearned":["--","--"],"Commreve":["271.1万","231.33%"],"Otherreve":["--","--"],"Totaloperatereveother":["--","--"],"Totaloperateexp":["474.1亿","36.43%"],"Operateexp":["389.7亿","38.46%"],"Intexp":["1516万","-53.79%"],"Commexp":["23.1万","27.34%"],"Rdexp":["17.86亿","44.57%"],"Surrenderpremium":["--","--"],"Netindemnityexp":["--","--"],"Netcontactreserve":["--","--"],"Policydiviexp":["--","--"],"Riexp":["--","--"],"Otherexp":["--","--"],"Operatetax":["5.113亿","67.35%"],"Saleexp":["62.42亿","49.66%"],"Manageexp":["9.82亿","11.05%"],"Financeexp":["-10.95亿","-4659.37%"],"Ofwintexp":["--","--"],"Ofwintreve":["--","--"],"Assetdevalueloss":["--","--"],"Creddevalueloss":["--","--"],"Totaloperateexpother":["--","--"],"Fvalueincome":["-2.37亿","-321.70%"],"Investincome":["-8554万","-139.36%"],"Investjointincome":["-78.86万","92.38%"],"Netexhedgincome":["--","--"],"Exchangeincome":["--","--"],"Adisposalincome":["--","--"],"Miotherincome":["--","--"],"Operateprofitother":["--","--"],"Operateprofitbalance":["--","--"],"Operateprofit":["97.87亿","48.15%"],"Nonoperatereve":["2208万","132.36%"],"Nonlassetreve":["--","--"],"Nonoperateexp":["809.2万","92.03%"],"Nonlassetnetloss":["--","--"],"Sumprofitbalance":["--","--"],"Sumprofit":["98.01亿","48.24%"],"Incometax":["16.83亿","84.16%"],"Sumprofitother":["--","--"],"Unconfirminvloss":["--","--"],"Netprofit":["81.18亿","42.48%"],"Combinednetprofitb":["--","--"],"Continuousonprofit":["--","--"],"Terminationonprofit":["--","--"],"Parentnetprofit":["80.78亿","42.43%"],"Minorityincome":["3969万","53.65%"],"Kcfjcxsyjlr":["82.8亿","61.91%"],"Netprofitother2":["--","--"],"Netprofitbalance":["--","--"],"Basiceps":["1.3500","43.62%"],"Dilutedeps":["1.3500","43.62%"],"Othercincome":["-1.033亿","-202.75%"],"Parentothercincome":["-1.03亿","-202.56%"],"Minorityothercincome":["-23.61万","-675.03%"],"Sumcincome":["80.15亿","38.23%"],"Parentcincome":["79.75亿","38.16%"],"Minoritycincome":["3945万","52.49%"],"Auditopinionsdomestic":["--","--"]}],"LiRunBiaoList_YinHang":null},"Status":0,"Message":"","OtherInfo":{}}
