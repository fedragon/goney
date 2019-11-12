package goney

// Currency represents an ISO4217 currency
type Currency struct {
	Code      string
	Precision int32
}

func (c Currency) String() string {
	return c.Code
}

// Find returns the currency having provided code, if found,
// or XXX (= no currency) if not found
func Find(code string) Currency {
	for _, curr := range currencies {
		if curr.Code == code {
			return curr
		}
	}

	return XXX
}

var currencies = []Currency{
	AED, AFN, ALL, AMD, ANG, AOA, ARS, AUD, AWG, AZN, BAM, BBD, BDT, BGN, BHD, BIF, BMD, BND, BOB, BOV, BRL, BSD, BTN, BWP, BYN, BZD, CAD, CDF, CHE, CHF, CHW, CLF, CLP, CNY, COP, COU, CRC, CUC, CUP, CVE, CZK, DJF, DKK, DOP, DZD, EGP, ERN, ETB, EUR, FJD, FKP, GBP, GEL, GHS, GIP, GMD, GNF, GTQ, GYD, HKD, HNL, HRK, HTG, HUF, IDR, ILS, INR, IQD, IRR, ISK, JMD, JOD, JPY, KES, KGS, KHR, KMF, KPW, KRW, KWD, KYD, KZT, LAK, LBP, LKR, LRD, LSL, LYD, MAD, MDL, MGA, MKD, MMK, MNT, MOP, MRU, MUR, MVR, MWK, MXN, MXV, MYR, MZN, NAD, NGN, NIO, NOK, NPR, NZD, OMR, PAB, PEN, PGK, PHP, PKR, PLN, PYG, QAR, RON, RSD, RUB, RWF, SAR, SBD, SCR, SDG, SEK, SGD, SHP, SLL, SOS, SRD, SSP, STN, SVC, SYP, SZL, THB, TJS, TMT, TND, TOP, TRY, TTD, TWD, TZS, UAH, UGX, USD, USN, UYI, UYU, UYW, UZS, VES, VND, VUV, WST, XAF, XAG, XAU, XBA, XBB, XBC, XBD, XCD, XDR, XOF, XPD, XPF, XPT, XSU, XTS, XUA, XXX, YER, ZAR, ZMW, ZWL,
}

// AED United Arab Emirates dirham
var AED = Currency{"AED", 2}

// AFN Afghan afghani
var AFN = Currency{"AFN", 2}

// ALL Albanian lek
var ALL = Currency{"ALL", 2}

// AMD Armenian dram
var AMD = Currency{"AMD", 2}

// ANG Netherlands Antillean guilder
var ANG = Currency{"ANG", 2}

// AOA Angolan kwanza
var AOA = Currency{"AOA", 2}

// ARS Argentine peso
var ARS = Currency{"ARS", 2}

// AUD Australian dollar
var AUD = Currency{"AUD", 2}

// AWG Aruban florin
var AWG = Currency{"AWG", 2}

// AZN Azerbaijani manat
var AZN = Currency{"AZN", 2}

// BAM Bosnia and Herzegovina convertible mark
var BAM = Currency{"BAM", 2}

// BBD Barbados dollar
var BBD = Currency{"BBD", 2}

// BDT Bangladeshi taka
var BDT = Currency{"BDT", 2}

// BGN Bulgarian lev
var BGN = Currency{"BGN", 2}

// BHD Bahraini dinar
var BHD = Currency{"BHD", 3}

// BIF Burundian franc
var BIF = Currency{"BIF", 0}

// BMD Bermudian dollar
var BMD = Currency{"BMD", 2}

// BND Brunei dollar
var BND = Currency{"BND", 2}

// BOB Boliviano
var BOB = Currency{"BOB", 2}

// BOV Bolivian Mvdol (funds code)
var BOV = Currency{"BOV", 2}

// BRL Brazilian real
var BRL = Currency{"BRL", 2}

// BSD Bahamian dollar
var BSD = Currency{"BSD", 2}

// BTN Bhutanese ngultrum
var BTN = Currency{"BTN", 2}

