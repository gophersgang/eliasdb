/*
 * Public Domain Software
 *
 * I (Matthias Ladkau) am the author of the source code in this file.
 * I have placed the source code in this file in the public domain. 
 *
 * For further information see: http://creativecommons.org/publicdomain/zero/1.0/
 */

/*
A persistent map storing binary values.
*/
package datautil

import (
	"encoding/gob"
	"os"
)

/*
Persistent map datastructure.
*/
type PersistentMap struct {
	filename string            // File of the persistent map
	Data     map[string]string // Data of the persistent map
}

/*
Create a new persistent map.
*/
func NewPersistentMap(filename string) (*PersistentMap, error) {
	pm := &PersistentMap{filename, make(map[string]string)}
	return pm, pm.Flush()
}

/*
Load persistent map from a file.
*/
func LoadPersistentMap(filename string) (*PersistentMap, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return nil, err
	}

	pm := &PersistentMap{filename, make(map[string]string)}

	de := gob.NewDecoder(file)
	de.Decode(&pm.Data)

	return pm, file.Close()
}

/*
Write contents of the persistent map to disk.
*/
func (pm *PersistentMap) Flush() error {
	file, err := os.OpenFile(pm.filename, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return err
	}

	en := gob.NewEncoder(file)
	en.Encode(pm.Data)

	return file.Close()
}