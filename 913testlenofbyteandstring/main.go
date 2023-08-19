package main

import (
	"fmt"
)

// opentelemetry,host=zhangub-OMEN-by-HP-Laptop-15-dc1xxx,operation=startRootSpan,resource_name=/,service=dktrace-otel-agent,service_name=dktrace-otel-agent,source_type=custom,span_type=web,status=ok duration=237013i,message="{\"trace_id\":\"RTXnv/SvgjdqYNQiWOT0kZSkluKPyIKR/lKR52i5\",\"span_id\":\"7peXV49LwnY=\",\"name\":\"startRootSpan\",\"kind\":1,\"start_time_unix_nano\":1686707957120476562,\"end_time_unix_nano\":1686707957357489730,\"attributes\":[{\"key\":\"resource.name\",\"value\":{\"Value\":{\"StringValue\":\"/\"}}},{\"key\":\"span.type\",\"value\":{\"Value\":{\"StringValue\":\"web\"}}}],\"status\":{}}",parent_id="0",priority=1i,resource="startRootSpan",span_id="ee9797578f4bc276",start=1686707957120476i,trace_id="4535e7bff4af82376a60d42258e4f49194a496e28fc88291fe5291e768b9" 1686707957120476562

func main() {
	strs := []string{
		"1234567890",
		"1234567890\t",
		"1234567890αλφάβητ",
		"1234567890一二三",
		"1234567890\n一二三",
	}
	for _, v := range strs {
		fmt.Println("错误", v, len(v), len([]byte(v)))

		if len(v) != len([]byte(v)) {
			panic("出错了")
		}
	}

	fmt.Println("执行结束")
}
