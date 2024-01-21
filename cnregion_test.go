package cnregion

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CnRegionTestSuite struct {
	suite.Suite
	instance *CnRegion
}

func TestCnRegionTestSuite(t *testing.T) {
	suite.Run(t, &CnRegionTestSuite{
		instance: NewCnRegion(),
	})
}

func (suite *CnRegionTestSuite) TestParseByCode_ValidCode() {
	province, city, area, street, err := suite.instance.ParseByCode("450305004")
	suite.NoError(err)
	suite.Equal("广西壮族自治区", province)
	suite.Equal("桂林市", city)
	suite.Equal("七星区", area)
	suite.Equal("漓东街道", street)
}

func (suite *CnRegionTestSuite) TestParseByCode_InvalidCode() {
	_, _, _, _, err := suite.instance.ParseByCode("invalidCode")
	suite.Error(err)
}

func (suite *CnRegionTestSuite) TestParseByName_ValidName() {
	code, err := suite.instance.ParseByName("广西壮族自治区", "桂林市", "七星区", "漓东街道")
	suite.NoError(err)
	suite.Equal("450305004", code)
}

func (suite *CnRegionTestSuite) TestParseByName_InvalidName() {
	_, err := suite.instance.ParseByName("invalidProvince", "invalidCity", "invalidArea", "invalidStreet")
	suite.Error(err)
}

func (suite *CnRegionTestSuite) TestSearch_ValidKeyword() {
	result := suite.instance.Search("广西壮族自治区")
	suite.NotEmpty(result)
	suite.Equal(14, len(result[0].Children))
	result = suite.instance.Search("广西壮族自治区桂林市七星区")
	suite.NotEmpty(result)
	suite.Equal(6, len(result[0].Children))
}

func (suite *CnRegionTestSuite) TestSearch_InvalidKeyword() {
	result := suite.instance.Search("invalidKeyword")
	suite.Empty(result)
}
