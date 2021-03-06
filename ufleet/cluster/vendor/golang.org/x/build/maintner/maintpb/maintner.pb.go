// Code generated by protoc-gen-go.
// source: maintner.proto
// DO NOT EDIT!

/*
Package maintpb is a generated protocol buffer package.

It is generated from these files:
	maintner.proto

It has these top-level messages:
	Mutation
	GithubIssueMutation
	GithubIssueCommentSyncStatus
	GithubIssueCommentMutation
	GithubUser
	GitMutation
	GitRepo
	GitCommit
	GitDiffTree
	GitDiffTreeFile
*/
package maintpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Mutation struct {
	GithubIssue *GithubIssueMutation `protobuf:"bytes,1,opt,name=github_issue,json=githubIssue" json:"github_issue,omitempty"`
	Git         *GitMutation         `protobuf:"bytes,2,opt,name=git" json:"git,omitempty"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Mutation) GetGithubIssue() *GithubIssueMutation {
	if m != nil {
		return m.GithubIssue
	}
	return nil
}

func (m *Mutation) GetGit() *GitMutation {
	if m != nil {
		return m.Git
	}
	return nil
}

type GithubIssueMutation struct {
	Owner  string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Number int32  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	// not_exist is set true if the issue has been found to not exist.
	// If true, the owner/repo/number fields above must still be set.
	// If a future issue mutation for the same number arrives without
	// not_exist set, then the issue comes back to life.
	NotExist         bool                          `protobuf:"varint,13,opt,name=not_exist,json=notExist" json:"not_exist,omitempty"`
	Id               int64                         `protobuf:"varint,12,opt,name=id" json:"id,omitempty"`
	User             *GithubUser                   `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Assignees        []*GithubUser                 `protobuf:"bytes,10,rep,name=assignees" json:"assignees,omitempty"`
	DeletedAssignees []int64                       `protobuf:"varint,11,rep,packed,name=deleted_assignees,json=deletedAssignees" json:"deleted_assignees,omitempty"`
	Created          *google_protobuf.Timestamp    `protobuf:"bytes,5,opt,name=created" json:"created,omitempty"`
	Updated          *google_protobuf.Timestamp    `protobuf:"bytes,6,opt,name=updated" json:"updated,omitempty"`
	Body             string                        `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	Title            string                        `protobuf:"bytes,9,opt,name=title" json:"title,omitempty"`
	Comment          []*GithubIssueCommentMutation `protobuf:"bytes,8,rep,name=comment" json:"comment,omitempty"`
	CommentStatus    *GithubIssueCommentSyncStatus `protobuf:"bytes,14,opt,name=comment_status,json=commentStatus" json:"comment_status,omitempty"`
}

func (m *GithubIssueMutation) Reset()                    { *m = GithubIssueMutation{} }
func (m *GithubIssueMutation) String() string            { return proto.CompactTextString(m) }
func (*GithubIssueMutation) ProtoMessage()               {}
func (*GithubIssueMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GithubIssueMutation) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GithubIssueMutation) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *GithubIssueMutation) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *GithubIssueMutation) GetNotExist() bool {
	if m != nil {
		return m.NotExist
	}
	return false
}

func (m *GithubIssueMutation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubIssueMutation) GetUser() *GithubUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GithubIssueMutation) GetAssignees() []*GithubUser {
	if m != nil {
		return m.Assignees
	}
	return nil
}

func (m *GithubIssueMutation) GetDeletedAssignees() []int64 {
	if m != nil {
		return m.DeletedAssignees
	}
	return nil
}

func (m *GithubIssueMutation) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *GithubIssueMutation) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *GithubIssueMutation) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GithubIssueMutation) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GithubIssueMutation) GetComment() []*GithubIssueCommentMutation {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *GithubIssueMutation) GetCommentStatus() *GithubIssueCommentSyncStatus {
	if m != nil {
		return m.CommentStatus
	}
	return nil
}

// GithubIssueCommentSyncStatus notes where syncing is at for comments
// on an issue,
// This mutation type is only made at/after the same top-level mutation
// which created the corresponding comments.
type GithubIssueCommentSyncStatus struct {
	// server_date is the "Date" response header from Github for the
	// final HTTP response.
	ServerDate *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=server_date,json=serverDate" json:"server_date,omitempty"`
}

