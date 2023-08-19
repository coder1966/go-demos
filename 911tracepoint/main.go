package main

import (
	"fmt"

	"github.com/GuanceCloud/cliutils/point"
)

// opentelemetry,host=zhangub-OMEN-by-HP-Laptop-15-dc1xxx,operation=startRootSpan,resource_name=/,service=dktrace-otel-agent,service_name=dktrace-otel-agent,source_type=custom,span_type=web,status=ok duration=237013i,message="{\"trace_id\":\"RTXnv/SvgjdqYNQiWOT0kZSkluKPyIKR/lKR52i5\",\"span_id\":\"7peXV49LwnY=\",\"name\":\"startRootSpan\",\"kind\":1,\"start_time_unix_nano\":1686707957120476562,\"end_time_unix_nano\":1686707957357489730,\"attributes\":[{\"key\":\"resource.name\",\"value\":{\"Value\":{\"StringValue\":\"/\"}}},{\"key\":\"span.type\",\"value\":{\"Value\":{\"StringValue\":\"web\"}}}],\"status\":{}}",parent_id="0",priority=1i,resource="startRootSpan",span_id="ee9797578f4bc276",start=1686707957120476i,trace_id="4535e7bff4af82376a60d42258e4f49194a496e28fc88291fe5291e768b9" 1686707957120476562

func main() {
	message := "{\"trace_id\":\"RTXnv/SvgjdqYNQiWOT0kZSkluKPyIKR/lKR52i5\",\"span_id\":\"7peXV49LwnY=\",\"name\":\"startRootSpan\",\"kind\":1,\"start_time_unix_nano\":1686707957120476562,\"end_time_unix_nano\":1686707957357489730,\"attributes\":[{\"key\":\"resource.name\",\"value\":{\"Value\":{\"StringValue\":\"/\"}}},{\"key\":\"span.type\",\"value\":{\"Value\":{\"StringValue\":\"web\"}}}],\"status\":{}}"
	_ = message
	metricName := "otel"

	opts := point.DefaultMetricOptions()
	tags := map[string]string{}
	fields := map[string]interface{}{}

	tags["host"] = "HP-serer"
	tags["user"] = "root"

	fields["float64Data"] = float64(1.234)
	fields["int64Data"] = int64(5678)
	// fields["stringData"] = message

	pt := point.NewPointV2([]byte(metricName), append(point.NewTags(tags), point.NewKVs(fields)...), opts...)
	_ = pt

	kvs := point.NewKVs(map[string]any{"f1": 123})

	kvs = kvs.Add([]byte(`t2`), []byte(message), false, true)

	for _, v := range kvs {
		pt.AddKV(v)
	}

	fmt.Println("pt.LineProto() == ", pt.LineProto())
	return
}
