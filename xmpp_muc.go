// Copyright 2013 Flo Lauber <dev@qatfy.at>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(flo):
//   - support password protected MUC rooms
//   - cleanup signatures of join/leave functions
package xmpp

import (
	"fmt"
)

const (
	nsMUC     = "http://jabber.org/protocol/muc"
	nsMUCUser = "http://jabber.org/protocol/muc#user"
)

// xep-0045 7.2
func (c *Client) JoinMUC(jid string) {
	writeMessageOut(c.tls, fmt.Sprintf("<presence to='%s'>"+
		"<x xmlns='%s'></x>i"+
		"</presence>",
		xmlEscape(jid), nsMUC))
}

func (c *Client) JoinMUCWithFrom(jid string, username string) {
    writeMessageOut(c.tls, fmt.Sprintf("<presence to='%s'\n"+
        "from='%s'>\n"+
		"<x xmlns='%s' />\n"+
		"</presence>",
		xmlEscape(jid), username, nsMUC))
}

// xep-0045 7.14
func (c *Client) LeaveMUC(jid string) {
	writeMessageOut(c.tls, fmt.Sprintf("<presence from='%s' to='%s' type='unavailable' />",
		c.jid, xmlEscape(jid)))
}
