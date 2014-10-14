package countrycodes

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
	name       string
	alpha2     string
	alpha3     string
	numeric    int
	assignment Assignment
}

var by_alpha2 map[string]CountryCode

var by_name map[string]CountryCode

var by_alpha3 map[string]CountryCode

var by_numeric map[int]CountryCode

func init() {

	by_name = make(map[string]CountryCode)
	by_alpha3 = make(map[string]CountryCode)
	by_numeric = make(map[int]CountryCode)

	by_alpha2 = map[string]CountryCode{
		/**
		 * <a href="http://en.wikipedia.org/wiki/Ascension_Island">Ascension Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AC">AC</a>, ASC, -1,
		 * Exceptionally reserved]
		 */
		"AC": CountryCode{
			name:       "Ascension Island",
			alpha2:     "AC",
			alpha3:     "ASC",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Andorra">Andorra</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AD">AD</a>, AND, 16,
		 * Officially assigned]
		 */
		"AD": CountryCode{
			name:       "Andorra",
			alpha2:     "AD",
			alpha3:     "AND",
			numeric:    20,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_Arab_Emirates">United Arab Emirates</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AE">AE</a>, AE, 784,
		 * Officially assigned]
		 */
		"AE": CountryCode{
			name:       "United Arab Emirates",
			alpha2:     "AE",
			alpha3:     "ARE",
			numeric:    784,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Afghanistan">Afghanistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AF">AF</a>, AFG, 4,
		 * Officially assigned]
		 */
		"AF": CountryCode{
			name:       "Afghanistan",
			alpha2:     "AF",
			alpha3:     "AFG",
			numeric:    4,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Antigua_and_Barbuda">Antigua and Barbuda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AG">AG</a>, ATG, 28,
		 * Officially assigned]
		 */
		"AG": CountryCode{
			name:       "Antigua and Barbuda",
			alpha2:     "AG",
			alpha3:     "ATG",
			numeric:    28,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Anguilla">Anguilla</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AI">AI</a>, AIA, 660,
		 * Officially assigned]
		 */
		"AI": CountryCode{
			name:       "Anguilla",
			alpha2:     "AI",
			alpha3:     "AIA",
			numeric:    660,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Albania">Albania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AL">AL</a>, ALB, 8,
		 * Officially assigned]
		 */
		"AL": CountryCode{
			name:       "Albania",
			alpha2:     "AL",
			alpha3:     "ALB",
			numeric:    8,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Armenia">Armenia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AM">AM</a>, ARM, 51,
		 * Officially assigned]
		 */
		"AM": CountryCode{
			name:       "Armenia",
			alpha2:     "AM",
			alpha3:     "ARM",
			numeric:    51,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Netherlands_Antilles">Netherlands Antilles</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AN">AN</a>, ANHH, 530,
		 * Traditionally reserved]
		 */
		"AN": CountryCode{
			name:       "Netherlands Antilles",
			alpha2:     "AN",
			alpha3:     "ANHH",
			numeric:    530,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Angola">Angola</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AO">AO</a>, AGO, 24,
		 * Officially assigned]
		 */
		"AO": CountryCode{
			name:       "Angola",
			alpha2:     "AO",
			alpha3:     "AGO",
			numeric:    24,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Antarctica">Antarctica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AQ">AQ</a>, ATA, 10,
		 * Officially assigned]
		 */
		"AQ": CountryCode{
			name:       "Antarctica",
			alpha2:     "AQ",
			alpha3:     "ATA",
			numeric:    10,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Argentina">Argentina</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AR">AR</a>, ARG, 32,
		 * Officially assigned]
		 */
		"AR": CountryCode{
			name:       "Argentina",
			alpha2:     "AR",
			alpha3:     "ARG",
			numeric:    32,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/American_Samoa">American Samoa</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AS">AS</a>, ASM, 16,
		 * Officially assigned]
		 */
		"AS": CountryCode{
			name:       "American Samoa",
			alpha2:     "AS",
			alpha3:     "ASM",
			numeric:    16,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Austria">Austria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AT">AT</a>, AUT, 40,
		 * Officially assigned]
		 */
		"AT": CountryCode{
			name:       "Austria",
			alpha2:     "AT",
			alpha3:     "AUT",
			numeric:    40,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Australia">Australia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AU">AU</a>, AUS, 36,
		 * Officially assigned]
		 */
		"AU": CountryCode{
			name:       "Australia",
			alpha2:     "AU",
			alpha3:     "AUS",
			numeric:    36,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Aruba">Aruba</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AW">AW</a>, ABW, 533,
		 * Officially assigned]
		 */
		"AW": CountryCode{
			name:       "Aruba",
			alpha2:     "AW",
			alpha3:     "ABW",
			numeric:    533,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/%C3%85land_Islands">&Aring;land Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AX">AX</a>, ALA, 248,
		 * Officially assigned]
		 */
		"AX": CountryCode{
			name:       "\u212Bland Islands",
			alpha2:     "AX",
			alpha3:     "ALA",
			numeric:    248,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Azerbaijan">Azerbaijan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#AZ">AZ</a>, AZE, 31,
		 * Officially assigned]
		 */
		"AZ": CountryCode{
			name:       "Azerbaijan",
			alpha2:     "AZ",
			alpha3:     "AZE",
			numeric:    31,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bosnia_and_Herzegovina">Bosnia and Herzegovina</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BA">BA</a>, BIH, 70,
		 * Officially assigned]
		 */
		"BA": CountryCode{
			name:       "Bosnia and Herzegovina",
			alpha2:     "BA",
			alpha3:     "BIH",
			numeric:    70,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Barbados">Barbados</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BB">BB</a>, BRB, 52,
		 * Officially assigned]
		 */
		"BB": CountryCode{
			name:       "Barbados",
			alpha2:     "BB",
			alpha3:     "BRB",
			numeric:    52,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bangladesh">Bangladesh</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BD">BD</a>, BGD, 50,
		 * Officially assigned]
		 */
		"BD": CountryCode{
			name:       "Bangladesh",
			alpha2:     "BD",
			alpha3:     "BGD",
			numeric:    50,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Belgium">Belgium</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BE">BE</a>, BEL, 56,
		 * Officially assigned]
		 */
		"BE": CountryCode{
			name:       "Belgium",
			alpha2:     "BE",
			alpha3:     "BEL",
			numeric:    56,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Burkina_Faso">Burkina Faso</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BF">BF</a>, BFA, 854,
		 * Officially assigned]
		 */
		"BF": CountryCode{
			name:       "Burkina Faso",
			alpha2:     "BF",
			alpha3:     "BFA",
			numeric:    854,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bulgaria">Bulgaria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BG">BG</a>, BGR, 100,
		 * Officially assigned]
		 */
		"BG": CountryCode{
			name:       "Bulgaria",
			alpha2:     "BG",
			alpha3:     "BGR",
			numeric:    100,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bahrain">Bahrain</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BH">BH</a>, BHR, 48,
		 * Officially assigned]
		 */
		"BH": CountryCode{
			name:       "Bahrain",
			alpha2:     "BH",
			alpha3:     "BHR",
			numeric:    48,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Burundi">Burundi</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BI">BI</a>, BDI, 108,
		 * Officially assigned]
		 */
		"BI": CountryCode{
			name:       "Burundi",
			alpha2:     "BI",
			alpha3:     "BDI",
			numeric:    108,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Benin">Benin</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BJ">BJ</a>, BEN, 204,
		 * Officially assigned]
		 */
		"BJ": CountryCode{
			name:       "Benin",
			alpha2:     "BJ",
			alpha3:     "BEN",
			numeric:    204,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Barth%C3%A9lemy">Saint Barth&eacute;lemy</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BL">BL</a>, BLM, 652,
		 * Officially assigned]
		 */
		"BL": CountryCode{
			name:       "Saint Barth\u00E9lemy",
			alpha2:     "BL",
			alpha3:     "BLM",
			numeric:    652,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bermuda">Bermuda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BM">BM</a>, BMU, 60,
		 * Officially assigned]
		 */
		"BM": CountryCode{
			name:       "Bermuda",
			alpha2:     "BM",
			alpha3:     "BMU",
			numeric:    60,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Brunei">Brunei Darussalam</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BN">BN</a>, BRN, 96,
		 * Officially assigned]
		 */
		"BN": CountryCode{
			name:       "Brunei Darussalam",
			alpha2:     "BN",
			alpha3:     "BRN",
			numeric:    96,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bolivia">Bolivia, Plurinational State of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BO">BO</a>, BOL, 68,
		 * Officially assigned]
		 */
		"BO": CountryCode{
			name:       "Bolivia, Plurinational State of",
			alpha2:     "BO",
			alpha3:     "BOL",
			numeric:    68,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Caribbean_Netherlands">Bonaire, Sint Eustatius and Saba</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BQ">BQ</a>, BES, 535,
		 * Officially assigned]
		 */
		"BQ": CountryCode{
			name:       "Bonaire, Sint Eustatius and Saba",
			alpha2:     "BQ",
			alpha3:     "BES",
			numeric:    535,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Brazil">Brazil</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BR">BR</a>, BRA, 76,
		 * Officially assigned]
		 */
		"BR": CountryCode{
			name:       "Brazil",
			alpha2:     "BR",
			alpha3:     "BRA",
			numeric:    76,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/The_Bahamas">Bahamas</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BS">BS</a>, BHS, 44,
		 * Officially assigned]
		 */
		"BS": CountryCode{
			name:       "Bahamas",
			alpha2:     "BS",
			alpha3:     "BHS",
			numeric:    44,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bhutan">Bhutan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BT">BT</a>, BTN, 64,
		 * Officially assigned]
		 */
		"BT": CountryCode{
			name:       "Bhutan",
			alpha2:     "BT",
			alpha3:     "BTN",
			numeric:    64,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Burma">Burma</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BU">BU</a>, BUMM, 104,
		 * Officially assigned]
		 *
		 * @see #MM
		 */
		"BU": CountryCode{
			name:       "Burma",
			alpha2:     "BU",
			alpha3:     "BUMM",
			numeric:    104,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Bouvet_Island">Bouvet Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BV">BV</a>, BVT, 74,
		 * Officially assigned]
		 */
		"BV": CountryCode{
			name:       "Bouvet Island",
			alpha2:     "BV",
			alpha3:     "BVT",
			numeric:    74,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Botswana">Botswana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BW">BW</a>, BWA, 72,
		 * Officially assigned]
		 */
		"BW": CountryCode{
			name:       "Botswana",
			alpha2:     "BW",
			alpha3:     "BWA",
			numeric:    72,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Belarus">Belarus</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BY">BY</a>, BLR, 112,
		 * Officially assigned]
		 */
		"BY": CountryCode{
			name:       "Belarus",
			alpha2:     "BY",
			alpha3:     "BLR",
			numeric:    112,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Belize">Belize</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#BZ">BZ</a>, BLZ, 84,
		 * Officially assigned]
		 */
		"BZ": CountryCode{
			name:       "Belize",
			alpha2:     "BZ",
			alpha3:     "BLZ",
			numeric:    84,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Canada">Canada</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CA">CA</a>, CAN, 124,
		 * Officially assigned]
		 */
		"CA": CountryCode{
			name:       "Canada",
			alpha2:     "CA",
			alpha3:     "CAN",
			numeric:    124,
			assignment: OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Cocos_(Keeling)_Islands">Cocos (Keeling) Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CC">CC</a>, CCK, 166,
		 * Officially assigned]
		 */
		"CC": CountryCode{
			name:       "Cocos (Keeling) Islands",
			alpha2:     "CC",
			alpha3:     "CCK",
			numeric:    166,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Democratic_Republic_of_the_Congo">Congo, the Democratic Republic of the</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CD">CD</a>, COD, 180,
		 * Officially assigned]
		 */
		"CD": CountryCode{
			name:       "Congo, the Democratic Republic of the",
			alpha2:     "CD",
			alpha3:     "COD",
			numeric:    180,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Central_African_Republic">Central African Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CF">CF</a>, CAF, 140,
		 * Officially assigned]
		 */
		"CF": CountryCode{
			name:       "Central African Republic",
			alpha2:     "CF",
			alpha3:     "CAF",
			numeric:    140,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Republic_of_the_Congo">Congo</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CG">CG</a>, COG, 178,
		 * Officially assigned]
		 */
		"CG": CountryCode{
			name:       "Congo",
			alpha2:     "CG",
			alpha3:     "COG",
			numeric:    178,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Switzerland">Switzerland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CH">CH</a>, CHE, 756,
		 * Officially assigned]
		 */
		"CH": CountryCode{
			name:       "Switzerland",
			alpha2:     "CH",
			alpha3:     "CHE",
			numeric:    756,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/C%C3%B4te_d%27Ivoire">C&ocirc;te d'Ivoire</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CI">CI</a>, CIV, 384,
		 * Officially assigned]
		 */
		"CI": CountryCode{
			name:       "C\u00F4te d'Ivoire",
			alpha2:     "CI",
			alpha3:     "CIV",
			numeric:    384,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cook_Islands">Cook Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CK">CK</a>, COK, 184,
		 * Officially assigned]
		 */
		"CK": CountryCode{
			name:       "Cook Islands",
			alpha2:     "CK",
			alpha3:     "COK",
			numeric:    184,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Chile">Chile</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CL">CL</a>, CHL, 152,
		 * Officially assigned]
		 */
		"CL": CountryCode{
			name:       "Chile",
			alpha2:     "CL",
			alpha3:     "CHL",
			numeric:    152,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cameroon">Cameroon</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CM">CM</a>, CMR, 120,
		 * Officially assigned]
		 */
		"CM": CountryCode{
			name:       "Cameroon",
			alpha2:     "CM",
			alpha3:     "CMR",
			numeric:    120,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/China">China</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CN">CN</a>, CHN, 156,
		 * Officially assigned]
		 */
		"CN": CountryCode{
			name:       "China",
			alpha2:     "CN",
			alpha3:     "CHN",
			numeric:    156,
			assignment: OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Colombia">Colombia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CO">CO</a>, COL, 170,
		 * Officially assigned]
		 */
		"CO": CountryCode{
			name:       "Colombia",
			alpha2:     "CO",
			alpha3:     "COL",
			numeric:    170,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Clipperton_Island">Clipperton Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CP">CP</a>, CPT, -1,
		 * Exceptionally reserved]
		 */
		"CP": CountryCode{
			name:       "Clipperton Island",
			alpha2:     "CP",
			alpha3:     "CPT",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Costa_Rica">Costa Rica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CR">CR</a>, CRI, 188,
		 * Officially assigned]
		 */
		"CR": CountryCode{
			name:       "Costa Rica",
			alpha2:     "CR",
			alpha3:     "CRI",
			numeric:    188,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Serbia_and_Montenegro">Serbia and Montenegro</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CS">CS</a>, CSXX, 891,
		 * Traditionally reserved]
		 */
		"CS": CountryCode{
			name:       "Serbia and Montenegro",
			alpha2:     "CS",
			alpha3:     "CSXX",
			numeric:    891,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cuba">Cuba</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CU">CU</a>, CUB, 192,
		 * Officially assigned]
		 */
		"CU": CountryCode{
			name:       "Cuba",
			alpha2:     "CU",
			alpha3:     "CUB",
			numeric:    192,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cape_Verde">Cape Verde</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CV">CV</a>, CPV, 132,
		 * Officially assigned]
		 */
		"CV": CountryCode{
			name:       "Cape Verde",
			alpha2:     "CV",
			alpha3:     "CPV",
			numeric:    132,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cura%C3%A7ao">Cura&ccedil;ao</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CW">CW</a>, CUW, 531,
		 * Officially assigned]
		 */
		"CW": CountryCode{
			name:       "Cura\u00E7ao",
			alpha2:     "CW",
			alpha3:     "CUW",
			numeric:    531,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Christmas_Island">Christmas Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CX">CX</a>, CXR, 162,
		 * Officially assigned]
		 */
		"CX": CountryCode{
			name:       "Christmas Island",
			alpha2:     "CX",
			alpha3:     "CXR",
			numeric:    162,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cyprus">Cyprus</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CY">CY</a>, CYP, 196,
		 * Officially assigned]
		 */
		"CY": CountryCode{
			name:       "Cyprus",
			alpha2:     "CY",
			alpha3:     "CYP",
			numeric:    196,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Czech_Republic">Czech Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#CZ">CZ</a>, CZE, 203,
		 * Officially assigned]
		 */
		"CZ": CountryCode{
			name:       "Czech Republic",
			alpha2:     "CZ",
			alpha3:     "CZE",
			numeric:    203,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Germany">Germany</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DE">DE</a>, DEU, 276,
		 * Officially assigned]
		 */
		"DE": CountryCode{
			name:       "Germany",
			alpha2:     "DE",
			alpha3:     "DEU",
			numeric:    276,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Diego_Garcia">Diego Garcia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DG">DG</a>, DGA, -1,
		 * Exceptionally reserved]
		 */
		"DG": CountryCode{
			name:       "Diego Garcia",
			alpha2:     "DG",
			alpha3:     "DGA",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Djibouti">Djibouti</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DJ">DJ</a>, DJI, 262,
		 * Officially assigned]
		 */
		"DJ": CountryCode{
			name:       "Djibouti",
			alpha2:     "DJ",
			alpha3:     "DJI",
			numeric:    262,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Denmark">Denmark</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DK">DK</a>, DNK, 208,
		 * Officially assigned]
		 */
		"DK": CountryCode{
			name:       "Denmark",
			alpha2:     "DK",
			alpha3:     "DNK",
			numeric:    208,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Dominica">Dominica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DM">DM</a>, DMA, 212,
		 * Officially assigned]
		 */
		"DM": CountryCode{
			name:       "Dominica",
			alpha2:     "DM",
			alpha3:     "DMA",
			numeric:    212,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Dominican_Republic">Dominican Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DO">DO</a>, DOM, 214,
		 * Officially assigned]
		 */
		"DO": CountryCode{
			name:       "Dominican Republic",
			alpha2:     "DO",
			alpha3:     "DOM",
			numeric:    214,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Algeria">Algeria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#DZ">DZ</a>, DZA, 12,
		 * Officially assigned]
		 */
		"DZ": CountryCode{
			name:       "Algeria",
			alpha2:     "DZ",
			alpha3:     "DZA",
			numeric:    12,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ceuta">Ceuta</a>,
		 * <a href="http://en.wikipedia.org/wiki/Melilla">Melilla</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EA">EA</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"EA": CountryCode{
			name:       "Ceuta, Melilla",
			alpha2:     "EA",
			alpha3:     "",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ecuador">Ecuador</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EC">EC</a>, ECU, 218,
		 * Officially assigned]
		 */
		"EC": CountryCode{
			name:       "Ecuador",
			alpha2:     "EC",
			alpha3:     "ECU",
			numeric:    218,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Estonia">Estonia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EE">EE</a>, EST, 233,
		 * Officially assigned]
		 */
		"EE": CountryCode{
			name:       "Estonia",
			alpha2:     "EE",
			alpha3:     "EST",
			numeric:    233,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Egypt">Egypt</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EG">EG</a>, EGY, 818,
		 * Officially assigned]
		 */
		"EG": CountryCode{
			name:       "Egypt",
			alpha2:     "EG",
			alpha3:     "EGY",
			numeric:    818,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Western_Sahara">Western Sahara</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EH">EH</a>, ESH, 732,
		 * Officially assigned]
		 */
		"EH": CountryCode{
			name:       "Western Sahara",
			alpha2:     "EH",
			alpha3:     "ESH",
			numeric:    732,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Eritrea">Eritrea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ER">ER</a>, ERI, 232,
		 * Officially assigned]
		 */
		"ER": CountryCode{
			name:       "Eritrea",
			alpha2:     "ER",
			alpha3:     "ERI",
			numeric:    232,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Spain">Spain</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ES">ES</a>, ESP, 724,
		 * Officially assigned]
		 */
		"ES": CountryCode{
			name:       "Spain",
			alpha2:     "ES",
			alpha3:     "ESP",
			numeric:    724,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ethiopia">Ethiopia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ET">ET</a>, ETH, 231,
		 * Officially assigned]
		 */
		"ET": CountryCode{
			name:       "Ethiopia",
			alpha2:     "ET",
			alpha3:     "ETH",
			numeric:    231,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/European_Union">European Union</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#EU">EU</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"EU": CountryCode{
			name:       "European Union",
			alpha2:     "EU",
			alpha3:     "",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Finland">Finland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FI">FI</a>, FIN, 246,
		 * Officially assigned]
		 *
		 * @see #SF
		 */
		"FI": CountryCode{
			name:       "Finland",
			alpha2:     "FI",
			alpha3:     "FIN",
			numeric:    246,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Fiji">Fiji</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#">FJ</a>, FJI, 242,
		 * Officially assigned]
		 */
		"FJ": CountryCode{
			name:       "Fiji",
			alpha2:     "FJ",
			alpha3:     "FJI",
			numeric:    242,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Falkland_Islands">Falkland Islands (Malvinas)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FK">FK</a>, FLK, 238,
		 * Officially assigned]
		 */
		"FK": CountryCode{
			name:       "Falkland Islands (Malvinas)",
			alpha2:     "FK",
			alpha3:     "FLK",
			numeric:    238,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Federated_States_of_Micronesia">Micronesia, Federated States of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FM">FM</a>, FSM, 583,
		 * Officially assigned]
		 */
		"FM": CountryCode{
			name:       "Micronesia, Federated States of",
			alpha2:     "FM",
			alpha3:     "FSM",
			numeric:    583,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Faroe_Islands">Faroe Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FO">FO</a>, FRO, 234,
		 * Officially assigned]
		 */
		"FO": CountryCode{
			name:       "Faroe Islands",
			alpha2:     "FO",
			alpha3:     "FRO",
			numeric:    234,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/France">France</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FR">FR</a>, FRA, 250,
		 * Officially assigned]
		 */
		"FR": CountryCode{
			name:       "France",
			alpha2:     "FR",
			alpha3:     "FRA",
			numeric:    250,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Metropolitan_France">France, Metropolitan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#FX">FX</a>, FXX, -1,
		 * Exceptionally reserved]
		 */
		"FX": CountryCode{
			name:       "France, Metropolitan",
			alpha2:     "FX",
			alpha3:     "FXX",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Gabon">Gabon </a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GA">GA</a>, GAB, 266,
		 * Officially assigned]
		 */
		"GA": CountryCode{
			name:       "Gabon",
			alpha2:     "GA",
			alpha3:     "GAB",
			numeric:    266,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_Kingdom">United Kingdom</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GB">GB</a>, GBR, 826,
		 * Officially assigned]
		 */
		"GB": CountryCode{
			name:       "United Kingdom",
			alpha2:     "GB",
			alpha3:     "GBR",
			numeric:    826,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Grenada">Grenada</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GD">GD</a>, GRD, 308,
		 * Officially assigned]
		 */
		"GD": CountryCode{
			name:       "Grenada",
			alpha2:     "GD",
			alpha3:     "GRD",
			numeric:    308,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Georgia_(country)">Georgia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GE">GE</a>, GEO, 268,
		 * Officially assigned]
		 */
		"GE": CountryCode{
			name:       "Georgia",
			alpha2:     "GE",
			alpha3:     "GEO",
			numeric:    268,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/French_Guiana">French Guiana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GF">GF</a>, GUF, 254,
		 * Officially assigned]
		 */
		"GF": CountryCode{
			name:       "French Guiana",
			alpha2:     "GF",
			alpha3:     "GUF",
			numeric:    254,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guernsey">Guernsey</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GG">GG</a>, GGY, 831,
		 * Officially assigned]
		 */
		"GG": CountryCode{
			name:       "Guernsey",
			alpha2:     "GG",
			alpha3:     "GGY",
			numeric:    831,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ghana">Ghana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GH">GH</a>, GHA, 288,
		 * Officially assigned]
		 */
		"GH": CountryCode{
			name:       "Ghana",
			alpha2:     "GH",
			alpha3:     "GHA",
			numeric:    288,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Gibraltar">Gibraltar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GI">GI</a>, GIB, 292,
		 * Officially assigned]
		 */
		"GI": CountryCode{
			name:       "Gibraltar",
			alpha2:     "GI",
			alpha3:     "GIB",
			numeric:    292,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Greenland">Greenland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GL">GL</a>, GRL, 304,
		 * Officially assigned]
		 */
		"GL": CountryCode{
			name:       "Greenland",
			alpha2:     "GL",
			alpha3:     "GRL",
			numeric:    304,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/The_Gambia">Gambia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GM">GM</a>, GMB, 270,
		 * Officially assigned]
		 */
		"GM": CountryCode{
			name:       "Gambia",
			alpha2:     "GM",
			alpha3:     "GMB",
			numeric:    270,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guinea">Guinea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GN">GN</a>, GIN, 324,
		 * Officially assigned]
		 */
		"GN": CountryCode{
			name:       "Guinea",
			alpha2:     "GN",
			alpha3:     "GIN",
			numeric:    324,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guadeloupe">Guadeloupe</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GP">GP</a>, GLP, 312,
		 * Officially assigned]
		 */
		"GP": CountryCode{
			name:       "Guadeloupe",
			alpha2:     "GP",
			alpha3:     "GLP",
			numeric:    312,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Equatorial_Guinea">Equatorial Guinea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GQ">GQ</a>, GNQ, 226,
		 * Officially assigned]
		 */
		"GQ": CountryCode{
			name:       "Equatorial Guinea",
			alpha2:     "GQ",
			alpha3:     "GNQ",
			numeric:    226,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Greece">Greece</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GR">GR</a>, GRC, 300,
		 * Officially assigned]
		 */
		"GR": CountryCode{
			name:       "Greece",
			alpha2:     "GR",
			alpha3:     "GRC",
			numeric:    300,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Georgia_and_the_South_Sandwich_Islands">South Georgia and the South Sandwich Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GS">GS</a>, SGS, 239,
		 * Officially assigned]
		 */
		"GS": CountryCode{
			name:       "South Georgia and the South Sandwich Islands",
			alpha2:     "GS",
			alpha3:     "SGS",
			numeric:    239,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guatemala">Guatemala</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GT">GT</a>, GTM, 320,
		 * Officially assigned]
		 */
		"GT": CountryCode{
			name:       "Guatemala",
			alpha2:     "GT",
			alpha3:     "GTM",
			numeric:    320,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guam">Guam</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GU">GU</a>, GUM, 316,
		 * Officially assigned]
		 */
		"GU": CountryCode{
			name:       "Guam",
			alpha2:     "GU",
			alpha3:     "GUM",
			numeric:    316,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guinea-Bissau">Guinea-Bissau</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GW">GW</a>, GNB, 624,
		 * Officially assigned]
		 */
		"GW": CountryCode{
			name:       "Guinea-Bissau",
			alpha2:     "GW",
			alpha3:     "GNB",
			numeric:    624,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Guyana">Guyana</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#GY">GY</a>, GUY, 328,
		 * Officially assigned]
		 */
		"GY": CountryCode{
			name:       "Guyana",
			alpha2:     "GY",
			alpha3:     "GUY",
			numeric:    328,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Hong_Kong">Hong Kong</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HK">HK</a>, HKG, 344,
		 * Officially assigned]
		 */
		"HK": CountryCode{
			name:       "Hong Kong",
			alpha2:     "HK",
			alpha3:     "HKG",
			numeric:    344,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Heard_Island_and_McDonald_Islands">Heard Island and McDonald Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HM">HM</a>, HMD, 334,
		 * Officially assigned]
		 */
		"HM": CountryCode{
			name:       "Heard Island and McDonald Islands",
			alpha2:     "HM",
			alpha3:     "HMD",
			numeric:    334,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Honduras">Honduras</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HN">HN</a>, HND, 340,
		 * Officially assigned]
		 */
		"HN": CountryCode{
			name:       "Honduras",
			alpha2:     "HN",
			alpha3:     "HND",
			numeric:    340,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Croatia">Croatia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HR">HR</a>, HRV, 191,
		 * Officially assigned]
		 */
		"HR": CountryCode{
			name:       "Croatia",
			alpha2:     "HR",
			alpha3:     "HRV",
			numeric:    191,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Haiti">Haiti</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HT">HT</a>, HTI, 332,
		 * Officially assigned]
		 */
		"HT": CountryCode{
			name:       "Haiti",
			alpha2:     "HT",
			alpha3:     "HTI",
			numeric:    332,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Hungary">Hungary</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#HU">HU</a>, HUN, 348,
		 * Officially assigned]
		 */
		"HU": CountryCode{
			name:       "Hungary",
			alpha2:     "HU",
			alpha3:     "HUN",
			numeric:    348,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Canary_Islands">Canary Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IC">IC</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"IC": CountryCode{
			name:       "Canary Islands",
			alpha2:     "IC",
			alpha3:     "",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Indonesia">Indonesia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ID">ID</a>, IDN, 360,
		 * Officially assigned]
		 */
		"ID": CountryCode{
			name:       "Indonesia",
			alpha2:     "ID",
			alpha3:     "IDN",
			numeric:    360,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Republic_of_Ireland">Ireland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IE">IE</a>, IRL, 372,
		 * Officially assigned]
		 */
		"IE": CountryCode{
			name:       "Ireland",
			alpha2:     "IE",
			alpha3:     "IRL",
			numeric:    372,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Israel">Israel</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IL">IL</a>, ISR, 376,
		 * Officially assigned]
		 */
		"IL": CountryCode{
			name:       "Israel",
			alpha2:     "IL",
			alpha3:     "ISR",
			numeric:    376,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Isle_of_Man">Isle of Man</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IM">IM</a>, IMN, 833,
		 * Officially assigned]
		 */
		"IM": CountryCode{
			name:       "Isle of Man",
			alpha2:     "IM",
			alpha3:     "IMN",
			numeric:    833,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/India">India</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IN">IN</a>, IND, 356,
		 * Officially assigned]
		 */
		"IN": CountryCode{
			name:       "India",
			alpha2:     "IN",
			alpha3:     "IND",
			numeric:    356,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/British_Indian_Ocean_Territory">British Indian Ocean Territory</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IO">IO</a>, IOT, 86,
		 * Officially assigned]
		 */
		"IO": CountryCode{
			name:       "British Indian Ocean Territory",
			alpha2:     "IO",
			alpha3:     "IOT",
			numeric:    86,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Iraq">Iraq</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IQ">IQ</a>, IRQ, 368,
		 * Officially assigned]
		 */
		"IQ": CountryCode{
			name:       "Iraq",
			alpha2:     "IQ",
			alpha3:     "IRQ",
			numeric:    368,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Iran">Iran, Islamic Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IR">IR</a>, IRN, 364,
		 * Officially assigned]
		 */
		"IR": CountryCode{
			name:       "Iran, Islamic Republic of",
			alpha2:     "IR",
			alpha3:     "IRN",
			numeric:    364,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Iceland">Iceland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IS">IS</a>, ISL, 352,
		 * Officially assigned]
		 */
		"IS": CountryCode{
			name:       "Iceland",
			alpha2:     "IS",
			alpha3:     "ISL",
			numeric:    352,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Italy">Italy</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#IT">IT</a>, ITA, 380,
		 * Officially assigned]
		 */
		"IT": CountryCode{
			name:       "Italy",
			alpha2:     "IT",
			alpha3:     "ITA",
			numeric:    380,
			assignment: OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Jersey">Jersey</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JE">JE</a>, JEY, 832,
		 * Officially assigned]
		 */
		"JE": CountryCode{
			name:       "Jersey",
			alpha2:     "JE",
			alpha3:     "JEY",
			numeric:    832,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Jamaica">Jamaica</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JM">JM</a>, JAM, 388,
		 * Officially assigned]
		 */
		"JM": CountryCode{
			name:       "Jamaica",
			alpha2:     "JM",
			alpha3:     "JAM",
			numeric:    388,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Jordan">Jordan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JO">JO</a>, JOR, 400,
		 * Officially assigned]
		 */
		"JO": CountryCode{
			name:       "Jordan",
			alpha2:     "JO",
			alpha3:     "JOR",
			numeric:    400,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Japan">Japan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#JP">JP</a>, JPN, 392,
		 * Officially assigned]
		 */
		"JP": CountryCode{
			name:       "Japan",
			alpha2:     "JP",
			alpha3:     "JPN",
			numeric:    392,
			assignment: OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Kenya">Kenya</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KE">KE</a>, KEN, 404,
		 * Officially assigned]
		 */
		"KE": CountryCode{
			name:       "Kenya",
			alpha2:     "KE",
			alpha3:     "KEN",
			numeric:    404,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kyrgyzstan">Kyrgyzstan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KG">KG</a>, KGZ, 417,
		 * Officially assigned]
		 */
		"KG": CountryCode{
			name:       "Kyrgyzstan",
			alpha2:     "KG",
			alpha3:     "KGZ",
			numeric:    417,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cambodia">Cambodia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KH">KH</a>, KHM, 116,
		 * Officially assigned]
		 */
		"KH": CountryCode{
			name:       "Cambodia",
			alpha2:     "KH",
			alpha3:     "KHM",
			numeric:    116,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kiribati">Kiribati</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KI">KI</a>, KIR, 296,
		 * Officially assigned]
		 */
		"KI": CountryCode{
			name:       "Kiribati",
			alpha2:     "KI",
			alpha3:     "KIR",
			numeric:    296,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Comoros">Comoros</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KM">KM</a>, COM, 174,
		 * Officially assigned]
		 */
		"KM": CountryCode{
			name:       "Comoros",
			alpha2:     "KM",
			alpha3:     "COM",
			numeric:    174,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Kitts_and_Nevis">Saint Kitts and Nevis</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KN">KN</a>, KNA, 659,
		 * Officially assigned]
		 */
		"KN": CountryCode{
			name:       "Saint Kitts and Nevis",
			alpha2:     "KN",
			alpha3:     "KNA",
			numeric:    659,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/North_Korea">Korea, Democratic People's Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KP">KP</a>, PRK, 408,
		 * Officially assigned]
		 */
		"KP": CountryCode{
			name:       "Korea, Democratic People's Republic of",
			alpha2:     "KP",
			alpha3:     "PRK",
			numeric:    408,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Korea">Korea, Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KR">KR</a>, KOR, 410,
		 * Officially assigned]
		 */
		"KR": CountryCode{
			name:       "Korea, Republic of",
			alpha2:     "KR",
			alpha3:     "KOR",
			numeric:    410,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kuwait">Kuwait</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KW">KW</a>, KWT, 414,
		 * Officially assigned]
		 */
		"KW": CountryCode{
			name:       "Kuwait",
			alpha2:     "KW",
			alpha3:     "KWT",
			numeric:    414,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Cayman_Islands">Cayman Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KY">KY</a>, CYM, 136,
		 * Officially assigned]
		 */
		"KY": CountryCode{
			name:       "Cayman Islands",
			alpha2:     "KY",
			alpha3:     "CYM",
			numeric:    136,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kazakhstan">Kazakhstan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#KZ">KZ</a>, KAZ, 398,
		 * Officially assigned]
		 */
		"KZ": CountryCode{
			name:       "Kazakhstan",
			alpha2:     "KZ",
			alpha3:     "KAZ",
			numeric:    398,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Laos">Lao People's Democratic Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LA">LA</a>, LAO, 418,
		 * Officially assigned]
		 */
		"LA": CountryCode{
			name:       "Lao People's Democratic Republic",
			alpha2:     "LA",
			alpha3:     "LAO",
			numeric:    418,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Lebanon">Lebanon</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LB">LB</a>, LBN, 422,
		 * Officially assigned]
		 */
		"LB": CountryCode{
			name:       "Lebanon",
			alpha2:     "LB",
			alpha3:     "LBN",
			numeric:    422,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Lucia">Saint Lucia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LC">LC</a>, LCA, 662,
		 * Officially assigned]
		 */
		"LC": CountryCode{
			name:       "Saint Lucia",
			alpha2:     "LC",
			alpha3:     "LCA",
			numeric:    662,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Liechtenstein">Liechtenstein</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LI">LI</a>, LIE, 438,
		 * Officially assigned]
		 */
		"LI": CountryCode{
			name:       "Liechtenstein",
			alpha2:     "LI",
			alpha3:     "LIE",
			numeric:    438,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sri_Lanka">Sri Lanka</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LK">LK</a>, LKA, 144,
		 * Officially assigned]
		 */
		"LK": CountryCode{
			name:       "Sri Lanka",
			alpha2:     "LK",
			alpha3:     "LKA",
			numeric:    144,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Liberia">Liberia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LR">LR</a>, LBR, 430,
		 * Officially assigned]
		 */
		"LR": CountryCode{
			name:       "Liberia",
			alpha2:     "LR",
			alpha3:     "LBR",
			numeric:    430,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Lesotho">Lesotho</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LS">LS</a>, LSO, 426,
		 * Officially assigned]
		 */
		"LS": CountryCode{
			name:       "Lesotho",
			alpha2:     "LS",
			alpha3:     "LSO",
			numeric:    426,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Lithuania">Lithuania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LT">LT</a>, LTU, 440,
		 * Officially assigned]
		 */
		"LT": CountryCode{
			name:       "Lithuania",
			alpha2:     "LT",
			alpha3:     "LTU",
			numeric:    440,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Luxembourg">Luxembourg</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LU">LU</a>, LUX, 442,
		 * Officially assigned]
		 */
		"LU": CountryCode{
			name:       "Luxembourg",
			alpha2:     "LU",
			alpha3:     "LUX",
			numeric:    442,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Latvia">Latvia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LV">LV</a>, LVA, 428,
		 * Officially assigned]
		 */
		"LV": CountryCode{
			name:       "Latvia",
			alpha2:     "LV",
			alpha3:     "LVA",
			numeric:    428,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Libya">Libya</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#LY">LY</a>, LBY, 434,
		 * Officially assigned]
		 */
		"LY": CountryCode{
			name:       "Libya",
			alpha2:     "LY",
			alpha3:     "LBY",
			numeric:    434,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Morocco">Morocco</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MA">MA</a>, MAR, 504,
		 * Officially assigned]
		 */
		"MA": CountryCode{
			name:       "Morocco",
			alpha2:     "MA",
			alpha3:     "MAR",
			numeric:    504,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Monaco">Monaco</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MC">MC</a>, MCO, 492,
		 * Officially assigned]
		 */
		"MC": CountryCode{
			name:       "Monaco",
			alpha2:     "MC",
			alpha3:     "MCO",
			numeric:    492,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Moldova">Moldova, Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MD">MD</a>, MDA, 498,
		 * Officially assigned]
		 */
		"MD": CountryCode{
			name:       "Moldova, Republic of",
			alpha2:     "MD",
			alpha3:     "MDA",
			numeric:    498,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Montenegro">Montenegro</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ME">ME</a>, MNE, 499,
		 * Officially assigned]
		 */
		"ME": CountryCode{
			name:       "Montenegro",
			alpha2:     "ME",
			alpha3:     "MNE",
			numeric:    499,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Collectivity_of_Saint_Martin">Saint Martin (French part)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MF">MF</a>, MAF, 663,
		 * Officially assigned]
		 */
		"MF": CountryCode{
			name:       "Saint Martin (French part)",
			alpha2:     "MF",
			alpha3:     "MAF",
			numeric:    663,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Madagascar">Madagascar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MG">MG</a>, MDG, 450,
		 * Officially assigned]
		 */
		"MG": CountryCode{
			name:       "Madagascar",
			alpha2:     "MG",
			alpha3:     "MDG",
			numeric:    450,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Marshall_Islands">Marshall Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MH">MH</a>, MHL, 584,
		 * Officially assigned]
		 */
		"MH": CountryCode{
			name:       "Marshall Islands",
			alpha2:     "MH",
			alpha3:     "MHL",
			numeric:    584,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Republic_of_Macedonia">Macedonia, the former Yugoslav Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MK">MK</a>, MKD, 807,
		 * Officially assigned]
		 */
		"MK": CountryCode{
			name:       "Macedonia, the former Yugoslav Republic of",
			alpha2:     "MK",
			alpha3:     "MKD",
			numeric:    807,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mali">Mali</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ML">ML</a>, MLI, 466,
		 * Officially assigned]
		 */
		"ML": CountryCode{
			name:       "Mali",
			alpha2:     "ML",
			alpha3:     "MLI",
			numeric:    466,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Myanmar">Myanmar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MM">MM</a>, MMR, 104,
		 * Officially assigned]
		 *
		 * @see #BU
		 */
		"MM": CountryCode{
			name:       "Myanmar",
			alpha2:     "MM",
			alpha3:     "MMR",
			numeric:    104,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mongolia">Mongolia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MN">MN</a>, MNG, 496,
		 * Officially assigned]
		 */
		"MN": CountryCode{
			name:       "Mongolia",
			alpha2:     "MN",
			alpha3:     "MNG",
			numeric:    496,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Macau">Macao</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MO">MO</a>, MCO, 492,
		 * Officially assigned]
		 */
		"MO": CountryCode{
			name:       "Macao",
			alpha2:     "MO",
			alpha3:     "MAC",
			numeric:    446,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Northern_Mariana_Islands">Northern Mariana Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MP">MP</a>, MNP, 580,
		 * Officially assigned]
		 */
		"MP": CountryCode{
			name:       "Northern Mariana Islands",
			alpha2:     "MP",
			alpha3:     "MNP",
			numeric:    580,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Martinique">Martinique</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MQ">MQ</a>, MTQ, 474,
		 * Officially assigned]
		 */
		"MQ": CountryCode{
			name:       "Martinique",
			alpha2:     "MQ",
			alpha3:     "MTQ",
			numeric:    474,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mauritania">Mauritania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MR">MR</a>, MRT, 478,
		 * Officially assigned]
		 */
		"MR": CountryCode{
			name:       "Mauritania",
			alpha2:     "MR",
			alpha3:     "MRT",
			numeric:    478,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Montserrat">Montserrat</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MS">MS</a>, MSR, 500,
		 * Officially assigned]
		 */
		"MS": CountryCode{
			name:       "Montserrat",
			alpha2:     "MS",
			alpha3:     "MSR",
			numeric:    500,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Malta">Malta</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MT">MT</a>, MLT, 470,
		 * Officially assigned]
		 */
		"MT": CountryCode{
			name:       "Malta",
			alpha2:     "MT",
			alpha3:     "MLT",
			numeric:    470,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mauritius">Mauritius</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MU">MU</a>, MUS, 480,
		 * Officially assigned]]
		 */
		"MU": CountryCode{
			name:       "Mauritius",
			alpha2:     "MU",
			alpha3:     "MUS",
			numeric:    480,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Maldives">Maldives</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MV">MV</a>, MDV, 462,
		 * Officially assigned]
		 */
		"MV": CountryCode{
			name:       "Maldives",
			alpha2:     "MV",
			alpha3:     "MDV",
			numeric:    462,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Malawi">Malawi</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MW">MW</a>, MWI, 454,
		 * Officially assigned]
		 */
		"MW": CountryCode{
			name:       "Malawi",
			alpha2:     "MW",
			alpha3:     "MWI",
			numeric:    454,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mexico">Mexico</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MX">MX</a>, MEX, 484,
		 * Officially assigned]
		 */
		"MX": CountryCode{
			name:       "Mexico",
			alpha2:     "MX",
			alpha3:     "MEX",
			numeric:    484,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Malaysia">Malaysia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MY">MY</a>, MYS, 458,
		 * Officially assigned]
		 */
		"MY": CountryCode{
			name:       "Malaysia",
			alpha2:     "MY",
			alpha3:     "MYS",
			numeric:    458,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mozambique">Mozambique</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#MZ">MZ</a>, MOZ, 508,
		 * Officially assigned]
		 */
		"MZ": CountryCode{
			name:       "Mozambique",
			alpha2:     "MZ",
			alpha3:     "MOZ",
			numeric:    508,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Namibia">Namibia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NA">NA</a>, NAM, 516,
		 * Officially assigned]
		 */
		"NA": CountryCode{
			name:       "Namibia",
			alpha2:     "NA",
			alpha3:     "NAM",
			numeric:    516,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/New_Caledonia">New Caledonia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NC">NC</a>, NCL, 540,
		 * Officially assigned]
		 */
		"NC": CountryCode{
			name:       "New Caledonia",
			alpha2:     "NC",
			alpha3:     "NCL",
			numeric:    540,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Niger">Niger</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NE">NE</a>, NER, 562,
		 * Officially assigned]
		 */
		"NE": CountryCode{
			name:       "Niger",
			alpha2:     "NE",
			alpha3:     "NER",
			numeric:    562,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Norfolk_Island">Norfolk Island</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NF">NF</a>, NFK, 574,
		 * Officially assigned]
		 */
		"NF": CountryCode{
			name:       "Norfolk Island",
			alpha2:     "NF",
			alpha3:     "NFK",
			numeric:    574,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nigeria">Nigeria</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NG">NG</a>, NGA, 566,
		 * Officially assigned]
		 */
		"NG": CountryCode{
			name:       "Nigera",
			alpha2:     "NG",
			alpha3:     "NGA",
			numeric:    566,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nicaragua">Nicaragua</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NI">NI</a>, NIC, 558,
		 * Officially assigned]
		 */
		"NI": CountryCode{
			name:       "Nicaragua",
			alpha2:     "NI",
			alpha3:     "NIC",
			numeric:    558,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Netherlands">Netherlands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NL">NL</a>, NLD, 528,
		 * Officially assigned]
		 */
		"NL": CountryCode{
			name:       "Netherlands",
			alpha2:     "NL",
			alpha3:     "NLD",
			numeric:    528,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Norway">Norway</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NO">NO</a>, NOR, 578,
		 * Officially assigned]
		 */
		"NO": CountryCode{
			name:       "Norway",
			alpha2:     "NO",
			alpha3:     "NOR",
			numeric:    578,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nepal">Nepal</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NP">NP</a>, NPL, 524,
		 * Officially assigned]
		 */
		"NP": CountryCode{
			name:       "Nepal",
			alpha2:     "NP",
			alpha3:     "NPL",
			numeric:    524,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Nauru">Nauru</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NR">NR</a>, NRU, 520,
		 * Officially assigned]
		 */
		"NR": CountryCode{
			name:       "Nauru",
			alpha2:     "NR",
			alpha3:     "NRU",
			numeric:    520,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saudi%E2%80%93Iraqi_neutral_zone">Neutral Zone</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NT">NT</a>, NTHH, 536,
		 * Traditionally reserved]
		 */
		"NT": CountryCode{
			name:       "Neutral Zone",
			alpha2:     "NT",
			alpha3:     "NTHH",
			numeric:    536,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Niue">Niue</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NU">NU</a>, NIU, 570,
		 * Officially assigned]
		 */
		"NU": CountryCode{
			name:       "Niue",
			alpha2:     "NU",
			alpha3:     "NIU",
			numeric:    570,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/New_Zealand">New Zealand</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#NZ">NZ</a>, NZL, 554,
		 * Officially assigned]
		 */
		"NZ": CountryCode{
			name:       "New Zealand",
			alpha2:     "NZ",
			alpha3:     "NZL",
			numeric:    554,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href=http://en.wikipedia.org/wiki/Oman"">Oman</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#OM">OM</a>, OMN, 512,
		 * Officially assigned]
		 */
		"OM": CountryCode{
			name:       "Oman",
			alpha2:     "OM",
			alpha3:     "OMN",
			numeric:    512,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Panama">Panama</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PA">PA</a>, PAN, 591,
		 * Officially assigned]
		 */
		"PA": CountryCode{
			name:       "Panama",
			alpha2:     "PA",
			alpha3:     "PAN",
			numeric:    591,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Peru">Peru</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PE">PE</a>, PER, 604,
		 * Officially assigned]
		 */
		"PE": CountryCode{
			name:       "Peru",
			alpha2:     "PE",
			alpha3:     "PER",
			numeric:    604,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/French_Polynesia">French Polynesia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PF">PF</a>, PYF, 258,
		 * Officially assigned]
		 */
		"PF": CountryCode{
			name:       "French Polynesia",
			alpha2:     "PF",
			alpha3:     "PYF",
			numeric:    258,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Papua_New_Guinea">Papua New Guinea</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PG">PG</a>, PNG, 598,
		 * Officially assigned]
		 */
		"PG": CountryCode{
			name:       "Papua New Guinea",
			alpha2:     "PG",
			alpha3:     "PNG",
			numeric:    598,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Philippines">Philippines</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PH">PH</a>, PHL, 608,
		 * Officially assigned]
		 */
		"PH": CountryCode{
			name:       "Philippines",
			alpha2:     "PH",
			alpha3:     "PHL",
			numeric:    608,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Pakistan">Pakistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PK">PK</a>, PAK, 586,
		 * Officially assigned]
		 */
		"PK": CountryCode{
			name:       "Pakistan",
			alpha2:     "PK",
			alpha3:     "PAK",
			numeric:    586,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Poland">Poland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PL">PL</a>, POL, 616,
		 * Officially assigned]
		 */
		"PL": CountryCode{
			name:       "Poland",
			alpha2:     "PL",
			alpha3:     "POL",
			numeric:    616,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Pierre_and_Miquelon">Saint Pierre and Miquelon</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PM">PM</a>, SPM, 666,
		 * Officially assigned]
		 */
		"PM": CountryCode{
			name:       "Saint Pierre and Miquelon",
			alpha2:     "PM",
			alpha3:     "SPM",
			numeric:    666,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Pitcairn_Islands">Pitcairn</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PN">PN</a>, PCN, 612,
		 * Officially assigned]
		 */
		"PN": CountryCode{
			name:       "Pitcairn",
			alpha2:     "PN",
			alpha3:     "PCN",
			numeric:    612,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Puerto_Rico">Puerto Rico</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PR">PR</a>, PRI, 630,
		 * Officially assigned]
		 */
		"PR": CountryCode{
			name:       "Puerto Rico",
			alpha2:     "PR",
			alpha3:     "PRI",
			numeric:    630,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Palestinian_territories">Palestine, State of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PS">PS</a>, PSE, 275,
		 * Officially assigned]
		 */
		"PS": CountryCode{
			name:       "Palestine, State of",
			alpha2:     "PS",
			alpha3:     "PSE",
			numeric:    275,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Portugal">Portugal</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PT">PT</a>, PRT, 620,
		 * Officially assigned]
		 */
		"PT": CountryCode{
			name:       "Portugal",
			alpha2:     "PT",
			alpha3:     "PRT",
			numeric:    620,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Palau">Palau</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PW">PW</a>, PLW, 585,
		 * Officially assigned]
		 */
		"PW": CountryCode{
			name:       "Palau",
			alpha2:     "PW",
			alpha3:     "PLW",
			numeric:    585,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Paraguay">Paraguay</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#PY">PY</a>, PRY, 600,
		 * Officially assigned]
		 */
		"PY": CountryCode{
			name:       "Paraguay",
			alpha2:     "PY",
			alpha3:     "PRY",
			numeric:    600,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Qatar">Qatar</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#QA">QA</a>, QAT, 634,
		 * Officially assigned]
		 */
		"QA": CountryCode{
			name:       "Qatar",
			alpha2:     "QA",
			alpha3:     "QAT",
			numeric:    634,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/R%C3%A9union">R&eacute;union</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RE">RE</a>, REU, 638,
		 * Officially assigned]
		 */
		"RE": CountryCode{
			name:       "R\u00E9union",
			alpha2:     "RE",
			alpha3:     "REU",
			numeric:    638,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Romania">Romania</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RO">RO</a>, ROU, 642,
		 * Officially assigned]
		 */
		"RO": CountryCode{
			name:       "Romania",
			alpha2:     "RO",
			alpha3:     "ROU",
			numeric:    642,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Serbia">Serbia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RS">RS</a>, SRB, 688,
		 * Officially assigned]
		 */
		"RS": CountryCode{
			name:       "Serbia",
			alpha2:     "RS",
			alpha3:     "SRB",
			numeric:    688,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Russia">Russian Federation</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RU">RU</a>, RUS, 643,
		 * Officially assigned]
		 */
		"RU": CountryCode{
			name:       "Russian Federation",
			alpha2:     "RU",
			alpha3:     "RUS",
			numeric:    643,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Rwanda">Rwanda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#RW">RW</a>, RWA, 646,
		 * Officially assigned]
		 */
		"RW": CountryCode{
			name:       "Rwanda",
			alpha2:     "RW",
			alpha3:     "RWA",
			numeric:    646,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saudi_Arabia">Saudi Arabia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SA">SA</a>, SAU, 682,
		 * Officially assigned]
		 */
		"SA": CountryCode{
			name:       "Saudi Arabia",
			alpha2:     "SA",
			alpha3:     "SAU",
			numeric:    682,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Solomon_Islands">Solomon Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SB">SB</a>, SLB, 90,
		 * Officially assigned]
		 */
		"SB": CountryCode{
			name:       "Solomon Islands",
			alpha2:     "SB",
			alpha3:     "SLB",
			numeric:    90,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Seychelles">Seychelles</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SC">SC</a>, SYC, 690,
		 * Officially assigned]
		 */
		"SC": CountryCode{
			name:       "Seychelles",
			alpha2:     "SC",
			alpha3:     "SYC",
			numeric:    690,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sudan">Sudan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SD">SD</a>, SDN, 729,
		 * Officially assigned]
		 */
		"SD": CountryCode{
			name:       "Sudan",
			alpha2:     "SD",
			alpha3:     "SDN",
			numeric:    729,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sweden">Sweden</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SE">SE</a>, SWE, 752,
		 * Officially assigned]
		 */
		"SE": CountryCode{
			name:       "Sweden",
			alpha2:     "SE",
			alpha3:     "SWE",
			numeric:    752,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Finland">Finland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SF">SF</a>, FIN, 246,
		 * Traditionally reserved]
		 *
		 * @see #FI
		 */
		"SF": CountryCode{
			name:       "Finland",
			alpha2:     "SF",
			alpha3:     "FIN",
			numeric:    246,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Singapore">Singapore</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SG">SG</a>, SGP, 702,
		 * Officially assigned]
		 */
		"SG": CountryCode{
			name:       "Singapore",
			alpha2:     "SG",
			alpha3:     "SGP",
			numeric:    702,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Helena,_Ascension_and_Tristan_da_Cunha">Saint Helena, Ascension and Tristan da Cunha</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SH">SH</a>, SHN, 654,
		 * Officially assigned]
		 */
		"SH": CountryCode{
			name:       "Saint Helena, Ascension and Tristan da Cunha",
			alpha2:     "SH",
			alpha3:     "SHN",
			numeric:    654,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Slovenia">Slovenia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SI">SI</a>, SVN, 705,
		 * Officially assigned]
		 */
		"SI": CountryCode{
			name:       "Slovenia",
			alpha2:     "SI",
			alpha3:     "SVN",
			numeric:    705,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Svalbard_and_Jan_Mayen">Svalbard and Jan Mayen</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SJ">SJ</a>, SJM, 744,
		 * Officially assigned]
		 */
		"SJ": CountryCode{
			name:       "Svalbard and Jan Mayen",
			alpha2:     "SJ",
			alpha3:     "SJM",
			numeric:    744,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Slovakia">Slovakia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SK">SK</a>, SVK, 703,
		 * Officially assigned]
		 */
		"SK": CountryCode{
			name:       "Slovakia",
			alpha2:     "SK",
			alpha3:     "SVK",
			numeric:    703,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sierra_Leone">Sierra Leone</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SL">SL</a>, SLE, 694,
		 * Officially assigned]
		 */
		"SL": CountryCode{
			name:       "Sierra Leone",
			alpha2:     "SL",
			alpha3:     "SLE",
			numeric:    694,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/San_Marino">San Marino</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SM">SM</a>, SMR, 674,
		 * Officially assigned]
		 */
		"SM": CountryCode{
			name:       "San Marino",
			alpha2:     "SM",
			alpha3:     "SMR",
			numeric:    674,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Senegal">Senegal</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SN">SN</a>, SEN, 686,
		 * Officially assigned]
		 */
		"SN": CountryCode{
			name:       "Senegal",
			alpha2:     "SN",
			alpha3:     "SEN",
			numeric:    686,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Somalia">Somalia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SO">SO</a>, SOM, 706,
		 * Officially assigned]
		 */
		"SO": CountryCode{
			name:       "Somalia",
			alpha2:     "SO",
			alpha3:     "SOM",
			numeric:    706,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Suriname">Suriname</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SR">SR</a>, SUR, 740,
		 * Officially assigned]
		 */
		"SR": CountryCode{
			name:       "Suriname",
			alpha2:     "SR",
			alpha3:     "SUR",
			numeric:    740,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Sudan">South Sudan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SS">SS</a>, SSD, 728,
		 * Officially assigned]
		 */
		"SS": CountryCode{
			name:       "South Sudan",
			alpha2:     "SS",
			alpha3:     "SSD",
			numeric:    728,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/S%C3%A3o_Tom%C3%A9_and_Pr%C3%ADncipe">Sao Tome and Principe</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ST">ST</a>, STP, 678,
		 * Officially assigned]
		 */
		"ST": CountryCode{
			name:       "Sao Tome and Principe",
			alpha2:     "ST",
			alpha3:     "STP",
			numeric:    678,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Soviet_Union">USSR</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SU">SU</a>, SUN, -1,
		 * Exceptionally reserved]
		 */
		"SU": CountryCode{
			name:       "USSR",
			alpha2:     "SU",
			alpha3:     "SUN",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/El_Salvador">El Salvador</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SV">SV</a>, SLV, 222,
		 * Officially assigned]
		 */
		"SV": CountryCode{
			name:       "El Salvador",
			alpha2:     "SV",
			alpha3:     "SLV",
			numeric:    222,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Sint_Maarten">Sint Maarten (Dutch part)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SX">SX</a>, SXM, 534,
		 * Officially assigned]
		 */
		"SX": CountryCode{
			name:       "Sint Maarten (Dutch part)",
			alpha2:     "SX",
			alpha3:     "SXM",
			numeric:    534,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Syria">Syrian Arab Republic</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SY">SY</a>, SYR, 760,
		 * Officially assigned]
		 */
		"SY": CountryCode{
			name:       "Syrian Arab Republic",
			alpha2:     "SY",
			alpha3:     "SYR",
			numeric:    760,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Swaziland">Swaziland</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#SZ">SZ</a>, SWZ, 748,
		 * Officially assigned]
		 */
		"SZ": CountryCode{
			name:       "Swaziland",
			alpha2:     "SZ",
			alpha3:     "SWZ",
			numeric:    748,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tristan_da_Cunha">Tristan da Cunha</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TA">TA</a>, TAA, -1,
		 * Exceptionally reserved.
		 */
		"TA": CountryCode{
			name:       "Tristan da Cunha",
			alpha2:     "TA",
			alpha3:     "TAA",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Turks_and_Caicos_Islands">Turks and Caicos Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TC">TC</a>, TCA, 796,
		 * Officially assigned]
		 */
		"TC": CountryCode{
			name:       "Turks and Caicos Islands",
			alpha2:     "TC",
			alpha3:     "TCA",
			numeric:    796,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Chad">Chad</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TD">TD</a>, TCD, 148,
		 * Officially assigned]
		 */
		"TD": CountryCode{
			name:       "Chad",
			alpha2:     "TD",
			alpha3:     "TCD",
			numeric:    148,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/French_Southern_and_Antarctic_Lands">French Southern Territories</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TF">TF</a>, ATF, 260,
		 * Officially assigned]
		 */
		"TF": CountryCode{
			name:       "French Southern Territories",
			alpha2:     "TF",
			alpha3:     "ATF",
			numeric:    260,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Togo">Togo</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TG">TG</a>, TGO, 768,
		 * Officially assigned]
		 */
		"TG": CountryCode{
			name:       "Togo",
			alpha2:     "TG",
			alpha3:     "TGO",
			numeric:    768,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Thailand">Thailand</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TH">TH</a>, THA, 764,
		 * Officially assigned]
		 */
		"TH": CountryCode{
			name:       "Thailand",
			alpha2:     "TH",
			alpha3:     "THA",
			numeric:    764,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tajikistan">Tajikistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TJ">TJ</a>, TJK, 762,
		 * Officially assigned]
		 */
		"TJ": CountryCode{
			name:       "Tajikistan",
			alpha2:     "TJ",
			alpha3:     "TJK",
			numeric:    762,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tokelau">Tokelau</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TK">TK</a>, TKL, 772,
		 * Officially assigned]
		 */
		"TK": CountryCode{
			name:       "Tokelau",
			alpha2:     "TK",
			alpha3:     "TKL",
			numeric:    772,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/East_Timor">Timor-Leste</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TL">TL</a>, TLS, 626,
		 * Officially assigned]
		 */
		"TL": CountryCode{
			name:       "Timor-Leste",
			alpha2:     "TL",
			alpha3:     "TLS",
			numeric:    626,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Turkmenistan">Turkmenistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TM">TM</a>, TKM, 795,
		 * Officially assigned]
		 */
		"TM": CountryCode{
			name:       "Turkmenistan",
			alpha2:     "TM",
			alpha3:     "TKM",
			numeric:    795,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tunisia">Tunisia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TN">TN</a>, TUN, 788,
		 * Officially assigned]
		 */
		"TN": CountryCode{
			name:       "Tunisia",
			alpha2:     "TN",
			alpha3:     "TUN",
			numeric:    788,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tonga">Tonga</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TO">TO</a>, TON, 776,
		 * Officially assigned]
		 */
		"TO": CountryCode{
			name:       "Tonga",
			alpha2:     "TO",
			alpha3:     "TON",
			numeric:    776,
			assignment: OFFICIALLY_ASSIGNED,
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
			name:       "East Timor",
			alpha2:     "TP",
			alpha3:     "TPTL",
			numeric:    0,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Turkey">Turkey</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TR">TR</a>, TUR, 792,
		 * Officially assigned]
		 */
		"TR": CountryCode{
			name:       "Turkey",
			alpha2:     "TR",
			alpha3:     "TUR",
			numeric:    792,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Trinidad_and_Tobago">Trinidad and Tobago</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TT">TT</a>, TTO, 780,
		 * Officially assigned]
		 */
		"TT": CountryCode{
			name:       "Trinidad and Tobago",
			alpha2:     "TT",
			alpha3:     "TTO",
			numeric:    780,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tuvalu">Tuvalu</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TV">TV</a>, TUV, 798,
		 * Officially assigned]
		 */
		"TV": CountryCode{
			name:       "Tuvalu",
			alpha2:     "TV",
			alpha3:     "TUV",
			numeric:    798,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Taiwan">Taiwan, Province of China</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TW">TW</a>, TWN, 158,
		 * Officially assigned]
		 */
		"TW": CountryCode{
			name:       "Taiwan, Province of China",
			alpha2:     "TW",
			alpha3:     "TWN",
			numeric:    158,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Tanzania">Tanzania, United Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#TZ">TZ</a>, TZA, 834,
		 * Officially assigned]
		 */
		"TZ": CountryCode{
			name:       "Tanzania, United Republic of",
			alpha2:     "TZ",
			alpha3:     "TZA",
			numeric:    834,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Ukraine">Ukraine</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UA">UA</a>, UKR, 804,
		 * Officially assigned]
		 */
		"UA": CountryCode{
			name:       "Ukraine",
			alpha2:     "UA",
			alpha3:     "UKR",
			numeric:    804,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Uganda">Uganda</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UG">UG</a>, UGA, 800,
		 * Officially assigned]
		 */
		"UG": CountryCode{
			name:       "Uganda",
			alpha2:     "UG",
			alpha3:     "UGA",
			numeric:    800,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_Kingdom">United Kingdom</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UK">UK</a>, null, -1,
		 * Exceptionally reserved]
		 */
		"UK": CountryCode{
			name:       "United Kingdom",
			alpha2:     "UK",
			alpha3:     "",
			numeric:    -1,
			assignment: EXCEPTIONALLY_RESERVED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/United_States_Minor_Outlying_Islands">United States Minor Outlying Islands</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UM">UM</a>, UMI, 581,
		 * Officially assigned]
		 */
		"UM": CountryCode{
			name:       "United States Minor Outlying Islands",
			alpha2:     "UM",
			alpha3:     "UMI",
			numeric:    581,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_States">United States</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#US">US</a>, USA, 840,
		 * Officially assigned]
		 */
		"US": CountryCode{
			name:       "United States",
			alpha2:     "US",
			alpha3:     "USA",
			numeric:    840,
			assignment: OFFICIALLY_ASSIGNED,
		},
		/**
		 * <a href="http://en.wikipedia.org/wiki/Uruguay">Uruguay</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UY">UY</a>, URY, 858,
		 * Officially assigned]
		 */
		"UY": CountryCode{
			name:       "Uruguay",
			alpha2:     "UY",
			alpha3:     "URY",
			numeric:    858,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Uzbekistan">Uzbekistan</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#UZ">UZ</a>, UZB, 860,
		 * Officially assigned]
		 */
		"UZ": CountryCode{
			name:       "Uzbekistan",
			alpha2:     "UZ",
			alpha3:     "UZB",
			numeric:    860,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Vatican_City">Holy See (Vatican City State)</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VA">VA</a>, VAT, 336,
		 * Officially assigned]
		 */
		"VA": CountryCode{
			name:       "Holy See (Vatican City State)",
			alpha2:     "VA",
			alpha3:     "VAT",
			numeric:    336,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Saint_Vincent_and_the_Grenadines">Saint Vincent and the Grenadines</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VC">VC</a>, VCT, 670,
		 * Officially assigned]
		 */
		"VC": CountryCode{
			name:       "Saint Vincent and the Grenadines",
			alpha2:     "VC",
			alpha3:     "VCT",
			numeric:    670,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Venezuela">Venezuela, Bolivarian Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VE">VE</a>, VEN, 862,
		 * Officially assigned]
		 */
		"VE": CountryCode{
			name:       "Venezuela, Bolivarian Republic of",
			alpha2:     "VE",
			alpha3:     "VEN",
			numeric:    862,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/British_Virgin_Islands">Virgin Islands, British</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VG">VG</a>, VGB, 92,
		 * Officially assigned]
		 */
		"VG": CountryCode{
			name:       "Virgin Islands, British",
			alpha2:     "VG",
			alpha3:     "VGB",
			numeric:    92,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/United_States_Virgin_Islands">Virgin Islands, U.S.</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VI">VI</a>, VIR, 850,
		 * Officially assigned]
		 */
		"VI": CountryCode{
			name:       "Virgin Islands, U.S.",
			alpha2:     "VI",
			alpha3:     "VIR",
			numeric:    850,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Vietnam">Viet Nam</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VN">VN</a>, VNM, 704,
		 * Officially assigned]
		 */
		"VN": CountryCode{
			name:       "Viet Nam",
			alpha2:     "VN",
			alpha3:     "VNM",
			numeric:    704,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Vanuatu">Vanuatu</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#VU">VU</a>, VUT, 548,
		 * Officially assigned]
		 */
		"VU": CountryCode{
			name:       "Vanuatu",
			alpha2:     "VU",
			alpha3:     "VUT",
			numeric:    548,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Wallis_and_Futuna">Wallis and Futuna</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#WF">WF</a>, WLF, 876,
		 * Officially assigned]
		 */
		"WF": CountryCode{
			name:       "Wallis and Futuna",
			alpha2:     "WF",
			alpha3:     "WLF",
			numeric:    876,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Samoa">Samoa</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#WS">WS</a>, WSM, 882,
		 * Officially assigned]
		 */
		"WS": CountryCode{
			name:       "Samoa",
			alpha2:     "WS",
			alpha3:     "WSM",
			numeric:    882,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Kosovo">Kosovo, Republic of</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#XK">XK</a>, XXK, -1,
		 * User assigned]
		 */
		"XK": CountryCode{
			name:       "Kosovo, Republic of",
			alpha2:     "XK",
			alpha3:     "XXK",
			numeric:    -1,
			assignment: USER_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Yemen">Yemen</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#YE">YE</a>, YEM, 887,
		 * Officially assigned]
		 */
		"YE": CountryCode{
			name:       "Yemen",
			alpha2:     "YE",
			alpha3:     "YEM",
			numeric:    887,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Mayotte">Mayotte</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#YT">YT</a>, MYT, 175,
		 * Officially assigned]
		 */
		"YT": CountryCode{
			name:       "Mayotte",
			alpha2:     "YT",
			alpha3:     "MYT",
			numeric:    175,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Yugoslavia">Yugoslavia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#YU">YU</a>, YUCS, 890,
		 * Traditionally reserved]
		 */
		"YU": CountryCode{
			name:       "Yugoslavia",
			alpha2:     "YU",
			alpha3:     "YUCS",
			numeric:    890,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/South_Africa">South Africa</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ZA">ZA</a>, ZAF, 710,
		 * Officially assigned]
		 */
		"ZA": CountryCode{
			name:       "South Africa",
			alpha2:     "ZA",
			alpha3:     "ZAF",
			numeric:    710,
			assignment: OFFICIALLY_ASSIGNED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Zambia">Zambia</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ZM">ZM</a>, ZMB, 894,
		 * Officially assigned]
		 */
		"ZM": CountryCode{
			name:       "Zambia",
			alpha2:     "ZM",
			alpha3:     "ZMB",
			numeric:    894,
			assignment: OFFICIALLY_ASSIGNED,
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
			name:       "Zaire",
			alpha2:     "ZR",
			alpha3:     "ZRCD",
			numeric:    0,
			assignment: TRANSITIONALLY_RESERVED,
		},

		/**
		 * <a href="http://en.wikipedia.org/wiki/Zimbabwe">Zimbabwe</a>
		 * [<a href="http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#ZW">ZW</a>, ZWE, 716,
		 * Officially assigned]
		 */
		"ZW": CountryCode{
			name:       "Zimbabwe",
			alpha2:     "ZW",
			alpha3:     "ZWE",
			numeric:    716,
			assignment: OFFICIALLY_ASSIGNED,
		},
	}

	for _, cc := range by_alpha2 {
		if cc.alpha3 != "" {
			by_alpha3[cc.alpha3] = cc
		}
		by_name[cc.name] = cc
		by_numeric[cc.numeric] = cc
	}
}
