package vlans

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/packethost/packngo"
)

type DataSourceModel struct {
	ProjectID   types.String `tfsdk:"project_id"`
	VlanID      types.String `tfsdk:"vlan_id"`
	Vxlan       types.Int64  `tfsdk:"vxlan"`
	Facility    types.String `tfsdk:"facility"`
	Metro       types.String `tfsdk:"metro"`
	Description types.String `tfsdk:"description"`
}

func (m *DataSourceModel) parse(vlan *packngo.VirtualNetwork) diag.Diagnostics {
	m.ProjectID = types.StringValue(vlan.Project.ID)
	m.VlanID = types.StringValue(vlan.ID)
	m.Facility = types.StringValue(vlan.FacilityCode)
	m.Metro = types.StringValue(vlan.MetroCode)
	m.Description = types.StringValue(vlan.Description)
	m.Vxlan = types.Int64Value(int64(vlan.VXLAN))
	return nil
}

type ResourceModel struct {
	ID          types.String `tfsdk:"id"`
	ProjectID   types.String `tfsdk:"project_id"`
	Vxlan       types.Int64  `tfsdk:"vxlan"`
	Facility    types.String `tfsdk:"facility"`
	Metro       types.String `tfsdk:"metro"`
	Description types.String `tfsdk:"description"`
}

func (m *ResourceModel) parse(vlan *packngo.VirtualNetwork) diag.Diagnostics {
	m.ID = types.StringValue(vlan.ID)
	m.ProjectID = types.StringValue(vlan.Project.ID)
	m.Facility = types.StringValue(vlan.FacilityCode)
	m.Metro = types.StringValue(vlan.MetroCode)
	m.Description = types.StringValue(vlan.Description)
	m.Vxlan = types.Int64Value(int64(vlan.VXLAN))
	return nil
}
