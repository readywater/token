package token

import (
	"time"
	"sync"
	// "github.com/coopernurse/gorp"
)

type TokenEntry struct {
	Id int64		`id`
	Tokenid int64	`tokenid`
	Timestamp int64	`timestamp`
	Value float32	`value`
}

type Token struct {
	Id int64				`id`
	Userid int64			`userid`
	Title string			`title`
	Quantity int64 			`quantity`
	Interval string 		`interval`
	Start int64 			`start`
	tokens []*TokenEntry	`db:"-"`
	mutex *sync.Mutex		`db:"-"`
}

/////// ENTRIES ///////
// Make Entry
func (t *Token) UseToken() int64 {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.IsTokenAvailable() { 
		return -1
	} else {
		newId := int64(len(t.tokens))
		currentTime := time.Now().Unix()

		entry := &TokenEntry{
			newId,
			t.Userid,
			currentTime,
			1.0,
		}

		t.tokens = append(t.tokens,entry)
		return 1
	}
}

// Get All Token Entries
func (t *Token) GetAllTokenEntries() []*TokenEntry{
	return nil
}

// Get most recent entry 
func (t *Token) GetMostRecentEntry() *TokenEntry {
	return t.tokens[len(t.tokens)]
}

// Get entry (id)
func (t *Token) GetEntry(entryid int64) *TokenEntry {
	if(entryid < 0 || entryid >= int64(len(t.tokens))) { return nil }
	return t.tokens[len(t.tokens)-1]
}

// Delete all Entries
func (t *Token) RemoveAllEntries() int {
	// Does this work?
	t.tokens = make([]*TokenEntry, 0)
	return 0
}

// Delete Last Entry (Undo)
func (t *Token) RemoveLastEntry() int {

	return 0
}

// Delete an entry (id)
func (t *Token) RemoveEntry(entryid int32) int {
	return 0
}