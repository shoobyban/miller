<!---  PLEASE DO NOT EDIT DIRECTLY. EDIT THE .md.in FILE PLEASE. --->
# Performance

## Disclaimer

In a previous version of this page, I compared Miller to some items in the Unix toolkit in terms of run time. But such comparisons are very much not apples-to-apples:

* Miller's principal strength is that it handles **key-value data in various formats** while the system tools **do not**. So if you time `mlr sort` on a CSV file against system `sort`, it's not relevant to say which is faster by how many percent -- Miller will respect the header line, leaving it in place, while the system sort will move it, sorting it along with all the other header lines. This would be comparing the run times of two programs produce different outputs.  Likewise, `awk` doesn't respect header lines, although you can code up some CSV-handling using `if (NR==1) { ... } else { ... }`. And that's just CSV: I don't know any simple way to get `sort`, `awk`, etc. to handle DKVP, JSON, etc. -- which is the main reason I wrote Miller.

* **Implementations differ by platform**: one `awk` may be fundamentally faster than another, and `mawk` has a very efficient bytecode implementation -- which handles positionally indexed data far faster than Miller does.

* The system `sort` command will, on some systems, handle too-large-for-RAM datasets by spilling to disk; Miller (as of version 5.2.0, mid-2017) does not. Miller sorts are always stable; GNU supports stable and unstable variants.

* Etc.

## Summary

Miller can do many kinds of processing on key-value-pair data using elapsed time roughly of the same order of magnitude as items in the Unix toolkit can handle positionally indexed data. Specific results vary widely by platform, implementation details, and multi-core use (or not). Lastly, specific special-purpose non-record-aware processing will run far faster if implemented in `grep`, `sed`, etc.
