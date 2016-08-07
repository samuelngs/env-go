# env-go
Environment variables utility for golang

```
Set(key string, vals ...string) error
Raw(name string, defs ...string) string
Multi(name, separator string, defs ...string) []string
Int(name string, defs ...int) int
I32(name string, defs ...int32) int32
I64(name string, defs ...int64) int64
Bool(name string, defs ...bool) bool
```
