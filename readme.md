

# playing around with bson in golang

modified utf8 defined here
https://docs.oracle.com/javase/8/docs/api/java/io/DataInput.html#modified-utf-8

main site: http://bsonspec.org/

using this: https://golang.org/pkg/encoding/binary/
this: https://www.jonathan-petitcolas.com/2014/09/25/parsing-binary-files-in-go.html

and this: https://golang.org/pkg/bytes/

Wondering about performance really.
Can you rewrite values in bson without a huge penalty?
it (looks) like it could be possible
I think you would need meta information about a bson object to do it quickly though.
And size of the resulting blob would change. I would think an append only file would handle this quite well?