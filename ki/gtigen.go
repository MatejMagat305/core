// Code generated by "core generate"; DO NOT EDIT.

package ki

import (
	"cogentcore.org/core/gti"
)

// NodeType is the [gti.Type] for [Node]
var NodeType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/ki.Node", IDName: "node", Doc: "The Node implements the Ki interface and provides the core functionality\nfor the Cogent Core tree -- use the Node as an embedded struct or as a struct\nfield -- the embedded version supports full JSON save / load.\n\nThe desc: key for fields is used by the Cogent Core GUI viewer for help / tooltip\ninfo -- add these to all your derived struct's fields.  See relevant docs\nfor other such tags controlling a wide range of GUI and other functionality\n-- Ki makes extensive use of such tags.", Fields: []gti.Field{{Name: "Nm", Doc: "Nm is the user-supplied name of this node, which can be empty and/or non-unique."}, {Name: "Flags", Doc: "Flags are bit flags for internal node state, which can be extended using the enums package."}, {Name: "Props", Doc: "Props is a property map for arbitrary extensible properties."}, {Name: "Par", Doc: "Par is the parent of this node, which is set automatically when this node is added as a child of a parent."}, {Name: "Kids", Doc: "Kids is the list of children of this node. All of them are set to have this node\nas their parent. They can be reordered, but you should generally use Ki Node methods\nto Add / Delete to ensure proper usage."}, {Name: "Ths", Doc: "Ths is a pointer to ourselves as a Ki. It can always be used to extract the true underlying type\nof an object when [Node] is embedded in other structs; function receivers do not have this ability\nso this is necessary. This is set to nil when deleted. Typically use [Ki.This] convenience accessor\nwhich protects against concurrent access."}, {Name: "NumLifetimeKids", Doc: "NumLifetimeKids is the number of children that have ever been added to this node, which is used for automatic unique naming."}, {Name: "index", Doc: "index is the last value of our index, which is used as a starting point for finding us in our parent next time.\nIt is not guaranteed to be accurate; use the [Ki.IndexInParent] method."}, {Name: "depth", Doc: "depth is an optional depth parameter of this node, which is only valid during specific contexts, not generally.\nFor example, it is used in the WalkBreadth function"}}, Instance: &Node{}})

// NewNode adds a new [Node] with the given name to the given parent:
// The Node implements the Ki interface and provides the core functionality
// for the Cogent Core tree -- use the Node as an embedded struct or as a struct
// field -- the embedded version supports full JSON save / load.
//
// The desc: key for fields is used by the Cogent Core GUI viewer for help / tooltip
// info -- add these to all your derived struct's fields.  See relevant docs
// for other such tags controlling a wide range of GUI and other functionality
// -- Ki makes extensive use of such tags.
func NewNode(par Ki, name ...string) *Node {
	return par.NewChild(NodeType, name...).(*Node)
}

// KiType returns the [*gti.Type] of [Node]
func (t *Node) KiType() *gti.Type { return NodeType }

// New returns a new [*Node] value
func (t *Node) New() Ki { return &Node{} }
