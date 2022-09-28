;; This WASM module was compiled with wat2wasm, but subsequently modified with
;; a HEX editor to change `global.get 0` and `global.set 0` to `global.get 1`
;; and `global.set 1`. Because global 1 is immutable, wat2wasm would have not
;; allowed compiling with `global.set 1` (but such a case was intended).
(module
  (type (;0;) (func (param i64)))
  (type (;1;) (func))
  (import "env" "int64finish" (func (;0;) (type 0)))
  (func (;1;) (type 1)
    global.get 0
    i64.const 1
    i64.add
    global.set 0

    global.get 1
    call 0
  )
  (memory (;0;) 2)
  (global (;0;) (mut i64) (i64.const 66560))
  (global (;1;) i64 (i64.const 42))
  (export "memory" (memory 0))
  (export "getglobal" (func 1))
)