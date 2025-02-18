// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// ListOfListsOfListsResponse is returned by ListOfListsOfLists on success.
type ListOfListsOfListsResponse struct {
	ListOfListsOfLists [][][]string `json:"listOfListsOfLists"`
}

// GetListOfListsOfLists returns ListOfListsOfListsResponse.ListOfListsOfLists, and is useful for accessing the field via an interface.
func (v *ListOfListsOfListsResponse) GetListOfListsOfLists() [][][]string {
	return v.ListOfListsOfLists
}

func ListOfListsOfLists(
	client graphql.Client,
) (*ListOfListsOfListsResponse, error) {
	var err error

	var retval ListOfListsOfListsResponse
	err = client.MakeRequest(
		nil,
		"ListOfListsOfLists",
		`
query ListOfListsOfLists {
	listOfListsOfLists
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

