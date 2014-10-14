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

var alpha2 map[string]CountryCode

func init() {
	alpha2 = map[string]CountryCode{
		"AC": CountryCode{
			name:       "Ascension Island",
			alpha2:     "AC",
			alpha3:     "ASC",
			numeric:    -1,
			assignment: OFFICIALLY_ASSIGNED,
		},
	}
}
