# UDP Forwarder

UDP Forwarder is a simple UDP packet forwarder which listen and forward from the source port to the destination port. With this utility you will be able to filter out the unwanted UDP packets before forwarding to the destination port. 

This utility uses simple regex pattern to match the payload data.

```
echo "uwsgi.worker.0.core.0.write_errors:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.1.core.0.write_errors:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.0.delta_requests:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.1.delta_requests:1|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.0.avg_response_time:10|c" | nc -w 1 -u localhost 8126
echo "uwsgi.worker.1.avg_response_time:10|c" | nc -w 1 -u localhost 8126
