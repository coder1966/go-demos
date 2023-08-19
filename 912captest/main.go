package main

import (
	"fmt"
)

// opentelemetry,host=zhangub-OMEN-by-HP-Laptop-15-dc1xxx,operation=startRootSpan,resource_name=/,service=dktrace-otel-agent,service_name=dktrace-otel-agent,source_type=custom,span_type=web,status=ok duration=237013i,message="{\"trace_id\":\"RTXnv/SvgjdqYNQiWOT0kZSkluKPyIKR/lKR52i5\",\"span_id\":\"7peXV49LwnY=\",\"name\":\"startRootSpan\",\"kind\":1,\"start_time_unix_nano\":1686707957120476562,\"end_time_unix_nano\":1686707957357489730,\"attributes\":[{\"key\":\"resource.name\",\"value\":{\"Value\":{\"StringValue\":\"/\"}}},{\"key\":\"span.type\",\"value\":{\"Value\":{\"StringValue\":\"web\"}}}],\"status\":{}}",parent_id="0",priority=1i,resource="startRootSpan",span_id="ee9797578f4bc276",start=1686707957120476i,trace_id="4535e7bff4af82376a60d42258e4f49194a496e28fc88291fe5291e768b9" 1686707957120476562

func main() {
	// A. array
	// B. slice
	// C. map
	// D. channel

	a := [5]int{}
	s := make([]int, 2, 5)
	m := make(map[string]int, 4)
	c := make(chan int, 12)
	fmt.Println("a = ", cap(a), a)
	fmt.Println("s = ", cap(s), s)
	// fmt.Println("m = ",cap(m),m)
	_ = m
	fmt.Println("c = ", cap(c), c)
}
