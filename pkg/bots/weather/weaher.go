package weather

import (
	"fmt"
	"os"
	"strings"

	owm "github.com/briandowns/openweathermap"
	"github.com/jdkato/prose/v2"
	log "github.com/sirupsen/logrus"

	"github.com/sevigo/kumabot/pkg/core"
)

type weatherBot struct {
	keyWord string
	name    string
	// TODO: add mutex here, not sure it's save
	weatherData *owm.CurrentWeatherData
}

// TODO: is it better to return (core.KumaBot, error) here? Not sure
func New() core.KumaBot {
	// TODO:  maybe move this to the cmd/config and pass the config variable here?
	apiKey := os.Getenv("OWM_API_KEY")
	// TODO: add this part to Run(), so we can change the configuration depending on
	// the user settings we can get from the profile
	w, err := owm.NewCurrent("C", "en", apiKey)
	if err != nil {
		panic(err)
	}

	return &weatherBot{
		keyWord:     "weather",
		name:        "weather",
		weatherData: w,
	}
}

func (w *weatherBot) Match(in string) bool {
	log.WithField("bot", w.name).Infof("Match(): %q", in)
	inNorm := strings.ToLower(in)
	// TODO: better to train the model to detect requests like:
	// "will it raind tomorrow?"
	if strings.Contains(inNorm, w.keyWord) {
		_, err := w.getPlace(in)
		if err != nil {
			log.WithField("bot", w.name).
				WithError(err).Error("Match(): error")
			return false
		}
		return true
	}
	return false
}

// TODO: probaly we will need data from the user profile here to cusomize the response
func (w *weatherBot) Run(in string) string {
	log.WithField("bot", w.name).Infof("Run(): %q", in)
	place, err := w.getPlace(in)
	if err != nil {
		log.WithField("bot", w.name).
			WithError(err).Error("Run(): error")
		return ""
	}
	w.weatherData.CurrentByName(place)
	return w.formatMessageTxt()
}

func (w *weatherBot) Name() string {
	return w.name
}

func (w *weatherBot) getPlace(in string) (string, error) {
	log.WithField("bot", w.name).Infof("getPlace(): %q", in)
	// Add ? to the end of the string, for wharever reason it's not working without
	if !strings.HasSuffix(in, "?") {
		in = in + "?"
	}
	// TODO: we need better traind model to detect places
	// and something like 'today', 'tomorrow', 'on the weekend'
	doc, err := prose.NewDocument(in)
	if err != nil {
		return "", err
	}

	places := []string{}
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// GPE: geographical/political Entities
		if ent.Label == "GPE" {
			places = append(places, ent.Text)
		}
	}
	if len(places) != 1 {
		return "", fmt.Errorf("no geo place found in the input")
	}
	return places[0], nil
}

// TODO: add this function to the interface
func (w *weatherBot) formatMessageTxt() string {
	name := w.weatherData.Name
	if name == "" {
		return "I don't know the weatcher there"
	}
	text := ""
	if len(w.weatherData.Weather) > 0 {
		text = w.weatherData.Weather[0].Description
	}
	return fmt.Sprintf("Weather in %s is:\n%s\nTemperature: %d °C\nFeels like: %d °C\nHumidity: %d%%\n",
		name, text,
		int(w.weatherData.Main.Temp),
		int(w.weatherData.Main.FeelsLike),
		w.weatherData.Main.Humidity)
}
