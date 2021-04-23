package metrics

const (
	// LBProvisionFailure is the OCI metric suffix for
	// LB provision
	LBProvision = "LB_PROVISION"
	// LBUpdateFailure is the OCI metric suffix for
	// LB update
	LBUpdate = "LB_UPDATE"
	// LBDeleteFailure is the OCI metric suffix for
	// LB delete
	LBDelete = "LB_DELETE"
)

const (
	// PVProvisionFailure is the metric for PV provision failures
	PVProvisionFailure = "PV_PROVISION_FAILURE"
	// PVAttachFailure is the metric for PV attach failures
	PVAttachFailure = "PV_ATTACH_FAILURE"
	// PVDetachFailure is the metric for PV detach failure
	PVDetachFailure = "PV_DETACH_FAILURE"
	// PVDeleteFailure is the metric for PV delete failure
	PVDeleteFailure = "PV_DELETE_FAILURE"

	// PVProvisionSuccess is the metric used to track the time
	// taken for the provision operation
	PVProvisionSuccess = "PV_PROVISION_SUCCESS"
	// PVAttachSuccess is the metric used to track the time
	// taken for the provision operation
	PVAttachSuccess = "PV_ATTACH_SUCCESS"
	// PVDetachSuccess is the metric used to track the time
	// taken for the provision operation
	PVDetachSuccess = "PV_DETACH_SUCCESS"
	// PVDeleteSuccess is the metric used to track the time
	// taken for the provision operation
	PVDeleteSuccess = "PV_DELETE_SUCCESS"
)
