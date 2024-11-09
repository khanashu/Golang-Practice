# Note Package

This `note` package provides functionality to create, save, and manage notes as JSON files. Each note includes a title, content, and a timestamp indicating when it was created. Notes are saved as JSON files with filenames derived from the note title.

## Features

- Create a note with a title and content.
- Save notes as JSON files.
- Automatically timestamp each note when it's created.

## Libraries and Functions Used

The package relies on several standard Go libraries for handling JSON encoding, error handling, time management, and file operations. Here’s an overview of each library used and how it's integrated:

1. **`encoding/json`**
   - Used to convert Go structs into JSON format, which allows notes to be saved in a structured and standardized format.
   - **Function used:** `json.Marshal` – Converts a `Note` struct into a JSON byte slice.

2. **`errors`**
   - Provides error handling for cases where required fields (title and content) are missing.
   - **Function used:** `errors.New` – Creates a new error message when either the title or content of a note is empty.

3. **`os`**
   - Handles file writing to save notes to the filesystem as JSON files.
   - **Function used:** `os.WriteFile` – Writes the JSON representation of a note to a file.

4. **`strings`**
   - Provides utility functions for manipulating the note's title to create a valid filename.
   - **Function used:** `strings.ReplaceAll` – Replaces spaces in the note title with underscores, ensuring filenames are space-free.
   - **Function used:** `strings.ToLower` – Converts the filename to lowercase.

5. **`time`**
   - Manages timestamps for notes, recording when each note is created.
   - **Function used:** `time.Now` – Captures the current time and assigns it to the note's `CreatedAt` field.

## Structs

### `Note`
Represents a note with the following fields:

- **`Title`** (`string`): The title of the note.
- **`Content`** (`string`): The content or body of the note.
- **`CreatedAt`** (`time.Time`): The timestamp for when the note was created.

## Functions

### `New(noteTitle, noteContent string) (Note, error)`

Creates a new note with the specified title and content.

- **Parameters:**
  - `noteTitle`: The title for the note. Cannot be empty.
  - `noteContent`: The content of the note. Cannot be empty.
- **Returns:**
  - `Note`: A `Note` struct populated with the provided title, content, and current timestamp.
  - `error`: Returns an error if the title or content is empty.

### `Save() error`

Saves the note as a JSON file with a filename derived from the title. Spaces are replaced with underscores, and the filename is converted to lowercase.

- **Returns:** `error` – If any error occurs during the save operation, it is returned.

## Usage Example

```go
package main

import (
	"fmt"
	"note"
)

func main() {
	note, err := note.New("Meeting Notes", "Discuss project milestones and deadlines.")
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	err = note.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
	} else {
		fmt.Println("Note saved successfully!")
	}
}
```

This example creates a new note with the title "Meeting Notes" and content about project milestones. It then saves the note to the current directory as a JSON file named `meeting_notes.json`.

## Struct Tags for Metadata

This code doesn’t currently implement JSON struct tags, but struct tags can be added to customize JSON keys. For example, to lowercase struct fields in JSON files:

```go
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
```

This will change the JSON keys to `title`, `content`, and `created_at` respectively.