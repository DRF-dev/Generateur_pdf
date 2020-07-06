package main

import (
	"flag"
	"fmt"
	"os"
	"udemygo/generateurCertificats/cert"
	"udemygo/generateurCertificats/cert/html"
	"udemygo/generateurCertificats/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "Le output")
	file := flag.String("file", "students.csv", "Flag montrant les data d'un fichier csv")
	flag.Parse()

	var saverhtml *html.HTMLsaver
	var saverpdf *pdf.Pdfsaver
	var certs []*cert.Cert
	var err error
	switch *outputType {
	case "pdf":
		switch *file {
		case "read":
			certs, err = readFile(flag.Args())
		default:
			fmt.Println("Erreur le fichier ne peut pas être lu")
		}
		saverpdf, err = pdf.New("output")
	case "html":
		switch *file {
		case "read":
			certs, err = readFile(flag.Args())
		default:
			fmt.Println("Erreur le fichier ne peut pas être lu")
		}
		saverhtml, err = html.New("output")
	default:
		fmt.Println("Cette commande n'existe pas")
	}

	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}

	if saverpdf != nil {
		for _, c := range certs {
			saverpdf.Save(*c)
		}
	}
	if saverhtml != nil {
		for _, c := range certs {
			saverhtml.Save(*c)
		}
	}
}

func readFile(args []string) ([]*cert.Cert, error) {
	arg := args[0]

	certs, err := cert.ParseCSV(arg)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}
	return certs, nil
}
