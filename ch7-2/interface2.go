package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"math"

	person "example.com/go-demo/ch7-2/person"
	eval "example.com/go-demo/ch7-2/eval"
)

func main() {
	interVal()
	typeAssert()
	sortInter()
	httpTest()
	evalTest()
}

func interVal() {
	// interface value = dtype + dvalue
	fmt.Println("interVal = type + value")
	var w io.Writer      // 接口变量声明（此时类型和值均为nil）
	w = os.Stdout        // 赋值后：类型=*os.File，值=os.Stdout
	fmt.Printf("%#v\n", w)
	w = new(bytes.Buffer) // 重新赋值：类型=*bytes.Buffer，值=新创建的Buffer
	fmt.Printf("%#v\n", w)
}

func typeAssert() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File)    // 成功：f=os.Stdout, ok=true
	// fmt.Println(f, ok)
	fmt.Printf("%T %#v\n", f, ok)
	b, ok := w.(*bytes.Buffer) // 失败：b=nil, ok=false
	// fmt.Println(b, ok)
	fmt.Printf("%#v %#v\n", b, ok)
}

func sortInter() {
	arr := []int{3, 1, 2, 4, 4}
	sort.Ints(arr)
	fmt.Println(arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
	sort.Stable(sort.IntSlice(arr))
	fmt.Println(arr)

	people := person.ByAge {
		person.Person{Name: "seuer1", Age: 1},
		person.Person{Name: "seuer2", Age: 2},
		person.Person{Name: "seuer3", Age: 3},
	}

	sort.Sort(people)
	fmt.Println(people)
	sort.Sort(sort.Reverse(person.ByAge(people)))
	fmt.Println(people)
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

// func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//     for item, price := range db {
//         fmt.Fprintf(w, "%s: %s\n", item, price)
//     }
// }

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
    case "/list":
        for item, price := range db {
            fmt.Fprintf(w, "%s: %s\n", item, price)
        }
    case "/price":
        item := req.URL.Query().Get("item")
        price, ok := db[item]
        if !ok {
            w.WriteHeader(http.StatusNotFound) // 404
            fmt.Fprintf(w, "no such item: %q\n", item)
            // 上面两行也可以替换为
            // msg := fmt.Sprintf("no such page: %s\n", req.URL)
// http.Error(w, msg, http.StatusNotFound) // 404
            return
        }
        fmt.Fprintf(w, "%s\n", price)
    default:
        w.WriteHeader(http.StatusNotFound) // 404
        fmt.Fprintf(w, "no such page: %s\n", req.URL)
    }
}

func httpTest() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func evalTest() {
	tests := []struct {
        expr string
        env  eval.Env
        want string
    }{
        {"sqrt(A / pi)", eval.Env{"A": 87616, "pi": math.Pi}, "167"},
        {"pow(x, 3) + pow(y, 3)", eval.Env{"x": 12, "y": 1}, "1729"},
        {"pow(x, 3) + pow(y, 3)", eval.Env{"x": 9, "y": 10}, "1729"},
        {"5 / 9 * (F - 32)", eval.Env{"F": -40}, "-40"},
        {"5 / 9 * (F - 32)", eval.Env{"F": 32}, "0"},
        {"5 / 9 * (F - 32)", eval.Env{"F": 212}, "100"},
    }
    var prevExpr string
    for _, test := range tests {
        // Print expr only when it changes.
        if test.expr != prevExpr {
            fmt.Printf("\n%s\n", test.expr)
            prevExpr = test.expr
        }
        expr, err := eval.Parse(test.expr)
        if err != nil {
            // t.Error(err) // parse error
            continue
        }
        got := fmt.Sprintf("%.6g", expr.Eval(test.env))
        fmt.Printf("\t%v => %s\n", test.env, got)
        if got != test.want {
            // t.Errorf("%s.Eval() in %v = %q, want %q\n",
            // test.expr, test.env, got, test.want)
			fmt.Println("戳啦！")
		}
    }
}