package cert

import (
	"fmt"
	"strings"
	"time"
)

//MaxLenCourse => Taille max d'un cours
var MaxLenCourse int = 20

//Cert => le structure de nos certificats
type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

//Saver => notre interface servant a sauvegarder
type Saver interface {
	Save()
}

//New => Nouveau certificat
func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is presented to",
		LabelParticipation: fmt.Sprintf("For the participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")), //Ceci n'est pas la date affiché mais le format uniquement
	}

	return cert, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, " course") { //On verifice que c se fini par " course"
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 {
		return c, fmt.Errorf("Invalid string got=%s, len=%d", c, len(c))
	}
	if len(c) > MaxLenCourse {
		return c, fmt.Errorf("Invalid string got=%s, len=%d too long", c, len(c))
	}
	return c, nil
}

func validateName(name string) (string, error) {
	n := strings.TrimSpace(name)
	if len(n) <= 0 {
		return n, fmt.Errorf("Error: name is empty")
	}
	if len(n) > 50 {
		return n, fmt.Errorf("Error : %v is too much caraters", len(n))
	}
	return n, nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date) //On met la date au format indiqué en premier lieu
	if err != nil {
		return t, err
	}
	return t, nil
}
