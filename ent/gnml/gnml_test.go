package gnml_test

import (
	"testing"

	"github.com/gnames/gnlib"
	"github.com/gnames/gnlib/ent/gnml"
	"github.com/stretchr/testify/assert"
)

var txt = `
Foreword 

Shell collecting is now taking its place as one of the major outdoor 
diversions. It has advantages over such pursuits as bird watching or fishing, 
for you may have even more pleasure in studying your catch at home than 
in the time spent afield. The thrill of finding a shell new to you, or of watch- 
ing some rare snail going about its watery aff^airs, is ample reward for the 
sunburn and stifi" neck you may have from wading around too long with a 
water-glass. Hours sieving dredgings are counted well spent if a fine volute 
or turrid turns up in the seaweed and rubbish. 

American Seashells gives a comprehensive and well-rounded view of the 
Mollusca in nontechnical language. It is easy reading for the beginner, but it 
contains also material indispensable to the advanced malacologist. The chap- 
ters on nudibranchs and pteropods are especially welcome, for these beautiful 
animals have always been slighted in American books. In chapters on the 
ife of the snail and the clam, with the author we "listen in" to the current 
of molluscan life. The shells become living things, moving and breathing, 
feeding and mating. 

One perplexity of the novice is that different books may give different 
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
cultivated the society of mollusks in many lands, from East Africa, the 

xiii 



Foreivord 

XIV 



A.AV 

of all of us who hunt the elusive moUusk. 



Henry A. Pilsbry 
Curator of Mollusks 
Academy of Natural Sciences 
of Philadelphia 



PART I 



The Natural History 
of Seashells 



CHAPTER I 

Man and Mollusks 



Seashells and man were closely associated even before the dawn of civiliza- 
tion when primitive man gathered snails, oysters, and other kinds of mollusks 
along the seashore for food, implements, ornaments, and money. The many 
kitchen-middens and burial sites in nearly every comer of the world reveal 
the great extent to which early peoples were dependent upon mollusks. On 
some coral islands, as, for instance, Barbados, where there was no available 
stone, nearly all domestic utensils, including knives and axes, were made from 
seashells. As civilization became more complex, specialization in the use of 
mollusks increased. From them were obtained dyes, inks, textiles and win- 
dowpanes. In the Mediterranean region there was a long period when an 
entire commercial empire owed its origin and continued success to the Tyrian 
purple obtained from a seashell. Later, in Roman times, the farming of 
oysters and edible snails became a major enterprise. 

Today the uses of moUuscan shells are legion. Jewelers, artists and but- 
ton manufacturers; biologists, geologists and archaeologists; bird and aquarium 
dealers; all daily use mollusks or their products. In recent years there has 
flourished m Florida a five-million-dollar-a-year seashell industry. Through- 
out the country, the hobby of shell collecting is enjoyed by countless thou- 
sands, and it now rivals the popularity of coin collecting. Local and federal 
agencies arc investing millions in research directed toward the more efficient 
cultivation and utilization of commercially important mollusks. 

From another standpoint of perhaps even greater importance mollusks 
have influenced the activities and welfare of man. Some are extremely de- 
structive to wooden structures in the sea, and others are a serious menace to 
health, mostly as intermediate hosts to dangerous parasites or as carriers of 

3 
`

func TestSplitText(t *testing.T) {
	assert := assert.New(t)
	res := gnml.SplitText(txt, 1000, 100)
	ls := gnlib.Map(res, func(s string) int {
		return len(s)
	})
	assert.Equal([]int{594, 656, 938, 871, 691, 715, 406}, ls)
}
