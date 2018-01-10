// Copyright 2018 Mark Wardle / Eldrix Ltd
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package snomed

import (
	"strconv"

	"github.com/wardle/go-terminology/verhoeff"
)

// Identifier (SCTID) is a checksummed (Verhoeff) globally unique persistent identifier
// See https://confluence.ihtsdotools.org/display/DOCTIG/3.1.4.2.+Component+features+-+Identifiers
// The SCTID data type is 64-bit Integer which is allocated and represented in accordance with a set of rules.
// These rules enable each Identifier to refer unambiguously to a unique component.
// They also support separate partitions for allocation of Identifiers for particular types of component and
// namespaces that distinguish between different issuing organizations.
type Identifier int

// AsInteger is a convenience method to convert to integer
func (id Identifier) AsInteger() int {
	return int(id)
}

// IsConcept will return true if this identifier refers to a concept
// TODO: add implementation
func (id Identifier) IsConcept() bool {
	pid := id.partitionIdentifier()
	return pid == "00" || pid == "10"
}

// IsDescription will return true if this identifier refers to a description.
// TODO: add implementation
func (id Identifier) IsDescription() bool {
	pid := id.partitionIdentifier()
	return pid == "01" || pid == "11"
}

// IsRelationship will return true if this identifier refers to a relationship.
// TODO: add implementation
func (id Identifier) isRelationship() bool {
	pid := id.partitionIdentifier()
	return pid == "02" || pid == "12"
}

// IsValid will return true if this is a valid SNOMED CT identifier
func (id Identifier) IsValid() bool {
	s := strconv.Itoa(int(id))
	return verhoeff.ValidateVerhoeffString(s)
}

// partitionIdentifier returns the penultimate last digit digits
// see https://confluence.ihtsdotools.org/display/DOCRELFMT/5.5.+Partition+Identifier
// 0123456789
// xxxxxxxppc
func (id Identifier) partitionIdentifier() string {
	s := strconv.Itoa(int(id))
	l := len(s)
	return s[l-3 : l-1]
}
