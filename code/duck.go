type Speaker interface {
    Speak(string)
}

type Person struct {
    name  string
}

func (p *Person) Speak(text string) {
    log.Println(p.name+": ", text)
}

func CheerForSponsor(speaker Speaker, sponsor string) {
    speaker.Speak("Thank you %s for supporting this conference", sponsor)
}

calvin := Person{name: "Calvin", title: "Why Go is awesome"}
CheerForSponsors(calvin, "Google Cloud")
