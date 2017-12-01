/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 01-12-2017
 * |
 * | File Name:     types.go
 * +===============================================
 */

package wikimedia

// Request represents query string parameters of wiki media api.
type Request struct {
	Action        string
	Generator     string
	Titles        string
	Prop          string
	Iiprop        string
	Format        string
	FormatVersion int
}
