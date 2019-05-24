# stats-forwarder

```
echo "uwsgi.worker.0.core.0.write_errors:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.1.core.0.write_errors:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.0.delta_requests:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.1.delta_requests:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.0.avg_response_time:10|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.1.avg_response_time:10|c" | nc -w 1 -u localhost 8126
