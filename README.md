# PB

Printable returns a human readable version of a byte array. The bytes that correspond with 
ASCII printable characters [32-127) are passed through. Other bytes are replaced with 
\x followed by a two character zero-padded hex code for byte.
