package custom

import (
	"rutasMap/v2/models"
)

type NodeSortingIndex []models.Point

func (n NodeSortingIndex) Less(i, j int) bool {

	return n[i].IndexNext == n[j].Index
}

func (n NodeSortingIndex) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NodeSortingIndex) Len() int {
	return len(n)

}