// BWP Botswana pula
var BWP = Currency{"BWP", 2}

// BYN Belarusian ruble
var BYN = Currency{"BYN", 2}

// BZD Belize dollar
var BZD = Currency{"BZD", 2}

// CAD Canadian dollar
var CAD = Currency{"CAD", 2}

// CDF Congolese franc
var CDF = Currency{"CDF", 2}

// CHE WIR Euro (complementary currency)
var CHE = Currency{"CHE", 2}

// CHF Swiss franc
var CHF = Currency{"CHF", 2}

// CHW WIR Franc (complementary currency)
var CHW = Currency{"CHW", 2}

// CLF Unidad de Fomento (funds code)
var CLF = Currency{"CLF", 4}

// CLP Chilean peso
var CLP = Currency{"CLP", 0}

// CNY Renminbi (Chinese) yuan
var CNY = Currency{"CNY", 2}

// COP Colombian peso
var COP = Currency{"COP", 2}

// COU Unidad de Valor Real (UVR) (funds code)
var COU = Currency{"COU", 2}

// CRC Costa Rican colon
var CRC = Currency{"CRC", 2}

// CUC Cuban convertible peso
var CUC = Currency{"CUC", 2}

// CUP Cuban peso
var CUP = Currency{"CUP", 2}

// CVE Cape Verdean escudo
var CVE = Currency{"CVE", 2}

// CZK Czech koruna
var CZK = Currency{"CZK", 2}

// DJF Djiboutian franc
var DJF = Currency{"DJF", 0}

// DKK Danish krone
var DKK = Currency{"DKK", 2}

// DOP Dominican peso
var DOP = Currency{"DOP", 2}

// DZD Algerian dinar
var DZD = Currency{"DZD", 2}

// EGP Egyptian pound
var EGP = Currency{"EGP", 2}

// ERN Eritrean nakfa
var ERN = Currency{"ERN", 2}

// ETB Ethiopian birr
var ETB = Currency{"ETB", 2}

// EUR Euro
var EUR = Currency{"EUR", 2}

// FJD Fiji dollar
var FJD = Currency{"FJD", 2}

// FKP Falkland Islands pound
var FKP = Currency{"FKP", 2}

// GBP Pound sterling
var GBP = Currency{"GBP", 2}

// GEL Georgian lari
var GEL = Currency{"GEL", 2}

// GHS Ghanaian cedi
var GHS = Currency{"GHS", 2}

// GIP Gibraltar pound
var GIP = Currency{"GIP", 2}

// GMD Gambian dalasi
var GMD = Currency{"GMD", 2}

// GNF Guinean franc
var GNF = Currency{"GNF", 0}

// GTQ Guatemalan quetzal
var GTQ = Currency{"GTQ", 2}

// GYD Guyanese dollar
var GYD = Currency{"GYD", 2}

// HKD Hong Kong dollar
var HKD = Currency{"HKD", 2}

// HNL Honduran lempira
var HNL = Currency{"HNL", 2}

// HRK Croatian kuna
var HRK = Currency{"HRK", 2}

// HTG Haitian gourde
var HTG = Currency{"HTG", 2}

// HUF Hungarian forint
var HUF = Currency{"HUF", 2}

// IDR Indonesian rupiah
var IDR = Currency{"IDR", 2}

// ILS Israeli new shekel
var ILS = Currency{"ILS", 2}

// INR Indian rupee
var INR = Currency{"INR", 2}

// IQD Iraqi dinar
var IQD = Currency{"IQD", 3}

// IRR Iranian rial
var IRR = Currency{"IRR", 2}

// ISK Icelandic króna
var ISK = Currency{"ISK", 0}

// JMD Jamaican dollar
var JMD = Currency{"JMD", 2}

// JOD Jordanian dinar
var JOD = Currency{"JOD", 3}

// JPY Japanese yen
var JPY = Currency{"JPY", 0}

// KES Kenyan shilling
var KES = Currency{"KES", 2}

// KGS Kyrgyzstani som
var KGS = Currency{"KGS", 2}

// KHR Cambodian riel
var KHR = Currency{"KHR", 2}

