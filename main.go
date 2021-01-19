package main

import "log"

/*
MKCOL, LOCK, GET, PUT, UNLOCK

 A write lock MUST prevent a principal without the lock from
   successfully executing a PUT, POST, PROPPATCH, LOCK, UNLOCK, MOVE,
   DELETE, or MKCOL on the locked resource.  All other current methods,
	 GET in particular, function independently of the lock.

	 PROPFIND

	 A client may submit a Depth header with a value of "0", "1", or
   "infinity" with a PROPFIND on a collection resource with internal
   member URIs.  DAV compliant servers MUST support the "0", "1" and
   "infinity" behaviors. By default, the PROPFIND method without a Depth
	 header MUST act as if a "Depth: infinity" header was included.


	 resource,
	 collection (of resources)
*/

func main() {
	log.Printf("check")
}
