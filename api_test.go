package peeringdb

import "testing"

func TestFormatSearchParameters(t *testing.T) {
	var searchMap map[string]interface{}
	var expected string
	var searchParameters string

	// Test for nil map
	expected = ""
	searchParameters = formatSearchParameters(nil)
	if searchParameters != expected {
		t.Errorf("formatSearchParameters, want '%s' got '%s'", expected,
			searchParameters)
	}

	// Test for empty map
	searchMap = make(map[string]interface{})
	expected = ""
	searchParameters = formatSearchParameters(searchMap)
	if searchParameters != expected {
		t.Errorf("formatSearchParameters, want '%s' got '%s'", expected,
			searchParameters)
	}

	// Test one value
	searchMap = make(map[string]interface{})
	searchMap["id"] = 10
	expected = "&id=10"
	searchParameters = formatSearchParameters(searchMap)
	if searchParameters != expected {
		t.Errorf("formatSearchParameters, want '%s' got '%s'", expected,
			searchParameters)
	}

	// Test two values
	searchMap = make(map[string]interface{})
	searchMap["id"] = 10
	searchMap["asn"] = 65536
	expected = "&asn=65536&id=10"
	searchParameters = formatSearchParameters(searchMap)
	if searchParameters != expected {
		t.Errorf("formatSearchParameters, want '%s' got '%s'", expected,
			searchParameters)
	}
}

func TestFormatURL(t *testing.T) {
	var expected string
	var url string

	searchMap := make(map[string]interface{})
	searchMap["id"] = 10

	// Test fac namespace with search parameter
	expected = "https://peeringdb.com/api/fac?depth=1&id=10"
	url = formatURL(facilityNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test ix namespace with search parameter
	expected = "https://peeringdb.com/api/ix?depth=1&id=10"
	url = formatURL(internetExchangeNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test ixfac namespace with search parameter
	expected = "https://peeringdb.com/api/ixfac?depth=1&id=10"
	url = formatURL(internetExchangeFacilityNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test ixlan namespace with search parameter
	expected = "https://peeringdb.com/api/ixlan?depth=1&id=10"
	url = formatURL(internetExchangeLANNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test ixpfx namespace with search parameter
	expected = "https://peeringdb.com/api/ixpfx?depth=1&id=10"
	url = formatURL(internetExchangePrefixNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test net namespace with search parameter
	expected = "https://peeringdb.com/api/net?depth=1&id=10"
	url = formatURL(networkNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test netfac namespace with search parameter
	expected = "https://peeringdb.com/api/netfac?depth=1&id=10"
	url = formatURL(networkFacilityNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test netixlan namespace with search parameter
	expected = "https://peeringdb.com/api/netixlan?depth=1&id=10"
	url = formatURL(networkInternetExchangeLANNamepsace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test org namespace with search parameter
	expected = "https://peeringdb.com/api/org?depth=1&id=10"
	url = formatURL(organizationNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}

	// Test poc namespace with search parameter
	expected = "https://peeringdb.com/api/poc?depth=1&id=10"
	url = formatURL(networkContactNamespace, searchMap)
	if url != expected {
		t.Errorf("formatURL, want '%s' got '%s'", expected, url)
	}
}

func TestGetASN(t *testing.T) {
	expectedASN := 29467
	net := GetASN(expectedASN)

	if net.ASN != expectedASN {
		t.Errorf("GetASN, want ASN '%d' got '%d'", expectedASN, net.ASN)
	}
}