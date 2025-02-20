// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package meta

import (
	"errors"
	"testing"

	"github.com/luci/gae/filter/featureBreaker"
	"github.com/luci/gae/impl/memory"
	"github.com/luci/gae/service/rawdatastore"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
)

func TestGetEntityGroupVersion(t *testing.T) {
	t.Parallel()

	Convey("GetEntityGroupVersion", t, func() {
		c := memory.Use(context.Background())
		c, fb := featureBreaker.FilterRDS(c, errors.New("INTERNAL_ERROR"))
		rds := rawdatastore.Get(c)

		aKey, err := rds.Put(rds.NewKey("A", "", 0, nil), rawdatastore.PropertyMap{
			"Val": {rawdatastore.MkProperty(10)},
		})
		So(err, ShouldBeNil)

		v, err := GetEntityGroupVersion(c, aKey)
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 1)

		So(rds.Delete(aKey), ShouldBeNil)

		v, err = GetEntityGroupVersion(c, rds.NewKey("madeUp", "thing", 0, aKey))
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 2)

		v, err = GetEntityGroupVersion(c, rds.NewKey("madeUp", "thing", 0, nil))
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 0)

		fb.BreakFeatures(nil, "Get")

		v, err = GetEntityGroupVersion(c, aKey)
		So(err.Error(), ShouldContainSubstring, "INTERNAL_ERROR")
	})
}
