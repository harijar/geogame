package prompts

// PromptData is created by GetPrompts to avoid repetitions in set of prompts returned to the user.
type promptData struct {
	Country string
	// This field is here to get fast access to the current country name without querying database

	Population   bool
	Area         bool
	GDP          bool
	GDPPerCapita bool
	HDI          bool
	Agricultural bool
	Industrial   bool
	Service      bool
	// GetNumeric checks one of these fields depending on its field parameter.
	// They are set true when GetNumeric returns a prompt for the first time for a corresponding field.
	// If they are true, GetNumeric returns empty string. This is the easiest way to avoid obvious repetitions. (maybe I'm gonna change it later, I don't know, for now it's like this)

	Location map[string]string
	// This field is checked by GetHemisphere.
	// Key is the direction received by GetLocation (e.g. if current country is norther than some other country, key is "Northern").
	// Value is hemisphere where that other country is located.
	// If key == value (or if value == "Equator"/"Greenwich"), information about hemisphere is not necessary, and GetHemisphere returns an empty string.

	HemisphereLatt string
	HemisphereLong string
	// GetLocation checks one of these fields depending on the extreme point column selected from database.
	// if other country is located in different hemisphere than current country, its relative location is obvious,
	// so GetLocation returns an empty string.

	Landlocked bool
	// GetBool checks this field if Island column is selected from database.
	// if this field == true, the fact that country is on continent is obvious, so GetBool returns an empty string.

	Island bool
	// GetBool checks this field if Landlocked column is selected from database.
	// if this field == true, the fact that country has access to sea is obvious, so GetBool returns an empty string.
}
