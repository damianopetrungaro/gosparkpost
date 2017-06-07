package main

import (
	"fmt"
	"strings"
	"time"
)

var ArfFormat string = `From: <%s>
Date: Mon, 02 Jan 2006 15:04:05 MST
Subject: FW: Earn money
To: <%s>
MIME-Version: 1.0
Content-Type: multipart/report; report-type=feedback-report;
      boundary="%s"

--%s
Content-Type: text/plain; charset="US-ASCII"
Content-Transfer-Encoding: 7bit

This is an email abuse report for an email message
received from IP 10.67.41.167 on Thu, 8 Mar 2005
14:00:00 EDT.
For more information about this format please see
http://www.mipassoc.org/arf/.

--%s
Content-Type: message/feedback-report

Feedback-Type: abuse
User-Agent: SomeGenerator/1.0
Version: 0.1

--%s
Content-Type: message/rfc822
Content-Disposition: inline

From: <bounces-%d@%s>
Received: from mailserver.example.net (mailserver.example.net
        [10.67.41.167])
        by example.com with ESMTP id M63d4137594e46;
        Thu, 08 Mar 2005 14:00:00 -0400
To: <Undisclosed Recipients>
Subject: Earn money
MIME-Version: 1.0
Content-type: text/plain
Message-ID: 8787KJKJ3K4J3K4J3K4J3.mail@%s
X-MSFBL: %s
Date: Thu, 02 Sep 2004 12:31:03 -0500

Spam Spam Spam
Spam Spam Spam
Spam Spam Spam
Spam Spam Spam

--%s--
`

func BuildArf(from, to, msfbl string, cid int) string {
	boundary := fmt.Sprintf("_----%d===_61/00-25439-267B0055", time.Now().Unix())
	domain := from[strings.Index(from, "@")+1:]
	msg := fmt.Sprintf(ArfFormat, from, to,
		boundary, boundary, boundary, boundary,
		cid, domain, domain, msfbl, boundary)
	return msg
}
