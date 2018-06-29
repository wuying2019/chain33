// Code generated by protoc-gen-go.
// source: transaction.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateTx struct {
	To          string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount      int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Fee         int64  `protobuf:"varint,3,opt,name=fee" json:"fee,omitempty"`
	Note        string `protobuf:"bytes,4,opt,name=note" json:"note,omitempty"`
	IsWithdraw  bool   `protobuf:"varint,5,opt,name=isWithdraw" json:"isWithdraw,omitempty"`
	IsToken     bool   `protobuf:"varint,6,opt,name=isToken" json:"isToken,omitempty"`
	TokenSymbol string `protobuf:"bytes,7,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
	ExecName    string `protobuf:"bytes,8,opt,name=execName" json:"execName,omitempty"`
}

func (m *CreateTx) Reset()                    { *m = CreateTx{} }
func (m *CreateTx) String() string            { return proto.CompactTextString(m) }
func (*CreateTx) ProtoMessage()               {}
func (*CreateTx) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *CreateTx) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CreateTx) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreateTx) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *CreateTx) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *CreateTx) GetIsWithdraw() bool {
	if m != nil {
		return m.IsWithdraw
	}
	return false
}

func (m *CreateTx) GetIsToken() bool {
	if m != nil {
		return m.IsToken
	}
	return false
}

func (m *CreateTx) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *CreateTx) GetExecName() string {
	if m != nil {
		return m.ExecName
	}
	return ""
}

type UnsignTx struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UnsignTx) Reset()                    { *m = UnsignTx{} }
func (m *UnsignTx) String() string            { return proto.CompactTextString(m) }
func (*UnsignTx) ProtoMessage()               {}
func (*UnsignTx) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *UnsignTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SignedTx struct {
	Unsign []byte `protobuf:"bytes,1,opt,name=unsign,proto3" json:"unsign,omitempty"`
	Sign   []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	Pubkey []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Ty     int32  `protobuf:"varint,4,opt,name=ty" json:"ty,omitempty"`
}

func (m *SignedTx) Reset()                    { *m = SignedTx{} }
func (m *SignedTx) String() string            { return proto.CompactTextString(m) }
func (*SignedTx) ProtoMessage()               {}
func (*SignedTx) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *SignedTx) GetUnsign() []byte {
	if m != nil {
		return m.Unsign
	}
	return nil
}

func (m *SignedTx) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *SignedTx) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *SignedTx) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

type Transaction struct {
	Execer    []byte     `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature *Signature `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	Fee       int64      `protobuf:"varint,4,opt,name=fee" json:"fee,omitempty"`
	Expire    int64      `protobuf:"varint,5,opt,name=expire" json:"expire,omitempty"`
	// 随机ID，可以防止payload 相同的时候，交易重复
	Nonce int64 `protobuf:"varint,6,opt,name=nonce" json:"nonce,omitempty"`
	// 对方地址，如果没有对方地址，可以为空
	To         string `protobuf:"bytes,7,opt,name=to" json:"to,omitempty"`
	GroupCount int32  `protobuf:"varint,8,opt,name=groupCount" json:"groupCount,omitempty"`
	Header     []byte `protobuf:"bytes,9,opt,name=header,proto3" json:"header,omitempty"`
	Next       []byte `protobuf:"bytes,10,opt,name=next,proto3" json:"next,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *Transaction) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Transaction) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Transaction) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *Transaction) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Transaction) GetGroupCount() int32 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *Transaction) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetNext() []byte {
	if m != nil {
		return m.Next
	}
	return nil
}

type Transactions struct {
	Txs []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *Transactions) Reset()                    { *m = Transactions{} }
func (m *Transactions) String() string            { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()               {}
func (*Transactions) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *Transactions) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

// 对于一个交易组中的交易，要么全部成功，要么全部失败
// 这个要好好设计一下
// 最好交易构成一个链条[prevhash].独立的交易构成链条
// 只要这个组中有一个执行是出错的，那么就执行不成功
// 三种签名支持
// ty = 1 -> secp256k1
// ty = 2 -> ed25519
// ty = 3 -> sm2
type Signature struct {
	Ty        int32  `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Pubkey    []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *Signature) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *Signature) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AddrOverview struct {
	Reciver int64 `protobuf:"varint,1,opt,name=reciver" json:"reciver,omitempty"`
	Balance int64 `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	TxCount int64 `protobuf:"varint,3,opt,name=txCount" json:"txCount,omitempty"`
}

