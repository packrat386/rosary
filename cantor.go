package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Cantor struct {
	lock    *sync.RWMutex
	display *bytes.Buffer
	beat    time.Duration
	pos     int64
}

func NewCantor() *Cantor {
	return &Cantor{
		lock:    new(sync.RWMutex),
		display: new(bytes.Buffer),
		beat:    configBeat(),
	}
}

func (c *Cantor) reset() {
	c.pos = 0
	c.resetDisplay()
}

func (c *Cantor) chant() {
	for {
		c.reset()

		mysteries := todaysMysteries()

		c.sayApostlesCreed()
		c.sayOurFather()
		c.sayHailMary()
		c.sayHailMary()
		c.sayHailMary()
		c.sayGloryBe()

		for _, m := range mysteries {
			c.announceMystery(m)
			c.sayDecade()
		}

		c.announceIntention()
		c.sayHailHolyQueen()
	}
}

func todaysMysteries() []string {
	w := time.Now().Weekday()

	switch w {
	case time.Monday, time.Saturday:
		return []string{
			"The first mystery is the Annunciation.\n",
			"The second mystery is the Visitation.\n",
			"The third mystery is the Nativity.\n",
			"The fourth mystery is the Presentation at the Temple.\n",
			"The fifth mystery is the Finding in the Temple.\n",
		}
	case time.Tuesday, time.Friday:
		return []string{
			"The first mystery is the Agony in the Garden.\n",
			"The second mystery is the Scourging at the Pillar.\n",
			"The third mystery is the Crowning with Thorns.\n",
			"The fourth mystery is the Carrying of the Cross.\n",
			"The fifth mystery is the Cruxifixion and Death.\n",
		}
	case time.Thursday:
		return []string{
			"The first mystery is the baptism of Christ in the Jordan.\n",
			"The second mystery is the wedding feast at Cana.\n",
			"The third mystery is the proclamation of the Kingdom of God.\n",
			"The fourth mystery is the Transfiguration.\n",
			"The fifth mystery is the Institution of the Eucharist.\n",
		}
	case time.Wednesday, time.Sunday:
		return []string{
			"The first mystery is the Resurrection.\n",
			"The second mystery is the Ascension.\n",
			"The third mystery is the descent of the Holy Spirit.\n",
			"The fourth mystery is the Assumption of the Virgin Mary.\n",
			"The fifth mystery is the Coronation of Mary.\n",
		}
	}

	panic(fmt.Sprintf("Unrecognized weekday: %d", w))

}

func (c *Cantor) announceMystery(l string) {
	c.resetDisplay()
	c.sayLine(l, 15)
	c.breath()
	c.breath()
	c.breath()
	c.breath()
}

func (c *Cantor) announceIntention() {
	c.resetDisplay()
	c.sayLine("This rosary is offered for peace. Now more than ever.\n", 15)
	c.breath()
	c.breath()
	c.breath()
}

func (c *Cantor) sayDecade() {
	c.sayOurFather()

	for i := 0; i < 10; i++ {
		c.sayHailMary()
	}

	c.sayGloryBe()
}

func (c *Cantor) sayApostlesCreed() {
	c.resetDisplay()
	c.sayLine("I believe in God,\n", 5)
	c.sayLine("the Father almighty,\n", 5)
	c.sayLine("Creator of heaven and earth,\n", 8)
	c.sayLine("and in Jesus Christ, his only Son, our Lord,\n", 10)
	c.sayLine("who was conceived by the Holy Spirit,\n", 10)
	c.sayLine("born of the Virgin Mary,\n", 6)
	c.sayLine("suffered under Pontius Pilate,\n", 8)
	c.sayLine("was crucified, died and was buried;\n", 9)
	c.sayLine("he descended into hell;\n", 6)
	c.sayLine("on the third day he rose again from the dead;\n", 11)
	c.sayLine("he ascended into heaven,\n", 7)
	c.sayLine("and is seated at the right hand of God the Father almighty;\n", 16)
	c.sayLine("from there he will come to judge the living and the dead.\n", 13)
	c.sayLine("I believe in the Holy Spirit,\n", 9)
	c.sayLine("the holy catholic Church,\n", 6)
	c.sayLine("the communion of saints,\n", 6)
	c.sayLine("the forgiveness of sins,\n", 6)
	c.sayLine("the resurrection of the body,\n", 8)
	c.sayLine("and life everlasting.\n", 6)
	c.sayLine("Amen.\n", 2)
	c.breath()
	c.breath()

	c.pos += 1
}

