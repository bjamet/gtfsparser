// Copyright 2015 geOps
// Authors: patrick.brosi@geops.de
//
// Use of this source code is governed by a GPL v2
// license that can be found in the LICENSE file

package gtfs

import (
	url "net/url"
)

type FeedInfo struct {
	Publisher_name string
	Publisher_url  *url.URL
	Lang           string
	Start_date     Date
	End_date       Date
	Phone          string
	Version        string
}
