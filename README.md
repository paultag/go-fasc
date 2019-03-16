go-fasc
=======

FASC, or the Federal Agency Smartcard Number is a legacy specification that
stores physical security information that's encoded into smartcards. Internally,
this is stored as a 5-bit-byte packed into 8 bit bytes. This library unpacks
the bytestream unto an idiomatic struct, to be further handled by the user.

This spec is nearly completely undocumented, and woefully hard to understand.
In so far as it's used within the federal government, some serious work
needs to be done in order to properly store and process data. This library
requires an in-depth understanding of FASC's quirks and internal representation.
