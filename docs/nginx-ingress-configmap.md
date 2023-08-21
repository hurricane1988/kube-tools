## Configuration options

The following table shows a configuration option's name, type, and the default value:

|name|type|default|notes|
|:---|:---|:------|:----|
|[add-headers](#add-headers)|string|""||
|[allow-backend-server-header](#allow-backend-server-header)|bool|"false"||
|[allow-cross-namespace-resources](#allow-cross-namespace-resources)|bool|"true"||
|[allow-snippet-annotations](#allow-snippet-annotations)|bool|true||
|[annotations-risk-level](#annotations-risk-level)|string|Critical||
|[annotation-value-word-blocklist](#annotation-value-word-blocklist)|string array|""||
|[hide-headers](#hide-headers)|string array|empty||
|[access-log-params](#access-log-params)|string|""||
|[access-log-path](#access-log-path)|string|"/var/log/nginx/access.log"||
|[http-access-log-path](#http-access-log-path)|string|""||
|[stream-access-log-path](#stream-access-log-path)|string|""||
|[enable-access-log-for-default-backend](#enable-access-log-for-default-backend)|bool|"false"||
|[error-log-path](#error-log-path)|string|"/var/log/nginx/error.log"||
|[enable-modsecurity](#enable-modsecurity)|bool|"false"||
|[modsecurity-snippet](#modsecurity-snippet)|string|""||
|[enable-owasp-modsecurity-crs](#enable-owasp-modsecurity-crs)|bool|"false"||
|[client-header-buffer-size](#client-header-buffer-size)|string|"1k"||
|[client-header-timeout](#client-header-timeout)|int|60||
|[client-body-buffer-size](#client-body-buffer-size)|string|"8k"||
|[client-body-timeout](#client-body-timeout)|int|60||
|[disable-access-log](#disable-access-log)|bool|false||
|[disable-ipv6](#disable-ipv6)|bool|false||
|[disable-ipv6-dns](#disable-ipv6-dns)|bool|false||
|[enable-underscores-in-headers](#enable-underscores-in-headers)|bool|false||
|[enable-ocsp](#enable-ocsp)|bool|false||
|[ignore-invalid-headers](#ignore-invalid-headers)|bool|true||
|[retry-non-idempotent](#retry-non-idempotent)|bool|"false"||
|[error-log-level](#error-log-level)|string|"notice"||
|[http2-max-field-size](#http2-max-field-size)|string|""|DEPRECATED in favour of [large_client_header_buffers](#large-client-header-buffers)|
|[http2-max-header-size](#http2-max-header-size)|string|""|DEPRECATED in favour of [large_client_header_buffers](#large-client-header-buffers)|
|[http2-max-requests](#http2-max-requests)|int|0|DEPRECATED in favour of [keepalive_requests](#keepalive-requests)|
|[http2-max-concurrent-streams](#http2-max-concurrent-streams)|int|128||
|[hsts](#hsts)|bool|"true"||
|[hsts-include-subdomains](#hsts-include-subdomains)|bool|"true"||
|[hsts-max-age](#hsts-max-age)|string|"15724800"||
|[hsts-preload](#hsts-preload)|bool|"false"||
|[keep-alive](#keep-alive)|int|75||
|[keep-alive-requests](#keep-alive-requests)|int|1000||
|[large-client-header-buffers](#large-client-header-buffers)|string|"4 8k"||
|[log-format-escape-none](#log-format-escape-none)|bool|"false"||
|[log-format-escape-json](#log-format-escape-json)|bool|"false"||
|[log-format-upstream](#log-format-upstream)|string|`$remote_addr - $remote_user [$time_local] "$request" $status $body_bytes_sent "$http_referer" "$http_user_agent" $request_length $request_time [$proxy_upstream_name] [$proxy_alternative_upstream_name] $upstream_addr $upstream_response_length $upstream_response_time $upstream_status $req_id`||
|[log-format-stream](#log-format-stream)|string|`[$remote_addr] [$time_local] $protocol $status $bytes_sent $bytes_received $session_time`||
|[enable-multi-accept](#enable-multi-accept)|bool|"true"||
|[max-worker-connections](#max-worker-connections)|int|16384||
|[max-worker-open-files](#max-worker-open-files)|int|0||
|[map-hash-bucket-size](#max-hash-bucket-size)|int|64||
|[nginx-status-ipv4-whitelist](#nginx-status-ipv4-whitelist)|[]string|"127.0.0.1"||
|[nginx-status-ipv6-whitelist](#nginx-status-ipv6-whitelist)|[]string|"::1"||
|[proxy-real-ip-cidr](#proxy-real-ip-cidr)|[]string|"0.0.0.0/0"||
|[proxy-set-headers](#proxy-set-headers)|string|""||
|[server-name-hash-max-size](#server-name-hash-max-size)|int|1024||
|[server-name-hash-bucket-size](#server-name-hash-bucket-size)|int|`<size of the processor’s cache line>`|
|[proxy-headers-hash-max-size](#proxy-headers-hash-max-size)|int|512||
|[proxy-headers-hash-bucket-size](#proxy-headers-hash-bucket-size)|int|64||
|[plugins](#plugins)|[]string| ||
|[reuse-port](#reuse-port)|bool|"true"||
|[server-tokens](#server-tokens)|bool|"false"||
|[ssl-ciphers](#ssl-ciphers)|string|"ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384"||
|[ssl-ecdh-curve](#ssl-ecdh-curve)|string|"auto"||
|[ssl-dh-param](#ssl-dh-param)|string|""||
|[ssl-protocols](#ssl-protocols)|string|"TLSv1.2 TLSv1.3"||
|[ssl-session-cache](#ssl-session-cache)|bool|"true"||
|[ssl-session-cache-size](#ssl-session-cache-size)|string|"10m"||
|[ssl-session-tickets](#ssl-session-tickets)|bool|"false"||
|[ssl-session-ticket-key](#ssl-session-ticket-key)|string|`<Randomly Generated>`|
|[ssl-session-timeout](#ssl-session-timeout)|string|"10m"||
|[ssl-buffer-size](#ssl-buffer-size)|string|"4k"||
|[use-proxy-protocol](#use-proxy-protocol)|bool|"false"||
|[proxy-protocol-header-timeout](#proxy-protocol-header-timeout)|string|"5s"||
|[use-gzip](#use-gzip)|bool|"false"||
|[use-geoip](#use-geoip)|bool|"true"||
|[use-geoip2](#use-geoip2)|bool|"false"||
|[enable-brotli](#enable-brotli)|bool|"false"||
|[brotli-level](#brotli-level)|int|4||
|[brotli-min-length](#brotli-min-length)|int|20||
|[brotli-types](#brotli-types)|string|"application/xml+rss application/atom+xml application/javascript application/x-javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/javascript text/plain text/x-component"||
|[use-http2](#use-http2)|bool|"true"||
|[gzip-disable](#gzip-disable)|string|""||
|[gzip-level](#gzip-level)|int|1||
|[gzip-min-length](#gzip-min-length)|int|256||
|[gzip-types](#gzip-types)|string|"application/atom+xml application/javascript application/x-javascript application/json application/rss+xml application/vnd.ms-fontobject application/x-font-ttf application/x-web-app-manifest+json application/xhtml+xml application/xml font/opentype image/svg+xml image/x-icon text/css text/javascript text/plain text/x-component"||
|[worker-processes](#worker-processes)|string|`<Number of CPUs>`||
|[worker-cpu-affinity](#worker-cpu-affinity)|string|""||
|[worker-shutdown-timeout](#worker-shutdown-timeout)|string|"240s"||
|[load-balance](#load-balance)|string|"round_robin"||
|[variables-hash-bucket-size](#variables-hash-bucket-size)|int|128||
|[variables-hash-max-size](#variables-hash-max-size)|int|2048||
|[upstream-keepalive-connections](#upstream-keepalive-connections)|int|320||
|[upstream-keepalive-time](#upstream-keepalive-time)|string|"1h"||
|[upstream-keepalive-timeout](#upstream-keepalive-timeout)|int|60||
|[upstream-keepalive-requests](#upstream-keepalive-requests)|int|10000||
|[limit-conn-zone-variable](#limit-conn-zone-variable)|string|"$binary_remote_addr"||
|[proxy-stream-timeout](#proxy-stream-timeout)|string|"600s"||
|[proxy-stream-next-upstream](#proxy-stream-next-upstream)|bool|"true"||
|[proxy-stream-next-upstream-timeout](#proxy-stream-next-upstream-timeout)|string|"600s"||
|[proxy-stream-next-upstream-tries](#proxy-stream-next-upstream-tries)|int|3||
|[proxy-stream-responses](#proxy-stream-responses)|int|1||
|[bind-address](#bind-address)|[]string|""||
|[use-forwarded-headers](#use-forwarded-headers)|bool|"false"||
|[enable-real-ip](#enable-real-ip)|bool|"false"||
|[forwarded-for-header](#forwarded-for-header)|string|"X-Forwarded-For"||
|[compute-full-forwarded-for](#compute-full-forwarded-for)|bool|"false"||
|[proxy-add-original-uri-header](#proxy-add-original-uri-header)|bool|"false"||
|[generate-request-id](#generate-request-id)|bool|"true"||
|[enable-opentracing](#enable-opentracing)|bool|"false"||
|[opentracing-operation-name](#opentracing-operation-name)|string|""||
|[opentracing-location-operation-name](#opentracing-location-operation-name)|string|""||
|[zipkin-collector-host](#zipkin-collector-host)|string|""||
|[zipkin-collector-port](#zipkin-collector-port)|int|9411||
|[zipkin-service-name](#zipkin-service-name)|string|"nginx"||
|[zipkin-sample-rate](#zipkin-sample-rate)|float|1.0||
|[jaeger-collector-host](#jaeger-collector-host)|string|""||
|[jaeger-collector-port](#jaeger-collector-port)|int|6831||
|[jaeger-endpoint](#jaeger-endpoint)|string|""||
|[jaeger-service-name](#jaeger-service-name)|string|"nginx"||
|[jaeger-propagation-format](#jaeger-propagation-format)|string|"jaeger"||
|[jaeger-sampler-type](#jaeger-sampler-type)|string|"const"||
|[jaeger-sampler-param](#jaeger-sampler-param)|string|"1"||
|[jaeger-sampler-host](#jaeger-sampler-host)|string|"http://127.0.0.1"||
|[jaeger-sampler-port](#jaeger-sampler-port)|int|5778||
|[jaeger-trace-context-header-name](#jaeger-trace-context-header-name)|string|uber-trace-id||
|[jaeger-debug-header](#jaeger-debug-header)|string|uber-debug-id||
|[jaeger-baggage-header](#jaeger-baggage-header)|string|jaeger-baggage||
|[jaeger-trace-baggage-header-prefix](#jaeger-trace-baggage-header-prefix)|string|uberctx-||
|[datadog-collector-host](#datadog-collector-host)|string|""||
|[datadog-collector-port](#datadog-collector-port)|int|8126||
|[datadog-service-name](#datadog-service-name)|string|"nginx"||
|[datadog-environment](#datadog-environment)|string|"prod"||
|[datadog-operation-name-override](#datadog-operation-name-override)|string|"nginx.handle"||
|[datadog-priority-sampling](#datadog-priority-sampling)|bool|"true"||
|[datadog-sample-rate](#datadog-sample-rate)|float|1.0||
|[enable-opentelemetry](#enable-opentelemetry)|bool|"false"||
|[opentelemetry-trust-incoming-span](#opentelemetry-trust-incoming-span)|bool|"true"||
|[opentelemetry-operation-name](#opentelemetry-operation-name)|string|""||
|[opentelemetry-config](#/etc/nginx/opentelemetry.toml)|string|"/etc/nginx/opentelemetry.toml"||
|[otlp-collector-host](#otlp-collector-host)|string|""||
|[otlp-collector-port](#otlp-collector-port)|int|4317||
|[otel-max-queuesize](#otel-max-queuesize)|int|||
|[otel-schedule-delay-millis](#otel-schedule-delay-millis)|int|||
|[otel-max-export-batch-size](#otel-max-export-batch-size)|int|||
|[otel-service-name](#otel-service-name)|string|"nginx"||
|[otel-sampler](#otel-sampler)|string|"AlwaysOff"||
|[otel-sampler-parent-based](#otel-sampler-parent-based)|bool|"false"||
|[otel-sampler-ratio](#otel-sampler-ratio)|float|0.01||
|[main-snippet](#main-snippet)|string|""||
|[http-snippet](#http-snippet)|string|""||
|[server-snippet](#server-snippet)|string|""||
|[stream-snippet](#stream-snippet)|string|""||
|[location-snippet](#location-snippet)|string|""||
|[custom-http-errors](#custom-http-errors)|[]int|[]int{}||
|[proxy-body-size](#proxy-body-size)|string|"1m"||
|[proxy-connect-timeout](#proxy-connect-timeout)|int|5||
|[proxy-read-timeout](#proxy-read-timeout)|int|60||
|[proxy-send-timeout](#proxy-send-timeout)|int|60||
|[proxy-buffers-number](#proxy-buffers-number)|int|4||
|[proxy-buffer-size](#proxy-buffer-size)|string|"4k"||
|[proxy-cookie-path](#proxy-cookie-path)|string|"off"||
|[proxy-cookie-domain](#proxy-cookie-domain)|string|"off"||
|[proxy-next-upstream](#proxy-next-upstream)|string|"error timeout"||
|[proxy-next-upstream-timeout](#proxy-next-upstream-timeout)|int|0||
|[proxy-next-upstream-tries](#proxy-next-upstream-tries)|int|3||
|[proxy-redirect-from](#proxy-redirect-from)|string|"off"||
|[proxy-request-buffering](#proxy-request-buffering)|string|"on"||
|[ssl-redirect](#ssl-redirect)|bool|"true"||
|[force-ssl-redirect](#force-ssl-redirect)|bool|"false"||
|[denylist-source-range](#denylist-source-range)|[]string|[]string{}||
|[whitelist-source-range](#whitelist-source-range)|[]string|[]string{}||
|[skip-access-log-urls](#skip-access-log-urls)|[]string|[]string{}||
|[limit-rate](#limit-rate)|int|0||
|[limit-rate-after](#limit-rate-after)|int|0||
|[lua-shared-dicts](#lua-shared-dicts)|string|""||
|[http-redirect-code](#http-redirect-code)|int|308||
|[proxy-buffering](#proxy-buffering)|string|"off"||
|[limit-req-status-code](#limit-req-status-code)|int|503||
|[limit-conn-status-code](#limit-conn-status-code)|int|503||
|[enable-syslog](#enable-syslog)|bool|false||
|[syslog-host](#syslog-host)|string|""||
|[syslog-port](#syslog-port)|int|514||
|[no-tls-redirect-locations](#no-tls-redirect-locations)|string|"/.well-known/acme-challenge"||
|[global-auth-url](#global-auth-url)|string|""||
|[global-auth-method](#global-auth-method)|string|""||
|[global-auth-signin](#global-auth-signin)|string|""||
|[global-auth-signin-redirect-param](#global-auth-signin-redirect-param)|string|"rd"||
|[global-auth-response-headers](#global-auth-response-headers)|string|""||
|[global-auth-request-redirect](#global-auth-request-redirect)|string|""||
|[global-auth-snippet](#global-auth-snippet)|string|""||
|[global-auth-cache-key](#global-auth-cache-key)|string|""||
|[global-auth-cache-duration](#global-auth-cache-duration)|string|"200 202 401 5m"||
|[no-auth-locations](#no-auth-locations)|string|"/.well-known/acme-challenge"||
|[block-cidrs](#block-cidrs)|[]string|""||
|[block-user-agents](#block-user-agents)|[]string|""||
|[block-referers](#block-referers)|[]string|""||
|[proxy-ssl-location-only](#proxy-ssl-location-only)|bool|"false"||
|[default-type](#default-type)|string|"text/html"||
|[global-rate-limit-memcached-host](#global-rate-limit)|string|""||
|[global-rate-limit-memcached-port](#global-rate-limit)|int|11211||
|[global-rate-limit-memcached-connect-timeout](#global-rate-limit)|int|50||
|[global-rate-limit-memcached-max-idle-timeout](#global-rate-limit)|int|10000||
|[global-rate-limit-memcached-pool-size](#global-rate-limit)|int|50||
|[global-rate-limit-status-code](#global-rate-limit)|int|429||
|[service-upstream](#service-upstream)|bool|"false"||
|[ssl-reject-handshake](#ssl-reject-handshake)|bool|"false"||
|[debug-connections](#debug-connections)|[]string|"127.0.0.1,1.1.1.1/24"||
|[strict-validate-path-type](#strict-validate-path-type)|bool|"false" (v1.7.x)||