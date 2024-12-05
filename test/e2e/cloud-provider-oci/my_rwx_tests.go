package e2e

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"

	. "github.com/onsi/ginkgo"
	csi_util "github.com/oracle/oci-cloud-controller-manager/pkg/csi-util"
	"github.com/oracle/oci-cloud-controller-manager/test/e2e/framework"
)

var _ = Describe("CSI RWX Raw Block Volume Creation - Immediate Volume Binding", func() {
	f := framework.NewDefaultFramework("csi-immediate")
	Context("[cloudprovider][storage][csi][raw-block][rwx]", func() {
		It("Create PVC without pod and wait to be bound.", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-immediate-bind")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "Immediate", true, "Delete", nil)
			pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimBound)
			err := f.DeleteStorageClass(f.Namespace.Name)
			if err != nil {
				Fail(fmt.Sprintf("deleting storage class failed %s", f.Namespace.Name))
			}
		})
	})
})

var _ = Describe("CSI RWX Raw Block Volume Creation", func() {
	f := framework.NewDefaultFramework("csi-basic")
	Context("[cloudprovider][storage][csi][system-tags][raw-block][rwx]", func() {
		It("Create RWX raw block PVC and POD for CSI.", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests")
			ctx := context.TODO()
			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app1", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)
			volumeName := pvcJig.GetVolumeNameFromPVC(pvc.GetName(), f.Namespace.Name)
			compartmentId := f.GetCompartmentId(*setupF)
			// read created BV
			volumes, err := f.Client.BlockStorage().GetVolumesByName(ctx, volumeName, compartmentId)
			framework.ExpectNoError(err)
			// volume name duplicate should not exist
			for _, volume := range volumes {
				framework.Logf("volume details %v :", volume)
				//framework.Logf("cluster ocid from setup is %s", setupF.ClusterOcid)
				if setupF.AddOkeSystemTags && !framework.HasOkeSystemTags(volume.SystemTags) {
					framework.Failf("the resource %s is expected to have oke system tags", *volume.Id)
				}
			}
		})

		It("Create RWX raw block PVC with VolumeSize 1Gi but should use default 50Gi", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-1gi")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.VolumeFss, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app2", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("50Gi", pvc.Name, f.Namespace.Name)
		})

		It("Create RWX raw block PVC with VolumeSize 100Gi should use 100Gi", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-100gi")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MaxVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app3", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("100Gi", pvc.Name, f.Namespace.Name)
		})

		It("Data should persist on CSI RWX raw block volume on pod restart", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-pod-restart-data-persistence")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.CheckDataPersistenceForRawBlockVolumeWithDeployment(pvc.Name, f.Namespace.Name)
		})
	})
})

var _ = Describe("CSI RWX Raw Block Volume MULTI_NODE", func() {
	f := framework.NewDefaultFramework("csi-basic")
	Context("[cloudprovider][storage][csi][raw-block][rwx]", func() {
		It("Create RWX raw block PVC, use 2 pods", func() { // TODO: force them to schedule on different nodes
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)

			pvcJig.NewPodForCSI("app1", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			pvcJig.NewPodForCSI("app1-2", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

		})

		It("Create RWX raw block PVC with VolumeSize 1Gi but should use default 50Gi", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-1gi")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.VolumeFss, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app2", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("50Gi", pvc.Name, f.Namespace.Name)
		})

		It("Create RWX raw block PVC with VolumeSize 100Gi should use 100Gi", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-100gi")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MaxVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.NewPodForCSI("app3", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("100Gi", pvc.Name, f.Namespace.Name)
		})

		It("Data should persist on CSI RWX raw block volume on pod restart", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-pod-restart-data-persistence")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			pvcJig.CheckDataPersistenceForRawBlockVolumeWithDeployment(pvc.Name, f.Namespace.Name)
		})
	})
})

var _ = Describe("CSI Raw Block Volume Expansion iSCSI", func() {
	f := framework.NewDefaultFramework("csi-expansion")
	Context("[cloudprovider][storage][csi][expand][iSCSI][raw-block][rwx]", func() {
		It("Expand raw block PVC VolumeSize from 50Gi to 100Gi and asserts size, file existence and file corruptions for iSCSI volumes with existing storage class", func() {
			var size = "100Gi"
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-resizer-pvc-expand-to-100gi-iscsi")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeISCSI},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			podName := pvcJig.NewPodForCSI("expanded-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			expandedPvc := pvcJig.UpdateAndAwaitPVCOrFailCSI(pvc, pvc.Namespace, size, nil)

			time.Sleep(120 * time.Second) //waiting for expanded pvc to be functional

			pvcJig.CheckVolumeCapacity("100Gi", expandedPvc.Name, f.Namespace.Name)
			pvcJig.CheckFileExists(f.Namespace.Name, podName, "/dev", "xvda")
			pvcJig.CheckExpandedRawBlockVolumeReadWrite(f.Namespace.Name, podName)
			pvcJig.ExtractDataFromBlockDevice(f.Namespace.Name, podName, "/dev/xvda", "/tmp/testdata.txt")
			// pvcJig.CheckFileCorruption(f.Namespace.Name, podName, "/tmp", "testdata.txt")

			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
		})
	})
})

