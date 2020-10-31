package planning

import "net/http"

func (suite *RestSuite) TestCreateReadDeleteCategory() {
	// ensure that category "abc" doesn't exist
	suite.
		Get("/category/get/abc").
		Expect(http.StatusNotFound, M{
			"success": false,
			"message": "category not found",
		})
	// create category
	suite.
		Post("/category/create", M{
			"name": "abc",
		}).
		Expect(http.StatusCreated, M{
			"success": true,
			"element": M{
				"name": "abc",
			},
		})
	// ensure that category is fetchable
	suite.
		Get("/category/get/abc").
		Expect(http.StatusOK, M{
			"success": true,
			"element": M{
				"name": "abc",
			},
		})
	// delete category "abc"
	suite.
		Get("/category/delete/abc").
		Expect(http.StatusOK, M{
			"success": true,
		})
	// ensure that category "abc" no longer exists
	suite.
		Get("/category/get/abc").
		Expect(http.StatusNotFound, M{
			"success": false,
			"message": "category not found",
		})
}
