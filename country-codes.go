package countrycodes

import (
	"github.com/tchap/go-patricia/patricia"
	"strings"
)

type Assignment int

const (
	/**
	 * <a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements"
	 * >Officially assigned</a>.
	 *
	 * Assigned to a country, territory, or area of geographical interest.
	 */
	OFFICIALLY_ASSIGNED Assignment = iota

	/**
	 * <a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#User-assigned_code_elements"
	 * >User assigned</a>.
	 *
	 * Free for assignment at the disposal of users.
	 */
	USER_ASSIGNED

	/**
	 * <a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Exceptional_reservations"
	 * >Exceptionally reserved</a>.
	 *
	 * Reserved on request for restricted use.
	 */
	EXCEPTIONALLY_RESERVED

	/**
	 * <a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Transitional_reservations"
	 * >Transitionally reserved</a>.
	 *
	 * Deleted from ISO 3166-1 but reserved transitionally.
	 */
	TRANSITIONALLY_RESERVED

	/**
	 * <a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Indeterminate_reservations"
	 * >Indeterminately reserved</a>.
	 *
	 * Used in coding systems associated with ISO 3166-1.
	 */
	INDETERMINATELY_RESERVED

	/**
	 * <a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Codes_currently_agreed_not_to_use"
	 * >Not used</a>.
	 *
	 * Not used in ISO 3166-1 in deference to international property
	 * organization names.
	 */
	NOT_USED
)

type CountryCode struct {
	Name        string
	Alpha2      string
	Alpha3      string
	Numeric     int
	DialingCode string
	Assignment  Assignment
}

var by_alpha2 map[string]CountryCode

var by_name map[string]CountryCode

var by_alpha3 map[string]CountryCode

var by_numeric map[int]CountryCode

var name_trie *patricia.Trie

