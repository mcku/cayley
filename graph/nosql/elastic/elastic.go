package elastic

import (
	"context"

	"github.com/cayleygraph/cayley/graph"

	"github.com/mcku/hidalgo/legacy/nosql"
	"github.com/mcku/hidalgo/legacy/nosql/elastic"

	//import hidal-go first so the registration of the no sql stores occurs before quadstore iterates for registration
	gnosql "github.com/cayleygraph/cayley/graph/nosql"
)

const Type = elastic.Name

func Create(addr string, opt graph.Options) (nosql.Database, error) {
	return elastic.Dial(context.TODO(), addr, gnosql.DefaultDBName, nosql.Options(opt))
}

func Open(addr string, opt graph.Options) (nosql.Database, error) {
	return elastic.Dial(context.TODO(), addr, gnosql.DefaultDBName, nosql.Options(opt))
}
