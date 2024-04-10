// Code generated by "core generate"; DO NOT EDIT.

package tree

import (
	"cogentcore.org/core/gti"
)

// NodeBaseType is the [gti.Type] for [NodeBase]
var NodeBaseType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/tree.NodeBase", IDName: "node-base", Doc: "NodeBase implements the [Node] interface and provides the core functionality\nfor the Cogent Core tree system. You should use NodeBase as an embedded struct\nin higher-level tree types.", Fields: []gti.Field{{Name: "Nm", Doc: "Nm is the user-supplied name of this node, which can be empty and/or non-unique.\nIt is typically accessed through [Node.Name]."}, {Name: "Flags", Doc: "Flags are bit flags for internal node state, which can be extended using\nthe enums package."}, {Name: "Props", Doc: "Props is a property map for arbitrary key-value properties.\nThey are typically accessed through the property methods on [Node]."}, {Name: "Par", Doc: "Par is the parent of this node, which is set automatically when this node is\nadded as a child of a parent. It is typically accessed through [Node.Parent]."}, {Name: "Kids", Doc: "Kids is the list of children of this node. All of them are set to have this node\nas their parent. They can be reordered, but you should generally use [Node]\nmethods when adding and deleting children to ensure everything gets updated.\nThey are typically accessed through [Node.Children]."}, {Name: "this", Doc: "this is a pointer to ourselves as a [Node]. It can always be used to extract the\ntrue underlying type of an object when [NodeBase] is embedded in other structs;\nfunction receivers do not have this ability, so this is necessary. This is set\nto nil when the node is deleted. It is typically accessed through [Node.This]."}, {Name: "numLifetimeChildren", Doc: "numLifetimeChildren is the number of children that have ever been added to this\nnode, which is used for automatic unique naming. It is typically accessed\nthrough [Node.NumLifetimeChildren]."}, {Name: "index", Doc: "index is the last value of our index, which is used as a starting point for\nfinding us in our parent next time. It is not guaranteed to be accurate;\nuse the [Node.IndexInParent] method."}, {Name: "depth", Doc: "depth is the depth of the node while using [Node.WalkDownBreadth]."}}, Instance: &NodeBase{}})

// NewNodeBase adds a new [NodeBase] with the given name to the given parent:
// NodeBase implements the [Node] interface and provides the core functionality
// for the Cogent Core tree system. You should use NodeBase as an embedded struct
// in higher-level tree types.
func NewNodeBase(parent Node, name ...string) *NodeBase {
	return parent.NewChild(NodeBaseType, name...).(*NodeBase)
}

// NodeType returns the [*gti.Type] of [NodeBase]
func (t *NodeBase) NodeType() *gti.Type { return NodeBaseType }

// New returns a new [*NodeBase] value
func (t *NodeBase) New() Node { return &NodeBase{} }
