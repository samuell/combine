Combinator generator
====================

Install
-------

```bash
go get github.com/samuell/comb/...
```

Usage
-----

```bash
$ ./combinator a,b x,y,z 1,2,3,4
a, a, a, a, a, a, a, a, a, a, a, a, b, b, b, b, b, b, b, b, b, b, b, b
x, x, x, x, y, y, y, y, z, z, z, z, x, x, x, x, y, y, y, y, z, z, z, z
1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4
```
