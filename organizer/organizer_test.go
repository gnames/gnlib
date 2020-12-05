package organizer_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	org "github.com/gnames/gnlib/organizer"
	"github.com/stretchr/testify/assert"
)

var txt = []rune(
	`One perplexity of the novice is that different books may give different
names for the same shell. The causes of this diversity are explained on a
later page. With the facilities of the largest museum in America, the author
has been able to speak with authority in those matters of nomenclature.
When the problem is zoological and still to be solved by further collections,
or by the study of living moUusks, then the cooperation of the keen collector
may give the answer sought. Professional malacologists are few. Their work
is largely in museums with dead animals. The interesting but long task of
collecting from a thousand miles of coast, and observing mollusks alive, has
always been in large part a labor of love by private naturalists. Our science
owes nearly as much to them as to the work of professional zoologists.

The author belongs to the younger group of malacologists, but he has
cultivated the society of mollusks in many lands, from East Africa, the`,
)

type ordRune struct {
	i int
	r rune
}

func (or ordRune) Index() int {
	return or.i
}

func (or ordRune) Unpack(v interface{}) error {
	switch r := v.(type) {
	case *rune:
		*r = or.r
	default:
		return fmt.Errorf("type %T is not a rune", r)
	}
	return nil
}

func newOrdered(i int, r rune) org.Ordered {
	return ordRune{i: i, r: r}
}

func TestOrganizer(t *testing.T) {
	chIn := make(chan org.Ordered)
	chOut := make(chan org.Ordered)
	res := make([]rune, 0, len(txt))
	var wg sync.WaitGroup
	wg.Add(1)
	shuffled := make([]org.Ordered, len(txt))
	for i, v := range txt {
		shuffled[i] = newOrdered(i, v)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	go org.Organize(chIn, chOut)
	go func() {
		defer wg.Done()
		for v := range chOut {
			var r rune
			err := v.Unpack(&r)
			if err != nil {
				panic(err)
			}
			res = append(res, r)
		}
	}()

	for _, v := range shuffled {
		chIn <- v
	}
	close(chIn)
	wg.Wait()
	assert.Equal(t, txt, res)
}
