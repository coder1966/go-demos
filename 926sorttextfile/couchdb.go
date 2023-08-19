package main

import (
	"fmt"
	"sort"
	"strings"
)

func couchdb() {
	strs := []string{}
	for k, v := range mesus {
		s := v.desc
		s = strings.TrimSpace(s)
		if !strings.HasSuffix(s, ".") {
			s = s + "."
		}
		s = FirstUpper(s)
		str := fmt.Sprintf(prnStr, k, s)
		// fmt.Println()
		strs = append(strs, str)
	}
	sort.Strings(strs)
	for _, v := range strs {
		fmt.Println(v)

	}
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

var prnStr = `"%s":&inputs.FieldInfo{DataType: inputs.Float, Type: inputs.Count, Unit: inputs.NCount, Desc: "%s"},`

type m struct {
	t    string
	desc string
}

var mesus = map[string]m{

	"auth_cache_hits_total":                               m{"counter", "Number of authentication cache hits."}, // add
	"auth_cache_misses_total":                             m{"counter", "Number of authentication cache misses."},
	"auth_cache_requests_total":                           m{"counter", "Microsecond latency for calls to couch_db:collect_results."},
	"collect_results_time_seconds":                        m{"summary", "microsecond latency for calls to couch_db:collect_results."},
	"couch_log_requests_total":                            m{"counter", "number of logged <level> messages. <level> = 'alert' 'critical' 'debug' 'emergency' 'error' 'info' 'notice' 'warning'"},
	"couch_replicator_changes_manager_deaths_total":       m{"counter", "number of failed replicator changes managers"},
	"couch_replicator_changes_queue_deaths_total":         m{"counter", "number of failed replicator changes work queues"},
	"couch_replicator_changes_read_failures_total":        m{"counter", "number of failed replicator changes read failures"},
	"couch_replicator_changes_reader_deaths_total":        m{"counter", "number of failed replicator changes readers"},
	"couch_replicator_checkpoints_failure_total":          m{"counter", "number of failed checkpoint saves"},
	"couch_replicator_checkpoints_total":                  m{"counter", "number of checkpoints successfully saves"},
	"couch_replicator_cluster_is_stable":                  m{"gauge", "1 if cluster is stable, 0 if unstable"},
	"couch_replicator_connection_acquires_total":          m{"counter", "number of times connections are shared"},
	"couch_replicator_connection_closes_total":            m{"counter", "number of times a worker is gracefully shut down"},
	"couch_replicator_connection_creates_total":           m{"counter", "number of connections created"},
	"couch_replicator_connection_owner_crashes_total":     m{"counter", "number of times a connection owner crashes while owning at least one connection"},
	"couch_replicator_connection_releases_total":          m{"counter", "number of times ownership of a connection is released"},
	"couch_replicator_connection_worker_crashes_total":    m{"counter", "number of times a worker unexpectedly terminates"},
	"couch_replicator_db_scans_total":                     m{"counter", "number of times replicator db scans have been started"},
	"couch_replicator_docs_completed_state_updates_total": m{"counter", "number of 'completed' state document updates"},
	"couch_replicator_docs_db_changes_total":              m{"counter", "number of db changes processed by replicator doc processor"},
	"couch_replicator_docs_dbs_created_total":             m{"counter", "number of db shard creations seen by replicator doc processor"},
	"couch_replicator_docs_dbs_deleted_total":             m{"counter", "number of db shard deletions seen by replicator doc processor"},
	"couch_replicator_docs_dbs_found_total":               m{"counter", "number of db shard found by replicator doc processor"},
	"couch_replicator_docs_failed_state_updates_total":    m{"counter", "number of 'failed' state document updates"},
	"couch_replicator_failed_starts_total":                m{"counter", "number of replications that have failed to start"},
	"couch_replicator_jobs_adds_total":                    m{"counter", "number of jobs added to replicator scheduler"},
	"couch_replicator_jobs_crashed":                       m{"gauge", "replicator scheduler crashed jobs"},
	"couch_replicator_jobs_crashes_total":                 m{"counter", "number of job crashed noticed by replicator scheduler."},
	"couch_replicator_jobs_duplicate_adds_total":          m{"counter", "number of duplicate jobs added to replicator scheduler"},
	"couch_replicator_jobs_pending":                       m{"gauge", "replicator scheduler pending jobs"},
	"couch_replicator_jobs_removes_total":                 m{"counter", "number of jobs removed from replicator scheduler"},
	"couch_replicator_jobs_running":                       m{"gauge", "replicator scheduler running jobs"},
	"couch_replicator_jobs_starts_total":                  m{"counter", "number of jobs started by replicator scheduler"},
	"couch_replicator_jobs_stops_total":                   m{"counter", "number of jobs stopped by replicator scheduler"},
	"couch_replicator_jobs_total":                         m{"gauge", "total number of replicator scheduler jobs"},
	"couch_replicator_requests_total":                     m{"counter", "number of HTTP requests made by the replicator"},
	"couch_replicator_responses_failure_total":            m{"counter", "number of failed HTTP responses received by the replicator"},
	"couch_replicator_responses_total":                    m{"counter", "number of successful HTTP responses received by the replicator"},
	"couch_replicator_stream_responses_failure_total":     m{"counter", "number of failed streaming HTTP responses received by the replicator"},
	"couch_replicator_stream_responses_total":             m{"counter", "number of successful streaming HTTP responses received by the replicator"},
	"couch_replicator_worker_deaths_total":                m{"counter", "number of failed replicator workers"},
	"couch_replicator_workers_started_total":              m{"counter", "number of replicator workers started"},
	"couch_server_lru_skip_total":                         m{"counter", "number of couch_server LRU operations skipped."},
	"database_purges_total":                               m{"counter", "Number of times a database was purged."},
	"database_reads_total":                                m{"counter", "Number of times a document was read from a database."},
	"database_writes_total":                               m{"counter", "Number of times a database was changed."},
	"db_open_time_seconds":                                m{"summary", "Milliseconds required to open a database."},
	"dbinfo_seconds":                                      m{"summary", "Milliseconds required to DB info."},
	"ddoc_cache_requests_failures_total":                  m{"counter", "number of design doc cache requests failures"},
	"ddoc_cache_requests_recovery_total":                  m{"counter", "number of design doc cache requests recoveries"},
	"ddoc_cache_requests_total":                           m{"counter", "number of design doc cache requests"},
	"ddoc_cache_hit_total":                                m{"counter", "number of design doc cache hits"},
	"ddoc_cache_miss_total":                               m{"counter", "number of design doc cache misses"},
	"ddoc_cache_recovery_total":                           m{"counter", "number of design doc cache recoveries"},
	"document_inserts_total":                              m{"counter", "number of documents inserted."},
	"document_purges_failure_total":                       m{"counter", "number of failed document purge operations."},
	"document_purges_success_total":                       m{"counter", "number of successful document purge operations."},
	"document_purges_total_total":                         m{"counter", "number of total document purge operations."},
	"document_writes_total":                               m{"counter", "number of document write operations."},
	"dreyfus_httpd_search_seconds":                        m{"summary", "Distribution of overall search request latency as experienced by the end user"},
	"dreyfus_index_await_seconds":                         m{"summary", "length of an dreyfus_index await request"},
	"dreyfus_index_group1_seconds":                        m{"summary", "length of an dreyfus_index group1 request"},
	"dreyfus_index_group2_seconds":                        m{"summary", "length of an dreyfus_index group2 request"},
	"dreyfus_index_info_seconds":                          m{"summary", "length of an dreyfus_index info request"},
	"dreyfus_index_search_seconds":                        m{"summary", "length of an dreyfus_index search request"},
	"dreyfus_rpc_group1_seconds":                          m{"summary", "length of a group1 RPC worker"},
	"dreyfus_rpc_group2_seconds":                          m{"summary", "length of a group2 RPC worker"},
	"dreyfus_rpc_info_seconds":                            m{"summary", "length of an info RPC worker"},
	"dreyfus_rpc_search_seconds":                          m{"summary", "length of a search RPC worker"},
	"erlang_context_switches_total":                       m{"counter", "total number of context switches"},
	"erlang_dirty_cpu_scheduler_queues":                   m{"gauge", "the total size of all dirty CPU scheduler run queues"},
	"erlang_ets_table":                                    m{"gauge", "number of ETS tables"},
	"erlang_gc_collections_total":                         m{"counter", "number of garbage collections by the Erlang emulator"},
	"erlang_gc_words_reclaimed_total":                     m{"counter", "number of words reclaimed by garbage collections"},
	"erlang_io_recv_bytes_total":                          m{"counter", "the total number of bytes received through ports"},
	"erlang_io_sent_bytes_total":                          m{"counter", "the total number of bytes output to ports"},
	"erlang_memory_bytes":                                 m{"gauge", "size of memory (in bytes) dynamically allocated by the Erlang emulator"},
	"erlang_message_queue_max":                            m{"gauge", "maximum size across all message queues"},
	"erlang_message_queue_min":                            m{"gauge", "minimum size across all message queues"},
	"erlang_message_queues":                               m{"gauge", "total size of all message queues"},
	"erlang_message_queue_size":                           m{"gauge", "size of message queue"},
	"erlang_process_limit":                                m{"gauge", "the maximum number of simultaneously existing Erlang processes"},
	"erlang_processes":                                    m{"gauge", "the number of Erlang processes"},
	"erlang_reductions_total":                             m{"counter", "total number of reductions"},
	"erlang_scheduler_queues":                             m{"gauge", "the total size of all normal run queues"},
	"erlang_distribution_recv_oct_bytes_total":            m{"counter", "Number of bytes received by the socket."},
	"erlang_distribution_recv_cnt_packets_total":          m{"counter", "number of packets received by the socket."},
	"erlang_distribution_recv_max_bytes":                  m{"gauge", "size of the largest packet, in bytes, received by the socket."},
	"erlang_distribution_recv_avg_bytes":                  m{"gauge", "average size of packets, in bytes, received by the socket."},
	"erlang_distribution_recv_dvi_bytes":                  m{"gauge", "average packet size deviation, in bytes, received by the socket."},
	"erlang_distribution_send_oct_bytes_total":            m{"counter", "Number of bytes sent by the socket"},
	"erlang_distribution_send_cnt_packets_total":          m{"counter", "number of packets sent by the socket."},
	"erlang_distribution_send_max_bytes":                  m{"gauge", "size of the largest packet, in bytes, sent by the socket."},
	"erlang_distribution_send_avg_bytes":                  m{"gauge", "average size of packets, in bytes, sent by the socket."},
	"erlang_distribution_send_pend_bytes":                 m{"gauge", "number of bytes waiting to be sent by the socket."},
	"fabric_doc_update_errors_total":                      m{"counter", "number of document update errors"},
	"fabric_doc_update_mismatched_errors_total":           m{"counter", "number of document update errors with multiple error types"},
	"fabric_doc_update_write_quorum_errors_total":         m{"counter", "number of write quorum errors"},
	"fabric_open_shard_timeouts_total":                    m{"counter", "number of open shard timeouts"},
	"fabric_read_repairs_failures_total":                  m{"counter", "number of failed read repair operations"},
	"fabric_read_repairs_total":                           m{"counter", "number of successful read repair operations"},
	"fabric_worker_timeouts_total":                        m{"counter", "number of worker timeouts"},
	"global_changes_db_writes_total":                      m{"counter", "number of db writes performed by global changes"},
	"global_changes_event_doc_conflict_total":             m{"counter", "number of conflicted event docs encountered by global changes"},
	"global_changes_listener_pending_updates":             m{"gauge", "number of global changes updates pending writes in global_changes_listener"},
	"global_changes_rpcs_total":                           m{"counter", "number of rpc operations performed by global_changes"},
	"global_changes_server_pending_updates":               m{"gauge", "number of global changes updates pending writes in global_changes_server"},
	"httpd_aborted_requests_total":                        m{"counter", "number of aborted requests"},
	"httpd_dbinfo":                                        m{"summary", "distribution of latencies for calls to retrieve DB info"},
	"httpd_all_docs_timeouts_total":                       m{"counter", "number of HTTP all_docs timeouts."},
	"httpd_bulk_docs_seconds":                             m{"summary", "distribution of the number of docs in _bulk_docs requests."},
	"httpd_bulk_requests_total":                           m{"counter", "number of bulk requests."},
	"httpd_clients_requesting_changes_total":              m{"counter", "number of clients for continuous _changes."},
	"httpd_explain_timeouts_total":                        m{"counter", "number of HTTP _explain timeouts."},
	"httpd_find_timeouts_total":                           m{"counter", "number of HTTP find timeouts."},
	"httpd_partition_all_docs_requests_total":             m{"counter", "number of partition HTTP _all_docs requests."},
	"httpd_partition_all_docs_timeouts_total":             m{"counter", "number of partition HTTP all_docs timeouts."},
	"httpd_partition_explain_requests_total":              m{"counter", "number of partition HTTP _explain requests."},
	"httpd_partition_explain_timeouts_total":              m{"counter", "number of partition HTTP _explain timeouts."},
	"httpd_partition_find_requests_total":                 m{"counter", "number of partition HTTP _find requests."},
	"httpd_partition_find_timeouts_total":                 m{"counter", "number of partition HTTP find timeouts."},
	"httpd_partition_view_requests_total":                 m{"counter", "number of partition HTTP view requests."},
	"httpd_partition_view_timeouts_total":                 m{"counter", "number of partition HTTP view timeouts."},
	"httpd_purge_requests_total":                          m{"counter", "number of purge requests."},
	"httpd_request_methods":                               m{"counter", "number of HTTP <option> requests. <option> = 'COPY' 'DELETE' 'GET' 'HEAD' 'OPTIONS' 'POST' 'PUT'"},
	"httpd_requests_total":                                m{"counter", "number of HTTP requests."},
	"httpd_status_codes":                                  m{"counter", "number of HTTP <status_codes> responses. <status_codes> = 200 201 202 204 206 301 304 400 403 404 405 406 409 412 414 415 416 417 500 501 503"},
	"httpd_temporary_view_reads_total":                    m{"counter", "number of temporary view reads."},
	"httpd_view_reads_total":                              m{"counter", "number of view reads."},
	"httpd_view_timeouts_total":                           m{"counter", "number of HTTP view timeouts."},
	"io_queue2_search_count_total":                        m{"counter", "Search IO directly triggered by client requests"},
	"io_queue_search_total":                               m{"counter", "Search IO directly triggered by client requests"},
	"local_document_writes_total":                         m{"counter", "number of document write operations."},
	"mango_keys_examined_total":                           m{"counter", "number of keys examined by mango queries coordinated by this node."},
	"mango_docs_examined_total":                           m{"counter", "number of documents examined by mango queries coordinated by this node."},
	"mango_evaluate_selector_total":                       m{"counter", "number of mango selector evaluations."},
	"mango_query_invalid_index_total":                     m{"counter", "number of mango queries that generated an invalid index warning."},
	"mango_query_time_seconds":                            m{"summary", "length of time processing a mango query."},
	"mango_quorum_docs_examined_total":                    m{"counter", "number of documents examined by mango queries, using cluster quorum."},
	"mango_results_returned_total":                        m{"counter", "number of rows returned by mango queries."},
	"mango_too_many_docs_scanned_total":                   m{"counter", "number of mango queries that generated an index scan warning."},
	"mango_unindexed_queries_total":                       m{"counter", "number of mango queries that could not use an index."},
	"mem3_shard_cache_eviction_total":                     m{"counter", "number of shard cache evictions"},
	"mem3_shard_cache_hit_total":                          m{"counter", "number of shard cache hits"},
	"mem3_shard_cache_miss_total":                         m{"counter", "number of shard cache misses"},
	"nouveau_search_latency":                              m{"summary", "Distribution of overall search request latency as experienced by the end user"},
	"nouveau_active_searches_total":                       m{"counter", "number of active search requests"},
	"mrview_emits_total":                                  m{"counter", "number of invocations of `emit' in map functions in the view server."},
	"mrview_map_doc_total":                                m{"counter", "number of documents mapped in the view server."},
	"open_databases_total":                                m{"counter", "number of open databases."},
	"open_os_files_total":                                 m{"counter", "number of file descriptors CouchDB has open."},
	"pread_exceed_eof_total":                              m{"counter", "number of the attempts to read beyond end of db file."},
	"pread_exceed_limit_total":                            m{"counter", "number of the attempts to read beyond set limit."},
	"fsync_time":                                          m{"summary", "microseconds to call fsync"},
	"fsync_count_total":                                   m{"counter", "number of fsync calls"},
	"query_server_vdu_process_time_seconds":               m{"summary", "duration of validate_doc_update function calls"},
	"query_server_vdu_rejects_total":                      m{"counter", "number of rejections by validate_doc_update function."},
	"query_server_acquired_processes_total":               m{"counter", "number of acquired external processes."},   // add
	"query_server_process_starts_total":                   m{"counter", "number of OS process starts."},             // add
	"query_server_process_exists_total":                   m{"counter", "number of OS normal process exits."},       // add
	"query_server_process_errors_total":                   m{"counter", "number of OS error process exits."},        // add
	"query_server_process_prompts_total":                  m{"counter", "number of successful OS process prompts."}, // add
	"query_server_process_prompt_errors_total":            m{"counter", "number of OS process prompt errors."},      // add
	"request_time_seconds":                                m{"summary", "length of a request inside CouchDB without MochiWeb."},
	"commits_total":                                       m{"counter", "number of commits performed."},                             // add
	"coalesced_updates_interactive":                       m{"counter", "number of coalesced interactive updates"},                  // add
	"coalesced_updates_replicated":                        m{"counter", "number of coalesced replicated updates"},                   // add
	"legacy_checksums":                                    m{"counter", "number of legacy checksums found in couch_file instances"}, // add
	"rexi_buffered_total":                                 m{"counter", "number of rexi messages buffered"},
	"rexi_down_total":                                     m{"counter", "number of rexi_DOWN messages handled"},
	"rexi_dropped_total":                                  m{"counter", "number of rexi messages dropped from buffers"},
	"rexi_streams_timeout_stream_total":                   m{"counter", "number of rexi stream timeouts"},
	"rexi_streams_timeout_total":                          m{"counter", "number of rexi stream initialization timeouts"},
	"rexi_streams_timeout_wait_for_ack_total":             m{"counter", "number of rexi stream timeouts while waiting for acks"},
	"uptime_seconds":                                      m{"counter", "couchdb uptime"},
}