func (m *AddrOverview) Reset()                    { *m = AddrOverview{} }
func (m *AddrOverview) String() string            { return proto.CompactTextString(m) }
func (*AddrOverview) ProtoMessage()               {}
func (*AddrOverview) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *AddrOverview) GetReciver() int64 {
	if m != nil {
		return m.Reciver
	}
	return 0
}

func (m *AddrOverview) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AddrOverview) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

type ReqAddr struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	// 表示取所有/from/to/其他的hash列表
	Flag      int32 `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
	Count     int32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Direction int32 `protobuf:"varint,4,opt,name=direction" json:"direction,omitempty"`
	Height    int64 `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Index     int64 `protobuf:"varint,6,opt,name=index" json:"index,omitempty"`
}

func (m *ReqAddr) Reset()                    { *m = ReqAddr{} }
func (m *ReqAddr) String() string            { return proto.CompactTextString(m) }
func (*ReqAddr) ProtoMessage()               {}
func (*ReqAddr) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *ReqAddr) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqAddr) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *ReqAddr) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqAddr) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqAddr) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqAddr) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type HexTx struct {
	Tx string `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
}

func (m *HexTx) Reset()                    { *m = HexTx{} }
func (m *HexTx) String() string            { return proto.CompactTextString(m) }
func (*HexTx) ProtoMessage()               {}
func (*HexTx) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *HexTx) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

type ReplyTxInfo struct {
	Hash   []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Index  int64  `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
}

func (m *ReplyTxInfo) Reset()                    { *m = ReplyTxInfo{} }
func (m *ReplyTxInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxInfo) ProtoMessage()               {}
func (*ReplyTxInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *ReplyTxInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ReplyTxInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReplyTxInfo) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ReqTxList struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *ReqTxList) Reset()                    { *m = ReqTxList{} }
func (m *ReqTxList) String() string            { return proto.CompactTextString(m) }
func (*ReqTxList) ProtoMessage()               {}
func (*ReqTxList) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *ReqTxList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyTxList struct {
	Txs []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *ReplyTxList) Reset()                    { *m = ReplyTxList{} }
func (m *ReplyTxList) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxList) ProtoMessage()               {}
func (*ReplyTxList) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{11} }

func (m *ReplyTxList) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxHashList struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
	Count  int64    `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *TxHashList) Reset()                    { *m = TxHashList{} }
func (m *TxHashList) String() string            { return proto.CompactTextString(m) }
func (*TxHashList) ProtoMessage()               {}
func (*TxHashList) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{12} }

func (m *TxHashList) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *TxHashList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyTxInfos struct {
	TxInfos []*ReplyTxInfo `protobuf:"bytes,1,rep,name=txInfos" json:"txInfos,omitempty"`
}

func (m *ReplyTxInfos) Reset()                    { *m = ReplyTxInfos{} }
func (m *ReplyTxInfos) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxInfos) ProtoMessage()               {}
func (*ReplyTxInfos) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{13} }

func (m *ReplyTxInfos) GetTxInfos() []*ReplyTxInfo {
	if m != nil {
		return m.TxInfos
	}
	return nil
}

type ReceiptLog struct {
	Ty  int32  `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Log []byte `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *ReceiptLog) Reset()                    { *m = ReceiptLog{} }
func (m *ReceiptLog) String() string            { return proto.CompactTextString(m) }
func (*ReceiptLog) ProtoMessage()               {}
func (*ReceiptLog) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{14} }

func (m *ReceiptLog) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *ReceiptLog) GetLog() []byte {
	if m != nil {
		return m.Log
	}
	return nil
}

// ty = 0 -> error Receipt
// ty = 1 -> CutFee //cut fee ,bug exec not ok
// ty = 2 -> exec ok
type Receipt struct {
	Ty   int32         `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	KV   []*KeyValue   `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
	Logs []*ReceiptLog `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
}

func (m *Receipt) Reset()                    { *m = Receipt{} }
func (m *Receipt) String() string            { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()               {}
func (*Receipt) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{15} }

