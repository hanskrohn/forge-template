package directories

import (
	"errors"
	"strings"
	"unicode"
)

func BuildTree(input string) (*Node, error) {
	lines := strings.Split(input, "\n")
	root := Node{Value: cleanUpLine(lines[0]), depth: 0, isDir: true}
	_, err := helpBuildTree(&root, lines, 0, 0)

	return &root, err
}

func helpBuildTree(rootNode *Node, lines []string, index int, depth int) (int, error) {
	for nextLineIndex := index + 1; nextLineIndex < len(lines); nextLineIndex++ {
		nextLine := lines[nextLineIndex]
		if strings.TrimSpace(nextLine) == "" {
			continue
		}

		nextDepth := getNextDepth(nextLine)

		if nextDepth == 0 {
			return -1, errors.New("can only have one project root")
		}

		isDir, err := valueIsDir(nextLine)
		if err != nil {
			return -1, err
		}

		if nextDepth == depth+1 {
			cleanedUpLine := cleanUpLine(nextLine)
			newNode := Node{Value: cleanedUpLine, depth: nextDepth, isDir: isDir}
			
			rootNode.Children = append(rootNode.Children, &newNode)
		} else if nextDepth == depth+2 {
			nextLineIndex, err = helpBuildTree(rootNode.Children[len(rootNode.Children)-1], lines, nextLineIndex-1, rootNode.Children[len(rootNode.Children)-1].depth)
			if err != nil {
				return -1, err
			}
		} else if nextDepth > depth+2 {
			return nextLineIndex - 1, errors.New("invalid structure")
		} else {
			return nextLineIndex - 1, nil
		}
	}
	return len(lines), nil
}

func cleanUpLine(line string) string {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "|--")
	line = strings.TrimPrefix(line, "|-")

	return strings.TrimSpace(line)
}

func getNextDepth(line string) int {
    count := 0
    for _, ch := range line {
        if ch == ' ' {
            count++
        } else {
            break
        }
    }
    return count / 3
}

func valueIsDir(line string) (bool, error) {
	trimmedStringLeft := strings.TrimLeftFunc(line, unicode.IsSpace)

	var isDir bool = true
	if strings.HasPrefix(trimmedStringLeft, "|--") {
		isDir = true
	} else if strings.HasPrefix(trimmedStringLeft, "|-") {
		isDir = false
	} else {
		return isDir, errors.New("please specify if the value is a directory or file")
	}

	return isDir, nil
}