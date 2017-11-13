package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	"github.com/postcert/entitlements/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (IdentityGroup)
// DB Table: Plural (identity_groups)
// Resource: Plural (IdentityGroups)
// Path: Plural (/identity_groups)
// View Template Folder: Plural (/templates/identity_groups/)

// IdentityGroupsResource is the resource for the IdentityGroup model
type IdentityGroupsResource struct {
	buffalo.Resource
}

// List gets all IdentityGroups. This function is mapped to the path
// GET /identity_groups
func (v IdentityGroupsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	identityGroups := &models.IdentityGroups{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all IdentityGroups from the DB
	if err := q.All(identityGroups); err != nil {
		return errors.WithStack(err)
	}

	// Make IdentityGroups available inside the html template
	c.Set("identityGroups", identityGroups)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("identity_groups/index.html"))
}

// Show gets the data for one IdentityGroup. This function is mapped to
// the path GET /identity_groups/{identity_group_id}
func (v IdentityGroupsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty IdentityGroup
	identityGroup := &models.IdentityGroup{}

	// To find the IdentityGroup the parameter identity_group_id is used.
	if err := tx.Find(identityGroup, c.Param("identity_group_id")); err != nil {
		return c.Error(404, err)
	}

	// Make identityGroup available inside the html template
	c.Set("identityGroup", identityGroup)

	return c.Render(200, r.HTML("identity_groups/show.html"))
}

// New renders the form for creating a new IdentityGroup.
// This function is mapped to the path GET /identity_groups/new
func (v IdentityGroupsResource) New(c buffalo.Context) error {
	// Make identityGroup available inside the html template
	c.Set("identityGroup", &models.IdentityGroup{})

	return c.Render(200, r.HTML("identity_groups/new.html"))
}

// Create adds a IdentityGroup to the DB. This function is mapped to the
// path POST /identity_groups
func (v IdentityGroupsResource) Create(c buffalo.Context) error {
	// Allocate an empty IdentityGroup
	identityGroup := &models.IdentityGroup{}

	// Bind identityGroup to the html form elements
	if err := c.Bind(identityGroup); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(identityGroup)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make identityGroup available inside the html template
		c.Set("identity_group", identityGroup)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("identity_groups/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "IdentityGroup was created successfully")

	// and redirect to the identity_groups index page
	return c.Redirect(302, "/identity_groups/%s", identityGroup.ID)
}

// Edit renders a edit form for a IdentityGroup. This function is
// mapped to the path GET /identity_groups/{identity_group_id}/edit
func (v IdentityGroupsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty IdentityGroup
	identityGroup := &models.IdentityGroup{}

	if err := tx.Find(identityGroup, c.Param("identity_group_id")); err != nil {
		return c.Error(404, err)
	}

	// Make identityGroup available inside the html template
	c.Set("identityGroup", identityGroup)
	return c.Render(200, r.HTML("identity_groups/edit.html"))
}

// Update changes a IdentityGroup in the DB. This function is mapped to
// the path PUT /identity_groups/{identity_group_id}
func (v IdentityGroupsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty IdentityGroup
	identityGroup := &models.IdentityGroup{}

	if err := tx.Find(identityGroup, c.Param("identity_group_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind IdentityGroup to the html form elements
	if err := c.Bind(identityGroup); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(identityGroup)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make identityGroup available inside the html template
		c.Set("identity_group", identityGroup)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("identity_groups/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "IdentityGroup was updated successfully")

	// and redirect to the identity_groups index page
	return c.Redirect(302, "/identity_groups/%s", identityGroup.ID)
}

// Destroy deletes a IdentityGroup from the DB. This function is mapped
// to the path DELETE /identity_groups/{identity_group_id}
func (v IdentityGroupsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty IdentityGroup
	identityGroup := &models.IdentityGroup{}

	// To find the IdentityGroup the parameter identity_group_id is used.
	if err := tx.Find(identityGroup, c.Param("identity_group_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(identityGroup); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "IdentityGroup was destroyed successfully")

	// Redirect to the identity_groups index page
	return c.Redirect(302, "/identity_groups")
}
