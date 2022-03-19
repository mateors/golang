# Group filter code

```go
func groupFilteredRows(rows []map[string]interface{}, gFields []string) []map[string]interface{} {

	var rmap = make([]map[string]interface{}, 0)
	var guvs = make(map[string][]interface{}) //group wise unique values

	//var sumFields = []string{"Amount"}
	var sumkv = make(map[string]float64)
	//var total float64

	for _, row := range rows {

		//fmt.Println(i, row)

		for _, grpname := range gFields {

			//fmt.Println(i, j, grpname)
			if vals, isExist := guvs[grpname]; !isExist { //new

				var uvals = make([]interface{}, 0)
				uvals = append(uvals, row[grpname])
				guvs[grpname] = uvals
				//var amounts = make([]interface{}, 0)
				//guvs["Amount"] = append(amounts, row["Amount"])
				//fmt.Println(i, j, vals, grpname, row[grpname])
				amountf, _ := strconv.ParseFloat(fmt.Sprintf("%v", row["Amount"]), 64)
				sumkv[grpname] = amountf

			} else { //exist

				isFound := FindExist(vals, row[grpname])
				if !isFound {
					uvals := append(vals, row[grpname])
					guvs[grpname] = uvals
					//fmt.Println(grpname, vals, row[grpname], isFound)
					amountf, _ := strconv.ParseFloat(fmt.Sprintf("%v", row["Amount"]), 64)
					eamountf := sumkv[grpname]
					sumkv[grpname] = eamountf + amountf
				}
			}
		}
		//fmt.Println(total)
	}

	for key, vals := range guvs {
		//fmt.Printf("%v %v %T\n", key, vali, vali)
		for _, val := range vals {
			kv := make(map[string]interface{})
			// kval, isOk := val.(string)
			// if isOk {
			// 	kv[kval] = key
			// }
			// if !isOk {
			// 	log.Println(val)
			// 	kv[fmt.Sprintf("%v", val)] = key
			// }
			kv[fmt.Sprintf("%v", val)] = key
			//kv["Amount"] = sumkv[key]
			rmap = append(rmap, kv)
		}
	}
	// fmt.Println(len(rmap))
	for key, val := range sumkv {
		fmt.Println(">>", key, val)
	}
	return rmap
}
```
