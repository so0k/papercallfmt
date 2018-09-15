package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Submission object represents papercall submission
type Submission struct {
	Abstract       string   `json:"abstract"`
	AdditionalInfo string   `json:"additional_info"`
	AudienceLevel  string   `json:"audience_level"`
	Avatar         string   `json:"avatar"`
	Bio            string   `json:"bio"`
	Confirmed      bool     `json:"confirmed"`
	CreatedAt      string   `json:"created_at"`
	Description    string   `json:"description"`
	Email          string   `json:"email"`
	Location       string   `json:"location"`
	Name           string   `json:"name"`
	Notes          string   `json:"notes"`
	Organization   string   `json:"organization"`
	Rating         float32  `json:"rating"`
	ShirtSize      string   `json:"shirt_size"`
	State          string   `json:"state"`
	Tags           []string `json:"tags"`
	TalkFormat     string   `json:"talk_format"`
	Title          string   `json:"title"`
	Twitter        string   `json:"twitter"`
	URL            string   `json:"url"`
}

var build = "0"

func main() {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "log-level",
			Value:  "error",
			Usage:  "Log level (panic, fatal, error, warn, info, or debug)",
			EnvVar: "LOG_LEVEL",
		},
		cli.StringFlag{
			Name:  "source, s",
			Value: "download.json",
			Usage: "Source `json`",
		},
		cli.StringFlag{
			Name:  "destination, d",
			Value: "output/speakers/",
			Usage: "Destination `directory` to render in - must exist",
		},
		cli.StringFlag{
			Name:  "template, t",
			Value: "templates/speaker.md.tpl",
			Usage: "Desired template used to render output",
		},
		cli.StringFlag{
			Name:  "state-filter, f",
			Value: "accepted",
			Usage: "Filter to only submissions of this state",
		},
	}
	app := cli.NewApp()
	app.Name = "Papercall Format"
	app.Usage = "Parse Papercall submissions from json"
	app.Action = run

	app.Version = fmt.Sprintf("0.1.%s", build)
	app.Author = "so0k"

	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	logLevelString := c.String("log-level")
	logLevel, err := log.ParseLevel(logLevelString)
	if err != nil {
		return err
	}
	log.SetLevel(logLevel)

	log.Debugf("source: %v", c.String("source"))
	log.Debugf("destination: %v", c.String("destination"))
	log.Debugf("template: %v", c.String("template"))

	//TODO: Add more validations
	dstDirName := path.Dir(c.String("destination"))
	log.Debugf("dstDirName: %v", dstDirName)
	if stat, err := os.Stat(dstDirName); err != nil || !stat.IsDir() {
		fmt.Printf("Invalid destination: %v\n", dstDirName)
		cli.ShowAppHelpAndExit(c, 1)
	}

	tpl := c.String("template") //expect name.xxx.tpl
	suffix := path.Ext(strings.TrimSuffix(tpl, ".tpl"))
	log.Debugf("suffix: %v", suffix)

	submissions, err := readSubmissions(c.String("source"))
	if err != nil {
		return err
	}

	for _, s := range *submissions {
		if s.State == c.String("state-filter") {
			linkName := convertName(s.Name)
			p := path.Join(dstDirName, fmt.Sprintf("%v%v", linkName, suffix))
			f, err := os.Create(p)
			if err != nil {
				return err
			}

			log.Debugf("rendering template for %s to %s", s.Name, p)
			err = RenderTemplate(&s, tpl, f)
			f.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func convertName(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "-", -1)
	return s
}

func readSubmissions(jsonfile string) (*[]Submission, error) {
	s := make([]Submission, 0)
	data, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
