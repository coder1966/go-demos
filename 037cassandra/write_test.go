/*
create keyspace testwrite with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 2 };
create table testwrite.tweet(timeline text, id UUID, text text, linenumber int, PRIMARY KEY(id));
create index on testwrite.tweet(timeline);
*/

package a0037cassandra

import "testing"

func Test_write(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			write()
		})
	}
}