// KMF Comoro franc
var KMF = Currency{"KMF", 0}

// KPW North Korean won
var KPW = Currency{"KPW", 2}

// KRW South Korean won
var KRW = Currency{"KRW", 0}

// KWD Kuwaiti dinar
var KWD = Currency{"KWD", 3}

// KYD Cayman Islands dollar
var KYD = Currency{"KYD", 2}

// KZT Kazakhstani tenge
var KZT = Currency{"KZT", 2}

// LAK Lao kip
var LAK = Currency{"LAK", 2}

// LBP Lebanese pound
var LBP = Currency{"LBP", 2}

// LKR Sri Lankan rupee
var LKR = Currency{"LKR", 2}

// LRD Liberian dollar
var LRD = Currency{"LRD", 2}

// LSL Lesotho loti
var LSL = Currency{"LSL", 2}

// LYD Libyan dinar
var LYD = Currency{"LYD", 3}

// MAD Moroccan dirham
var MAD = Currency{"MAD", 2}

// MDL Moldovan leu
var MDL = Currency{"MDL", 2}

// MGA 11] 	Malagasy ariary
var MGA = Currency{"MGA", 2}

// MKD Macedonian denar
var MKD = Currency{"MKD", 2}

// MMK Myanmar kyat
var MMK = Currency{"MMK", 2}

// MNT Mongolian tögrög
var MNT = Currency{"MNT", 2}

// MOP Macanese pataca
var MOP = Currency{"MOP", 2}

// MRU Mauritanian ouguiya
var MRU = Currency{"MRU", 2}

// MUR Mauritian rupee
var MUR = Currency{"MUR", 2}

// MVR Maldivian rufiyaa
var MVR = Currency{"MVR", 2}

// MWK Malawian kwacha
var MWK = Currency{"MWK", 2}

// MXN Mexican peso
var MXN = Currency{"MXN", 2}

// MXV Mexican Unidad de Inversion (UDI) (funds code)
var MXV = Currency{"MXV", 2}

// MYR Malaysian ringgit
var MYR = Currency{"MYR", 2}

// MZN Mozambican metical
var MZN = Currency{"MZN", 2}

// NAD Namibian dollar
var NAD = Currency{"NAD", 2}

// NGN Nigerian naira
var NGN = Currency{"NGN", 2}

// NIO Nicaraguan córdoba
var NIO = Currency{"NIO", 2}

// NOK Norwegian krone
var NOK = Currency{"NOK", 2}

// NPR Nepalese rupee
var NPR = Currency{"NPR", 2}

// NZD New Zealand dollar
var NZD = Currency{"NZD", 2}

// OMR Omani rial
var OMR = Currency{"OMR", 3}

// PAB Panamanian balboa
var PAB = Currency{"PAB", 2}

// PEN Peruvian sol
var PEN = Currency{"PEN", 2}

// PGK Papua New Guinean kina
var PGK = Currency{"PGK", 2}

// PHP Philippine peso
var PHP = Currency{"PHP", 2}

// PKR Pakistani rupee
var PKR = Currency{"PKR", 2}

// PLN Polish złoty
var PLN = Currency{"PLN", 2}

// PYG Paraguayan guaraní
var PYG = Currency{"PYG", 0}

// QAR Qatari riyal
var QAR = Currency{"QAR", 2}

// RON Romanian leu
var RON = Currency{"RON", 2}

// RSD Serbian dinar
var RSD = Currency{"RSD", 2}

// RUB Russian ruble
var RUB = Currency{"RUB", 2}

// RWF Rwandan franc
var RWF = Currency{"RWF", 0}

// SAR Saudi riyal
var SAR = Currency{"SAR", 2}

// SBD Solomon Islands dollar
var SBD = Currency{"SBD", 2}

// SCR Seychelles rupee
var SCR = Currency{"SCR", 2}

// SDG Sudanese pound
var SDG = Currency{"SDG", 2}

// SEK Swedish krona/kronor
var SEK = Currency{"SEK", 2}

// SGD Singapore dollar
var SGD = Currency{"SGD", 2}

