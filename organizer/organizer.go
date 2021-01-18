// package organizer tries to introduce a 'generic' method that will allow to
// reorder any kind of elements that are returning back from a multiple
// parallel workers. A system of this sort should be able to rebuild order
// without keepint too many objects in memory and able to work with millions of
// elements.
package organizer

import (
	"context"
	"fmt"
)

// Organize takes disordered data from the chIn channel and resends them in
// the original order to chOut channel.
func Organize(
	ctx context.Context,
	chIn <-chan Ordered,
	chOut chan<- Ordered,
) error {
	var currIndex, count int
	tempStorage := make(map[int]Ordered)
	for o := range chIn {
		count++
		if o.Index() == currIndex {
			select {
			case <-ctx.Done():
				return nil
			case chOut <- o:
			}
			currIndex++
			continue
		}

		tempStorage[o.Index()] = o

		if oMap, ok := tempStorage[currIndex]; ok {
			select {
			case <-ctx.Done():
				return nil
			case chOut <- oMap:
			}
			delete(tempStorage, currIndex)
			currIndex++
			continue
		}
	}
	for currIndex < count {
		if oMap, ok := tempStorage[currIndex]; ok {
			select {
			case <-ctx.Done():
				return nil
			case chOut <- oMap:
			}
			delete(tempStorage, currIndex)
		}
		currIndex++
	}
	close(chOut)
	if len(tempStorage) > 0 {
		return fmt.Errorf("Could not assemble all elements, %d of them left",
			len(tempStorage))
	}
	return nil
}
