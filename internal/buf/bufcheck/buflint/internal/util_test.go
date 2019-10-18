package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageHasVersionSuffix(t *testing.T) {
	// note we allow this in the linter as we check this in PACKAGE_DEFINED
	// however, for the purposes of packageHasVersionSuffix, this does not
	testPackageHasVersionSuffix(t, false, "")
	testPackageHasVersionSuffix(t, false, "foo")
	testPackageHasVersionSuffix(t, false, "foo.bar")
	testPackageHasVersionSuffix(t, false, "foo.v1.bar")
	testPackageHasVersionSuffix(t, false, "v1")
	testPackageHasVersionSuffix(t, false, "v1beta1")
	testPackageHasVersionSuffix(t, true, "foo.v1")
	testPackageHasVersionSuffix(t, true, "foo.v2")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1alpha1")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1alpha2")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1beta1")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1beta2")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1p1alpha1")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1p1alpha2")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1p1beta1")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1p1beta2")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1test")
	testPackageHasVersionSuffix(t, true, "foo.bar.v1testfoo")
	testPackageHasVersionSuffix(t, false, "foo.v0")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0alpha1")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0alpha2")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0beta1")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0beta2")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0p1alpha1")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0p1alpha2")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0p1beta1")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0p1beta2")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0test")
	testPackageHasVersionSuffix(t, false, "foo.bar.v0testfoo")
	testPackageHasVersionSuffix(t, false, "foo.bar.v1alpha0")
	testPackageHasVersionSuffix(t, false, "foo.bar.v1beta0")
	testPackageHasVersionSuffix(t, false, "foo.bar.v1p1alpha0")
	testPackageHasVersionSuffix(t, false, "foo.bar.v1p1beta0")
	testPackageHasVersionSuffix(t, false, "foo.bar.v1p0alpha1")
	testPackageHasVersionSuffix(t, false, "foo.bar.v1p0beta1")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1alpha1")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1alpha2")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1beta1")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1beta2")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1p1alpha1")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1p1alpha2")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1p1beta1")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1p1beta2")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1test")
	testPackageHasVersionSuffix(t, false, "foo.bar.vv1testfoo")
	testPackageHasVersionSuffix(t, false, "foo.bar.v1aalpha1")
}

func testPackageHasVersionSuffix(t *testing.T, expected bool, pkg string) {
	assert.Equal(t, expected, packageHasVersionSuffix(pkg), pkg)
}