// SHP Saint Helena pound
var SHP = Currency{"SHP", 2}

// SLL Sierra Leonean leone
var SLL = Currency{"SLL", 2}

// SOS Somali shilling
var SOS = Currency{"SOS", 2}

// SRD Surinamese dollar
var SRD = Currency{"SRD", 2}

// SSP South Sudanese pound
var SSP = Currency{"SSP", 2}

// STN São Tomé and Príncipe dobra
var STN = Currency{"STN", 2}

// SVC Salvadoran colón
var SVC = Currency{"SVC", 2}

// SYP Syrian pound
var SYP = Currency{"SYP", 2}

// SZL Swazi lilangeni
var SZL = Currency{"SZL", 2}

// THB Thai baht
var THB = Currency{"THB", 2}

// TJS Tajikistani somoni
var TJS = Currency{"TJS", 2}

// TMT Turkmenistan manat
var TMT = Currency{"TMT", 2}

// TND Tunisian dinar
var TND = Currency{"TND", 3}

// TOP Tongan paʻanga
var TOP = Currency{"TOP", 2}

// TRY Turkish lira
var TRY = Currency{"TRY", 2}

// TTD Trinidad and Tobago dollar
var TTD = Currency{"TTD", 2}

// TWD New Taiwan dollar
var TWD = Currency{"TWD", 2}

// TZS Tanzanian shilling
var TZS = Currency{"TZS", 2}

// UAH Ukrainian hryvnia
var UAH = Currency{"UAH", 2}

// UGX Ugandan shilling
var UGX = Currency{"UGX", 0}

// USD United States dollar
var USD = Currency{"USD", 2}

// USN United States dollar (next day) (funds code)
var USN = Currency{"USN", 2}

// UYI Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)
var UYI = Currency{"UYI", 0}

// UYU Uruguayan peso
var UYU = Currency{"UYU", 2}

// UYW Unidad previsional[15]
var UYW = Currency{"UYW", 4}

// UZS Uzbekistan som
var UZS = Currency{"UZS", 2}

// VES Venezuelan bolívar soberano
var VES = Currency{"VES", 2}

// VND Vietnamese đồng
var VND = Currency{"VND", 0}

// VUV Vanuatu vatu
var VUV = Currency{"VUV", 0}

// WST Samoan tala
var WST = Currency{"WST", 2}

// XAF CFA franc BEAC
var XAF = Currency{"XAF", 0}

// XAG Silver (one troy
var XAG = Currency{"XAG", 0}

// XAU Gold (one troy
var XAU = Currency{"XAU", 0}

// XBA European Composite Unit (EURCO) (bond market)
var XBA = Currency{"XBA", 0}

// XBB European Monetary Unit (E.M.U.-6) (bond market)
var XBB = Currency{"XBB", 0}

// XBC European Unit of Account 9 (E.U.A.-9) (bond market)
var XBC = Currency{"XBC", 0}

// XBD European Unit of Account 17 (E.U.A.-17) (bond market)
var XBD = Currency{"XBD", 0}

// XCD East Caribbean dollar
var XCD = Currency{"XCD", 2}

// XDR Special drawing rights
var XDR = Currency{"XDR", 0}

// XOF CFA franc BCEAO
var XOF = Currency{"XOF", 0}

// XPD Palladium
var XPD = Currency{"XPD", 0}

// XPF CFP franc
var XPF = Currency{"XPF", 0}

// XPT Platinum
var XPT = Currency{"XPT", 0}

// XSU SUCRE
var XSU = Currency{"XSU", 0}

// XTS Code reserved for testing
var XTS = Currency{"XTS", 0}

// XUA ADB Unit of Account
var XUA = Currency{"XUA", 0}

// XXX No currency
var XXX = Currency{"XXX", 0}

// YER Yemeni rial
var YER = Currency{"YER", 2}

// ZAR South African rand
var ZAR = Currency{"ZAR", 2}

// ZMW Zambian kwacha
var ZMW = Currency{"ZMW", 2}

// ZWL Zimbabwean dollar
var ZWL = Currency{"ZWL", 2}
