/*
 *    This file is part of Disgo-Commons library.
 *
 *    The Disgo-Commons library is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    The Disgo-Commons library is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with the Disgo-Commons library.  If not, see <http://www.gnu.org/licenses/>.
 */
package tree

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
)

//MerkleTreeContent represents the data that is stored and verified by the tree. A type that
//implements this interface can be used as an item in the tree.
type MerkleTreeContent interface {
	CalculateHash() []byte
	Equals(other MerkleTreeContent) bool
}

//MerkleTree is the container for the tree. It holds a pointer to the root of the tree,
//a list of pointers to the leaf nodes, and the merkle root.
type MerkleTree struct {
	Root       *Node
	merkleRoot []byte
	Leafs      []*Node
}

//Node represents a node, root, or leaf in the tree. It stores pointers to its immediate
//relationships, a hash, the content stored if it is a leaf, and other metadata.
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	leaf   bool
	dup    bool
	Hash   []byte
	C      MerkleTreeContent
}

//verifyNode walks down the tree until hitting a leaf, calculating the hash at each level
//and returning the resulting hash of Node n.
func (n *Node) verifyNode() []byte {
	if n.leaf {
		return n.C.CalculateHash()
	}
	h := sha256.New()
	h.Write(append(n.Left.verifyNode(), n.Right.verifyNode()...))
	return h.Sum(nil)
}

//calculateNodeHash is a helper function that calculates the hash of the node.
func (n *Node) calculateNodeHash() []byte {
	if n.leaf {
		return n.C.CalculateHash()
	}
	h := sha256.New()
	h.Write(append(n.Left.Hash, n.Right.Hash...))
	return h.Sum(nil)
}

//NewTree creates a new Merkle Tree using the content cs.
func NewTree(cs []MerkleTreeContent) (*MerkleTree, error) {
	root, leafs, err := buildWithContent(cs)
	if err != nil {
		return nil, err
	}
	t := &MerkleTree{
		Root:       root,
		merkleRoot: root.Hash,
		Leafs:      leafs,
	}
	return t, nil
}

//buildWithContent is a helper function that for a given set of Contents, generates a
//corresponding tree and returns the root node, a list of leaf nodes, and a possible error.
//Returns an error if cs contains no Contents.
func buildWithContent(cs []MerkleTreeContent) (*Node, []*Node, error) {
	if len(cs) == 0 {
		return nil, nil, errors.New("Error: cannot construct tree with no content.")
	}
	var leafs []*Node
	for _, c := range cs {
		leafs = append(leafs, &Node{
			Hash: c.CalculateHash(),
			C:    c,
			leaf: true,
		})
	}
	if len(leafs)%2 == 1 {
		duplicate := &Node{
			Hash: leafs[len(leafs)-1].Hash,
			C:    leafs[len(leafs)-1].C,
			leaf: true,
			dup:  true,
		}
		leafs = append(leafs, duplicate)
	}
	root := buildIntermediate(leafs)
	return root, leafs, nil
}

//buildIntermediate is a helper function that for a given list of leaf nodes, constructs
//the intermediate and root levels of the tree. Returns the resulting root node of the tree.
func buildIntermediate(nl []*Node) *Node {
	var nodes []*Node
	for i := 0; i < len(nl); i += 2 {
		h := sha256.New()
		var left, right int = i, i + 1
		if i+1 == len(nl) {
			right = i
		}
		chash := append(nl[left].Hash, nl[right].Hash...)
		h.Write(chash)
		n := &Node{
			Left:  nl[left],
			Right: nl[right],
			Hash:  h.Sum(nil),
		}
		nodes = append(nodes, n)
		nl[left].Parent = n
		nl[right].Parent = n
		if len(nl) == 2 {
			return n
		}
	}
	return buildIntermediate(nodes)
}

//MerkleRoot returns the unverified Merkle Root (hash of the root node) of the tree.
func (m *MerkleTree) MerkleRoot() []byte {
	return m.merkleRoot
}

//RebuildTree is a helper function that will rebuild the tree reusing only the content that
//it holds in the leaves.
func (m *MerkleTree) RebuildTree() error {
	var cs []MerkleTreeContent
	for _, c := range m.Leafs {
		cs = append(cs, c.C)
	}
	root, leafs, err := buildWithContent(cs)
	if err != nil {
		return err
	}
	m.Root = root
	m.Leafs = leafs
	m.merkleRoot = root.Hash
	return nil
}

//RebuildTreeWith replaces the content of the tree and does a complete rebuild; while the root of
//the tree will be replaced the MerkleTree completely survives this operation. Returns an error if the
//list of content cs contains no entries.
func (m *MerkleTree) RebuildTreeWith(cs []MerkleTreeContent) error {
	root, leafs, err := buildWithContent(cs)
	if err != nil {
		return err
	}
	m.Root = root
	m.Leafs = leafs
	m.merkleRoot = root.Hash
	return nil
}

//VerifyTree verify tree validates the hashes at each level of the tree and returns true if the
//resulting hash at the root of the tree matches the resulting root hash; returns false otherwise.
func (m *MerkleTree) VerifyTree() bool {
	calculatedMerkleRoot := m.Root.verifyNode()
	if bytes.Compare(m.merkleRoot, calculatedMerkleRoot) == 0 {
		return true
	}
	return false
}

//VerifyContent indicates whether a given content is in the tree and the hashes are valid for that content.
//Returns true if the expected Merkle Root is equivalent to the Merkle root calculated on the critical path
//for a given content. Returns true if valid and false otherwise.
func (m *MerkleTree) VerifyContent(expectedMerkleRoot []byte, content MerkleTreeContent) bool {
	for _, l := range m.Leafs {
		if l.C.Equals(content) {
			currentParent := l.Parent
			for currentParent != nil {
				h := sha256.New()
				if currentParent.Left.leaf && currentParent.Right.leaf {
					h.Write(append(currentParent.Left.calculateNodeHash(), currentParent.Right.calculateNodeHash()...))
					if bytes.Compare(h.Sum(nil), currentParent.Hash) != 0 {
						return false
					}
					currentParent = currentParent.Parent
				} else {
					h.Write(append(currentParent.Left.calculateNodeHash(), currentParent.Right.calculateNodeHash()...))
					if bytes.Compare(h.Sum(nil), currentParent.Hash) != 0 {
						return false
					}
					currentParent = currentParent.Parent
				}
			}
			return true
		}
	}
	return false
}

//String returns a string representation of the tree. Only leaf nodes are included
//in the output.
func (m *MerkleTree) String() string {
	s := ""
	for _, l := range m.Leafs {
		s += fmt.Sprint(l)
		s += "\n"
	}
	return s
}
