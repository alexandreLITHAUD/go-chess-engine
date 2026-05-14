package board

import (
	"fmt"
	"strings"
)

// Print renders the board with corrected alignment and looping labels.
func (b *Board) Print() {
	height := len(b.Squares)
	if height == 0 {
		fmt.Println("Board is empty.")
		return
	}
	width := len(b.Squares[0])

	// 1. Configuration
	const cellWidth = 3 // Width of the content inside the pipes (e.g., " K ")

	// Determine the padding needed for the rank numbers (e.g., "100" needs more than "8")
	rankPadding := len(fmt.Sprintf("%d", height))

	// Create the horizontal separator line
	// We use "+" at intersections for better visual structure
	// Format: [Spaces for Rank] + [---+] repeated
	padding := strings.Repeat(" ", rankPadding+1)
	lineSeparator := padding + "+" + strings.Repeat(strings.Repeat("-", cellWidth)+"+", width)

	fmt.Println(lineSeparator)

	for i := range height {
		// 2. Print Rank number with dynamic padding to keep columns straight
		rankNum := height - i
		fmt.Printf("%*d |", rankPadding, rankNum)

		for j := range width {
			piece := b.Squares[i][j]
			var displayChar string

			if piece == nil {
				displayChar = " " // Changed from "." to " " for a cleaner look
			} else {
				displayChar = piece.CharRepresentation
				if strings.ToLower(piece.Color) == "black" {
					displayChar = strings.ToLower(displayChar)
				} else {
					displayChar = strings.ToUpper(displayChar)
				}
			}

			// Center the character within the cellWidth
			// This ensures " r " stays aligned with "---"
			fmt.Printf(" %s |", displayChar)
		}

		fmt.Println()
		fmt.Println(lineSeparator)
	}

	// 3. Print File labels with looping logic (a-z, then A-Z)
	fmt.Print(strings.Repeat(" ", rankPadding+3)) // Align with the first cell
	for j := range width {
		fmt.Printf("%s   ", getFileLabel(j))
	}
	fmt.Println()
}

// getFileLabel handles the looping from a-z to A-Z
func getFileLabel(index int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Loop back to 'a' if we exceed 52 columns
	return string(alphabet[index%len(alphabet)])
}

// Print renders the board with strict horizontal alignment.
// func (b *Board) Print() {
// 	height := len(b.Squares)
// 	if height == 0 {
// 		fmt.Println("Board is empty.")
// 		return
// 	}
// 	width := len(b.Squares[0])
//
// 	// Define cell width (e.g., 3 characters for " K ", " . ", " B ")
// 	const cellWidth = 3
// 	// Pre-calculate the horizontal separator line
// 	// (width of cells + width of pipes)
// 	lineSeparator := "   " + strings.Repeat(strings.Repeat("-", cellWidth)+"+", width)
//
// 	fmt.Println(lineSeparator)
//
// 	for i := range height {
// 		// Print Rank number with padding
// 		fmt.Printf("%2d |", height-i)
//
// 		for j := range width {
// 			piece := b.Squares[i][j]
// 			var displayChar string
//
// 			if piece == nil {
// 				displayChar = " "
// 			} else {
// 				displayChar = piece.CharRepresentation
// 				// Casing logic: White Upper, Black Lower
// 				if strings.ToLower(piece.Color) == "black" {
// 					displayChar = strings.ToLower(displayChar)
// 				} else {
// 					displayChar = strings.ToUpper(displayChar)
// 				}
// 			}
//
// 			// %-*s aligns the string to the left with a fixed width
// 			// This keeps the pipes '|' perfectly vertical
// 			fmt.Printf(" %-*s|", cellWidth-1, displayChar)
// 		}
//
// 		fmt.Println()
// 		fmt.Println(lineSeparator)
// 	}
//
// 	// Print File labels (a, b, c...)
// 	fmt.Print("    ") // Initial padding for rank numbers
// 	for j := range width {
// 		// Center the letter under the cell
// 		label := fmt.Sprintf("%c", 'a'+j)
// 		fmt.Printf(" %-*s ", cellWidth-1, label)
// 	}
//
// }
