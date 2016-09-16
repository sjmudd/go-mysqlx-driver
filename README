This is a project I started to see what would be required to create
a database/sql driver for the MySQL X protocol which was announced
as part of the MySQL 5.7.12 server release.

Documentation of the MySQL X protocol can be found here:
https://dev.mysql.com/doc/internals/en/x-protocol.html.

The code in this project is based very heavily on the original MySQL
driver which was available on https://github.com/go-sql-driver/mysql
as of the time that the MySQL X protocol was announced.

Copyright (C) 2016 Simon J Mudd <sjmudd@pobox.com>

KNOWN ISSUES
============

This file lists the current known issues with the driver where I
have bothered to find them. They can be split into the following
types:

[A] Issues with the driver itself
---------------------------------

A1. Use of TLS mode is not possible yet if the appropriate seting
is provided with the DSN.

A2. No way to determine mysqlx_max_allowed_packet short of doing a
full query after authorisation. This will add several round trips
to the connection to get this information and is needed to set
internal buffers but to also to prevent sending "too large packets"
to the server.

A3. No way to provide parameters to SQL queries

A4. Support of understanding all datatypes is incomplete

A5. Documentation of the mapping of the MySQL datatypes to the value
provided back to the caller is not done. This is actually quite
important to ensure that behavour is clearly defined.

A6. No character set support.  The X protocol is also rather vague
about expectations for client and server initial setups which does
not help. However some character data that's received from the MySQL
server may require conversion to Go's utf8 etc and that is not done
as the Go driver does not expose the server's internal datatypes.

[B] Issues with using the X protocol using github.com/golang/go-sql-driver because the
driver does not support the functionality needed by X protocol or expose them to the
client.
---------------------------------

B1. No way to disconnect a session and re-authenticate, allowed by
X protocol but not by go driver. Close() is likely to drop the
connection to ensure that previous state does not remain in the
"previous connection".

B2. No way to determine capabilities prior to logging in. Perhaps
not critical but not currently possible. X protocol allows this
which may be convenient.

B3. No way to use the XAPI type JSON/document store interface. This
probably makes sense as database/sql assumes you're going to talk
SQL.

B4. No way to pipeline commands. You must wait for a single statement
to complete prior to sending the next one. The pool functionality
of the go driver would allow you to use another connection transparently
but that's not the same.

B5. No way to change TLS setting once connected. The attribute is
a connection setting and while it may not be common to want to
change behaviour it might be convenient.

B6. No support for anything like init_connect which can be used
frequently. I think really this should be negoatiated as part of
the X protocol but that functionality is not currently present.
