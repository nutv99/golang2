package nutvjson

import (
	"encoding/json"
	"fmt"
	"strings"
)

var stAll string = ""

var AllQueryText string = ""
var DSN string = ""

var JSONString string = ""

var MainTableName string = "headOrder"
var sqlMain string = "INSERT INTO " + MainTableName + "("
var ValueMain string = " VALUES("
var QueryChildArray []string

func Mainnutv() {
	jsonString := `{ 
		"name": "John",
		"age": "30",
		"dddd" : [
			{
				"dddField1": "dddData",
				"dddField2": "dddData2"
			}
		],
		"mmm" : [
			{
				"mmmfirld1": "mmmData",
				"mmmfield2": "mmmData2"
			}
		],
		"address": [{
			"addressStreet": "123 Main St",
			"addresscity": "New York"
			
		},{
			"addressStreet": "999 กรุงเทพ-ปทุม",
			"addresscity": "New York"
		 
		}
		],
		"WareHouse": [{
			"WareHouseaddressStreet": "123 Main St",
			"WareHouseaddresscity": "New York"
			
		},{
			"WareHouseaddressStreet": "999 กรุงเทพ-ปทุม",
			"WareHouseaddresscity": "New York"
		 
		}
		]
		,
		"sendplace": [{
			"SendPlacestreet": "SendPlace123 Main St",
			"SendPlacecity": "SendPlaceNew York"
			
		},{
			"SendPlacestreet": "SendPlace999 กรุงเทพ-ปทุม",
			"SendPlacecity": "SendPlaceNew York"
			
		}
		],
		"email": "john@example.com"		
	}`

	var data interface{}
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		panic(err)
	}

	process(data, "")
	Process2()
	//fmt.Println("Final", stAll)
	fmt.Println("QueryChildArray", QueryChildArray)

	for i := 0; i <= len(QueryChildArray)-1; i++ {
		AllQueryText += QueryChildArray[i] + " ; "
	}

}

func appendIfNotExists(arr []string, elem string) []string {
	for _, a := range arr {
		if a == elem {
			return arr // If element already exists, return the original array
		}
	}
	// If element doesn't exist, append it to the array
	return append(arr, elem)
}

func searchArray(arr []string, elem string) bool {
	for _, a := range arr {
		if a == elem {
			return true // If element already exists, return the original array
		}
	}
	// If element doesn't exist, append it to the array
	return false
}

func Process2() {

	var dataSeries string = ""
	// import strings
	stAll = strings.TrimSuffix(stAll, "@#")
	//fmt.Println("newStALL", stAll)

	var sqlChild string = ""
	var ValueChild string = ""
	var fieldAppend []string
	var oldTableName string
	var newChildQuery bool = false

	res1 := strings.Split(stAll, "@#")
	for i := 0; i <= len(res1)-1; i++ {
		if res1[i] != "" {
			res2 := strings.Split(res1[i], ":")
			res3 := strings.Split(res2[0], ".")
			if len(res3) == 1 {
				// case 1 MainTable
				sqlMain += res2[0] + ","
				ValueMain += "'" + res2[1] + "',"
				sqlChild2 := strings.TrimSuffix(sqlChild, ",") + ")"
				ValueChild2 := strings.TrimSuffix(ValueChild, ",") + ")"
				thisSqlChild := sqlChild2 + ValueChild2
				fmt.Println(thisSqlChild)
				QueryChildArray = appendIfNotExists(QueryChildArray, thisSqlChild)
			} else {
				// case 2 ChildTable
				tablenameTmp := strings.Split(res3[0], "[")
				tablename := tablenameTmp[0]

				if sqlChild == "" {
					sqlChild = "INSERT INTO " + tablename + "(" + res3[1] + ","
					ValueChild = "VALUES ('" + res2[1] + "',"
					newChildQuery = true
				} else {
					if oldTableName != tablename && oldTableName != "" {
						sqlChild = strings.TrimSuffix(sqlChild, ",") + ")"
						ValueChild = strings.TrimSuffix(ValueChild, ",") + ")"
						thisSqlChild := sqlChild + ValueChild
						fmt.Println(thisSqlChild)
						QueryChildArray = appendIfNotExists(QueryChildArray, thisSqlChild)
						sqlChild = "INSERT INTO " + tablename + "(" + res3[1] + ","
						ValueChild = "VALUES ('" + res2[1] + "',"
						fieldAppend = make([]string, 0, cap(fieldAppend)) // Clear Array
						newChildQuery = true
					}

					thisDataSeries := strings.TrimSuffix(tablenameTmp[1], "]")
					if thisDataSeries != dataSeries && oldTableName == tablename {
						// import strings
						newChildQuery = false
						if dataSeries != "" {
							sqlChild = strings.TrimSuffix(sqlChild, ",") + ")"
							ValueChild = strings.TrimSuffix(ValueChild, ",") + ")"
							thisSqlChild := sqlChild + ValueChild
							fmt.Println(thisSqlChild)
							QueryChildArray = appendIfNotExists(QueryChildArray, thisSqlChild)
							sqlChild = "INSERT INTO " + tablename + "(" + res3[1] + ","
							ValueChild = "VALUES ('" + res2[1] + "',"
							fieldAppend = make([]string, 0, cap(fieldAppend)) // Clear Array
						} else {
							sqlChild += res3[1] + ","
							ValueChild += "'" + res2[1] + "',"
							fieldAppend = appendIfNotExists(fieldAppend, res3[1])
						}

						dataSeries = thisDataSeries
					} else {
						if !searchArray(fieldAppend, res3[1]) && newChildQuery == false {
							sqlChild += res3[1] + ","
							ValueChild += "'" + res2[1] + "',"
							fieldAppend = appendIfNotExists(fieldAppend, res3[1])
						}
						dataSeries = thisDataSeries
						newChildQuery = true

					}
					oldTableName = tablename

				}

				// import strings

				//fmt.Println(tablename, " Series=", dataSeries)

			}
			aa := 2
			newChildQuery = false
			aa = aa
		}
		bb := 22
		bb = bb

	} // end for แต่ละ  for คือ 1 record

	// import strings
	sqlMain = strings.TrimSuffix(sqlMain, ",") + ")"

	ValueMain = strings.TrimSuffix(ValueMain, ",") + ")"

	sqlMain += ValueMain
	fmt.Println("sqlMain", sqlMain)
	//fmt.Println("ValueMain", ValueMain)

}

