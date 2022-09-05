module laonanhai

go 1.16

require (
	github.com/Shopify/sarama v1.32.0
	github.com/astaxie/beego v1.12.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/hpcloud/tail v1.0.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/olivere/elastic/v7 v7.0.32
	go.etcd.io/etcd/api/v3 v3.5.2
	go.etcd.io/etcd/client/v3 v3.5.2
//gopkg.in/olivere/elastic/v7
)

//replace github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6

//replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
//replace (
//	github.com/olivere/elastic/v7 => gopkg.in/olivere/elastic.v7
//)