func (m *Receipt) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *Receipt) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *Receipt) GetLogs() []*ReceiptLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type ReceiptData struct {
	Ty   int32         `protobuf:"varint,1,opt,name=ty" json:"ty,omitempty"`
	Logs []*ReceiptLog `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
}

func (m *ReceiptData) Reset()                    { *m = ReceiptData{} }
func (m *ReceiptData) String() string            { return proto.CompactTextString(m) }
func (*ReceiptData) ProtoMessage()               {}
func (*ReceiptData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{16} }

func (m *ReceiptData) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *ReceiptData) GetLogs() []*ReceiptLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type TxResult struct {
	Height      int64        `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	Index       int32        `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Tx          *Transaction `protobuf:"bytes,3,opt,name=tx" json:"tx,omitempty"`
	Receiptdate *ReceiptData `protobuf:"bytes,4,opt,name=receiptdate" json:"receiptdate,omitempty"`
	Blocktime   int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	ActionName  string       `protobuf:"bytes,6,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *TxResult) Reset()                    { *m = TxResult{} }
func (m *TxResult) String() string            { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()               {}
func (*TxResult) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{17} }

func (m *TxResult) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TxResult) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TxResult) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TxResult) GetReceiptdate() *ReceiptData {
	if m != nil {
		return m.Receiptdate
	}
	return nil
}

func (m *TxResult) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *TxResult) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type TransactionDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Proofs     [][]byte     `protobuf:"bytes,3,rep,name=proofs,proto3" json:"proofs,omitempty"`
	Height     int64        `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,6,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,7,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,8,opt,name=fromaddr" json:"fromaddr,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
}

func (m *TransactionDetail) Reset()                    { *m = TransactionDetail{} }
func (m *TransactionDetail) String() string            { return proto.CompactTextString(m) }
func (*TransactionDetail) ProtoMessage()               {}
func (*TransactionDetail) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{18} }

func (m *TransactionDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TransactionDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *TransactionDetail) GetProofs() [][]byte {
	if m != nil {
		return m.Proofs
	}
	return nil
}

func (m *TransactionDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TransactionDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TransactionDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *TransactionDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransactionDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *TransactionDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type TransactionDetails struct {
	Txs []*TransactionDetail `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *TransactionDetails) Reset()                    { *m = TransactionDetails{} }
func (m *TransactionDetails) String() string            { return proto.CompactTextString(m) }
func (*TransactionDetails) ProtoMessage()               {}
func (*TransactionDetails) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{19} }

func (m *TransactionDetails) GetTxs() []*TransactionDetail {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ReqAddrs struct {
	Addrs []string `protobuf:"bytes,1,rep,name=addrs" json:"addrs,omitempty"`
}

func (m *ReqAddrs) Reset()                    { *m = ReqAddrs{} }
func (m *ReqAddrs) String() string            { return proto.CompactTextString(m) }
func (*ReqAddrs) ProtoMessage()               {}
func (*ReqAddrs) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{20} }

func (m *ReqAddrs) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type ReqDecodeRawTransaction struct {
	TxHex string `protobuf:"bytes,1,opt,name=txHex" json:"txHex,omitempty"`
}

func (m *ReqDecodeRawTransaction) Reset()                    { *m = ReqDecodeRawTransaction{} }
func (m *ReqDecodeRawTransaction) String() string            { return proto.CompactTextString(m) }
func (*ReqDecodeRawTransaction) ProtoMessage()               {}
func (*ReqDecodeRawTransaction) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{21} }

