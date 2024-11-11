// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/groth16.proto

package groth16

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClientState defines a groth16 client that holds two keys for verifying state
// transition proofs needed to verify IBC packets
type ClientState struct {
	// latest height of the client state
	LatestHeight uint64 `protobuf:"varint,1,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height,omitempty"`
	// groth16 state transition proof verifier key. Verifies proofs on a rollups
	// state root after the state transition has been applied. Only BN254 curve is
	// supported
	StateTransitionVerifierKey []byte `protobuf:"bytes,2,opt,name=state_transition_verifier_key,json=stateTransitionVerifierKey,proto3" json:"state_transition_verifier_key,omitempty" yaml:"stp_verifier_key"`
	// Provided during initialization of the IBC Client
	CodeCommitment []byte `protobuf:"bytes,3,opt,name=code_commitment,json=codeCommitment,proto3" json:"code_commitment,omitempty"`
	// Provided during initialization of the IBC Client
	GenesisStateRoot []byte `protobuf:"bytes,4,opt,name=genesis_state_root,json=genesisStateRoot,proto3" json:"genesis_state_root,omitempty"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e462ca8f61a2a26, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

// ConsensusState defines a groth16 consensus state.
type ConsensusState struct {
	// timestamp that corresponds to the block height in which the ConsensusState
	// was stored.
	Timestamp time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// state root of the rollup
	StateRoot []byte `protobuf:"bytes,2,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty" yaml:"state_root"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e462ca8f61a2a26, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

// Header defines a groth16 header for updating the trusted state root of a
// rollup
type Header struct {
	// serialized groth16 proof that the given state transition is valid
	StateTransitionProof []byte `protobuf:"bytes,1,opt,name=state_transition_proof,json=stateTransitionProof,proto3" json:"state_transition_proof,omitempty" yaml:"state_transition_proof"`
	// last verified height of the rollup. This is used to retrieve the previous
	// state root with which the proof is verified against
	TrustedHeight int64 `protobuf:"varint,2,opt,name=trusted_height,json=trustedHeight,proto3" json:"trusted_height,omitempty" yaml:"trusted_height"`
	// trusted header hash passed by the relayer
	TrustedCelestiaHeaderHash []byte `protobuf:"bytes,3,opt,name=trusted_celestia_header_hash,json=trustedCelestiaHeaderHash,proto3" json:"trusted_celestia_header_hash,omitempty" yaml:"trusted_celestia_header_hash"`
	// new state root, height and header hash of the rollup after the state transition has been
	// applied
	NewStateRoot          []byte `protobuf:"bytes,4,opt,name=new_state_root,json=newStateRoot,proto3" json:"new_state_root,omitempty" yaml:"new_state_root"`
	NewHeight             int64  `protobuf:"varint,5,opt,name=new_height,json=newHeight,proto3" json:"new_height,omitempty" yaml:"new_height"`
	NewCelestiaHeaderHash []byte `protobuf:"bytes,6,opt,name=new_celestia_header_hash,json=newCelestiaHeaderHash,proto3" json:"new_celestia_header_hash,omitempty" yaml:"new_celestia_header_hash"`
	// TODO: This is provided by the user at the moment but we can't trust them
	// with this data. We need to get all the data roots from the
	// the store.
	DataRoots [][]byte `protobuf:"bytes,7,rep,name=data_roots,json=dataRoots,proto3" json:"data_roots,omitempty" yaml:"data_roots"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e462ca8f61a2a26, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "celestia.ibc.lightclients.groth16.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "celestia.ibc.lightclients.groth16.v1.ConsensusState")
	proto.RegisterType((*Header)(nil), "celestia.ibc.lightclients.groth16.v1.Header")
}

func init() { proto.RegisterFile("proto/groth16.proto", fileDescriptor_7e462ca8f61a2a26) }

var fileDescriptor_7e462ca8f61a2a26 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0x8c, 0x69, 0x28, 0x64, 0x9b, 0x06, 0x30, 0x2d, 0xb8, 0x81, 0xda, 0xc5, 0x45, 0x6a, 0x0f,
	0xd4, 0x56, 0xa1, 0xe2, 0xd0, 0x0b, 0xc8, 0x39, 0x50, 0x89, 0x0b, 0x5a, 0x2a, 0x90, 0x10, 0xc2,
	0x5a, 0x3b, 0x5f, 0xed, 0x55, 0x6d, 0x6f, 0xe4, 0xdd, 0x34, 0x2a, 0x4f, 0xc0, 0xb1, 0x4f, 0x80,
	0x78, 0x9c, 0x1e, 0x7b, 0xe4, 0x64, 0x50, 0xfb, 0x00, 0x48, 0x7e, 0x02, 0xe4, 0xb5, 0x9d, 0x9f,
	0x36, 0xb7, 0xec, 0x37, 0xb3, 0x93, 0x99, 0xf9, 0x56, 0x46, 0x0f, 0x07, 0x29, 0x13, 0xcc, 0x0e,
	0x52, 0x26, 0xc2, 0xdd, 0xd7, 0x96, 0x3c, 0xa9, 0xcf, 0x7d, 0x88, 0x80, 0x0b, 0x4a, 0x2c, 0xea,
	0xf9, 0x56, 0x44, 0x83, 0x50, 0xf8, 0x11, 0x85, 0x44, 0x70, 0xab, 0x26, 0x9e, 0xec, 0x76, 0x57,
	0x02, 0x16, 0xb0, 0xea, 0x3a, 0x0b, 0x58, 0x79, 0xb7, 0x6b, 0x04, 0x8c, 0x05, 0x11, 0xd8, 0xf2,
	0xe4, 0x0d, 0x8f, 0x6c, 0x41, 0x63, 0xe0, 0x82, 0xc4, 0x83, 0x92, 0x60, 0xfe, 0x53, 0xd0, 0x52,
	0x4f, 0xaa, 0x7d, 0x14, 0x44, 0x80, 0xba, 0x89, 0x96, 0x23, 0x22, 0x80, 0x0b, 0x37, 0x84, 0xe2,
	0x9f, 0x34, 0x65, 0x43, 0xd9, 0x6e, 0xe2, 0x76, 0x39, 0x3c, 0x90, 0x33, 0xf5, 0x1b, 0x5a, 0xe7,
	0x05, 0xdb, 0x15, 0x29, 0x49, 0x38, 0x15, 0x94, 0x25, 0xee, 0x09, 0xa4, 0xf4, 0x88, 0x42, 0xea,
	0x1e, 0xc3, 0xa9, 0x76, 0x6b, 0x43, 0xd9, 0x6e, 0x3b, 0x4f, 0xf2, 0xcc, 0x78, 0x7c, 0x4a, 0xe2,
	0x68, 0xdf, 0xe4, 0x62, 0x30, 0xc3, 0x30, 0x71, 0x57, 0x2a, 0x1c, 0x8e, 0x05, 0x3e, 0x55, 0xe8,
	0x7b, 0x38, 0x55, 0xb7, 0xd0, 0x3d, 0x9f, 0xf5, 0xc1, 0xf5, 0x59, 0x1c, 0x53, 0x11, 0x43, 0x22,
	0xb4, 0x85, 0x42, 0x11, 0x77, 0x8a, 0x71, 0x6f, 0x3c, 0x55, 0x5f, 0x20, 0x35, 0x80, 0x04, 0x38,
	0xe5, 0x6e, 0x69, 0x28, 0x65, 0x4c, 0x68, 0x4d, 0xc9, 0xbd, 0x5f, 0x21, 0x32, 0x17, 0x66, 0x4c,
	0xec, 0x37, 0x7f, 0xfc, 0x32, 0x1a, 0xe6, 0x99, 0x82, 0x3a, 0x3d, 0x96, 0x70, 0x48, 0xf8, 0xb0,
	0x04, 0x55, 0x07, 0xb5, 0xc6, 0xbd, 0xc8, 0xc0, 0x4b, 0x2f, 0xbb, 0x56, 0xd9, 0x9c, 0x55, 0x37,
	0x67, 0x1d, 0xd6, 0x0c, 0xe7, 0xee, 0x79, 0x66, 0x34, 0xce, 0xfe, 0x18, 0x0a, 0x9e, 0x5c, 0x53,
	0xf7, 0x10, 0x9a, 0xb2, 0x50, 0x16, 0xb0, 0x9a, 0x67, 0xc6, 0x83, 0xba, 0x80, 0x1a, 0x33, 0x71,
	0x8b, 0x5f, 0xb3, 0xf4, 0xb3, 0x89, 0x16, 0x0f, 0x80, 0xf4, 0x21, 0x55, 0x3f, 0xa3, 0x47, 0x37,
	0xaa, 0x1d, 0xa4, 0x8c, 0x1d, 0x49, 0x5f, 0x6d, 0xe7, 0x59, 0x9e, 0x19, 0xeb, 0xd3, 0x92, 0xd7,
	0x79, 0x26, 0x5e, 0xb9, 0xd6, 0xec, 0x87, 0x62, 0xac, 0xbe, 0x45, 0x1d, 0x91, 0x0e, 0xb9, 0x80,
	0x7e, 0xbd, 0xd9, 0xc2, 0xe3, 0x82, 0xb3, 0x96, 0x67, 0xc6, 0x6a, 0x29, 0x38, 0x8b, 0x9b, 0x78,
	0xb9, 0x1a, 0x54, 0x5b, 0x0f, 0xd1, 0xd3, 0x9a, 0x51, 0xbf, 0x48, 0x37, 0x94, 0xae, 0xdd, 0x90,
	0xf0, 0xb0, 0x5c, 0x91, 0xb3, 0x95, 0x67, 0xc6, 0xe6, 0xac, 0xde, 0x3c, 0xb6, 0x89, 0xd7, 0x2a,
	0xb8, 0x57, 0xa1, 0x65, 0x01, 0x07, 0x84, 0x87, 0xea, 0x1b, 0xd4, 0x49, 0x60, 0x74, 0x63, 0xa5,
	0xd3, 0x5e, 0x67, 0x71, 0x13, 0xb7, 0x13, 0x18, 0x8d, 0x37, 0x5d, 0x2c, 0xa3, 0x20, 0x54, 0x41,
	0x6f, 0xcb, 0xa0, 0x53, 0xcb, 0x98, 0x60, 0x26, 0x6e, 0x25, 0x30, 0xaa, 0x02, 0x7e, 0x45, 0x5a,
	0x81, 0xcc, 0x0d, 0xb7, 0x28, 0x0d, 0x6c, 0xe6, 0x99, 0x61, 0x4c, 0x34, 0xe6, 0x07, 0x5b, 0x4d,
	0x60, 0x34, 0x27, 0xd4, 0x1e, 0x42, 0x7d, 0x22, 0x88, 0xf4, 0xcb, 0xb5, 0x3b, 0x1b, 0x0b, 0xb3,
	0x0f, 0x64, 0x82, 0x99, 0xb8, 0x55, 0x1c, 0x8a, 0x20, 0xbc, 0x7c, 0x20, 0x0e, 0x39, 0xbf, 0xd4,
	0x95, 0x8b, 0x4b, 0x5d, 0xf9, 0x7b, 0xa9, 0x2b, 0x67, 0x57, 0x7a, 0xe3, 0xe2, 0x4a, 0x6f, 0xfc,
	0xbe, 0xd2, 0x1b, 0x5f, 0xde, 0x05, 0x54, 0x84, 0x43, 0xcf, 0xf2, 0x59, 0x6c, 0xd7, 0x76, 0x58,
	0x1a, 0x8c, 0x7f, 0xef, 0x7c, 0x3f, 0x86, 0x93, 0x78, 0x87, 0x7a, 0xfe, 0x4e, 0x1f, 0x62, 0x66,
	0x53, 0xcf, 0xb7, 0xa7, 0x3f, 0x21, 0xf5, 0xb7, 0xc6, 0x5b, 0x94, 0x0f, 0xfd, 0xd5, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x2a, 0x09, 0x26, 0x83, 0x04, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisStateRoot) > 0 {
		i -= len(m.GenesisStateRoot)
		copy(dAtA[i:], m.GenesisStateRoot)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.GenesisStateRoot)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CodeCommitment) > 0 {
		i -= len(m.CodeCommitment)
		copy(dAtA[i:], m.CodeCommitment)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.CodeCommitment)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StateTransitionVerifierKey) > 0 {
		i -= len(m.StateTransitionVerifierKey)
		copy(dAtA[i:], m.StateTransitionVerifierKey)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.StateTransitionVerifierKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.LatestHeight != 0 {
		i = encodeVarintGroth16(dAtA, i, uint64(m.LatestHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGroth16(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DataRoots) > 0 {
		for iNdEx := len(m.DataRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DataRoots[iNdEx])
			copy(dAtA[i:], m.DataRoots[iNdEx])
			i = encodeVarintGroth16(dAtA, i, uint64(len(m.DataRoots[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.NewCelestiaHeaderHash) > 0 {
		i -= len(m.NewCelestiaHeaderHash)
		copy(dAtA[i:], m.NewCelestiaHeaderHash)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.NewCelestiaHeaderHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.NewHeight != 0 {
		i = encodeVarintGroth16(dAtA, i, uint64(m.NewHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.NewStateRoot) > 0 {
		i -= len(m.NewStateRoot)
		copy(dAtA[i:], m.NewStateRoot)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.NewStateRoot)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TrustedCelestiaHeaderHash) > 0 {
		i -= len(m.TrustedCelestiaHeaderHash)
		copy(dAtA[i:], m.TrustedCelestiaHeaderHash)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.TrustedCelestiaHeaderHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TrustedHeight != 0 {
		i = encodeVarintGroth16(dAtA, i, uint64(m.TrustedHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.StateTransitionProof) > 0 {
		i -= len(m.StateTransitionProof)
		copy(dAtA[i:], m.StateTransitionProof)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.StateTransitionProof)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGroth16(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroth16(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LatestHeight != 0 {
		n += 1 + sovGroth16(uint64(m.LatestHeight))
	}
	l = len(m.StateTransitionVerifierKey)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	l = len(m.CodeCommitment)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	l = len(m.GenesisStateRoot)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovGroth16(uint64(l))
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StateTransitionProof)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	if m.TrustedHeight != 0 {
		n += 1 + sovGroth16(uint64(m.TrustedHeight))
	}
	l = len(m.TrustedCelestiaHeaderHash)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	l = len(m.NewStateRoot)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	if m.NewHeight != 0 {
		n += 1 + sovGroth16(uint64(m.NewHeight))
	}
	l = len(m.NewCelestiaHeaderHash)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	if len(m.DataRoots) > 0 {
		for _, b := range m.DataRoots {
			l = len(b)
			n += 1 + l + sovGroth16(uint64(l))
		}
	}
	return n
}

func sovGroth16(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroth16(x uint64) (n int) {
	return sovGroth16(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroth16
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			m.LatestHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateTransitionVerifierKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateTransitionVerifierKey = append(m.StateTransitionVerifierKey[:0], dAtA[iNdEx:postIndex]...)
			if m.StateTransitionVerifierKey == nil {
				m.StateTransitionVerifierKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeCommitment", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeCommitment = append(m.CodeCommitment[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeCommitment == nil {
				m.CodeCommitment = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisStateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisStateRoot = append(m.GenesisStateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.GenesisStateRoot == nil {
				m.GenesisStateRoot = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroth16(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroth16
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroth16
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroth16(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroth16
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroth16
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateTransitionProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateTransitionProof = append(m.StateTransitionProof[:0], dAtA[iNdEx:postIndex]...)
			if m.StateTransitionProof == nil {
				m.StateTransitionProof = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedHeight", wireType)
			}
			m.TrustedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrustedHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedCelestiaHeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedCelestiaHeaderHash = append(m.TrustedCelestiaHeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TrustedCelestiaHeaderHash == nil {
				m.TrustedCelestiaHeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewStateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewStateRoot = append(m.NewStateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.NewStateRoot == nil {
				m.NewStateRoot = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewHeight", wireType)
			}
			m.NewHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCelestiaHeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewCelestiaHeaderHash = append(m.NewCelestiaHeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NewCelestiaHeaderHash == nil {
				m.NewCelestiaHeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRoots = append(m.DataRoots, make([]byte, postIndex-iNdEx))
			copy(m.DataRoots[len(m.DataRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroth16(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroth16
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGroth16(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroth16
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGroth16
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroth16
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroth16
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroth16        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroth16          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroth16 = fmt.Errorf("proto: unexpected end of group")
)
