    resources:
      cpu:
        metrics:
          - traffic: cpu_util_percent
          - latency: runqueue_latency
          - saturation: queue_length
          - other: ipc, pressure
        traffic:
          - cpu_util_percent
        latency:
          - runqueue_latency
        errors:
        saturation:
          - queue_length
        other:
          - ipc
          - pressure
      memory:
        traffic:
          - mem_util_percent
        errors:
          - oom_kill
        other:
          - pressure
      diskio:
        traffic:
          - read_bps
          - write_bps
        latency:
          - await_time
        errors:
          - filesystem_errors
          - disk_errors
        saturation:
          - io_queue_depth
        other:
          - pressure
      network:
        traffic:
          - net_in
          - net_out
        latency:
          - driver_queue
        errors:
          - tcp_drops
          - device_errors
        saturation:
          - overruns
          - retransmissions