func (m *ReqDecodeRawTransaction) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTx)(nil), "types.CreateTx")
	proto.RegisterType((*UnsignTx)(nil), "types.UnsignTx")
	proto.RegisterType((*SignedTx)(nil), "types.SignedTx")
	proto.RegisterType((*Transaction)(nil), "types.Transaction")
	proto.RegisterType((*Transactions)(nil), "types.Transactions")
	proto.RegisterType((*Signature)(nil), "types.Signature")
	proto.RegisterType((*AddrOverview)(nil), "types.AddrOverview")
	proto.RegisterType((*ReqAddr)(nil), "types.ReqAddr")
	proto.RegisterType((*HexTx)(nil), "types.HexTx")
	proto.RegisterType((*ReplyTxInfo)(nil), "types.ReplyTxInfo")
	proto.RegisterType((*ReqTxList)(nil), "types.ReqTxList")
	proto.RegisterType((*ReplyTxList)(nil), "types.ReplyTxList")
	proto.RegisterType((*TxHashList)(nil), "types.TxHashList")
	proto.RegisterType((*ReplyTxInfos)(nil), "types.ReplyTxInfos")
	proto.RegisterType((*ReceiptLog)(nil), "types.ReceiptLog")
	proto.RegisterType((*Receipt)(nil), "types.Receipt")
	proto.RegisterType((*ReceiptData)(nil), "types.ReceiptData")
	proto.RegisterType((*TxResult)(nil), "types.TxResult")
	proto.RegisterType((*TransactionDetail)(nil), "types.TransactionDetail")
	proto.RegisterType((*TransactionDetails)(nil), "types.TransactionDetails")
	proto.RegisterType((*ReqAddrs)(nil), "types.ReqAddrs")
	proto.RegisterType((*ReqDecodeRawTransaction)(nil), "types.ReqDecodeRawTransaction")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 932 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0x86, 0x24, 0xff, 0x8e, 0x8d, 0x36, 0x21, 0x16, 0xbb, 0x42, 0x50, 0xa4, 0xae, 0xd0, 0x02,
	0x41, 0xb1, 0x70, 0x81, 0xdd, 0x3d, 0x15, 0x3d, 0xb4, 0xd8, 0x1c, 0x52, 0x64, 0xd1, 0x45, 0x19,
	0x35, 0xbd, 0x14, 0x05, 0x68, 0x69, 0x6c, 0x0b, 0x91, 0x45, 0x47, 0xa2, 0xb3, 0xf2, 0x33, 0xf4,
	0xd8, 0x73, 0xdf, 0xa9, 0x8f, 0x54, 0x70, 0x48, 0xda, 0x4c, 0x6c, 0x14, 0xb9, 0xf1, 0x1b, 0x0e,
	0xf9, 0xcd, 0x7c, 0x33, 0x1c, 0x09, 0x4e, 0x55, 0x2d, 0xaa, 0x46, 0x64, 0xaa, 0x90, 0xd5, 0x74,
	0x5d, 0x4b, 0x25, 0x59, 0x57, 0x6d, 0xd7, 0xd8, 0x9c, 0x8d, 0x33, 0xb9, 0x5a, 0x39, 0x63, 0xf2,
	0x6f, 0x00, 0x83, 0xf7, 0x35, 0x0a, 0x85, 0x69, 0xcb, 0x3e, 0x83, 0x50, 0xc9, 0x38, 0x98, 0x04,
	0x17, 0x43, 0x1e, 0x2a, 0xc9, 0x5e, 0x42, 0x4f, 0xac, 0xe4, 0xa6, 0x52, 0x71, 0x38, 0x09, 0x2e,
	0x22, 0x6e, 0x11, 0x3b, 0x81, 0x68, 0x8e, 0x18, 0x47, 0x64, 0xd4, 0x4b, 0xc6, 0xa0, 0x53, 0x49,
	0x85, 0x71, 0x87, 0xce, 0xd2, 0x9a, 0x9d, 0x03, 0x14, 0xcd, 0xef, 0x85, 0x5a, 0xe6, 0xb5, 0xf8,
	0x14, 0x77, 0x27, 0xc1, 0xc5, 0x80, 0x7b, 0x16, 0x16, 0x43, 0xbf, 0x68, 0x52, 0x79, 0x87, 0x55,
	0xdc, 0xa3, 0x4d, 0x07, 0xd9, 0x04, 0x46, 0x4a, 0x2f, 0x6e, 0xb6, 0xab, 0x99, 0x2c, 0xe3, 0x3e,
	0x5d, 0xea, 0x9b, 0xd8, 0x19, 0x0c, 0xb0, 0xc5, 0xec, 0x17, 0xb1, 0xc2, 0x78, 0x40, 0xdb, 0x3b,
	0x9c, 0x9c, 0xc3, 0xe0, 0xb7, 0xaa, 0x29, 0x16, 0x55, 0xda, 0xea, 0xb8, 0x72, 0xa1, 0x04, 0xe5,
	0x34, 0xe6, 0xb4, 0x4e, 0xfe, 0x84, 0xc1, 0x4d, 0xb1, 0xa8, 0x30, 0x4f, 0x5b, 0x9d, 0xe1, 0x86,
	0x7c, 0xad, 0x87, 0x45, 0xfa, 0x1c, 0x59, 0x43, 0x73, 0x8e, 0x6c, 0x2f, 0xa1, 0xb7, 0xde, 0xcc,
	0xee, 0x70, 0x4b, 0x89, 0x8f, 0xb9, 0x45, 0xa4, 0xda, 0x96, 0x32, 0xef, 0xf2, 0x50, 0x6d, 0x93,
	0xbf, 0x42, 0x18, 0xa5, 0x7b, 0xf5, 0xf5, 0x39, 0x1d, 0x1b, 0xd6, 0x8e, 0xc3, 0x20, 0x9d, 0xff,
	0x5a, 0x6c, 0x4b, 0x29, 0x72, 0x4b, 0xe3, 0x20, 0x9b, 0xc2, 0x50, 0x33, 0x0a, 0xb5, 0xa9, 0x8d,
	0xca, 0xa3, 0x37, 0x27, 0x53, 0xaa, 0xde, 0xf4, 0xc6, 0xd9, 0xf9, 0xde, 0xc5, 0xd5, 0xa3, 0xb3,
	0xaf, 0x07, 0x71, 0xae, 0x8b, 0x1a, 0x49, 0xf7, 0x88, 0x5b, 0xc4, 0x5e, 0x40, 0xb7, 0x92, 0x55,
	0x86, 0xa4, 0x78, 0xc4, 0x0d, 0xb0, 0x75, 0xef, 0xef, 0xea, 0x7e, 0x0e, 0xb0, 0xa8, 0xe5, 0x66,
	0xfd, 0x9e, 0x6a, 0x3f, 0xa0, 0xcc, 0x3c, 0x8b, 0xbe, 0x7d, 0x89, 0x22, 0xc7, 0x3a, 0x1e, 0x9a,
	0x8c, 0x0c, 0xa2, 0x2e, 0xc0, 0x56, 0xc5, 0x60, 0x54, 0xd3, 0xeb, 0xe4, 0x1d, 0x8c, 0x3d, 0x31,
	0x1a, 0xf6, 0x35, 0x44, 0xaa, 0x6d, 0xe2, 0x60, 0x12, 0x5d, 0x8c, 0xde, 0x30, 0x9b, 0x95, 0xe7,
	0xc1, 0xf5, 0x76, 0xf2, 0x2b, 0x0c, 0x77, 0x99, 0x5a, 0x81, 0x03, 0x27, 0xb0, 0x57, 0x88, 0xf0,
	0x51, 0x21, 0xbe, 0x78, 0x2a, 0xdb, 0xd8, 0x13, 0x29, 0xf9, 0x03, 0xc6, 0x3f, 0xe5, 0x79, 0xfd,
	0xf1, 0x01, 0xeb, 0x87, 0x02, 0xa9, 0xfd, 0x6a, 0xcc, 0x8a, 0x07, 0x5b, 0x97, 0x88, 0x3b, 0xa8,
	0x77, 0x66, 0xa2, 0x14, 0x5a, 0x26, 0xd3, 0xf7, 0x0e, 0xea, 0x1d, 0xd5, 0x1a, 0x55, 0x4c, 0xf3,
	0x3b, 0x98, 0xfc, 0x1d, 0x40, 0x9f, 0xe3, 0xbd, 0x66, 0xd0, 0x32, 0x88, 0x3c, 0xaf, 0xed, 0x43,
	0xa2, 0xb5, 0xb6, 0xcd, 0x4b, 0xb1, 0xa0, 0x0b, 0xbb, 0x9c, 0xd6, 0xba, 0x18, 0xd9, 0xee, 0xae,
	0x2e, 0x37, 0x40, 0x67, 0x91, 0x17, 0x35, 0x92, 0x18, 0xb6, 0xab, 0xf6, 0x06, 0x23, 0x7d, 0xb1,
	0x58, 0x2a, 0x57, 0x58, 0x83, 0xf4, 0x5d, 0x45, 0x95, 0x63, 0xeb, 0x0a, 0x4b, 0x20, 0x79, 0x05,
	0xdd, 0x2b, 0x6c, 0xed, 0xcb, 0x6e, 0x77, 0x2f, 0xbb, 0x4d, 0x3e, 0xc2, 0x88, 0xe3, 0xba, 0xdc,
	0xa6, 0xed, 0xcf, 0xd5, 0x5c, 0xea, 0xe8, 0x96, 0xa2, 0x59, 0xba, 0x67, 0xa2, 0xd7, 0x1e, 0x53,
	0x78, 0x9c, 0x29, 0xf2, 0x99, 0xbe, 0x82, 0x21, 0xc7, 0xfb, 0xb4, 0xfd, 0x50, 0x34, 0x6a, 0x9f,
	0x98, 0x11, 0xd6, 0x80, 0xe4, 0xed, 0x8e, 0x93, 0x9c, 0x9e, 0xd7, 0x08, 0xdf, 0x03, 0xa4, 0xed,
	0x95, 0x68, 0x96, 0x74, 0x46, 0xc7, 0x24, 0x9a, 0x25, 0x9a, 0x63, 0xba, 0xf1, 0x08, 0xed, 0x09,
	0x43, 0x9f, 0xf0, 0x07, 0x18, 0x7b, 0x49, 0x36, 0xec, 0xb5, 0xae, 0x1e, 0x2d, 0x9f, 0xb0, 0x7a,
	0x5e, 0xdc, 0xb9, 0x24, 0x53, 0x00, 0x8e, 0x19, 0x16, 0x6b, 0xf5, 0x41, 0x2e, 0x0e, 0x7a, 0xf0,
	0x04, 0xa2, 0x52, 0x2e, 0x6c, 0x03, 0xea, 0x65, 0x22, 0x74, 0x03, 0x90, 0xff, 0x81, 0xf3, 0x97,
	0x10, 0x5e, 0xdf, 0xc6, 0x21, 0x71, 0x7e, 0x6e, 0x39, 0xaf, 0x71, 0x7b, 0x2b, 0xca, 0x0d, 0xf2,
	0xf0, 0xfa, 0x96, 0x7d, 0x03, 0x9d, 0x52, 0x2e, 0x9a, 0x38, 0x22, 0x97, 0xd3, 0x5d, 0x58, 0x8e,
	0x9e, 0xd3, 0x76, 0x72, 0xa9, 0x15, 0x24, 0xdb, 0xa5, 0x50, 0xe2, 0x80, 0xe6, 0x99, 0xb7, 0xe8,
	0x91, 0x9f, 0xb6, 0x1c, 0x9b, 0x4d, 0xa9, 0xbc, 0x2a, 0x07, 0xc7, 0xab, 0x6c, 0x1a, 0xd6, 0x00,
	0x96, 0x50, 0x1b, 0x99, 0x89, 0x74, 0xac, 0x64, 0xa1, 0x6a, 0xd9, 0x3b, 0x18, 0xd5, 0x86, 0x32,
	0x17, 0xf6, 0x8b, 0xe0, 0x2b, 0xbd, 0x0b, 0x9f, 0xfb, 0x6e, 0xba, 0xeb, 0x67, 0xa5, 0xcc, 0xee,
	0x54, 0xb1, 0x72, 0x33, 0x6b, 0x6f, 0xd0, 0x03, 0xc9, 0x30, 0xd0, 0xc0, 0xef, 0x51, 0x1b, 0x7b,
	0x96, 0xe4, 0x9f, 0x10, 0x4e, 0xbd, 0x38, 0x2e, 0x51, 0x89, 0xa2, 0xb4, 0xd1, 0x06, 0xff, 0x1b,
	0xed, 0x6b, 0x9a, 0x02, 0x3a, 0x0c, 0xca, 0xf4, 0x78, 0xa4, 0xce, 0x85, 0x26, 0x4f, 0x2d, 0xe5,
	0xdc, 0x68, 0xac, 0x27, 0x0f, 0x21, 0x4f, 0xc5, 0xce, 0x71, 0x15, 0xbb, 0xde, 0x5b, 0x79, 0x9c,
	0x6b, 0xef, 0x69, 0xae, 0xfb, 0x8f, 0x6e, 0xff, 0xd1, 0x47, 0xf7, 0x0c, 0x06, 0xf3, 0x5a, 0xae,
	0x68, 0xb2, 0xd8, 0x4f, 0x9e, 0xc3, 0x4f, 0xf4, 0x19, 0x1e, 0xe8, 0xf3, 0x23, 0xb0, 0x03, 0x79,
	0x1a, 0xf6, 0xad, 0xff, 0x02, 0xe3, 0x43, 0x81, 0x8c, 0x9f, 0x79, 0x87, 0x13, 0x18, 0xd8, 0xf1,
	0x46, 0xaf, 0x4d, 0xb3, 0x9a, 0x93, 0x43, 0x6e, 0x40, 0xf2, 0x1d, 0xbc, 0xe2, 0x78, 0x7f, 0x89,
	0x99, 0xcc, 0x91, 0x8b, 0x4f, 0xfe, 0x17, 0xf0, 0x05, 0x74, 0x55, 0x7b, 0x85, 0x6e, 0x00, 0x19,
	0x30, 0xeb, 0xd1, 0x1f, 0xc8, 0xdb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x08, 0xd0, 0xa4,
	0xab, 0x08, 0x00, 0x00,
}