func process3() {

	var dataSeries string = ""
	// import strings
	stAll = strings.TrimSuffix(stAll, "@#")
	//fmt.Println("newStALL", stAll)

	var sqlChild string = ""
	var ValueChild string = ""
	var fieldAppend []string

	res1 := strings.Split(stAll, "@#")
	for i := 0; i <= len(res1)-1; i++ {
		if res1[i] != "" {
			res2 := strings.Split(res1[i], ":")
			res3 := strings.Split(res2[0], ".")
			if len(res3) == 1 {
				// case 1 MainTable
				sqlMain += res2[0] + ","
				ValueMain += "'" + res2[1] + "',"
				sqlChild2 := strings.TrimSuffix(sqlChild, ",") + ")"
				ValueChild2 := strings.TrimSuffix(ValueChild, ",") + ")"
				thisSqlChild := sqlChild2 + ValueChild2
				fmt.Println(thisSqlChild)
				QueryChildArray = appendIfNotExists(QueryChildArray, thisSqlChild)
			} else {
				// case 2 ChildTable
				tablenameTmp := strings.Split(res3[0], "[")
				tablename := tablenameTmp[0]
				if sqlChild == "" {
					sqlChild = "INSERT INTO " + tablename + "(" + res3[1] + ","
					ValueChild = "VALUES ('" + res2[1] + "',"
				} else {
					thisDataSeries := strings.TrimSuffix(tablenameTmp[1], "]")
					if thisDataSeries != dataSeries {
						// import strings
						if dataSeries != "" {
							sqlChild = strings.TrimSuffix(sqlChild, ",") + ")"
							ValueChild = strings.TrimSuffix(ValueChild, ",") + ")"
							thisSqlChild := sqlChild + ValueChild
							fmt.Println(thisSqlChild)
							QueryChildArray = appendIfNotExists(QueryChildArray, thisSqlChild)
							sqlChild = "INSERT INTO " + tablename + "(" + res3[1] + ","
							ValueChild = "VALUES ('" + res2[1] + "',"
							fieldAppend = make([]string, 0, cap(fieldAppend)) // Clear Array
						} else {
							sqlChild += res3[1] + ","
							ValueChild += "'" + res2[1] + "',"
							fieldAppend = appendIfNotExists(fieldAppend, res3[1])
						}
						dataSeries = thisDataSeries
					} else {
						if !searchArray(fieldAppend, res3[1]) {
							sqlChild += res3[1] + ","
							ValueChild += "'" + res2[1] + "',"
							fieldAppend = appendIfNotExists(fieldAppend, res3[1])
						}
					}
					// dataSeries := strings.TrimSuffix(tablenameTmp[1], "]")
					// dataSeries = dataSeries
					// if !searchArray(fieldAppend, res3[1]) {
					// 	sqlChild += res3[1] + ","
					// 	fieldAppend = appendIfNotExists(fieldAppend, res3[1])
					// }
				}
				// import strings

				//fmt.Println(tablename, " Series=", dataSeries)

			}
			aa := 2
			aa = aa
		}
		bb := 22
		bb = bb

	} // end for แต่ละ  for คือ 1 record

	// import strings
	sqlMain = strings.TrimSuffix(sqlMain, ",") + ")"

	ValueMain = strings.TrimSuffix(ValueMain, ",") + ")"

	sqlMain += ValueMain
	fmt.Println("sqlMain", sqlMain)
	//fmt.Println("ValueMain", ValueMain)

}

func process(data interface{}, parentKey string) {
	switch val := data.(type) {
	case map[string]interface{}:
		for key, value := range val {
			newKey := ""
			if parentKey != "" {
				newKey = parentKey + "." + key

			} else {
				newKey = key
			}

			process(value, newKey)
		}
	case []interface{}:
		for i, value := range val {
			newKey := fmt.Sprintf("%s[%d]", parentKey, i)
			//stAll += parentKey + ":" + value.(string) + "@#"
			process(value, newKey)
		}
	default:
		if parentKey != "" {
			fmt.Printf("%s: %v\n", parentKey, val)
			stAll += parentKey + ":" + val.(string) + "@#"
		} else {
			fmt.Printf("Value: %v\n", val)
			stAll += val.(string) + "@#"
		}
	}

}