func init() {

	by_name = make(map[string]CountryCode)
	by_alpha3 = make(map[string]CountryCode)
	by_numeric = make(map[int]CountryCode)
	name_trie = patricia.NewTrie()

	by_alpha2 = map[string]CountryCode{
		/**
		 * <a href="http://en.wikipedia.org/wiki/Ascension_Island">Ascension Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AC">AC</a>, ASC, -1,
		 * Exceptionally reserved]
		 */
		"AC": CountryCode{
			Name:        "Ascension Island",
			Alpha2:      "AC",
			Alpha3:      "ASC",
			Numeric:     -1,
			DialingCode: "+247",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Andorra">Andorra</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AD">AD</a>, AND, 16,
		 * Officially assigned]
		 */
		"AD": CountryCode{
			Name:        "Andorra",
			Alpha2:      "AD",
			Alpha3:      "AND",
			Numeric:     20,
			DialingCode: "+376",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_Arab_Emirates">United Arab Emirates</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AE">AE</a>, AE, 784,
		 * Officially assigned]
		 */
		"AE": CountryCode{
			Name:        "United Arab Emirates",
			Alpha2:      "AE",
			Alpha3:      "ARE",
			Numeric:     784,
			DialingCode: "+971",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Afghanistan">Afghanistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AF">AF</a>, AFG, 4,
		 * Officially assigned]
		 */
		"AF": CountryCode{
			Name:        "Afghanistan",
			Alpha2:      "AF",
			Alpha3:      "AFG",
			Numeric:     4,
			DialingCode: "+93",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Antigua_and_Barbuda">Antigua and Barbuda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AG">AG</a>, ATG, 28,
		 * Officially assigned]
		 */
		"AG": CountryCode{
			Name:        "Antigua and Barbuda",
			Alpha2:      "AG",
			Alpha3:      "ATG",
			Numeric:     28,
			DialingCode: "+1-268",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Anguilla">Anguilla</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AI">AI</a>, AIA, 660,
		 * Officially assigned]
		 */
		"AI": CountryCode{
			Name:        "Anguilla",
			Alpha2:      "AI",
			Alpha3:      "AIA",
			Numeric:     660,
			DialingCode: "+1-264",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Albania">Albania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AL">AL</a>, ALB, 8,
		 * Officially assigned]
		 */
		"AL": CountryCode{
			Name:        "Albania",
			Alpha2:      "AL",
			Alpha3:      "ALB",
			Numeric:     8,
			DialingCode: "+355",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Armenia">Armenia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AM">AM</a>, ARM, 51,
		 * Officially assigned]
		 */
		"AM": CountryCode{
			Name:        "Armenia",
			Alpha2:      "AM",
			Alpha3:      "ARM",
			Numeric:     51,
			DialingCode: "+374",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Netherlands_Antilles">Netherlands Antilles</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AN">AN</a>, ANHH, 530,
		 * Traditionally reserved]
		 */
		"AN": CountryCode{
			Name:        "Netherlands Antilles",
			Alpha2:      "AN",
			Alpha3:      "ANHH",
			Numeric:     530,
			DialingCode: "+599",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Angola">Angola</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AO">AO</a>, AGO, 24,
		 * Officially assigned]
		 */
		"AO": CountryCode{
			Name:        "Angola",
			Alpha2:      "AO",
			Alpha3:      "AGO",
			Numeric:     24,
			DialingCode: "+244",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Antarctica">Antarctica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AQ">AQ</a>, ATA, 10,
		 * Officially assigned]
		 */
		"AQ": CountryCode{
			Name:        "Antarctica",
			Alpha2:      "AQ",
			Alpha3:      "ATA",
			Numeric:     10,
			DialingCode: "+672",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Argentina">Argentina</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AR">AR</a>, ARG, 32,
		 * Officially assigned]
		 */
		"AR": CountryCode{
			Name:        "Argentina",
			Alpha2:      "AR",
			Alpha3:      "ARG",
			Numeric:     32,
			DialingCode: "+54",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/American_Samoa">American Samoa</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AS">AS</a>, ASM, 16,
		 * Officially assigned]
		 */
		"AS": CountryCode{
			Name:        "American Samoa",
			Alpha2:      "AS",
			Alpha3:      "ASM",
			Numeric:     16,
			DialingCode: "+1-684",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Austria">Austria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AT">AT</a>, AUT, 40,
		 * Officially assigned]
		 */
		"AT": CountryCode{
			Name:        "Austria",
			Alpha2:      "AT",
			Alpha3:      "AUT",
			Numeric:     40,
			DialingCode: "+43",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Australia">Australia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AU">AU</a>, AUS, 36,
		 * Officially assigned]
		 */
		"AU": CountryCode{
			Name:        "Australia",
			Alpha2:      "AU",
			Alpha3:      "AUS",
			Numeric:     36,
			DialingCode: "+61",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Aruba">Aruba</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AW">AW</a>, ABW, 533,
		 * Officially assigned]
		 */
		"AW": CountryCode{
			Name:        "Aruba",
			Alpha2:      "AW",
			Alpha3:      "ABW",
			Numeric:     533,
			DialingCode: "+297",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/%C3%85land_Islands">&Aring;land Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AX">AX</a>, ALA, 248,
		 * Officially assigned]
		 */
		"AX": CountryCode{
			Name:        "\u212Bland Islands",
			Alpha2:      "AX",
			Alpha3:      "ALA",
			Numeric:     248,
			DialingCode: "",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Azerbaijan">Azerbaijan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AZ">AZ</a>, AZE, 31,
		 * Officially assigned]
		 */
		"AZ": CountryCode{
			Name:        "Azerbaijan",
			Alpha2:      "AZ",
			Alpha3:      "AZE",
			Numeric:     31,
			DialingCode: "+994",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bosnia_and_Herzegovina">Bosnia and Herzegovina</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BA">BA</a>, BIH, 70,
		 * Officially assigned]
		 */
		"BA": CountryCode{
			Name:        "Bosnia and Herzegovina",
			Alpha2:      "BA",
			Alpha3:      "BIH",
			Numeric:     70,
			DialingCode: "+387",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Barbados">Barbados</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BB">BB</a>, BRB, 52,
		 * Officially assigned]
		 */
		"BB": CountryCode{
			Name:        "Barbados",
			Alpha2:      "BB",
			Alpha3:      "BRB",
			Numeric:     52,
			DialingCode: "+1-246",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bangladesh">Bangladesh</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BD">BD</a>, BGD, 50,
		 * Officially assigned]
		 */
		"BD": CountryCode{
			Name:        "Bangladesh",
			Alpha2:      "BD",
			Alpha3:      "BGD",
			Numeric:     50,
			DialingCode: "+880",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Belgium">Belgium</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BE">BE</a>, BEL, 56,
		 * Officially assigned]
		 */
		"BE": CountryCode{
			Name:        "Belgium",
			Alpha2:      "BE",
			Alpha3:      "BEL",
			Numeric:     56,
			DialingCode: "+32",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Burkina_Faso">Burkina Faso</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BF">BF</a>, BFA, 854,
		 * Officially assigned]
		 */
		"BF": CountryCode{
			Name:        "Burkina Faso",
			Alpha2:      "BF",
			Alpha3:      "BFA",
			Numeric:     854,
			DialingCode: "+226",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bulgaria">Bulgaria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BG">BG</a>, BGR, 100,
		 * Officially assigned]
		 */
		"BG": CountryCode{
			Name:        "Bulgaria",
			Alpha2:      "BG",
			Alpha3:      "BGR",
			Numeric:     100,
			DialingCode: "+359",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bahrain">Bahrain</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BH">BH</a>, BHR, 48,
		 * Officially assigned]
		 */
		"BH": CountryCode{
			Name:        "Bahrain",
			Alpha2:      "BH",
			Alpha3:      "BHR",
			Numeric:     48,
			DialingCode: "+973",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Burundi">Burundi</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BI">BI</a>, BDI, 108,
		 * Officially assigned]
		 */
		"BI": CountryCode{
			Name:        "Burundi",
			Alpha2:      "BI",
			Alpha3:      "BDI",
			Numeric:     108,
			DialingCode: "+257",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Benin">Benin</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BJ">BJ</a>, BEN, 204,
		 * Officially assigned]
		 */
		"BJ": CountryCode{
			Name:        "Benin",
			Alpha2:      "BJ",
			Alpha3:      "BEN",
			Numeric:     204,
			DialingCode: "+229",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Barth%C3%A9lemy">Saint Barth&eacute;lemy</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BL">BL</a>, BLM, 652,
		 * Officially assigned]
		 */
		"BL": CountryCode{
			Name:        "Saint Barth\u00E9lemy",
			Alpha2:      "BL",
			Alpha3:      "BLM",
			Numeric:     652,
			DialingCode: "+590",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bermuda">Bermuda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BM">BM</a>, BMU, 60,
		 * Officially assigned]
		 */
		"BM": CountryCode{
			Name:        "Bermuda",
			Alpha2:      "BM",
			Alpha3:      "BMU",
			Numeric:     60,
			DialingCode: "+1-441",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Brunei">Brunei Darussalam</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BN">BN</a>, BRN, 96,
		 * Officially assigned]
		 */
		"BN": CountryCode{
			Name:        "Brunei Darussalam",
			Alpha2:      "BN",
			Alpha3:      "BRN",
			Numeric:     96,
			DialingCode: "+673",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bolivia">Bolivia, Plurinational State of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BO">BO</a>, BOL, 68,
		 * Officially assigned]
		 */
		"BO": CountryCode{
			Name:        "Bolivia, Plurinational State of",
			Alpha2:      "BO",
			Alpha3:      "BOL",
			Numeric:     68,
			DialingCode: "+591",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Caribbean_Netherlands">Bonaire, Sint Eustatius and Saba</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BQ">BQ</a>, BES, 535,
		 * Officially assigned]
		 */
		"BQ": CountryCode{
			Name:        "Bonaire, Sint Eustatius and Saba",
			Alpha2:      "BQ",
			Alpha3:      "BES",
			Numeric:     535,
			DialingCode: "+599",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Brazil">Brazil</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BR">BR</a>, BRA, 76,
		 * Officially assigned]
		 */
		"BR": CountryCode{
			Name:        "Brazil",
			Alpha2:      "BR",
			Alpha3:      "BRA",
			Numeric:     76,
			DialingCode: "+55",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/The_Bahamas">Bahamas</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BS">BS</a>, BHS, 44,
		 * Officially assigned]
		 */
		"BS": CountryCode{
			Name:        "Bahamas",
			Alpha2:      "BS",
			Alpha3:      "BHS",
			Numeric:     44,
			DialingCode: "+1-242",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bhutan">Bhutan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BT">BT</a>, BTN, 64,
		 * Officially assigned]
		 */
		"BT": CountryCode{
			Name:        "Bhutan",
			Alpha2:      "BT",
			Alpha3:      "BTN",
			Numeric:     64,
			DialingCode: "+975",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Burma">Burma</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BU">BU</a>, BUMM, 104,
		 * Officially assigned]
		 *
		 * @see #MM
		 */
		"BU": CountryCode{
			Name:        "Burma",
			Alpha2:      "BU",
			Alpha3:      "BUMM",
			Numeric:     104,
			DialingCode: "+95",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bouvet_Island">Bouvet Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BV">BV</a>, BVT, 74,
		 * Officially assigned]
		 */
		"BV": CountryCode{
			Name:        "Bouvet Island",
			Alpha2:      "BV",
			Alpha3:      "BVT",
			Numeric:     74,
			DialingCode: "",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Botswana">Botswana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BW">BW</a>, BWA, 72,
		 * Officially assigned]
		 */
		"BW": CountryCode{
			Name:        "Botswana",
			Alpha2:      "BW",
			Alpha3:      "BWA",
			Numeric:     72,
			DialingCode: "+267",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Belarus">Belarus</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BY">BY</a>, BLR, 112,
		 * Officially assigned]
		 */
		"BY": CountryCode{
			Name:        "Belarus",
			Alpha2:      "BY",
			Alpha3:      "BLR",
			Numeric:     112,
			DialingCode: "+375",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Belize">Belize</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BZ">BZ</a>, BLZ, 84,
		 * Officially assigned]
		 */
		"BZ": CountryCode{
			Name:        "Belize",
			Alpha2:      "BZ",
			Alpha3:      "BLZ",
			Numeric:     84,
			DialingCode: "+501",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Canada">Canada</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CA">CA</a>, CAN, 124,
		 * Officially assigned]
		 */
		"CA": CountryCode{
			Name:        "Canada",
			Alpha2:      "CA",
			Alpha3:      "CAN",
			Numeric:     124,
			DialingCode: "+1",
			Assignment:  OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Cocos_(Keeling)_Islands">Cocos (Keeling) Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CC">CC</a>, CCK, 166,
		 * Officially assigned]
		 */
		"CC": CountryCode{
			Name:        "Cocos (Keeling) Islands",
			Alpha2:      "CC",
			Alpha3:      "CCK",
			Numeric:     166,
			DialingCode: "+61",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Democratic_Republic_of_the_Congo">Congo, the Democratic Republic of the</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CD">CD</a>, COD, 180,
		 * Officially assigned]
		 */
		"CD": CountryCode{
			Name:        "Congo, the Democratic Republic of the",
			Alpha2:      "CD",
			Alpha3:      "COD",
			Numeric:     180,
			DialingCode: "+243",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Central_African_Republic">Central African Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CF">CF</a>, CAF, 140,
		 * Officially assigned]
		 */
		"CF": CountryCode{
			Name:        "Central African Republic",
			Alpha2:      "CF",
			Alpha3:      "CAF",
			Numeric:     140,
			DialingCode: "+236",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Republic_of_the_Congo">Congo</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CG">CG</a>, COG, 178,
		 * Officially assigned]
		 */
		"CG": CountryCode{
			Name:        "Congo",
			Alpha2:      "CG",
			Alpha3:      "COG",
			Numeric:     178,
			DialingCode: "+242",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Switzerland">Switzerland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CH">CH</a>, CHE, 756,
		 * Officially assigned]
		 */
		"CH": CountryCode{
			Name:        "Switzerland",
			Alpha2:      "CH",
			Alpha3:      "CHE",
			Numeric:     756,
			DialingCode: "+41",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/C%C3%B4te_d%27Ivoire">C&ocirc;te d'Ivoire</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CI">CI</a>, CIV, 384,
		 * Officially assigned]
		 */
		"CI": CountryCode{
			Name:        "C\u00F4te d'Ivoire",
			Alpha2:      "CI",
			Alpha3:      "CIV",
			Numeric:     384,
			DialingCode: "+225",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cook_Islands">Cook Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CK">CK</a>, COK, 184,
		 * Officially assigned]
		 */
		"CK": CountryCode{
			Name:        "Cook Islands",
			Alpha2:      "CK",
			Alpha3:      "COK",
			Numeric:     184,
			DialingCode: "+682",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Chile">Chile</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CL">CL</a>, CHL, 152,
		 * Officially assigned]
		 */
		"CL": CountryCode{
			Name:        "Chile",
			Alpha2:      "CL",
			Alpha3:      "CHL",
			Numeric:     152,
			DialingCode: "+56",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cameroon">Cameroon</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CM">CM</a>, CMR, 120,
		 * Officially assigned]
		 */
		"CM": CountryCode{
			Name:        "Cameroon",
			Alpha2:      "CM",
			Alpha3:      "CMR",
			Numeric:     120,
			DialingCode: "+237",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/China">China</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CN">CN</a>, CHN, 156,
		 * Officially assigned]
		 */
		"CN": CountryCode{
			Name:        "China",
			Alpha2:      "CN",
			Alpha3:      "CHN",
			Numeric:     156,
			DialingCode: "+86",
			Assignment:  OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Colombia">Colombia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CO">CO</a>, COL, 170,
		 * Officially assigned]
		 */
		"CO": CountryCode{
			Name:        "Colombia",
			Alpha2:      "CO",
			Alpha3:      "COL",
			Numeric:     170,
			DialingCode: "+57",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Clipperton_Island">Clipperton Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CP">CP</a>, CPT, -1,
		 * Exceptionally reserved]
		 */
		"CP": CountryCode{
			Name:        "Clipperton Island",
			Alpha2:      "CP",
			Alpha3:      "CPT",
			Numeric:     -1,
			DialingCode: "",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Costa_Rica">Costa Rica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CR">CR</a>, CRI, 188,
		 * Officially assigned]
		 */
		"CR": CountryCode{
			Name:        "Costa Rica",
			Alpha2:      "CR",
			Alpha3:      "CRI",
			Numeric:     188,
			DialingCode: "+506",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Serbia_and_Montenegro">Serbia and Montenegro</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CS">CS</a>, CSXX, 891,
		 * Traditionally reserved]
		 */
		"CS": CountryCode{
			Name:        "Serbia and Montenegro",
			Alpha2:      "CS",
			Alpha3:      "CSXX",
			Numeric:     891,
			DialingCode: "+381",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cuba">Cuba</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CU">CU</a>, CUB, 192,
		 * Officially assigned]
		 */
		"CU": CountryCode{
			Name:        "Cuba",
			Alpha2:      "CU",
			Alpha3:      "CUB",
			Numeric:     192,
			DialingCode: "+53",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cape_Verde">Cape Verde</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CV">CV</a>, CPV, 132,
		 * Officially assigned]
		 */
		"CV": CountryCode{
			Name:        "Cape Verde",
			Alpha2:      "CV",
			Alpha3:      "CPV",
			Numeric:     132,
			DialingCode: "+238",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cura%C3%A7ao">Cura&ccedil;ao</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CW">CW</a>, CUW, 531,
		 * Officially assigned]
		 */
		"CW": CountryCode{
			Name:        "Cura\u00E7ao",
			Alpha2:      "CW",
			Alpha3:      "CUW",
			Numeric:     531,
			DialingCode: "+599",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Christmas_Island">Christmas Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CX">CX</a>, CXR, 162,
		 * Officially assigned]
		 */
		"CX": CountryCode{
			Name:        "Christmas Island",
			Alpha2:      "CX",
			Alpha3:      "CXR",
			Numeric:     162,
			DialingCode: "+61",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cyprus">Cyprus</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CY">CY</a>, CYP, 196,
		 * Officially assigned]
		 */
		"CY": CountryCode{
			Name:        "Cyprus",
			Alpha2:      "CY",
			Alpha3:      "CYP",
			Numeric:     196,
			DialingCode: "+357",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Czech_Republic">Czech Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CZ">CZ</a>, CZE, 203,
		 * Officially assigned]
		 */
		"CZ": CountryCode{
			Name:        "Czech Republic",
			Alpha2:      "CZ",
			Alpha3:      "CZE",
			Numeric:     203,
			DialingCode: "+420",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Germany">Germany</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DE">DE</a>, DEU, 276,
		 * Officially assigned]
		 */
		"DE": CountryCode{
			Name:        "Germany",
			Alpha2:      "DE",
			Alpha3:      "DEU",
			Numeric:     276,
			DialingCode: "+49",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Diego_Garcia">Diego Garcia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DG">DG</a>, DGA, -1,
		 * Exceptionally reserved]
		 */
		"DG": CountryCode{
			Name:        "Diego Garcia",
			Alpha2:      "DG",
			Alpha3:      "DGA",
			Numeric:     -1,
			DialingCode: "+246",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Djibouti">Djibouti</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DJ">DJ</a>, DJI, 262,
		 * Officially assigned]
		 */
		"DJ": CountryCode{
			Name:        "Djibouti",
			Alpha2:      "DJ",
			Alpha3:      "DJI",
			Numeric:     262,
			DialingCode: "+253",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Denmark">Denmark</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DK">DK</a>, DNK, 208,
		 * Officially assigned]
		 */
		"DK": CountryCode{
			Name:        "Denmark",
			Alpha2:      "DK",
			Alpha3:      "DNK",
			Numeric:     208,
			DialingCode: "+45",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Dominica">Dominica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DM">DM</a>, DMA, 212,
		 * Officially assigned]
		 */
		"DM": CountryCode{
			Name:        "Dominica",
			Alpha2:      "DM",
			Alpha3:      "DMA",
			Numeric:     212,
			DialingCode: "+1-767",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Dominican_Republic">Dominican Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DO">DO</a>, DOM, 214,
		 * Officially assigned]
		 */
		"DO": CountryCode{
			Name:        "Dominican Republic",
			Alpha2:      "DO",
			Alpha3:      "DOM",
			Numeric:     214,
			DialingCode: "+1-809, +1-829, +1-849",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Algeria">Algeria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DZ">DZ</a>, DZA, 12,
		 * Officially assigned]
		 */
		"DZ": CountryCode{
			Name:        "Algeria",
			Alpha2:      "DZ",
			Alpha3:      "DZA",
			Numeric:     12,
			DialingCode: "+213",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ceuta">Ceuta</a>,
		 * <a href="http://en.wikipedia.org/wiki/Melilla">Melilla</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EA">EA</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"EA": CountryCode{
			Name:        "Ceuta, Melilla",
			Alpha2:      "EA",
			Alpha3:      "",
			Numeric:     -1,
			DialingCode: "",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ecuador">Ecuador</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EC">EC</a>, ECU, 218,
		 * Officially assigned]
		 */
		"EC": CountryCode{
			Name:        "Ecuador",
			Alpha2:      "EC",
			Alpha3:      "ECU",
			Numeric:     218,
			DialingCode: "+593",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Estonia">Estonia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EE">EE</a>, EST, 233,
		 * Officially assigned]
		 */
		"EE": CountryCode{
			Name:        "Estonia",
			Alpha2:      "EE",
			Alpha3:      "EST",
			Numeric:     233,
			DialingCode: "+372",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Egypt">Egypt</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EG">EG</a>, EGY, 818,
		 * Officially assigned]
		 */
		"EG": CountryCode{
			Name:        "Egypt",
			Alpha2:      "EG",
			Alpha3:      "EGY",
			Numeric:     818,
			DialingCode: "+20",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Western_Sahara">Western Sahara</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EH">EH</a>, ESH, 732,
		 * Officially assigned]
		 */
		"EH": CountryCode{
			Name:        "Western Sahara",
			Alpha2:      "EH",
			Alpha3:      "ESH",
			Numeric:     732,
			DialingCode: "+212",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Eritrea">Eritrea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ER">ER</a>, ERI, 232,
		 * Officially assigned]
		 */
		"ER": CountryCode{
			Name:        "Eritrea",
			Alpha2:      "ER",
			Alpha3:      "ERI",
			Numeric:     232,
			DialingCode: "+291",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Spain">Spain</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ES">ES</a>, ESP, 724,
		 * Officially assigned]
		 */
		"ES": CountryCode{
			Name:        "Spain",
			Alpha2:      "ES",
			Alpha3:      "ESP",
			Numeric:     724,
			DialingCode: "+34",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ethiopia">Ethiopia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ET">ET</a>, ETH, 231,
		 * Officially assigned]
		 */
		"ET": CountryCode{
			Name:        "Ethiopia",
			Alpha2:      "ET",
			Alpha3:      "ETH",
			Numeric:     231,
			DialingCode: "+251",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/European_Union">European Union</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EU">EU</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"EU": CountryCode{
			Name:        "European Union",
			Alpha2:      "EU",
			Alpha3:      "",
			Numeric:     -1,
			DialingCode: "",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Finland">Finland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FI">FI</a>, FIN, 246,
		 * Officially assigned]
		 *
		 * @see #SF
		 */
		"FI": CountryCode{
			Name:        "Finland",
			Alpha2:      "FI",
			Alpha3:      "FIN",
			Numeric:     246,
			DialingCode: "+358",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Fiji">Fiji</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#">FJ</a>, FJI, 242,
		 * Officially assigned]
		 */
		"FJ": CountryCode{
			Name:        "Fiji",
			Alpha2:      "FJ",
			Alpha3:      "FJI",
			Numeric:     242,
			DialingCode: "+679",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Falkland_Islands">Falkland Islands (Malvinas)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FK">FK</a>, FLK, 238,
		 * Officially assigned]
		 */
		"FK": CountryCode{
			Name:        "Falkland Islands (Malvinas)",
			Alpha2:      "FK",
			Alpha3:      "FLK",
			Numeric:     238,
			DialingCode: "+500",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Federated_States_of_Micronesia">Micronesia, Federated States of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FM">FM</a>, FSM, 583,
		 * Officially assigned]
		 */
		"FM": CountryCode{
			Name:        "Micronesia, Federated States of",
			Alpha2:      "FM",
			Alpha3:      "FSM",
			Numeric:     583,
			DialingCode: "+691",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Faroe_Islands">Faroe Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FO">FO</a>, FRO, 234,
		 * Officially assigned]
		 */
		"FO": CountryCode{
			Name:        "Faroe Islands",
			Alpha2:      "FO",
			Alpha3:      "FRO",
			Numeric:     234,
			DialingCode: "+298",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/France">France</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FR">FR</a>, FRA, 250,
		 * Officially assigned]
		 */
		"FR": CountryCode{
			Name:        "France",
			Alpha2:      "FR",
			Alpha3:      "FRA",
			Numeric:     250,
			DialingCode: "+33",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Metropolitan_France">France, Metropolitan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FX">FX</a>, FXX, -1,
		 * Exceptionally reserved]
		 */
		"FX": CountryCode{
			Name:        "France, Metropolitan",
			Alpha2:      "FX",
			Alpha3:      "FXX",
			Numeric:     -1,
			DialingCode: "",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Gabon">Gabon </a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GA">GA</a>, GAB, 266,
		 * Officially assigned]
		 */
		"GA": CountryCode{
			Name:        "Gabon",
			Alpha2:      "GA",
			Alpha3:      "GAB",
			Numeric:     266,
			DialingCode: "+241",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_Kingdom">United Kingdom</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GB">GB</a>, GBR, 826,
		 * Officially assigned]
		 */
		"GB": CountryCode{
			Name:        "United Kingdom",
			Alpha2:      "GB",
			Alpha3:      "GBR",
			Numeric:     826,
			DialingCode: "+44",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Grenada">Grenada</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GD">GD</a>, GRD, 308,
		 * Officially assigned]
		 */
		"GD": CountryCode{
			Name:        "Grenada",
			Alpha2:      "GD",
			Alpha3:      "GRD",
			Numeric:     308,
			DialingCode: "+1-473",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Georgia_(country)">Georgia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GE">GE</a>, GEO, 268,
		 * Officially assigned]
		 */
		"GE": CountryCode{
			Name:        "Georgia",
			Alpha2:      "GE",
			Alpha3:      "GEO",
			Numeric:     268,
			DialingCode: "+995",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/French_Guiana">French Guiana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GF">GF</a>, GUF, 254,
		 * Officially assigned]
		 */
		"GF": CountryCode{
			Name:        "French Guiana",
			Alpha2:      "GF",
			Alpha3:      "GUF",
			Numeric:     254,
			DialingCode: "+594",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guernsey">Guernsey</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GG">GG</a>, GGY, 831,
		 * Officially assigned]
		 */
		"GG": CountryCode{
			Name:        "Guernsey",
			Alpha2:      "GG",
			Alpha3:      "GGY",
			Numeric:     831,
			DialingCode: "+44-1481",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ghana">Ghana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GH">GH</a>, GHA, 288,
		 * Officially assigned]
		 */
		"GH": CountryCode{
			Name:        "Ghana",
			Alpha2:      "GH",
			Alpha3:      "GHA",
			Numeric:     288,
			DialingCode: "+233",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Gibraltar">Gibraltar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GI">GI</a>, GIB, 292,
		 * Officially assigned]
		 */
		"GI": CountryCode{
			Name:        "Gibraltar",
			Alpha2:      "GI",
			Alpha3:      "GIB",
			Numeric:     292,
			DialingCode: "+350",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Greenland">Greenland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GL">GL</a>, GRL, 304,
		 * Officially assigned]
		 */
		"GL": CountryCode{
			Name:        "Greenland",
			Alpha2:      "GL",
			Alpha3:      "GRL",
			Numeric:     304,
			DialingCode: "+299",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/The_Gambia">Gambia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GM">GM</a>, GMB, 270,
		 * Officially assigned]
		 */
		"GM": CountryCode{
			Name:        "Gambia",
			Alpha2:      "GM",
			Alpha3:      "GMB",
			Numeric:     270,
			DialingCode: "+220",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guinea">Guinea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GN">GN</a>, GIN, 324,
		 * Officially assigned]
		 */
		"GN": CountryCode{
			Name:        "Guinea",
			Alpha2:      "GN",
			Alpha3:      "GIN",
			Numeric:     324,
			DialingCode: "+224",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guadeloupe">Guadeloupe</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GP">GP</a>, GLP, 312,
		 * Officially assigned]
		 */
		"GP": CountryCode{
			Name:        "Guadeloupe",
			Alpha2:      "GP",
			Alpha3:      "GLP",
			Numeric:     312,
			DialingCode: "+590",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Equatorial_Guinea">Equatorial Guinea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GQ">GQ</a>, GNQ, 226,
		 * Officially assigned]
		 */
		"GQ": CountryCode{
			Name:        "Equatorial Guinea",
			Alpha2:      "GQ",
			Alpha3:      "GNQ",
			Numeric:     226,
			DialingCode: "+240",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Greece">Greece</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GR">GR</a>, GRC, 300,
		 * Officially assigned]
		 */
		"GR": CountryCode{
			Name:        "Greece",
			Alpha2:      "GR",
			Alpha3:      "GRC",
			Numeric:     300,
			DialingCode: "+30",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Georgia_and_the_South_Sandwich_Islands">South Georgia and the South Sandwich Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GS">GS</a>, SGS, 239,
		 * Officially assigned]
		 */
		"GS": CountryCode{
			Name:        "South Georgia and the South Sandwich Islands",
			Alpha2:      "GS",
			Alpha3:      "SGS",
			Numeric:     239,
			DialingCode: "+500",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guatemala">Guatemala</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GT">GT</a>, GTM, 320,
		 * Officially assigned]
		 */
		"GT": CountryCode{
			Name:        "Guatemala",
			Alpha2:      "GT",
			Alpha3:      "GTM",
			Numeric:     320,
			DialingCode: "+502",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guam">Guam</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GU">GU</a>, GUM, 316,
		 * Officially assigned]
		 */
		"GU": CountryCode{
			Name:        "Guam",
			Alpha2:      "GU",
			Alpha3:      "GUM",
			Numeric:     316,
			DialingCode: "+1-671",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guinea-Bissau">Guinea-Bissau</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GW">GW</a>, GNB, 624,
		 * Officially assigned]
		 */
		"GW": CountryCode{
			Name:        "Guinea-Bissau",
			Alpha2:      "GW",
			Alpha3:      "GNB",
			Numeric:     624,
			DialingCode: "+245",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guyana">Guyana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GY">GY</a>, GUY, 328,
		 * Officially assigned]
		 */
		"GY": CountryCode{
			Name:        "Guyana",
			Alpha2:      "GY",
			Alpha3:      "GUY",
			Numeric:     328,
			DialingCode: "+592",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Hong_Kong">Hong Kong</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HK">HK</a>, HKG, 344,
		 * Officially assigned]
		 */
		"HK": CountryCode{
			Name:        "Hong Kong",
			Alpha2:      "HK",
			Alpha3:      "HKG",
			Numeric:     344,
			DialingCode: "+852",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Heard_Island_and_McDonald_Islands">Heard Island and McDonald Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HM">HM</a>, HMD, 334,
		 * Officially assigned]
		 */
		"HM": CountryCode{
			Name:        "Heard Island and McDonald Islands",
			Alpha2:      "HM",
			Alpha3:      "HMD",
			Numeric:     334,
			DialingCode: "",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Honduras">Honduras</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HN">HN</a>, HND, 340,
		 * Officially assigned]
		 */
		"HN": CountryCode{
			Name:        "Honduras",
			Alpha2:      "HN",
			Alpha3:      "HND",
			Numeric:     340,
			DialingCode: "+504",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Croatia">Croatia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HR">HR</a>, HRV, 191,
		 * Officially assigned]
		 */
		"HR": CountryCode{
			Name:        "Croatia",
			Alpha2:      "HR",
			Alpha3:      "HRV",
			Numeric:     191,
			DialingCode: "+385",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Haiti">Haiti</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HT">HT</a>, HTI, 332,
		 * Officially assigned]
		 */
		"HT": CountryCode{
			Name:        "Haiti",
			Alpha2:      "HT",
			Alpha3:      "HTI",
			Numeric:     332,
			DialingCode: "+509",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Hungary">Hungary</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HU">HU</a>, HUN, 348,
		 * Officially assigned]
		 */
		"HU": CountryCode{
			Name:        "Hungary",
			Alpha2:      "HU",
			Alpha3:      "HUN",
			Numeric:     348,
			DialingCode: "+36",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Canary_Islands">Canary Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IC">IC</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"IC": CountryCode{
			Name:        "Canary Islands",
			Alpha2:      "IC",
			Alpha3:      "",
			Numeric:     -1,
			DialingCode: "",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Indonesia">Indonesia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ID">ID</a>, IDN, 360,
		 * Officially assigned]
		 */
		"ID": CountryCode{
			Name:        "Indonesia",
			Alpha2:      "ID",
			Alpha3:      "IDN",
			Numeric:     360,
			DialingCode: "+62",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Republic_of_Ireland">Ireland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IE">IE</a>, IRL, 372,
		 * Officially assigned]
		 */
		"IE": CountryCode{
			Name:        "Ireland",
			Alpha2:      "IE",
			Alpha3:      "IRL",
			Numeric:     372,
			DialingCode: "+353",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Israel">Israel</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IL">IL</a>, ISR, 376,
		 * Officially assigned]
		 */
		"IL": CountryCode{
			Name:        "Israel",
			Alpha2:      "IL",
			Alpha3:      "ISR",
			Numeric:     376,
			DialingCode: "+972",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Isle_of_Man">Isle of Man</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IM">IM</a>, IMN, 833,
		 * Officially assigned]
		 */
		"IM": CountryCode{
			Name:        "Isle of Man",
			Alpha2:      "IM",
			Alpha3:      "IMN",
			Numeric:     833,
			DialingCode: "+44-1624",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/India">India</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IN">IN</a>, IND, 356,
		 * Officially assigned]
		 */
		"IN": CountryCode{
			Name:        "India",
			Alpha2:      "IN",
			Alpha3:      "IND",
			Numeric:     356,
			DialingCode: "+91",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/British_Indian_Ocean_Territory">British Indian Ocean Territory</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IO">IO</a>, IOT, 86,
		 * Officially assigned]
		 */
		"IO": CountryCode{
			Name:        "British Indian Ocean Territory",
			Alpha2:      "IO",
			Alpha3:      "IOT",
			Numeric:     86,
			DialingCode: "+246",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Iraq">Iraq</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IQ">IQ</a>, IRQ, 368,
		 * Officially assigned]
		 */
		"IQ": CountryCode{
			Name:        "Iraq",
			Alpha2:      "IQ",
			Alpha3:      "IRQ",
			Numeric:     368,
			DialingCode: "+964",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Iran">Iran, Islamic Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IR">IR</a>, IRN, 364,
		 * Officially assigned]
		 */
		"IR": CountryCode{
			Name:        "Iran, Islamic Republic of",
			Alpha2:      "IR",
			Alpha3:      "IRN",
			Numeric:     364,
			DialingCode: "+98",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Iceland">Iceland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IS">IS</a>, ISL, 352,
		 * Officially assigned]
		 */
		"IS": CountryCode{
			Name:        "Iceland",
			Alpha2:      "IS",
			Alpha3:      "ISL",
			Numeric:     352,
			DialingCode: "+354",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Italy">Italy</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IT">IT</a>, ITA, 380,
		 * Officially assigned]
		 */
		"IT": CountryCode{
			Name:        "Italy",
			Alpha2:      "IT",
			Alpha3:      "ITA",
			Numeric:     380,
			DialingCode: "+39",
			Assignment:  OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Jersey">Jersey</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JE">JE</a>, JEY, 832,
		 * Officially assigned]
		 */
		"JE": CountryCode{
			Name:        "Jersey",
			Alpha2:      "JE",
			Alpha3:      "JEY",
			Numeric:     832,
			DialingCode: "+44-1534",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Jamaica">Jamaica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JM">JM</a>, JAM, 388,
		 * Officially assigned]
		 */
		"JM": CountryCode{
			Name:        "Jamaica",
			Alpha2:      "JM",
			Alpha3:      "JAM",
			Numeric:     388,
			DialingCode: "+1-876",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Jordan">Jordan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JO">JO</a>, JOR, 400,
		 * Officially assigned]
		 */
		"JO": CountryCode{
			Name:        "Jordan",
			Alpha2:      "JO",
			Alpha3:      "JOR",
			Numeric:     400,
			DialingCode: "+962",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Japan">Japan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JP">JP</a>, JPN, 392,
		 * Officially assigned]
		 */
		"JP": CountryCode{
			Name:        "Japan",
			Alpha2:      "JP",
			Alpha3:      "JPN",
			Numeric:     392,
			DialingCode: "+81",
			Assignment:  OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Kenya">Kenya</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KE">KE</a>, KEN, 404,
		 * Officially assigned]
		 */
		"KE": CountryCode{
			Name:        "Kenya",
			Alpha2:      "KE",
			Alpha3:      "KEN",
			Numeric:     404,
			DialingCode: "+254",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kyrgyzstan">Kyrgyzstan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KG">KG</a>, KGZ, 417,
		 * Officially assigned]
		 */
		"KG": CountryCode{
			Name:        "Kyrgyzstan",
			Alpha2:      "KG",
			Alpha3:      "KGZ",
			Numeric:     417,
			DialingCode: "+996",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cambodia">Cambodia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KH">KH</a>, KHM, 116,
		 * Officially assigned]
		 */
		"KH": CountryCode{
			Name:        "Cambodia",
			Alpha2:      "KH",
			Alpha3:      "KHM",
			Numeric:     116,
			DialingCode: "+855",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kiribati">Kiribati</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KI">KI</a>, KIR, 296,
		 * Officially assigned]
		 */
		"KI": CountryCode{
			Name:        "Kiribati",
			Alpha2:      "KI",
			Alpha3:      "KIR",
			Numeric:     296,
			DialingCode: "+686",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Comoros">Comoros</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KM">KM</a>, COM, 174,
		 * Officially assigned]
		 */
		"KM": CountryCode{
			Name:        "Comoros",
			Alpha2:      "KM",
			Alpha3:      "COM",
			Numeric:     174,
			DialingCode: "+269",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Kitts_and_Nevis">Saint Kitts and Nevis</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KN">KN</a>, KNA, 659,
		 * Officially assigned]
		 */
		"KN": CountryCode{
			Name:        "Saint Kitts and Nevis",
			Alpha2:      "KN",
			Alpha3:      "KNA",
			Numeric:     659,
			DialingCode: "+1-869",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/North_Korea">Korea, Democratic People's Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KP">KP</a>, PRK, 408,
		 * Officially assigned]
		 */
		"KP": CountryCode{
			Name:        "Korea, Democratic People's Republic of",
			Alpha2:      "KP",
			Alpha3:      "PRK",
			Numeric:     408,
			DialingCode: "+850",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Korea">Korea, Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KR">KR</a>, KOR, 410,
		 * Officially assigned]
		 */
		"KR": CountryCode{
			Name:        "Korea, Republic of",
			Alpha2:      "KR",
			Alpha3:      "KOR",
			Numeric:     410,
			DialingCode: "+82",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kuwait">Kuwait</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KW">KW</a>, KWT, 414,
		 * Officially assigned]
		 */
		"KW": CountryCode{
			Name:        "Kuwait",
			Alpha2:      "KW",
			Alpha3:      "KWT",
			Numeric:     414,
			DialingCode: "+965",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cayman_Islands">Cayman Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KY">KY</a>, CYM, 136,
		 * Officially assigned]
		 */
		"KY": CountryCode{
			Name:        "Cayman Islands",
			Alpha2:      "KY",
			Alpha3:      "CYM",
			Numeric:     136,
			DialingCode: "+1-345",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kazakhstan">Kazakhstan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KZ">KZ</a>, KAZ, 398,
		 * Officially assigned]
		 */
		"KZ": CountryCode{
			Name:        "Kazakhstan",
			Alpha2:      "KZ",
			Alpha3:      "KAZ",
			Numeric:     398,
			DialingCode: "+7",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Laos">Lao People's Democratic Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LA">LA</a>, LAO, 418,
		 * Officially assigned]
		 */
		"LA": CountryCode{
			Name:        "Lao People's Democratic Republic",
			Alpha2:      "LA",
			Alpha3:      "LAO",
			Numeric:     418,
			DialingCode: "+856",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Lebanon">Lebanon</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LB">LB</a>, LBN, 422,
		 * Officially assigned]
		 */
		"LB": CountryCode{
			Name:        "Lebanon",
			Alpha2:      "LB",
			Alpha3:      "LBN",
			Numeric:     422,
			DialingCode: "+961",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Lucia">Saint Lucia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LC">LC</a>, LCA, 662,
		 * Officially assigned]
		 */
		"LC": CountryCode{
			Name:        "Saint Lucia",
			Alpha2:      "LC",
			Alpha3:      "LCA",
			Numeric:     662,
			DialingCode: "+1-758",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Liechtenstein">Liechtenstein</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LI">LI</a>, LIE, 438,
		 * Officially assigned]
		 */
		"LI": CountryCode{
			Name:        "Liechtenstein",
			Alpha2:      "LI",
			Alpha3:      "LIE",
			Numeric:     438,
			DialingCode: "+423",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sri_Lanka">Sri Lanka</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LK">LK</a>, LKA, 144,
		 * Officially assigned]
		 */
		"LK": CountryCode{
			Name:        "Sri Lanka",
			Alpha2:      "LK",
			Alpha3:      "LKA",
			Numeric:     144,
			DialingCode: "+94",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Liberia">Liberia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LR">LR</a>, LBR, 430,
		 * Officially assigned]
		 */
		"LR": CountryCode{
			Name:        "Liberia",
			Alpha2:      "LR",
			Alpha3:      "LBR",
			Numeric:     430,
			DialingCode: "+231",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Lesotho">Lesotho</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LS">LS</a>, LSO, 426,
		 * Officially assigned]
		 */
		"LS": CountryCode{
			Name:        "Lesotho",
			Alpha2:      "LS",
			Alpha3:      "LSO",
			Numeric:     426,
			DialingCode: "+266",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Lithuania">Lithuania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LT">LT</a>, LTU, 440,
		 * Officially assigned]
		 */
		"LT": CountryCode{
			Name:        "Lithuania",
			Alpha2:      "LT",
			Alpha3:      "LTU",
			Numeric:     440,
			DialingCode: "+370",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Luxembourg">Luxembourg</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LU">LU</a>, LUX, 442,
		 * Officially assigned]
		 */
		"LU": CountryCode{
			Name:        "Luxembourg",
			Alpha2:      "LU",
			Alpha3:      "LUX",
			Numeric:     442,
			DialingCode: "+352",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Latvia">Latvia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LV">LV</a>, LVA, 428,
		 * Officially assigned]
		 */
		"LV": CountryCode{
			Name:        "Latvia",
			Alpha2:      "LV",
			Alpha3:      "LVA",
			Numeric:     428,
			DialingCode: "+371",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Libya">Libya</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LY">LY</a>, LBY, 434,
		 * Officially assigned]
		 */
		"LY": CountryCode{
			Name:        "Libya",
			Alpha2:      "LY",
			Alpha3:      "LBY",
			Numeric:     434,
			DialingCode: "+218",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Morocco">Morocco</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MA">MA</a>, MAR, 504,
		 * Officially assigned]
		 */
		"MA": CountryCode{
			Name:        "Morocco",
			Alpha2:      "MA",
			Alpha3:      "MAR",
			Numeric:     504,
			DialingCode: "+212",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Monaco">Monaco</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MC">MC</a>, MCO, 492,
		 * Officially assigned]
		 */
		"MC": CountryCode{
			Name:        "Monaco",
			Alpha2:      "MC",
			Alpha3:      "MCO",
			Numeric:     492,
			DialingCode: "+377",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Moldova">Moldova, Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MD">MD</a>, MDA, 498,
		 * Officially assigned]
		 */
		"MD": CountryCode{
			Name:        "Moldova, Republic of",
			Alpha2:      "MD",
			Alpha3:      "MDA",
			Numeric:     498,
			DialingCode: "+373",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Montenegro">Montenegro</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ME">ME</a>, MNE, 499,
		 * Officially assigned]
		 */
		"ME": CountryCode{
			Name:        "Montenegro",
			Alpha2:      "ME",
			Alpha3:      "MNE",
			Numeric:     499,
			DialingCode: "+382",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Collectivity_of_Saint_Martin">Saint Martin (French part)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MF">MF</a>, MAF, 663,
		 * Officially assigned]
		 */
		"MF": CountryCode{
			Name:        "Saint Martin (French part)",
			Alpha2:      "MF",
			Alpha3:      "MAF",
			Numeric:     663,
			DialingCode: "+590",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Madagascar">Madagascar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MG">MG</a>, MDG, 450,
		 * Officially assigned]
		 */
		"MG": CountryCode{
			Name:        "Madagascar",
			Alpha2:      "MG",
			Alpha3:      "MDG",
			Numeric:     450,
			DialingCode: "+261",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Marshall_Islands">Marshall Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MH">MH</a>, MHL, 584,
		 * Officially assigned]
		 */
		"MH": CountryCode{
			Name:        "Marshall Islands",
			Alpha2:      "MH",
			Alpha3:      "MHL",
			Numeric:     584,
			DialingCode: "+692",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Republic_of_Macedonia">Macedonia, the former Yugoslav Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MK">MK</a>, MKD, 807,
		 * Officially assigned]
		 */
		"MK": CountryCode{
			Name:        "Macedonia, the former Yugoslav Republic of",
			Alpha2:      "MK",
			Alpha3:      "MKD",
			Numeric:     807,
			DialingCode: "+389",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mali">Mali</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ML">ML</a>, MLI, 466,
		 * Officially assigned]
		 */
		"ML": CountryCode{
			Name:        "Mali",
			Alpha2:      "ML",
			Alpha3:      "MLI",
			Numeric:     466,
			DialingCode: "+223",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Myanmar">Myanmar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MM">MM</a>, MMR, 104,
		 * Officially assigned]
		 *
		 * @see #BU
		 */
		"MM": CountryCode{
			Name:        "Myanmar",
			Alpha2:      "MM",
			Alpha3:      "MMR",
			Numeric:     104,
			DialingCode: "+95",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mongolia">Mongolia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MN">MN</a>, MNG, 496,
		 * Officially assigned]
		 */
		"MN": CountryCode{
			Name:        "Mongolia",
			Alpha2:      "MN",
			Alpha3:      "MNG",
			Numeric:     496,
			DialingCode: "+976",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Macau">Macao</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MO">MO</a>, MCO, 492,
		 * Officially assigned]
		 */
		"MO": CountryCode{
			Name:        "Macao",
			Alpha2:      "MO",
			Alpha3:      "MAC",
			Numeric:     446,
			DialingCode: "+853",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Northern_Mariana_Islands">Northern Mariana Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MP">MP</a>, MNP, 580,
		 * Officially assigned]
		 */
		"MP": CountryCode{
			Name:        "Northern Mariana Islands",
			Alpha2:      "MP",
			Alpha3:      "MNP",
			Numeric:     580,
			DialingCode: "+1-670",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Martinique">Martinique</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MQ">MQ</a>, MTQ, 474,
		 * Officially assigned]
		 */
		"MQ": CountryCode{
			Name:        "Martinique",
			Alpha2:      "MQ",
			Alpha3:      "MTQ",
			Numeric:     474,
			DialingCode: "+596",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mauritania">Mauritania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MR">MR</a>, MRT, 478,
		 * Officially assigned]
		 */
		"MR": CountryCode{
			Name:        "Mauritania",
			Alpha2:      "MR",
			Alpha3:      "MRT",
			Numeric:     478,
			DialingCode: "+222",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Montserrat">Montserrat</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MS">MS</a>, MSR, 500,
		 * Officially assigned]
		 */
		"MS": CountryCode{
			Name:        "Montserrat",
			Alpha2:      "MS",
			Alpha3:      "MSR",
			Numeric:     500,
			DialingCode: "+1-664",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Malta">Malta</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MT">MT</a>, MLT, 470,
		 * Officially assigned]
		 */
		"MT": CountryCode{
			Name:        "Malta",
			Alpha2:      "MT",
			Alpha3:      "MLT",
			Numeric:     470,
			DialingCode: "+356",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mauritius">Mauritius</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MU">MU</a>, MUS, 480,
		 * Officially assigned]]
		 */
		"MU": CountryCode{
			Name:        "Mauritius",
			Alpha2:      "MU",
			Alpha3:      "MUS",
			Numeric:     480,
			DialingCode: "+230",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Maldives">Maldives</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MV">MV</a>, MDV, 462,
		 * Officially assigned]
		 */
		"MV": CountryCode{
			Name:        "Maldives",
			Alpha2:      "MV",
			Alpha3:      "MDV",
			Numeric:     462,
			DialingCode: "+960",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Malawi">Malawi</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MW">MW</a>, MWI, 454,
		 * Officially assigned]
		 */
		"MW": CountryCode{
			Name:        "Malawi",
			Alpha2:      "MW",
			Alpha3:      "MWI",
			Numeric:     454,
			DialingCode: "+265",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mexico">Mexico</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MX">MX</a>, MEX, 484,
		 * Officially assigned]
		 */
		"MX": CountryCode{
			Name:        "Mexico",
			Alpha2:      "MX",
			Alpha3:      "MEX",
			Numeric:     484,
			DialingCode: "+52",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Malaysia">Malaysia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MY">MY</a>, MYS, 458,
		 * Officially assigned]
		 */
		"MY": CountryCode{
			Name:        "Malaysia",
			Alpha2:      "MY",
			Alpha3:      "MYS",
			Numeric:     458,
			DialingCode: "+60",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mozambique">Mozambique</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MZ">MZ</a>, MOZ, 508,
		 * Officially assigned]
		 */
		"MZ": CountryCode{
			Name:        "Mozambique",
			Alpha2:      "MZ",
			Alpha3:      "MOZ",
			Numeric:     508,
			DialingCode: "+258",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Namibia">Namibia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NA">NA</a>, NAM, 516,
		 * Officially assigned]
		 */
		"NA": CountryCode{
			Name:        "Namibia",
			Alpha2:      "NA",
			Alpha3:      "NAM",
			Numeric:     516,
			DialingCode: "+264",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/New_Caledonia">New Caledonia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NC">NC</a>, NCL, 540,
		 * Officially assigned]
		 */
		"NC": CountryCode{
			Name:        "New Caledonia",
			Alpha2:      "NC",
			Alpha3:      "NCL",
			Numeric:     540,
			DialingCode: "+687",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Niger">Niger</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NE">NE</a>, NER, 562,
		 * Officially assigned]
		 */
		"NE": CountryCode{
			Name:        "Niger",
			Alpha2:      "NE",
			Alpha3:      "NER",
			Numeric:     562,
			DialingCode: "+227",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Norfolk_Island">Norfolk Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NF">NF</a>, NFK, 574,
		 * Officially assigned]
		 */
		"NF": CountryCode{
			Name:        "Norfolk Island",
			Alpha2:      "NF",
			Alpha3:      "NFK",
			Numeric:     574,
			DialingCode: "+672",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nigeria">Nigeria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NG">NG</a>, NGA, 566,
		 * Officially assigned]
		 */
		"NG": CountryCode{
			Name:        "Nigeria",
			Alpha2:      "NG",
			Alpha3:      "NGA",
			Numeric:     566,
			DialingCode: "+234",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nicaragua">Nicaragua</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NI">NI</a>, NIC, 558,
		 * Officially assigned]
		 */
		"NI": CountryCode{
			Name:        "Nicaragua",
			Alpha2:      "NI",
			Alpha3:      "NIC",
			Numeric:     558,
			DialingCode: "+505",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Netherlands">Netherlands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NL">NL</a>, NLD, 528,
		 * Officially assigned]
		 */
		"NL": CountryCode{
			Name:        "Netherlands",
			Alpha2:      "NL",
			Alpha3:      "NLD",
			Numeric:     528,
			DialingCode: "+31",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Norway">Norway</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NO">NO</a>, NOR, 578,
		 * Officially assigned]
		 */
		"NO": CountryCode{
			Name:        "Norway",
			Alpha2:      "NO",
			Alpha3:      "NOR",
			Numeric:     578,
			DialingCode: "+47",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nepal">Nepal</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NP">NP</a>, NPL, 524,
		 * Officially assigned]
		 */
		"NP": CountryCode{
			Name:        "Nepal",
			Alpha2:      "NP",
			Alpha3:      "NPL",
			Numeric:     524,
			DialingCode: "+977",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nauru">Nauru</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NR">NR</a>, NRU, 520,
		 * Officially assigned]
		 */
		"NR": CountryCode{
			Name:        "Nauru",
			Alpha2:      "NR",
			Alpha3:      "NRU",
			Numeric:     520,
			DialingCode: "+674",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saudi%E2%80%93Iraqi_neutral_zone">Neutral Zone</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NT">NT</a>, NTHH, 536,
		 * Traditionally reserved]
		 */
		"NT": CountryCode{
			Name:        "Neutral Zone",
			Alpha2:      "NT",
			Alpha3:      "NTHH",
			Numeric:     536,
			DialingCode: "",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Niue">Niue</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NU">NU</a>, NIU, 570,
		 * Officially assigned]
		 */
		"NU": CountryCode{
			Name:        "Niue",
			Alpha2:      "NU",
			Alpha3:      "NIU",
			Numeric:     570,
			DialingCode: "+683",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/New_Zealand">New Zealand</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NZ">NZ</a>, NZL, 554,
		 * Officially assigned]
		 */
		"NZ": CountryCode{
			Name:        "New Zealand",
			Alpha2:      "NZ",
			Alpha3:      "NZL",
			Numeric:     554,
			DialingCode: "+64",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href=http://en.wikipedia.org/wiki/Oman"">Oman</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#OM">OM</a>, OMN, 512,
		 * Officially assigned]
		 */
		"OM": CountryCode{
			Name:        "Oman",
			Alpha2:      "OM",
			Alpha3:      "OMN",
			Numeric:     512,
			DialingCode: "+968",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Panama">Panama</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PA">PA</a>, PAN, 591,
		 * Officially assigned]
		 */
		"PA": CountryCode{
			Name:        "Panama",
			Alpha2:      "PA",
			Alpha3:      "PAN",
			Numeric:     591,
			DialingCode: "+507",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Peru">Peru</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PE">PE</a>, PER, 604,
		 * Officially assigned]
		 */
		"PE": CountryCode{
			Name:        "Peru",
			Alpha2:      "PE",
			Alpha3:      "PER",
			Numeric:     604,
			DialingCode: "+51",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/French_Polynesia">French Polynesia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PF">PF</a>, PYF, 258,
		 * Officially assigned]
		 */
		"PF": CountryCode{
			Name:        "French Polynesia",
			Alpha2:      "PF",
			Alpha3:      "PYF",
			Numeric:     258,
			DialingCode: "+689",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Papua_New_Guinea">Papua New Guinea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PG">PG</a>, PNG, 598,
		 * Officially assigned]
		 */
		"PG": CountryCode{
			Name:        "Papua New Guinea",
			Alpha2:      "PG",
			Alpha3:      "PNG",
			Numeric:     598,
			DialingCode: "+675",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Philippines">Philippines</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PH">PH</a>, PHL, 608,
		 * Officially assigned]
		 */
		"PH": CountryCode{
			Name:        "Philippines",
			Alpha2:      "PH",
			Alpha3:      "PHL",
			Numeric:     608,
			DialingCode: "+63",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Pakistan">Pakistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PK">PK</a>, PAK, 586,
		 * Officially assigned]
		 */
		"PK": CountryCode{
			Name:        "Pakistan",
			Alpha2:      "PK",
			Alpha3:      "PAK",
			Numeric:     586,
			DialingCode: "+92",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Poland">Poland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PL">PL</a>, POL, 616,
		 * Officially assigned]
		 */
		"PL": CountryCode{
			Name:        "Poland",
			Alpha2:      "PL",
			Alpha3:      "POL",
			Numeric:     616,
			DialingCode: "+48",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Pierre_and_Miquelon">Saint Pierre and Miquelon</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PM">PM</a>, SPM, 666,
		 * Officially assigned]
		 */
		"PM": CountryCode{
			Name:        "Saint Pierre and Miquelon",
			Alpha2:      "PM",
			Alpha3:      "SPM",
			Numeric:     666,
			DialingCode: "+508",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Pitcairn_Islands">Pitcairn</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PN">PN</a>, PCN, 612,
		 * Officially assigned]
		 */
		"PN": CountryCode{
			Name:        "Pitcairn",
			Alpha2:      "PN",
			Alpha3:      "PCN",
			Numeric:     612,
			DialingCode: "+64",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Puerto_Rico">Puerto Rico</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PR">PR</a>, PRI, 630,
		 * Officially assigned]
		 */
		"PR": CountryCode{
			Name:        "Puerto Rico",
			Alpha2:      "PR",
			Alpha3:      "PRI",
			Numeric:     630,
			DialingCode: "+1-787, +1-939",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Palestinian_territories">Palestine, State of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PS">PS</a>, PSE, 275,
		 * Officially assigned]
		 */
		"PS": CountryCode{
			Name:        "Palestine, State of",
			Alpha2:      "PS",
			Alpha3:      "PSE",
			Numeric:     275,
			DialingCode: "+970",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Portugal">Portugal</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PT">PT</a>, PRT, 620,
		 * Officially assigned]
		 */
		"PT": CountryCode{
			Name:        "Portugal",
			Alpha2:      "PT",
			Alpha3:      "PRT",
			Numeric:     620,
			DialingCode: "+351",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Palau">Palau</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PW">PW</a>, PLW, 585,
		 * Officially assigned]
		 */
		"PW": CountryCode{
			Name:        "Palau",
			Alpha2:      "PW",
			Alpha3:      "PLW",
			Numeric:     585,
			DialingCode: "+680",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Paraguay">Paraguay</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PY">PY</a>, PRY, 600,
		 * Officially assigned]
		 */
		"PY": CountryCode{
			Name:        "Paraguay",
			Alpha2:      "PY",
			Alpha3:      "PRY",
			Numeric:     600,
			DialingCode: "+595",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Qatar">Qatar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#QA">QA</a>, QAT, 634,
		 * Officially assigned]
		 */
		"QA": CountryCode{
			Name:        "Qatar",
			Alpha2:      "QA",
			Alpha3:      "QAT",
			Numeric:     634,
			DialingCode: "+974",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/R%C3%A9union">R&eacute;union</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RE">RE</a>, REU, 638,
		 * Officially assigned]
		 */
		"RE": CountryCode{
			Name:        "R\u00E9union",
			Alpha2:      "RE",
			Alpha3:      "REU",
			Numeric:     638,
			DialingCode: "+262",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Romania">Romania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RO">RO</a>, ROU, 642,
		 * Officially assigned]
		 */
		"RO": CountryCode{
			Name:        "Romania",
			Alpha2:      "RO",
			Alpha3:      "ROU",
			Numeric:     642,
			DialingCode: "+40",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Serbia">Serbia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RS">RS</a>, SRB, 688,
		 * Officially assigned]
		 */
		"RS": CountryCode{
			Name:        "Serbia",
			Alpha2:      "RS",
			Alpha3:      "SRB",
			Numeric:     688,
			DialingCode: "+381",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Russia">Russian Federation</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RU">RU</a>, RUS, 643,
		 * Officially assigned]
		 */
		"RU": CountryCode{
			Name:        "Russian Federation",
			Alpha2:      "RU",
			Alpha3:      "RUS",
			Numeric:     643,
			DialingCode: "+7",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Rwanda">Rwanda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RW">RW</a>, RWA, 646,
		 * Officially assigned]
		 */
		"RW": CountryCode{
			Name:        "Rwanda",
			Alpha2:      "RW",
			Alpha3:      "RWA",
			Numeric:     646,
			DialingCode: "+250",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saudi_Arabia">Saudi Arabia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SA">SA</a>, SAU, 682,
		 * Officially assigned]
		 */
		"SA": CountryCode{
			Name:        "Saudi Arabia",
			Alpha2:      "SA",
			Alpha3:      "SAU",
			Numeric:     682,
			DialingCode: "+966",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Solomon_Islands">Solomon Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SB">SB</a>, SLB, 90,
		 * Officially assigned]
		 */
		"SB": CountryCode{
			Name:        "Solomon Islands",
			Alpha2:      "SB",
			Alpha3:      "SLB",
			Numeric:     90,
			DialingCode: "+677",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Seychelles">Seychelles</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SC">SC</a>, SYC, 690,
		 * Officially assigned]
		 */
		"SC": CountryCode{
			Name:        "Seychelles",
			Alpha2:      "SC",
			Alpha3:      "SYC",
			Numeric:     690,
			DialingCode: "+248",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sudan">Sudan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SD">SD</a>, SDN, 729,
		 * Officially assigned]
		 */
		"SD": CountryCode{
			Name:        "Sudan",
			Alpha2:      "SD",
			Alpha3:      "SDN",
			Numeric:     729,
			DialingCode: "+249",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sweden">Sweden</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SE">SE</a>, SWE, 752,
		 * Officially assigned]
		 */
		"SE": CountryCode{
			Name:        "Sweden",
			Alpha2:      "SE",
			Alpha3:      "SWE",
			Numeric:     752,
			DialingCode: "+46",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Finland">Finland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SF">SF</a>, FIN, 246,
		 * Traditionally reserved]
		 *
		 * @see #FI
		 */
		"SF": CountryCode{
			Name:        "Finland",
			Alpha2:      "SF",
			Alpha3:      "FIN",
			Numeric:     246,
			DialingCode: "+358",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Singapore">Singapore</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SG">SG</a>, SGP, 702,
		 * Officially assigned]
		 */
		"SG": CountryCode{
			Name:        "Singapore",
			Alpha2:      "SG",
			Alpha3:      "SGP",
			Numeric:     702,
			DialingCode: "+65",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Helena,_Ascension_and_Tristan_da_Cunha">Saint Helena, Ascension and Tristan da Cunha</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SH">SH</a>, SHN, 654,
		 * Officially assigned]
		 */
		"SH": CountryCode{
			Name:        "Saint Helena, Ascension and Tristan da Cunha",
			Alpha2:      "SH",
			Alpha3:      "SHN",
			Numeric:     654,
			DialingCode: "+290",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Slovenia">Slovenia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SI">SI</a>, SVN, 705,
		 * Officially assigned]
		 */
		"SI": CountryCode{
			Name:        "Slovenia",
			Alpha2:      "SI",
			Alpha3:      "SVN",
			Numeric:     705,
			DialingCode: "+386",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Svalbard_and_Jan_Mayen">Svalbard and Jan Mayen</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SJ">SJ</a>, SJM, 744,
		 * Officially assigned]
		 */
		"SJ": CountryCode{
			Name:        "Svalbard and Jan Mayen",
			Alpha2:      "SJ",
			Alpha3:      "SJM",
			Numeric:     744,
			DialingCode: "+47",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Slovakia">Slovakia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SK">SK</a>, SVK, 703,
		 * Officially assigned]
		 */
		"SK": CountryCode{
			Name:        "Slovakia",
			Alpha2:      "SK",
			Alpha3:      "SVK",
			Numeric:     703,
			DialingCode: "+421",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sierra_Leone">Sierra Leone</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SL">SL</a>, SLE, 694,
		 * Officially assigned]
		 */
		"SL": CountryCode{
			Name:        "Sierra Leone",
			Alpha2:      "SL",
			Alpha3:      "SLE",
			Numeric:     694,
			DialingCode: "+232",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/San_Marino">San Marino</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SM">SM</a>, SMR, 674,
		 * Officially assigned]
		 */
		"SM": CountryCode{
			Name:        "San Marino",
			Alpha2:      "SM",
			Alpha3:      "SMR",
			Numeric:     674,
			DialingCode: "+378",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Senegal">Senegal</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SN">SN</a>, SEN, 686,
		 * Officially assigned]
		 */
		"SN": CountryCode{
			Name:        "Senegal",
			Alpha2:      "SN",
			Alpha3:      "SEN",
			Numeric:     686,
			DialingCode: "+221",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Somalia">Somalia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SO">SO</a>, SOM, 706,
		 * Officially assigned]
		 */
		"SO": CountryCode{
			Name:        "Somalia",
			Alpha2:      "SO",
			Alpha3:      "SOM",
			Numeric:     706,
			DialingCode: "+252",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Suriname">Suriname</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SR">SR</a>, SUR, 740,
		 * Officially assigned]
		 */
		"SR": CountryCode{
			Name:        "Suriname",
			Alpha2:      "SR",
			Alpha3:      "SUR",
			Numeric:     740,
			DialingCode: "+597",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Sudan">South Sudan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SS">SS</a>, SSD, 728,
		 * Officially assigned]
		 */
		"SS": CountryCode{
			Name:        "South Sudan",
			Alpha2:      "SS",
			Alpha3:      "SSD",
			Numeric:     728,
			DialingCode: "+211",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/S%C3%A3o_Tom%C3%A9_and_Pr%C3%ADncipe">Sao Tome and Principe</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ST">ST</a>, STP, 678,
		 * Officially assigned]
		 */
		"ST": CountryCode{
			Name:        "Sao Tome and Principe",
			Alpha2:      "ST",
			Alpha3:      "STP",
			Numeric:     678,
			DialingCode: "+239",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Soviet_Union">USSR</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SU">SU</a>, SUN, -1,
		 * Exceptionally reserved]
		 */
		"SU": CountryCode{
			Name:        "USSR",
			Alpha2:      "SU",
			Alpha3:      "SUN",
			Numeric:     -1,
			DialingCode: "+7",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/El_Salvador">El Salvador</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SV">SV</a>, SLV, 222,
		 * Officially assigned]
		 */
		"SV": CountryCode{
			Name:        "El Salvador",
			Alpha2:      "SV",
			Alpha3:      "SLV",
			Numeric:     222,
			DialingCode: "+503",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sint_Maarten">Sint Maarten (Dutch part)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SX">SX</a>, SXM, 534,
		 * Officially assigned]
		 */
		"SX": CountryCode{
			Name:        "Sint Maarten (Dutch part)",
			Alpha2:      "SX",
			Alpha3:      "SXM",
			Numeric:     534,
			DialingCode: "+1-721",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Syria">Syrian Arab Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SY">SY</a>, SYR, 760,
		 * Officially assigned]
		 */
		"SY": CountryCode{
			Name:        "Syrian Arab Republic",
			Alpha2:      "SY",
			Alpha3:      "SYR",
			Numeric:     760,
			DialingCode: "+963",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Swaziland">Swaziland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SZ">SZ</a>, SWZ, 748,
		 * Officially assigned]
		 */
		"SZ": CountryCode{
			Name:        "Swaziland",
			Alpha2:      "SZ",
			Alpha3:      "SWZ",
			Numeric:     748,
			DialingCode: "+268",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tristan_da_Cunha">Tristan da Cunha</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TA">TA</a>, TAA, -1,
		 * Exceptionally reserved.
		 */
		"TA": CountryCode{
			Name:        "Tristan da Cunha",
			Alpha2:      "TA",
			Alpha3:      "TAA",
			Numeric:     -1,
			DialingCode: "+290-8",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Turks_and_Caicos_Islands">Turks and Caicos Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TC">TC</a>, TCA, 796,
		 * Officially assigned]
		 */
		"TC": CountryCode{
			Name:        "Turks and Caicos Islands",
			Alpha2:      "TC",
			Alpha3:      "TCA",
			Numeric:     796,
			DialingCode: "+1-649",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Chad">Chad</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TD">TD</a>, TCD, 148,
		 * Officially assigned]
		 */
		"TD": CountryCode{
			Name:        "Chad",
			Alpha2:      "TD",
			Alpha3:      "TCD",
			Numeric:     148,
			DialingCode: "+235",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/French_Southern_and_Antarctic_Lands">French Southern Territories</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TF">TF</a>, ATF, 260,
		 * Officially assigned]
		 */
		"TF": CountryCode{
			Name:        "French Southern Territories",
			Alpha2:      "TF",
			Alpha3:      "ATF",
			Numeric:     260,
			DialingCode: "",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Togo">Togo</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TG">TG</a>, TGO, 768,
		 * Officially assigned]
		 */
		"TG": CountryCode{
			Name:        "Togo",
			Alpha2:      "TG",
			Alpha3:      "TGO",
			Numeric:     768,
			DialingCode: "228",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Thailand">Thailand</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TH">TH</a>, THA, 764,
		 * Officially assigned]
		 */
		"TH": CountryCode{
			Name:        "Thailand",
			Alpha2:      "TH",
			Alpha3:      "THA",
			Numeric:     764,
			DialingCode: "+66",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tajikistan">Tajikistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TJ">TJ</a>, TJK, 762,
		 * Officially assigned]
		 */
		"TJ": CountryCode{
			Name:        "Tajikistan",
			Alpha2:      "TJ",
			Alpha3:      "TJK",
			Numeric:     762,
			DialingCode: "+992",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tokelau">Tokelau</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TK">TK</a>, TKL, 772,
		 * Officially assigned]
		 */
		"TK": CountryCode{
			Name:        "Tokelau",
			Alpha2:      "TK",
			Alpha3:      "TKL",
			Numeric:     772,
			DialingCode: "+690",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/East_Timor">Timor-Leste</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TL">TL</a>, TLS, 626,
		 * Officially assigned]
		 */
		"TL": CountryCode{
			Name:        "Timor-Leste",
			Alpha2:      "TL",
			Alpha3:      "TLS",
			Numeric:     626,
			DialingCode: "+670",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Turkmenistan">Turkmenistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TM">TM</a>, TKM, 795,
		 * Officially assigned]
		 */
		"TM": CountryCode{
			Name:        "Turkmenistan",
			Alpha2:      "TM",
			Alpha3:      "TKM",
			Numeric:     795,
			DialingCode: "+993",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tunisia">Tunisia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TN">TN</a>, TUN, 788,
		 * Officially assigned]
		 */
		"TN": CountryCode{
			Name:        "Tunisia",
			Alpha2:      "TN",
			Alpha3:      "TUN",
			Numeric:     788,
			DialingCode: "+216",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tonga">Tonga</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TO">TO</a>, TON, 776,
		 * Officially assigned]
		 */
		"TO": CountryCode{
			Name:        "Tonga",
			Alpha2:      "TO",
			Alpha3:      "TON",
			Numeric:     776,
			DialingCode: "+676",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/East_Timor">East Timor</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TP">TP</a>, TPTL, 0,
		 * Traditionally reserved]
		 *
		 * <p>
		 * ISO 3166-1 numeric code is unknown.
		 * </p>
		 */
		"TP": CountryCode{
			Name:        "East Timor",
			Alpha2:      "TP",
			Alpha3:      "TPTL",
			Numeric:     0,
			DialingCode: "+670",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Turkey">Turkey</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TR">TR</a>, TUR, 792,
		 * Officially assigned]
		 */
		"TR": CountryCode{
			Name:        "Turkey",
			Alpha2:      "TR",
			Alpha3:      "TUR",
			Numeric:     792,
			DialingCode: "+90",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Trinidad_and_Tobago">Trinidad and Tobago</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TT">TT</a>, TTO, 780,
		 * Officially assigned]
		 */
		"TT": CountryCode{
			Name:        "Trinidad and Tobago",
			Alpha2:      "TT",
			Alpha3:      "TTO",
			Numeric:     780,
			DialingCode: "+1-868",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tuvalu">Tuvalu</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TV">TV</a>, TUV, 798,
		 * Officially assigned]
		 */
		"TV": CountryCode{
			Name:        "Tuvalu",
			Alpha2:      "TV",
			Alpha3:      "TUV",
			Numeric:     798,
			DialingCode: "+688",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Taiwan">Taiwan, Province of China</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TW">TW</a>, TWN, 158,
		 * Officially assigned]
		 */
		"TW": CountryCode{
			Name:        "Taiwan, Province of China",
			Alpha2:      "TW",
			Alpha3:      "TWN",
			Numeric:     158,
			DialingCode: "+886",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tanzania">Tanzania, United Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TZ">TZ</a>, TZA, 834,
		 * Officially assigned]
		 */
		"TZ": CountryCode{
			Name:        "Tanzania, United Republic of",
			Alpha2:      "TZ",
			Alpha3:      "TZA",
			Numeric:     834,
			DialingCode: "+255",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ukraine">Ukraine</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UA">UA</a>, UKR, 804,
		 * Officially assigned]
		 */
		"UA": CountryCode{
			Name:        "Ukraine",
			Alpha2:      "UA",
			Alpha3:      "UKR",
			Numeric:     804,
			DialingCode: "+380",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Uganda">Uganda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UG">UG</a>, UGA, 800,
		 * Officially assigned]
		 */
		"UG": CountryCode{
			Name:        "Uganda",
			Alpha2:      "UG",
			Alpha3:      "UGA",
			Numeric:     800,
			DialingCode: "+256",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_Kingdom">United Kingdom</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UK">UK</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"UK": CountryCode{
			Name:        "United Kingdom",
			Alpha2:      "UK",
			Alpha3:      "",
			Numeric:     -1,
			DialingCode: "+44",
			Assignment:  EXCEPTIONALLY_RESERVED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/United_States_Minor_Outlying_Islands">United States Minor Outlying Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UM">UM</a>, UMI, 581,
		 * Officially assigned]
		 */
		"UM": CountryCode{
			Name:        "United States Minor Outlying Islands",
			Alpha2:      "UM",
			Alpha3:      "UMI",
			Numeric:     581,
			DialingCode: "+1",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_States">United States</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#US">US</a>, USA, 840,
		 * Officially assigned]
		 */
		"US": CountryCode{
			Name:        "United States",
			Alpha2:      "US",
			Alpha3:      "USA",
			Numeric:     840,
			DialingCode: "+1",
			Assignment:  OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Uruguay">Uruguay</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UY">UY</a>, URY, 858,
		 * Officially assigned]
		 */
		"UY": CountryCode{
			Name:        "Uruguay",
			Alpha2:      "UY",
			Alpha3:      "URY",
			Numeric:     858,
			DialingCode: "+598",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Uzbekistan">Uzbekistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UZ">UZ</a>, UZB, 860,
		 * Officially assigned]
		 */
		"UZ": CountryCode{
			Name:        "Uzbekistan",
			Alpha2:      "UZ",
			Alpha3:      "UZB",
			Numeric:     860,
			DialingCode: "+998",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Vatican_City">Holy See (Vatican City State)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VA">VA</a>, VAT, 336,
		 * Officially assigned]
		 */
		"VA": CountryCode{
			Name:        "Holy See (Vatican City State)",
			Alpha2:      "VA",
			Alpha3:      "VAT",
			Numeric:     336,
			DialingCode: "+379",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Vincent_and_the_Grenadines">Saint Vincent and the Grenadines</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VC">VC</a>, VCT, 670,
		 * Officially assigned]
		 */
		"VC": CountryCode{
			Name:        "Saint Vincent and the Grenadines",
			Alpha2:      "VC",
			Alpha3:      "VCT",
			Numeric:     670,
			DialingCode: "+1-784",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Venezuela">Venezuela, Bolivarian Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VE">VE</a>, VEN, 862,
		 * Officially assigned]
		 */
		"VE": CountryCode{
			Name:        "Venezuela, Bolivarian Republic of",
			Alpha2:      "VE",
			Alpha3:      "VEN",
			Numeric:     862,
			DialingCode: "+58",

			Assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/British_Virgin_Islands">Virgin Islands, British</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VG">VG</a>, VGB, 92,
		 * Officially assigned]
		 */
		"VG": CountryCode{
			Name:        "Virgin Islands, British",
			Alpha2:      "VG",
			Alpha3:      "VGB",
			Numeric:     92,
			DialingCode: "+1-284",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_States_Virgin_Islands">Virgin Islands, U.S.</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VI">VI</a>, VIR, 850,
		 * Officially assigned]
		 */
		"VI": CountryCode{
			Name:        "Virgin Islands, U.S.",
			Alpha2:      "VI",
			Alpha3:      "VIR",
			Numeric:     850,
			DialingCode: "+1-340",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Vietnam">Viet Nam</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VN">VN</a>, VNM, 704,
		 * Officially assigned]
		 */
		"VN": CountryCode{
			Name:        "Viet Nam",
			Alpha2:      "VN",
			Alpha3:      "VNM",
			Numeric:     704,
			DialingCode: "+84",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Vanuatu">Vanuatu</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VU">VU</a>, VUT, 548,
		 * Officially assigned]
		 */
		"VU": CountryCode{
			Name:        "Vanuatu",
			Alpha2:      "VU",
			Alpha3:      "VUT",
			Numeric:     548,
			DialingCode: "+678",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Wallis_and_Futuna">Wallis and Futuna</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#WF">WF</a>, WLF, 876,
		 * Officially assigned]
		 */
		"WF": CountryCode{
			Name:        "Wallis and Futuna",
			Alpha2:      "WF",
			Alpha3:      "WLF",
			Numeric:     876,
			DialingCode: "+681",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Samoa">Samoa</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#WS">WS</a>, WSM, 882,
		 * Officially assigned]
		 */
		"WS": CountryCode{
			Name:        "Samoa",
			Alpha2:      "WS",
			Alpha3:      "WSM",
			Numeric:     882,
			DialingCode: "+685",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kosovo">Kosovo, Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#XK">XK</a>, XXK, -1,
		 * User assigned]
		 */
		"XK": CountryCode{
			Name:        "Kosovo, Republic of",
			Alpha2:      "XK",
			Alpha3:      "XXK",
			Numeric:     -1,
			DialingCode: "+383",
			Assignment:  USER_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Yemen">Yemen</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#YE">YE</a>, YEM, 887,
		 * Officially assigned]
		 */
		"YE": CountryCode{
			Name:        "Yemen",
			Alpha2:      "YE",
			Alpha3:      "YEM",
			Numeric:     887,
			DialingCode: "+967",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mayotte">Mayotte</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#YT">YT</a>, MYT, 175,
		 * Officially assigned]
		 */
		"YT": CountryCode{
			Name:        "Mayotte",
			Alpha2:      "YT",
			Alpha3:      "MYT",
			Numeric:     175,
			DialingCode: "+262",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Yugoslavia">Yugoslavia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#YU">YU</a>, YUCS, 890,
		 * Traditionally reserved]
		 */
		"YU": CountryCode{
			Name:        "Yugoslavia",
			Alpha2:      "YU",
			Alpha3:      "YUCS",
			Numeric:     890,
			DialingCode: "+38",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Africa">South Africa</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ZA">ZA</a>, ZAF, 710,
		 * Officially assigned]
		 */
		"ZA": CountryCode{
			Name:        "South Africa",
			Alpha2:      "ZA",
			Alpha3:      "ZAF",
			Numeric:     710,
			DialingCode: "+27",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Zambia">Zambia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ZM">ZM</a>, ZMB, 894,
		 * Officially assigned]
		 */
		"ZM": CountryCode{
			Name:        "Zambia",
			Alpha2:      "ZM",
			Alpha3:      "ZMB",
			Numeric:     894,
			DialingCode: "+260",
			Assignment:  OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Zaire">Zaire</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ZR">ZR</a>, ZRCD, 0,
		 * Traditionally reserved]
		 *
		 * <p>
		 * ISO 3166-1 numeric code is unknown.
		 * </p>
		 */
		"ZR": CountryCode{
			Name:        "Zaire",
			Alpha2:      "ZR",
			Alpha3:      "ZRCD",
			Numeric:     0,
			DialingCode: "+243",
			Assignment:  TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Zimbabwe">Zimbabwe</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ZW">ZW</a>, ZWE, 716,
		 * Officially assigned]
		 */
		"ZW": CountryCode{
			Name:        "Zimbabwe",
			Alpha2:      "ZW",
			Alpha3:      "ZWE",
			Numeric:     716,
			DialingCode: "+263",
			Assignment:  OFFICIALLY_ASSIGNED,
		},
	}

	for _, cc := range by_alpha2 {
		if cc.Alpha3 != "" {
			by_alpha3[cc.Alpha3] = cc
		}
		by_name[cc.Name] = cc
		by_numeric[cc.Numeric] = cc
		name_trie.Insert(patricia.Prefix(strings.ToLower(cc.Name)), cc)
	}
}

func GetByAlpha2(a2 string) (CountryCode, bool) {
	code := by_alpha2[a2]

	return code, code.Alpha2 != ""
}

func GetByAlpha3(a3 string) (CountryCode, bool) {
	code := by_alpha3[a3]

	return code, code.Alpha2 != ""
}

func GetByName(name string) (CountryCode, bool) {
	code := by_name[name]

	return code, code.Alpha2 != ""
}

func FindByName(prefix string) (matches []CountryCode) {
	matches = make([]CountryCode, 0)

	visit := func(prefix patricia.Prefix, item patricia.Item) error {
		matches = append(matches, item.(CountryCode))
		return nil
	}

	name_trie.VisitSubtree(patricia.Prefix(strings.ToLower(prefix)), visit)

	return
}
