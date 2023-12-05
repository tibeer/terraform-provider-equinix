package metal_device

import (
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    "github.com/packethost/packngo"
)

type MetalDeviceResourceModel struct {
    ID types.String `tfsdk:"id"`
    Metro types.String `tfsdk:"metro"`
    Project_id types.String `tfsdk:"project_id"`
    Locked types.Bool   `tfsdk:"locked"`
}

func (rm *MetalDeviceResourceModel) parse(device *packngo.Device) diag.Diagnostics {
    var diags diag.Diagnostics

    // Populate the model fields from the device
    rm.Metro = types.StringValue(device.Metro.ID)
    rm.Project_id = types.StringValue(device.Project.ID)
    rm.Locked = types.BoolValue(device.Locked)

    return diags
}
