## Prerequisite

- install cobra

```

https://github.com/spf13/cobra

```

- Change the path to the todo file in the todo.go

```
const (
	todoFile = "~/.todo"
)

```

## Run

### Add

```

go run main.go add -t "Finish the iOS build"

```

### List

```

go run main.go list

```

### Complete

```


 go run main.go complete -c 1


```