func (m *GithubIssueCommentSyncStatus) Reset()                    { *m = GithubIssueCommentSyncStatus{} }
func (m *GithubIssueCommentSyncStatus) String() string            { return proto.CompactTextString(m) }
func (*GithubIssueCommentSyncStatus) ProtoMessage()               {}
func (*GithubIssueCommentSyncStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GithubIssueCommentSyncStatus) GetServerDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ServerDate
	}
	return nil
}

type GithubIssueCommentMutation struct {
	Id      int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	User    *GithubUser                `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Body    string                     `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	Updated *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated" json:"updated,omitempty"`
}

func (m *GithubIssueCommentMutation) Reset()                    { *m = GithubIssueCommentMutation{} }
func (m *GithubIssueCommentMutation) String() string            { return proto.CompactTextString(m) }
func (*GithubIssueCommentMutation) ProtoMessage()               {}
func (*GithubIssueCommentMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GithubIssueCommentMutation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubIssueCommentMutation) GetUser() *GithubUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GithubIssueCommentMutation) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GithubIssueCommentMutation) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *GithubIssueCommentMutation) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

type GithubUser struct {
	Id    int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login" json:"login,omitempty"`
}

func (m *GithubUser) Reset()                    { *m = GithubUser{} }
func (m *GithubUser) String() string            { return proto.CompactTextString(m) }
func (*GithubUser) ProtoMessage()               {}
func (*GithubUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GithubUser) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubUser) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type GitMutation struct {
	Repo *GitRepo `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	// commit adds a commit, or adds new information to a commit if fields
	// are added in the future.
	Commit *GitCommit `protobuf:"bytes,2,opt,name=commit" json:"commit,omitempty"`
}

func (m *GitMutation) Reset()                    { *m = GitMutation{} }
func (m *GitMutation) String() string            { return proto.CompactTextString(m) }
func (*GitMutation) ProtoMessage()               {}
func (*GitMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GitMutation) GetRepo() *GitRepo {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *GitMutation) GetCommit() *GitCommit {
	if m != nil {
		return m.Commit
	}
	return nil
}

// GitRepo identifies a git repo being mutated.
type GitRepo struct {
	// If go_repo is set, it identifies a go.googlesource.com/<go_repo> repo.
	GoRepo string `protobuf:"bytes,1,opt,name=go_repo,json=goRepo" json:"go_repo,omitempty"`
}

func (m *GitRepo) Reset()                    { *m = GitRepo{} }
func (m *GitRepo) String() string            { return proto.CompactTextString(m) }
func (*GitRepo) ProtoMessage()               {}
func (*GitRepo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GitRepo) GetGoRepo() string {
	if m != nil {
		return m.GoRepo
	}
	return ""
}

type GitCommit struct {
	Sha1 string `protobuf:"bytes,1,opt,name=sha1" json:"sha1,omitempty"`
	// raw is the "git cat-file commit $sha1" output.
	Raw      []byte       `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	DiffTree *GitDiffTree `protobuf:"bytes,3,opt,name=diff_tree,json=diffTree" json:"diff_tree,omitempty"`
}

func (m *GitCommit) Reset()                    { *m = GitCommit{} }
func (m *GitCommit) String() string            { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()               {}
func (*GitCommit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GitCommit) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func (m *GitCommit) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *GitCommit) GetDiffTree() *GitDiffTree {
	if m != nil {
		return m.DiffTree
	}
	return nil
}

