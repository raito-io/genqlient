// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InputEnumQueryResponse is returned by InputEnumQuery on success.
type InputEnumQueryResponse struct {
	// usersWithRole looks a user up by role.
	UsersWithRole []InputEnumQueryUsersWithRoleUser `json:"usersWithRole"`
}

// GetUsersWithRole returns InputEnumQueryResponse.UsersWithRole, and is useful for accessing the field via an interface.
func (v *InputEnumQueryResponse) GetUsersWithRole() []InputEnumQueryUsersWithRoleUser {
	return v.UsersWithRole
}

// InputEnumQueryUsersWithRoleUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type InputEnumQueryUsersWithRoleUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns InputEnumQueryUsersWithRoleUser.Id, and is useful for accessing the field via an interface.
func (v *InputEnumQueryUsersWithRoleUser) GetId() testutil.ID { return v.Id }

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// __InputEnumQueryInput is used internally by genqlient
type __InputEnumQueryInput struct {
	Role Role `json:"role"`
}

// GetRole returns __InputEnumQueryInput.Role, and is useful for accessing the field via an interface.
func (v *__InputEnumQueryInput) GetRole() Role { return v.Role }

// The query or mutation executed by InputEnumQuery.
const InputEnumQueryOperation = `
query InputEnumQuery ($role: Role!) {
	usersWithRole(role: $role) {
		id
	}
}
`

func InputEnumQuery(
	client graphql.Client,
	role Role,
) (*InputEnumQueryResponse, error) {
	req := &graphql.Request{
		OpName: "InputEnumQuery",
		Query:  InputEnumQueryOperation,
		Variables: &__InputEnumQueryInput{
			Role: role,
		},
	}
	var err error

	var data InputEnumQueryResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

