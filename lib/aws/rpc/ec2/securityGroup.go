// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for SecurityGroup resource provider */

// SecurityGroupToken is the type token corresponding to the SecurityGroup package type.
const SecurityGroupToken = tokens.Type("aws:ec2/securityGroup:SecurityGroup")

// SecurityGroupProviderOps is a pluggable interface for SecurityGroup-related management functionality.
type SecurityGroupProviderOps interface {
    Check(ctx context.Context, obj *SecurityGroup) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *SecurityGroup) (resource.ID, *SecurityGroupOuts, error)
    Get(ctx context.Context, id resource.ID) (*SecurityGroup, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *SecurityGroup, new *SecurityGroup, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *SecurityGroup, new *SecurityGroup, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// SecurityGroupProvider is a dynamic gRPC-based plugin for managing SecurityGroup resources.
type SecurityGroupProvider struct {
    ops SecurityGroupProviderOps
}

// NewSecurityGroupProvider allocates a resource provider that delegates to a ops instance.
func NewSecurityGroupProvider(ops SecurityGroupProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &SecurityGroupProvider{ops: ops}
}

func (p *SecurityGroupProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *SecurityGroupProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *SecurityGroupProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, outs, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   string(id),
        Outputs: resource.MarshalProperties(
            nil, resource.NewPropertyMap(outs), resource.MarshalOptions{},
        ),
    }, nil
}

func (p *SecurityGroupProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *SecurityGroupProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("groupDescription") {
            replaces = append(replaces, "groupDescription")
        }
        if diff.Changed("vpc") {
            replaces = append(replaces, "vpc")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *SecurityGroupProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(SecurityGroupToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *SecurityGroupProvider) Unmarshal(
    v *pbstruct.Struct) (*SecurityGroup, resource.PropertyMap, mapper.DecodeError) {
    var obj SecurityGroup
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable SecurityGroup structure(s) */

// SecurityGroup is a marshalable representation of its corresponding IDL type.
type SecurityGroup struct {
    Name string `json:"name"`
    GroupDescription string `json:"groupDescription"`
    VPC *resource.ID `json:"vpc,omitempty"`
    SecurityGroupEgress *[]SecurityGroupRule `json:"securityGroupEgress,omitempty"`
    SecurityGroupIngress *[]SecurityGroupRule `json:"securityGroupIngress,omitempty"`
    GroupID string `json:"groupID,omitempty"`
}

// SecurityGroupOuts is a marshalable representation of its IDL type's output properties.
type SecurityGroupOuts struct {
    GroupID string `json:"groupID"`
}

// SecurityGroup's properties have constants to make dealing with diffs and property bags easier.
const (
    SecurityGroup_Name = "name"
    SecurityGroup_GroupDescription = "groupDescription"
    SecurityGroup_VPC = "vpc"
    SecurityGroup_SecurityGroupEgress = "securityGroupEgress"
    SecurityGroup_SecurityGroupIngress = "securityGroupIngress"
    SecurityGroup_GroupID = "groupID"
)

/* Marshalable SecurityGroupRule structure(s) */

// SecurityGroupRule is a marshalable representation of its corresponding IDL type.
type SecurityGroupRule struct {
    IPProtocol string `json:"ipProtocol"`
    CIDRIP *string `json:"cidrIp,omitempty"`
    FromPort *float64 `json:"fromPort,omitempty"`
    ToPort *float64 `json:"toPort,omitempty"`
}

// SecurityGroupRule's properties have constants to make dealing with diffs and property bags easier.
const (
    SecurityGroupRule_IPProtocol = "ipProtocol"
    SecurityGroupRule_CIDRIP = "cidrIp"
    SecurityGroupRule_FromPort = "fromPort"
    SecurityGroupRule_ToPort = "toPort"
)