func (c *Cantor) sayOurFather() {
	c.resetDisplay()
	c.sayLine("Our Father, who art in heaven,\n", 7)
	c.sayLine("hallowed be thy name;\n", 5)
	c.sayLine("thy kingdom come;\n", 4)
	c.sayLine("thy will be done on earth as it is in heaven.\n", 12)
	c.sayLine("Give us this day our daily bread;\n", 8)
	c.sayLine("and forgive us our trespasses\n", 8)
	c.sayLine("as we forgive those who trespass against us;\n", 11)
	c.sayLine("and lead us not into temptation,\n", 9)
	c.sayLine("but deliver us from evil.\n", 8)
	c.sayLine("Amen\n", 2)
	c.breath()
	c.breath()

	c.pos += 1
}

func (c *Cantor) sayHailMary() {
	c.resetDisplay()
	c.sayLine("Hail, Mary,\n", 3)
	c.sayLine("full of grace,\n", 3)
	c.sayLine("the Lord is with thee.\n", 5)
	c.sayLine("Blessed art thou amongst women\n", 6)
	c.sayLine("and blessed is the fruit of thy womb, Jesus.\n", 10)
	c.sayLine("Holy Mary, Mother of God,\n", 4)
	c.sayLine("pray for us sinners,\n", 4)
	c.sayLine("now and at the hour of our death.\n", 9)
	c.sayLine("Amen.\n", 2)
	c.breath()

	c.pos += 1
}

func (c *Cantor) sayGloryBe() {
	c.resetDisplay()
	c.sayLine("Glory be to the Father, the Son, and the Holy Spirit;\n", 15)
	c.sayLine("as it was in the beginning, is now, and ever shall be,\n", 15)
	c.sayLine("world without end.\n", 4)
	c.sayLine("Amen.\n", 2)
	c.breath()
	c.breath()

	c.pos += 1
}

func (c *Cantor) sayHailHolyQueen() {
	c.resetDisplay()
	c.sayLine("Hail, holy Queen, mother of mercy,\n", 9)
	c.sayLine("our life, our sweetness, and our hope.\n", 8)
	c.sayLine("To you we cry, poor banished children of Eve;\n", 11)
	c.sayLine("to you we send up our sighs,\n", 7)
	c.sayLine("mourning and weeping in this valley of tears.\n", 11)
	c.sayLine("Turn, then, most gracious advocate,\n", 8)
	c.sayLine("your eyes of mercy toward us;\n", 7)
	c.sayLine("and after this, our exile,\n", 7)
	c.sayLine("show unto us the blessed fruit of your womb, Jesus.\n", 13)
	c.sayLine("O clement, O loving, O sweet Virgin Mary.\n", 12)
	c.sayLine("Amen.\n", 2)
	c.breath()
	c.breath()
}

func (c *Cantor) breath() {
	c.sayLine(".\n", 4)
}

func (c *Cantor) getDisplay() []byte {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.display.Bytes()
}

func (c *Cantor) resetDisplay() {
	c.lock.RLock()
	defer c.lock.RUnlock()

	c.display.Reset()
	c.display.WriteString("\u001B[2J\u001B[3J\u001B[H\n\n")

	c.display.WriteString("+Oooo-Ooooooooooo-Ooooooooooo-Ooooooooooo-Ooooooooooo-Ooooooooooo-@\n")

	spaces := []byte(strings.Repeat(" ", 67))
	spaces[c.pos] = '^'
	c.display.Write(spaces)
	c.display.WriteString("\n\n")
}

func (c *Cantor) appendDisplay(s string) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	c.display.WriteString(s)
}

func (c *Cantor) sayLine(line string, beats int64) {
	c.appendDisplay(line)
	time.Sleep(time.Duration(beats) * c.beat)
}
