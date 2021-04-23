# sort example

~~~go
import (
	"fmt"

	"github.com/gakkiyomi/galang/net"
	"github.com/gakkiyomi/galang/sort"
)

func main() {
	source := []int{1, 4, 2, 44, 22, 13, 222, 441, 3, 54}
	sort.BubbleSort(source)
	sort.HeapSort(source, true)
	sort.QuickSort(source)
	sort.MergeSort(source)
	sort.ShellSort(source)
}

~~~