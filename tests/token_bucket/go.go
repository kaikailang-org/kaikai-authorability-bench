import "fmt"
import "os"

func main() {
	ok := true
	b, o := NewBucket(5, 2)
	ok = ok && o && b.Tokens == 5 && b.Capacity == 5 && b.Refill == 2

	b2 := Tick(Tick(Drain(b)))
	ok = ok && b2.Tokens == 4

	b3 := Tick(b2)
	ok = ok && b3.Tokens == 5 // cap

	h3 := Bucket{Capacity: 5, Tokens: 3, Refill: 2}
	r1, o1 := TryAcquire(h3, 3)
	ok = ok && o1 && r1.Tokens == 0

	r2, o2 := TryAcquire(h3, 4)
	ok = ok && !o2 && r2.Tokens == 3

	r3, o3 := TryAcquire(h3, 0)
	ok = ok && !o3 && r3.Tokens == 3

	r4, o4 := TryAcquire(h3, -2)
	ok = ok && !o4 && r4.Tokens == 3

	_, o5 := NewBucket(0, 1)
	ok = ok && !o5
	_, o6 := NewBucket(3, -1)
	ok = ok && !o6
	z, o7 := NewBucket(3, 0)
	ok = ok && o7 && Tick(Drain(z)).Tokens == 0

	if !ok {
		fmt.Println("FAIL")
		os.Exit(1)
	}
	fmt.Println("OK")
}
