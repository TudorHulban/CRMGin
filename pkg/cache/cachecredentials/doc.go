/*
Package cachecredentials provides a KV cache where the key is the user code.
The value consists of:
a. user ID
b. user password (hashed)
c. user roles

Time to live for the cache items could be 24 hours.
*/
package cachecredentials
