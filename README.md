# basic-latin
Convert strings to Unicode's Basic Latin in Go

### Example

```
package main

import (
	"fmt"

	"github.com/alex-ant/basic-latin"
)

func main() {
	s1 := "ÀƂĈĐÊƑĢĦĨĴĶĽMŇØƤQŘŞŦŨƲŴXŶƵ"
	fmt.Printf("Capital letters: %s\n", bl.ToBasicLatin(s1)) // Capital letters: ABCDEFGHIJKLMNOPQRSTUVWXYZ

	s2 := "áƃćđèƒġĥĩǰƙŀmņòƥqȓșţùvŵxýź"
	fmt.Printf("Small letters: %s\n", bl.ToBasicLatin(s2)) // Small letters: abcdefghijklmnopqrstuvwxyz
}
```
