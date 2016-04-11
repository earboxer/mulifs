// Copyright 2016 Danko Miocevic. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Danko Miocevic

// Package playlistmgr contains all the tools to read and modify
// playlists files.
package playlistmgr

import (
	"errors"
	"os"
)

// FileTags defines the tags found in a specific music file.
type PlaylistFile struct {
	Title  string
	Artist string
	Album  string
	Path   string
}

/** CheckPlaylistFile opens a Playlist file and checks that
 * it really is a valid Playlist.
 */
func CheckPlaylistFile(path string) error {
	f, err := os.Open(path)

	if err != nil {
		return err
	}

	firstLine := make([]byte, len("#EXTM3U"))
	_, err = f.Read(firstLine)
	if err != nil {
		return err
	}

	if string(firstLine) != "#EXTM3U" {
		return errors.New("Not a playlist!")
	}

	return nil
}

// ProcessPlaylist function receives the path of a playlist
// and adds all the information into the database.
// It process every line in the file and reads all the
// songs in it.
func ProcessPlaylist(path string) error {
	return nil
}

/** AddPlaylistSong receives an item from a playlist and
 * adds it to the playlist Bucket in the database.
 * It also checks that the song exists in the database
 * and adds the playlist link to the specific song.
 */
func AddPlaylistSong(playlist, artist, album, song string) error {
	return nil
}

/** AddPlaylistFile iterates throught the lines of the playlist
 * and finds MuLi entries to get the playlist items.
 */
func AddPlaylistFile(path string) error {
	return nil
}

/** DeletePlaylist deletes a playlist file.
 * It also deletes the contents on the database and deletes
 * the playlist entries on each individual file.
 */
func DeletePlaylist(playlist string) error {
	return nil
}

/** DeletePlaylistArtist iterates through an artist in the playlist
 * and deletes all the contents.
 * It also deletes the contents on the database and deletes
 * the playlist entries on each individual file.
 */
func DeletePlaylistArtist(playlist, artist string) error {
	return nil
}

/** DeletePlaylistArtist iterates through a specific album from a
 * specific artist in the playlist and deletes all the contents.
 * It also deletes the contents on the database and deletes
 * the playlist entries on each individual file.
 */
func DeletePlaylistAlbum(playlist, artist, album string) error {
	return nil
}

/** RegeneratePlaylistFile creates the playlist file from the
 * information in the database.
 */
func RegeneratePlaylistFile(playlist string) error {
	return nil
}
