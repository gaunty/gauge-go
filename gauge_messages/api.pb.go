// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package gauge_messages is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetProjectRootRequest
	GetProjectRootResponse
	GetInstallationRootRequest
	GetInstallationRootResponse
	GetAllStepsRequest
	GetAllStepsResponse
	GetAllSpecsRequest
	GetAllSpecsResponse
	GetAllConceptsRequest
	GetAllConceptsResponse
	ConceptInfo
	GetStepValueRequest
	GetStepValueResponse
	GetLanguagePluginLibPathRequest
	GetLanguagePluginLibPathResponse
	ErrorResponse
	PerformRefactoringRequest
	PerformRefactoringResponse
	ExtractConceptInfoRequest
	ExtractConceptRequest
	TextInfo
	Step
	ExtractConceptResponse
	FormatSpecsRequest
	FormatSpecsResponse
	UnsupportedApiMessageResponse
	APIMessage
*/
package gauge_messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type APIMessage_APIMessageType int32

const (
	APIMessage_GetProjectRootRequest            APIMessage_APIMessageType = 1
	APIMessage_GetProjectRootResponse           APIMessage_APIMessageType = 2
	APIMessage_GetInstallationRootRequest       APIMessage_APIMessageType = 3
	APIMessage_GetInstallationRootResponse      APIMessage_APIMessageType = 4
	APIMessage_GetAllStepsRequest               APIMessage_APIMessageType = 5
	APIMessage_GetAllStepResponse               APIMessage_APIMessageType = 6
	APIMessage_GetAllSpecsRequest               APIMessage_APIMessageType = 7
	APIMessage_GetAllSpecsResponse              APIMessage_APIMessageType = 8
	APIMessage_GetStepValueRequest              APIMessage_APIMessageType = 9
	APIMessage_GetStepValueResponse             APIMessage_APIMessageType = 10
	APIMessage_GetLanguagePluginLibPathRequest  APIMessage_APIMessageType = 11
	APIMessage_GetLanguagePluginLibPathResponse APIMessage_APIMessageType = 12
	APIMessage_ErrorResponse                    APIMessage_APIMessageType = 13
	APIMessage_GetAllConceptsRequest            APIMessage_APIMessageType = 14
	APIMessage_GetAllConceptsResponse           APIMessage_APIMessageType = 15
	APIMessage_PerformRefactoringRequest        APIMessage_APIMessageType = 16
	APIMessage_PerformRefactoringResponse       APIMessage_APIMessageType = 17
	APIMessage_ExtractConceptRequest            APIMessage_APIMessageType = 18
	APIMessage_ExtractConceptResponse           APIMessage_APIMessageType = 19
	APIMessage_FormatSpecsRequest               APIMessage_APIMessageType = 20
	APIMessage_FormatSpecsResponse              APIMessage_APIMessageType = 21
	APIMessage_UnsupportedApiMessageResponse    APIMessage_APIMessageType = 22
)

var APIMessage_APIMessageType_name = map[int32]string{
	1:  "GetProjectRootRequest",
	2:  "GetProjectRootResponse",
	3:  "GetInstallationRootRequest",
	4:  "GetInstallationRootResponse",
	5:  "GetAllStepsRequest",
	6:  "GetAllStepResponse",
	7:  "GetAllSpecsRequest",
	8:  "GetAllSpecsResponse",
	9:  "GetStepValueRequest",
	10: "GetStepValueResponse",
	11: "GetLanguagePluginLibPathRequest",
	12: "GetLanguagePluginLibPathResponse",
	13: "ErrorResponse",
	14: "GetAllConceptsRequest",
	15: "GetAllConceptsResponse",
	16: "PerformRefactoringRequest",
	17: "PerformRefactoringResponse",
	18: "ExtractConceptRequest",
	19: "ExtractConceptResponse",
	20: "FormatSpecsRequest",
	21: "FormatSpecsResponse",
	22: "UnsupportedApiMessageResponse",
}
var APIMessage_APIMessageType_value = map[string]int32{
	"GetProjectRootRequest":            1,
	"GetProjectRootResponse":           2,
	"GetInstallationRootRequest":       3,
	"GetInstallationRootResponse":      4,
	"GetAllStepsRequest":               5,
	"GetAllStepResponse":               6,
	"GetAllSpecsRequest":               7,
	"GetAllSpecsResponse":              8,
	"GetStepValueRequest":              9,
	"GetStepValueResponse":             10,
	"GetLanguagePluginLibPathRequest":  11,
	"GetLanguagePluginLibPathResponse": 12,
	"ErrorResponse":                    13,
	"GetAllConceptsRequest":            14,
	"GetAllConceptsResponse":           15,
	"PerformRefactoringRequest":        16,
	"PerformRefactoringResponse":       17,
	"ExtractConceptRequest":            18,
	"ExtractConceptResponse":           19,
	"FormatSpecsRequest":               20,
	"FormatSpecsResponse":              21,
	"UnsupportedApiMessageResponse":    22,
}

