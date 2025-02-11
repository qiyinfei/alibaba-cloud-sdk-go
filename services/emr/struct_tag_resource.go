package emr

// TagResource is a nested struct in emr response
type TagResource struct {
	TagKey       string `json:"TagKey" xml:"TagKey"`
	TagValue     string `json:"TagValue" xml:"TagValue"`
	ResourceType string `json:"ResourceType" xml:"ResourceType"`
	ResourceId   string `json:"ResourceId" xml:"ResourceId"`
}