// git diff-tree --numstat oldtree newtree
type GitDiffTree struct {
	File []*GitDiffTreeFile `protobuf:"bytes,1,rep,name=file" json:"file,omitempty"`
}

func (m *GitDiffTree) Reset()                    { *m = GitDiffTree{} }
func (m *GitDiffTree) String() string            { return proto.CompactTextString(m) }
func (*GitDiffTree) ProtoMessage()               {}
func (*GitDiffTree) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GitDiffTree) GetFile() []*GitDiffTreeFile {
	if m != nil {
		return m.File
	}
	return nil
}

// GitDiffTreeFile represents one line of `git diff-tree --numstat` output.
type GitDiffTreeFile struct {
	File    string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Added   int64  `protobuf:"varint,2,opt,name=added" json:"added,omitempty"`
	Deleted int64  `protobuf:"varint,3,opt,name=deleted" json:"deleted,omitempty"`
	Binary  bool   `protobuf:"varint,4,opt,name=binary" json:"binary,omitempty"`
}

func (m *GitDiffTreeFile) Reset()                    { *m = GitDiffTreeFile{} }
func (m *GitDiffTreeFile) String() string            { return proto.CompactTextString(m) }
func (*GitDiffTreeFile) ProtoMessage()               {}
func (*GitDiffTreeFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GitDiffTreeFile) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *GitDiffTreeFile) GetAdded() int64 {
	if m != nil {
		return m.Added
	}
	return 0
}

func (m *GitDiffTreeFile) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func (m *GitDiffTreeFile) GetBinary() bool {
	if m != nil {
		return m.Binary
	}
	return false
}

func init() {
	proto.RegisterType((*Mutation)(nil), "maintpb.Mutation")
	proto.RegisterType((*GithubIssueMutation)(nil), "maintpb.GithubIssueMutation")
	proto.RegisterType((*GithubIssueCommentSyncStatus)(nil), "maintpb.GithubIssueCommentSyncStatus")
	proto.RegisterType((*GithubIssueCommentMutation)(nil), "maintpb.GithubIssueCommentMutation")
	proto.RegisterType((*GithubUser)(nil), "maintpb.GithubUser")
	proto.RegisterType((*GitMutation)(nil), "maintpb.GitMutation")
	proto.RegisterType((*GitRepo)(nil), "maintpb.GitRepo")
	proto.RegisterType((*GitCommit)(nil), "maintpb.GitCommit")
	proto.RegisterType((*GitDiffTree)(nil), "maintpb.GitDiffTree")
	proto.RegisterType((*GitDiffTreeFile)(nil), "maintpb.GitDiffTreeFile")
}

