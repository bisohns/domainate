package main

import (
	"flag"
	"github.com/bisoncorps/domainate/pkg/generate"
	"github.com/bisoncorps/domainate/pkg/thesaurus"
	"github.com/deven96/gosock/pkg/custlog"
)

// LogFile : global logfile to be used
var LogFile = flag.String("log", "domainate.log", "Name of the log file to save to")
var apikey = flag.String("apikey", "asdaksjdjfsjasdjasd", "API key of BigAPI")

func main() {
	flag.Parse() //parse the flags
	//    get the default writers for log
	defwriters := custlog.DefaultWriters(*LogFile, false)
	word := "chow"
	custlog.LogInit(defwriters)
	custlog.Info.Println("Domainate started")
	custlog.Info.Printf(`Example domain names using %v as keyword => %v`, word, generate.Domains(word))
	thesaurus := &thesaurus.BigHuge{APIKey: *apikey}
	syns, err := thesaurus.Synonyms(word)
	if err != nil {
		custlog.Error.Println("Failed when looking for synonyms for" + word + " ")
	}
	if len(syns) == 0 {
		custlog.Error.Println("Couldn't find any synonyms for" + word + " ")
	}
	if len(syns) > 0 {
		custlog.Info.Printf(`Synonyms found are %v`, syns)
	}
}
