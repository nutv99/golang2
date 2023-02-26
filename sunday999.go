package sunday999

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

/*
echo "# golang2" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/nutv99/golang2.git
git push -u origin main
*/

/*
git remote add origin https://github.com/nutv99/golang2.git
git branch -M main
git push -u origin main
*/