func (x APIMessage_APIMessageType) Enum() *APIMessage_APIMessageType {
	p := new(APIMessage_APIMessageType)
	*p = x
	return p
}
func (x APIMessage_APIMessageType) String() string {
	return proto.EnumName(APIMessage_APIMessageType_name, int32(x))
}
func (x *APIMessage_APIMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(APIMessage_APIMessageType_value, data, "APIMessage_APIMessageType")
	if err != nil {
		return err
	}
	*x = APIMessage_APIMessageType(value)
	return nil
}

// / Request to get the Root Directory of the project
type GetProjectRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetProjectRootRequest) Reset()         { *m = GetProjectRootRequest{} }
func (m *GetProjectRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootRequest) ProtoMessage()    {}

// / Response of GetProjectRootRequest.
type GetProjectRootResponse struct {
	// / Holds the absolute path of the Project Root directory.
	ProjectRoot      *string `protobuf:"bytes,1,req,name=projectRoot" json:"projectRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetProjectRootResponse) Reset()         { *m = GetProjectRootResponse{} }
func (m *GetProjectRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootResponse) ProtoMessage()    {}

func (m *GetProjectRootResponse) GetProjectRoot() string {
	if m != nil && m.ProjectRoot != nil {
		return *m.ProjectRoot
	}
	return ""
}

// / Request to get the Root Directory of the Gauge installation
type GetInstallationRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetInstallationRootRequest) Reset()         { *m = GetInstallationRootRequest{} }
func (m *GetInstallationRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstallationRootRequest) ProtoMessage()    {}

// / Response of GetInstallationRootRequest
type GetInstallationRootResponse struct {
	// / Holds the absolute path of the Gauge installation directory
	InstallationRoot *string `protobuf:"bytes,1,req,name=installationRoot" json:"installationRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetInstallationRootResponse) Reset()         { *m = GetInstallationRootResponse{} }
func (m *GetInstallationRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetInstallationRootResponse) ProtoMessage()    {}

func (m *GetInstallationRootResponse) GetInstallationRoot() string {
	if m != nil && m.InstallationRoot != nil {
		return *m.InstallationRoot
	}
	return ""
}

// / Request to get all Steps in the project
type GetAllStepsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllStepsRequest) Reset()         { *m = GetAllStepsRequest{} }
func (m *GetAllStepsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsRequest) ProtoMessage()    {}

// / Response to GetAllStepsRequest
type GetAllStepsResponse struct {
	// / Holds a collection of Steps that are defined in the project.
	AllSteps         []*ProtoStepValue `protobuf:"bytes,1,rep,name=allSteps" json:"allSteps,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *GetAllStepsResponse) Reset()         { *m = GetAllStepsResponse{} }
func (m *GetAllStepsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsResponse) ProtoMessage()    {}

func (m *GetAllStepsResponse) GetAllSteps() []*ProtoStepValue {
	if m != nil {
		return m.AllSteps
	}
	return nil
}

// / Request to get all Specs in the project
type GetAllSpecsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllSpecsRequest) Reset()         { *m = GetAllSpecsRequest{} }
func (m *GetAllSpecsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsRequest) ProtoMessage()    {}

// / Response to GetAllSpecsRequest
type GetAllSpecsResponse struct {
	// / Holds a collection of Specs that are defined in the project.
	Specs            []*ProtoSpec `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GetAllSpecsResponse) Reset()         { *m = GetAllSpecsResponse{} }
