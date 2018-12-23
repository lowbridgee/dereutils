package main

import (
	"bytes"
	"log"
	"os"
	"github.com/urfave/cli"
	"github.com/knakk/sparql"
)

func main() {
	app := cli.NewApp()

	app.Name = "dereutils"
	app.Usage = "The interface of im@sparql"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name: "find_unit",
			Usage: "Find unit name by idol name.",
			Action: func(c *cli.Context) error {
				name := os.Args[2]
				res := findUnitByMemberName(name)
				m := res.Results.Bindings
				for _, v := range m {
					println(v["ユニット名"].Value)
				}
				return nil
			},
		},
		{
			Name: "grep",
			Usage: "find idol",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func findUnitByMemberName(name string) *sparql.Results {
	repo, err := sparql.NewRepo("https://sparql.crssnky.xyz/spql/imas/query")
	if err != nil {
		log.Fatal(err)
	}
	const query = `
# tag: find-unit-by-member-idol
	PREFIX schema: <http://schema.org/>
	PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
	PREFIX imas: <https://sparql.crssnky.xyz/imasrdf/URIs/imas-schema.ttl#>
	PREFIX imasrdf: <https://sparql.crssnky.xyz/imasrdf/RDFs/detail/>
	PREFIX foaf: <http://xmlns.com/foaf/0.1/>
	PREFIX math: <http://www.w3.org/2005/xpath-functions/math#>
	PREFIX xsd: <https://www.w3.org/TR/xmlschema11-2/#>
	PREFIX rdfs:  <http://www.w3.org/2000/01/rdf-schema#>
	SELECT  ?ユニット名 (group_concat(?名前;separator=", ")as ?メンバー)
	WHERE {
	  ?s rdf:type imas:Unit;
		 schema:name ?ユニット名;
		 schema:member ?m.
	  ?m schema:name ?名前.
	  filter contains (?名前, "{{.Name}}").
	}group by (?ユニット名) order by(?ユニット名)`
	
	f := bytes.NewBufferString(query)
	bank := sparql.LoadBank(f)
	
	sql, err := bank.Prepare("find-unit-by-member-idol", struct{ Name string }{name})
	if err != nil {
		log.Fatal(err)
	}
	
	res, err := repo.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return res
}