module database-example

go 1.18

require (
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.25.7-0.20240204074919-46816ad31dde
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.32.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.32.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.7.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	github.com/XSAM/otelsql v0.14.1
)

require github.com/felixge/httpsnoop v1.0.3 // indirect

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/handlers v1.5.2
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lib/pq v1.10.9
	github.com/neo4j/neo4j-go-driver/v5 v5.19.0
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gorm.io/driver/postgres v1.5.6 // indirect
)
