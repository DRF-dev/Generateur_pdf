package pdf

import (
	"fmt"
	"os"
	"path"
	"udemygo/generateurCertificats/cert"

	"github.com/jung-kurt/gofpdf"
)

//Pdfsaver => pdfSaver
type Pdfsaver struct {
	OutputDir string
}

//New => nouveau dir
func New(outputdir string) (*Pdfsaver, error) {
	var p *Pdfsaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &Pdfsaver{
		OutputDir: outputdir,
	}
	return p, nil
}

//Save => Sauvegarde de nos PDF
func (p *Pdfsaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "") // Orienté paysage, taille en mm, format A4, font par défaut
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	//Background
	background(pdf)

	//Header
	header(pdf, &cert)
	pdf.Ln(30) //Saut de ligne

	//body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	//body - student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	//body - Participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	//body - Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	//Footer
	footer(pdf)

	//Sauvegarder le fichier
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate %v\n", path)
	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeigth := pdf.GetPageSize()
	pdf.ImageOptions("img/téléchargement.png", 0, 0, pageWidth, pageHeigth, false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	ImageWidth := 30.0
	filename := "img/600_331101642.png"
	pdf.ImageOptions(filename, x+margin, 20, ImageWidth, 0, false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - ImageWidth
	pdf.ImageOptions(filename, x-margin, 20, ImageWidth, 0, false, opts, 0, "")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	filename := "img/stamp.png"
	largeur, hauteur := pdf.GetPageSize()
	imageWidth := 50.0
	x := largeur - imageWidth - 20
	y := hauteur - imageWidth - 10

	pdf.ImageOptions(filename, x, y, imageWidth, 0, false, opts, 0, "")
}
