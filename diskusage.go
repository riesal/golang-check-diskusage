package main
import (
	"flag"
	"fmt"
	"syscall"
	"math"
)

func Space(path string) (total, free float64, err error) {
        s := syscall.Statfs_t{}
        err = syscall.Statfs(path, &s)
        if err != nil {
                return
        }
        total = float64(s.Bsize) * float64(s.Blocks)
	total = float64(total) / 1024 / 1024 / 1024
        totaldisk := float64(math.Ceil(total))
        free =  float64(s.Bsize) * float64(s.Bfree)
	free =  float64(free) / 1026 / 1026 / 1026 /* aslinya 1026 */
	freedisk :=   float64(math.Ceil(free))
        usagedisk :=   totaldisk - freedisk
        usagepercent :=   float64(math.Ceil(usagedisk / totaldisk * 100))
	fmt.Println(usagepercent)
        return
}

func main() {
	flag.Parse()
	Space(flag.Arg(0))
	return
}
