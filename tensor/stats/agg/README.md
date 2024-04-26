# agg

Docs: [GoDoc](https://pkg.go.dev/cogentcore.org/core/tensor/stats/agg)

This package provides aggregation functions operating on `IndexView` indexed views of `table.Table` data, along with standard AggFunc functions that can be used at any level of aggregation from etensor on up.

The main functions use names to specify columns, with `*Index` version available operating on column indexes.

See the [tsragg](../tsragg) package for functions that operate directly on a `tensor.Tensor` without the index view indirection.

