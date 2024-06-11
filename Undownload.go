package github.com/jroden2/dumb

import (
	"fmt"
	"strings"
	"time"
)

func Undownload() {
	const col = 30
	bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar+" %s", strings.Repeat("=", i)+">", "Downloading...")
		time.Sleep(100 * time.Millisecond)
	}
  // [==============================] Downloaded!
	fmt.Printf(bar+" Downloaded!", strings.Repeat("=", col))


  // Prints an '[ERROR                         ] Downloaded!' message
	fmt.Printf(bar+" Downloaded!", "ERROR")
	time.Sleep(500 * time.Millisecond)

  // Wipes the bar [==============================>] Downloaded!
	for i := col; i > 0; i-- {
		fmt.Printf(bar+" %s", strings.Repeat("=", i)+">", "Undownloading...")
		time.Sleep(100 * time.Millisecond)
	}
  // [                              ] Undownloaded!
	fmt.Printf(bar+" Undownloaded!", strings.Repeat(" ", col))
}
