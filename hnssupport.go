package hcsshim

import (
	"github.com/sirupsen/logrus"
)

type HNSSupportedFeatures struct {
	Acl HNSAclFeatures `json:"ACL"`
}

type HNSAclFeatures struct {
	AclAddressLists       bool `json:"AclAddressLists"`
	AclNoHostRulePriority bool `json:"AclHostRulePriority"`
	AclPortRanges         bool `json:"AclPortRanges"`
	AclRuleId             bool `json:"AclRuleId"`
}

func GetHNSSupportedFeatures() HNSSupportedFeatures {
	var hnsFeatures HNSSupportedFeatures

	globals, err := GetHNSGlobals()
	if err != nil {
		// Expected on pre-17060 builds, all features will be false/unsupported
		logrus.Debugf("Unable to obtain HNS globals: %s", err)
		return hnsFeatures
	}

	hnsFeatures.Acl = HNSAclFeatures{
		AclAddressLists:       isHNSFeatureSupported(globals.Version, HNSVersion1804),
		AclNoHostRulePriority: isHNSFeatureSupported(globals.Version, HNSVersion1804),
		AclPortRanges:         isHNSFeatureSupported(globals.Version, HNSVersion1804),
		AclRuleId:             isHNSFeatureSupported(globals.Version, HNSVersion1804),
	}

	return hnsFeatures
}

func isHNSFeatureSupported(currentVersion HNSVersion, minVersionSupported HNSVersion) bool {
	if currentVersion.Major < minVersionSupported.Major {
		return false
	}
	if currentVersion.Major > minVersionSupported.Major {
		return true
	}
	if currentVersion.Minor < minVersionSupported.Minor {
		return false
	}
	return true
}
