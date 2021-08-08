# HelloObjectPoolPattern

> [Source](https://golangbyexample.com/golang-object-pool/)

## Explain on functions in `pool.go`

- `func initPool(poolObjects []iPoolObject) (*pool, error) `

Init pool with `poolObjects` and if `len(poolObjects) = 0` it returns error.

- `func (p *pool) loan() (iPoolObject, error)`

Pick a `iPoolObject` from `p.idle` after that update `p.idle` and `p.active` , if `len(p.idle) = 0` it returns error.
It's thread safe.

- `func (p *pool) receive(target iPoolObject) error`

It receives `iPoolObject` and do check if it belongs `p.active` if yes remove it from `p.active`, after check add it
to `p.idle` , call it `recycle` better? It's thread safe.

- `func (p *pool) remove(target iPoolObject) error`

Check if `target` belongs to `p.active` if not returns error