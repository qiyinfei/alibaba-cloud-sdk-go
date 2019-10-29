package emr

// TagResources is a nested struct in emr response
type TagResources struct {
	TagResource []TagResource `json:"TagResource" xml:"TagResource"`
}
