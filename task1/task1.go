package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)
func main() {


elveid := 0
scanner := bufio.NewScanner(os.Stdin)
foods := make(map[int][]int)


for scanner.Scan() {  
  line := scanner.Text()
  if line == "" {
    elveid +=1
  } else {
    kcal, err := strconv.Atoi(line)
    if err != nil {
      fmt.Println("Error during conversion")
      return
    }
    foods[elveid] = append(foods[elveid], kcal)
  }
    
}

if err := scanner.Err(); err != nil {
  panic(err)
}


var ss []int
for _, callories := range foods {
  sum := 0 
  for _, val := range callories {
    sum += val
  }
  ss = append(ss, sum)
}
sort.Slice(ss, func(i, j int) bool {
  return ss[i] > ss[j]
})

for _, kv := range ss[:1] {
  fmt.Printf("%d\n", kv)
}

sum :=0
for _, kv := range ss[:3] {
  sum += kv
}
fmt.Printf("%d\n", sum)

}