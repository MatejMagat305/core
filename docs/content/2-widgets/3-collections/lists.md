# Lists

Cogent Core provides interactive lists that allow you to display a slice value to users and have them edit it. Users can use context menus and drag-and-drop to add, remove, and move rows.

You can make a list from any slice pointer:

```Go
core.NewList(parent).SetSlice(&[]int{1, 3, 5})
```

You can make a list that fits in one line:

```Go
core.NewInlineList(parent).SetSlice(&[]int{1, 3, 5})
```

You can detect when the user changes the value of the slice:

```Go
sl := []int{1, 3, 5}
core.NewList(parent).SetSlice(&sl).OnChange(func(e events.Event) {
    core.MessageSnackbar(parent, fmt.Sprintf("Slice: %v", sl))
})
```

You can prevent the user from editing the slice:

```Go
core.NewList(parent).SetSlice(&[]int{1, 3, 5}).SetReadOnly(true)
```

When you use [[core.NewValue]] with a slice value, it will create an inline list if the slice has four or fewer elements:

```Go
// core.NewValue(parent, &[]int{1, 3, 5})
```

Otherwise, it will create a button that opens a dialog with a normal list:

```Go
// core.NewValue(parent, &[]int{1, 3, 5, 7, 9})
```