func init() { proto.RegisterFile("maintner.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x56, 0xe2, 0x7c, 0x4e, 0xda, 0xbc, 0x7d, 0xb7, 0x15, 0xac, 0x4a, 0x0f, 0x91, 0xf9, 0x8a,
	0x00, 0x39, 0x6a, 0xe1, 0x56, 0x21, 0x84, 0x5a, 0x40, 0x48, 0x70, 0xd9, 0x96, 0x13, 0x07, 0xcb,
	0x8e, 0x27, 0xee, 0x4a, 0xb6, 0xd7, 0xda, 0x5d, 0x53, 0xf2, 0xaf, 0xf8, 0x33, 0xfc, 0x1f, 0xb4,
	0xeb, 0x75, 0x12, 0xf5, 0x13, 0x71, 0x9b, 0xf1, 0x3c, 0xcf, 0x78, 0xe6, 0x99, 0x99, 0x85, 0x71,
	0x1e, 0xf1, 0x42, 0x17, 0x28, 0x83, 0x52, 0x0a, 0x2d, 0x48, 0xdf, 0xfa, 0x65, 0xbc, 0x7f, 0x9c,
	0x72, 0x7d, 0x51, 0xc5, 0xc1, 0x5c, 0xe4, 0xb3, 0x54, 0x64, 0x51, 0x91, 0xce, 0x2c, 0x22, 0xae,
	0x16, 0xb3, 0x52, 0x2f, 0x4b, 0x54, 0x33, 0xcd, 0x73, 0x54, 0x3a, 0xca, 0xcb, 0xb5, 0x55, 0x67,
	0xf1, 0x15, 0x0c, 0xbe, 0x56, 0x3a, 0xd2, 0x5c, 0x14, 0xe4, 0x1d, 0x6c, 0xd5, 0xa9, 0x42, 0xae,
	0x54, 0x85, 0xb4, 0x35, 0x69, 0x4d, 0x47, 0x47, 0x07, 0x81, 0xfb, 0x51, 0xf0, 0xc9, 0x06, 0x3f,
	0x9b, 0x58, 0xc3, 0x61, 0xa3, 0x74, 0xfd, 0x91, 0x3c, 0x03, 0x2f, 0xe5, 0x9a, 0xb6, 0x2d, 0x6f,
	0x6f, 0x93, 0xb7, 0xc2, 0x1b, 0x80, 0xff, 0xab, 0x03, 0xbb, 0x37, 0x24, 0x23, 0x7b, 0xd0, 0x15,
	0x97, 0x05, 0x4a, 0xfb, 0xe7, 0x21, 0xab, 0x1d, 0x42, 0xa0, 0x23, 0xb1, 0x14, 0x36, 0xed, 0x90,
	0x59, 0x9b, 0x3c, 0x80, 0x5e, 0x51, 0xe5, 0x31, 0x4a, 0xea, 0x4d, 0x5a, 0xd3, 0x2e, 0x73, 0x1e,
	0x79, 0x04, 0xc3, 0x42, 0xe8, 0x10, 0x7f, 0x72, 0xa5, 0xe9, 0xf6, 0xa4, 0x35, 0x1d, 0xb0, 0x41,
	0x21, 0xf4, 0x07, 0xe3, 0x93, 0x31, 0xb4, 0x79, 0x42, 0xb7, 0x26, 0xad, 0xa9, 0xc7, 0xda, 0x3c,
	0x21, 0xcf, 0xa1, 0x53, 0x29, 0x94, 0xb4, 0x63, 0xeb, 0xdd, 0xbd, 0xd2, 0xe7, 0x37, 0x85, 0x92,
	0x59, 0x00, 0x39, 0x84, 0x61, 0xa4, 0x14, 0x4f, 0x0b, 0x44, 0x45, 0x61, 0xe2, 0xdd, 0x86, 0x5e,
	0xa3, 0xc8, 0x4b, 0xf8, 0x3f, 0xc1, 0x0c, 0x35, 0x26, 0xe1, 0x9a, 0x3a, 0x9a, 0x78, 0x53, 0x8f,
	0xed, 0xb8, 0xc0, 0xfb, 0x15, 0xf8, 0x0d, 0xf4, 0xe7, 0x12, 0x23, 0x8d, 0x09, 0xed, 0xda, 0x5a,
	0xf6, 0x83, 0x54, 0x88, 0x34, 0xc3, 0xa0, 0x19, 0x64, 0x70, 0xde, 0xcc, 0x8d, 0x35, 0x50, 0xc3,
	0xaa, 0xca, 0xc4, 0xb2, 0x7a, 0xf7, 0xb3, 0x1c, 0xd4, 0xa8, 0x19, 0x8b, 0x64, 0x49, 0xfb, 0xb5,
	0x9a, 0xc6, 0x36, 0xba, 0x6b, 0xae, 0x33, 0xa4, 0xc3, 0x5a, 0x77, 0xeb, 0x90, 0xb7, 0xd0, 0x9f,
	0x8b, 0x3c, 0xc7, 0x42, 0xd3, 0x81, 0xed, 0xf9, 0xf1, 0x4d, 0x9b, 0x70, 0x52, 0x43, 0x56, 0x03,
	0x6e, 0x38, 0xe4, 0x0b, 0x8c, 0x9d, 0x19, 0x2a, 0x1d, 0xe9, 0x4a, 0xd1, 0xb1, 0xad, 0xf2, 0xe9,
	0x1d, 0x59, 0xce, 0x96, 0xc5, 0xfc, 0xcc, 0x82, 0xd9, 0xb6, 0x23, 0xd7, 0xae, 0xff, 0x1d, 0x0e,
	0xee, 0x82, 0x93, 0x63, 0x18, 0x29, 0x94, 0x3f, 0x50, 0x86, 0xa6, 0x4d, 0xb7, 0xba, 0x77, 0x09,
	0x02, 0x35, 0xfc, 0x34, 0xd2, 0xe8, 0xff, 0x6e, 0xc1, 0xfe, 0xed, 0x2d, 0xb9, 0xbd, 0x69, 0x5d,
	0xdb, 0x9b, 0xf6, 0x7d, 0x7b, 0xd3, 0x68, 0xed, 0x6d, 0x68, 0xbd, 0x31, 0xeb, 0xce, 0x3f, 0xcd,
	0xba, 0xfb, 0xd7, 0xb3, 0xf6, 0x8f, 0x00, 0xd6, 0x35, 0x5d, 0x6b, 0x63, 0x0f, 0xba, 0x99, 0x48,
	0x79, 0xe1, 0x0e, 0xab, 0x76, 0xfc, 0x10, 0x46, 0x1b, 0xf7, 0x4a, 0x9e, 0xb8, 0xe3, 0xab, 0x05,
	0xdd, 0xd9, 0xec, 0x95, 0x61, 0x29, 0xdc, 0x39, 0xbe, 0x80, 0x9e, 0x19, 0xd7, 0xea, 0xf6, 0xc9,
	0x26, 0xee, 0xc4, 0x46, 0x98, 0x43, 0xf8, 0x3e, 0xf4, 0x1d, 0x99, 0x3c, 0x84, 0x7e, 0x2a, 0xc2,
	0x55, 0xfe, 0x21, 0xeb, 0xa5, 0xc2, 0x04, 0xfc, 0x04, 0x86, 0x2b, 0xa2, 0x51, 0x51, 0x5d, 0x44,
	0x87, 0x0e, 0x62, 0x6d, 0xb2, 0x03, 0x9e, 0x8c, 0x2e, 0xed, 0xdf, 0xb6, 0x98, 0x31, 0xcd, 0x8d,
	0x26, 0x7c, 0xb1, 0x08, 0xb5, 0x44, 0xb4, 0x82, 0x5f, 0x79, 0x81, 0x4e, 0xf9, 0x62, 0x71, 0x2e,
	0x11, 0xd9, 0x20, 0x71, 0x96, 0x7f, 0x6c, 0x5b, 0x6d, 0x02, 0xe4, 0x15, 0x74, 0x16, 0x3c, 0x33,
	0xbb, 0x63, 0x96, 0x9d, 0xde, 0x44, 0xfe, 0xc8, 0x33, 0x64, 0x16, 0xe5, 0xe7, 0xf0, 0xdf, 0x95,
	0x80, 0x29, 0xd4, 0x25, 0xb0, 0x85, 0x1a, 0xdb, 0x88, 0x1c, 0x25, 0x09, 0x26, 0xb6, 0x54, 0x8f,
	0xd5, 0x0e, 0xa1, 0xd0, 0x77, 0x8f, 0x80, 0x2d, 0xd5, 0x63, 0x8d, 0x6b, 0x1e, 0xb6, 0x98, 0x17,
	0x91, 0x5c, 0xda, 0xed, 0x18, 0x30, 0xe7, 0xc5, 0x3d, 0x3b, 0xe7, 0xd7, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf3, 0xbd, 0x1b, 0x44, 0x06, 0x06, 0x00, 0x00,
}