var _ = Describe("CSI Raw Block Volume Expansion iSCSI", func() {
	f := framework.NewDefaultFramework("csi-expansion")
	Context("[cloudprovider][storage][csi][expand][iSCSI][raw-block][rwx]", func() {
		It("Expand raw block PVC VolumeSize from 50Gi to 100Gi and asserts size, file existence and file corruptions for iSCSI volumes with new storage class", func() {
			var size = "100Gi"
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-resizer-pvc-expand-to-100gi-iscsi")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeISCSI},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			podName := pvcJig.NewPodForCSI("expanded-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			expandedPvc := pvcJig.UpdateAndAwaitPVCOrFailCSI(pvc, pvc.Namespace, size, nil)

			time.Sleep(120 * time.Second) //waiting for expanded pvc to be functional

			pvcJig.CheckVolumeCapacity("100Gi", expandedPvc.Name, f.Namespace.Name)
			pvcJig.CheckFileExists(f.Namespace.Name, podName, "/dev", "xvda")
			pvcJig.CheckExpandedRawBlockVolumeReadWrite(f.Namespace.Name, podName)
			pvcJig.ExtractDataFromBlockDevice(f.Namespace.Name, podName, "/dev/xvda", "/tmp/testdata.txt")
			// pvcJig.CheckFileCorruption(f.Namespace.Name, podName, "/tmp", "testdata.txt")

			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
	})
})

var _ = Describe("CSI Raw Block Volume Expansion Paravirtualized", func() {
	f := framework.NewDefaultFramework("csi-expansion")
	Context("[cloudprovider][storage][csi][expand][paravirtualized][raw-block][rwx]", func() {
		It("Expand raw block PVC VolumeSize from 50Gi to 100Gi and asserts size, file existence and file corruptions for paravirtualized volumes with new storage class", func() {
			var size = "100Gi"
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-resizer-pvc-expand-to-100gi-paravirtualized")

			scParameter := map[string]string{
				framework.KmsKey:         setupF.CMEKKMSKey,
				framework.AttachmentType: framework.AttachmentTypeParavirtualized,
			}
			scName := f.CreateStorageClassOrFail(f.Namespace.Name,
				"blockvolume.csi.oraclecloud.com", scParameter, pvcJig.Labels,
				"WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			podName := pvcJig.NewPodForCSI("expanded-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			expandedPvc := pvcJig.UpdateAndAwaitPVCOrFailCSI(pvc, pvc.Namespace, size, nil)

			time.Sleep(120 * time.Second) //waiting for expanded pvc to be functional

			pvcJig.CheckVolumeCapacity("100Gi", expandedPvc.Name, f.Namespace.Name)
			pvcJig.CheckFileExists(f.Namespace.Name, podName, "/dev", "xvda")
			pvcJig.CheckExpandedRawBlockVolumeReadWrite(f.Namespace.Name, podName)
			pvcJig.ExtractDataFromBlockDevice(f.Namespace.Name, podName, "/dev/xvda", "/tmp/testdata.txt")
			// pvcJig.CheckFileCorruption(f.Namespace.Name, podName, "/tmp", "testdata.txt")

			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
	})
})

var _ = Describe("CSI Raw Block Volume Performance Level", func() {
	f := framework.NewBackupFramework("csi-perf-level")
	Context("[cloudprovider][storage][csi][perf][iSCSI][raw-block][rwx]", func() {
		It("Create CSI raw block volume with Performance Level as Low Cost", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-perf-iscsi-lowcost")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeISCSI, csi_util.VpusPerGB: "0"},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			pvcJig.NewPodForCSI("low-cost-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumePerformanceLevel(f.BlockStorageClient, pvc.Namespace, pvc.Name, csi_util.LowCostPerformanceOption)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
		It("Create CSI raw block volume with no Performance Level and verify default", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-perf-iscsi-default")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com", nil, pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			pvcJig.NewPodForCSI("default-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumePerformanceLevel(f.BlockStorageClient, pvc.Namespace, pvc.Name, csi_util.BalancedPerformanceOption)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
		})
		It("Create CSI raw block volume with Performance Level as High", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-perf-iscsi-high")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeISCSI, csi_util.VpusPerGB: "20"},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			podName := pvcJig.NewPodForCSI("high-perf-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running
			pvcJig.CheckVolumePerformanceLevel(f.BlockStorageClient, pvc.Namespace, pvc.Name, csi_util.HigherPerformanceOption)
			pvcJig.CheckISCSIQueueDepthOnNode(f.Namespace.Name, podName)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
	})
	Context("[cloudprovider][storage][csi][perf][paravirtualized][raw-block][rwx]", func() {
		It("Create CSI raw block volume with Performance Level as Low Cost", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-perf-paravirtual-lowcost")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeParavirtualized, csi_util.VpusPerGB: "0"},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			pvcJig.NewPodForCSI("low-cost-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumePerformanceLevel(f.BlockStorageClient, pvc.Namespace, pvc.Name, csi_util.LowCostPerformanceOption)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
		It("Create CSI raw block volume with no Performance Level and verify default", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-perf-paravirtual-balanced")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeParavirtualized, csi_util.VpusPerGB: "10"},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			pvcJig.NewPodForCSI("default-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumePerformanceLevel(f.BlockStorageClient, pvc.Namespace, pvc.Name, csi_util.BalancedPerformanceOption)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
		It("Create CSI raw block volume with Performance Level as High", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-perf-paravirtual-high")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeParavirtualized, csi_util.VpusPerGB: "20"},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)
			pvc := pvcJig.CreateAndAwaitPVCOrFailCSI(f.Namespace.Name, framework.MinVolumeBlock, scName, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			pvcJig.NewPodForCSI("high-perf-pvc-app", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running
			pvcJig.CheckVolumePerformanceLevel(f.BlockStorageClient, pvc.Namespace, pvc.Name, csi_util.HigherPerformanceOption)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
	})
	Context("[cloudprovider][storage][csi][perf][static][raw-block]", func() {
		It("High Performance CSI raw block volume Static Provisioning", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-perf-static-high")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				map[string]string{framework.AttachmentType: framework.AttachmentTypeISCSI, csi_util.VpusPerGB: "20"},
				pvcJig.Labels, "WaitForFirstConsumer", true, "Delete", nil)

			compartmentId := ""
			if setupF.Compartment1 != "" {
				compartmentId = setupF.Compartment1
			} else if f.CloudProviderConfig.CompartmentID != "" {
				compartmentId = f.CloudProviderConfig.CompartmentID
			} else if f.CloudProviderConfig.Auth.CompartmentID != "" {
				compartmentId = f.CloudProviderConfig.Auth.CompartmentID
			} else {
				framework.Failf("Compartment Id undefined.")
			}
			pvc, volumeId := pvcJig.CreateAndAwaitStaticPVCOrFailCSI(f.BlockStorageClient, f.Namespace.Name, framework.MinVolumeBlock, csi_util.HigherPerformanceOption, scName, setupF.AdLocation, compartmentId, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			f.VolumeIds = append(f.VolumeIds, pvc.Spec.VolumeName)
			podName := pvcJig.NewPodForCSI("app4", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcJig.CheckVolumeCapacity("50Gi", pvc.Name, f.Namespace.Name)
			pvcJig.CheckISCSIQueueDepthOnNode(pvc.Namespace, podName)
			f.VolumeIds = append(f.VolumeIds, volumeId)
			_ = f.DeleteStorageClass(f.Namespace.Name)
		})
	})
})

var _ = Describe("CSI Static Raw Block Volume Creation", func() {
	f := framework.NewBackupFramework("csi-static")
	Context("[cloudprovider][storage][csi][static][raw-block][rwx]", func() {
		It("Static Provisioning of a raw block CSI", func() {
			pvcJig := framework.NewPVCTestJig(f.ClientSet, "csi-provisioner-e2e-tests-pvc-with-static")

			scName := f.CreateStorageClassOrFail(f.Namespace.Name, "blockvolume.csi.oraclecloud.com",
				nil, pvcJig.Labels, "WaitForFirstConsumer", false, "Delete", nil)

			compartmentId := ""
			if setupF.Compartment1 != "" {
				compartmentId = setupF.Compartment1
			} else if f.CloudProviderConfig.CompartmentID != "" {
				compartmentId = f.CloudProviderConfig.CompartmentID
			} else if f.CloudProviderConfig.Auth.CompartmentID != "" {
				compartmentId = f.CloudProviderConfig.Auth.CompartmentID
			} else {
				framework.Failf("Compartment Id undefined.")
			}
			pvc, volumeId := pvcJig.CreateAndAwaitStaticPVCOrFailCSI(f.BlockStorageClient, f.Namespace.Name, framework.MinVolumeBlock, 10, scName, setupF.AdLocation, compartmentId, nil, v1.PersistentVolumeBlock, v1.ReadWriteMany, v1.ClaimPending)
			pvcJig.NewPodForCSI("app4", f.Namespace.Name, pvc.Name, setupF.AdLabel, v1.PersistentVolumeBlock)

			time.Sleep(60 * time.Second) //waiting for pod to up and running

			pvcObj := pvcJig.GetPVCByName(pvc.Name, f.Namespace.Name)
			f.VolumeIds = append(f.VolumeIds, pvcObj.Spec.VolumeName)
			pvcJig.CheckVolumeCapacity("50Gi", pvc.Name, f.Namespace.Name)
			f.VolumeIds = append(f.VolumeIds, volumeId)
		})
	})
})
