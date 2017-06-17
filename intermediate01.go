// Write a program that outputs all possibilities to put + or - or nothing
// between the numbers 1,2,…,9 (in this order) such that the result is 100. For
// example 1 + 2 + 3 - 4 + 5 + 6 + 78 + 9 = 100.
//
// There are 3 possibilities for each position, and 8 positions.
// This gives a total amount of possibilities 3**8 = 6561, which is very
// feasible to loop trough.
package main

import "fmt"
import "strconv"
import "github.com/roessland/gopkg/iterutil"

const (
    PLUS = iota
    MINUS
    CONCAT
)

func ToRunes(ops []int64) []rune {
    str := make([]rune, len(ops))
    for i := 0; i < len(ops); i++ {
        var r rune
        switch ops[i] {
        case PLUS:
            r = '+'
        case MINUS:
            r = '-'
        case CONCAT:
            r = ' '
        default:
            r = 'Æ'
        }
        str[i] = r
    }
    return str
}

func TestToRunes() {
    var r []rune
    r = ToRunes([]int64{PLUS})
    if len(r) != 1 {
        panic("wrong length")
    }
    if r[0] != '+' {
        panic("wrong op!")
    }

    r = ToRunes([]int64{MINUS, PLUS, CONCAT})
    if len(r) != 3 {
        panic("wrong length")
    }
    if r[0] != '-' {
        panic("wrong op!")
    }
    if r[1] != '+' {
        panic("wrong op!")
    }
    if r[2] != ' ' {
        panic("wrong op!")
    }
}

func ConcatInts(a, b int64) int64 {
    ab, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
    if err != nil {
        panic("concatints")
    }
    return int64(ab)
}

func TestConcatInts() {
    if ConcatInts(5, 7) != 57 {
        panic("5 . 7 != 57!")
    }

    if ConcatInts(123456, 789) != 123456789 {
        panic("123456 . 789 != 123456789")
    }
}

func Sum(nums[]int64, ops []int64) int64 {
	if len(nums) - 1 != len(ops) {
        panic("There are not enough operations")
    }
    if len(nums) == 1 {
        return nums[0]
    }
    foundConcat := false
    var i int64
    for i = 0; i < int64(len(ops)); i = i+1 {
        if ops[i] == CONCAT {
            foundConcat = true
            break
        }
    }

    // Make copies of input slices so we don't change them
    nextNums := make([]int64, len(nums))
    copy(nextNums, nums)

    if foundConcat {
        // Recurse at the point where the operator was found
        nextNums[i] = ConcatInts(nums[i], nums[i+1])
        nextNums = append(nextNums[:i+1], nextNums[i+2:]...)
        nextOps := append(ops[:i], ops[i+1:]...)
        return Sum(nextNums, nextOps)
    } else {
        // Recurse from start
        if ops[0] == PLUS {
            nextNums[1] = nextNums[0] + nextNums[1]
        } else if ops[0] == MINUS {
            nextNums[1] = nextNums[0] - nextNums[1]
        } else {
            fmt.Printf("Encountered unknown operator%d.\n", ops[0])
            fmt.Printf("    %v\n", ops)
            panic("unknown operator")
        }
        return Sum(nextNums[1:], ops[1:])
    }
    panic("This should never happen")
}

func TestSum() {
    var S int64

    S = Sum([]int64{1,2}, []int64{CONCAT})
    if S != 12 {
        fmt.Printf("1.2 was %v, but wanted 12\n", S)
        panic("S12")
    }

    S = Sum([]int64{1,2,3}, []int64{CONCAT, CONCAT})
    if S != 123 {
        fmt.Printf("1.2.3 was %v, but wanted 123\n", S)
        panic("S123")
    }

    S = Sum([]int64{1,2}, []int64{PLUS})
    if S != 3 {
        fmt.Printf("1+2 was %v, but wanted 3\n", S)
        panic("S3")
    }

    S = Sum([]int64{1,2,3,4,5,6,7,8,9}, []int64{MINUS, PLUS, CONCAT, CONCAT, MINUS, MINUS, PLUS, PLUS})
    if S != 348 {
        fmt.Printf("1-2+3.4.5-6-7+8+9 was %v, but wanted 348\n", S)
        panic("123456789")
    }

    S = Sum([]int64{1,2,3,4,5,6,7,8,9}, []int64{PLUS, PLUS, MINUS, PLUS, PLUS, PLUS, PLUS, PLUS})
    if S != 37 {
        fmt.Printf("1+2+3-4+5+6+7+8+9 was %v but wanted 37\n", S)
        panic("BANANA")
    }

    S = Sum([]int64{1,2,3,4,5,6,7,8,9}, []int64{MINUS, PLUS, PLUS, CONCAT, CONCAT, CONCAT, CONCAT, CONCAT})
    if S != 456791 {
        fmt.Printf("1-2+3+456789 was %v but wanted 456791\n", S)
        panic("YODEL")
    }
}

func main() {
    TestConcatInts()
    TestSum()
    TestToRunes()

    combChan := iterutil.CartesianPower(3, 8)
	for comb := range combChan {
        r := ToRunes(comb)
        S := Sum([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, comb)
        if S == 100 {
            fmt.Printf("1 %c 2 %c 3 %c 4 %c 5 %c 6 %c 7 %c 8 %c 9 = %d\n", r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], S)
		}
	}
}
