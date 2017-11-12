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
// Model: Singular (Grant)
// DB Table: Plural (grants)
// Resource: Plural (Grants)
// Path: Plural (/grants)
// View Template Folder: Plural (/templates/grants/)

// GrantsResource is the resource for the Grant model
type GrantsResource struct {
	buffalo.Resource
}

// List gets all Grants. This function is mapped to the path
// GET /grants
func (v GrantsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	grants := &models.Grants{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Grants from the DB
	if err := q.All(grants); err != nil {
		return errors.WithStack(err)
	}

	// Make Grants available inside the html template
	c.Set("grants", grants)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("grants/index.html"))
}

// Show gets the data for one Grant. This function is mapped to
// the path GET /grants/{grant_id}
func (v GrantsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Grant
	grant := &models.Grant{}

	// To find the Grant the parameter grant_id is used.
	if err := tx.Find(grant, c.Param("grant_id")); err != nil {
		return c.Error(404, err)
	}

	// Make grant available inside the html template
	c.Set("grant", grant)

	return c.Render(200, r.HTML("grants/show.html"))
}

// New renders the form for creating a new Grant.
// This function is mapped to the path GET /grants/new
func (v GrantsResource) New(c buffalo.Context) error {
	// Make grant available inside the html template
	c.Set("grant", &models.Grant{})

	return c.Render(200, r.HTML("grants/new.html"))
}

// Create adds a Grant to the DB. This function is mapped to the
// path POST /grants
func (v GrantsResource) Create(c buffalo.Context) error {
	// Allocate an empty Grant
	grant := &models.Grant{}

	// Bind grant to the html form elements
	if err := c.Bind(grant); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(grant)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make grant available inside the html template
		c.Set("grant", grant)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("grants/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Grant was created successfully")

	// and redirect to the grants index page
	return c.Redirect(302, "/grants/%s", grant.ID)
}

// Edit renders a edit form for a Grant. This function is
// mapped to the path GET /grants/{grant_id}/edit
func (v GrantsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Grant
	grant := &models.Grant{}

	if err := tx.Find(grant, c.Param("grant_id")); err != nil {
		return c.Error(404, err)
	}

	// Make grant available inside the html template
	c.Set("grant", grant)
	return c.Render(200, r.HTML("grants/edit.html"))
}

// Update changes a Grant in the DB. This function is mapped to
// the path PUT /grants/{grant_id}
func (v GrantsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Grant
	grant := &models.Grant{}

	if err := tx.Find(grant, c.Param("grant_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Grant to the html form elements
	if err := c.Bind(grant); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(grant)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make grant available inside the html template
		c.Set("grant", grant)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("grants/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Grant was updated successfully")

	// and redirect to the grants index page
	return c.Redirect(302, "/grants/%s", grant.ID)
}

// Destroy deletes a Grant from the DB. This function is mapped
// to the path DELETE /grants/{grant_id}
func (v GrantsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Grant
	grant := &models.Grant{}

	// To find the Grant the parameter grant_id is used.
	if err := tx.Find(grant, c.Param("grant_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(grant); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Grant was destroyed successfully")

	// Redirect to the grants index page
	return c.Redirect(302, "/grants")
}