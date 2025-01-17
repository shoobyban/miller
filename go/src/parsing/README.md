This directory contains a single source file, `mlr.bnf`, which is the lexical/semantic grammar file for the Miller `put`/`filter` DSL using the GOCC framework. (In a classical Lex/Yacc framework, there would be separate `mlr.l` and `mlr.y` files; using GOCC, there is a single `mlr.bnf` file.)

All subdirectories of `src/parsing/` are autogen code created by GOCC's processing of `mlr.bnf`. They are nonetheless committed to source control, since running GOCC takes quite a bit longer than the `go build` does, and the BNF file doesn't often change. See the top-level `miller/go` build scripts for how to rerun GOCC. As of this writing, it's `bin/gocc -o src/parsing src/parsing/mlr.bnf` as invoked from the `miller/go` base directory.

Making changes to `mlr.bnf` requires several minutes to re-run GOCC. For experimental changes, please see the [experiments](../../../experiments/dsl-parser) directory.
