# Reflect

## String variable
```go
	name := "mostain"
	rv := reflect.ValueOf(name)
	fmt.Println("string:", rv.String())
	fmt.Println("kind:", rv.Kind())
	fmt.Println("type:", rv.Type())
	fmt.Println("length:", rv.Len())
```

## Array variable
```go
	stds := []string{"Istiyak", "Hasan", "Dipu"}
	fmt.Println("kind:", reflect.ValueOf(stds).Kind()) //slice
	fmt.Println("type:", reflect.ValueOf(stds).Type())
	fmt.Println("length:", reflect.ValueOf(stds).Len())
	len := reflect.ValueOf(stds).Len()
	for i := 0; i < len; i++ {
		fmt.Println(i, "->", reflect.ValueOf(stds).Index(i))
	}
  ```

## Func
```go
	greet := func(s, t, u string) string { return fmt.Sprintf("%s %s %s", s, t, u) }
	frv := reflect.ValueOf(greet)
	fmt.Println("string:", frv.String())
	fmt.Println("kind:", frv.Kind())
	fmt.Println("type:", frv.Type())
	fmt.Println("parameter count:", frv.Type().NumIn())
```
	
## Struct
```go
	func Details(inf interface{}) []string {

		var fields []string
		typof := reflect.TypeOf(inf)
		count := typof.NumField()
		for i := 0; i < count; i++ {
			third := typof.Field(i).Name
			fields = append(fields, third)
		}
		return fields
	}

	//call
	var p Person
	fmt.Println(Details(p)) //[Name Age Height Gender]
	
```

## Map
```go
	row := make(map[string]interface{})
	row["name"] = "Arisha"
	row["age"] = 6
	row["height"] = 4.2

	trow := reflect.TypeOf(row)
	fmt.Println("string:", trow.String())
	fmt.Println("kind:", trow.Kind())
	fmt.Println("keyType:", trow.Key())    //key type
	fmt.Println("valueType:", trow.Elem()) //value type
	//melm := trow.Elem()

	fmt.Println("--------------------")

	vrow := reflect.ValueOf(row)
	fmt.Println("string:", vrow.String())
	fmt.Println("type:", vrow.Type())
	fmt.Println("kind:", vrow.Kind())
	fmt.Println("keys:", vrow.MapKeys())


	for key, val := range vrow.MapKeys() {
		fmt.Println(">>", key, val)
	}

	iter := vrow.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Println(k, v)
	}
```