func (m *GetAllSpecsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsResponse) ProtoMessage()    {}

func (m *GetAllSpecsResponse) GetSpecs() []*ProtoSpec {
	if m != nil {
		return m.Specs
	}
	return nil
}

// / Request to get all Concepts in the project
type GetAllConceptsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllConceptsRequest) Reset()         { *m = GetAllConceptsRequest{} }
func (m *GetAllConceptsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllConceptsRequest) ProtoMessage()    {}

// / Response to GetAllConceptsResponse
type GetAllConceptsResponse struct {
	// / Holds a collection of Concepts that are defined in the project.
	Concepts         []*ConceptInfo `protobuf:"bytes,1,rep,name=concepts" json:"concepts,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GetAllConceptsResponse) Reset()         { *m = GetAllConceptsResponse{} }
func (m *GetAllConceptsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllConceptsResponse) ProtoMessage()    {}

func (m *GetAllConceptsResponse) GetConcepts() []*ConceptInfo {
	if m != nil {
		return m.Concepts
	}
	return nil
}

// / Details of a Concept
type ConceptInfo struct {
	// / The text that defines a concept
	StepValue *ProtoStepValue `protobuf:"bytes,1,req,name=stepValue" json:"stepValue,omitempty"`
	// / The absolute path to the file that contains the Concept
	Filepath *string `protobuf:"bytes,2,req,name=filepath" json:"filepath,omitempty"`
	// / The line number in the file where the concept is defined.
	LineNumber       *int32 `protobuf:"varint,3,req,name=lineNumber" json:"lineNumber,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ConceptInfo) Reset()         { *m = ConceptInfo{} }
func (m *ConceptInfo) String() string { return proto.CompactTextString(m) }
func (*ConceptInfo) ProtoMessage()    {}

func (m *ConceptInfo) GetStepValue() *ProtoStepValue {
	if m != nil {
		return m.StepValue
	}
	return nil
}

func (m *ConceptInfo) GetFilepath() string {
	if m != nil && m.Filepath != nil {
		return *m.Filepath
	}
	return ""
}

func (m *ConceptInfo) GetLineNumber() int32 {
	if m != nil && m.LineNumber != nil {
		return *m.LineNumber
	}
	return 0
}

// / Request to get a Step Value.
type GetStepValueRequest struct {
	// / The text of the Step.
	StepText *string `protobuf:"bytes,1,req,name=stepText" json:"stepText,omitempty"`
	// / Flag to indicate if the Step has an inline table.
	HasInlineTable   *bool  `protobuf:"varint,2,opt,name=hasInlineTable" json:"hasInlineTable,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetStepValueRequest) Reset()         { *m = GetStepValueRequest{} }
func (m *GetStepValueRequest) String() string { return proto.CompactTextString(m) }
func (*GetStepValueRequest) ProtoMessage()    {}

func (m *GetStepValueRequest) GetStepText() string {
	if m != nil && m.StepText != nil {
		return *m.StepText
	}
	return ""
}

func (m *GetStepValueRequest) GetHasInlineTable() bool {
	if m != nil && m.HasInlineTable != nil {
		return *m.HasInlineTable
	}
	return false
}

// / Response to GetStepValueRequest
type GetStepValueResponse struct {
	// / The Step corresponding to the request provided.
	StepValue        *ProtoStepValue `protobuf:"bytes,1,req,name=stepValue" json:"stepValue,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetStepValueResponse) Reset()         { *m = GetStepValueResponse{} }
func (m *GetStepValueResponse) String() string { return proto.CompactTextString(m) }
func (*GetStepValueResponse) ProtoMessage()    {}

func (m *GetStepValueResponse) GetStepValue() *ProtoStepValue {
	if m != nil {
		return m.StepValue
	}
	return nil
}

// / Request to get the location of language plugin's Lib directory
type GetLanguagePluginLibPathRequest struct {
	// / The language to locate the lib directory for.
	Language         *string `protobuf:"bytes,1,req,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetLanguagePluginLibPathRequest) Reset()         { *m = GetLanguagePluginLibPathRequest{} }
func (m *GetLanguagePluginLibPathRequest) String() string { return proto.CompactTextString(m) }
func (*GetLanguagePluginLibPathRequest) ProtoMessage()    {}

func (m *GetLanguagePluginLibPathRequest) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

// / Response to GetLanguagePluginLibPathRequest
type GetLanguagePluginLibPathResponse struct {
	// / Absolute path to the Lib directory of the language.
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetLanguagePluginLibPathResponse) Reset()         { *m = GetLanguagePluginLibPathResponse{} }
func (m *GetLanguagePluginLibPathResponse) String() string { return proto.CompactTextString(m) }
func (*GetLanguagePluginLibPathResponse) ProtoMessage()    {}

func (m *GetLanguagePluginLibPathResponse) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

// / A generic failure response
type ErrorResponse struct {
	// / Actual error message
	Error            *string `protobuf:"bytes,1,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}

func (m *ErrorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// / Request to perform a Refactor
type PerformRefactoringRequest struct {
	// / Step to refactor
	OldStep *string `protobuf:"bytes,1,req,name=oldStep" json:"oldStep,omitempty"`
	// / Change to be made
	NewStep          *string `protobuf:"bytes,2,req,name=newStep" json:"newStep,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PerformRefactoringRequest) Reset()         { *m = PerformRefactoringRequest{} }
func (m *PerformRefactoringRequest) String() string { return proto.CompactTextString(m) }
func (*PerformRefactoringRequest) ProtoMessage()    {}

func (m *PerformRefactoringRequest) GetOldStep() string {
	if m != nil && m.OldStep != nil {
		return *m.OldStep
	}
	return ""
}

func (m *PerformRefactoringRequest) GetNewStep() string {
	if m != nil && m.NewStep != nil {
		return *m.NewStep
	}
	return ""
}

// / Response to PerformRefactoringRequest
type PerformRefactoringResponse struct {
	// / Flag indicating Success
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// / Error message if the refactoring was unsuccessful.
	Errors []string `protobuf:"bytes,2,rep,name=errors" json:"errors,omitempty"`
	// / Collection of files that were changed as part of the Refactoring.
	FilesChanged     []string `protobuf:"bytes,3,rep,name=filesChanged" json:"filesChanged,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PerformRefactoringResponse) Reset()         { *m = PerformRefactoringResponse{} }
func (m *PerformRefactoringResponse) String() string { return proto.CompactTextString(m) }
func (*PerformRefactoringResponse) ProtoMessage()    {}

func (m *PerformRefactoringResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *PerformRefactoringResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PerformRefactoringResponse) GetFilesChanged() []string {
	if m != nil {
		return m.FilesChanged
	}
	return nil
}

// / Request to perform Extract to Concept refactoring
// / The runner does not do the refactoring here, instead it provides inputs enabling the IDE to do refactoring
type ExtractConceptInfoRequest struct {
	// / The text blob containing steps that should be refactored to concept.
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExtractConceptInfoRequest) Reset()         { *m = ExtractConceptInfoRequest{} }
func (m *ExtractConceptInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ExtractConceptInfoRequest) ProtoMessage()    {}

func (m *ExtractConceptInfoRequest) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

// / Request to perform Extract to Concept refactoring
type ExtractConceptRequest struct {
	// / The Concept name given by the user
	ConceptName *Step `protobuf:"bytes,1,req,name=conceptName" json:"conceptName,omitempty"`
	// / steps to extract
	Steps []*Step `protobuf:"bytes,2,rep,name=steps" json:"steps,omitempty"`
	// / Flag indicating if refactoring should be done across project
	ChangeAcrossProject *bool `protobuf:"varint,3,req,name=changeAcrossProject" json:"changeAcrossProject,omitempty"`
	// / The concept filename in which extracted concept will be added
	ConceptFileName *string `protobuf:"bytes,4,req,name=conceptFileName" json:"conceptFileName,omitempty"`
	// / Info related to selected text, required only if changeAcrossProject is false
	SelectedTextInfo *TextInfo `protobuf:"bytes,5,opt,name=selectedTextInfo" json:"selectedTextInfo,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ExtractConceptRequest) Reset()         { *m = ExtractConceptRequest{} }
func (m *ExtractConceptRequest) String() string { return proto.CompactTextString(m) }
func (*ExtractConceptRequest) ProtoMessage()    {}

func (m *ExtractConceptRequest) GetConceptName() *Step {
	if m != nil {
		return m.ConceptName
	}
	return nil
}

func (m *ExtractConceptRequest) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *ExtractConceptRequest) GetChangeAcrossProject() bool {
	if m != nil && m.ChangeAcrossProject != nil {
		return *m.ChangeAcrossProject
	}
	return false
}

func (m *ExtractConceptRequest) GetConceptFileName() string {
	if m != nil && m.ConceptFileName != nil {
		return *m.ConceptFileName
	}
	return ""
}

func (m *ExtractConceptRequest) GetSelectedTextInfo() *TextInfo {
	if m != nil {
		return m.SelectedTextInfo
	}
	return nil
}

type TextInfo struct {
	// / The filename from where concept is being extracted
	FileName *string `protobuf:"bytes,1,req,name=fileName" json:"fileName,omitempty"`
	// / storing the starting and ending line number of selected text
	StartingLineNo   *int32 `protobuf:"varint,2,req,name=startingLineNo" json:"startingLineNo,omitempty"`
	EndLineNo        *int32 `protobuf:"varint,3,req,name=endLineNo" json:"endLineNo,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TextInfo) Reset()         { *m = TextInfo{} }
func (m *TextInfo) String() string { return proto.CompactTextString(m) }
func (*TextInfo) ProtoMessage()    {}

func (m *TextInfo) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *TextInfo) GetStartingLineNo() int32 {
	if m != nil && m.StartingLineNo != nil {
		return *m.StartingLineNo
	}
	return 0
}

func (m *TextInfo) GetEndLineNo() int32 {
	if m != nil && m.EndLineNo != nil {
		return *m.EndLineNo
	}
	return 0
}

type Step struct {
	// / name of the step
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// /  table present in step as parameter
	Table *string `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
	// / name of table in concept heading, if it comes as a param to concept
	ParamTableName   *string `protobuf:"bytes,3,opt,name=paramTableName" json:"paramTableName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}

func (m *Step) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Step) GetTable() string {
	if m != nil && m.Table != nil {
		return *m.Table
	}
	return ""
}

func (m *Step) GetParamTableName() string {
	if m != nil && m.ParamTableName != nil {
		return *m.ParamTableName
	}
	return ""
}

// / Response to perform Extract to Concept refactoring
type ExtractConceptResponse struct {
	// / Flag indicating Success
	IsSuccess *bool `protobuf:"varint,1,req,name=isSuccess" json:"isSuccess,omitempty"`
	// / Error message if the refactoring was unsuccessful.
	Error *string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	// / Collection of files that were changed as part of the Refactoring.
	FilesChanged     []string `protobuf:"bytes,3,rep,name=filesChanged" json:"filesChanged,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ExtractConceptResponse) Reset()         { *m = ExtractConceptResponse{} }
func (m *ExtractConceptResponse) String() string { return proto.CompactTextString(m) }
func (*ExtractConceptResponse) ProtoMessage()    {}

func (m *ExtractConceptResponse) GetIsSuccess() bool {
	if m != nil && m.IsSuccess != nil {
		return *m.IsSuccess
	}
	return false
}

func (m *ExtractConceptResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *ExtractConceptResponse) GetFilesChanged() []string {
	if m != nil {
		return m.FilesChanged
	}
	return nil
}

// / Request to format spec files
type FormatSpecsRequest struct {
	// / Specs to be formatted
	Specs            []string `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FormatSpecsRequest) Reset()         { *m = FormatSpecsRequest{} }
func (m *FormatSpecsRequest) String() string { return proto.CompactTextString(m) }
func (*FormatSpecsRequest) ProtoMessage()    {}

func (m *FormatSpecsRequest) GetSpecs() []string {
	if m != nil {
		return m.Specs
	}
	return nil
}

// / Response on formatting spec files
type FormatSpecsResponse struct {
	// / Errors occurred on formatting
	Errors []string `protobuf:"bytes,1,rep,name=errors" json:"errors,omitempty"`
	// / Warnings occurred on formatting
	Warnings         []string `protobuf:"bytes,2,rep,name=warnings" json:"warnings,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FormatSpecsResponse) Reset()         { *m = FormatSpecsResponse{} }
func (m *FormatSpecsResponse) String() string { return proto.CompactTextString(m) }
func (*FormatSpecsResponse) ProtoMessage()    {}

func (m *FormatSpecsResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *FormatSpecsResponse) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

// / Response when a API message request is not supported.
type UnsupportedApiMessageResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *UnsupportedApiMessageResponse) Reset()         { *m = UnsupportedApiMessageResponse{} }
func (m *UnsupportedApiMessageResponse) String() string { return proto.CompactTextString(m) }
func (*UnsupportedApiMessageResponse) ProtoMessage()    {}

// / A generic message composing of all possible operations.
// / One of the Request/Response fields will have value, depending on the MessageType set.
type APIMessage struct {
	// / Type of API call being made
	MessageType *APIMessage_APIMessageType `protobuf:"varint,1,req,name=messageType,enum=gauge.messages.APIMessage_APIMessageType" json:"messageType,omitempty"`
	// / A unique id to represent this message. A response to the message should copy over this value.
	// / This is used to synchronize messages & responses
	MessageId *int64 `protobuf:"varint,2,req,name=messageId" json:"messageId,omitempty"`
	// / [GetProjectRootRequest](#gauge.messages.GetProjectRootRequest)
	ProjectRootRequest *GetProjectRootRequest `protobuf:"bytes,3,opt,name=projectRootRequest" json:"projectRootRequest,omitempty"`
	// / [GetProjectRootResponse](#gauge.messages.GetProjectRootResponse)
	ProjectRootResponse *GetProjectRootResponse `protobuf:"bytes,4,opt,name=projectRootResponse" json:"projectRootResponse,omitempty"`
	// / [GetInstallationRootRequest](#gauge.messages.GetInstallationRootRequest)
	InstallationRootRequest *GetInstallationRootRequest `protobuf:"bytes,5,opt,name=installationRootRequest" json:"installationRootRequest,omitempty"`
	// / [GetInstallationRootResponse](#gauge.messages.GetInstallationRootResponse)
	InstallationRootResponse *GetInstallationRootResponse `protobuf:"bytes,6,opt,name=installationRootResponse" json:"installationRootResponse,omitempty"`
	// / [GetAllStepsRequest](#gauge.messages.GetAllStepsRequest)
	AllStepsRequest *GetAllStepsRequest `protobuf:"bytes,7,opt,name=allStepsRequest" json:"allStepsRequest,omitempty"`
	// / [GetAllStepsResponse](#gauge.messages.GetAllStepsResponse)
	AllStepsResponse *GetAllStepsResponse `protobuf:"bytes,8,opt,name=allStepsResponse" json:"allStepsResponse,omitempty"`
	// / [GetAllSpecsRequest](#gauge.messages.GetAllSpecsRequest)
	AllSpecsRequest *GetAllSpecsRequest `protobuf:"bytes,9,opt,name=allSpecsRequest" json:"allSpecsRequest,omitempty"`
	// / [GetAllSpecsResponse](#gauge.messages.GetAllSpecsResponse)
	AllSpecsResponse *GetAllSpecsResponse `protobuf:"bytes,10,opt,name=allSpecsResponse" json:"allSpecsResponse,omitempty"`
	// / [GetStepValueRequest](#gauge.messages.GetStepValueRequest)
	StepValueRequest *GetStepValueRequest `protobuf:"bytes,11,opt,name=stepValueRequest" json:"stepValueRequest,omitempty"`
	// / [GetStepValueResponse](#gauge.messages.GetStepValueResponse)
	StepValueResponse *GetStepValueResponse `protobuf:"bytes,12,opt,name=stepValueResponse" json:"stepValueResponse,omitempty"`
	// / [GetLanguagePluginLibPathRequest](#gauge.messages.GetLanguagePluginLibPathRequest)
	LibPathRequest *GetLanguagePluginLibPathRequest `protobuf:"bytes,13,opt,name=libPathRequest" json:"libPathRequest,omitempty"`
	// / [GetLanguagePluginLibPathResponse](#gauge.messages.GetLanguagePluginLibPathResponse)
	LibPathResponse *GetLanguagePluginLibPathResponse `protobuf:"bytes,14,opt,name=libPathResponse" json:"libPathResponse,omitempty"`
	// / [ErrorResponse](#gauge.messages.ErrorResponse)
	Error *ErrorResponse `protobuf:"bytes,15,opt,name=error" json:"error,omitempty"`
	// / [GetAllConceptsRequest](#gauge.messages.GetAllConceptsRequest)
	AllConceptsRequest *GetAllConceptsRequest `protobuf:"bytes,16,opt,name=allConceptsRequest" json:"allConceptsRequest,omitempty"`
	// / [GetAllConceptsResponse](#gauge.messages.GetAllConceptsResponse)
	AllConceptsResponse *GetAllConceptsResponse `protobuf:"bytes,17,opt,name=allConceptsResponse" json:"allConceptsResponse,omitempty"`
	// / [PerformRefactoringRequest](#gauge.messages.PerformRefactoringRequest)
	PerformRefactoringRequest *PerformRefactoringRequest `protobuf:"bytes,18,opt,name=performRefactoringRequest" json:"performRefactoringRequest,omitempty"`
	// / [PerformRefactoringResponse](#gauge.messages.PerformRefactoringResponse)
	PerformRefactoringResponse *PerformRefactoringResponse `protobuf:"bytes,19,opt,name=performRefactoringResponse" json:"performRefactoringResponse,omitempty"`
	// / [ExtractConceptRequest](#gauge.messages.ExtractConceptRequest)
	ExtractConceptRequest *ExtractConceptRequest `protobuf:"bytes,20,opt,name=extractConceptRequest" json:"extractConceptRequest,omitempty"`
	// / [ExtractConceptResponse](#gauge.messages.ExtractConceptResponse)
	ExtractConceptResponse *ExtractConceptResponse `protobuf:"bytes,21,opt,name=extractConceptResponse" json:"extractConceptResponse,omitempty"`
	// / [FormatSpecsRequest] (#gauge.messages.FormatSpecsRequest)
	FormatSpecsRequest *FormatSpecsRequest `protobuf:"bytes,22,opt,name=formatSpecsRequest" json:"formatSpecsRequest,omitempty"`
	// / [FormatSpecsResponse] (#gauge.messages.FormatSpecsResponse)
	FormatSpecsResponse *FormatSpecsResponse `protobuf:"bytes,23,opt,name=formatSpecsResponse" json:"formatSpecsResponse,omitempty"`
	// / [UnsupportedApiMessageResponse] (#gauge.messages.UnsupportedApiMessageResponse)
	UnsupportedApiMessageResponse *UnsupportedApiMessageResponse `protobuf:"bytes,24,opt,name=unsupportedApiMessageResponse" json:"unsupportedApiMessageResponse,omitempty"`
	XXX_unrecognized              []byte                         `json:"-"`
}

func (m *APIMessage) Reset()         { *m = APIMessage{} }
func (m *APIMessage) String() string { return proto.CompactTextString(m) }
func (*APIMessage) ProtoMessage()    {}

func (m *APIMessage) GetMessageType() APIMessage_APIMessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return APIMessage_GetProjectRootRequest
}

func (m *APIMessage) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *APIMessage) GetProjectRootRequest() *GetProjectRootRequest {
	if m != nil {
		return m.ProjectRootRequest
	}
	return nil
}

func (m *APIMessage) GetProjectRootResponse() *GetProjectRootResponse {
	if m != nil {
		return m.ProjectRootResponse
	}
	return nil
}

func (m *APIMessage) GetInstallationRootRequest() *GetInstallationRootRequest {
	if m != nil {
		return m.InstallationRootRequest
	}
	return nil
}

func (m *APIMessage) GetInstallationRootResponse() *GetInstallationRootResponse {
	if m != nil {
		return m.InstallationRootResponse
	}
	return nil
}

func (m *APIMessage) GetAllStepsRequest() *GetAllStepsRequest {
	if m != nil {
		return m.AllStepsRequest
	}
	return nil
}

func (m *APIMessage) GetAllStepsResponse() *GetAllStepsResponse {
	if m != nil {
		return m.AllStepsResponse
	}
	return nil
}

func (m *APIMessage) GetAllSpecsRequest() *GetAllSpecsRequest {
	if m != nil {
		return m.AllSpecsRequest
	}
	return nil
}

func (m *APIMessage) GetAllSpecsResponse() *GetAllSpecsResponse {
	if m != nil {
		return m.AllSpecsResponse
	}
	return nil
}

func (m *APIMessage) GetStepValueRequest() *GetStepValueRequest {
	if m != nil {
		return m.StepValueRequest
	}
	return nil
}

func (m *APIMessage) GetStepValueResponse() *GetStepValueResponse {
	if m != nil {
		return m.StepValueResponse
	}
	return nil
}

func (m *APIMessage) GetLibPathRequest() *GetLanguagePluginLibPathRequest {
	if m != nil {
		return m.LibPathRequest
	}
	return nil
}

func (m *APIMessage) GetLibPathResponse() *GetLanguagePluginLibPathResponse {
	if m != nil {
		return m.LibPathResponse
	}
	return nil
}

func (m *APIMessage) GetError() *ErrorResponse {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *APIMessage) GetAllConceptsRequest() *GetAllConceptsRequest {
	if m != nil {
		return m.AllConceptsRequest
	}
	return nil
}

func (m *APIMessage) GetAllConceptsResponse() *GetAllConceptsResponse {
	if m != nil {
		return m.AllConceptsResponse
	}
	return nil
}

func (m *APIMessage) GetPerformRefactoringRequest() *PerformRefactoringRequest {
	if m != nil {
		return m.PerformRefactoringRequest
	}
	return nil
}

func (m *APIMessage) GetPerformRefactoringResponse() *PerformRefactoringResponse {
	if m != nil {
		return m.PerformRefactoringResponse
	}
	return nil
}

func (m *APIMessage) GetExtractConceptRequest() *ExtractConceptRequest {
	if m != nil {
		return m.ExtractConceptRequest
	}
	return nil
}

func (m *APIMessage) GetExtractConceptResponse() *ExtractConceptResponse {
	if m != nil {
		return m.ExtractConceptResponse
	}
	return nil
}

func (m *APIMessage) GetFormatSpecsRequest() *FormatSpecsRequest {
	if m != nil {
		return m.FormatSpecsRequest
	}
	return nil
}

func (m *APIMessage) GetFormatSpecsResponse() *FormatSpecsResponse {
	if m != nil {
		return m.FormatSpecsResponse
	}
	return nil
}

func (m *APIMessage) GetUnsupportedApiMessageResponse() *UnsupportedApiMessageResponse {
	if m != nil {
		return m.UnsupportedApiMessageResponse
	}
	return nil
}

func init() {
	proto.RegisterEnum("gauge.messages.APIMessage_APIMessageType", APIMessage_APIMessageType_name, APIMessage_APIMessageType_value)
}
