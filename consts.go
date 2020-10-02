// {{{ Copyright (c) Paul R. Tagliamonte <paultag@gmail.com>, 2019
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

package fasc // import "pault.ag/go/fasc"

// OrganizationalCategory is the cateogry of organization referenced, such
// as the Federal Government, State Government, or a Company.
type OrganizationalCategory int

// AssociationCategory is the type of relationship between the individual and
// the Organization. This is something like "employee", or "contractor".
type AssociationCategory int

const (
	// OrganizationalCategoryFederalGoverment indicates the organization
	// is a part of the federal government.
	OrganizationalCategoryFederalGoverment OrganizationalCategory = 1

	// OrganizationalCategoryStateGovernment indicates the organization
	// is a part of a state government.
	OrganizationalCategoryStateGovernment OrganizationalCategory = 2

	// OrganizationalCategoryCommercialEnterprise indicates the organization
	// is a company, or other commercial enterprise.
	OrganizationalCategoryCommercialEnterprise OrganizationalCategory = 3

	// OrganizationalCategoryForeignGovernment indicates the organization
	// is a foreign government.
	OrganizationalCategoryForeignGovernment OrganizationalCategory = 4

	// AssociationCategoryEmployee indicates the individual is an employee
	// of the organization in question.
	AssociationCategoryEmployee AssociationCategory = 1

	// AssociationCategoryCivil indidicates the individual is a civilian
	// employee of a branch of the military.
	AssociationCategoryCivil AssociationCategory = 2

	// AssociationCategoryExecutiveStaff indidicates the individual is a
	// member of the executive staff.
	AssociationCategoryExecutiveStaff AssociationCategory = 3

	// AssociationCategoryUniformedServivce indidicates the individual is a
	// member of the uniformed service of a branch of the military.
	AssociationCategoryUniformedServivce AssociationCategory = 4

	// AssociationCategoryContractor indidicates the individual is acting
	// as a contractor for the organization.
	AssociationCategoryContractor AssociationCategory = 5

	// AssociationCategoryAffiliate indidicates the individual was issued
	// a credential on the basis of being affiliated with the organization.
	AssociationCategoryAffiliate AssociationCategory = 6

	// AssociationCategoryBeneficiary indidicates the individual is a
	// beneficiary of the organization.
	AssociationCategoryBeneficiary AssociationCategory = 7
)

// String will return a human readable string version of the OrganizationalCategory.
func (oc OrganizationalCategory) String() string {
	switch oc {
	case OrganizationalCategoryFederalGoverment:
		return "federal government"
	case OrganizationalCategoryStateGovernment:
		return "state government"
	case OrganizationalCategoryCommercialEnterprise:
		return "commercial enterprise"
	case OrganizationalCategoryForeignGovernment:
		return "foreign government"
	}
	return "unknown"
}

// String will return a human readable string version of the AssociationCategory.
func (ac AssociationCategory) String() string {
	switch ac {
	case AssociationCategoryEmployee:
		return "employee"
	case AssociationCategoryCivil:
		return "civil"
	case AssociationCategoryExecutiveStaff:
		return "execstaff"
	case AssociationCategoryUniformedServivce:
		return "uniformed service"
	case AssociationCategoryContractor:
		return "contractor"
	case AssociationCategoryAffiliate:
		return "affiliate"
	case AssociationCategoryBeneficiary:
		return "beneficiary"
	}
	return "unknown"
}

// vim: foldmethod=marker
